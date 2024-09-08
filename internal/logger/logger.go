package logger

import (
	"errors"
	"os"

	"github.com/tommjj/ql-kho-lua/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logLevelMap = map[string]zapcore.Level{
	"Debug":  zap.DebugLevel,
	"Info":   zap.InfoLevel,
	"Warn":   zap.WarnLevel,
	"Error":  zap.ErrorLevel,
	"DPanic": zap.DPanicLevel,
	"Panic":  zap.PanicLevel,
	"Fatal":  zap.FatalLevel,
}

func Set(conf config.Logger) error {
	log, err := New(conf)
	if err != nil {
		return err
	}

	zap.ReplaceGlobals(log)
	return nil
}

func New(conf config.Logger) (*zap.Logger, error) {
	level, ok := logLevelMap[conf.Level]
	if !ok {
		return nil, errors.New("logger level is not valid")
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(newJSONEncoder(conf.Encoder)),
		newWriteSyncer(conf.LogFileWriter),
		level,
	)

	return zap.New(core), nil
}

func newWriteSyncer(conf *config.LogFileWriter) zapcore.WriteSyncer {
	if conf == nil {
		return zapcore.AddSync(os.Stdout)
	}

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   conf.FileName,
		MaxSize:    conf.MaxSize,
		MaxBackups: conf.MaxBackups,
		MaxAge:     conf.MaxAge,
	})
	return zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(w),
		zapcore.AddSync(os.Stdout),
	)
}

func newJSONEncoder(mode string) zapcore.EncoderConfig {
	if mode != "production" {
		return zap.NewDevelopmentEncoderConfig()
	}
	return zap.NewProductionEncoderConfig()
}
