package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	authModule "github.com/hifat/mallow-sale-api/internal/auth"
	"github.com/hifat/mallow-sale-api/pkg/config"
)

type TokenType int64

var (
	ErrInvalidToken      = errors.New("invalid token")
	ErrTokenExpired      = errors.New("token expired")
	ErrNotFoundTokenType = errors.New("not found token type")
)

const (
	REFRESH TokenType = iota
	ACCESS
)

const (
	RefreshToken = "refresh-token"
	AccessToken  = "access-token"
)

type AuthClaims struct {
	Passport authModule.Passport
	jwt.RegisteredClaims
}

func (t TokenType) name() (string, error) {
	tokenTypes := map[TokenType]string{
		REFRESH: "refresh-token",
		ACCESS:  "access-token",
	}
	if _, ok := tokenTypes[t]; !ok {
		return "", ErrNotFoundTokenType
	}

	return tokenTypes[t], nil
}

func (t TokenType) duration(cfg config.Auth) (time.Duration, error) {
	durations := map[TokenType]time.Duration{
		REFRESH: cfg.RefreshTokenExpires,
		ACCESS:  cfg.AccessTokenExpires,
	}
	if _, ok := durations[t]; !ok {
		return 0, ErrNotFoundTokenType
	}

	return durations[t], nil
}

func (t TokenType) secret(cfg config.Auth) (string, error) {
	secrets := map[TokenType]string{
		REFRESH: cfg.RefreshTokenSecret,
		ACCESS:  cfg.AccessTokenSecret,
	}
	if _, ok := secrets[t]; !ok {
		return "", ErrNotFoundTokenType
	}

	return secrets[t], nil
}

type handler struct {
	cfg      *config.Config
	passport authModule.Passport
}

func New(cfg *config.Config, passport authModule.Passport) *handler {
	return &handler{
		cfg,
		passport,
	}
}

func (h *handler) Signed(tokenType TokenType) (*AuthClaims, string, error) {
	subject, err := tokenType.name()
	if err != nil {
		return nil, "", err
	}

	duration, err := tokenType.duration(h.cfg.Auth)
	if err != nil {
		return nil, "", err
	}

	secret, err := tokenType.secret(h.cfg.Auth)
	if err != nil {
		return nil, "", err
	}

	authClaims := AuthClaims{
		Passport: h.passport,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    h.cfg.App.Name,
			Subject:   subject,
			Audience:  []string{"*"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims)

	signed, err := token.SignedString([]byte(secret))

	return &authClaims, signed, err
}

func Claims(cfg config.Auth, tokenType TokenType, tokenString string) (*AuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secret, err := tokenType.secret(cfg)
		if err != nil {
			return nil, err
		}

		return []byte(secret), nil
	})
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenMalformed), errors.Is(err, jwt.ErrTokenSignatureInvalid), errors.Is(err, jwt.ErrTokenNotValidYet):
			return nil, ErrInvalidToken
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, ErrTokenExpired
		default:
			return nil, err
		}
	}

	if claims, ok := token.Claims.(*AuthClaims); ok {
		return claims, nil
	}

	return nil, err
}
