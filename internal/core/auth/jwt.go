package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tommjj/ql-kho-lua/internal/config"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

var jwtMethod *jwt.SigningMethodHMAC = jwt.SigningMethodHS256

type CustomClaims struct {
	ID    int         `json:"id"`
	Name  string      `json:"name"`
	Email string      `json:"email"`
	Role  domain.Role `json:"role"`
	jwt.RegisteredClaims
}

type JWTService struct {
	key      []byte
	keyFunc  func(token *jwt.Token) (interface{}, error)
	duration time.Duration
}

func NewJWTTokenService(conf config.Auth) ports.ITokenService {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.ErrInvalidToken
		}

		return []byte(conf.SecretKey), nil
	}

	return &JWTService{
		key:      []byte(conf.SecretKey),
		keyFunc:  keyFunc,
		duration: conf.Duration,
	}
}

func (j *JWTService) CreateToken(user *domain.User) (string, error) {
	claims := jwt.NewWithClaims(jwtMethod, CustomClaims{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "blog-api",
		},
	})

	str, err := claims.SignedString(j.key)
	if err != nil {
		return "", domain.ErrTokenCreation
	}

	return str, nil
}

func (j *JWTService) VerifyToken(tokenString string) (*domain.TokenPayload, error) {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, j.keyFunc)

	switch {
	case token.Valid:
		return &domain.TokenPayload{
			ID:    claims.ID,
			Name:  claims.Name,
			Email: claims.Email,
			Role:  claims.Role,
		}, nil
	case errors.Is(err, jwt.ErrTokenMalformed) || errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return nil, domain.ErrInvalidToken
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		return nil, domain.ErrExpiredToken
	default:
		return nil, err
	}
}
