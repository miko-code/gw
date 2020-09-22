package conf

import "go.uber.org/zap"

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}
func (l *Logger) GetLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./game.log",
	}

	return cfg.Build()
}
