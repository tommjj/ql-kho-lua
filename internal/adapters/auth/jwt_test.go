package auth_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tommjj/ql-kho-lua/internal/adapters/auth"
	"github.com/tommjj/ql-kho-lua/internal/config"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

type MockKeyRepository struct {
	mock.Mock
}

func (m *MockKeyRepository) SetKey(ctx context.Context, userID int, key string) error {
	args := m.Called(ctx, userID, key)
	return args.Error(0)
}

func (m *MockKeyRepository) GetKey(ctx context.Context, userID int) (string, error) {
	args := m.Called(ctx, userID)
	return args.String(0), args.Error(1)
}

func (m *MockKeyRepository) DelKey(ctx context.Context, userID int) error {
	args := m.Called(ctx, userID)
	return args.Error(0)
}

func TestImplementsIAuthService(t *testing.T) {
	mockRepo := new(MockKeyRepository)
	conf := config.Auth{SecretKey: "secret", Duration: time.Hour}
	service := auth.NewJWTTokenService(conf, mockRepo)

	assert.Implements(t, (*ports.ITokenService)(nil), service)
}

func TestCreateToken(t *testing.T) {
	mockRepo := new(MockKeyRepository)
	conf := config.Auth{SecretKey: "secret", Duration: time.Hour}
	service := auth.NewJWTTokenService(conf, mockRepo)

	user := &domain.User{ID: 1, Name: "Test", Email: "test@example.com", Role: domain.Root}
	mockRepo.On("SetKey", mock.Anything, user.ID, mock.Anything).Return(nil)

	token, err := service.CreateToken(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	mockRepo.AssertExpectations(t)
}

func TestVerifyToken(t *testing.T) {
	mockRepo := new(MockKeyRepository)
	key := ""
	conf := config.Auth{SecretKey: "secret", Duration: time.Hour}
	service := auth.NewJWTTokenService(conf, mockRepo)

	user := &domain.User{ID: 1, Name: "Test", Email: "test@example.com", Phone: "+5555555555", Role: domain.Root}
	mockRepo.On("SetKey", mock.Anything, user.ID, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		key = args.String(2)
	})

	token, _ := service.CreateToken(user)
	mockRepo.On("GetKey", mock.Anything, user.ID).Return(key, nil)

	payload, err := service.VerifyToken(token)

	assert.NoError(t, err)
	assert.Equal(t, user.ID, payload.ID)
	assert.Equal(t, user.Email, payload.Email)
	assert.Equal(t, user.Name, payload.Name)
	assert.Equal(t, user.Role, payload.Role)
}

func TestVerifyToken_InvalidKey(t *testing.T) {
	mockRepo := new(MockKeyRepository)
	conf := config.Auth{SecretKey: "secret", Duration: time.Hour}
	service := auth.NewJWTTokenService(conf, mockRepo)

	user := &domain.User{ID: 1, Name: "Test", Email: "test@example.com", Phone: "+5555555555", Role: domain.Root}
	mockRepo.On("SetKey", mock.Anything, user.ID, mock.Anything).Return(nil)
	mockRepo.On("GetKey", mock.Anything, user.ID).Return("wrong-key", nil)

	token, _ := service.CreateToken(user)
	mockRepo.On("GetKey", mock.Anything, user.ID).Return("wrong-key", nil)

	_, err := service.VerifyToken(token)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrInvalidToken, err)
}

func TestVerifyToken_Expired(t *testing.T) {
	mockRepo := new(MockKeyRepository)
	key := ""
	conf := config.Auth{SecretKey: "secret", Duration: -time.Hour}
	service := auth.NewJWTTokenService(conf, mockRepo)

	user := &domain.User{ID: 1, Name: "Test", Email: "test@example.com", Phone: "+5555555555", Role: domain.Root}
	mockRepo.On("SetKey", mock.Anything, user.ID, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		key = args.String(2)
	})

	token, _ := service.CreateToken(user)
	mockRepo.On("GetKey", mock.Anything, user.ID).Return(key, nil)

	_, err := service.VerifyToken(token)
	assert.Error(t, err)
	assert.Equal(t, domain.ErrExpiredToken, err)
}
