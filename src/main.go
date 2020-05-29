package main

import (
	"StockData/src/handler"
	pb "StockData/src/idl"
	"fmt"

	"go.uber.org/yarpc/transport/grpc"

	"go.uber.org/yarpc/yarpcconfig"

	"go.uber.org/config"
	"go.uber.org/fx"
	"go.uber.org/yarpc"
)

func main() {
	cfg, err := config.NewYAMLProviderFromFiles("C:\\Users\\wmz66\\go\\src\\StockData\\src\\config\\base.yaml")
	if err != nil {
		fmt.Println("cfg error")
		panic(err.Error())
	}

	yCfg, err := buildYARPCConfig(cfg)
	if err != nil {
		panic("yCfg error " + err.Error())
	}
	dispatcher := yarpc.NewDispatcher(yCfg)
	if err := dispatcher.Start(); err != nil {
		panic(err.Error())
	}
	defer dispatcher.Stop()

	fx.New(opts()).Run()
}

func opts() fx.Option {
	return fx.Options(
		// fx.Provide(config.NewYAMLProviderFromFiles("C:\\Users\\wmz66\\go\\src\\StockData\\src\\config\\base.yaml")),
		fx.Provide(pb.NewFxStockDataYARPCProcedures()),
		handler.Module,
	)
}

func buildYARPCConfig(cfg config.Provider) (yarpc.Config, error) {
	var cfgData map[string]interface{}
	if err := cfg.Get("yarpc").Populate(&cfgData); err != nil {
		return yarpc.Config{}, fmt.Errorf("get yarpc config error " + err.Error())
	}

	configurator := yarpcconfig.New()
	configurator.RegisterTransport(grpcTransportSpec())
	yCfg, err := (*configurator).LoadConfig("stock-data", cfgData)
	if err != nil {
		return yarpc.Config{}, fmt.Errorf("load config failed " + err.Error())
	}

	return yCfg, nil
}

func grpcTransportSpec() yarpcconfig.TransportSpec {
	spec := grpc.TransportSpec()
	return spec
}
