package services

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
	mockRepo "github.com/tommjj/ql-kho-lua/internal/core/services/mock"
)

func TestImplements(t *testing.T) {
	assert.Implements(t, (*ports.IRiceService)(nil), new(riceService))
}

func TestCreateRice_Success(t *testing.T) {
	mockRepo := new(mockRepo.MockRiceRepository)
	mockRepo.On("CreateRice", mock.Anything, mock.Anything).Return(&domain.Rice{}, nil)

	service := NewRiceService(mockRepo)
	_, err := service.CreateRice(context.TODO(), &domain.Rice{})
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateRice_FailConflicting(t *testing.T) {
	mockRepo := new(mockRepo.MockRiceRepository)
	mockRepo.On("CreateRice", mock.Anything, mock.Anything).Return(nil, domain.ErrConflictingData)

	service := NewRiceService(mockRepo)
	_, err := service.CreateRice(context.TODO(), &domain.Rice{})
	assert.Equal(t, domain.ErrConflictingData, err)

	mockRepo.AssertExpectations(t)
}

func TestCreateRice_FailUnknownErr(t *testing.T) {
	mockRepo := new(mockRepo.MockRiceRepository)
	mockRepo.On("CreateRice", mock.Anything, mock.Anything).Return(nil, errors.New("unknown error"))

	service := NewRiceService(mockRepo)
	_, err := service.CreateRice(context.TODO(), &domain.Rice{})
	assert.Equal(t, domain.ErrInternal, err)

	mockRepo.AssertExpectations(t)
}

func TestGetRiceByID_Success(t *testing.T) {
	mockRepo := new(mockRepo.MockRiceRepository)
	mockRepo.On("GetRiceByID", mock.Anything, 1).Return(&domain.Rice{
		ID:   1,
		Name: "Rice",
	}, nil)

	service := NewRiceService(mockRepo)
	rice, err := service.GetRiceByID(context.TODO(), 1)
	assert.Nil(t, err)
	assert.Equal(t, rice.ID, 1)
	assert.Equal(t, rice.Name, "Rice")
	mockRepo.AssertExpectations(t)
}

func TestGetRiceByID_FailNotFound(t *testing.T) {
	mockRepo := new(mockRepo.MockRiceRepository)
	mockRepo.On("GetRiceByID", mock.Anything, 1).Return(nil, domain.ErrDataNotFound)

	service := NewRiceService(mockRepo)
	_, err := service.GetRiceByID(context.TODO(), 1)
	assert.Equal(t, domain.ErrDataNotFound, err)

	mockRepo.AssertExpectations(t)
}

func TestGetRiceByID_FailUnknownErr(t *testing.T) {
	mockRepo := new(mockRepo.MockRiceRepository)
	mockRepo.On("GetRiceByID", mock.Anything, 1).Return(nil, errors.New("unknown error"))

	service := NewRiceService(mockRepo)
	_, err := service.GetRiceByID(context.TODO(), 1)
	assert.Equal(t, domain.ErrInternal, err)

	mockRepo.AssertExpectations(t)
}

func TestCountRice_Success(t *testing.T) {
	mockRepo := new(mockRepo.MockRiceRepository)
	mockRepo.On("CountRice", mock.Anything, "name").Return(int64(1), nil)

	service := NewRiceService(mockRepo)
	count, err := service.CountRice(context.TODO(), "name")
	assert.Nil(t, err)
	assert.Equal(t, count, int64(1))
	mockRepo.AssertExpectations(t)
}

func TestCountRice_FailNotFound(t *testing.T) {
	mockRepo := new(mockRepo.MockRiceRepository)
	mockRepo.On("CountRice", mock.Anything, "name").Return(int64(0), domain.ErrDataNotFound)

	service := NewRiceService(mockRepo)
	count, _ := service.CountRice(context.TODO(), "name")
	assert.Equal(t, count, int64(0))

	mockRepo.AssertExpectations(t)
}

func TestCountRice_FailUnknownErr(t *testing.T) {
	mockRepo := new(mockRepo.MockRiceRepository)
	mockRepo.On("CountRice", mock.Anything, "name").Return(int64(0), errors.New("unknown error"))

	service := NewRiceService(mockRepo)
	_, err := service.CountRice(context.TODO(), "name")
	assert.Equal(t, domain.ErrInternal, err)

	mockRepo.AssertExpectations(t)
}

func TestCountRice_SuccessNoQuery(t *testing.T) {
	mockRepo := new(mockRepo.MockRiceRepository)
	mockRepo.On("CountRice", mock.Anything, "").Return(int64(1), nil)

	service := NewRiceService(mockRepo)
	count, err := service.CountRice(context.TODO(), "")
	assert.Nil(t, err)
	assert.Equal(t, count, int64(1))
	mockRepo.AssertExpectations(t)
}

func TestGetListRice_Success(t *testing.T) {
	mockRepo := new(mockRepo.MockRiceRepository)
	mockRepo.On("GetListRice", mock.Anything, "name", 10, 0).Return([]domain.Rice{
		{
			ID:   1,
			Name: "Rice",
		},
	}, nil)

	service := NewRiceService(mockRepo)
	rice, err := service.GetListRice(context.TODO(), "name", 10, 0)
	assert.Nil(t, err)
	assert.Equal(t, rice[0].ID, 1)
	assert.Equal(t, rice[0].Name, "Rice")
	mockRepo.AssertExpectations(t)
}

func TestGetListRice_FailNotFound(t *testing.T) {
	mockRepo := new(mockRepo.MockRiceRepository)
	mockRepo.On("GetListRice", mock.Anything, "name", 10, 0).Return(nil, domain.ErrDataNotFound)

	service := NewRiceService(mockRepo)
	_, err := service.GetListRice(context.TODO(), "name", 10, 0)
	assert.Equal(t, domain.ErrDataNotFound, err)

	mockRepo.AssertExpectations(t)
}

func TestGetListRice_FailUnknownErr(t *testing.T) {
	mockRepo := new(mockRepo.MockRiceRepository)
	mockRepo.On("GetListRice", mock.Anything, "name", 10, 0).Return(nil, errors.New("unknown error"))

	service := NewRiceService(mockRepo)
	_, err := service.GetListRice(context.TODO(), "name", 10, 0)
	assert.Equal(t, domain.ErrInternal, err)

	mockRepo.AssertExpectations(t)
}

func TestUpdateRice_Success(t *testing.T) {
	mockRepo := new(mockRepo.MockRiceRepository)
	mockRepo.On("GetRiceByID", mock.Anything, mock.Anything).Return(nil, nil)
	mockRepo.On("UpdateRice", mock.Anything, mock.Anything).Return(&domain.Rice{}, nil)

	service := NewRiceService(mockRepo)
	_, err := service.UpdateRice(context.TODO(), &domain.Rice{})
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateRice_FailNotFound(t *testing.T) {
	mockRepo := new(mockRepo.MockRiceRepository)
	mockRepo.On("GetRiceByID", mock.Anything, mock.Anything).Return(nil, domain.ErrDataNotFound)

	service := NewRiceService(mockRepo)
	_, err := service.UpdateRice(context.TODO(), &domain.Rice{})
	assert.Equal(t, domain.ErrDataNotFound, err)

	mockRepo.AssertExpectations(t)
}

func TestUpdateRice_FailUnknownErr(t *testing.T) {
	mockRepo := new(mockRepo.MockRiceRepository)
	mockRepo.On("GetRiceByID", mock.Anything, mock.Anything).Return(nil, errors.New("unknown error"))

	service := NewRiceService(mockRepo)
	_, err := service.UpdateRice(context.TODO(), &domain.Rice{})
	assert.Equal(t, domain.ErrInternal, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteRice_Success(t *testing.T) {
	mockRepo := new(mockRepo.MockRiceRepository)
	mockRepo.On("DeleteRice", mock.Anything, 1).Return(nil)

	service := NewRiceService(mockRepo)
	err := service.DeleteRice(context.TODO(), 1)
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteRice_FailNotFound(t *testing.T) {
	mockRepo := new(mockRepo.MockRiceRepository)
	mockRepo.On("DeleteRice", mock.Anything, 1).Return(domain.ErrDataNotFound)

	service := NewRiceService(mockRepo)
	err := service.DeleteRice(context.TODO(), 1)
	assert.Equal(t, domain.ErrDataNotFound, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteRice_FailUnknownErr(t *testing.T) {
	mockRepo := new(mockRepo.MockRiceRepository)
	mockRepo.On("DeleteRice", mock.Anything, 1).Return(errors.New("unknown error"))

	service := NewRiceService(mockRepo)
	err := service.DeleteRice(context.TODO(), 1)
	assert.Equal(t, domain.ErrInternal, err)

	mockRepo.AssertExpectations(t)
}
