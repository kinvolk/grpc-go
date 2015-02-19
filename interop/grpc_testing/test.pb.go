/*
 *
 * Copyright 2014, Google Inc.
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 *     * Redistributions of source code must retain the above copyright
 * notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above
 * copyright notice, this list of conditions and the following disclaimer
 * in the documentation and/or other materials provided with the
 * distribution.
 *     * Neither the name of Google Inc. nor the names of its
 * contributors may be used to endorse or promote products derived from
 * this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 * A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */

// Code generated by protoc-gen-go.
// source: google.golang.org/grpc/test/grpc_testing/test.proto
// DO NOT EDIT!

/*
Package grpc_testing is a generated protocol buffer package.

It is generated from these files:
	third_party/golang/grpc/test/grpc_testing/test.proto

It has these top-level messages:
	Empty
	Payload
	SimpleRequest
	SimpleResponse
	StreamingInputCallRequest
	StreamingInputCallResponse
	ResponseParameters
	StreamingOutputCallRequest
	StreamingOutputCallResponse
*/
package grpc_testing

import proto "github.com/golang/protobuf/proto"
import math "math"

import (
	errors "errors"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = errors.New
var _ = io.EOF
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// The type of payload that should be returned.
type PayloadType int32

const (
	// Compressable text format.
	PayloadType_COMPRESSABLE PayloadType = 0
	// Uncompressable binary format.
	PayloadType_UNCOMPRESSABLE PayloadType = 1
	// Randomly chosen from all other formats defined in this enum.
	PayloadType_RANDOM PayloadType = 2
)

var PayloadType_name = map[int32]string{
	0: "COMPRESSABLE",
	1: "UNCOMPRESSABLE",
	2: "RANDOM",
}
var PayloadType_value = map[string]int32{
	"COMPRESSABLE":   0,
	"UNCOMPRESSABLE": 1,
	"RANDOM":         2,
}

func (x PayloadType) Enum() *PayloadType {
	p := new(PayloadType)
	*p = x
	return p
}
func (x PayloadType) String() string {
	return proto.EnumName(PayloadType_name, int32(x))
}
func (x *PayloadType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PayloadType_value, data, "PayloadType")
	if err != nil {
		return err
	}
	*x = PayloadType(value)
	return nil
}
func (PayloadType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Empty struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// A block of data, to simply increase gRPC message size.
type Payload struct {
	// The type of data in body.
	Type *PayloadType `protobuf:"varint,1,opt,name=type,enum=grpc.testing.PayloadType" json:"type,omitempty"`
	// Primary contents of payload.
	Body             []byte `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Payload) GetType() PayloadType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return PayloadType_COMPRESSABLE
}

func (m *Payload) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

// Unary request.
type SimpleRequest struct {
	// Desired payload type in the response from the server.
	// If response_type is RANDOM, server randomly chooses one from other formats.
	ResponseType *PayloadType `protobuf:"varint,1,opt,name=response_type,enum=grpc.testing.PayloadType" json:"response_type,omitempty"`
	// Desired payload size in the response from the server.
	// If response_type is COMPRESSABLE, this denotes the size before compression.
	ResponseSize *int32 `protobuf:"varint,2,opt,name=response_size" json:"response_size,omitempty"`
	// Optional input payload sent along with the request.
	Payload *Payload `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
	// Whether SimpleResponse should include username.
	FillUsername *bool `protobuf:"varint,4,opt,name=fill_username" json:"fill_username,omitempty"`
	// Whether SimpleResponse should include OAuth scope.
	FillOauthScope   *bool  `protobuf:"varint,5,opt,name=fill_oauth_scope" json:"fill_oauth_scope,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SimpleRequest) Reset()                    { *m = SimpleRequest{} }
func (m *SimpleRequest) String() string            { return proto.CompactTextString(m) }
func (*SimpleRequest) ProtoMessage()               {}
func (*SimpleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SimpleRequest) GetResponseType() PayloadType {
	if m != nil && m.ResponseType != nil {
		return *m.ResponseType
	}
	return PayloadType_COMPRESSABLE
}

func (m *SimpleRequest) GetResponseSize() int32 {
	if m != nil && m.ResponseSize != nil {
		return *m.ResponseSize
	}
	return 0
}

func (m *SimpleRequest) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *SimpleRequest) GetFillUsername() bool {
	if m != nil && m.FillUsername != nil {
		return *m.FillUsername
	}
	return false
}

func (m *SimpleRequest) GetFillOauthScope() bool {
	if m != nil && m.FillOauthScope != nil {
		return *m.FillOauthScope
	}
	return false
}

// Unary response, as configured by the request.
type SimpleResponse struct {
	// Payload to increase message size.
	Payload *Payload `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
	// The user the request came from, for verifying authentication was
	// successful when the client expected it.
	Username *string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	// OAuth scope.
	OauthScope       *string `protobuf:"bytes,3,opt,name=oauth_scope" json:"oauth_scope,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SimpleResponse) Reset()                    { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string            { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()               {}
func (*SimpleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SimpleResponse) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *SimpleResponse) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *SimpleResponse) GetOauthScope() string {
	if m != nil && m.OauthScope != nil {
		return *m.OauthScope
	}
	return ""
}

// Client-streaming request.
type StreamingInputCallRequest struct {
	// Optional input payload sent along with the request.
	Payload          *Payload `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *StreamingInputCallRequest) Reset()                    { *m = StreamingInputCallRequest{} }
func (m *StreamingInputCallRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamingInputCallRequest) ProtoMessage()               {}
func (*StreamingInputCallRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StreamingInputCallRequest) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Client-streaming response.
type StreamingInputCallResponse struct {
	// Aggregated size of payloads received from the client.
	AggregatedPayloadSize *int32 `protobuf:"varint,1,opt,name=aggregated_payload_size" json:"aggregated_payload_size,omitempty"`
	XXX_unrecognized      []byte `json:"-"`
}

func (m *StreamingInputCallResponse) Reset()                    { *m = StreamingInputCallResponse{} }
func (m *StreamingInputCallResponse) String() string            { return proto.CompactTextString(m) }
func (*StreamingInputCallResponse) ProtoMessage()               {}
func (*StreamingInputCallResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *StreamingInputCallResponse) GetAggregatedPayloadSize() int32 {
	if m != nil && m.AggregatedPayloadSize != nil {
		return *m.AggregatedPayloadSize
	}
	return 0
}

// Configuration for a particular response.
type ResponseParameters struct {
	// Desired payload sizes in responses from the server.
	// If response_type is COMPRESSABLE, this denotes the size before compression.
	Size *int32 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	// Desired interval between consecutive responses in the response stream in
	// microseconds.
	IntervalUs       *int32 `protobuf:"varint,2,opt,name=interval_us" json:"interval_us,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ResponseParameters) Reset()                    { *m = ResponseParameters{} }
func (m *ResponseParameters) String() string            { return proto.CompactTextString(m) }
func (*ResponseParameters) ProtoMessage()               {}
func (*ResponseParameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ResponseParameters) GetSize() int32 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

func (m *ResponseParameters) GetIntervalUs() int32 {
	if m != nil && m.IntervalUs != nil {
		return *m.IntervalUs
	}
	return 0
}

// Server-streaming request.
type StreamingOutputCallRequest struct {
	// Desired payload type in the response from the server.
	// If response_type is RANDOM, the payload from each response in the stream
	// might be of different types. This is to simulate a mixed type of payload
	// stream.
	ResponseType *PayloadType `protobuf:"varint,1,opt,name=response_type,enum=grpc.testing.PayloadType" json:"response_type,omitempty"`
	// Configuration for each expected response message.
	ResponseParameters []*ResponseParameters `protobuf:"bytes,2,rep,name=response_parameters" json:"response_parameters,omitempty"`
	// Optional input payload sent along with the request.
	Payload          *Payload `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *StreamingOutputCallRequest) Reset()                    { *m = StreamingOutputCallRequest{} }
func (m *StreamingOutputCallRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamingOutputCallRequest) ProtoMessage()               {}
func (*StreamingOutputCallRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *StreamingOutputCallRequest) GetResponseType() PayloadType {
	if m != nil && m.ResponseType != nil {
		return *m.ResponseType
	}
	return PayloadType_COMPRESSABLE
}

func (m *StreamingOutputCallRequest) GetResponseParameters() []*ResponseParameters {
	if m != nil {
		return m.ResponseParameters
	}
	return nil
}

func (m *StreamingOutputCallRequest) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Server-streaming response, as configured by the request and parameters.
type StreamingOutputCallResponse struct {
	// Payload to increase response size.
	Payload          *Payload `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *StreamingOutputCallResponse) Reset()                    { *m = StreamingOutputCallResponse{} }
func (m *StreamingOutputCallResponse) String() string            { return proto.CompactTextString(m) }
func (*StreamingOutputCallResponse) ProtoMessage()               {}
func (*StreamingOutputCallResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *StreamingOutputCallResponse) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterEnum("grpc.testing.PayloadType", PayloadType_name, PayloadType_value)
}

// Client API for TestService service

type TestServiceClient interface {
	// One empty request followed by one empty response.
	EmptyCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// One request followed by one response.
	// The server returns the client payload as-is.
	UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	// One request followed by a sequence of responses (streamed download).
	// The server returns the payload with client desired type and sizes.
	StreamingOutputCall(ctx context.Context, in *StreamingOutputCallRequest, opts ...grpc.CallOption) (TestService_StreamingOutputCallClient, error)
	// A sequence of requests followed by one response (streamed upload).
	// The server returns the aggregated size of client payload as the result.
	StreamingInputCall(ctx context.Context, opts ...grpc.CallOption) (TestService_StreamingInputCallClient, error)
	// A sequence of requests with each request served by the server immediately.
	// As one request could lead to multiple responses, this interface
	// demonstrates the idea of full duplexing.
	FullDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_FullDuplexCallClient, error)
	// A sequence of requests followed by a sequence of responses.
	// The server buffers all the client requests and then serves them in order. A
	// stream of responses are returned to the client when the server starts with
	// first request.
	HalfDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_HalfDuplexCallClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) EmptyCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/grpc.testing.TestService/EmptyCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := grpc.Invoke(ctx, "/grpc.testing.TestService/UnaryCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) StreamingOutputCall(ctx context.Context, in *StreamingOutputCallRequest, opts ...grpc.CallOption) (TestService_StreamingOutputCallClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[0], c.cc, "/grpc.testing.TestService/StreamingOutputCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceStreamingOutputCallClient{stream}
	if err := x.ClientStream.SendProto(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_StreamingOutputCallClient interface {
	Recv() (*StreamingOutputCallResponse, error)
	grpc.ClientStream
}

type testServiceStreamingOutputCallClient struct {
	grpc.ClientStream
}

func (x *testServiceStreamingOutputCallClient) Recv() (*StreamingOutputCallResponse, error) {
	m := new(StreamingOutputCallResponse)
	if err := x.ClientStream.RecvProto(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) StreamingInputCall(ctx context.Context, opts ...grpc.CallOption) (TestService_StreamingInputCallClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[1], c.cc, "/grpc.testing.TestService/StreamingInputCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceStreamingInputCallClient{stream}
	return x, nil
}

type TestService_StreamingInputCallClient interface {
	Send(*StreamingInputCallRequest) error
	CloseAndRecv() (*StreamingInputCallResponse, error)
	grpc.ClientStream
}

type testServiceStreamingInputCallClient struct {
	grpc.ClientStream
}

func (x *testServiceStreamingInputCallClient) Send(m *StreamingInputCallRequest) error {
	return x.ClientStream.SendProto(m)
}

func (x *testServiceStreamingInputCallClient) CloseAndRecv() (*StreamingInputCallResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamingInputCallResponse)
	if err := x.ClientStream.RecvProto(m); err != io.EOF {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) FullDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_FullDuplexCallClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[2], c.cc, "/grpc.testing.TestService/FullDuplexCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceFullDuplexCallClient{stream}
	return x, nil
}

type TestService_FullDuplexCallClient interface {
	Send(*StreamingOutputCallRequest) error
	Recv() (*StreamingOutputCallResponse, error)
	grpc.ClientStream
}

type testServiceFullDuplexCallClient struct {
	grpc.ClientStream
}

func (x *testServiceFullDuplexCallClient) Send(m *StreamingOutputCallRequest) error {
	return x.ClientStream.SendProto(m)
}

func (x *testServiceFullDuplexCallClient) Recv() (*StreamingOutputCallResponse, error) {
	m := new(StreamingOutputCallResponse)
	if err := x.ClientStream.RecvProto(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) HalfDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_HalfDuplexCallClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[3], c.cc, "/grpc.testing.TestService/HalfDuplexCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceHalfDuplexCallClient{stream}
	return x, nil
}

type TestService_HalfDuplexCallClient interface {
	Send(*StreamingOutputCallRequest) error
	Recv() (*StreamingOutputCallResponse, error)
	grpc.ClientStream
}

type testServiceHalfDuplexCallClient struct {
	grpc.ClientStream
}

func (x *testServiceHalfDuplexCallClient) Send(m *StreamingOutputCallRequest) error {
	return x.ClientStream.SendProto(m)
}

func (x *testServiceHalfDuplexCallClient) Recv() (*StreamingOutputCallResponse, error) {
	m := new(StreamingOutputCallResponse)
	if err := x.ClientStream.RecvProto(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TestService service

type TestServiceServer interface {
	// One empty request followed by one empty response.
	EmptyCall(context.Context, *Empty) (*Empty, error)
	// One request followed by one response.
	// The server returns the client payload as-is.
	UnaryCall(context.Context, *SimpleRequest) (*SimpleResponse, error)
	// One request followed by a sequence of responses (streamed download).
	// The server returns the payload with client desired type and sizes.
	StreamingOutputCall(*StreamingOutputCallRequest, TestService_StreamingOutputCallServer) error
	// A sequence of requests followed by one response (streamed upload).
	// The server returns the aggregated size of client payload as the result.
	StreamingInputCall(TestService_StreamingInputCallServer) error
	// A sequence of requests with each request served by the server immediately.
	// As one request could lead to multiple responses, this interface
	// demonstrates the idea of full duplexing.
	FullDuplexCall(TestService_FullDuplexCallServer) error
	// A sequence of requests followed by a sequence of responses.
	// The server buffers all the client requests and then serves them in order. A
	// stream of responses are returned to the client when the server starts with
	// first request.
	HalfDuplexCall(TestService_HalfDuplexCallServer) error
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_EmptyCall_Handler(srv interface{}, ctx context.Context, buf []byte) (proto.Message, error) {
	in := new(Empty)
	if err := proto.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(TestServiceServer).EmptyCall(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _TestService_UnaryCall_Handler(srv interface{}, ctx context.Context, buf []byte) (proto.Message, error) {
	in := new(SimpleRequest)
	if err := proto.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(TestServiceServer).UnaryCall(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _TestService_StreamingOutputCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamingOutputCallRequest)
	if err := stream.RecvProto(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).StreamingOutputCall(m, &testServiceStreamingOutputCallServer{stream})
}

type TestService_StreamingOutputCallServer interface {
	Send(*StreamingOutputCallResponse) error
	grpc.ServerStream
}

type testServiceStreamingOutputCallServer struct {
	grpc.ServerStream
}

func (x *testServiceStreamingOutputCallServer) Send(m *StreamingOutputCallResponse) error {
	return x.ServerStream.SendProto(m)
}

func _TestService_StreamingInputCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).StreamingInputCall(&testServiceStreamingInputCallServer{stream})
}

type TestService_StreamingInputCallServer interface {
	SendAndClose(*StreamingInputCallResponse) error
	Recv() (*StreamingInputCallRequest, error)
	grpc.ServerStream
}

type testServiceStreamingInputCallServer struct {
	grpc.ServerStream
}

func (x *testServiceStreamingInputCallServer) SendAndClose(m *StreamingInputCallResponse) error {
	return x.ServerStream.SendProto(m)
}

func (x *testServiceStreamingInputCallServer) Recv() (*StreamingInputCallRequest, error) {
	m := new(StreamingInputCallRequest)
	if err := x.ServerStream.RecvProto(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_FullDuplexCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).FullDuplexCall(&testServiceFullDuplexCallServer{stream})
}

type TestService_FullDuplexCallServer interface {
	Send(*StreamingOutputCallResponse) error
	Recv() (*StreamingOutputCallRequest, error)
	grpc.ServerStream
}

type testServiceFullDuplexCallServer struct {
	grpc.ServerStream
}

func (x *testServiceFullDuplexCallServer) Send(m *StreamingOutputCallResponse) error {
	return x.ServerStream.SendProto(m)
}

func (x *testServiceFullDuplexCallServer) Recv() (*StreamingOutputCallRequest, error) {
	m := new(StreamingOutputCallRequest)
	if err := x.ServerStream.RecvProto(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_HalfDuplexCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).HalfDuplexCall(&testServiceHalfDuplexCallServer{stream})
}

type TestService_HalfDuplexCallServer interface {
	Send(*StreamingOutputCallResponse) error
	Recv() (*StreamingOutputCallRequest, error)
	grpc.ServerStream
}

type testServiceHalfDuplexCallServer struct {
	grpc.ServerStream
}

func (x *testServiceHalfDuplexCallServer) Send(m *StreamingOutputCallResponse) error {
	return x.ServerStream.SendProto(m)
}

func (x *testServiceHalfDuplexCallServer) Recv() (*StreamingOutputCallRequest, error) {
	m := new(StreamingOutputCallRequest)
	if err := x.ServerStream.RecvProto(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmptyCall",
			Handler:    _TestService_EmptyCall_Handler,
		},
		{
			MethodName: "UnaryCall",
			Handler:    _TestService_UnaryCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingOutputCall",
			Handler:       _TestService_StreamingOutputCall_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamingInputCall",
			Handler:       _TestService_StreamingInputCall_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FullDuplexCall",
			Handler:       _TestService_FullDuplexCall_Handler,
			ClientStreams: true,
			ServerStreams: true,
		},
		{
			StreamName:    "HalfDuplexCall",
			Handler:       _TestService_HalfDuplexCall_Handler,
			ClientStreams: true,
			ServerStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x53, 0x5f, 0x6f, 0xd2, 0x50,
	0x14, 0xf7, 0x0e, 0x18, 0xd9, 0x81, 0x11, 0x72, 0xc8, 0x94, 0x75, 0x46, 0x97, 0x3e, 0x38, 0xf4,
	0x01, 0x08, 0xd1, 0xf8, 0xb4, 0xe8, 0x64, 0x2c, 0x9a, 0xb8, 0x41, 0xe8, 0x96, 0xf8, 0xd6, 0x5c,
	0xe1, 0x52, 0x9b, 0x94, 0xf6, 0xee, 0xf6, 0x76, 0xb1, 0x3e, 0xf9, 0x51, 0x7c, 0xdc, 0x17, 0xf0,
	0xa3, 0xf8, 0x7d, 0xbc, 0x6d, 0x61, 0xb6, 0xd8, 0x29, 0x7b, 0xd0, 0x27, 0x9a, 0x73, 0x7e, 0xff,
	0xce, 0x39, 0x5c, 0x78, 0x2e, 0x3f, 0xd9, 0x62, 0x6a, 0x72, 0x2a, 0x64, 0xd8, 0xb1, 0x3c, 0x87,
	0xba, 0x56, 0xc7, 0x12, 0x7c, 0xd2, 0x91, 0xcc, 0x97, 0xf1, 0x97, 0x19, 0x7d, 0xd9, 0xaa, 0x1c,
	0xfd, 0xb6, 0xb9, 0xf0, 0xa4, 0x87, 0xd5, 0xa8, 0xd1, 0x5e, 0x34, 0xf4, 0x32, 0x94, 0x06, 0x73,
	0x2e, 0x43, 0xfd, 0x35, 0x94, 0x47, 0x34, 0x74, 0x3c, 0x3a, 0xc5, 0x03, 0x28, 0xca, 0x90, 0xb3,
	0x26, 0xd9, 0x27, 0xad, 0x5a, 0x6f, 0xb7, 0x9d, 0x26, 0xb4, 0x17, 0xa0, 0x73, 0x05, 0xc0, 0x2a,
	0x14, 0x3f, 0x7a, 0xd3, 0xb0, 0xb9, 0xa1, 0x80, 0x55, 0xfd, 0x2b, 0x81, 0x6d, 0xc3, 0x9e, 0x73,
	0x87, 0x8d, 0xd9, 0x65, 0xa0, 0xe0, 0xd8, 0x85, 0x6d, 0xc1, 0x7c, 0xee, 0xb9, 0x3e, 0x33, 0xd7,
	0x53, 0xdc, 0x49, 0x31, 0x7c, 0xfb, 0x0b, 0x8b, 0xa5, 0x4b, 0xf8, 0x04, 0xca, 0x3c, 0x41, 0x35,
	0x0b, 0xaa, 0x50, 0xe9, 0xed, 0xe4, 0x4a, 0xe8, 0x1f, 0xa0, 0xb6, 0x4c, 0x90, 0x88, 0xa4, 0x99,
	0xe4, 0x0f, 0x4c, 0x7c, 0x04, 0xf7, 0xd9, 0x6c, 0xc6, 0x26, 0xd2, 0xbe, 0x62, 0xa6, 0x45, 0x6d,
	0x6a, 0x06, 0x3e, 0x13, 0xa6, 0x3d, 0x8d, 0x13, 0x14, 0xf4, 0x3e, 0xec, 0x1a, 0x52, 0x30, 0x3a,
	0x57, 0xa4, 0x77, 0x2e, 0x0f, 0x64, 0x9f, 0x3a, 0xce, 0x72, 0xce, 0x35, 0x4d, 0xf4, 0x43, 0xd0,
	0xf2, 0x44, 0x16, 0x51, 0x1f, 0xc3, 0x03, 0x6a, 0x59, 0x82, 0x59, 0x54, 0xb2, 0xe8, 0xaa, 0x31,
	0x27, 0xd9, 0x42, 0xa4, 0x5a, 0xd2, 0x5f, 0x02, 0x2e, 0xc1, 0x23, 0x2a, 0xe8, 0x9c, 0x49, 0x26,
	0xfc, 0xe8, 0x08, 0xbf, 0x30, 0xd8, 0x80, 0x8a, 0xed, 0xaa, 0xfa, 0x15, 0x75, 0xd4, 0x04, 0xc9,
	0xfa, 0xf4, 0xef, 0x24, 0x65, 0x3c, 0x0c, 0xe4, 0x4a, 0xfc, 0xbb, 0x9f, 0xe9, 0x10, 0x1a, 0x37,
	0x0c, 0x7e, 0x13, 0x45, 0xb9, 0x15, 0xd4, 0xf0, 0xfb, 0x59, 0x5e, 0x4e, 0xe4, 0x75, 0xcf, 0x39,
	0x80, 0xbd, 0xdc, 0xd8, 0x77, 0xbb, 0xed, 0xb3, 0x57, 0x50, 0x49, 0x87, 0xaf, 0x43, 0xb5, 0x3f,
	0x3c, 0x1d, 0x8d, 0x07, 0x86, 0x71, 0xf4, 0xe6, 0xfd, 0xa0, 0x7e, 0x0f, 0x11, 0x6a, 0x17, 0x67,
	0x99, 0x1a, 0x41, 0x80, 0xcd, 0xf1, 0xd1, 0xd9, 0xf1, 0xf0, 0xb4, 0xbe, 0xd1, 0xfb, 0x51, 0x84,
	0xca, 0xb9, 0x12, 0x35, 0xd4, 0x5e, 0xed, 0x09, 0xc3, 0x17, 0xb0, 0x15, 0x3f, 0x9a, 0x28, 0x0d,
	0x36, 0xb2, 0xa6, 0x71, 0x43, 0xcb, 0x2b, 0xe2, 0x09, 0x6c, 0x5d, 0xb8, 0x54, 0x24, 0xb4, 0xbd,
	0x2c, 0x22, 0xf3, 0x70, 0xb4, 0x87, 0xf9, 0xcd, 0xc5, 0xdc, 0x97, 0xd0, 0xc8, 0x59, 0x0b, 0xb6,
	0x56, 0x48, 0xb7, 0x1e, 0x5c, 0x7b, 0xba, 0x06, 0x32, 0xf1, 0xd2, 0x0b, 0xd7, 0x84, 0x74, 0x09,
	0xba, 0x80, 0xbf, 0xff, 0x73, 0xf1, 0xe0, 0x16, 0x9d, 0xd5, 0x07, 0xa2, 0xb5, 0xfe, 0x0e, 0x5c,
	0xfa, 0x7d, 0x23, 0xa4, 0x45, 0xd4, 0x88, 0xb5, 0x93, 0xc0, 0x71, 0x8e, 0x03, 0x35, 0xf7, 0xe7,
	0x7f, 0x37, 0xdd, 0xa6, 0x72, 0xbb, 0x8e, 0x0c, 0xbb, 0xb1, 0xe5, 0x5b, 0xea, 0xcc, 0xfe, 0xa3,
	0xe5, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x98, 0x02, 0x21, 0x8d, 0xc1, 0x05, 0x00, 0x00,
}
