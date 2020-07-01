package stock_data

import (
	"StockData/src/entity"
	"StockData/src/idl/tdx"
	"StockData/src/mapper"
	"context"
	"go.uber.org/fx"
)

var Module = fx.Provide(new)

type Controller interface {
	HelloTdxClient(ctx context.Context, name string) (string, error)
	ReadLocalTdxStockData(ctx context.Context, request entity.ReadLocalStockDataRequest) (*entity.StockData, error)
}

type controller struct {
	tdxClient tdxreader.TdxReaderYARPCClient
}

func new(tdxClient tdxreader.TdxReaderYARPCClient) Controller {
	return &controller{
		tdxClient: tdxClient,
	}
}

func (c *controller) HelloTdxClient(ctx context.Context, name string) (string, error) {
	res, err := c.tdxClient.Hello(ctx, &tdxreader.HelloRequest{
		Name: name,
	})

	return res.Message, err
}

func (c *controller) ReadLocalTdxStockData(
	ctx context.Context, request entity.ReadLocalStockDataRequest) (*entity.StockData, error) {
	res, err := c.tdxClient.ReadTdxFile(ctx, &tdxreader.ReadTdxFileRequest{
		FilePath: request.LocalPath,
		Metric: mapper.MapMetricFromEntityToProto(request.DataMetric),
	})

	if err != nil {
		return nil, err
	}

	return mapper.MapStockDataFromProtoToEntity(res), nil
}
