// Code generated by protoc-gen-yarpc-go. DO NOT EDIT.
// source: stock_data.proto

package muzwang_stock_data

import (
	"context"
	"io/ioutil"
	"reflect"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"go.uber.org/fx"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/api/x/restriction"
	"go.uber.org/yarpc/encoding/protobuf"
	"go.uber.org/yarpc/encoding/protobuf/reflection"
)

var _ = ioutil.NopCloser

// StockDataYARPCClient is the YARPC client-side interface for the StockData service.
type StockDataYARPCClient interface {
	Test(context.Context, *TestRequest, ...yarpc.CallOption) (*TestResponse, error)
	GetStockData(context.Context, *GetStockDataRequest, ...yarpc.CallOption) (*GetStockDataResponse, error)
}

func newStockDataYARPCClient(clientConfig transport.ClientConfig, anyResolver jsonpb.AnyResolver, options ...protobuf.ClientOption) StockDataYARPCClient {
	return &_StockDataYARPCCaller{protobuf.NewStreamClient(
		protobuf.ClientParams{
			ServiceName:  "muzwang.stock_data.StockData",
			ClientConfig: clientConfig,
			AnyResolver:  anyResolver,
			Options:      options,
		},
	)}
}

// NewStockDataYARPCClient builds a new YARPC client for the StockData service.
func NewStockDataYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) StockDataYARPCClient {
	return newStockDataYARPCClient(clientConfig, nil, options...)
}

// StockDataYARPCServer is the YARPC server-side interface for the StockData service.
type StockDataYARPCServer interface {
	Test(context.Context, *TestRequest) (*TestResponse, error)
	GetStockData(context.Context, *GetStockDataRequest) (*GetStockDataResponse, error)
}

type buildStockDataYARPCProceduresParams struct {
	Server      StockDataYARPCServer
	AnyResolver jsonpb.AnyResolver
}

func buildStockDataYARPCProcedures(params buildStockDataYARPCProceduresParams) []transport.Procedure {
	handler := &_StockDataYARPCHandler{params.Server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName: "muzwang.stock_data.StockData",
			UnaryHandlerParams: []protobuf.BuildProceduresUnaryHandlerParams{
				{
					MethodName: "Test",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:      handler.Test,
							NewRequest:  newStockDataServiceTestYARPCRequest,
							AnyResolver: params.AnyResolver,
						},
					),
				},
				{
					MethodName: "GetStockData",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:      handler.GetStockData,
							NewRequest:  newStockDataServiceGetStockDataYARPCRequest,
							AnyResolver: params.AnyResolver,
						},
					),
				},
			},
			OnewayHandlerParams: []protobuf.BuildProceduresOnewayHandlerParams{},
			StreamHandlerParams: []protobuf.BuildProceduresStreamHandlerParams{},
		},
	)
}

// BuildStockDataYARPCProcedures prepares an implementation of the StockData service for YARPC registration.
func BuildStockDataYARPCProcedures(server StockDataYARPCServer) []transport.Procedure {
	return buildStockDataYARPCProcedures(buildStockDataYARPCProceduresParams{Server: server})
}

// FxStockDataYARPCClientParams defines the input
// for NewFxStockDataYARPCClient. It provides the
// paramaters to get a StockDataYARPCClient in an
// Fx application.
type FxStockDataYARPCClientParams struct {
	fx.In

	Provider    yarpc.ClientConfig
	AnyResolver jsonpb.AnyResolver  `name:"yarpcfx" optional:"true"`
	Restriction restriction.Checker `optional:"true"`
}

