package handlers

import "github.com/tommjj/ql-kho-lua/internal/core/ports"

type UserHandler struct {
	svc ports.IUserRepository
}

func NewUserHandler(userService ports.IUserRepository) *UserHandler {
	return &UserHandler{
		svc: userService,
	}
}
