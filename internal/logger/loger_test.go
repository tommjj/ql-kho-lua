package logger

import (
	"testing"

	"github.com/tommjj/ql-kho-lua/internal/config"
	"go.uber.org/zap"
)

func TestLogger(t *testing.T) {
	err := Set(config.Logger{
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
	zap.L().Info("aaa")
}
