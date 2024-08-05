// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: author.proto

package author

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AuthorService_GetAuthors_FullMethodName = "/author.AuthorService/GetAuthors"
)

// AuthorServiceClient is the client API for AuthorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorServiceClient interface {
	GetAuthors(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Authors, error)
}

type authorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorServiceClient(cc grpc.ClientConnInterface) AuthorServiceClient {
	return &authorServiceClient{cc}
}

func (c *authorServiceClient) GetAuthors(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Authors, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Authors)
	err := c.cc.Invoke(ctx, AuthorService_GetAuthors_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorServiceServer is the server API for AuthorService service.
// All implementations must embed UnimplementedAuthorServiceServer
// for forward compatibility.
type AuthorServiceServer interface {
	GetAuthors(context.Context, *Ids) (*Authors, error)
	mustEmbedUnimplementedAuthorServiceServer()
}

// UnimplementedAuthorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthorServiceServer struct{}

func (UnimplementedAuthorServiceServer) GetAuthors(context.Context, *Ids) (*Authors, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthors not implemented")
}
func (UnimplementedAuthorServiceServer) mustEmbedUnimplementedAuthorServiceServer() {}
func (UnimplementedAuthorServiceServer) testEmbeddedByValue()                       {}

// UnsafeAuthorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorServiceServer will
// result in compilation errors.
type UnsafeAuthorServiceServer interface {
	mustEmbedUnimplementedAuthorServiceServer()
}

func RegisterAuthorServiceServer(s grpc.ServiceRegistrar, srv AuthorServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuthorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthorService_ServiceDesc, srv)
}

func _AuthorService_GetAuthors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ids)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).GetAuthors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorService_GetAuthors_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).GetAuthors(ctx, req.(*Ids))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthorService_ServiceDesc is the grpc.ServiceDesc for AuthorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "author.AuthorService",
	HandlerType: (*AuthorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthors",
			Handler:    _AuthorService_GetAuthors_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "author.proto",
}
