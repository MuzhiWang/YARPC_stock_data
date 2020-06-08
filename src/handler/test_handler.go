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

func (h *handler) Test(context.Context, *pb.TestRequest) (
	*pb.TestResponse, error) {
	return nil, nil
}
