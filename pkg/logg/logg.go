package logg

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLogger is a struct for ZapLogger
type ZapLogger struct {
	Logger *zap.Logger
}

// NewZapLogger creates a new instance of ZapLogger with default settings.
func NewZapLogger(level zapcore.Level) (*ZapLogger, error) {
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		Development: false,
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:     "time",
			LevelKey:    "level",
			NameKey:     "logger",
			MessageKey:  "msg",
			LineEnding:  zapcore.DefaultLineEnding,
			EncodeLevel: zapcore.LowercaseColorLevelEncoder,
			EncodeTime:  zapcore.ISO8601TimeEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return &ZapLogger{Logger: logger}, nil
}

// Debug logs a message at DebugLevel.
func (zl *ZapLogger) Debug(msg string, tags ...zap.Field) {
	zl.Logger.Debug(msg, tags...)
}

// Info logs a message at InfoLevel.
func (zl *ZapLogger) Info(msg string, tags ...zap.Field) {
	zl.Logger.Info(msg, tags...)
}

// Warn logs a message at WarnLevel.
func (zl *ZapLogger) Warn(msg string, tags ...zap.Field) {
	zl.Logger.Warn(msg, tags...)
}

// Error logs a message at ErrorLevel.
func (zl *ZapLogger) Error(msg string, tags ...zap.Field) {
	zl.Logger.Error(msg, tags...)
}

// Fatal logs a message at FatalLevel.
func (zl *ZapLogger) Fatal(msg string, tags ...zap.Field) {
	zl.Logger.Fatal(msg, tags...)
}

// Panic logs a message at PanicLevel.
func (zl *ZapLogger) Panic(msg string, tags ...zap.Field) {
	zl.Logger.Panic(msg, tags...)
}

// Close closes the logger
func (zl *ZapLogger) Close() error {
	err := zl.Logger.Sync()
	if err != nil {
		return err
	}
	return nil
}
