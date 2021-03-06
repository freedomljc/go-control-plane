// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/v3alpha/number.proto

package envoy_type_matcher_v3alpha

import (
	fmt "fmt"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/type/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type DoubleMatcher struct {
	// Types that are valid to be assigned to MatchPattern:
	//	*DoubleMatcher_Range
	//	*DoubleMatcher_Exact
	MatchPattern         isDoubleMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *DoubleMatcher) Reset()         { *m = DoubleMatcher{} }
func (m *DoubleMatcher) String() string { return proto.CompactTextString(m) }
func (*DoubleMatcher) ProtoMessage()    {}
func (*DoubleMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_50221d5482f481a5, []int{0}
}

func (m *DoubleMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleMatcher.Unmarshal(m, b)
}
func (m *DoubleMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleMatcher.Marshal(b, m, deterministic)
}
func (m *DoubleMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleMatcher.Merge(m, src)
}
func (m *DoubleMatcher) XXX_Size() int {
	return xxx_messageInfo_DoubleMatcher.Size(m)
}
func (m *DoubleMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleMatcher proto.InternalMessageInfo

type isDoubleMatcher_MatchPattern interface {
	isDoubleMatcher_MatchPattern()
}

type DoubleMatcher_Range struct {
	Range *v3alpha.DoubleRange `protobuf:"bytes,1,opt,name=range,proto3,oneof"`
}

type DoubleMatcher_Exact struct {
	Exact float64 `protobuf:"fixed64,2,opt,name=exact,proto3,oneof"`
}

func (*DoubleMatcher_Range) isDoubleMatcher_MatchPattern() {}

func (*DoubleMatcher_Exact) isDoubleMatcher_MatchPattern() {}

func (m *DoubleMatcher) GetMatchPattern() isDoubleMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *DoubleMatcher) GetRange() *v3alpha.DoubleRange {
	if x, ok := m.GetMatchPattern().(*DoubleMatcher_Range); ok {
		return x.Range
	}
	return nil
}

func (m *DoubleMatcher) GetExact() float64 {
	if x, ok := m.GetMatchPattern().(*DoubleMatcher_Exact); ok {
		return x.Exact
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DoubleMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DoubleMatcher_Range)(nil),
		(*DoubleMatcher_Exact)(nil),
	}
}

func init() {
	proto.RegisterType((*DoubleMatcher)(nil), "envoy.type.matcher.v3alpha.DoubleMatcher")
}

func init() {
	proto.RegisterFile("envoy/type/matcher/v3alpha/number.proto", fileDescriptor_50221d5482f481a5)
}

var fileDescriptor_50221d5482f481a5 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0x4d, 0x2c, 0x49, 0xce, 0x48, 0x2d, 0xd2, 0x2f,
	0x33, 0x4e, 0xcc, 0x29, 0xc8, 0x48, 0xd4, 0xcf, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0xd2, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x92, 0x02, 0x2b, 0xd4, 0x03, 0x29, 0xd4, 0x83, 0x2a, 0xd4, 0x83, 0x2a,
	0x94, 0x92, 0x43, 0x32, 0x04, 0xa6, 0xb9, 0x28, 0x31, 0x2f, 0x3d, 0x15, 0xa2, 0x57, 0x4a, 0xbc,
	0x2c, 0x31, 0x27, 0x33, 0x25, 0xb1, 0x24, 0x55, 0x1f, 0xc6, 0x80, 0x48, 0x28, 0x95, 0x71, 0xf1,
	0xba, 0xe4, 0x97, 0x26, 0xe5, 0xa4, 0xfa, 0x42, 0x4c, 0x14, 0x32, 0xe7, 0x62, 0x05, 0x6b, 0x94,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x92, 0xd7, 0x43, 0xb2, 0x15, 0x6a, 0xb2, 0x1e, 0x44, 0x47,
	0x10, 0x48, 0x99, 0x07, 0x43, 0x10, 0x44, 0xbd, 0x90, 0x18, 0x17, 0x6b, 0x6a, 0x45, 0x62, 0x72,
	0x89, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x23, 0x48, 0x1c, 0xcc, 0x75, 0x12, 0xe1, 0xe2, 0x05, 0xbb,
	0x36, 0xbe, 0x20, 0xb1, 0xa4, 0x24, 0xb5, 0x28, 0x4f, 0x88, 0xf9, 0x87, 0x13, 0xa3, 0x93, 0x25,
	0x97, 0x46, 0x66, 0x3e, 0xc4, 0xec, 0x82, 0xa2, 0xfc, 0x8a, 0x4a, 0x3d, 0xdc, 0x9e, 0x73, 0xe2,
	0xf6, 0x03, 0x07, 0x43, 0x00, 0xc8, 0xc1, 0x01, 0x8c, 0x49, 0x6c, 0x60, 0x97, 0x1b, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xdd, 0x8c, 0x57, 0x82, 0x39, 0x01, 0x00, 0x00,
}
