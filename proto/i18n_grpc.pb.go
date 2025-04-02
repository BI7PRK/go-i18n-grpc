// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: proto/i18n.proto

package proto

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
	I18NService_CultureFeature_FullMethodName                  = "/i18n.I18nService/CultureFeature"
	I18NService_CulturesResourceTypeFeature_FullMethodName     = "/i18n.I18nService/CulturesResourceTypeFeature"
	I18NService_CulturesResourceKeyFeature_FullMethodName      = "/i18n.I18nService/CulturesResourceKeyFeature"
	I18NService_CulturesResourceKeyValueFeature_FullMethodName = "/i18n.I18nService/CulturesResourceKeyValueFeature"
	I18NService_AddResourceKeyValue_FullMethodName             = "/i18n.I18nService/AddResourceKeyValue"
)

// I18NServiceClient is the client API for I18NService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type I18NServiceClient interface {
	CultureFeature(ctx context.Context, in *CulturesRequest, opts ...grpc.CallOption) (*CulturesReply, error)
	CulturesResourceTypeFeature(ctx context.Context, in *CultureTypesRequest, opts ...grpc.CallOption) (*CulturesTypesReply, error)
	CulturesResourceKeyFeature(ctx context.Context, in *CultureKeysRequest, opts ...grpc.CallOption) (*CultureKeysReply, error)
	CulturesResourceKeyValueFeature(ctx context.Context, in *CultureKeyValuesRequest, opts ...grpc.CallOption) (*CultureKeyValuesReply, error)
	AddResourceKeyValue(ctx context.Context, in *AddCultureKeyValueRequest, opts ...grpc.CallOption) (*CultureBaseReply, error)
}

type i18NServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewI18NServiceClient(cc grpc.ClientConnInterface) I18NServiceClient {
	return &i18NServiceClient{cc}
}

