package mock

import (
	"io"

	"github.com/stretchr/testify/mock"
)

type MockFileStorage struct {
	mock.Mock
}

func (m *MockFileStorage) SaveTempFile(src io.Reader, filename string) (string, error) {
	args := m.Called(src, filename)
	return args.String(0), args.Error(1)
}

func (m *MockFileStorage) SavePermanentFile(filename string) error {
	args := m.Called(filename)
	return args.Error(0)
}

func (m *MockFileStorage) DeleteFile(filename string) error {
	args := m.Called(filename)
	return args.Error(0)
}

func (m *MockFileStorage) DeleteTempFile(filename string) error {
	args := m.Called(filename)
	return args.Error(0)
}

func (m *MockFileStorage) CleanupTempFiles() error {
	args := m.Called()
	return args.Error(0)
}
