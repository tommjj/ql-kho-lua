package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tommjj/ql-kho-lua/internal/config"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
	"github.com/tommjj/ql-kho-lua/internal/core/utils"
)

var jwtMethod *jwt.SigningMethodHMAC = jwt.SigningMethodHS256

type CustomClaims struct {
	ID    int         `json:"id"`
	Name  string      `json:"name"`
	Email string      `json:"email"`
	Role  domain.Role `json:"role"`
	Key   string      `json:"key"`
	jwt.RegisteredClaims
}

type JWTService struct {
	key      []byte
	keyFunc  func(token *jwt.Token) (interface{}, error)
	duration time.Duration
	keyRepo  ports.IKeyRepository
}

func NewJWTTokenService(conf config.Auth, keyRepo ports.IKeyRepository) ports.ITokenService {
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
		keyRepo:  keyRepo,
	}
}

func (j *JWTService) CreateToken(user *domain.User) (string, error) {
	key := utils.CreateRandomString(64)

	err := j.keyRepo.SetKey(context.Background(), user.ID, key)
	if err != nil {
		return "", err
	}

	claims := jwt.NewWithClaims(jwtMethod, CustomClaims{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		Key:   key,
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
		key, err := j.keyRepo.GetKey(context.Background(), claims.ID)
		if err != nil {
			return nil, domain.ErrInvalidToken
		}

		if key == "" {
			return nil, domain.ErrInvalidToken
		}

		if key != claims.Key {
			return nil, domain.ErrInvalidToken
		}

		return &domain.TokenPayload{
			ID:    claims.ID,
			Name:  claims.Name,
			Email: claims.Email,
			Role:  claims.Role,
			Key:   claims.Key,
		}, nil
	case errors.Is(err, jwt.ErrTokenMalformed) || errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return nil, domain.ErrInvalidToken
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		return nil, domain.ErrExpiredToken
	default:
		return nil, err
	}
}
