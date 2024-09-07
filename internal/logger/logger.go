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

var L *zap.Logger = zap.NewExample()
var S *zap.SugaredLogger = zap.NewExample().Sugar()

func Set(conf config.Logger) error {
	log, err := New(conf)
	if err != nil {
		return err
	}

	L = log
	S = log.Sugar()
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
	if conf != nil {
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

	return zapcore.AddSync(os.Stdout)
}

func newJSONEncoder(mode string) zapcore.EncoderConfig {
	if mode != "production" {
		return zap.NewDevelopmentEncoderConfig()
	}
	return zap.NewProductionEncoderConfig()
}

func Debug(msg string, fields ...zapcore.Field) {
	L.Debug(msg, fields...)
}

func Info(msg string, fields ...zapcore.Field) {
	L.Info(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	L.Warn(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	L.Error(msg, fields...)
}

func DPanic(msg string, fields ...zapcore.Field) {
	L.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zapcore.Field) {
	L.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zapcore.Field) {
	L.Fatal(msg, fields...)
}

func Debugf(template string, args ...interface{}) {
	S.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	S.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	S.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	S.Errorf(template, args...)
}

func DPanicf(template string, args ...interface{}) {
	S.DPanicf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	S.Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	S.Fatalf(template, args...)
}

func Sync() error {
	var errs error = nil
	err := L.Sync()
	if err != nil {
		errs = errors.Join(errs, err)
	}

	err = S.Sync()
	if err != nil {
		errs = errors.Join(errs, err)
	}
	return errs
}
