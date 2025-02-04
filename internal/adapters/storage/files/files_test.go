package files

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

func teardownTestDirs(baseDir, tempDir string) {
	os.RemoveAll(baseDir)
	os.RemoveAll(tempDir)
}

func TestNewFileStorage(t *testing.T) {
	baseDir := "./static"
	tempDir := "./temp"

	t.Cleanup(func() {
		teardownTestDirs(baseDir, tempDir)
	})

	fileStorage, err := NewFileStorage(baseDir, tempDir, time.Hour)
	assert.NoError(t, err)
	assert.NotNil(t, fileStorage, "fileStorage cannot be nil")
}

func TestLocalFileStorage_ImplementsIFileStorage(t *testing.T) {
	fileStorage := &localFileStorage{}

	assert.Implements(t, (*ports.IFileStorage)(nil), fileStorage, "fileStorage must implements IFileStorage")
}

func TestSaveTempFile(t *testing.T) {
	baseDir := "./static"
	tempDir := "./temp"

	defer teardownTestDirs(baseDir, tempDir)

	fileStorage, err := NewFileStorage(baseDir, tempDir, time.Hour)
	assert.NoError(t, err)

	data := "test content"
	src := strings.NewReader(data)
	filename := "test.txt"

	path, err := fileStorage.SaveTempFile(src, filename)
	assert.NoError(t, err)
	assert.FileExists(t, path)
}

func TestSavePermanentFile(t *testing.T) {
	baseDir := "./static"
	tempDir := "./temp"

	defer teardownTestDirs(baseDir, tempDir)

	fileStorage, err := NewFileStorage(baseDir, tempDir, time.Hour)
	assert.NoError(t, err)

	data := "test content"
	src := strings.NewReader(data)
	filename := "test.txt"

	fileStorage.SaveTempFile(src, filename)

	err = fileStorage.SavePermanentFile(filename)
	assert.NoError(t, err)
	assert.FileExists(t, filepath.Join(baseDir, filename))
}

func TestDeleteFile(t *testing.T) {
	baseDir := "./static"
	tempDir := "./temp"

	defer teardownTestDirs(baseDir, tempDir)

	fileStorage, err := NewFileStorage(baseDir, tempDir, time.Hour)
	assert.NoError(t, err)

	data := "test content"
	src := strings.NewReader(data)
	filename := "test.txt"

	fileStorage.SaveTempFile(src, filename)
	fileStorage.SavePermanentFile(filename)

	err = fileStorage.DeleteFile(filename)
	assert.NoError(t, err)
	assert.NoFileExists(t, filepath.Join(baseDir, filename))
}

func TestCleanupTempFiles(t *testing.T) {
	baseDir := "./static"
	tempDir := "./temp"

	defer teardownTestDirs(baseDir, tempDir)

	fs, err := NewFileStorage(baseDir, tempDir, time.Second)
	assert.NoError(t, err)

	data := "test content"
	src := strings.NewReader(data)
	filename := "test.txt"

	fs.SaveTempFile(src, filename)
	time.Sleep(2 * time.Second)
	err = fs.CleanupTempFiles()
	assert.NoError(t, err)
	assert.NoFileExists(t, filepath.Join(tempDir, filename))
}

func TestCleanupTempFiles_NotExpired(t *testing.T) {
	baseDir := "./static"
	tempDir := "./temp"

	defer teardownTestDirs(baseDir, tempDir)

	fs, err := NewFileStorage(baseDir, tempDir, time.Hour)
	assert.NoError(t, err)

	data := "test content"
	src := strings.NewReader(data)
	filename := "test.txt"

	fs.SaveTempFile(src, filename)
	err = fs.CleanupTempFiles()
	assert.NoError(t, err)
	assert.FileExists(t, filepath.Join(tempDir, filename))
}
