// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: redpanda/api/dataplane/v1alpha2/pipeline.proto

package dataplanev1alpha2connect

import (
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"

	connect "connectrpc.com/connect"

	v1alpha2 "github.com/redpanda-data/console/backend/pkg/protogen/redpanda/api/dataplane/v1alpha2"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// PipelineServiceName is the fully-qualified name of the PipelineService service.
	PipelineServiceName = "redpanda.api.dataplane.v1alpha2.PipelineService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PipelineServiceCreatePipelineProcedure is the fully-qualified name of the PipelineService's
	// CreatePipeline RPC.
	PipelineServiceCreatePipelineProcedure = "/redpanda.api.dataplane.v1alpha2.PipelineService/CreatePipeline"
	// PipelineServiceGetPipelineProcedure is the fully-qualified name of the PipelineService's
	// GetPipeline RPC.
	PipelineServiceGetPipelineProcedure = "/redpanda.api.dataplane.v1alpha2.PipelineService/GetPipeline"
	// PipelineServiceListPipelinesProcedure is the fully-qualified name of the PipelineService's
	// ListPipelines RPC.
	PipelineServiceListPipelinesProcedure = "/redpanda.api.dataplane.v1alpha2.PipelineService/ListPipelines"
	// PipelineServiceUpdatePipelineProcedure is the fully-qualified name of the PipelineService's
	// UpdatePipeline RPC.
	PipelineServiceUpdatePipelineProcedure = "/redpanda.api.dataplane.v1alpha2.PipelineService/UpdatePipeline"
	// PipelineServiceDeletePipelineProcedure is the fully-qualified name of the PipelineService's
	// DeletePipeline RPC.
	PipelineServiceDeletePipelineProcedure = "/redpanda.api.dataplane.v1alpha2.PipelineService/DeletePipeline"
	// PipelineServiceStopPipelineProcedure is the fully-qualified name of the PipelineService's
	// StopPipeline RPC.
	PipelineServiceStopPipelineProcedure = "/redpanda.api.dataplane.v1alpha2.PipelineService/StopPipeline"
	// PipelineServiceStartPipelineProcedure is the fully-qualified name of the PipelineService's
	// StartPipeline RPC.
	PipelineServiceStartPipelineProcedure = "/redpanda.api.dataplane.v1alpha2.PipelineService/StartPipeline"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	pipelineServiceServiceDescriptor              = v1alpha2.File_redpanda_api_dataplane_v1alpha2_pipeline_proto.Services().ByName("PipelineService")
	pipelineServiceCreatePipelineMethodDescriptor = pipelineServiceServiceDescriptor.Methods().ByName("CreatePipeline")
	pipelineServiceGetPipelineMethodDescriptor    = pipelineServiceServiceDescriptor.Methods().ByName("GetPipeline")
	pipelineServiceListPipelinesMethodDescriptor  = pipelineServiceServiceDescriptor.Methods().ByName("ListPipelines")
	pipelineServiceUpdatePipelineMethodDescriptor = pipelineServiceServiceDescriptor.Methods().ByName("UpdatePipeline")
	pipelineServiceDeletePipelineMethodDescriptor = pipelineServiceServiceDescriptor.Methods().ByName("DeletePipeline")
	pipelineServiceStopPipelineMethodDescriptor   = pipelineServiceServiceDescriptor.Methods().ByName("StopPipeline")
	pipelineServiceStartPipelineMethodDescriptor  = pipelineServiceServiceDescriptor.Methods().ByName("StartPipeline")
)

