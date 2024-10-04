package services

import (
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

type uploadService struct {
	storage ports.IFileStorage
}

func NewUploadService(fileStorage ports.IFileStorage) ports.IUploadService {
	return &uploadService{
		storage: fileStorage,
	}
}

func (u *uploadService) SaveTemp(file *multipart.FileHeader) (string, error) {
	filename := file.Filename

	f, err := file.Open()
	if err != nil {
		return "", domain.ErrInternal
	}

	ext := filepath.Ext(filename)
	id := uuid.NewString()
	newFilename := id + ext

	_, err = u.storage.SaveTempFile(f, newFilename)
	if err != nil {
		return "", domain.ErrInternal
	}

	return newFilename, nil
}
