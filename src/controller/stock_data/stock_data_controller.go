package stock_data

import (
	"StockData/src/idl/tdx"
	"context"
	"go.uber.org/fx"
)

var Module = fx.Provide(new)

type Controller interface {
	ReadStockData(ctx context.Context, name string) (string, error)
}

type controller struct {
	tdxClient tdxreader.TdxReaderYARPCClient
}

func new(tdxClient tdxreader.TdxReaderYARPCClient) Controller {
	return &controller{
		tdxClient: tdxClient,
	}
}

func (c *controller) ReadStockData(ctx context.Context, name string) (string, error) {
	res, err := c.tdxClient.Hello(ctx, &tdxreader.HelloRequest{
		Name: name,
	})

	return res.Message, err
}
