package logger

import (
	"go.uber.org/zap"
	"time"
)

var log *zap.Logger

func Init() error {
	var err error
	log, err = zap.NewProduction()
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
