package middlewareHandler

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"golang.org/x/time/rate"
)

const (
	// Per-IP: 20 requests per second, burst of 40
	rateLimitPerIP  = 20
	burstPerIP      = 40
	// Global: 500 requests per second, burst of 1000
	rateLimitGlobal = 500
	burstGlobal     = 1000
	// Cleanup idle visitors after 3 minutes
	idleTimeout = 3 * time.Minute
)

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type rateLimiter struct {
	mu        sync.Mutex
	visitors  map[string]*visitor
	global    *rate.Limiter
}

func newRateLimiter() *rateLimiter {
	rl := &rateLimiter{
		visitors: make(map[string]*visitor),
		global:   rate.NewLimiter(rate.Limit(rateLimitGlobal), burstGlobal),
	}
	go rl.cleanupLoop()
	return rl
}

// getVisitor returns (or creates) the per-IP limiter for the given IP.
func (rl *rateLimiter) getVisitor(ip string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	v, exists := rl.visitors[ip]
	if !exists {
		limiter := rate.NewLimiter(rate.Limit(rateLimitPerIP), burstPerIP)
		rl.visitors[ip] = &visitor{limiter: limiter, lastSeen: time.Now()}
		return limiter
	}

	v.lastSeen = time.Now()
	return v.limiter
}

// cleanupLoop removes IP entries that have been idle longer than idleTimeout.
func (rl *rateLimiter) cleanupLoop() {
	ticker := time.NewTicker(idleTimeout)
	defer ticker.Stop()
	for range ticker.C {
		rl.mu.Lock()
		for ip, v := range rl.visitors {
			if time.Since(v.lastSeen) > idleTimeout {
				delete(rl.visitors, ip)
			}
		}
		rl.mu.Unlock()
	}
}

// RateLimit returns a Gin middleware that enforces per-IP and global rate limits.
func RateLimit() gin.HandlerFunc {
	rl := newRateLimiter()

	return func(c *gin.Context) {
		ip := c.ClientIP()

		// Check global limit first
		if !rl.global.Allow() {
			errRes := handling.ThrowErrByCode(define.CodeRateLimitExceeded)
			c.Header("Retry-After", "1")
			c.AbortWithStatusJSON(http.StatusTooManyRequests, errRes)
			return
		}

		// Check per-IP limit
		if !rl.getVisitor(ip).Allow() {
			errRes := handling.ThrowErrByCode(define.CodeRateLimitExceeded)
			c.Header("Retry-After", "1")
			c.AbortWithStatusJSON(http.StatusTooManyRequests, errRes)
			return
		}

		c.Next()
	}
}
