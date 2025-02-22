package mock

import (
	"context"

	"github.com/stretchr/testify/mock"
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
