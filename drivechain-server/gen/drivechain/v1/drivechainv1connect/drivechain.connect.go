// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: drivechain/v1/drivechain.proto

package drivechainv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/LayerTwo-Labs/sidesail/drivechain-server/gen/drivechain/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// DrivechainServiceName is the fully-qualified name of the DrivechainService service.
	DrivechainServiceName = "drivechain.v1.DrivechainService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// DrivechainServiceListTransactionsProcedure is the fully-qualified name of the DrivechainService's
	// ListTransactions RPC.
	DrivechainServiceListTransactionsProcedure = "/drivechain.v1.DrivechainService/ListTransactions"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	drivechainServiceServiceDescriptor                = v1.File_drivechain_v1_drivechain_proto.Services().ByName("DrivechainService")
	drivechainServiceListTransactionsMethodDescriptor = drivechainServiceServiceDescriptor.Methods().ByName("ListTransactions")
)

// DrivechainServiceClient is a client for the drivechain.v1.DrivechainService service.
type DrivechainServiceClient interface {
	ListTransactions(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.ListTransactionsResponse], error)
}

// NewDrivechainServiceClient constructs a client for the drivechain.v1.DrivechainService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewDrivechainServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) DrivechainServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &drivechainServiceClient{
		listTransactions: connect.NewClient[emptypb.Empty, v1.ListTransactionsResponse](
			httpClient,
			baseURL+DrivechainServiceListTransactionsProcedure,
			connect.WithSchema(drivechainServiceListTransactionsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// drivechainServiceClient implements DrivechainServiceClient.
type drivechainServiceClient struct {
	listTransactions *connect.Client[emptypb.Empty, v1.ListTransactionsResponse]
}

// ListTransactions calls drivechain.v1.DrivechainService.ListTransactions.
func (c *drivechainServiceClient) ListTransactions(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[v1.ListTransactionsResponse], error) {
	return c.listTransactions.CallUnary(ctx, req)
}

// DrivechainServiceHandler is an implementation of the drivechain.v1.DrivechainService service.
type DrivechainServiceHandler interface {
	ListTransactions(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.ListTransactionsResponse], error)
}

// NewDrivechainServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDrivechainServiceHandler(svc DrivechainServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	drivechainServiceListTransactionsHandler := connect.NewUnaryHandler(
		DrivechainServiceListTransactionsProcedure,
		svc.ListTransactions,
		connect.WithSchema(drivechainServiceListTransactionsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/drivechain.v1.DrivechainService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case DrivechainServiceListTransactionsProcedure:
			drivechainServiceListTransactionsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedDrivechainServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedDrivechainServiceHandler struct{}

func (UnimplementedDrivechainServiceHandler) ListTransactions(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.ListTransactionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("drivechain.v1.DrivechainService.ListTransactions is not implemented"))
}
