// Code generated by protoc-gen-go-triple. DO NOT EDIT.
// versions:
// - protoc-gen-go-triple v1.0.5
// - protoc             v5.29.0--rc3
// source: pb/dci.proto

package dci

import (
	context "context"
	protocol "dubbo.apache.org/dubbo-go/v3/protocol"
	dubbo3 "dubbo.apache.org/dubbo-go/v3/protocol/dubbo3"
	invocation "dubbo.apache.org/dubbo-go/v3/protocol/invocation"
	grpc_go "github.com/dubbogo/grpc-go"
	codes "github.com/dubbogo/grpc-go/codes"
	metadata "github.com/dubbogo/grpc-go/metadata"
	status "github.com/dubbogo/grpc-go/status"
	common "github.com/dubbogo/triple/pkg/common"
	constant "github.com/dubbogo/triple/pkg/common/constant"
	triple "github.com/dubbogo/triple/pkg/triple"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc_go.SupportPackageIsVersion7

// DciClient is the client API for Dci service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DciClient interface {
	// 通用接口 用于文件上传
	GetUploadUrl(ctx context.Context, in *GetUploadUrlRequest, opts ...grpc_go.CallOption) (*GetUploadUrlResponse, common.ErrorWithAttachment)
	// dci user
	AddDciUser(ctx context.Context, in *AddDciUserRequest, opts ...grpc_go.CallOption) (*AddDciUserResponse, common.ErrorWithAttachment)
	UpdateDciUser(ctx context.Context, in *UpdateDciUserRequest, opts ...grpc_go.CallOption) (*UpdateDciUserResponse, common.ErrorWithAttachment)
	QueryDciUser(ctx context.Context, in *QueryDciUserRequest, opts ...grpc_go.CallOption) (*QueryDciUserResponse, common.ErrorWithAttachment)
	// dci work
	CreateDciPreregistration(ctx context.Context, in *CreateDciPreregistrationRequest, opts ...grpc_go.CallOption) (*CreateDciPreregistrationResponse, common.ErrorWithAttachment)
	QueryDciPreregistration(ctx context.Context, in *QueryDciPreregistrationRequest, opts ...grpc_go.CallOption) (*QueryDciPreregistrationResponse, common.ErrorWithAttachment)
	SubmitDciFeedback(ctx context.Context, in *SubmitDciFeedbackRequest, opts ...grpc_go.CallOption) (*SubmitDciFeedbackResponse, common.ErrorWithAttachment)
	QueryDciFeedback(ctx context.Context, in *QueryDciFeedbackRequest, opts ...grpc_go.CallOption) (*QueryDciFeedbackResponse, common.ErrorWithAttachment)
	// 数登
	CreateDciRegistration(ctx context.Context, in *CreateDciRegistrationRequest, opts ...grpc_go.CallOption) (*CreateDciRegistrationResponse, common.ErrorWithAttachment)
	QueryDciRegistration(ctx context.Context, in *QueryDciRegistrationRequest, opts ...grpc_go.CallOption) (*QueryDciRegistrationResponse, common.ErrorWithAttachment)
	GetDciPayUrl(ctx context.Context, in *GetDciPayUrlRequest, opts ...grpc_go.CallOption) (*GetDciPayUrlResponse, common.ErrorWithAttachment)
	QueryDciPay(ctx context.Context, in *QueryDciPayRequest, opts ...grpc_go.CallOption) (*QueryDciPayResponse, common.ErrorWithAttachment)
	GetDciRegistrationcert(ctx context.Context, in *GetDciRegistrationcertRequest, opts ...grpc_go.CallOption) (*GetDciRegistrationcertResponse, common.ErrorWithAttachment)
	RetryDciRegistration(ctx context.Context, in *RetryDciRegistrationRequest, opts ...grpc_go.CallOption) (*RetryDciRegistrationResponse, common.ErrorWithAttachment)
	CloseDciRegistration(ctx context.Context, in *CloseDciRegistrationRequest, opts ...grpc_go.CallOption) (*CloseDciRegistrationResponse, common.ErrorWithAttachment)
}

