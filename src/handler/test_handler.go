package handler

import (
	pb "StockData/src/idl"
	"context"

	"go.uber.org/fx"
)

var Module = fx.Provide(New)

type handler struct {
}

func New() pb.StockDataYARPCServer {
	return &handler{}
}

func (h *handler) Test(context.Context, *pb.TestRequest) (
	*pb.TestResponse, error) {
	return nil, nil
}