// FxStockDataYARPCClientResult defines the output
// of NewFxStockDataYARPCClient. It provides a
// StockDataYARPCClient to an Fx application.
type FxStockDataYARPCClientResult struct {
	fx.Out

	Client StockDataYARPCClient

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// NewFxStockDataYARPCClient provides a StockDataYARPCClient
// to an Fx application using the given name for routing.
//
//  fx.Provide(
//    muzwang_stock_data.NewFxStockDataYARPCClient("service-name"),
//    ...
//  )
func NewFxStockDataYARPCClient(name string, options ...protobuf.ClientOption) interface{} {
	return func(params FxStockDataYARPCClientParams) FxStockDataYARPCClientResult {
		cc := params.Provider.ClientConfig(name)

		if params.Restriction != nil {
			if namer, ok := cc.GetUnaryOutbound().(transport.Namer); ok {
				if err := params.Restriction.Check(protobuf.Encoding, namer.TransportName()); err != nil {
					panic(err.Error())
				}
			}
		}

		return FxStockDataYARPCClientResult{
			Client: newStockDataYARPCClient(cc, params.AnyResolver, options...),
		}
	}
}

// FxStockDataYARPCProceduresParams defines the input
// for NewFxStockDataYARPCProcedures. It provides the
// paramaters to get StockDataYARPCServer procedures in an
// Fx application.
type FxStockDataYARPCProceduresParams struct {
	fx.In

	Server      StockDataYARPCServer
	AnyResolver jsonpb.AnyResolver `name:"yarpcfx" optional:"true"`
}

// FxStockDataYARPCProceduresResult defines the output
// of NewFxStockDataYARPCProcedures. It provides
// StockDataYARPCServer procedures to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group.
// Dig 1.2 or newer must be used for this feature to work.
type FxStockDataYARPCProceduresResult struct {
	fx.Out

	Procedures     []transport.Procedure
	ReflectionMeta reflection.ServerMeta
}

// NewFxStockDataYARPCProcedures provides StockDataYARPCServer procedures to an Fx application.
// It expects a StockDataYARPCServer to be present in the container.
//
//  fx.Provide(
//    muzwang_stock_data.NewFxStockDataYARPCProcedures(),
//    ...
//  )
func NewFxStockDataYARPCProcedures() interface{} {
	return func(params FxStockDataYARPCProceduresParams) FxStockDataYARPCProceduresResult {
		return FxStockDataYARPCProceduresResult{
			Procedures: buildStockDataYARPCProcedures(buildStockDataYARPCProceduresParams{
				Server:      params.Server,
				AnyResolver: params.AnyResolver,
			}),
			ReflectionMeta: reflection.ServerMeta{
				ServiceName:     "muzwang.stock_data.StockData",
				FileDescriptors: yarpcFileDescriptorClosure46d4f64edd7c98cf,
			},
		}
	}
}

type _StockDataYARPCCaller struct {
	streamClient protobuf.StreamClient
}

func (c *_StockDataYARPCCaller) Test(ctx context.Context, request *TestRequest, options ...yarpc.CallOption) (*TestResponse, error) {
	responseMessage, err := c.streamClient.Call(ctx, "Test", request, newStockDataServiceTestYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*TestResponse)
	if !ok {
		return nil, protobuf.CastError(emptyStockDataServiceTestYARPCResponse, responseMessage)
	}
	return response, err
}

func (c *_StockDataYARPCCaller) GetStockData(ctx context.Context, request *GetStockDataRequest, options ...yarpc.CallOption) (*GetStockDataResponse, error) {
	responseMessage, err := c.streamClient.Call(ctx, "GetStockData", request, newStockDataServiceGetStockDataYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*GetStockDataResponse)
	if !ok {
		return nil, protobuf.CastError(emptyStockDataServiceGetStockDataYARPCResponse, responseMessage)
	}
	return response, err
}

type _StockDataYARPCHandler struct {
	server StockDataYARPCServer
}

func (h *_StockDataYARPCHandler) Test(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *TestRequest
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*TestRequest)
		if !ok {
			return nil, protobuf.CastError(emptyStockDataServiceTestYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.Test(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func (h *_StockDataYARPCHandler) GetStockData(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *GetStockDataRequest
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*GetStockDataRequest)
		if !ok {
			return nil, protobuf.CastError(emptyStockDataServiceGetStockDataYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.GetStockData(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func newStockDataServiceTestYARPCRequest() proto.Message {
	return &TestRequest{}
}

func newStockDataServiceTestYARPCResponse() proto.Message {
	return &TestResponse{}
}

func newStockDataServiceGetStockDataYARPCRequest() proto.Message {
	return &GetStockDataRequest{}
}

func newStockDataServiceGetStockDataYARPCResponse() proto.Message {
	return &GetStockDataResponse{}
}

var (
	emptyStockDataServiceTestYARPCRequest          = &TestRequest{}
	emptyStockDataServiceTestYARPCResponse         = &TestResponse{}
	emptyStockDataServiceGetStockDataYARPCRequest  = &GetStockDataRequest{}
	emptyStockDataServiceGetStockDataYARPCResponse = &GetStockDataResponse{}
)

var yarpcFileDescriptorClosure46d4f64edd7c98cf = [][]byte{
	// stock_data.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0x4b, 0x6f, 0xd3, 0x40,
		0x10, 0xc6, 0xce, 0xab, 0x99, 0x94, 0xd4, 0x9d, 0x06, 0xb0, 0x42, 0x11, 0x55, 0x40, 0xa2, 0xea,
		0x21, 0xa0, 0x70, 0x41, 0xc0, 0x09, 0xca, 0x23, 0x87, 0xaa, 0x68, 0x9b, 0x3b, 0x5a, 0xea, 0xa1,
		0xb1, 0x9a, 0x78, 0x83, 0xbd, 0x4e, 0x08, 0xff, 0x87, 0x33, 0x57, 0xf8, 0x75, 0xa0, 0x9d, 0xf5,
		0xa3, 0x05, 0x57, 0x70, 0xe5, 0x36, 0xf3, 0xcd, 0x37, 0x8f, 0x9d, 0x6f, 0x77, 0xc1, 0x4b, 0xb4,
		0x3a, 0x3d, 0x7f, 0x1f, 0x48, 0x2d, 0x87, 0x8b, 0x58, 0x69, 0x85, 0x38, 0x4f, 0xbf, 0xac, 0x64,
		0x74, 0x36, 0x2c, 0x23, 0x83, 0x7b, 0xd0, 0x99, 0x50, 0xa2, 0x05, 0x7d, 0x4a, 0x29, 0xd1, 0xd8,
		0x83, 0xc6, 0x52, 0xce, 0x52, 0xf2, 0x9d, 0x3d, 0x67, 0xbf, 0x2d, 0xac, 0x33, 0xb8, 0x0f, 0x9b,
		0x96, 0x94, 0x2c, 0x54, 0x94, 0xd0, 0x15, 0xac, 0x97, 0xd0, 0x3e, 0x94, 0x9a, 0x84, 0x8c, 0xce,
		0x08, 0x77, 0xa1, 0x9d, 0x68, 0x19, 0x6b, 0x83, 0x64, 0xb4, 0x12, 0x40, 0x1f, 0x5a, 0x14, 0x05,
		0x1c, 0x73, 0x39, 0x96, 0xbb, 0x83, 0x9f, 0x2e, 0xec, 0xbc, 0x21, 0x7d, 0x62, 0x26, 0x3c, 0x94,
		0x5a, 0xe6, 0x83, 0x8d, 0xa0, 0x39, 0x27, 0x1d, 0x87, 0xa7, 0x5c, 0xac, 0x3b, 0xea, 0x0f, 0xff,
		0x3c, 0xcc, 0xf0, 0x88, 0x19, 0x22, 0x63, 0xe2, 0x33, 0x68, 0x07, 0xf9, 0x40, 0xdc, 0xa7, 0x33,
		0xba, 0x53, 0x95, 0x56, 0x4c, 0x2d, 0x4a, 0xbe, 0x49, 0x66, 0xca, 0x64, 0xbd, 0x20, 0xbf, 0xc6,
		0x3d, 0x2b, 0x93, 0x4f, 0x72, 0x92, 0x28, 0xf9, 0xd8, 0x05, 0x37, 0x0c, 0xfc, 0x3a, 0x1f, 0xcd,
		0x0d, 0x03, 0xfc, 0x08, 0x5b, 0x4a, 0x4f, 0x29, 0x7e, 0x17, 0xab, 0x05, 0xc5, 0x3a, 0xa4, 0xc4,
		0x6f, 0xec, 0xd5, 0xf6, 0x3b, 0xa3, 0xe7, 0x55, 0x25, 0x2b, 0xce, 0x3f, 0x3c, 0xbe, 0x9c, 0xfe,
		0x2a, 0xd2, 0xf1, 0x5a, 0xfc, 0x5e, 0xb4, 0xff, 0x02, 0x7a, 0x55, 0x44, 0xf4, 0xa0, 0x76, 0x4e,
		0xeb, 0x4c, 0x07, 0x63, 0x96, 0x12, 0xba, 0x17, 0x24, 0x7c, 0xea, 0x3e, 0x71, 0x06, 0xdf, 0x1d,
		0xd8, 0x38, 0x22, 0x2d, 0x4d, 0xf7, 0xff, 0x6b, 0xed, 0x83, 0xaf, 0x0e, 0xb4, 0x84, 0x5a, 0xf1,
		0xe4, 0x7d, 0xd8, 0x30, 0x55, 0x27, 0xe1, 0x3c, 0xbf, 0x7f, 0x85, 0x8f, 0x08, 0x75, 0xb5, 0xa0,
		0x88, 0x87, 0x73, 0x04, 0xdb, 0x06, 0x9b, 0x86, 0x67, 0x53, 0xee, 0xe9, 0x08, 0xb6, 0xcd, 0xda,
		0x66, 0x6a, 0xc5, 0x3a, 0x3a, 0xc2, 0x98, 0x66, 0x6d, 0xa7, 0x33, 0x95, 0x90, 0xdf, 0x60, 0xcc,
		0x3a, 0x78, 0x13, 0x9a, 0x72, 0xae, 0xd2, 0x48, 0xfb, 0x4d, 0x86, 0x33, 0xcf, 0xe0, 0x4b, 0x35,
		0x4b, 0xe7, 0xe4, 0xb7, 0x2c, 0x6e, 0xbd, 0xc1, 0x1a, 0x7a, 0x97, 0x35, 0xce, 0xde, 0xd5, 0x23,
		0xa8, 0xcf, 0x49, 0x4b, 0x9e, 0xb7, 0x33, 0xda, 0xbd, 0x62, 0xd7, 0xac, 0x8c, 0x60, 0x26, 0x3e,
		0x84, 0xba, 0x81, 0x7d, 0x97, 0x6f, 0xd3, 0xed, 0xaa, 0x8c, 0x6c, 0x21, 0x82, 0x89, 0x07, 0xdf,
		0x1c, 0x68, 0x5a, 0xbd, 0x70, 0x1b, 0xae, 0x5b, 0x6b, 0x1c, 0x2d, 0xe5, 0x2c, 0x0c, 0xbc, 0x6b,
		0xb8, 0x03, 0x5b, 0x16, 0x3a, 0x8e, 0xe8, 0x28, 0x8c, 0x52, 0x4d, 0x9e, 0x83, 0x37, 0x60, 0xdb,
		0x82, 0xaf, 0xc3, 0x65, 0x86, 0x26, 0x9e, 0x8b, 0xb7, 0x60, 0xc7, 0xc2, 0x93, 0x69, 0x18, 0xeb,
		0x75, 0x1e, 0xa8, 0x95, 0x75, 0x8f, 0x23, 0x7a, 0xab, 0xd2, 0xd8, 0xab, 0x23, 0x42, 0x37, 0xe3,
		0xae, 0x94, 0x81, 0x12, 0xaf, 0x81, 0x3d, 0xf0, 0xf2, 0xfc, 0x98, 0xc8, 0xa2, 0x4d, 0xf4, 0x60,
		0xb3, 0x48, 0x3e, 0x94, 0x6b, 0xaf, 0x75, 0x30, 0x86, 0x76, 0x21, 0xb6, 0x49, 0x2a, 0x9c, 0x72,
		0x6c, 0x84, 0x6e, 0x81, 0xb2, 0xe1, 0x39, 0x97, 0xb0, 0x71, 0x14, 0xd0, 0x67, 0xcf, 0x1d, 0xfd,
		0x70, 0xb2, 0x5a, 0x7c, 0x43, 0xc6, 0x50, 0x37, 0xbf, 0x1a, 0xde, 0xad, 0xda, 0xda, 0x85, 0x4f,
		0xb1, 0xbf, 0x77, 0x35, 0x21, 0x13, 0x4e, 0xc2, 0xe6, 0x45, 0x41, 0xf1, 0xc1, 0x3f, 0x3e, 0xeb,
		0xfe, 0xfe, 0xdf, 0x89, 0xb6, 0xc5, 0x87, 0x26, 0xff, 0xe1, 0x8f, 0x7f, 0x05, 0x00, 0x00, 0xff,
		0xff, 0x32, 0xf8, 0x9f, 0x4b, 0xd7, 0x05, 0x00, 0x00,
	},
}

func init() {
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) StockDataYARPCClient {
			return NewStockDataYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
}
