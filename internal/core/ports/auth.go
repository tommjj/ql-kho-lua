package ports

import (
	"context"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

type IAuthService interface {
	Login(ctx context.Context, username, password string) (string, error)
}

type ITokenService interface {
	// CreateToken create an new token
	CreateToken(user *domain.User) (string, error)
	// VerifyToken verify string token
	VerifyToken(token string) (*domain.TokenPayload, error)
}
