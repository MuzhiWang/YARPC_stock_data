package main

import (
	"StockData/src/idl"
	"context"
	"fmt"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/grpc"
	"time"
)

func main() {
	//client := muzwang_stock_data.NewFxStockDataYARPCClient("stock_data")
	//
	//ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2 * time.Second))
	//defer cancel()
	//
	//_, err := client()


	dispatcher := yarpc.NewDispatcher(yarpc.Config{
		Name: "stockdata-client",
		Outbounds: yarpc.Outbounds{
			"stockdata": {
				Unary: grpc.NewTransport().NewSingleOutbound(":5432"),
			},
		},
	})

	client := muzwang_stock_data.NewStockDataYARPCClient(dispatcher.MustOutboundConfig("stockdata"))

	if err := dispatcher.Start(); err != nil {
		fmt.Printf("\nfailed to start client dispatcher. Error %s\n", err.Error())
	}
	defer dispatcher.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	res, err := client.Test(ctx, &muzwang_stock_data.TestRequest{
		Value: "muzhi!!",
	})
	if err != nil {
		fmt.Printf("\n failed to call . Error : %s \n", err.Error())
	} else {
		fmt.Println("no error for request! ")
	}

	fmt.Printf("\n succeeded to call. Res: %s", res.Value)
}
