package files

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/tommjj/ql-kho-lua/internal/core/domain"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

type localFileStorage struct {
	baseDir string
	tempDir string
	maxAge  time.Duration
}

func NewFileStorage(baseDir string, tempDir string, maxAge time.Duration) (ports.IFileStorage, error) {
	if baseDir == tempDir {
		return nil, domain.ErrConflictingDirectory
	}

	err := os.MkdirAll(baseDir, 0777)
	if err != nil {
		return nil, domain.ErrCreateBaseDirectory
	}

	err = os.MkdirAll(tempDir, 0777)
	if err != nil {
		return nil, domain.ErrCreateTempDirectory
	}

	return &localFileStorage{
		baseDir: baseDir,
		tempDir: tempDir,
		maxAge:  maxAge,
	}, nil
}

func (s *localFileStorage) SaveTempFile(src io.Reader, filename string) (string, error) {
	tempFilePath := filepath.Join(s.tempDir, filename)

	out, err := os.Create(tempFilePath)
	if err != nil {
		return "", domain.ErrCreateFile
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return "", domain.ErrSaveFile
	}

	return tempFilePath, nil
}

func (s *localFileStorage) SavePermanentFile(filename string) error {
	tempFilePath := filepath.Join(s.tempDir, filename)
	permanentFilePath := filepath.Join(s.baseDir, filename)

	src, err := os.Open(tempFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return domain.ErrFileIsNotExist
		}
		return err
	}

	out, err := os.Create(permanentFilePath)
	if err != nil {
		return domain.ErrCreateFile
	}

	_, err = io.Copy(out, src)
	if err != nil {
		return domain.ErrSaveFile
	}
	out.Close()
	src.Close()

	err = s.deleteTempFile(filename)
	if err != nil {
		if err == domain.ErrFileIsNotExist {
			return nil
		}
		return err
	}

	return nil
}

func (s *localFileStorage) DeleteFile(filename string) error {
	permanentFilePath := filepath.Join(s.baseDir, filename)

	err := os.Remove(permanentFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return domain.ErrFileIsNotExist
		}
		return err
	}
	return nil
}

func (s *localFileStorage) deleteTempFile(filename string) error {
	tempFilePath := filepath.Join(s.tempDir, filename)

	err := os.Remove(tempFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return domain.ErrFileIsNotExist
		}
		return err
	}
	return nil
}

func (s *localFileStorage) CleanupTempFiles() error {
	files, err := os.ReadDir(s.tempDir)
	if err != nil {
		return err
	}

	for _, dir := range files {
		file, err := dir.Info()
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return err
		}

		modTime := file.ModTime()
		isExpired := time.Since(modTime) > s.maxAge
		if isExpired {
			err := s.deleteTempFile(file.Name())
			if err != nil && err != domain.ErrFileIsNotExist {
				return err
			}
		}
	}

	return nil
}
