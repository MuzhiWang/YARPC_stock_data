package handler

import (
	"StockData/src/common/request_processor"
	pb "StockData/src/idl"
	"context"
	"fmt"
	"go.uber.org/zap"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(New),
	fx.Provide(pb.NewFxStockDataYARPCProcedures()),
)

type Params struct {
	fx.In

	RequestProcessor request_processor.Processor
	Logger           *zap.Logger
}

type handler struct {
	Logger           *zap.Logger
	RequestProcessor request_processor.Processor
}

func New(p Params) pb.StockDataYARPCServer {
	fmt.Println("start test handler! .......")
	return &handler{
		Logger:           p.Logger,
		RequestProcessor: p.RequestProcessor,
	}
}

func (h *handler) Test(ctx context.Context, request *pb.TestRequest) (
	res *pb.TestResponse, err error) {

	defer h.RequestProcessor.Handle(ctx, request, res, err)

	res = &pb.TestResponse{
		Value: request.Value,
	}

	return res, nil
}
