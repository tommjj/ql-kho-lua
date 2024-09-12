package repository

import (
	"context"
	"errors"

	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb"
	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb/schema"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// implement ports.IUserRepository
type userRepo struct {
	db *mysqldb.MysqlDB
}

func NewUserRepo(db *mysqldb.MysqlDB) ports.IUserRepository {
	return &userRepo{
		db: db,
	}
}

// Create create an new user
func (ur *userRepo) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	createdUser := &schema.User{
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: user.Password,
		Role:     user.Role,
	}

	err := ur.db.WithContext(ctx).Create(createdUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, domain.ErrConflictingData
		}
		return nil, err
	}

	return convertToDomainUser(createdUser), nil
}

func (ur *userRepo) GetByID(ctx context.Context, id int) (*domain.User, error) {
	user := &schema.User{}

	err := ur.db.Where("id = ?", id).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	return convertToDomainUser(user), nil
}

func (ur *userRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &schema.User{}

	err := ur.db.Where("email = ?", email).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	return convertToDomainUser(user), nil
}

func (ur *userRepo) GetList(ctx context.Context, skip, limit int) ([]domain.User, error) {
	users := []schema.User{}

	err := ur.db.WithContext(ctx).Omit("password").Limit(limit).Offset((skip - 1) * limit).Find(&users).Error

	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, domain.ErrDataNotFound
	}

	domainUsers := make([]domain.User, 0, len(users))
	for _, user := range users {
		domainUsers = append(domainUsers, *convertToDomainUser(&user))
	}

	return domainUsers, nil
}

func (ur *userRepo) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	updateData := &schema.User{
		ID:       user.ID,
		Name:     user.Name,
		Phone:    user.Phone,
		Role:     user.Role,
		Email:    user.Email,
		Password: user.Password,
	}

	updatedUser := &schema.User{}

	result := ur.db.WithContext(ctx).Clauses(clause.Returning{}).Model(updatedUser).Where("id = ?", user.ID).Updates(updateData)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrDataNotFound
		}
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, domain.ErrNoUpdatedData
	}

	return convertToDomainUser(updatedUser), nil
}

func (ur *userRepo) Delete(ctx context.Context, id int) error {
	user := &schema.User{}

	result := ur.db.Where("id = ?", id).Delete(user)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain.ErrDataNotFound
	}

	return nil
}
