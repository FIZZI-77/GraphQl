package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var log *zap.Logger

func Init() error {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Encoding:    "json", // важно для ELK
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "timestamp",
			LevelKey:      "level",
			NameKey:       "logger",
			CallerKey:     "caller",
			MessageKey:    "message",
			StacktraceKey: "stacktrace",
			EncodeLevel:   zapcore.LowercaseLevelEncoder,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"}, // логи идут в stdout
		ErrorOutputPaths: []string{"stderr"},
	}

	var err error
	log, err = cfg.Build()
	if err != nil {
		return err
	}
	return nil
}

func Sync() {
	_ = log.Sync()
}

func RecordMetric(name string, value float64, labels map[string]string) {
	log.Info("metric",
		zap.String("timestamp", time.Now().Format(time.RFC3339)),
		zap.String("type", "metric"),
		zap.String("name", name),
		zap.Float64("value", value),
		zap.Any("labels", labels),
	)
}

func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	log.Warn(msg, fields...)
}
