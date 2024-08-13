// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpc_demo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RouteGuidClient is the client API for RouteGuid service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteGuidClient interface {
	// unary
	GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error)
	// Server-side streaming
	ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuid_ListFeaturesClient, error)
	// Client-side streaming
	RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuid_RecordRouteClient, error)
	// Bidirectional streaming
	Recommend(ctx context.Context, opts ...grpc.CallOption) (RouteGuid_RecommendClient, error)
}

type routeGuidClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteGuidClient(cc grpc.ClientConnInterface) RouteGuidClient {
	return &routeGuidClient{cc}
}

func (c *routeGuidClient) GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error) {
	out := new(Feature)
	err := c.cc.Invoke(ctx, "/route.RouteGuid/GetFeature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuidClient) ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuid_ListFeaturesClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteGuid_ServiceDesc.Streams[0], "/route.RouteGuid/ListFeatures", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuidListFeaturesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RouteGuid_ListFeaturesClient interface {
	Recv() (*Feature, error)
	grpc.ClientStream
}

type routeGuidListFeaturesClient struct {
	grpc.ClientStream
}

func (x *routeGuidListFeaturesClient) Recv() (*Feature, error) {
	m := new(Feature)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuidClient) RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuid_RecordRouteClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteGuid_ServiceDesc.Streams[1], "/route.RouteGuid/RecordRoute", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuidRecordRouteClient{stream}
	return x, nil
}

type RouteGuid_RecordRouteClient interface {
	Send(*Point) error
	CloseAndRecv() (*RouteSummary, error)
	grpc.ClientStream
}

type routeGuidRecordRouteClient struct {
	grpc.ClientStream
}

func (x *routeGuidRecordRouteClient) Send(m *Point) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuidRecordRouteClient) CloseAndRecv() (*RouteSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RouteSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuidClient) Recommend(ctx context.Context, opts ...grpc.CallOption) (RouteGuid_RecommendClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteGuid_ServiceDesc.Streams[2], "/route.RouteGuid/Recommend", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuidRecommendClient{stream}
	return x, nil
}

type RouteGuid_RecommendClient interface {
	Send(*RecommendationRequest) error
	Recv() (*Feature, error)
	grpc.ClientStream
}

type routeGuidRecommendClient struct {
	grpc.ClientStream
}

func (x *routeGuidRecommendClient) Send(m *RecommendationRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuidRecommendClient) Recv() (*Feature, error) {
	m := new(Feature)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteGuidServer is the server API for RouteGuid service.
// All implementations must embed UnimplementedRouteGuidServer
// for forward compatibility
type RouteGuidServer interface {
	// unary
	GetFeature(context.Context, *Point) (*Feature, error)
	// Server-side streaming
	ListFeatures(*Rectangle, RouteGuid_ListFeaturesServer) error
	// Client-side streaming
	RecordRoute(RouteGuid_RecordRouteServer) error
	// Bidirectional streaming
	Recommend(RouteGuid_RecommendServer) error
	mustEmbedUnimplementedRouteGuidServer()
}

// UnimplementedRouteGuidServer must be embedded to have forward compatible implementations.
type UnimplementedRouteGuidServer struct {
}

func (UnimplementedRouteGuidServer) GetFeature(context.Context, *Point) (*Feature, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeature not implemented")
}
func (UnimplementedRouteGuidServer) ListFeatures(*Rectangle, RouteGuid_ListFeaturesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListFeatures not implemented")
}
func (UnimplementedRouteGuidServer) RecordRoute(RouteGuid_RecordRouteServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordRoute not implemented")
}
func (UnimplementedRouteGuidServer) Recommend(RouteGuid_RecommendServer) error {
	return status.Errorf(codes.Unimplemented, "method Recommend not implemented")
}
func (UnimplementedRouteGuidServer) mustEmbedUnimplementedRouteGuidServer() {}

// UnsafeRouteGuidServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteGuidServer will
// result in compilation errors.
type UnsafeRouteGuidServer interface {
	mustEmbedUnimplementedRouteGuidServer()
}

func RegisterRouteGuidServer(s grpc.ServiceRegistrar, srv RouteGuidServer) {
	s.RegisterService(&RouteGuid_ServiceDesc, srv)
}

func _RouteGuid_GetFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuidServer).GetFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/route.RouteGuid/GetFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuidServer).GetFeature(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuid_ListFeatures_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Rectangle)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouteGuidServer).ListFeatures(m, &routeGuidListFeaturesServer{stream})
}

type RouteGuid_ListFeaturesServer interface {
	Send(*Feature) error
	grpc.ServerStream
}

type routeGuidListFeaturesServer struct {
	grpc.ServerStream
}

func (x *routeGuidListFeaturesServer) Send(m *Feature) error {
	return x.ServerStream.SendMsg(m)
}

func _RouteGuid_RecordRoute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuidServer).RecordRoute(&routeGuidRecordRouteServer{stream})
}

type RouteGuid_RecordRouteServer interface {
	SendAndClose(*RouteSummary) error
	Recv() (*Point, error)
	grpc.ServerStream
}

type routeGuidRecordRouteServer struct {
	grpc.ServerStream
}

func (x *routeGuidRecordRouteServer) SendAndClose(m *RouteSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuidRecordRouteServer) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteGuid_Recommend_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuidServer).Recommend(&routeGuidRecommendServer{stream})
}

type RouteGuid_RecommendServer interface {
	Send(*Feature) error
	Recv() (*RecommendationRequest, error)
	grpc.ServerStream
}

type routeGuidRecommendServer struct {
	grpc.ServerStream
}

func (x *routeGuidRecommendServer) Send(m *Feature) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuidRecommendServer) Recv() (*RecommendationRequest, error) {
	m := new(RecommendationRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteGuid_ServiceDesc is the grpc.ServiceDesc for RouteGuid service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouteGuid_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "route.RouteGuid",
	HandlerType: (*RouteGuidServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeature",
			Handler:    _RouteGuid_GetFeature_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFeatures",
			Handler:       _RouteGuid_ListFeatures_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecordRoute",
			Handler:       _RouteGuid_RecordRoute_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Recommend",
			Handler:       _RouteGuid_Recommend_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "route.proto",
}