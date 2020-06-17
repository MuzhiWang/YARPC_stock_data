package request_processor

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In

	Logger *zap.Logger
}


type Processor struct {
	logger *zap.Logger
}

func New(p Params) Processor {
	return Processor{
		logger: p.Logger,
	}
}


func (p *Processor) Handle(ctx context.Context, request, response interface{}, err error) {
	if recoverErr := recover(); recoverErr != nil {
		p.logger.Sugar().With(zap.Any("error", err))
		panic(err)
	}

	p.logger.Sugar().
		With(zap.Any("request", request)).
		With(zap.Any("response", response)).
		Info("request succeeded! ")
}
