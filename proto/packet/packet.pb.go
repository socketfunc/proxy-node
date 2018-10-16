// Code generated by protoc-gen-go. DO NOT EDIT.
// source: packet/packet.proto

package packet // import "github.com/socketfunc/proto/packet"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Packet struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Topic                string   `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Event                string   `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
	Payload              []byte   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_packet_3ba6270a2537819b, []int{0}
}
func (m *Packet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet.Unmarshal(m, b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
}
func (dst *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(dst, src)
}
func (m *Packet) XXX_Size() int {
	return xxx_messageInfo_Packet.Size(m)
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

func (m *Packet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Packet) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Packet) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *Packet) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Packet)(nil), "packet.Packet")
}

func init() { proto.RegisterFile("packet/packet.proto", fileDescriptor_packet_3ba6270a2537819b) }

var fileDescriptor_packet_3ba6270a2537819b = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x48, 0x4c, 0xce,
	0x4e, 0x2d, 0xd1, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x52,
	0x1c, 0x17, 0x5b, 0x00, 0x98, 0x25, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x19, 0xc4, 0x94, 0x99, 0x22, 0x24, 0xc2, 0xc5, 0x5a, 0x92, 0x5f, 0x90, 0x99, 0x2c, 0xc1,
	0x04, 0x16, 0x82, 0x70, 0x40, 0xa2, 0xa9, 0x65, 0xa9, 0x79, 0x25, 0x12, 0xcc, 0x10, 0x51, 0x30,
	0x47, 0x48, 0x82, 0x8b, 0xbd, 0x20, 0xb1, 0x32, 0x27, 0x3f, 0x31, 0x45, 0x82, 0x45, 0x81, 0x51,
	0x83, 0x27, 0x08, 0xc6, 0x75, 0x52, 0x89, 0x52, 0x4a, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b,
	0xce, 0xcf, 0xd5, 0x2f, 0xce, 0x07, 0x59, 0x95, 0x56, 0x9a, 0x97, 0xac, 0x0f, 0x76, 0x06, 0xd4,
	0x4d, 0x49, 0x6c, 0x60, 0x9e, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x04, 0x0f, 0x20, 0xab,
	0x00, 0x00, 0x00,
}
