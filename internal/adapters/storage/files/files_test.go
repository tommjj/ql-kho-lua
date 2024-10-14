package files

import (
	"os"
	"testing"
	"time"
)

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
