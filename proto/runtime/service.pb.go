// Code generated by protoc-gen-go. DO NOT EDIT.
// source: runtime/service.proto

package runtime // import "github.com/socketfunc/proto/runtime"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import packet "github.com/socketfunc/proto/packet"
import store "github.com/socketfunc/proto/store"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StreamCmd int32

const (
	StreamCmd_MESSAGE StreamCmd = 0
	StreamCmd_STORE   StreamCmd = 1
)

var StreamCmd_name = map[int32]string{
	0: "MESSAGE",
	1: "STORE",
}
var StreamCmd_value = map[string]int32{
	"MESSAGE": 0,
	"STORE":   1,
}

func (x StreamCmd) String() string {
	return proto.EnumName(StreamCmd_name, int32(x))
}
func (StreamCmd) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_service_5e76dd191cbd0f5d, []int{0}
}

type StoreCmd int32

const (
	StoreCmd_GET    StoreCmd = 0
	StoreCmd_PUT    StoreCmd = 1
	StoreCmd_MODIFY StoreCmd = 2
	StoreCmd_DEL    StoreCmd = 3
)

var StoreCmd_name = map[int32]string{
	0: "GET",
	1: "PUT",
	2: "MODIFY",
	3: "DEL",
}
var StoreCmd_value = map[string]int32{
	"GET":    0,
	"PUT":    1,
	"MODIFY": 2,
	"DEL":    3,
}

func (x StoreCmd) String() string {
	return proto.EnumName(StoreCmd_name, int32(x))
}
func (StoreCmd) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_service_5e76dd191cbd0f5d, []int{1}
}

type MessageResponse_Type int32

const (
	MessageResponse_SEND      MessageResponse_Type = 0
	MessageResponse_REPLY     MessageResponse_Type = 1
	MessageResponse_BROADCAST MessageResponse_Type = 2
)

var MessageResponse_Type_name = map[int32]string{
	0: "SEND",
	1: "REPLY",
	2: "BROADCAST",
}
var MessageResponse_Type_value = map[string]int32{
	"SEND":      0,
	"REPLY":     1,
	"BROADCAST": 2,
}

func (x MessageResponse_Type) String() string {
	return proto.EnumName(MessageResponse_Type_name, int32(x))
}
func (MessageResponse_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_service_5e76dd191cbd0f5d, []int{3, 0}
}

type StreamRequest struct {
	Cmd                  StreamCmd       `protobuf:"varint,1,opt,name=cmd,proto3,enum=runtime.StreamCmd" json:"cmd,omitempty"`
	MessageRequest       *MessageRequest `protobuf:"bytes,2,opt,name=message_request,json=messageRequest,proto3" json:"message_request,omitempty"`
	StoreResponse        *StoreResponse  `protobuf:"bytes,3,opt,name=store_response,json=storeResponse,proto3" json:"store_response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5e76dd191cbd0f5d, []int{0}
}
func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (dst *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(dst, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetCmd() StreamCmd {
	if m != nil {
		return m.Cmd
	}
	return StreamCmd_MESSAGE
}

func (m *StreamRequest) GetMessageRequest() *MessageRequest {
	if m != nil {
		return m.MessageRequest
	}
	return nil
}

func (m *StreamRequest) GetStoreResponse() *StoreResponse {
	if m != nil {
		return m.StoreResponse
	}
	return nil
}

type StreamSend struct {
	Cmd                  StreamCmd        `protobuf:"varint,1,opt,name=cmd,proto3,enum=runtime.StreamCmd" json:"cmd,omitempty"`
	MessageResponse      *MessageResponse `protobuf:"bytes,2,opt,name=message_response,json=messageResponse,proto3" json:"message_response,omitempty"`
	StoreRequest         *StoreRequest    `protobuf:"bytes,3,opt,name=store_request,json=storeRequest,proto3" json:"store_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *StreamSend) Reset()         { *m = StreamSend{} }
func (m *StreamSend) String() string { return proto.CompactTextString(m) }
func (*StreamSend) ProtoMessage()    {}
func (*StreamSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5e76dd191cbd0f5d, []int{1}
}
func (m *StreamSend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamSend.Unmarshal(m, b)
}
func (m *StreamSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamSend.Marshal(b, m, deterministic)
}
func (dst *StreamSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamSend.Merge(dst, src)
}
func (m *StreamSend) XXX_Size() int {
	return xxx_messageInfo_StreamSend.Size(m)
}
func (m *StreamSend) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamSend.DiscardUnknown(m)
}