type dciClient struct {
	cc *triple.TripleConn
}

type DciClientImpl struct {
	GetUploadUrl             func(ctx context.Context, in *GetUploadUrlRequest) (*GetUploadUrlResponse, error)
	AddDciUser               func(ctx context.Context, in *AddDciUserRequest) (*AddDciUserResponse, error)
	UpdateDciUser            func(ctx context.Context, in *UpdateDciUserRequest) (*UpdateDciUserResponse, error)
	QueryDciUser             func(ctx context.Context, in *QueryDciUserRequest) (*QueryDciUserResponse, error)
	CreateDciPreregistration func(ctx context.Context, in *CreateDciPreregistrationRequest) (*CreateDciPreregistrationResponse, error)
	QueryDciPreregistration  func(ctx context.Context, in *QueryDciPreregistrationRequest) (*QueryDciPreregistrationResponse, error)
	SubmitDciFeedback        func(ctx context.Context, in *SubmitDciFeedbackRequest) (*SubmitDciFeedbackResponse, error)
	QueryDciFeedback         func(ctx context.Context, in *QueryDciFeedbackRequest) (*QueryDciFeedbackResponse, error)
	CreateDciRegistration    func(ctx context.Context, in *CreateDciRegistrationRequest) (*CreateDciRegistrationResponse, error)
	QueryDciRegistration     func(ctx context.Context, in *QueryDciRegistrationRequest) (*QueryDciRegistrationResponse, error)
	GetDciPayUrl             func(ctx context.Context, in *GetDciPayUrlRequest) (*GetDciPayUrlResponse, error)
	QueryDciPay              func(ctx context.Context, in *QueryDciPayRequest) (*QueryDciPayResponse, error)
	GetDciRegistrationcert   func(ctx context.Context, in *GetDciRegistrationcertRequest) (*GetDciRegistrationcertResponse, error)
	RetryDciRegistration     func(ctx context.Context, in *RetryDciRegistrationRequest) (*RetryDciRegistrationResponse, error)
	CloseDciRegistration     func(ctx context.Context, in *CloseDciRegistrationRequest) (*CloseDciRegistrationResponse, error)
}

func (c *DciClientImpl) GetDubboStub(cc *triple.TripleConn) DciClient {
	return NewDciClient(cc)
}

func (c *DciClientImpl) XXX_InterfaceName() string {
	return "dci.Dci"
}

func NewDciClient(cc *triple.TripleConn) DciClient {
	return &dciClient{cc}
}

