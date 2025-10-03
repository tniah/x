package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const (
	logLevelNone  = "none"
	logLevelDebug = "debug"
	logLevelInfo  = "info"
	logLevelWarn  = "warn"
	logLevelError = "error"
	logLevelPanic = "panic"
	logLevelFatal = "fatal"
)

type ZapLogger struct {
	*zap.Logger
}

func (z *ZapLogger) Debug(msg string, fields ...zap.Field) {
	z.Logger.Debug(msg, fields...)
}

func (z *ZapLogger) Info(msg string, fields ...zap.Field) {
	z.Logger.Info(msg, fields...)
}

func (z *ZapLogger) Warn(msg string, fields ...zap.Field) {
	z.Logger.Warn(msg, fields...)
}

func (z *ZapLogger) Error(msg string, fields ...zap.Field) {
	z.Logger.Error(msg, fields...)
}

func (z *ZapLogger) Panic(msg string, fields ...zap.Field) {
	z.Logger.Panic(msg, fields...)
}

func (z *ZapLogger) Fatal(msg string, fields ...zap.Field) {
	z.Logger.Fatal(msg, fields...)
}

func NewNoopLogger() *ZapLogger {
	return &ZapLogger{
		zap.NewNop(),
	}
}

func NewZapLogger(level string) (*ZapLogger, error) {
	if level == logLevelNone {
		return NewNoopLogger(), nil
	}

	var logLevel zapcore.Level
	switch level {
	case logLevelDebug:
		logLevel = zapcore.DebugLevel
	case logLevelInfo:
		logLevel = zapcore.InfoLevel
	case logLevelWarn:
		logLevel = zapcore.WarnLevel
	case logLevelError:
		logLevel = zapcore.ErrorLevel
	case logLevelPanic:
		logLevel = zapcore.PanicLevel
	case logLevelFatal:
		logLevel = zapcore.FatalLevel
	default:
		return nil, fmt.Errorf("logger - NewZapLogger - Unknown logger level: %s", level)
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.MessageKey = "message"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(logLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: true,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	logger, err := config.Build()
	if err != nil {
		return nil, fmt.Errorf("logger - NewZapLogger - config.Build: %w", err)
	}

	logger = logger.WithOptions(zap.AddCallerSkip(1))
	return &ZapLogger{logger}, nil
}

func MustNewZapLogger(level string) *ZapLogger {
	logger, err := NewZapLogger(level)
	if err != nil {
		panic(err)
	}

	return logger
}
