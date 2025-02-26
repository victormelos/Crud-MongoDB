package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log        *zap.Logger
	LOG_OUTPUT = "LOG_OUTPUT"
	LOG_LEVEL  = "LOG_LEVEL"
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{getOutputLogger()},
		Level:       zap.NewAtomicLevelAt(getLevelLogger()),
		Development: true,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			CallerKey:    "caller",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = logConfig.Build()

}
func getOutputLogger() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))
	if output == "" {
		return "stdout"
	}
	return output
}
func getLevelLogger() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	log.Sync()
}
func Error(message string, err error, tags ...zap.Field) {
	log.Error(message, append(tags, zap.NamedError("error", err))...)
	log.Sync()
}
func Debug(message string, tags ...zap.Field) {
	log.Debug(message, tags...)
	log.Sync()
}
func Warn(message string, tags ...zap.Field) {
	log.Warn(message, tags...)
	log.Sync()
}
func Panic(message string, tags ...zap.Field) {
	log.Panic(message, tags...)
	log.Sync()
}
func Fatal(message string, tags ...zap.Field) {
	log.Fatal(message, tags...)
	log.Sync()
}
