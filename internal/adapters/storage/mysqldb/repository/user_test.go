package repository

import (
	"context"
	"testing"
	"time"

	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb"
	"github.com/tommjj/ql-kho-lua/internal/config"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

func NewDefaultUserRepo() (ports.IUserRepository, error) {
	db, err := mysqldb.NewMysqlDB(config.DB{
		DSN:             "root:@tcp(127.0.0.1:3306)/ql?charset=utf8mb4&parseTime=True&loc=Local",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: time.Hour,
	})
	if err != nil {
		return nil, err
	}

	return NewUserRepository(db), nil
}

func TestUserRepo_Create(t *testing.T) {
	repo, err := NewDefaultUserRepo()
	if err != nil {
		t.Fatal(err)
	}

	user, err := repo.CreateUser(context.TODO(), &domain.User{
		Name:     "fiammetta",
		Phone:    "012345678",
		Email:    "fiammetta@gmail.com",
		Password: "123456789",
		Role:     domain.Staff,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(*user)
}

func TestUserRepo_GetByID(t *testing.T) {
	repo, err := NewDefaultUserRepo()
	if err != nil {
		t.Fatal(err)
	}

	user, err := repo.GetUserByID(context.TODO(), 1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(*user)
}

func TestUserRepo_GetByEmail(t *testing.T) {
	repo, err := NewDefaultUserRepo()
	if err != nil {
		t.Fatal(err)
	}

	user, err := repo.GetUserByEmail(context.TODO(), "mus.update@gmail.com")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(*user)
}

func TestUserRepo_GetList(t *testing.T) {
	repo, err := NewDefaultUserRepo()
	if err != nil {
		t.Fatal(err)
	}

	users, err := repo.GetListUsers(context.TODO(), "", 5, 1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", users)
}

func TestUserRepo_Update(t *testing.T) {
	repo, err := NewDefaultUserRepo()
	if err != nil {
		t.Fatal(err)
	}

	updated, err := repo.UpdateUser(context.TODO(), &domain.User{
		ID:       1,
		Name:     "mostima",
		Phone:    "1256",
		Password: "123456789",
		Role:     domain.Staff,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(updated)
}

func TestUserRepo_Delete(t *testing.T) {
	repo, err := NewDefaultUserRepo()
	if err != nil {
		t.Fatal(err)
	}

	err = repo.DeleteUser(context.TODO(), 1)
	if err != nil {
		t.Fatal(err)
	}
}
