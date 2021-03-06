// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/retry/previous_priorities/previous_priorities_config.proto

package envoy_config_retry_previous_priorities

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

type PreviousPrioritiesConfig struct {
	UpdateFrequency      int32    `protobuf:"varint,1,opt,name=update_frequency,json=updateFrequency,proto3" json:"update_frequency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PreviousPrioritiesConfig) Reset()         { *m = PreviousPrioritiesConfig{} }
func (m *PreviousPrioritiesConfig) String() string { return proto.CompactTextString(m) }
func (*PreviousPrioritiesConfig) ProtoMessage()    {}
func (*PreviousPrioritiesConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_763b66cb0fac7d2f, []int{0}
}

func (m *PreviousPrioritiesConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PreviousPrioritiesConfig.Unmarshal(m, b)
}
func (m *PreviousPrioritiesConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PreviousPrioritiesConfig.Marshal(b, m, deterministic)
}
func (m *PreviousPrioritiesConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreviousPrioritiesConfig.Merge(m, src)
}
func (m *PreviousPrioritiesConfig) XXX_Size() int {
	return xxx_messageInfo_PreviousPrioritiesConfig.Size(m)
}
func (m *PreviousPrioritiesConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PreviousPrioritiesConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PreviousPrioritiesConfig proto.InternalMessageInfo

func (m *PreviousPrioritiesConfig) GetUpdateFrequency() int32 {
	if m != nil {
		return m.UpdateFrequency
	}
	return 0
}

func init() {
	proto.RegisterType((*PreviousPrioritiesConfig)(nil), "envoy.config.retry.previous_priorities.PreviousPrioritiesConfig")
}

func init() {
	proto.RegisterFile("envoy/config/retry/previous_priorities/previous_priorities_config.proto", fileDescriptor_763b66cb0fac7d2f)
}

var fileDescriptor_763b66cb0fac7d2f = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4f, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2f, 0x4a, 0x2d, 0x29, 0xaa, 0xd4, 0x2f,
	0x28, 0x4a, 0x2d, 0xcb, 0xcc, 0x2f, 0x2d, 0x8e, 0x2f, 0x28, 0xca, 0xcc, 0x2f, 0xca, 0x2c, 0xc9,
	0x4c, 0x2d, 0xc6, 0x26, 0x16, 0x0f, 0xd1, 0xa4, 0x57, 0x50, 0x94, 0x5f, 0x92, 0x2f, 0xa4, 0x06,
	0x36, 0x48, 0x0f, 0x2a, 0x06, 0x36, 0x48, 0x0f, 0x8b, 0x26, 0x25, 0x57, 0x2e, 0x89, 0x00, 0xa8,
	0x70, 0x00, 0x5c, 0xd4, 0x19, 0xac, 0x4b, 0x48, 0x93, 0x4b, 0xa0, 0xb4, 0x20, 0x25, 0xb1, 0x24,
	0x35, 0x3e, 0xad, 0x28, 0xb5, 0xb0, 0x34, 0x35, 0x2f, 0xb9, 0x52, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x35, 0x88, 0x1f, 0x22, 0xee, 0x06, 0x13, 0x76, 0x0a, 0xe7, 0x32, 0xc9, 0xcc, 0xd7, 0x03, 0xdb,
	0x59, 0x50, 0x94, 0x5f, 0x51, 0xa9, 0x47, 0x9c, 0xf5, 0x4e, 0xb2, 0xb8, 0x2c, 0x0f, 0x00, 0xf9,
	0x22, 0x80, 0x31, 0x89, 0x0d, 0xec, 0x1d, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x90, 0x11,
	0xea, 0x22, 0x19, 0x01, 0x00, 0x00,
}
