package logger_test

import (
	"testing"

	"github.com/tommjj/ql-kho-lua/internal/config"
	"github.com/tommjj/ql-kho-lua/internal/logger"
)

func TestLogger(t *testing.T) {
	log, err := logger.New(config.Logger{
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
