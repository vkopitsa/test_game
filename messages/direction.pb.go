// Code generated by protoc-gen-go. DO NOT EDIT.
// source: direction.proto

package messages

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Direction_Type int32

const (
	Direction_STOP  Direction_Type = 0
	Direction_UP    Direction_Type = 1
	Direction_DOWN  Direction_Type = 2
	Direction_LEFT  Direction_Type = 3
	Direction_RIGHT Direction_Type = 4
	Direction_NONE  Direction_Type = 5
)

var Direction_Type_name = map[int32]string{
	0: "STOP",
	1: "UP",
	2: "DOWN",
	3: "LEFT",
	4: "RIGHT",
	5: "NONE",
}

var Direction_Type_value = map[string]int32{
	"STOP":  0,
	"UP":    1,
	"DOWN":  2,
	"LEFT":  3,
	"RIGHT": 4,
	"NONE":  5,
}

func (x Direction_Type) String() string {
	return proto.EnumName(Direction_Type_name, int32(x))
}

func (Direction_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b0916f7e0e5c3a69, []int{0, 0}
}

type Direction struct {
	Type                 Direction_Type `protobuf:"varint,1,opt,name=type,proto3,enum=messages.Direction_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Direction) Reset()         { *m = Direction{} }
func (m *Direction) String() string { return proto.CompactTextString(m) }
func (*Direction) ProtoMessage()    {}
func (*Direction) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0916f7e0e5c3a69, []int{0}
}

func (m *Direction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Direction.Unmarshal(m, b)
}
func (m *Direction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Direction.Marshal(b, m, deterministic)
}
func (m *Direction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Direction.Merge(m, src)
}
func (m *Direction) XXX_Size() int {
	return xxx_messageInfo_Direction.Size(m)
}
func (m *Direction) XXX_DiscardUnknown() {
	xxx_messageInfo_Direction.DiscardUnknown(m)
}

var xxx_messageInfo_Direction proto.InternalMessageInfo

func (m *Direction) GetType() Direction_Type {
	if m != nil {
		return m.Type
	}
	return Direction_STOP
}

func init() {
	proto.RegisterEnum("messages.Direction_Type", Direction_Type_name, Direction_Type_value)
	proto.RegisterType((*Direction)(nil), "messages.Direction")
}

func init() { proto.RegisterFile("direction.proto", fileDescriptor_b0916f7e0e5c3a69) }

var fileDescriptor_b0916f7e0e5c3a69 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xc9, 0x2c, 0x4a,
	0x4d, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0x4d, 0x2d,
	0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0x56, 0xaa, 0xe1, 0xe2, 0x74, 0x81, 0x49, 0x0a, 0xe9, 0x70, 0xb1,
	0x94, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x19, 0x49, 0xe8, 0xc1, 0x54, 0xe9,
	0xc1, 0x95, 0xe8, 0x85, 0x54, 0x16, 0xa4, 0x06, 0x81, 0x55, 0x29, 0x39, 0x72, 0xb1, 0x80, 0x78,
	0x42, 0x1c, 0x5c, 0x2c, 0xc1, 0x21, 0xfe, 0x01, 0x02, 0x0c, 0x42, 0x6c, 0x5c, 0x4c, 0xa1, 0x01,
	0x02, 0x8c, 0x20, 0x11, 0x17, 0xff, 0x70, 0x3f, 0x01, 0x26, 0x10, 0xcb, 0xc7, 0xd5, 0x2d, 0x44,
	0x80, 0x59, 0x88, 0x93, 0x8b, 0x35, 0xc8, 0xd3, 0xdd, 0x23, 0x44, 0x80, 0x05, 0x24, 0xe8, 0xe7,
	0xef, 0xe7, 0x2a, 0xc0, 0x9a, 0xc4, 0x06, 0x76, 0x8e, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x55,
	0x01, 0x80, 0x15, 0xa1, 0x00, 0x00, 0x00,
}