var xxx_messageInfo_StreamSend proto.InternalMessageInfo

func (m *StreamSend) GetCmd() StreamCmd {
	if m != nil {
		return m.Cmd
	}
	return StreamCmd_MESSAGE
}

func (m *StreamSend) GetMessageResponse() *MessageResponse {
	if m != nil {
		return m.MessageResponse
	}
	return nil
}

func (m *StreamSend) GetStoreRequest() *StoreRequest {
	if m != nil {
		return m.StoreRequest
	}
	return nil
}

type MessageRequest struct {
	Packet               *packet.Packet `protobuf:"bytes,1,opt,name=packet,proto3" json:"packet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MessageRequest) Reset()         { *m = MessageRequest{} }
func (m *MessageRequest) String() string { return proto.CompactTextString(m) }
func (*MessageRequest) ProtoMessage()    {}
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5e76dd191cbd0f5d, []int{2}
}
func (m *MessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageRequest.Unmarshal(m, b)
}
func (m *MessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageRequest.Marshal(b, m, deterministic)
}
func (dst *MessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageRequest.Merge(dst, src)
}
func (m *MessageRequest) XXX_Size() int {
	return xxx_messageInfo_MessageRequest.Size(m)
}
func (m *MessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MessageRequest proto.InternalMessageInfo

func (m *MessageRequest) GetPacket() *packet.Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

type MessageResponse struct {
	Type                 MessageResponse_Type `protobuf:"varint,1,opt,name=type,proto3,enum=runtime.MessageResponse_Type" json:"type,omitempty"`
	Packet               *packet.Packet       `protobuf:"bytes,2,opt,name=packet,proto3" json:"packet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MessageResponse) Reset()         { *m = MessageResponse{} }
func (m *MessageResponse) String() string { return proto.CompactTextString(m) }
func (*MessageResponse) ProtoMessage()    {}
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5e76dd191cbd0f5d, []int{3}
}
func (m *MessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageResponse.Unmarshal(m, b)
}
func (m *MessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageResponse.Marshal(b, m, deterministic)
}
func (dst *MessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageResponse.Merge(dst, src)
}
func (m *MessageResponse) XXX_Size() int {
	return xxx_messageInfo_MessageResponse.Size(m)
}
func (m *MessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MessageResponse proto.InternalMessageInfo

func (m *MessageResponse) GetType() MessageResponse_Type {
	if m != nil {
		return m.Type
	}
	return MessageResponse_SEND
}

func (m *MessageResponse) GetPacket() *packet.Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

type StoreRequest struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Cmd                  StoreCmd        `protobuf:"varint,2,opt,name=cmd,proto3,enum=runtime.StoreCmd" json:"cmd,omitempty"`
	Key                  string          `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Entity               *store.Entity   `protobuf:"bytes,4,opt,name=entity,proto3" json:"entity,omitempty"`
	Filters              []*store.Filter `protobuf:"bytes,5,rep,name=filters,proto3" json:"filters,omitempty"`
	Updates              []*store.Update `protobuf:"bytes,6,rep,name=updates,proto3" json:"updates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StoreRequest) Reset()         { *m = StoreRequest{} }
func (m *StoreRequest) String() string { return proto.CompactTextString(m) }
func (*StoreRequest) ProtoMessage()    {}
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5e76dd191cbd0f5d, []int{4}
}
func (m *StoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreRequest.Unmarshal(m, b)
}
func (m *StoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreRequest.Marshal(b, m, deterministic)
}
func (dst *StoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreRequest.Merge(dst, src)
}
func (m *StoreRequest) XXX_Size() int {
	return xxx_messageInfo_StoreRequest.Size(m)
}
func (m *StoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoreRequest proto.InternalMessageInfo

func (m *StoreRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StoreRequest) GetCmd() StoreCmd {
	if m != nil {
		return m.Cmd
	}
	return StoreCmd_GET
}

func (m *StoreRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StoreRequest) GetEntity() *store.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *StoreRequest) GetFilters() []*store.Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *StoreRequest) GetUpdates() []*store.Update {
	if m != nil {
		return m.Updates
	}
	return nil
}

