package tdx

import (
	"StockData/src/idl/tdx"
	"go.uber.org/fx"
	"go.uber.org/yarpc"
)

//var Module = fx.Provide(new)

type Params struct {
	fx.In

	Dispatcher yarpc.Dispatcher
}

type Result struct {
	fx.Out

	Client tdxreader.TdxReaderYARPCClient
}

func new(p Params) Result {
	return Result{
		Client: tdxreader.NewTdxReaderYARPCClient(
			p.Dispatcher.MustOutboundConfig("TdxReader")),
	}
}
