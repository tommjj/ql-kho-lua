package ports

import (
	"context"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

type IKeyRepository interface {
	// SetKey set the key for the user
	SetKey(ctx context.Context, id int, key string) error
	// GetKey get key by user id
	GetKey(ctx context.Context, id int) (string, error)
	// DelKey delete key
	DelKey(ctx context.Context, id int) error
}

type IAuthService interface {
	Login(ctx context.Context, email, password string) (string, error)
}

type ITokenService interface {
	// CreateToken create an new token
	CreateToken(user *domain.User) (string, error)
	// VerifyToken verify string token
	VerifyToken(token string) (*domain.TokenPayload, error)
}