// PipelineServiceClient is a client for the redpanda.api.dataplane.v1alpha2.PipelineService
// service.
type PipelineServiceClient interface {
	CreatePipeline(context.Context, *connect.Request[v1alpha2.CreatePipelineRequest]) (*connect.Response[v1alpha2.CreatePipelineResponse], error)
	GetPipeline(context.Context, *connect.Request[v1alpha2.GetPipelineRequest]) (*connect.Response[v1alpha2.GetPipelineResponse], error)
	ListPipelines(context.Context, *connect.Request[v1alpha2.ListPipelinesRequest]) (*connect.Response[v1alpha2.ListPipelinesResponse], error)
	UpdatePipeline(context.Context, *connect.Request[v1alpha2.UpdatePipelineRequest]) (*connect.Response[v1alpha2.UpdatePipelineResponse], error)
	DeletePipeline(context.Context, *connect.Request[v1alpha2.DeletePipelineRequest]) (*connect.Response[v1alpha2.DeletePipelineResponse], error)
	StopPipeline(context.Context, *connect.Request[v1alpha2.StopPipelineRequest]) (*connect.Response[v1alpha2.StopPipelineResponse], error)
	StartPipeline(context.Context, *connect.Request[v1alpha2.StartPipelineRequest]) (*connect.Response[v1alpha2.StartPipelineResponse], error)
}

