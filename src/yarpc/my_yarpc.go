package myyarpc

import (
	"context"
	"fmt"
	"go.uber.org/yarpc/transport/grpc"
	"net"

	// "google.golang.org/grpc"

	"go.uber.org/yarpc/encoding/protobuf/reflection"

	"go.uber.org/config"
	"go.uber.org/fx"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
)

var Module = fx.Options(
	fx.Provide(New),
	fx.Invoke(RunDispatcher),
)

type Params struct {
	fx.In

	LifeCycle       fx.Lifecycle
	Cfg             config.Provider
}

type Result struct {
	fx.Out

	Dispatcher yarpc.Dispatcher
}

func New(p Params) (Result, error) {
	fmt.Println("*****#####++++++++++")
	dispatcher, err := newServerDispatcher()
	if err != nil {
		return Result{}, err
	}

	return Result{
		Dispatcher: *dispatcher,
	}, nil
}

func newServerDispatcher() (*yarpc.Dispatcher, error) {
	listener, err := net.Listen("tcp", ":5432")
	if err != nil {
		return nil, err
	}
	dispatcher := yarpc.NewDispatcher(
		yarpc.Config{
			Name: "stockdata",
			Inbounds: yarpc.Inbounds{
				grpc.NewTransport().NewInbound(listener),
			},
			Outbounds: yarpc.Outbounds{
				"TdxReader": transport.Outbounds{
					Unary: grpc.NewTransport().NewSingleOutbound(":50051"),
				},
			},
		},
	)


	//ggrefl.Register()
	return dispatcher, nil
}

type RunDispatcherParams struct {
	fx.In

	LifeCycle fx.Lifecycle

	Transports      []transport.Procedure
	ReflectionMeta  reflection.ServerMeta
	//TdxClient tdxreader.TdxReaderYARPCClient
	Dispatcher yarpc.Dispatcher
}

func RunDispatcher(p RunDispatcherParams) {
	p.LifeCycle.Append(fx.Hook{
		//OnStart: func(context.Context) error {
		//	if err := dispatcher.Start(); err != nil {
		//		return err
		//	}
		//	return nil
		//},
		OnStop: func(context.Context) error {
			if err := p.Dispatcher.Stop(); err != nil {
				return err
			}
			return nil
		},
	})

	p.Dispatcher.Register(p.Transports)

	fmt.Println("start to run dispatcher !!!!!!")
	if err := p.Dispatcher.Start(); err != nil {
		panic("errorrrrrrr start dispatcher")
	}
}

//func buildYARPCConfig(cfg config.Provider) (yarpc.Config, error) {
//	var cfgData map[string]interface{}
//	if err := cfg.Get("yarpc").Populate(&cfgData); err != nil {
//		return yarpc.Config{}, fmt.Errorf("get yarpc config error " + err.Error())
//	}
//
//	configurator := yarpcconfig.New()
//	configurator.RegisterTransport(grpcTransportSpec())
//	yCfg, err := (*configurator).LoadConfig("stock-data", cfgData)
//	if err != nil {
//		return yarpc.Config{}, fmt.Errorf("load config failed " + err.Error())
//	}
//
//	return yCfg, nil
//}
//
//func grpcTransportSpec() yarpcconfig.TransportSpec {
//	spec := grpc.TransportSpec()
//	return spec
//}
