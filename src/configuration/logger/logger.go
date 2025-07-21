package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log        *zap.Logger
	LOG_LEVEL  = "LOG_LEVEL"
	LOG_OUTPUT = "LOG_OUTPUT"
)

func init() {
	output := getOutputLogs()
	logLevel := getLogLevel()

	logConfig := zap.Config{
		OutputPaths: []string{output},
		Level:       zap.NewAtomicLevelAt(logLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = logConfig.Build()
}

func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	log.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Info(message, tags...)
	log.Sync()
}

func getOutputLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))
	if output == "" {
		return "stdout"
	}
	return output
}

func getLogLevel() zapcore.Level {
	switch strings.ToLower(os.Getenv(LOG_LEVEL)) {
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	case "debug":
		return zap.DebugLevel
	case "warn":
		return zap.WarnLevel
	default:
		return zap.InfoLevel
	}
}