// NewPipelineServiceClient constructs a client for the
// redpanda.api.dataplane.v1alpha2.PipelineService service. By default, it uses the Connect protocol
// with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To
// use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb()
// options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPipelineServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) PipelineServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &pipelineServiceClient{
		createPipeline: connect.NewClient[v1alpha2.CreatePipelineRequest, v1alpha2.CreatePipelineResponse](
			httpClient,
			baseURL+PipelineServiceCreatePipelineProcedure,
			connect.WithSchema(pipelineServiceCreatePipelineMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getPipeline: connect.NewClient[v1alpha2.GetPipelineRequest, v1alpha2.GetPipelineResponse](
			httpClient,
			baseURL+PipelineServiceGetPipelineProcedure,
			connect.WithSchema(pipelineServiceGetPipelineMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listPipelines: connect.NewClient[v1alpha2.ListPipelinesRequest, v1alpha2.ListPipelinesResponse](
			httpClient,
			baseURL+PipelineServiceListPipelinesProcedure,
			connect.WithSchema(pipelineServiceListPipelinesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updatePipeline: connect.NewClient[v1alpha2.UpdatePipelineRequest, v1alpha2.UpdatePipelineResponse](
			httpClient,
			baseURL+PipelineServiceUpdatePipelineProcedure,
			connect.WithSchema(pipelineServiceUpdatePipelineMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deletePipeline: connect.NewClient[v1alpha2.DeletePipelineRequest, v1alpha2.DeletePipelineResponse](
			httpClient,
			baseURL+PipelineServiceDeletePipelineProcedure,
			connect.WithSchema(pipelineServiceDeletePipelineMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		stopPipeline: connect.NewClient[v1alpha2.StopPipelineRequest, v1alpha2.StopPipelineResponse](
			httpClient,
			baseURL+PipelineServiceStopPipelineProcedure,
			connect.WithSchema(pipelineServiceStopPipelineMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		startPipeline: connect.NewClient[v1alpha2.StartPipelineRequest, v1alpha2.StartPipelineResponse](
			httpClient,
			baseURL+PipelineServiceStartPipelineProcedure,
			connect.WithSchema(pipelineServiceStartPipelineMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// pipelineServiceClient implements PipelineServiceClient.
type pipelineServiceClient struct {
	createPipeline *connect.Client[v1alpha2.CreatePipelineRequest, v1alpha2.CreatePipelineResponse]
	getPipeline    *connect.Client[v1alpha2.GetPipelineRequest, v1alpha2.GetPipelineResponse]
	listPipelines  *connect.Client[v1alpha2.ListPipelinesRequest, v1alpha2.ListPipelinesResponse]
	updatePipeline *connect.Client[v1alpha2.UpdatePipelineRequest, v1alpha2.UpdatePipelineResponse]
	deletePipeline *connect.Client[v1alpha2.DeletePipelineRequest, v1alpha2.DeletePipelineResponse]
	stopPipeline   *connect.Client[v1alpha2.StopPipelineRequest, v1alpha2.StopPipelineResponse]
	startPipeline  *connect.Client[v1alpha2.StartPipelineRequest, v1alpha2.StartPipelineResponse]
}

// CreatePipeline calls redpanda.api.dataplane.v1alpha2.PipelineService.CreatePipeline.
func (c *pipelineServiceClient) CreatePipeline(ctx context.Context, req *connect.Request[v1alpha2.CreatePipelineRequest]) (*connect.Response[v1alpha2.CreatePipelineResponse], error) {
	return c.createPipeline.CallUnary(ctx, req)
}

// GetPipeline calls redpanda.api.dataplane.v1alpha2.PipelineService.GetPipeline.
func (c *pipelineServiceClient) GetPipeline(ctx context.Context, req *connect.Request[v1alpha2.GetPipelineRequest]) (*connect.Response[v1alpha2.GetPipelineResponse], error) {
	return c.getPipeline.CallUnary(ctx, req)
}

// ListPipelines calls redpanda.api.dataplane.v1alpha2.PipelineService.ListPipelines.
func (c *pipelineServiceClient) ListPipelines(ctx context.Context, req *connect.Request[v1alpha2.ListPipelinesRequest]) (*connect.Response[v1alpha2.ListPipelinesResponse], error) {
	return c.listPipelines.CallUnary(ctx, req)
}

// UpdatePipeline calls redpanda.api.dataplane.v1alpha2.PipelineService.UpdatePipeline.
func (c *pipelineServiceClient) UpdatePipeline(ctx context.Context, req *connect.Request[v1alpha2.UpdatePipelineRequest]) (*connect.Response[v1alpha2.UpdatePipelineResponse], error) {
	return c.updatePipeline.CallUnary(ctx, req)
}

// DeletePipeline calls redpanda.api.dataplane.v1alpha2.PipelineService.DeletePipeline.
func (c *pipelineServiceClient) DeletePipeline(ctx context.Context, req *connect.Request[v1alpha2.DeletePipelineRequest]) (*connect.Response[v1alpha2.DeletePipelineResponse], error) {
	return c.deletePipeline.CallUnary(ctx, req)
}

// StopPipeline calls redpanda.api.dataplane.v1alpha2.PipelineService.StopPipeline.
func (c *pipelineServiceClient) StopPipeline(ctx context.Context, req *connect.Request[v1alpha2.StopPipelineRequest]) (*connect.Response[v1alpha2.StopPipelineResponse], error) {
	return c.stopPipeline.CallUnary(ctx, req)
}

// StartPipeline calls redpanda.api.dataplane.v1alpha2.PipelineService.StartPipeline.
func (c *pipelineServiceClient) StartPipeline(ctx context.Context, req *connect.Request[v1alpha2.StartPipelineRequest]) (*connect.Response[v1alpha2.StartPipelineResponse], error) {
	return c.startPipeline.CallUnary(ctx, req)
}

// PipelineServiceHandler is an implementation of the
// redpanda.api.dataplane.v1alpha2.PipelineService service.
type PipelineServiceHandler interface {
	CreatePipeline(context.Context, *connect.Request[v1alpha2.CreatePipelineRequest]) (*connect.Response[v1alpha2.CreatePipelineResponse], error)
	GetPipeline(context.Context, *connect.Request[v1alpha2.GetPipelineRequest]) (*connect.Response[v1alpha2.GetPipelineResponse], error)
	ListPipelines(context.Context, *connect.Request[v1alpha2.ListPipelinesRequest]) (*connect.Response[v1alpha2.ListPipelinesResponse], error)
	UpdatePipeline(context.Context, *connect.Request[v1alpha2.UpdatePipelineRequest]) (*connect.Response[v1alpha2.UpdatePipelineResponse], error)
	DeletePipeline(context.Context, *connect.Request[v1alpha2.DeletePipelineRequest]) (*connect.Response[v1alpha2.DeletePipelineResponse], error)
	StopPipeline(context.Context, *connect.Request[v1alpha2.StopPipelineRequest]) (*connect.Response[v1alpha2.StopPipelineResponse], error)
	StartPipeline(context.Context, *connect.Request[v1alpha2.StartPipelineRequest]) (*connect.Response[v1alpha2.StartPipelineResponse], error)
}

// NewPipelineServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPipelineServiceHandler(svc PipelineServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	pipelineServiceCreatePipelineHandler := connect.NewUnaryHandler(
		PipelineServiceCreatePipelineProcedure,
		svc.CreatePipeline,
		connect.WithSchema(pipelineServiceCreatePipelineMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	pipelineServiceGetPipelineHandler := connect.NewUnaryHandler(
		PipelineServiceGetPipelineProcedure,
		svc.GetPipeline,
		connect.WithSchema(pipelineServiceGetPipelineMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	pipelineServiceListPipelinesHandler := connect.NewUnaryHandler(
		PipelineServiceListPipelinesProcedure,
		svc.ListPipelines,
		connect.WithSchema(pipelineServiceListPipelinesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	pipelineServiceUpdatePipelineHandler := connect.NewUnaryHandler(
		PipelineServiceUpdatePipelineProcedure,
		svc.UpdatePipeline,
		connect.WithSchema(pipelineServiceUpdatePipelineMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	pipelineServiceDeletePipelineHandler := connect.NewUnaryHandler(
		PipelineServiceDeletePipelineProcedure,
		svc.DeletePipeline,
		connect.WithSchema(pipelineServiceDeletePipelineMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	pipelineServiceStopPipelineHandler := connect.NewUnaryHandler(
		PipelineServiceStopPipelineProcedure,
		svc.StopPipeline,
		connect.WithSchema(pipelineServiceStopPipelineMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	pipelineServiceStartPipelineHandler := connect.NewUnaryHandler(
		PipelineServiceStartPipelineProcedure,
		svc.StartPipeline,
		connect.WithSchema(pipelineServiceStartPipelineMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/redpanda.api.dataplane.v1alpha2.PipelineService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PipelineServiceCreatePipelineProcedure:
			pipelineServiceCreatePipelineHandler.ServeHTTP(w, r)
		case PipelineServiceGetPipelineProcedure:
			pipelineServiceGetPipelineHandler.ServeHTTP(w, r)
		case PipelineServiceListPipelinesProcedure:
			pipelineServiceListPipelinesHandler.ServeHTTP(w, r)
		case PipelineServiceUpdatePipelineProcedure:
			pipelineServiceUpdatePipelineHandler.ServeHTTP(w, r)
		case PipelineServiceDeletePipelineProcedure:
			pipelineServiceDeletePipelineHandler.ServeHTTP(w, r)
		case PipelineServiceStopPipelineProcedure:
			pipelineServiceStopPipelineHandler.ServeHTTP(w, r)
		case PipelineServiceStartPipelineProcedure:
			pipelineServiceStartPipelineHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPipelineServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPipelineServiceHandler struct{}

func (UnimplementedPipelineServiceHandler) CreatePipeline(context.Context, *connect.Request[v1alpha2.CreatePipelineRequest]) (*connect.Response[v1alpha2.CreatePipelineResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha2.PipelineService.CreatePipeline is not implemented"))
}

func (UnimplementedPipelineServiceHandler) GetPipeline(context.Context, *connect.Request[v1alpha2.GetPipelineRequest]) (*connect.Response[v1alpha2.GetPipelineResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha2.PipelineService.GetPipeline is not implemented"))
}

func (UnimplementedPipelineServiceHandler) ListPipelines(context.Context, *connect.Request[v1alpha2.ListPipelinesRequest]) (*connect.Response[v1alpha2.ListPipelinesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha2.PipelineService.ListPipelines is not implemented"))
}

func (UnimplementedPipelineServiceHandler) UpdatePipeline(context.Context, *connect.Request[v1alpha2.UpdatePipelineRequest]) (*connect.Response[v1alpha2.UpdatePipelineResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha2.PipelineService.UpdatePipeline is not implemented"))
}

func (UnimplementedPipelineServiceHandler) DeletePipeline(context.Context, *connect.Request[v1alpha2.DeletePipelineRequest]) (*connect.Response[v1alpha2.DeletePipelineResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha2.PipelineService.DeletePipeline is not implemented"))
}

func (UnimplementedPipelineServiceHandler) StopPipeline(context.Context, *connect.Request[v1alpha2.StopPipelineRequest]) (*connect.Response[v1alpha2.StopPipelineResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha2.PipelineService.StopPipeline is not implemented"))
}

func (UnimplementedPipelineServiceHandler) StartPipeline(context.Context, *connect.Request[v1alpha2.StartPipelineRequest]) (*connect.Response[v1alpha2.StartPipelineResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha2.PipelineService.StartPipeline is not implemented"))
}