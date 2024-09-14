package services

import (
	"context"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
	"github.com/tommjj/ql-kho-lua/internal/core/utils"
)

type authService struct {
	userRepo     ports.IUserRepository
	tokenService ports.ITokenService
}

func NewAuthService(userRepo ports.IUserRepository, token ports.ITokenService) ports.IAuthService {
	return &authService{
		userRepo:     userRepo,
		tokenService: token,
	}
}

func (as *authService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := as.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	err = utils.ComparePassword(password, user.Password)
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	token, err := as.tokenService.CreateToken(user)
	if err != nil {
		return "", domain.ErrInternal
	}

	return token, nil
}
