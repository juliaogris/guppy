// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package bank

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

// SupportClient is the client API for Support service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SupportClient interface {
	// ChatCustomer is used by a customer-facing app to send the customer's messages
	// to a chat session. The customer is how initiates and terminates (via "hangup")
	// a chat session. Only customers may invoke this method (e.g. requests must
	// include customer auth credentials).
	ChatCustomer(ctx context.Context, opts ...grpc.CallOption) (Support_ChatCustomerClient, error)
	// ChatAgent is used by an agent-facing app to allow an agent to reply to a
	// customer's messages in a chat session. The agent may accept a chat session,
	// which defaults to the session awaiting an agent for the longest period of time
	// (FIFO queue).
	ChatAgent(ctx context.Context, opts ...grpc.CallOption) (Support_ChatAgentClient, error)
}

type supportClient struct {
	cc grpc.ClientConnInterface
}

func NewSupportClient(cc grpc.ClientConnInterface) SupportClient {
	return &supportClient{cc}
}

func (c *supportClient) ChatCustomer(ctx context.Context, opts ...grpc.CallOption) (Support_ChatCustomerClient, error) {
	stream, err := c.cc.NewStream(ctx, &Support_ServiceDesc.Streams[0], "/bank.Support/ChatCustomer", opts...)
	if err != nil {
		return nil, err
	}
	x := &supportChatCustomerClient{stream}
	return x, nil
}

type Support_ChatCustomerClient interface {
	Send(*ChatCustomerRequest) error
	Recv() (*ChatCustomerResponse, error)
	grpc.ClientStream
}

type supportChatCustomerClient struct {
	grpc.ClientStream
}

func (x *supportChatCustomerClient) Send(m *ChatCustomerRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *supportChatCustomerClient) Recv() (*ChatCustomerResponse, error) {
	m := new(ChatCustomerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *supportClient) ChatAgent(ctx context.Context, opts ...grpc.CallOption) (Support_ChatAgentClient, error) {
	stream, err := c.cc.NewStream(ctx, &Support_ServiceDesc.Streams[1], "/bank.Support/ChatAgent", opts...)
	if err != nil {
		return nil, err
	}
	x := &supportChatAgentClient{stream}
	return x, nil
}

type Support_ChatAgentClient interface {
	Send(*ChatAgentRequest) error
	Recv() (*ChatAgentResponse, error)
	grpc.ClientStream
}

type supportChatAgentClient struct {
	grpc.ClientStream
}

func (x *supportChatAgentClient) Send(m *ChatAgentRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *supportChatAgentClient) Recv() (*ChatAgentResponse, error) {
	m := new(ChatAgentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SupportServer is the server API for Support service.
// All implementations must embed UnimplementedSupportServer
// for forward compatibility
type SupportServer interface {
	// ChatCustomer is used by a customer-facing app to send the customer's messages
	// to a chat session. The customer is how initiates and terminates (via "hangup")
	// a chat session. Only customers may invoke this method (e.g. requests must
	// include customer auth credentials).
	ChatCustomer(Support_ChatCustomerServer) error
	// ChatAgent is used by an agent-facing app to allow an agent to reply to a
	// customer's messages in a chat session. The agent may accept a chat session,
	// which defaults to the session awaiting an agent for the longest period of time
	// (FIFO queue).
	ChatAgent(Support_ChatAgentServer) error
	mustEmbedUnimplementedSupportServer()
}

// UnimplementedSupportServer must be embedded to have forward compatible implementations.
type UnimplementedSupportServer struct {
}

func (UnimplementedSupportServer) ChatCustomer(Support_ChatCustomerServer) error {
	return status.Errorf(codes.Unimplemented, "method ChatCustomer not implemented")
}
func (UnimplementedSupportServer) ChatAgent(Support_ChatAgentServer) error {
	return status.Errorf(codes.Unimplemented, "method ChatAgent not implemented")
}
func (UnimplementedSupportServer) mustEmbedUnimplementedSupportServer() {}

// UnsafeSupportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SupportServer will
// result in compilation errors.
type UnsafeSupportServer interface {
	mustEmbedUnimplementedSupportServer()
}

func RegisterSupportServer(s grpc.ServiceRegistrar, srv SupportServer) {
	s.RegisterService(&Support_ServiceDesc, srv)
}

func _Support_ChatCustomer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SupportServer).ChatCustomer(&supportChatCustomerServer{stream})
}

type Support_ChatCustomerServer interface {
	Send(*ChatCustomerResponse) error
	Recv() (*ChatCustomerRequest, error)
	grpc.ServerStream
}

type supportChatCustomerServer struct {
	grpc.ServerStream
}

func (x *supportChatCustomerServer) Send(m *ChatCustomerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *supportChatCustomerServer) Recv() (*ChatCustomerRequest, error) {
	m := new(ChatCustomerRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Support_ChatAgent_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SupportServer).ChatAgent(&supportChatAgentServer{stream})
}

type Support_ChatAgentServer interface {
	Send(*ChatAgentResponse) error
	Recv() (*ChatAgentRequest, error)
	grpc.ServerStream
}

type supportChatAgentServer struct {
	grpc.ServerStream
}

func (x *supportChatAgentServer) Send(m *ChatAgentResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *supportChatAgentServer) Recv() (*ChatAgentRequest, error) {
	m := new(ChatAgentRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Support_ServiceDesc is the grpc.ServiceDesc for Support service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Support_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bank.Support",
	HandlerType: (*SupportServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ChatCustomer",
			Handler:       _Support_ChatCustomer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ChatAgent",
			Handler:       _Support_ChatAgent_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "bank/support.proto",
}
