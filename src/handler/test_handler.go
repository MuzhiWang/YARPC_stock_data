package handler

import (
	"StockData/src/common/request_processor"
	"StockData/src/controller/stock_data"
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

	StockDataController stock_data.Controller
}

type handler struct {
	logger           *zap.Logger
	requestProcessor request_processor.Processor

	controller stock_data.Controller
}

func New(p Params) pb.StockDataYARPCServer {
	fmt.Println("start test handler! .......")
	return &handler{
		logger:           p.Logger,
		requestProcessor: p.RequestProcessor,
		controller: p.StockDataController,
	}
}

func (h *handler) Test(ctx context.Context, request *pb.TestRequest) (
	res *pb.TestResponse, err error) {

	defer h.requestProcessor.Handle(ctx, request, res, err)


	err = h.controller.ReadStockData(ctx)
	if err != nil {
		return nil, err
	}

	//res = &pb.TestResponse{
	//	Value: ret.Message,
	//}

	return res, nil
}
