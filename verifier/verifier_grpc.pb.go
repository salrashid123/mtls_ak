// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.1
// source: verifier/verifier.proto

package verifier

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

const (
	Verifier_GetEK_FullMethodName    = "/verifier.Verifier/GetEK"
	Verifier_GetAK_FullMethodName    = "/verifier.Verifier/GetAK"
	Verifier_Attest_FullMethodName   = "/verifier.Verifier/Attest"
	Verifier_Quote_FullMethodName    = "/verifier.Verifier/Quote"
	Verifier_NewKey_FullMethodName   = "/verifier.Verifier/NewKey"
	Verifier_Sign_FullMethodName     = "/verifier.Verifier/Sign"
	Verifier_StartTLS_FullMethodName = "/verifier.Verifier/StartTLS"
)

// VerifierClient is the client API for Verifier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VerifierClient interface {
	GetEK(ctx context.Context, in *GetEKRequest, opts ...grpc.CallOption) (*GetEKResponse, error)
	GetAK(ctx context.Context, in *GetAKRequest, opts ...grpc.CallOption) (*GetAKResponse, error)
	Attest(ctx context.Context, in *AttestRequest, opts ...grpc.CallOption) (*AttestResponse, error)
	Quote(ctx context.Context, in *QuoteRequest, opts ...grpc.CallOption) (*QuoteResponse, error)
	NewKey(ctx context.Context, in *NewKeyRequest, opts ...grpc.CallOption) (*NewKeyResponse, error)
	Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error)
	StartTLS(ctx context.Context, in *StartTLSRequest, opts ...grpc.CallOption) (*StartTLSResponse, error)
}

type verifierClient struct {
	cc grpc.ClientConnInterface
}

func NewVerifierClient(cc grpc.ClientConnInterface) VerifierClient {
	return &verifierClient{cc}
}

func (c *verifierClient) GetEK(ctx context.Context, in *GetEKRequest, opts ...grpc.CallOption) (*GetEKResponse, error) {
	out := new(GetEKResponse)
	err := c.cc.Invoke(ctx, Verifier_GetEK_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifierClient) GetAK(ctx context.Context, in *GetAKRequest, opts ...grpc.CallOption) (*GetAKResponse, error) {
	out := new(GetAKResponse)
	err := c.cc.Invoke(ctx, Verifier_GetAK_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifierClient) Attest(ctx context.Context, in *AttestRequest, opts ...grpc.CallOption) (*AttestResponse, error) {
	out := new(AttestResponse)
	err := c.cc.Invoke(ctx, Verifier_Attest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifierClient) Quote(ctx context.Context, in *QuoteRequest, opts ...grpc.CallOption) (*QuoteResponse, error) {
	out := new(QuoteResponse)
	err := c.cc.Invoke(ctx, Verifier_Quote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifierClient) NewKey(ctx context.Context, in *NewKeyRequest, opts ...grpc.CallOption) (*NewKeyResponse, error) {
	out := new(NewKeyResponse)
	err := c.cc.Invoke(ctx, Verifier_NewKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifierClient) Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error) {
	out := new(SignResponse)
	err := c.cc.Invoke(ctx, Verifier_Sign_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifierClient) StartTLS(ctx context.Context, in *StartTLSRequest, opts ...grpc.CallOption) (*StartTLSResponse, error) {
	out := new(StartTLSResponse)
	err := c.cc.Invoke(ctx, Verifier_StartTLS_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VerifierServer is the server API for Verifier service.
// All implementations should embed UnimplementedVerifierServer
// for forward compatibility
type VerifierServer interface {
	GetEK(context.Context, *GetEKRequest) (*GetEKResponse, error)
	GetAK(context.Context, *GetAKRequest) (*GetAKResponse, error)
	Attest(context.Context, *AttestRequest) (*AttestResponse, error)
	Quote(context.Context, *QuoteRequest) (*QuoteResponse, error)
	NewKey(context.Context, *NewKeyRequest) (*NewKeyResponse, error)
	Sign(context.Context, *SignRequest) (*SignResponse, error)
	StartTLS(context.Context, *StartTLSRequest) (*StartTLSResponse, error)
}

// UnimplementedVerifierServer should be embedded to have forward compatible implementations.
type UnimplementedVerifierServer struct {
}

func (UnimplementedVerifierServer) GetEK(context.Context, *GetEKRequest) (*GetEKResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEK not implemented")
}
func (UnimplementedVerifierServer) GetAK(context.Context, *GetAKRequest) (*GetAKResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAK not implemented")
}
func (UnimplementedVerifierServer) Attest(context.Context, *AttestRequest) (*AttestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Attest not implemented")
}
func (UnimplementedVerifierServer) Quote(context.Context, *QuoteRequest) (*QuoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Quote not implemented")
}
func (UnimplementedVerifierServer) NewKey(context.Context, *NewKeyRequest) (*NewKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewKey not implemented")
}
func (UnimplementedVerifierServer) Sign(context.Context, *SignRequest) (*SignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}
func (UnimplementedVerifierServer) StartTLS(context.Context, *StartTLSRequest) (*StartTLSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartTLS not implemented")
}

// UnsafeVerifierServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VerifierServer will
// result in compilation errors.
type UnsafeVerifierServer interface {
	mustEmbedUnimplementedVerifierServer()
}

func RegisterVerifierServer(s grpc.ServiceRegistrar, srv VerifierServer) {
	s.RegisterService(&Verifier_ServiceDesc, srv)
}

func _Verifier_GetEK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEKRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifierServer).GetEK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Verifier_GetEK_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifierServer).GetEK(ctx, req.(*GetEKRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Verifier_GetAK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAKRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifierServer).GetAK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Verifier_GetAK_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifierServer).GetAK(ctx, req.(*GetAKRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Verifier_Attest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifierServer).Attest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Verifier_Attest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifierServer).Attest(ctx, req.(*AttestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Verifier_Quote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifierServer).Quote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Verifier_Quote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifierServer).Quote(ctx, req.(*QuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Verifier_NewKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifierServer).NewKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Verifier_NewKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifierServer).NewKey(ctx, req.(*NewKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Verifier_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifierServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Verifier_Sign_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifierServer).Sign(ctx, req.(*SignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Verifier_StartTLS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartTLSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifierServer).StartTLS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Verifier_StartTLS_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifierServer).StartTLS(ctx, req.(*StartTLSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Verifier_ServiceDesc is the grpc.ServiceDesc for Verifier service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Verifier_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "verifier.Verifier",
	HandlerType: (*VerifierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEK",
			Handler:    _Verifier_GetEK_Handler,
		},
		{
			MethodName: "GetAK",
			Handler:    _Verifier_GetAK_Handler,
		},
		{
			MethodName: "Attest",
			Handler:    _Verifier_Attest_Handler,
		},
		{
			MethodName: "Quote",
			Handler:    _Verifier_Quote_Handler,
		},
		{
			MethodName: "NewKey",
			Handler:    _Verifier_NewKey_Handler,
		},
		{
			MethodName: "Sign",
			Handler:    _Verifier_Sign_Handler,
		},
		{
			MethodName: "StartTLS",
			Handler:    _Verifier_StartTLS_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "verifier/verifier.proto",
}
