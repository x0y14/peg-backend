// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: operation/v1/operation.proto

package operationv1connect

import (
	v1 "backend/gen/operation/v1"
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// OperationServiceName is the fully-qualified name of the OperationService service.
	OperationServiceName = "operation.v1.OperationService"
)

// OperationServiceClient is a client for the operation.v1.OperationService service.
type OperationServiceClient interface {
	Ping(context.Context, *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error)
	FetchOperations(context.Context, *connect_go.Request[v1.FetchOperationsRequest]) (*connect_go.ServerStreamForClient[v1.FetchOperationsResponse], error)
}

// NewOperationServiceClient constructs a client for the operation.v1.OperationService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewOperationServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) OperationServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &operationServiceClient{
		ping: connect_go.NewClient[v1.PingRequest, v1.PingResponse](
			httpClient,
			baseURL+"/operation.v1.OperationService/Ping",
			opts...,
		),
		fetchOperations: connect_go.NewClient[v1.FetchOperationsRequest, v1.FetchOperationsResponse](
			httpClient,
			baseURL+"/operation.v1.OperationService/FetchOperations",
			opts...,
		),
	}
}

// operationServiceClient implements OperationServiceClient.
type operationServiceClient struct {
	ping            *connect_go.Client[v1.PingRequest, v1.PingResponse]
	fetchOperations *connect_go.Client[v1.FetchOperationsRequest, v1.FetchOperationsResponse]
}

// Ping calls operation.v1.OperationService.Ping.
func (c *operationServiceClient) Ping(ctx context.Context, req *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error) {
	return c.ping.CallUnary(ctx, req)
}

// FetchOperations calls operation.v1.OperationService.FetchOperations.
func (c *operationServiceClient) FetchOperations(ctx context.Context, req *connect_go.Request[v1.FetchOperationsRequest]) (*connect_go.ServerStreamForClient[v1.FetchOperationsResponse], error) {
	return c.fetchOperations.CallServerStream(ctx, req)
}

// OperationServiceHandler is an implementation of the operation.v1.OperationService service.
type OperationServiceHandler interface {
	Ping(context.Context, *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error)
	FetchOperations(context.Context, *connect_go.Request[v1.FetchOperationsRequest], *connect_go.ServerStream[v1.FetchOperationsResponse]) error
}

// NewOperationServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewOperationServiceHandler(svc OperationServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/operation.v1.OperationService/Ping", connect_go.NewUnaryHandler(
		"/operation.v1.OperationService/Ping",
		svc.Ping,
		opts...,
	))
	mux.Handle("/operation.v1.OperationService/FetchOperations", connect_go.NewServerStreamHandler(
		"/operation.v1.OperationService/FetchOperations",
		svc.FetchOperations,
		opts...,
	))
	return "/operation.v1.OperationService/", mux
}

// UnimplementedOperationServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedOperationServiceHandler struct{}

func (UnimplementedOperationServiceHandler) Ping(context.Context, *connect_go.Request[v1.PingRequest]) (*connect_go.Response[v1.PingResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("operation.v1.OperationService.Ping is not implemented"))
}

func (UnimplementedOperationServiceHandler) FetchOperations(context.Context, *connect_go.Request[v1.FetchOperationsRequest], *connect_go.ServerStream[v1.FetchOperationsResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("operation.v1.OperationService.FetchOperations is not implemented"))
}