func (c *i18NServiceClient) CultureFeature(ctx context.Context, in *CulturesRequest, opts ...grpc.CallOption) (*CulturesReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CulturesReply)
	err := c.cc.Invoke(ctx, I18NService_CultureFeature_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *i18NServiceClient) CulturesResourceTypeFeature(ctx context.Context, in *CultureTypesRequest, opts ...grpc.CallOption) (*CulturesTypesReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CulturesTypesReply)
	err := c.cc.Invoke(ctx, I18NService_CulturesResourceTypeFeature_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *i18NServiceClient) CulturesResourceKeyFeature(ctx context.Context, in *CultureKeysRequest, opts ...grpc.CallOption) (*CultureKeysReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CultureKeysReply)
	err := c.cc.Invoke(ctx, I18NService_CulturesResourceKeyFeature_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *i18NServiceClient) CulturesResourceKeyValueFeature(ctx context.Context, in *CultureKeyValuesRequest, opts ...grpc.CallOption) (*CultureKeyValuesReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CultureKeyValuesReply)
	err := c.cc.Invoke(ctx, I18NService_CulturesResourceKeyValueFeature_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *i18NServiceClient) AddResourceKeyValue(ctx context.Context, in *AddCultureKeyValueRequest, opts ...grpc.CallOption) (*CultureBaseReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CultureBaseReply)
	err := c.cc.Invoke(ctx, I18NService_AddResourceKeyValue_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// I18NServiceServer is the server API for I18NService service.
// All implementations must embed UnimplementedI18NServiceServer
// for forward compatibility.
type I18NServiceServer interface {
	CultureFeature(context.Context, *CulturesRequest) (*CulturesReply, error)
	CulturesResourceTypeFeature(context.Context, *CultureTypesRequest) (*CulturesTypesReply, error)
	CulturesResourceKeyFeature(context.Context, *CultureKeysRequest) (*CultureKeysReply, error)
	CulturesResourceKeyValueFeature(context.Context, *CultureKeyValuesRequest) (*CultureKeyValuesReply, error)
	AddResourceKeyValue(context.Context, *AddCultureKeyValueRequest) (*CultureBaseReply, error)
	mustEmbedUnimplementedI18NServiceServer()
}

// UnimplementedI18NServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedI18NServiceServer struct{}

func (UnimplementedI18NServiceServer) CultureFeature(context.Context, *CulturesRequest) (*CulturesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CultureFeature not implemented")
}
func (UnimplementedI18NServiceServer) CulturesResourceTypeFeature(context.Context, *CultureTypesRequest) (*CulturesTypesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CulturesResourceTypeFeature not implemented")
}
func (UnimplementedI18NServiceServer) CulturesResourceKeyFeature(context.Context, *CultureKeysRequest) (*CultureKeysReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CulturesResourceKeyFeature not implemented")
}
func (UnimplementedI18NServiceServer) CulturesResourceKeyValueFeature(context.Context, *CultureKeyValuesRequest) (*CultureKeyValuesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CulturesResourceKeyValueFeature not implemented")
}
func (UnimplementedI18NServiceServer) AddResourceKeyValue(context.Context, *AddCultureKeyValueRequest) (*CultureBaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddResourceKeyValue not implemented")
}
func (UnimplementedI18NServiceServer) mustEmbedUnimplementedI18NServiceServer() {}
func (UnimplementedI18NServiceServer) testEmbeddedByValue()                     {}

// UnsafeI18NServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to I18NServiceServer will
// result in compilation errors.
type UnsafeI18NServiceServer interface {
	mustEmbedUnimplementedI18NServiceServer()
}

func RegisterI18NServiceServer(s grpc.ServiceRegistrar, srv I18NServiceServer) {
	// If the following call pancis, it indicates UnimplementedI18NServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&I18NService_ServiceDesc, srv)
}

func _I18NService_CultureFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CulturesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(I18NServiceServer).CultureFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: I18NService_CultureFeature_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(I18NServiceServer).CultureFeature(ctx, req.(*CulturesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _I18NService_CulturesResourceTypeFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CultureTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(I18NServiceServer).CulturesResourceTypeFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: I18NService_CulturesResourceTypeFeature_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(I18NServiceServer).CulturesResourceTypeFeature(ctx, req.(*CultureTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _I18NService_CulturesResourceKeyFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CultureKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(I18NServiceServer).CulturesResourceKeyFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: I18NService_CulturesResourceKeyFeature_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(I18NServiceServer).CulturesResourceKeyFeature(ctx, req.(*CultureKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _I18NService_CulturesResourceKeyValueFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CultureKeyValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(I18NServiceServer).CulturesResourceKeyValueFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: I18NService_CulturesResourceKeyValueFeature_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(I18NServiceServer).CulturesResourceKeyValueFeature(ctx, req.(*CultureKeyValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _I18NService_AddResourceKeyValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCultureKeyValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(I18NServiceServer).AddResourceKeyValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: I18NService_AddResourceKeyValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(I18NServiceServer).AddResourceKeyValue(ctx, req.(*AddCultureKeyValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// I18NService_ServiceDesc is the grpc.ServiceDesc for I18NService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var I18NService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "i18n.I18nService",
	HandlerType: (*I18NServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CultureFeature",
			Handler:    _I18NService_CultureFeature_Handler,
		},
		{
			MethodName: "CulturesResourceTypeFeature",
			Handler:    _I18NService_CulturesResourceTypeFeature_Handler,
		},
		{
			MethodName: "CulturesResourceKeyFeature",
			Handler:    _I18NService_CulturesResourceKeyFeature_Handler,
		},
		{
			MethodName: "CulturesResourceKeyValueFeature",
			Handler:    _I18NService_CulturesResourceKeyValueFeature_Handler,
		},
		{
			MethodName: "AddResourceKeyValue",
			Handler:    _I18NService_AddResourceKeyValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/i18n.proto",
}
