// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: projects/gloo/api/external/envoy/extensions/jwt/solo_jwt_authn.proto

package jwt

import (
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
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

type SoloJwtAuthnPerRoute struct {
	Requirement     string                                          `protobuf:"bytes,1,opt,name=requirement,proto3" json:"requirement,omitempty"`
	ClaimsToHeaders map[string]*SoloJwtAuthnPerRoute_ClaimToHeaders `protobuf:"bytes,2,rep,name=claims_to_headers,json=claimsToHeaders,proto3" json:"claims_to_headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// clear the route cache if claims were added to the header
	ClearRouteCache      bool     `protobuf:"varint,3,opt,name=clear_route_cache,json=clearRouteCache,proto3" json:"clear_route_cache,omitempty"`
	PayloadInMetadata    string   `protobuf:"bytes,4,opt,name=payload_in_metadata,json=payloadInMetadata,proto3" json:"payload_in_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoloJwtAuthnPerRoute) Reset()         { *m = SoloJwtAuthnPerRoute{} }
func (m *SoloJwtAuthnPerRoute) String() string { return proto.CompactTextString(m) }
func (*SoloJwtAuthnPerRoute) ProtoMessage()    {}
func (*SoloJwtAuthnPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_d64a2643ddeeddde, []int{0}
}
func (m *SoloJwtAuthnPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoloJwtAuthnPerRoute.Unmarshal(m, b)
}
func (m *SoloJwtAuthnPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoloJwtAuthnPerRoute.Marshal(b, m, deterministic)
}
func (m *SoloJwtAuthnPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoloJwtAuthnPerRoute.Merge(m, src)
}
func (m *SoloJwtAuthnPerRoute) XXX_Size() int {
	return xxx_messageInfo_SoloJwtAuthnPerRoute.Size(m)
}
func (m *SoloJwtAuthnPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_SoloJwtAuthnPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_SoloJwtAuthnPerRoute proto.InternalMessageInfo

func (m *SoloJwtAuthnPerRoute) GetRequirement() string {
	if m != nil {
		return m.Requirement
	}
	return ""
}

func (m *SoloJwtAuthnPerRoute) GetClaimsToHeaders() map[string]*SoloJwtAuthnPerRoute_ClaimToHeaders {
	if m != nil {
		return m.ClaimsToHeaders
	}
	return nil
}

func (m *SoloJwtAuthnPerRoute) GetClearRouteCache() bool {
	if m != nil {
		return m.ClearRouteCache
	}
	return false
}

func (m *SoloJwtAuthnPerRoute) GetPayloadInMetadata() string {
	if m != nil {
		return m.PayloadInMetadata
	}
	return ""
}

// If this is specified, one of the claims will be copied to a header
// and the route cache will be cleared.
type SoloJwtAuthnPerRoute_ClaimToHeader struct {
	Claim                string   `protobuf:"bytes,1,opt,name=claim,proto3" json:"claim,omitempty"`
	Header               string   `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	Append               bool     `protobuf:"varint,3,opt,name=append,proto3" json:"append,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoloJwtAuthnPerRoute_ClaimToHeader) Reset()         { *m = SoloJwtAuthnPerRoute_ClaimToHeader{} }
func (m *SoloJwtAuthnPerRoute_ClaimToHeader) String() string { return proto.CompactTextString(m) }
func (*SoloJwtAuthnPerRoute_ClaimToHeader) ProtoMessage()    {}
func (*SoloJwtAuthnPerRoute_ClaimToHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_d64a2643ddeeddde, []int{0, 0}
}
func (m *SoloJwtAuthnPerRoute_ClaimToHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeader.Unmarshal(m, b)
}
func (m *SoloJwtAuthnPerRoute_ClaimToHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeader.Marshal(b, m, deterministic)
}
func (m *SoloJwtAuthnPerRoute_ClaimToHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeader.Merge(m, src)
}
func (m *SoloJwtAuthnPerRoute_ClaimToHeader) XXX_Size() int {
	return xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeader.Size(m)
}
func (m *SoloJwtAuthnPerRoute_ClaimToHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeader.DiscardUnknown(m)
}

var xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeader proto.InternalMessageInfo

func (m *SoloJwtAuthnPerRoute_ClaimToHeader) GetClaim() string {
	if m != nil {
		return m.Claim
	}
	return ""
}

func (m *SoloJwtAuthnPerRoute_ClaimToHeader) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

func (m *SoloJwtAuthnPerRoute_ClaimToHeader) GetAppend() bool {
	if m != nil {
		return m.Append
	}
	return false
}

type SoloJwtAuthnPerRoute_ClaimToHeaders struct {
	Claims               []*SoloJwtAuthnPerRoute_ClaimToHeader `protobuf:"bytes,1,rep,name=claims,proto3" json:"claims,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *SoloJwtAuthnPerRoute_ClaimToHeaders) Reset()         { *m = SoloJwtAuthnPerRoute_ClaimToHeaders{} }
func (m *SoloJwtAuthnPerRoute_ClaimToHeaders) String() string { return proto.CompactTextString(m) }
func (*SoloJwtAuthnPerRoute_ClaimToHeaders) ProtoMessage()    {}
func (*SoloJwtAuthnPerRoute_ClaimToHeaders) Descriptor() ([]byte, []int) {
	return fileDescriptor_d64a2643ddeeddde, []int{0, 1}
}
func (m *SoloJwtAuthnPerRoute_ClaimToHeaders) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeaders.Unmarshal(m, b)
}
func (m *SoloJwtAuthnPerRoute_ClaimToHeaders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeaders.Marshal(b, m, deterministic)
}
func (m *SoloJwtAuthnPerRoute_ClaimToHeaders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeaders.Merge(m, src)
}
func (m *SoloJwtAuthnPerRoute_ClaimToHeaders) XXX_Size() int {
	return xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeaders.Size(m)
}
func (m *SoloJwtAuthnPerRoute_ClaimToHeaders) XXX_DiscardUnknown() {
	xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeaders.DiscardUnknown(m)
}

var xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeaders proto.InternalMessageInfo

func (m *SoloJwtAuthnPerRoute_ClaimToHeaders) GetClaims() []*SoloJwtAuthnPerRoute_ClaimToHeader {
	if m != nil {
		return m.Claims
	}
	return nil
}

func init() {
	proto.RegisterType((*SoloJwtAuthnPerRoute)(nil), "envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute")
	proto.RegisterMapType((map[string]*SoloJwtAuthnPerRoute_ClaimToHeaders)(nil), "envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimsToHeadersEntry")
	proto.RegisterType((*SoloJwtAuthnPerRoute_ClaimToHeader)(nil), "envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader")
	proto.RegisterType((*SoloJwtAuthnPerRoute_ClaimToHeaders)(nil), "envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders")
}

func init() {
	proto.RegisterFile("projects/gloo/api/external/envoy/extensions/jwt/solo_jwt_authn.proto", fileDescriptor_d64a2643ddeeddde)
}

var fileDescriptor_d64a2643ddeeddde = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0x96, 0x93, 0x36, 0xfa, 0xff, 0x89, 0xb8, 0x64, 0x88, 0xc0, 0xca, 0x2a, 0x62, 0x15, 0x55,
	0x62, 0x2c, 0x85, 0x4d, 0xc5, 0x8e, 0x16, 0x24, 0x2e, 0x02, 0x2a, 0x43, 0x37, 0x6c, 0xac, 0xa9,
	0x7d, 0x12, 0x4f, 0x3a, 0x99, 0x33, 0x8c, 0x8f, 0x93, 0x78, 0xc9, 0x33, 0xb0, 0xe5, 0x5d, 0x41,
	0x9e, 0x71, 0x24, 0x82, 0x2a, 0x01, 0x52, 0x77, 0x73, 0x2e, 0xfe, 0x2e, 0x47, 0xfe, 0xd8, 0x0b,
	0xeb, 0x70, 0x05, 0x39, 0x55, 0xc9, 0x52, 0x23, 0x26, 0xd2, 0xaa, 0x04, 0x76, 0x04, 0xce, 0x48,
	0x9d, 0x80, 0xd9, 0x60, 0xe3, 0x4b, 0x53, 0x29, 0x34, 0x55, 0xb2, 0xda, 0x52, 0x52, 0xa1, 0xc6,
	0x6c, 0xb5, 0xa5, 0x4c, 0xd6, 0x54, 0x1a, 0x61, 0x1d, 0x12, 0xf2, 0x13, 0xbf, 0x2a, 0x72, 0x34,
	0x0b, 0xb5, 0x14, 0x0b, 0xa5, 0x09, 0x9c, 0x28, 0x89, 0xac, 0xf8, 0x6d, 0x7d, 0x33, 0x9f, 0x3c,
	0xda, 0x48, 0xad, 0x0a, 0x49, 0x90, 0xec, 0x1f, 0x01, 0xe4, 0xf1, 0x8f, 0x23, 0x36, 0xfe, 0x88,
	0x1a, 0xdf, 0x6c, 0xe9, 0x79, 0xbb, 0x7c, 0x01, 0x2e, 0xc5, 0x9a, 0x80, 0x4f, 0xd9, 0xd0, 0xc1,
	0x97, 0x5a, 0x39, 0x58, 0x83, 0xa1, 0x38, 0x9a, 0x46, 0xb3, 0xff, 0xd3, 0x5f, 0x5b, 0xfc, 0x6b,
	0xc4, 0x46, 0xb9, 0x96, 0x6a, 0x5d, 0x65, 0x84, 0x59, 0x09, 0xb2, 0x00, 0x57, 0xc5, 0xbd, 0x69,
	0x7f, 0x36, 0x9c, 0x5f, 0x8a, 0xbf, 0x17, 0x27, 0x6e, 0xe2, 0x17, 0xe7, 0x1e, 0xf9, 0x13, 0xbe,
	0x0a, 0xb8, 0x2f, 0x0d, 0xb9, 0x26, 0xbd, 0x97, 0x1f, 0x76, 0xf9, 0x49, 0x2b, 0x01, 0xa4, 0xcb,
	0x5c, 0xfb, 0x51, 0x96, 0xcb, 0xbc, 0x84, 0xb8, 0x3f, 0x8d, 0x66, 0xff, 0xb5, 0xbb, 0x20, 0x03,
	0xd8, 0x79, 0xdb, 0xe6, 0x82, 0x3d, 0xb0, 0xb2, 0xd1, 0x28, 0x8b, 0x4c, 0x99, 0x6c, 0x0d, 0x24,
	0x0b, 0x49, 0x32, 0x3e, 0xf2, 0xce, 0x46, 0xdd, 0xe8, 0xb5, 0x79, 0xd7, 0x0d, 0x26, 0x97, 0xec,
	0x8e, 0x17, 0xb1, 0x67, 0xe3, 0x63, 0x76, 0xec, 0xf9, 0xbb, 0x63, 0x84, 0x82, 0x3f, 0x64, 0x83,
	0xe0, 0x3d, 0xee, 0xf9, 0x76, 0x57, 0xb5, 0x7d, 0x69, 0x2d, 0x98, 0xa2, 0xd3, 0xd3, 0x55, 0x93,
	0x1d, 0xbb, 0x7b, 0x00, 0x5b, 0xf1, 0x05, 0x1b, 0x04, 0x5f, 0x71, 0xe4, 0x8f, 0xf7, 0xfe, 0x76,
	0x8e, 0xb7, 0x27, 0x48, 0x3b, 0xf4, 0xc9, 0xb7, 0x88, 0x8d, 0x6f, 0x3a, 0x2b, 0xbf, 0xcf, 0xfa,
	0xd7, 0xd0, 0x74, 0xb6, 0xda, 0x27, 0x07, 0x76, 0xbc, 0x91, 0xba, 0x06, 0xef, 0x69, 0x38, 0xff,
	0x70, 0xbb, 0x8a, 0xaa, 0x34, 0xa0, 0x3f, 0xeb, 0x9d, 0x46, 0x67, 0xdf, 0x23, 0x76, 0xaa, 0x30,
	0x10, 0x58, 0x87, 0xbb, 0xe6, 0x1f, 0xb8, 0xce, 0x46, 0x07, 0x64, 0xed, 0x1f, 0x7d, 0x11, 0x7d,
	0x7e, 0xbb, 0x54, 0x54, 0xd6, 0x57, 0x22, 0xc7, 0xb5, 0x4f, 0xce, 0x13, 0x85, 0x21, 0x68, 0x87,
	0xb1, 0xb3, 0xd7, 0xcb, 0x3f, 0x47, 0xef, 0x6a, 0xe0, 0x73, 0xf2, 0xf4, 0x67, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xbd, 0x43, 0xc5, 0xce, 0xb4, 0x03, 0x00, 0x00,
}
