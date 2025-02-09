package utils

import (
	"context"

	"github.com/tommjj/ql-kho-lua/internal/config"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

// AutoCreateRootUser is a helper func auto create a root user if user table is empty
func AutoCreateRootUser(userService ports.IUserService, user *config.DefaultRootUser) error {
	ctx := context.Background()

	count, err := userService.CountUsers(ctx, "")
	if err != nil {
		return err
	}

	if count != 0 {
		return nil
	}

	_, err = userService.CreateUser(ctx, &domain.User{
		Name:     user.Name,
		Phone:    user.Phone,
		Password: user.Password,
		Email:    user.Email,
		Role:     domain.Root,
	})

	return err
}
