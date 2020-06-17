package main

import (
	"StockData/src/common/logger"
	"StockData/src/common/request_processor"
	"StockData/src/config"
	"StockData/src/handler"
	myyarpc "StockData/src/yarpc"
	"fmt"
	"go.uber.org/fx"
)

func main() {
	fmt.Println("start runnnnnnnnning....")
	fx.New(opts()).Run()
}

func opts() fx.Option {
	return fx.Options(
		fx.Provide(config.Module),
		handler.Module,
		myyarpc.Module,
		logger.Module,
		request_processor.Module,
	)
}
