// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: projects/gloo/api/v1/options/subset_spec.proto

package options

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SubsetSpec struct {
	Selectors            []*Selector `protobuf:"bytes,1,rep,name=selectors,proto3" json:"selectors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SubsetSpec) Reset()         { *m = SubsetSpec{} }
func (m *SubsetSpec) String() string { return proto.CompactTextString(m) }
func (*SubsetSpec) ProtoMessage()    {}
func (*SubsetSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d34062a0ca6b210, []int{0}
}
func (m *SubsetSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubsetSpec.Unmarshal(m, b)
}
func (m *SubsetSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubsetSpec.Marshal(b, m, deterministic)
}
func (m *SubsetSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubsetSpec.Merge(m, src)
}
func (m *SubsetSpec) XXX_Size() int {
	return xxx_messageInfo_SubsetSpec.Size(m)
}
func (m *SubsetSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_SubsetSpec.DiscardUnknown(m)
}

var xxx_messageInfo_SubsetSpec proto.InternalMessageInfo

func (m *SubsetSpec) GetSelectors() []*Selector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

type Selector struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Selector) Reset()         { *m = Selector{} }
func (m *Selector) String() string { return proto.CompactTextString(m) }
func (*Selector) ProtoMessage()    {}
func (*Selector) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d34062a0ca6b210, []int{1}
}
func (m *Selector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Selector.Unmarshal(m, b)
}
func (m *Selector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Selector.Marshal(b, m, deterministic)
}
func (m *Selector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Selector.Merge(m, src)
}
func (m *Selector) XXX_Size() int {
	return xxx_messageInfo_Selector.Size(m)
}
func (m *Selector) XXX_DiscardUnknown() {
	xxx_messageInfo_Selector.DiscardUnknown(m)
}

var xxx_messageInfo_Selector proto.InternalMessageInfo

func (m *Selector) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func init() {
	proto.RegisterType((*SubsetSpec)(nil), "options.gloo.solo.io.SubsetSpec")
	proto.RegisterType((*Selector)(nil), "options.gloo.solo.io.Selector")
}

func init() {
	proto.RegisterFile("projects/gloo/api/v1/options/subset_spec.proto", fileDescriptor_0d34062a0ca6b210)
}

var fileDescriptor_0d34062a0ca6b210 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2b, 0x28, 0xca, 0xcf,
	0x4a, 0x4d, 0x2e, 0x29, 0xd6, 0x4f, 0xcf, 0xc9, 0xcf, 0xd7, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x33,
	0xd4, 0xcf, 0x2f, 0x28, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2f, 0x2e, 0x4d, 0x2a, 0x4e, 0x2d, 0x89,
	0x2f, 0x2e, 0x48, 0x4d, 0x06, 0x29, 0x2c, 0xc9, 0x17, 0x12, 0x81, 0x4a, 0xe9, 0x81, 0x94, 0xeb,
	0x15, 0xe7, 0xe7, 0xe4, 0xeb, 0x65, 0xe6, 0x4b, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x15, 0xe8,
	0x83, 0x58, 0x10, 0xb5, 0x52, 0x42, 0xa9, 0x15, 0x25, 0x10, 0xc1, 0xd4, 0x8a, 0x12, 0x88, 0x98,
	0x92, 0x17, 0x17, 0x57, 0x30, 0xd8, 0xd0, 0xe0, 0x82, 0xd4, 0x64, 0x21, 0x1b, 0x2e, 0xce, 0xe2,
	0xd4, 0x9c, 0xd4, 0xe4, 0x92, 0xfc, 0xa2, 0x62, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x39,
	0x3d, 0x6c, 0x36, 0xe8, 0x05, 0x43, 0x95, 0x05, 0x21, 0x34, 0x28, 0xc9, 0x71, 0x71, 0xc0, 0x84,
	0x85, 0x84, 0xb8, 0x58, 0xb2, 0x53, 0x2b, 0x21, 0x86, 0x70, 0x06, 0x81, 0xd9, 0x4e, 0x4e, 0x3b,
	0xbe, 0xb2, 0x30, 0xae, 0x78, 0x24, 0xc7, 0x18, 0x65, 0x91, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4,
	0x97, 0x9c, 0x9f, 0xab, 0x0f, 0x32, 0x51, 0x37, 0x33, 0x1f, 0xe2, 0x5f, 0x54, 0xdf, 0x17, 0x64,
	0xa7, 0xa3, 0x85, 0x40, 0x12, 0x1b, 0xd8, 0xd9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf1,
	0xa3, 0xf8, 0x72, 0x28, 0x01, 0x00, 0x00,
}

func (this *SubsetSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SubsetSpec)
	if !ok {
		that2, ok := that.(SubsetSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Selectors) != len(that1.Selectors) {
		return false
	}
	for i := range this.Selectors {
		if !this.Selectors[i].Equal(that1.Selectors[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Selector) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Selector)
	if !ok {
		that2, ok := that.(Selector)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Keys) != len(that1.Keys) {
		return false
	}
	for i := range this.Keys {
		if this.Keys[i] != that1.Keys[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