type StoreResponse struct {
	Id                   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Cmd                  StoreCmd      `protobuf:"varint,2,opt,name=cmd,proto3,enum=runtime.StoreCmd" json:"cmd,omitempty"`
	Success              bool          `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	Entity               *store.Entity `protobuf:"bytes,4,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *StoreResponse) Reset()         { *m = StoreResponse{} }
func (m *StoreResponse) String() string { return proto.CompactTextString(m) }
func (*StoreResponse) ProtoMessage()    {}
func (*StoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5e76dd191cbd0f5d, []int{5}
}
func (m *StoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreResponse.Unmarshal(m, b)
}
func (m *StoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreResponse.Marshal(b, m, deterministic)
}
func (dst *StoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreResponse.Merge(dst, src)
}
func (m *StoreResponse) XXX_Size() int {
	return xxx_messageInfo_StoreResponse.Size(m)
}
func (m *StoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StoreResponse proto.InternalMessageInfo

func (m *StoreResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StoreResponse) GetCmd() StoreCmd {
	if m != nil {
		return m.Cmd
	}
	return StoreCmd_GET
}

func (m *StoreResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *StoreResponse) GetEntity() *store.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type HealthzRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthzRequest) Reset()         { *m = HealthzRequest{} }
func (m *HealthzRequest) String() string { return proto.CompactTextString(m) }
func (*HealthzRequest) ProtoMessage()    {}
func (*HealthzRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5e76dd191cbd0f5d, []int{6}
}
func (m *HealthzRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthzRequest.Unmarshal(m, b)
}
func (m *HealthzRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthzRequest.Marshal(b, m, deterministic)
}
func (dst *HealthzRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthzRequest.Merge(dst, src)
}
func (m *HealthzRequest) XXX_Size() int {
	return xxx_messageInfo_HealthzRequest.Size(m)
}
func (m *HealthzRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthzRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthzRequest proto.InternalMessageInfo

type HealthzResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthzResponse) Reset()         { *m = HealthzResponse{} }
func (m *HealthzResponse) String() string { return proto.CompactTextString(m) }
func (*HealthzResponse) ProtoMessage()    {}
func (*HealthzResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_5e76dd191cbd0f5d, []int{7}
}
func (m *HealthzResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthzResponse.Unmarshal(m, b)
}
func (m *HealthzResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthzResponse.Marshal(b, m, deterministic)
}
func (dst *HealthzResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthzResponse.Merge(dst, src)
}
func (m *HealthzResponse) XXX_Size() int {
	return xxx_messageInfo_HealthzResponse.Size(m)
}
func (m *HealthzResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthzResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthzResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StreamRequest)(nil), "runtime.StreamRequest")
	proto.RegisterType((*StreamSend)(nil), "runtime.StreamSend")
	proto.RegisterType((*MessageRequest)(nil), "runtime.MessageRequest")
	proto.RegisterType((*MessageResponse)(nil), "runtime.MessageResponse")
	proto.RegisterType((*StoreRequest)(nil), "runtime.StoreRequest")
	proto.RegisterType((*StoreResponse)(nil), "runtime.StoreResponse")
	proto.RegisterType((*HealthzRequest)(nil), "runtime.HealthzRequest")
	proto.RegisterType((*HealthzResponse)(nil), "runtime.HealthzResponse")
	proto.RegisterEnum("runtime.StreamCmd", StreamCmd_name, StreamCmd_value)
	proto.RegisterEnum("runtime.StoreCmd", StoreCmd_name, StoreCmd_value)
	proto.RegisterEnum("runtime.MessageResponse_Type", MessageResponse_Type_name, MessageResponse_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RuntimeServiceClient is the client API for RuntimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RuntimeServiceClient interface {
	Stream(ctx context.Context, opts ...grpc.CallOption) (RuntimeService_StreamClient, error)
	Healthz(ctx context.Context, in *HealthzRequest, opts ...grpc.CallOption) (*HealthzResponse, error)
}

type runtimeServiceClient struct {
	cc *grpc.ClientConn
}

func NewRuntimeServiceClient(cc *grpc.ClientConn) RuntimeServiceClient {
	return &runtimeServiceClient{cc}
}

func (c *runtimeServiceClient) Stream(ctx context.Context, opts ...grpc.CallOption) (RuntimeService_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RuntimeService_serviceDesc.Streams[0], "/runtime.RuntimeService/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &runtimeServiceStreamClient{stream}
	return x, nil
}

type RuntimeService_StreamClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamSend, error)
	grpc.ClientStream
}

