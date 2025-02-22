package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb"
	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb/schema"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
	"gorm.io/gorm"
)

// implement ports.IUserRepository
type userRepository struct {
	db *mysqldb.MysqlDB
}

func NewUserRepository(db *mysqldb.MysqlDB) ports.IUserRepository {
	return &userRepository{
		db: db,
	}
}

// Create create an new user
func (ur *userRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
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

	return convertToUser(createdUser), nil
}

func (ur *userRepository) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	user := &schema.User{}

	err := ur.db.WithContext(ctx).Where("id = ?", id).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	return convertToUser(user), nil
}

func (ur *userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &schema.User{}

	err := ur.db.WithContext(ctx).Where("email = ?", email).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	return convertToUser(user), nil
}

func (ur *userRepository) CountUsers(ctx context.Context, query string) (int64, error) {
	var count int64
	var err error

	q := ur.db.WithContext(ctx).Table("users")

	trimQuery := strings.TrimSpace(query)
	if trimQuery == "" {
		err = q.Count(&count).Error
	} else {
		err = q.Where("name LIKE ?", fmt.Sprintf("%%%v%%", trimQuery)).Count(&count).Error
	}

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (ur *userRepository) GetListUsers(ctx context.Context, query string, limit, skip int) ([]domain.User, error) {
	users := []domain.User{}
	var err error

	sql := ur.db.WithContext(ctx).Table("users").
		Select("id", "name", "email", "phone", "role").
		Limit(limit).Offset((skip - 1) * limit).Order("id desc").Where("deleted_at is NULL")

	trimQuery := strings.TrimSpace(query)
	if trimQuery == "" {
		err = sql.Scan(&users).Error
	} else {
		err = sql.Where("name LIKE ?", fmt.Sprintf("%%%v%%", trimQuery)).Scan(&users).Error
	}

	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, domain.ErrDataNotFound
	}

	return users, nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	updateData := &schema.User{
		ID:       user.ID,
		Name:     user.Name,
		Phone:    user.Phone,
		Role:     user.Role,
		Email:    user.Email,
		Password: user.Password,
	}

	result := ur.db.WithContext(ctx).Model(&schema.User{}).Where("id = ?", user.ID).Updates(updateData)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, domain.ErrConflictingData
		}
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, domain.ErrNoUpdatedData
	}

	return ur.GetUserByID(ctx, user.ID)
}

func (ur *userRepository) DeleteUser(ctx context.Context, id int) error {
	result := ur.db.WithContext(ctx).Where("id = ?", id).Delete(&schema.User{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain.ErrDataNotFound
	}

	return nil
}
