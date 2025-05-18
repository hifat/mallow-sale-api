package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/hifat/mallow-sale-api/config"
)

type TokenType string

const AccessToken TokenType = "access_token"
const RefreshToken TokenType = "refresh_token"
const APIKey TokenType = "api_key"

var ErrInvalidToken = errors.New("invalid token")
var ErrExpired = errors.New("token expired")
var ErrInvalidPayload = errors.New("invalid token payload")

type IToken interface {
	Claims(tokenStr string, tokenType TokenType) (*Credentials, error)
}

type tokenAuth struct {
	cfg *config.Config
}

type Credentials struct {
	UserID string `json:"userID"`
	jwt.RegisteredClaims
}

func NewToken(cfg *config.Config) IToken {
	return &tokenAuth{cfg}
}

func (t *tokenAuth) getSecret(tokenType TokenType) string {
	secretMap := map[TokenType]string{
		AccessToken:  t.cfg.Auth.AccessToken,
		RefreshToken: t.cfg.Auth.RefreshToken,
		APIKey:       t.cfg.Auth.APIKey,
	}

	return secretMap[tokenType]
}

func (t *tokenAuth) getExpires(tokenType TokenType) time.Duration {
	secretMap := map[TokenType]time.Duration{
		AccessToken:  t.cfg.Auth.AccessTokenExpires,
		RefreshToken: t.cfg.Auth.RefreshTokenExpires,
	}

	return secretMap[tokenType]
}

// TODO: Change payload type to UserPrototype
func (t *tokenAuth) Sign(tokenType TokenType, payload *Credentials) (string, error) {
	var claims *Credentials
	if payload != nil {
		claims = &Credentials{
			UserID: payload.UserID,
		}
	}

	tokenDuration := t.getExpires(tokenType)
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ID:        uuid.NewString(),
		Issuer:    "mallow-sale",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenDuration)),
	}

	secret := t.getSecret(tokenType)
	token := jwt.NewWithClaims(jwt.SigningMethodES512, claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (t *tokenAuth) Claims(tokenStr string, tokenType TokenType) (*Credentials, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Credentials{}, func(jt *jwt.Token) (interface{}, error) {
		if _, ok := jt.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jt.Header["alg"])
		}

		return []byte(t.getSecret(tokenType)), nil
	})
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenMalformed), errors.Is(err, jwt.ErrTokenSignatureInvalid), errors.Is(err, jwt.ErrTokenNotValidYet):
			return nil, ErrInvalidToken
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, ErrExpired
		default:
			return nil, err
		}
	}

	claims, ok := token.Claims.(*Credentials)
	if !ok {
		return nil, ErrInvalidPayload
	}

	return claims, nil
}
