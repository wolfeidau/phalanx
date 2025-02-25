// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/index.proto

package proto

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

// IndexClient is the client API for Index service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IndexClient interface {
	LivenessCheck(ctx context.Context, in *LivenessCheckRequest, opts ...grpc.CallOption) (*LivenessCheckResponse, error)
	ReadinessCheck(ctx context.Context, in *ReadinessCheckRequest, opts ...grpc.CallOption) (*ReadinessCheckResponse, error)
	Metrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error)
	Cluster(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*ClusterResponse, error)
	CreateIndex(ctx context.Context, in *CreateIndexRequest, opts ...grpc.CallOption) (*CreateIndexResponse, error)
	DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*DeleteIndexResponse, error)
	AddDocuments(ctx context.Context, in *AddDocumentsRequest, opts ...grpc.CallOption) (*AddDocumentsResponse, error)
	DeleteDocuments(ctx context.Context, in *DeleteDocumentsRequest, opts ...grpc.CallOption) (*DeleteDocumentsResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type indexClient struct {
	cc grpc.ClientConnInterface
}

func NewIndexClient(cc grpc.ClientConnInterface) IndexClient {
	return &indexClient{cc}
}

func (c *indexClient) LivenessCheck(ctx context.Context, in *LivenessCheckRequest, opts ...grpc.CallOption) (*LivenessCheckResponse, error) {
	out := new(LivenessCheckResponse)
	err := c.cc.Invoke(ctx, "/index.Index/LivenessCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) ReadinessCheck(ctx context.Context, in *ReadinessCheckRequest, opts ...grpc.CallOption) (*ReadinessCheckResponse, error) {
	out := new(ReadinessCheckResponse)
	err := c.cc.Invoke(ctx, "/index.Index/ReadinessCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) Metrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error) {
	out := new(MetricsResponse)
	err := c.cc.Invoke(ctx, "/index.Index/Metrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) Cluster(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*ClusterResponse, error) {
	out := new(ClusterResponse)
	err := c.cc.Invoke(ctx, "/index.Index/Cluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) CreateIndex(ctx context.Context, in *CreateIndexRequest, opts ...grpc.CallOption) (*CreateIndexResponse, error) {
	out := new(CreateIndexResponse)
	err := c.cc.Invoke(ctx, "/index.Index/CreateIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*DeleteIndexResponse, error) {
	out := new(DeleteIndexResponse)
	err := c.cc.Invoke(ctx, "/index.Index/DeleteIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) AddDocuments(ctx context.Context, in *AddDocumentsRequest, opts ...grpc.CallOption) (*AddDocumentsResponse, error) {
	out := new(AddDocumentsResponse)
	err := c.cc.Invoke(ctx, "/index.Index/AddDocuments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) DeleteDocuments(ctx context.Context, in *DeleteDocumentsRequest, opts ...grpc.CallOption) (*DeleteDocumentsResponse, error) {
	out := new(DeleteDocumentsResponse)
	err := c.cc.Invoke(ctx, "/index.Index/DeleteDocuments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/index.Index/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexServer is the server API for Index service.
// All implementations must embed UnimplementedIndexServer
// for forward compatibility
type IndexServer interface {
	LivenessCheck(context.Context, *LivenessCheckRequest) (*LivenessCheckResponse, error)
	ReadinessCheck(context.Context, *ReadinessCheckRequest) (*ReadinessCheckResponse, error)
	Metrics(context.Context, *MetricsRequest) (*MetricsResponse, error)
	Cluster(context.Context, *ClusterRequest) (*ClusterResponse, error)
	CreateIndex(context.Context, *CreateIndexRequest) (*CreateIndexResponse, error)
	DeleteIndex(context.Context, *DeleteIndexRequest) (*DeleteIndexResponse, error)
	AddDocuments(context.Context, *AddDocumentsRequest) (*AddDocumentsResponse, error)
	DeleteDocuments(context.Context, *DeleteDocumentsRequest) (*DeleteDocumentsResponse, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	mustEmbedUnimplementedIndexServer()
}

// UnimplementedIndexServer must be embedded to have forward compatible implementations.
type UnimplementedIndexServer struct {
}

func (UnimplementedIndexServer) LivenessCheck(context.Context, *LivenessCheckRequest) (*LivenessCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LivenessCheck not implemented")
}
func (UnimplementedIndexServer) ReadinessCheck(context.Context, *ReadinessCheckRequest) (*ReadinessCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadinessCheck not implemented")
}
func (UnimplementedIndexServer) Metrics(context.Context, *MetricsRequest) (*MetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Metrics not implemented")
}
func (UnimplementedIndexServer) Cluster(context.Context, *ClusterRequest) (*ClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cluster not implemented")
}
func (UnimplementedIndexServer) CreateIndex(context.Context, *CreateIndexRequest) (*CreateIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIndex not implemented")
}
func (UnimplementedIndexServer) DeleteIndex(context.Context, *DeleteIndexRequest) (*DeleteIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIndex not implemented")
}
func (UnimplementedIndexServer) AddDocuments(context.Context, *AddDocumentsRequest) (*AddDocumentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDocuments not implemented")
}
func (UnimplementedIndexServer) DeleteDocuments(context.Context, *DeleteDocumentsRequest) (*DeleteDocumentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDocuments not implemented")
}
func (UnimplementedIndexServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedIndexServer) mustEmbedUnimplementedIndexServer() {}

// UnsafeIndexServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IndexServer will
// result in compilation errors.
type UnsafeIndexServer interface {
	mustEmbedUnimplementedIndexServer()
}

func RegisterIndexServer(s grpc.ServiceRegistrar, srv IndexServer) {
	s.RegisterService(&Index_ServiceDesc, srv)
}

func _Index_LivenessCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LivenessCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).LivenessCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/LivenessCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).LivenessCheck(ctx, req.(*LivenessCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_ReadinessCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadinessCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).ReadinessCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/ReadinessCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).ReadinessCheck(ctx, req.(*ReadinessCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_Metrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).Metrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/Metrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).Metrics(ctx, req.(*MetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_Cluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).Cluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/Cluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).Cluster(ctx, req.(*ClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_CreateIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).CreateIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/CreateIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).CreateIndex(ctx, req.(*CreateIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_DeleteIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).DeleteIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/DeleteIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).DeleteIndex(ctx, req.(*DeleteIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_AddDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDocumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).AddDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/AddDocuments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).AddDocuments(ctx, req.(*AddDocumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_DeleteDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDocumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).DeleteDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/DeleteDocuments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).DeleteDocuments(ctx, req.(*DeleteDocumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Index_ServiceDesc is the grpc.ServiceDesc for Index service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Index_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "index.Index",
	HandlerType: (*IndexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LivenessCheck",
			Handler:    _Index_LivenessCheck_Handler,
		},
		{
			MethodName: "ReadinessCheck",
			Handler:    _Index_ReadinessCheck_Handler,
		},
		{
			MethodName: "Metrics",
			Handler:    _Index_Metrics_Handler,
		},
		{
			MethodName: "Cluster",
			Handler:    _Index_Cluster_Handler,
		},
		{
			MethodName: "CreateIndex",
			Handler:    _Index_CreateIndex_Handler,
		},
		{
			MethodName: "DeleteIndex",
			Handler:    _Index_DeleteIndex_Handler,
		},
		{
			MethodName: "AddDocuments",
			Handler:    _Index_AddDocuments_Handler,
		},
		{
			MethodName: "DeleteDocuments",
			Handler:    _Index_DeleteDocuments_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Index_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/index.proto",
}