type runtimeServiceStreamClient struct {
	grpc.ClientStream
}

func (x *runtimeServiceStreamClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *runtimeServiceStreamClient) Recv() (*StreamSend, error) {
	m := new(StreamSend)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *runtimeServiceClient) Healthz(ctx context.Context, in *HealthzRequest, opts ...grpc.CallOption) (*HealthzResponse, error) {
	out := new(HealthzResponse)
	err := c.cc.Invoke(ctx, "/runtime.RuntimeService/Healthz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RuntimeServiceServer is the server API for RuntimeService service.
type RuntimeServiceServer interface {
	Stream(RuntimeService_StreamServer) error
	Healthz(context.Context, *HealthzRequest) (*HealthzResponse, error)
}

func RegisterRuntimeServiceServer(s *grpc.Server, srv RuntimeServiceServer) {
	s.RegisterService(&_RuntimeService_serviceDesc, srv)
}

func _RuntimeService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RuntimeServiceServer).Stream(&runtimeServiceStreamServer{stream})
}

type RuntimeService_StreamServer interface {
	Send(*StreamSend) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type runtimeServiceStreamServer struct {
	grpc.ServerStream
}

func (x *runtimeServiceStreamServer) Send(m *StreamSend) error {
	return x.ServerStream.SendMsg(m)
}

func (x *runtimeServiceStreamServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RuntimeService_Healthz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthzRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeServiceServer).Healthz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runtime.RuntimeService/Healthz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeServiceServer).Healthz(ctx, req.(*HealthzRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RuntimeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "runtime.RuntimeService",
	HandlerType: (*RuntimeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Healthz",
			Handler:    _RuntimeService_Healthz_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _RuntimeService_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "runtime/service.proto",
}

func init() { proto.RegisterFile("runtime/service.proto", fileDescriptor_service_5e76dd191cbd0f5d) }

var fileDescriptor_service_5e76dd191cbd0f5d = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x51, 0x4f, 0xdb, 0x3c,
	0x14, 0xad, 0xd3, 0xd2, 0xd2, 0x0b, 0x4d, 0x83, 0x11, 0x7c, 0x51, 0xa5, 0x4f, 0x42, 0x61, 0x6c,
	0x88, 0x87, 0x64, 0x74, 0x2f, 0xd3, 0xa6, 0x4d, 0x03, 0x1a, 0xd8, 0x24, 0x18, 0xc8, 0x29, 0x0f,
	0xec, 0x05, 0x85, 0xd4, 0x40, 0x04, 0x69, 0xb2, 0xd8, 0x41, 0xea, 0xde, 0xf7, 0xb8, 0xf7, 0xfd,
	0x8f, 0x49, 0xfb, 0x19, 0xfb, 0x4d, 0x53, 0x6c, 0xa7, 0x24, 0x45, 0x9b, 0xd0, 0x9e, 0x12, 0x9f,
	0x73, 0xed, 0x7b, 0x8e, 0xef, 0xbd, 0x86, 0x95, 0x34, 0x1b, 0xf3, 0x30, 0xa2, 0x0e, 0xa3, 0xe9,
	0x5d, 0x18, 0x50, 0x3b, 0x49, 0x63, 0x1e, 0xe3, 0x96, 0x82, 0x7b, 0xcb, 0x89, 0x1f, 0xdc, 0x50,
	0xee, 0xc8, 0x8f, 0x64, 0x7b, 0x98, 0xf1, 0x38, 0xa5, 0xce, 0x65, 0x78, 0xcb, 0x69, 0x5a, 0xc5,
	0xb2, 0x64, 0xe4, 0x73, 0x75, 0x4a, 0x6f, 0x49, 0x62, 0x77, 0xfe, 0x6d, 0xa6, 0x20, 0xeb, 0x27,
	0x82, 0x8e, 0xc7, 0x53, 0xea, 0x47, 0x84, 0x7e, 0xce, 0x28, 0xe3, 0xf8, 0x09, 0xd4, 0x83, 0x68,
	0x64, 0xa2, 0x35, 0xb4, 0xa9, 0xf7, 0xb1, 0xad, 0x12, 0xdb, 0x32, 0x68, 0x2f, 0x1a, 0x91, 0x9c,
	0xc6, 0xef, 0xa0, 0x1b, 0x51, 0xc6, 0xfc, 0x2b, 0x7a, 0x9e, 0xca, 0x8d, 0xa6, 0xb6, 0x86, 0x36,
	0x17, 0xfa, 0xff, 0x4d, 0x77, 0x1c, 0x49, 0x5e, 0x9d, 0x4b, 0xf4, 0xa8, 0xb2, 0xc6, 0x6f, 0x40,
	0x17, 0x72, 0xce, 0x53, 0xca, 0x92, 0x78, 0xcc, 0xa8, 0x59, 0x17, 0x07, 0xac, 0x96, 0x52, 0xc6,
	0x29, 0x25, 0x8a, 0x25, 0x1d, 0x56, 0x5e, 0x5a, 0x3f, 0x10, 0x80, 0xd4, 0xe4, 0xd1, 0xf1, 0xe8,
	0x91, 0xaa, 0xf7, 0xc0, 0xb8, 0x57, 0xad, 0xb2, 0x4a, 0xd9, 0xe6, 0x43, 0xd9, 0x2a, 0x6f, 0x37,
	0xaa, 0x02, 0xf8, 0x15, 0x74, 0x0a, 0xe1, 0xd2, 0xb8, 0xd4, 0xbd, 0x32, 0xab, 0x5b, 0xda, 0x5e,
	0x64, 0xa5, 0x95, 0xf5, 0x12, 0xf4, 0xea, 0xb5, 0xe0, 0xa7, 0xd0, 0x94, 0xb5, 0x14, 0xda, 0x17,
	0xfa, 0xba, 0xad, 0x4a, 0x7b, 0x22, 0x3e, 0x44, 0xb1, 0xd6, 0x77, 0x04, 0xdd, 0x19, 0x69, 0x78,
	0x1b, 0x1a, 0x7c, 0x92, 0x50, 0xe5, 0xfa, 0xff, 0x3f, 0x59, 0xb0, 0x87, 0x93, 0x84, 0x12, 0x11,
	0x5a, 0x4a, 0xa7, 0xfd, 0x35, 0xdd, 0x16, 0x34, 0xf2, 0x5d, 0x78, 0x1e, 0x1a, 0x9e, 0xfb, 0x71,
	0x60, 0xd4, 0x70, 0x1b, 0xe6, 0x88, 0x7b, 0x72, 0x78, 0x66, 0x20, 0xdc, 0x81, 0xf6, 0x2e, 0x39,
	0xde, 0x19, 0xec, 0xed, 0x78, 0x43, 0x43, 0xb3, 0x7e, 0x21, 0x58, 0x2c, 0x7b, 0xc6, 0x3a, 0x68,
	0xa1, 0xac, 0x45, 0x9b, 0x68, 0xe1, 0x08, 0xaf, 0xcb, 0xe2, 0x68, 0x42, 0xe6, 0x52, 0xf5, 0x9e,
	0xa6, 0xb5, 0x31, 0xa0, 0x7e, 0x43, 0x27, 0xe2, 0x32, 0xdb, 0x24, 0xff, 0xc5, 0x1b, 0xd0, 0xa4,
	0x63, 0x1e, 0xf2, 0x89, 0xd9, 0x10, 0x5a, 0x3b, 0xb6, 0xb8, 0x4b, 0xdb, 0x15, 0x20, 0x51, 0x24,
	0x7e, 0x06, 0x2d, 0xd9, 0xf9, 0xcc, 0x9c, 0x5b, 0xab, 0x97, 0xe2, 0xf6, 0x05, 0x4a, 0x0a, 0x36,
	0x0f, 0x94, 0xe3, 0xc0, 0xcc, 0x66, 0x25, 0xf0, 0x54, 0xa0, 0xa4, 0x60, 0xad, 0xaf, 0x62, 0x28,
	0x4a, 0xdd, 0xf6, 0x6f, 0x8e, 0x4c, 0x68, 0xb1, 0x2c, 0x08, 0x28, 0x63, 0xc2, 0xd5, 0x3c, 0x29,
	0x96, 0x8f, 0x74, 0x66, 0x19, 0xa0, 0xbf, 0xa7, 0xfe, 0x2d, 0xbf, 0xfe, 0x52, 0xf4, 0xcf, 0x12,
	0x74, 0xa7, 0x88, 0x94, 0xb6, 0xb5, 0x0e, 0xed, 0x69, 0x97, 0xe3, 0x05, 0x68, 0x1d, 0xb9, 0x9e,
	0xb7, 0x73, 0xe0, 0xca, 0x8a, 0x79, 0xc3, 0x63, 0xe2, 0x1a, 0x68, 0x6b, 0x1b, 0xe6, 0x0b, 0x6d,
	0xb8, 0x05, 0xf5, 0x03, 0x77, 0x68, 0xd4, 0xf2, 0x9f, 0x93, 0xd3, 0xa1, 0x81, 0x30, 0x40, 0xf3,
	0xe8, 0x78, 0xf0, 0x61, 0xff, 0xcc, 0xd0, 0x72, 0x70, 0xe0, 0x1e, 0x1a, 0xf5, 0xfe, 0x37, 0x04,
	0x3a, 0x91, 0xbe, 0x3c, 0xf9, 0x16, 0xe1, 0xd7, 0xd0, 0x94, 0xa9, 0xf0, 0xea, 0xcc, 0x84, 0x29,
	0x7d, 0xbd, 0xe5, 0x19, 0x3c, 0x9f, 0x4d, 0xab, 0xb6, 0x89, 0x9e, 0x23, 0xfc, 0x16, 0x5a, 0x4a,
	0x3a, 0xbe, 0x7f, 0x23, 0xaa, 0xf6, 0x7a, 0xe6, 0x43, 0x42, 0x8d, 0x7b, 0x6d, 0x77, 0xe3, 0xd3,
	0xfa, 0x55, 0xc8, 0xaf, 0xb3, 0x0b, 0x3b, 0x88, 0x23, 0x87, 0xc5, 0x79, 0x9b, 0x5e, 0x66, 0xe3,
	0xc0, 0x11, 0x0f, 0x99, 0xa3, 0x36, 0x5e, 0x34, 0xc5, 0xf2, 0xc5, 0xef, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xaa, 0x80, 0xad, 0x06, 0x49, 0x05, 0x00, 0x00,
}