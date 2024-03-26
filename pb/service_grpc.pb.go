// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: pb/service.proto

package pb

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

// GossipClient is the client API for Gossip service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GossipClient interface {
	ShakeHands(ctx context.Context, in *Handshake, opts ...grpc.CallOption) (*Handshake, error)
	TakeSeat(ctx context.Context, in *TakeSeatMsg, opts ...grpc.CallOption) (*Ack, error)
	ShuffleAndEncrypt(ctx context.Context, in *ShuffleAndEncryptMsg, opts ...grpc.CallOption) (*Ack, error)
	SetGameStatus(ctx context.Context, in *SetGameStatusMsg, opts ...grpc.CallOption) (*Ack, error)
	TakeAction(ctx context.Context, in *TakeActionMsg, opts ...grpc.CallOption) (*Ack, error)
}

type gossipClient struct {
	cc grpc.ClientConnInterface
}

func NewGossipClient(cc grpc.ClientConnInterface) GossipClient {
	return &gossipClient{cc}
}

func (c *gossipClient) ShakeHands(ctx context.Context, in *Handshake, opts ...grpc.CallOption) (*Handshake, error) {
	out := new(Handshake)
	err := c.cc.Invoke(ctx, "/Gossip/ShakeHands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gossipClient) TakeSeat(ctx context.Context, in *TakeSeatMsg, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/Gossip/TakeSeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gossipClient) ShuffleAndEncrypt(ctx context.Context, in *ShuffleAndEncryptMsg, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/Gossip/ShuffleAndEncrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gossipClient) SetGameStatus(ctx context.Context, in *SetGameStatusMsg, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/Gossip/SetGameStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gossipClient) TakeAction(ctx context.Context, in *TakeActionMsg, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/Gossip/TakeAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GossipServer is the server API for Gossip service.
// All implementations must embed UnimplementedGossipServer
// for forward compatibility
type GossipServer interface {
	ShakeHands(context.Context, *Handshake) (*Handshake, error)
	TakeSeat(context.Context, *TakeSeatMsg) (*Ack, error)
	ShuffleAndEncrypt(context.Context, *ShuffleAndEncryptMsg) (*Ack, error)
	SetGameStatus(context.Context, *SetGameStatusMsg) (*Ack, error)
	TakeAction(context.Context, *TakeActionMsg) (*Ack, error)
	mustEmbedUnimplementedGossipServer()
}

// UnimplementedGossipServer must be embedded to have forward compatible implementations.
type UnimplementedGossipServer struct {
}

func (UnimplementedGossipServer) ShakeHands(context.Context, *Handshake) (*Handshake, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShakeHands not implemented")
}
func (UnimplementedGossipServer) TakeSeat(context.Context, *TakeSeatMsg) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TakeSeat not implemented")
}
func (UnimplementedGossipServer) ShuffleAndEncrypt(context.Context, *ShuffleAndEncryptMsg) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShuffleAndEncrypt not implemented")
}
func (UnimplementedGossipServer) SetGameStatus(context.Context, *SetGameStatusMsg) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGameStatus not implemented")
}
func (UnimplementedGossipServer) TakeAction(context.Context, *TakeActionMsg) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TakeAction not implemented")
}
func (UnimplementedGossipServer) mustEmbedUnimplementedGossipServer() {}

// UnsafeGossipServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GossipServer will
// result in compilation errors.
type UnsafeGossipServer interface {
	mustEmbedUnimplementedGossipServer()
}

func RegisterGossipServer(s grpc.ServiceRegistrar, srv GossipServer) {
	s.RegisterService(&Gossip_ServiceDesc, srv)
}

func _Gossip_ShakeHands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Handshake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GossipServer).ShakeHands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Gossip/ShakeHands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GossipServer).ShakeHands(ctx, req.(*Handshake))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gossip_TakeSeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TakeSeatMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GossipServer).TakeSeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Gossip/TakeSeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GossipServer).TakeSeat(ctx, req.(*TakeSeatMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gossip_ShuffleAndEncrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShuffleAndEncryptMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GossipServer).ShuffleAndEncrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Gossip/ShuffleAndEncrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GossipServer).ShuffleAndEncrypt(ctx, req.(*ShuffleAndEncryptMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gossip_SetGameStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGameStatusMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GossipServer).SetGameStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Gossip/SetGameStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GossipServer).SetGameStatus(ctx, req.(*SetGameStatusMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gossip_TakeAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TakeActionMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GossipServer).TakeAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Gossip/TakeAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GossipServer).TakeAction(ctx, req.(*TakeActionMsg))
	}
	return interceptor(ctx, in, info, handler)
}

// Gossip_ServiceDesc is the grpc.ServiceDesc for Gossip service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gossip_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Gossip",
	HandlerType: (*GossipServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShakeHands",
			Handler:    _Gossip_ShakeHands_Handler,
		},
		{
			MethodName: "TakeSeat",
			Handler:    _Gossip_TakeSeat_Handler,
		},
		{
			MethodName: "ShuffleAndEncrypt",
			Handler:    _Gossip_ShuffleAndEncrypt_Handler,
		},
		{
			MethodName: "SetGameStatus",
			Handler:    _Gossip_SetGameStatus_Handler,
		},
		{
			MethodName: "TakeAction",
			Handler:    _Gossip_TakeAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/service.proto",
}
