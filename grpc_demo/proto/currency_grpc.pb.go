// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: proto/currency.proto

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

// CurrencyServiceClient is the client API for CurrencyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CurrencyServiceClient interface {
	CreateCurrency(ctx context.Context, in *Currency, opts ...grpc.CallOption) (*Currency, error)
	ReadCurrency(ctx context.Context, in *ReadCurrencyRequest, opts ...grpc.CallOption) (*Currency, error)
	UpdateCurrency(ctx context.Context, in *Currency, opts ...grpc.CallOption) (*Currency, error)
	DeleteCurrency(ctx context.Context, in *DeleteCurrencyRequest, opts ...grpc.CallOption) (*DeleteCurrencyResponse, error)
	ListCurrencies(ctx context.Context, in *Empty, opts ...grpc.CallOption) (CurrencyService_ListCurrenciesClient, error)
}

type currencyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyServiceClient(cc grpc.ClientConnInterface) CurrencyServiceClient {
	return &currencyServiceClient{cc}
}

func (c *currencyServiceClient) CreateCurrency(ctx context.Context, in *Currency, opts ...grpc.CallOption) (*Currency, error) {
	out := new(Currency)
	err := c.cc.Invoke(ctx, "/currency.CurrencyService/CreateCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyServiceClient) ReadCurrency(ctx context.Context, in *ReadCurrencyRequest, opts ...grpc.CallOption) (*Currency, error) {
	out := new(Currency)
	err := c.cc.Invoke(ctx, "/currency.CurrencyService/ReadCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyServiceClient) UpdateCurrency(ctx context.Context, in *Currency, opts ...grpc.CallOption) (*Currency, error) {
	out := new(Currency)
	err := c.cc.Invoke(ctx, "/currency.CurrencyService/UpdateCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyServiceClient) DeleteCurrency(ctx context.Context, in *DeleteCurrencyRequest, opts ...grpc.CallOption) (*DeleteCurrencyResponse, error) {
	out := new(DeleteCurrencyResponse)
	err := c.cc.Invoke(ctx, "/currency.CurrencyService/DeleteCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyServiceClient) ListCurrencies(ctx context.Context, in *Empty, opts ...grpc.CallOption) (CurrencyService_ListCurrenciesClient, error) {
	stream, err := c.cc.NewStream(ctx, &CurrencyService_ServiceDesc.Streams[0], "/currency.CurrencyService/ListCurrencies", opts...)
	if err != nil {
		return nil, err
	}
	x := &currencyServiceListCurrenciesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CurrencyService_ListCurrenciesClient interface {
	Recv() (*Currency, error)
	grpc.ClientStream
}

type currencyServiceListCurrenciesClient struct {
	grpc.ClientStream
}

func (x *currencyServiceListCurrenciesClient) Recv() (*Currency, error) {
	m := new(Currency)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CurrencyServiceServer is the server API for CurrencyService service.
// All implementations must embed UnimplementedCurrencyServiceServer
// for forward compatibility
type CurrencyServiceServer interface {
	CreateCurrency(context.Context, *Currency) (*Currency, error)
	ReadCurrency(context.Context, *ReadCurrencyRequest) (*Currency, error)
	UpdateCurrency(context.Context, *Currency) (*Currency, error)
	DeleteCurrency(context.Context, *DeleteCurrencyRequest) (*DeleteCurrencyResponse, error)
	ListCurrencies(*Empty, CurrencyService_ListCurrenciesServer) error
	mustEmbedUnimplementedCurrencyServiceServer()
}

// UnimplementedCurrencyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCurrencyServiceServer struct {
}

func (UnimplementedCurrencyServiceServer) CreateCurrency(context.Context, *Currency) (*Currency, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCurrency not implemented")
}
func (UnimplementedCurrencyServiceServer) ReadCurrency(context.Context, *ReadCurrencyRequest) (*Currency, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCurrency not implemented")
}
func (UnimplementedCurrencyServiceServer) UpdateCurrency(context.Context, *Currency) (*Currency, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCurrency not implemented")
}
func (UnimplementedCurrencyServiceServer) DeleteCurrency(context.Context, *DeleteCurrencyRequest) (*DeleteCurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCurrency not implemented")
}
func (UnimplementedCurrencyServiceServer) ListCurrencies(*Empty, CurrencyService_ListCurrenciesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListCurrencies not implemented")
}
func (UnimplementedCurrencyServiceServer) mustEmbedUnimplementedCurrencyServiceServer() {}

// UnsafeCurrencyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CurrencyServiceServer will
// result in compilation errors.
type UnsafeCurrencyServiceServer interface {
	mustEmbedUnimplementedCurrencyServiceServer()
}

func RegisterCurrencyServiceServer(s grpc.ServiceRegistrar, srv CurrencyServiceServer) {
	s.RegisterService(&CurrencyService_ServiceDesc, srv)
}

func _CurrencyService_CreateCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Currency)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).CreateCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/currency.CurrencyService/CreateCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).CreateCurrency(ctx, req.(*Currency))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyService_ReadCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).ReadCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/currency.CurrencyService/ReadCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).ReadCurrency(ctx, req.(*ReadCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyService_UpdateCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Currency)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).UpdateCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/currency.CurrencyService/UpdateCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).UpdateCurrency(ctx, req.(*Currency))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyService_DeleteCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).DeleteCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/currency.CurrencyService/DeleteCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).DeleteCurrency(ctx, req.(*DeleteCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyService_ListCurrencies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CurrencyServiceServer).ListCurrencies(m, &currencyServiceListCurrenciesServer{stream})
}

type CurrencyService_ListCurrenciesServer interface {
	Send(*Currency) error
	grpc.ServerStream
}

type currencyServiceListCurrenciesServer struct {
	grpc.ServerStream
}

func (x *currencyServiceListCurrenciesServer) Send(m *Currency) error {
	return x.ServerStream.SendMsg(m)
}

// CurrencyService_ServiceDesc is the grpc.ServiceDesc for CurrencyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CurrencyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "currency.CurrencyService",
	HandlerType: (*CurrencyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCurrency",
			Handler:    _CurrencyService_CreateCurrency_Handler,
		},
		{
			MethodName: "ReadCurrency",
			Handler:    _CurrencyService_ReadCurrency_Handler,
		},
		{
			MethodName: "UpdateCurrency",
			Handler:    _CurrencyService_UpdateCurrency_Handler,
		},
		{
			MethodName: "DeleteCurrency",
			Handler:    _CurrencyService_DeleteCurrency_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListCurrencies",
			Handler:       _CurrencyService_ListCurrencies_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/currency.proto",
}
