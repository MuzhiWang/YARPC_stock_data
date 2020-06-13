package handler

import (
	pb "StockData/src/idl"
	"context"
	"fmt"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(New),
	fx.Provide(pb.NewFxStockDataYARPCProcedures()),
)

type handler struct {
}

func New() pb.StockDataYARPCServer {
	fmt.Println("start test handler! .......")
	return &handler{}
}

func (h *handler) Test(ctx context.Context, request *pb.TestRequest) (
	*pb.TestResponse, error) {
	return &pb.TestResponse{
		Value: request.Value,
	}, nil
}
