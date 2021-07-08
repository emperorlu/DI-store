// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package storage_server

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

// ObjectStoreClient is the client API for ObjectStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObjectStoreClient interface {
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type objectStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewObjectStoreClient(cc grpc.ClientConnInterface) ObjectStoreClient {
	return &objectStoreClient{cc}
}

func (c *objectStoreClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error) {
	out := new(FetchResponse)
	err := c.cc.Invoke(ctx, "/di_store.storage_server.ObjectStore/fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectStoreClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/di_store.storage_server.ObjectStore/get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectStoreClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/di_store.storage_server.ObjectStore/delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectStoreServer is the server API for ObjectStore service.
// All implementations must embed UnimplementedObjectStoreServer
// for forward compatibility
type ObjectStoreServer interface {
	Fetch(context.Context, *FetchRequest) (*FetchResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedObjectStoreServer()
}

// UnimplementedObjectStoreServer must be embedded to have forward compatible implementations.
type UnimplementedObjectStoreServer struct {
}

func (UnimplementedObjectStoreServer) Fetch(context.Context, *FetchRequest) (*FetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedObjectStoreServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedObjectStoreServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedObjectStoreServer) mustEmbedUnimplementedObjectStoreServer() {}

// UnsafeObjectStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObjectStoreServer will
// result in compilation errors.
type UnsafeObjectStoreServer interface {
	mustEmbedUnimplementedObjectStoreServer()
}

func RegisterObjectStoreServer(s grpc.ServiceRegistrar, srv ObjectStoreServer) {
	s.RegisterService(&ObjectStore_ServiceDesc, srv)
}

func _ObjectStore_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStoreServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/di_store.storage_server.ObjectStore/fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStoreServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectStore_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStoreServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/di_store.storage_server.ObjectStore/get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStoreServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectStore_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStoreServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/di_store.storage_server.ObjectStore/delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStoreServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ObjectStore_ServiceDesc is the grpc.ServiceDesc for ObjectStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ObjectStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "di_store.storage_server.ObjectStore",
	HandlerType: (*ObjectStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "fetch",
			Handler:    _ObjectStore_Fetch_Handler,
		},
		{
			MethodName: "get",
			Handler:    _ObjectStore_Get_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _ObjectStore_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "object_store.proto",
}
