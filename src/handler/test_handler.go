package handler

import (
	"StockData/src/common/request_processor"
	"StockData/src/controller/stock_data"
	"StockData/src/entity"
	pb "StockData/src/idl"
	"StockData/src/mapper"
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


	ret, err := h.controller.HelloTdxClient(ctx, request.Value)
	if err != nil {
		return nil, err
	}

	res = &pb.TestResponse{
		Value: ret,
	}

	return res, nil
}

func (h *handler) GetStockData(
	ctx context.Context, request *pb.GetStockDataRequest) (*pb.GetStockDataResponse, error) {
	testPath := "/Users/superegg/PycharmProjects/QuantTest/Tests/LC1/SZ/sz000001.lc1"
	stockData, err := h.controller.ReadLocalTdxStockData(ctx, entity.ReadLocalStockDataRequest{
		DataMetric: mapper.MapMetricFromProtoToEntity(request.Metric),
		LocalPath: testPath,
	})

	if err != nil {
		return nil, err
	}

	return mapper.MapStockDataFromEntityToProto(stockData), nil
}