func (c *dciClient) GetUploadUrl(ctx context.Context, in *GetUploadUrlRequest, opts ...grpc_go.CallOption) (*GetUploadUrlResponse, common.ErrorWithAttachment) {
	out := new(GetUploadUrlResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/GetUploadUrl", in, out)
}

func (c *dciClient) AddDciUser(ctx context.Context, in *AddDciUserRequest, opts ...grpc_go.CallOption) (*AddDciUserResponse, common.ErrorWithAttachment) {
	out := new(AddDciUserResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/AddDciUser", in, out)
}

func (c *dciClient) UpdateDciUser(ctx context.Context, in *UpdateDciUserRequest, opts ...grpc_go.CallOption) (*UpdateDciUserResponse, common.ErrorWithAttachment) {
	out := new(UpdateDciUserResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/UpdateDciUser", in, out)
}

func (c *dciClient) QueryDciUser(ctx context.Context, in *QueryDciUserRequest, opts ...grpc_go.CallOption) (*QueryDciUserResponse, common.ErrorWithAttachment) {
	out := new(QueryDciUserResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/QueryDciUser", in, out)
}

func (c *dciClient) CreateDciPreregistration(ctx context.Context, in *CreateDciPreregistrationRequest, opts ...grpc_go.CallOption) (*CreateDciPreregistrationResponse, common.ErrorWithAttachment) {
	out := new(CreateDciPreregistrationResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/CreateDciPreregistration", in, out)
}

func (c *dciClient) QueryDciPreregistration(ctx context.Context, in *QueryDciPreregistrationRequest, opts ...grpc_go.CallOption) (*QueryDciPreregistrationResponse, common.ErrorWithAttachment) {
	out := new(QueryDciPreregistrationResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/QueryDciPreregistration", in, out)
}

func (c *dciClient) SubmitDciFeedback(ctx context.Context, in *SubmitDciFeedbackRequest, opts ...grpc_go.CallOption) (*SubmitDciFeedbackResponse, common.ErrorWithAttachment) {
	out := new(SubmitDciFeedbackResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/SubmitDciFeedback", in, out)
}

func (c *dciClient) QueryDciFeedback(ctx context.Context, in *QueryDciFeedbackRequest, opts ...grpc_go.CallOption) (*QueryDciFeedbackResponse, common.ErrorWithAttachment) {
	out := new(QueryDciFeedbackResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/QueryDciFeedback", in, out)
}

func (c *dciClient) CreateDciRegistration(ctx context.Context, in *CreateDciRegistrationRequest, opts ...grpc_go.CallOption) (*CreateDciRegistrationResponse, common.ErrorWithAttachment) {
	out := new(CreateDciRegistrationResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/CreateDciRegistration", in, out)
}

func (c *dciClient) QueryDciRegistration(ctx context.Context, in *QueryDciRegistrationRequest, opts ...grpc_go.CallOption) (*QueryDciRegistrationResponse, common.ErrorWithAttachment) {
	out := new(QueryDciRegistrationResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/QueryDciRegistration", in, out)
}

func (c *dciClient) GetDciPayUrl(ctx context.Context, in *GetDciPayUrlRequest, opts ...grpc_go.CallOption) (*GetDciPayUrlResponse, common.ErrorWithAttachment) {
	out := new(GetDciPayUrlResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/GetDciPayUrl", in, out)
}

func (c *dciClient) QueryDciPay(ctx context.Context, in *QueryDciPayRequest, opts ...grpc_go.CallOption) (*QueryDciPayResponse, common.ErrorWithAttachment) {
	out := new(QueryDciPayResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/QueryDciPay", in, out)
}

func (c *dciClient) GetDciRegistrationcert(ctx context.Context, in *GetDciRegistrationcertRequest, opts ...grpc_go.CallOption) (*GetDciRegistrationcertResponse, common.ErrorWithAttachment) {
	out := new(GetDciRegistrationcertResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/GetDciRegistrationcert", in, out)
}

func (c *dciClient) RetryDciRegistration(ctx context.Context, in *RetryDciRegistrationRequest, opts ...grpc_go.CallOption) (*RetryDciRegistrationResponse, common.ErrorWithAttachment) {
	out := new(RetryDciRegistrationResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/RetryDciRegistration", in, out)
}

func (c *dciClient) CloseDciRegistration(ctx context.Context, in *CloseDciRegistrationRequest, opts ...grpc_go.CallOption) (*CloseDciRegistrationResponse, common.ErrorWithAttachment) {
	out := new(CloseDciRegistrationResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/CloseDciRegistration", in, out)
}

// DciServer is the server API for Dci service.
// All implementations must embed UnimplementedDciServer
// for forward compatibility
type DciServer interface {
	// 通用接口 用于文件上传
	GetUploadUrl(context.Context, *GetUploadUrlRequest) (*GetUploadUrlResponse, error)
	// dci user
	AddDciUser(context.Context, *AddDciUserRequest) (*AddDciUserResponse, error)
	UpdateDciUser(context.Context, *UpdateDciUserRequest) (*UpdateDciUserResponse, error)
	QueryDciUser(context.Context, *QueryDciUserRequest) (*QueryDciUserResponse, error)
	// dci work
	CreateDciPreregistration(context.Context, *CreateDciPreregistrationRequest) (*CreateDciPreregistrationResponse, error)
	QueryDciPreregistration(context.Context, *QueryDciPreregistrationRequest) (*QueryDciPreregistrationResponse, error)
	SubmitDciFeedback(context.Context, *SubmitDciFeedbackRequest) (*SubmitDciFeedbackResponse, error)
	QueryDciFeedback(context.Context, *QueryDciFeedbackRequest) (*QueryDciFeedbackResponse, error)
	// 数登
	CreateDciRegistration(context.Context, *CreateDciRegistrationRequest) (*CreateDciRegistrationResponse, error)
	QueryDciRegistration(context.Context, *QueryDciRegistrationRequest) (*QueryDciRegistrationResponse, error)
	GetDciPayUrl(context.Context, *GetDciPayUrlRequest) (*GetDciPayUrlResponse, error)
	QueryDciPay(context.Context, *QueryDciPayRequest) (*QueryDciPayResponse, error)
	GetDciRegistrationcert(context.Context, *GetDciRegistrationcertRequest) (*GetDciRegistrationcertResponse, error)
	RetryDciRegistration(context.Context, *RetryDciRegistrationRequest) (*RetryDciRegistrationResponse, error)
	CloseDciRegistration(context.Context, *CloseDciRegistrationRequest) (*CloseDciRegistrationResponse, error)
	mustEmbedUnimplementedDciServer()
}

// UnimplementedDciServer must be embedded to have forward compatible implementations.
type UnimplementedDciServer struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedDciServer) GetUploadUrl(context.Context, *GetUploadUrlRequest) (*GetUploadUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUploadUrl not implemented")
}
func (UnimplementedDciServer) AddDciUser(context.Context, *AddDciUserRequest) (*AddDciUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDciUser not implemented")
}
func (UnimplementedDciServer) UpdateDciUser(context.Context, *UpdateDciUserRequest) (*UpdateDciUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDciUser not implemented")
}
func (UnimplementedDciServer) QueryDciUser(context.Context, *QueryDciUserRequest) (*QueryDciUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDciUser not implemented")
}
func (UnimplementedDciServer) CreateDciPreregistration(context.Context, *CreateDciPreregistrationRequest) (*CreateDciPreregistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDciPreregistration not implemented")
}
func (UnimplementedDciServer) QueryDciPreregistration(context.Context, *QueryDciPreregistrationRequest) (*QueryDciPreregistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDciPreregistration not implemented")
}
func (UnimplementedDciServer) SubmitDciFeedback(context.Context, *SubmitDciFeedbackRequest) (*SubmitDciFeedbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitDciFeedback not implemented")
}
func (UnimplementedDciServer) QueryDciFeedback(context.Context, *QueryDciFeedbackRequest) (*QueryDciFeedbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDciFeedback not implemented")
}
func (UnimplementedDciServer) CreateDciRegistration(context.Context, *CreateDciRegistrationRequest) (*CreateDciRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDciRegistration not implemented")
}
func (UnimplementedDciServer) QueryDciRegistration(context.Context, *QueryDciRegistrationRequest) (*QueryDciRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDciRegistration not implemented")
}
func (UnimplementedDciServer) GetDciPayUrl(context.Context, *GetDciPayUrlRequest) (*GetDciPayUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDciPayUrl not implemented")
}
func (UnimplementedDciServer) QueryDciPay(context.Context, *QueryDciPayRequest) (*QueryDciPayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDciPay not implemented")
}
func (UnimplementedDciServer) GetDciRegistrationcert(context.Context, *GetDciRegistrationcertRequest) (*GetDciRegistrationcertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDciRegistrationcert not implemented")
}
func (UnimplementedDciServer) RetryDciRegistration(context.Context, *RetryDciRegistrationRequest) (*RetryDciRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetryDciRegistration not implemented")
}
func (UnimplementedDciServer) CloseDciRegistration(context.Context, *CloseDciRegistrationRequest) (*CloseDciRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseDciRegistration not implemented")
}
func (s *UnimplementedDciServer) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedDciServer) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (s *UnimplementedDciServer) XXX_ServiceDesc() *grpc_go.ServiceDesc {
	return &Dci_ServiceDesc
}
func (s *UnimplementedDciServer) XXX_InterfaceName() string {
	return "dci.Dci"
}

func (UnimplementedDciServer) mustEmbedUnimplementedDciServer() {}

// UnsafeDciServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DciServer will
// result in compilation errors.
type UnsafeDciServer interface {
	mustEmbedUnimplementedDciServer()
}

func RegisterDciServer(s grpc_go.ServiceRegistrar, srv DciServer) {
	s.RegisterService(&Dci_ServiceDesc, srv)
}

func _Dci_GetUploadUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUploadUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("GetUploadUrl", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _Dci_AddDciUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDciUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("AddDciUser", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _Dci_UpdateDciUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDciUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("UpdateDciUser", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _Dci_QueryDciUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDciUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("QueryDciUser", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _Dci_CreateDciPreregistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDciPreregistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("CreateDciPreregistration", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _Dci_QueryDciPreregistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDciPreregistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("QueryDciPreregistration", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _Dci_SubmitDciFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitDciFeedbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("SubmitDciFeedback", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _Dci_QueryDciFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDciFeedbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("QueryDciFeedback", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _Dci_CreateDciRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDciRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("CreateDciRegistration", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _Dci_QueryDciRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDciRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("QueryDciRegistration", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _Dci_GetDciPayUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDciPayUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("GetDciPayUrl", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _Dci_QueryDciPay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDciPayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("QueryDciPay", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _Dci_GetDciRegistrationcert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDciRegistrationcertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("GetDciRegistrationcert", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _Dci_RetryDciRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetryDciRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("RetryDciRegistration", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _Dci_CloseDciRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseDciRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("CloseDciRegistration", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

// Dci_ServiceDesc is the grpc_go.ServiceDesc for Dci service.
// It's only intended for direct use with grpc_go.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dci_ServiceDesc = grpc_go.ServiceDesc{
	ServiceName: "dci.Dci",
	HandlerType: (*DciServer)(nil),
	Methods: []grpc_go.MethodDesc{
		{
			MethodName: "GetUploadUrl",
			Handler:    _Dci_GetUploadUrl_Handler,
		},
		{
			MethodName: "AddDciUser",
			Handler:    _Dci_AddDciUser_Handler,
		},
		{
			MethodName: "UpdateDciUser",
			Handler:    _Dci_UpdateDciUser_Handler,
		},
		{
			MethodName: "QueryDciUser",
			Handler:    _Dci_QueryDciUser_Handler,
		},
		{
			MethodName: "CreateDciPreregistration",
			Handler:    _Dci_CreateDciPreregistration_Handler,
		},
		{
			MethodName: "QueryDciPreregistration",
			Handler:    _Dci_QueryDciPreregistration_Handler,
		},
		{
			MethodName: "SubmitDciFeedback",
			Handler:    _Dci_SubmitDciFeedback_Handler,
		},
		{
			MethodName: "QueryDciFeedback",
			Handler:    _Dci_QueryDciFeedback_Handler,
		},
		{
			MethodName: "CreateDciRegistration",
			Handler:    _Dci_CreateDciRegistration_Handler,
		},
		{
			MethodName: "QueryDciRegistration",
			Handler:    _Dci_QueryDciRegistration_Handler,
		},
		{
			MethodName: "GetDciPayUrl",
			Handler:    _Dci_GetDciPayUrl_Handler,
		},
		{
			MethodName: "QueryDciPay",
			Handler:    _Dci_QueryDciPay_Handler,
		},
		{
			MethodName: "GetDciRegistrationcert",
			Handler:    _Dci_GetDciRegistrationcert_Handler,
		},
		{
			MethodName: "RetryDciRegistration",
			Handler:    _Dci_RetryDciRegistration_Handler,
		},
		{
			MethodName: "CloseDciRegistration",
			Handler:    _Dci_CloseDciRegistration_Handler,
		},
	},
	Streams:  []grpc_go.StreamDesc{},
	Metadata: "pb/dci.proto",
}
