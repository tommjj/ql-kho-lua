package ports

import "mime/multipart"

type IUploadService interface {
	// SaveTemp save temp file
	SaveTemp(file *multipart.FileHeader) (string, error)
}
