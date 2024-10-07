package services

import (
	"context"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
	"github.com/tommjj/ql-kho-lua/internal/core/utils"
)

type userService struct {
	repo ports.IUserRepository
}

func NewUserService(userRepo ports.IUserRepository) *userService {
	return &userService{
		repo: userRepo,
	}
}

func (us *userService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	hashPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, domain.ErrInternal
	}

	newUser, err := us.repo.CreateUser(ctx, &domain.User{
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: hashPass,
		Role:     user.Role,
	})
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	newUser.RemovePass()

	return newUser, err
}

func (us *userService) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	var user *domain.User
	var err error

	user, err = us.repo.GetUserByID(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	user.RemovePass()

	return user, nil
}

func (us *userService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user *domain.User
	var err error

	user, err = us.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	user.RemovePass()

	return user, nil
}

func (us *userService) GetListUsers(ctx context.Context, q string, limit, skip int) ([]domain.User, error) {
	var user []domain.User
	var err error

	user, err = us.repo.GetListUsers(ctx, q, limit, skip)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return user, nil
}

func (us *userService) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	_, err := us.repo.GetUserByID(ctx, user.ID)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	var hashedPass string
	if user.Password != "" {
		hashedPass, err = utils.HashPassword(user.Password)
		if err != nil {
			return nil, domain.ErrInternal
		}
	}

	updatedUser, err := us.repo.UpdateUser(ctx, &domain.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: hashedPass,
		Role:     user.Role,
	})
	if err != nil {
		switch err {
		case domain.ErrConflictingData:
			return nil, err
		case domain.ErrNoUpdatedData:
			return nil, err
		default:
			return nil, domain.ErrInternal
		}
	}
	updatedUser.RemovePass()

	return updatedUser, err
}

func (us *userService) DeleteUser(ctx context.Context, id int) error {
	err := us.repo.DeleteUser(ctx, id)
	if err != nil {
		switch err {
		case domain.ErrDataNotFound:
			return err
		default:
			return domain.ErrInternal
		}
	}
	return nil
}
