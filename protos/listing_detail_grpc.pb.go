// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: protos/listing_detail.proto

package listing

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

// ListingServiceClient is the client API for ListingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListingServiceClient interface {
	GetEntity(ctx context.Context, in *GetEntityRequest, opts ...grpc.CallOption) (*Entity, error)
}

type listingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewListingServiceClient(cc grpc.ClientConnInterface) ListingServiceClient {
	return &listingServiceClient{cc}
}

func (c *listingServiceClient) GetEntity(ctx context.Context, in *GetEntityRequest, opts ...grpc.CallOption) (*Entity, error) {
	out := new(Entity)
	err := c.cc.Invoke(ctx, "/listing.ListingService/GetEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListingServiceServer is the server API for ListingService service.
// All implementations must embed UnimplementedListingServiceServer
// for forward compatibility
type ListingServiceServer interface {
	GetEntity(context.Context, *GetEntityRequest) (*Entity, error)
	mustEmbedUnimplementedListingServiceServer()
}

// UnimplementedListingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedListingServiceServer struct {
}

func (UnimplementedListingServiceServer) GetEntity(context.Context, *GetEntityRequest) (*Entity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntity not implemented")
}
func (UnimplementedListingServiceServer) mustEmbedUnimplementedListingServiceServer() {}

// UnsafeListingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListingServiceServer will
// result in compilation errors.
type UnsafeListingServiceServer interface {
	mustEmbedUnimplementedListingServiceServer()
}

func RegisterListingServiceServer(s grpc.ServiceRegistrar, srv ListingServiceServer) {
	s.RegisterService(&ListingService_ServiceDesc, srv)
}

func _ListingService_GetEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListingServiceServer).GetEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listing.ListingService/GetEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListingServiceServer).GetEntity(ctx, req.(*GetEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ListingService_ServiceDesc is the grpc.ServiceDesc for ListingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ListingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "listing.ListingService",
	HandlerType: (*ListingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEntity",
			Handler:    _ListingService_GetEntity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/listing_detail.proto",
}
