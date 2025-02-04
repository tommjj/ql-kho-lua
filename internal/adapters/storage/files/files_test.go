package files

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
)

func TestNewFileStorage(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll("./static")
		os.RemoveAll("./temp")
	})

	fileStorage, err := NewFileStorage("./static", "./temp", time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, fileStorage, "fileStorage cannot be nil")
}

func TestLocalFileStorage_ImplementsIFileStorage(t *testing.T) {
	fileStorage := &localFileStorage{}

	assert.Implements(t, (*ports.IFileStorage)(nil), fileStorage, "fileStorage must implements IFileStorage")
}

//**

func TestSaveTempFile(t *testing.T) {
	fileStorage, err := NewFileStorage("./static", "./temp", time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.Open("./files.go")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	f, err := fileStorage.SaveTempFile(file, "test(2).go")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(f)
}

func TestSavePermanentFile(t *testing.T) {
	fileStorage, err := NewFileStorage("./static", "./temp", time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	err = fileStorage.SavePermanentFile("test(2).go")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteFile(t *testing.T) {
	fileStorage, err := NewFileStorage("./static", "./temp", time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	err = fileStorage.DeleteFile("test.go")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCleanupTempFiles(t *testing.T) {
	fileStorage, err := NewFileStorage("./static", "./temp", time.Second)
	if err != nil {
		t.Fatal(err)
	}

	err = fileStorage.CleanupTempFiles()
	if err != nil {
		t.Fatal(err)
	}
}
