package mock

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
)

type MockRiceRepository struct {
	mock.Mock
}

func (m *MockRiceRepository) CreateRice(ctx context.Context, rice *domain.Rice) (*domain.Rice, error) {
	args := m.Called(ctx, rice)

	if rice, ok := args.Get(0).(*domain.Rice); ok {
		return rice, args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (m *MockRiceRepository) GetRiceByID(ctx context.Context, id int) (*domain.Rice, error) {
	args := m.Called(ctx, id)
	if rice, ok := args.Get(0).(*domain.Rice); ok {
		return rice, args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (m *MockRiceRepository) CountRice(ctx context.Context, query string) (int64, error) {
	args := m.Called(ctx, query)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockRiceRepository) GetListRice(ctx context.Context, query string, limit, skip int) ([]domain.Rice, error) {
	args := m.Called(ctx, query, limit, skip)
	if rice, ok := args.Get(0).([]domain.Rice); ok {
		return rice, args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (m *MockRiceRepository) UpdateRice(ctx context.Context, rice *domain.Rice) (*domain.Rice, error) {
	args := m.Called(ctx, rice)
	if rice, ok := args.Get(0).(*domain.Rice); ok {
		return rice, args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (m *MockRiceRepository) DeleteRice(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
