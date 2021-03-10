// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// PortServiceClient is the client API for PortService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortServiceClient interface {
	SubmitPort(ctx context.Context, in *Port, opts ...grpc.CallOption) (*SubmitPortResponse, error)
}

type portServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPortServiceClient(cc grpc.ClientConnInterface) PortServiceClient {
	return &portServiceClient{cc}
}

func (c *portServiceClient) SubmitPort(ctx context.Context, in *Port, opts ...grpc.CallOption) (*SubmitPortResponse, error) {
	out := new(SubmitPortResponse)
	err := c.cc.Invoke(ctx, "/proto.PortService/SubmitPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortServiceServer is the server API for PortService service.
// All implementations must embed UnimplementedPortServiceServer
// for forward compatibility
type PortServiceServer interface {
	SubmitPort(context.Context, *Port) (*SubmitPortResponse, error)
	mustEmbedUnimplementedPortServiceServer()
}

// UnimplementedPortServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPortServiceServer struct {
}

func (UnimplementedPortServiceServer) SubmitPort(context.Context, *Port) (*SubmitPortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitPort not implemented")
}
func (UnimplementedPortServiceServer) mustEmbedUnimplementedPortServiceServer() {}

// UnsafePortServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortServiceServer will
// result in compilation errors.
type UnsafePortServiceServer interface {
	mustEmbedUnimplementedPortServiceServer()
}

func RegisterPortServiceServer(s grpc.ServiceRegistrar, srv PortServiceServer) {
	s.RegisterService(&PortService_ServiceDesc, srv)
}

func _PortService_SubmitPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Port)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortServiceServer).SubmitPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PortService/SubmitPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortServiceServer).SubmitPort(ctx, req.(*Port))
	}
	return interceptor(ctx, in, info, handler)
}

// PortService_ServiceDesc is the grpc.ServiceDesc for PortService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PortService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PortService",
	HandlerType: (*PortServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitPort",
			Handler:    _PortService_SubmitPort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "port.proto",
}