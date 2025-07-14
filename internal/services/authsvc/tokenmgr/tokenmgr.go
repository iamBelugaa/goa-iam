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

type tokenType string

var (
	AccessToken  tokenType = "ACCESS_TOKEN"
	RefreshToken tokenType = "REFRESH_TOKEN"
)

type Claims struct {
	jwt.RegisteredClaims
	TokenType tokenType `json:"tokenType"`
}

type JWTTokenManager struct {
	cfg    *config.Auth
	parser *jwt.Parser
	method jwt.SigningMethod
}

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

func (tm *JWTTokenManager) StandardClaims(sub string, tokenType tokenType) Claims {
	if tokenType == AccessToken {
		return Claims{
			TokenType: AccessToken,
			RegisteredClaims: jwt.RegisteredClaims{
				ID:        uuid.New().String(),
				Subject:   sub,
				Issuer:    tm.cfg.Issuer,
				Audience:  jwt.ClaimStrings{tm.cfg.Audience},
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(tm.cfg.AccessTokenExpTime)),
				NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		}
	}

	return Claims{
		TokenType: RefreshToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			Subject:   sub,
			Issuer:    tm.cfg.Issuer,
			Audience:  jwt.ClaimStrings{tm.cfg.Audience},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tm.cfg.RefreshTokenExpTime)),
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}

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

func (tm *JWTTokenManager) ParseWithClaims(bearerToken string) (Claims, error) {
	parts := strings.Split(bearerToken, " ")

	if len(parts) != 2 || parts[1] != "Bearer" {
		return Claims{}, auth.MakeInvalidToken(fmt.Errorf("authorization header missing"))
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
