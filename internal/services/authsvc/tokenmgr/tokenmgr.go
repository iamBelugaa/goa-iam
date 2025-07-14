// Package tokenmgr provides JWT-based token management utilities for the IAM system.
package tokenmgr

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/iamBelugaa/goa-iam/gen/auth/gen/auth"
	"github.com/iamBelugaa/goa-iam/internal/config"
	"github.com/iamBelugaa/goa-iam/internal/domain/codes"
)

// tokenType defines a custom string type used to distinguish token purposes.
type tokenType string

// Supported token types.
var (
	AccessToken  tokenType = "ACCESS_TOKEN"
	RefreshToken tokenType = "REFRESH_TOKEN"
)

// Claims wraps jwt.RegisteredClaims and adds a custom TokenType field.
type Claims struct {
	jwt.RegisteredClaims
	TokenType tokenType `json:"tokenType"`
}

// JWTTokenManager is responsible for creating and validating JWT tokens
// based on configuration such as issuer, audience, expiration, and secret.
type JWTTokenManager struct {
	cfg    *config.Auth      // Auth config containing secrets and expiration durations
	parser *jwt.Parser       // Configured JWT parser
	method jwt.SigningMethod // Signing method
}

// NewJWTManager creates a new JWTTokenManager with validation rules
// based on the given auth configuration.
func NewJWTManager(cfg *config.Auth) *JWTTokenManager {
	return &JWTTokenManager{
		cfg:    cfg,
		method: jwt.GetSigningMethod(jwt.SigningMethodHS256.Name),
		parser: jwt.NewParser(
			jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
			jwt.WithAudience(cfg.Audience),
			jwt.WithIssuer(cfg.Issuer),
			jwt.WithIssuedAt(),
			jwt.WithExpirationRequired(),
		),
	}
}

// StandardClaims creates a JWT Claims object for the given subject and token type.
// It sets fields like issuer, audience, issue time, expiration, etc.
func (tm *JWTTokenManager) StandardClaims(sub string, tokenType tokenType) Claims {
	expiration := tm.cfg.AccessTokenExpTime
	if tokenType == RefreshToken {
		expiration = tm.cfg.RefreshTokenExpTime
	}

	return Claims{
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			Subject:   sub,
			Issuer:    tm.cfg.Issuer,
			Audience:  jwt.ClaimStrings{tm.cfg.Audience},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}

// Generate signs the given claims and returns the corresponding JWT as a string.
func (tm *JWTTokenManager) Generate(claims Claims) (string, error) {
	token := jwt.NewWithClaims(tm.method, claims)

	signedToken, err := token.SignedString(tm.cfg.Secret)
	if err != nil {
		return "", &auth.InternalServerError{
			Message: "Failed to sign token string",
			Code:    auth.ErrorCode(codes.InternalServerErrCode),
		}
	}

	return signedToken, nil
}

// ParseWithClaims extracts and validates claims from the given Bearer token string.
func (tm *JWTTokenManager) ParseWithClaims(bearerToken string) (Claims, error) {
	parts := strings.Split(bearerToken, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		return Claims{}, auth.MakeInvalidToken(fmt.Errorf("authorization header missing or malformed"))
	}

	var claims Claims
	tokenStr := parts[1]

	if _, err := tm.parser.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (any, error) {
		if t.Method != tm.method {
			return "", auth.MakeInvalidToken(fmt.Errorf("invalid token signature"))
		}
		return tm.cfg.Secret, nil
	}); err != nil {
		return Claims{}, auth.MakeInvalidToken(fmt.Errorf("invalid token"))
	}

	return claims, nil
}
