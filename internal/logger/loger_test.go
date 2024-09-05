package logger

import (
	"testing"

	"github.com/tommjj/ql-kho-lua/internal/config"
)

func TestLogger(t *testing.T) {
	log, err := New(config.Logger{
		Level:   "Info",
		Encoder: "development",
		LogFileWriter: &config.LogFileWriter{
			FileName:   "test.log",
			MaxSize:    10,
			MaxBackups: 3,
			MaxAge:     20,
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	log.Info("abc")
}
