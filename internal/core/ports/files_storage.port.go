package ports

import "io"

type IFileStorage interface {
	// SaveTempFile save file in temp folder
	SaveTempFile(src io.Reader, filename string) (string, error)
	// SavePermanentFile move file from temp to permanent
	SavePermanentFile(filename string) error
	// DeleteFile delete permanent file
	DeleteFile(filename string) error
	// CleanupTempFiles cleanup expired file
	CleanupTempFiles() error
}
