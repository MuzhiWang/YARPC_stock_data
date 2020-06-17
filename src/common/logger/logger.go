package logger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Provide(New)

type Result struct {
	fx.Out

	Logger *zap.Logger
}

func New() Result {
	logger, _ := zap.NewProduction()

	return Result{
		Logger: logger,
	}
}
