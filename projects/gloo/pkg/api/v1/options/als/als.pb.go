// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: projects/gloo/api/v1/options/als/als.proto

package als

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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

// Contains various settings for Envoy's access logging service.
// See here for more information: https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/filter/accesslog/v2/accesslog.proto#envoy-api-msg-config-filter-accesslog-v2-accesslog
type AccessLoggingService struct {
	AccessLog            []*AccessLog `protobuf:"bytes,1,rep,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AccessLoggingService) Reset()         { *m = AccessLoggingService{} }
func (m *AccessLoggingService) String() string { return proto.CompactTextString(m) }
func (*AccessLoggingService) ProtoMessage()    {}
func (*AccessLoggingService) Descriptor() ([]byte, []int) {
	return fileDescriptor_26db446233c13dd1, []int{0}
}
func (m *AccessLoggingService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLoggingService.Unmarshal(m, b)
}
func (m *AccessLoggingService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLoggingService.Marshal(b, m, deterministic)
}
func (m *AccessLoggingService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLoggingService.Merge(m, src)
}
func (m *AccessLoggingService) XXX_Size() int {
	return xxx_messageInfo_AccessLoggingService.Size(m)
}
func (m *AccessLoggingService) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLoggingService.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLoggingService proto.InternalMessageInfo

func (m *AccessLoggingService) GetAccessLog() []*AccessLog {
	if m != nil {
		return m.AccessLog
	}
	return nil
}

type AccessLog struct {
	// type of Access Logging service to implement
	//
	// Types that are valid to be assigned to OutputDestination:
	//	*AccessLog_FileSink
	//	*AccessLog_GrpcService
	OutputDestination    isAccessLog_OutputDestination `protobuf_oneof:"OutputDestination"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *AccessLog) Reset()         { *m = AccessLog{} }
func (m *AccessLog) String() string { return proto.CompactTextString(m) }
func (*AccessLog) ProtoMessage()    {}
func (*AccessLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_26db446233c13dd1, []int{1}
}
func (m *AccessLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessLog.Unmarshal(m, b)
}
func (m *AccessLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessLog.Marshal(b, m, deterministic)
}
func (m *AccessLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessLog.Merge(m, src)
}
func (m *AccessLog) XXX_Size() int {
	return xxx_messageInfo_AccessLog.Size(m)
}
func (m *AccessLog) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessLog.DiscardUnknown(m)
}

var xxx_messageInfo_AccessLog proto.InternalMessageInfo

type isAccessLog_OutputDestination interface {
	isAccessLog_OutputDestination()
	Equal(interface{}) bool
}

type AccessLog_FileSink struct {
	FileSink *FileSink `protobuf:"bytes,2,opt,name=file_sink,json=fileSink,proto3,oneof" json:"file_sink,omitempty"`
}
type AccessLog_GrpcService struct {
	GrpcService *GrpcService `protobuf:"bytes,3,opt,name=grpc_service,json=grpcService,proto3,oneof" json:"grpc_service,omitempty"`
}

func (*AccessLog_FileSink) isAccessLog_OutputDestination()    {}
func (*AccessLog_GrpcService) isAccessLog_OutputDestination() {}

func (m *AccessLog) GetOutputDestination() isAccessLog_OutputDestination {
	if m != nil {
		return m.OutputDestination
	}
	return nil
}

func (m *AccessLog) GetFileSink() *FileSink {
	if x, ok := m.GetOutputDestination().(*AccessLog_FileSink); ok {
		return x.FileSink
	}
	return nil
}

func (m *AccessLog) GetGrpcService() *GrpcService {
	if x, ok := m.GetOutputDestination().(*AccessLog_GrpcService); ok {
		return x.GrpcService
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AccessLog) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AccessLog_FileSink)(nil),
		(*AccessLog_GrpcService)(nil),
	}
}

type FileSink struct {
	// the file path to which the file access logging service will sink
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// the format which the logs should be outputted by
	//
	// Types that are valid to be assigned to OutputFormat:
	//	*FileSink_StringFormat
	//	*FileSink_JsonFormat
	OutputFormat         isFileSink_OutputFormat `protobuf_oneof:"output_format"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FileSink) Reset()         { *m = FileSink{} }
func (m *FileSink) String() string { return proto.CompactTextString(m) }
func (*FileSink) ProtoMessage()    {}
func (*FileSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_26db446233c13dd1, []int{2}
}
func (m *FileSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileSink.Unmarshal(m, b)
}
func (m *FileSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileSink.Marshal(b, m, deterministic)
}
func (m *FileSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileSink.Merge(m, src)
}
func (m *FileSink) XXX_Size() int {
	return xxx_messageInfo_FileSink.Size(m)
}
func (m *FileSink) XXX_DiscardUnknown() {
	xxx_messageInfo_FileSink.DiscardUnknown(m)
}

var xxx_messageInfo_FileSink proto.InternalMessageInfo

type isFileSink_OutputFormat interface {
	isFileSink_OutputFormat()
	Equal(interface{}) bool
}

type FileSink_StringFormat struct {
	StringFormat string `protobuf:"bytes,2,opt,name=string_format,json=stringFormat,proto3,oneof" json:"string_format,omitempty"`
}
type FileSink_JsonFormat struct {
	JsonFormat *types.Struct `protobuf:"bytes,3,opt,name=json_format,json=jsonFormat,proto3,oneof" json:"json_format,omitempty"`
}

func (*FileSink_StringFormat) isFileSink_OutputFormat() {}
func (*FileSink_JsonFormat) isFileSink_OutputFormat()   {}

func (m *FileSink) GetOutputFormat() isFileSink_OutputFormat {
	if m != nil {
		return m.OutputFormat
	}
	return nil
}

func (m *FileSink) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *FileSink) GetStringFormat() string {
	if x, ok := m.GetOutputFormat().(*FileSink_StringFormat); ok {
		return x.StringFormat
	}
	return ""
}

func (m *FileSink) GetJsonFormat() *types.Struct {
	if x, ok := m.GetOutputFormat().(*FileSink_JsonFormat); ok {
		return x.JsonFormat
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FileSink) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FileSink_StringFormat)(nil),
		(*FileSink_JsonFormat)(nil),
	}
}

type GrpcService struct {
	// name of log stream
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	// The static cluster defined in bootstrap config to route to
	//
	// Types that are valid to be assigned to ServiceRef:
	//	*GrpcService_StaticClusterName
	ServiceRef                      isGrpcService_ServiceRef `protobuf_oneof:"service_ref"`
	AdditionalRequestHeadersToLog   []string                 `protobuf:"bytes,4,rep,name=additional_request_headers_to_log,json=additionalRequestHeadersToLog,proto3" json:"additional_request_headers_to_log,omitempty"`
	AdditionalResponseHeadersToLog  []string                 `protobuf:"bytes,5,rep,name=additional_response_headers_to_log,json=additionalResponseHeadersToLog,proto3" json:"additional_response_headers_to_log,omitempty"`
	AdditionalResponseTrailersToLog []string                 `protobuf:"bytes,6,rep,name=additional_response_trailers_to_log,json=additionalResponseTrailersToLog,proto3" json:"additional_response_trailers_to_log,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}                 `json:"-"`
	XXX_unrecognized                []byte                   `json:"-"`
	XXX_sizecache                   int32                    `json:"-"`
}

func (m *GrpcService) Reset()         { *m = GrpcService{} }
func (m *GrpcService) String() string { return proto.CompactTextString(m) }
func (*GrpcService) ProtoMessage()    {}
func (*GrpcService) Descriptor() ([]byte, []int) {
	return fileDescriptor_26db446233c13dd1, []int{3}
}
func (m *GrpcService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcService.Unmarshal(m, b)
}
func (m *GrpcService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcService.Marshal(b, m, deterministic)
}
func (m *GrpcService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcService.Merge(m, src)
}
func (m *GrpcService) XXX_Size() int {
	return xxx_messageInfo_GrpcService.Size(m)
}
func (m *GrpcService) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcService.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcService proto.InternalMessageInfo

type isGrpcService_ServiceRef interface {
	isGrpcService_ServiceRef()
	Equal(interface{}) bool
}

type GrpcService_StaticClusterName struct {
	StaticClusterName string `protobuf:"bytes,2,opt,name=static_cluster_name,json=staticClusterName,proto3,oneof" json:"static_cluster_name,omitempty"`
}

func (*GrpcService_StaticClusterName) isGrpcService_ServiceRef() {}

func (m *GrpcService) GetServiceRef() isGrpcService_ServiceRef {
	if m != nil {
		return m.ServiceRef
	}
	return nil
}

func (m *GrpcService) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

func (m *GrpcService) GetStaticClusterName() string {
	if x, ok := m.GetServiceRef().(*GrpcService_StaticClusterName); ok {
		return x.StaticClusterName
	}
	return ""
}

func (m *GrpcService) GetAdditionalRequestHeadersToLog() []string {
	if m != nil {
		return m.AdditionalRequestHeadersToLog
	}
	return nil
}

func (m *GrpcService) GetAdditionalResponseHeadersToLog() []string {
	if m != nil {
		return m.AdditionalResponseHeadersToLog
	}
	return nil
}

func (m *GrpcService) GetAdditionalResponseTrailersToLog() []string {
	if m != nil {
		return m.AdditionalResponseTrailersToLog
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GrpcService) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GrpcService_StaticClusterName)(nil),
	}
}

func init() {
	proto.RegisterType((*AccessLoggingService)(nil), "als.options.gloo.solo.io.AccessLoggingService")
	proto.RegisterType((*AccessLog)(nil), "als.options.gloo.solo.io.AccessLog")
	proto.RegisterType((*FileSink)(nil), "als.options.gloo.solo.io.FileSink")
	proto.RegisterType((*GrpcService)(nil), "als.options.gloo.solo.io.GrpcService")
}

func init() {
	proto.RegisterFile("projects/gloo/api/v1/options/als/als.proto", fileDescriptor_26db446233c13dd1)
}

var fileDescriptor_26db446233c13dd1 = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0xe3, 0x26, 0x5f, 0xbf, 0x78, 0xdd, 0x08, 0x75, 0x5b, 0x89, 0xb4, 0x82, 0x12, 0x5c,
	0x55, 0x8a, 0x90, 0xb0, 0xa1, 0xdc, 0x10, 0x97, 0x06, 0x14, 0xac, 0xaa, 0x02, 0xc9, 0xe9, 0xa9,
	0x17, 0x6b, 0xe3, 0xac, 0x37, 0xdb, 0x6c, 0x3c, 0x66, 0x77, 0x5d, 0xf5, 0x39, 0x78, 0x0a, 0x6e,
	0x5c, 0x79, 0x09, 0x5e, 0x82, 0x77, 0xe0, 0x8e, 0xbc, 0x6b, 0xb7, 0x09, 0x34, 0x87, 0x48, 0xe3,
	0x99, 0xff, 0xfc, 0x76, 0x46, 0xff, 0x0c, 0x7a, 0x51, 0x48, 0xb8, 0xa6, 0xa9, 0x56, 0x21, 0x13,
	0x00, 0x21, 0x29, 0x78, 0x78, 0xf3, 0x3a, 0x84, 0x42, 0x73, 0xc8, 0x55, 0x48, 0x84, 0xf9, 0x05,
	0x85, 0x04, 0x0d, 0xb8, 0x5f, 0x85, 0x75, 0x29, 0xa8, 0xe4, 0x81, 0x02, 0x01, 0x01, 0x87, 0xc3,
	0x7d, 0x06, 0x0c, 0x8c, 0x28, 0xac, 0x22, 0xab, 0x3f, 0xc4, 0xf4, 0x56, 0xdb, 0x24, 0xbd, 0xd5,
	0x75, 0xee, 0xa0, 0x6a, 0x79, 0xb9, 0xe0, 0xba, 0x79, 0x4a, 0xd2, 0xac, 0x2e, 0x3d, 0x61, 0x00,
	0x4c, 0xd0, 0xd0, 0x7c, 0x4d, 0xcb, 0x2c, 0x54, 0x5a, 0x96, 0x69, 0xdd, 0xe8, 0x5f, 0xa1, 0xfd,
	0xb3, 0x34, 0xa5, 0x4a, 0x5d, 0x00, 0x63, 0x3c, 0x67, 0x13, 0x2a, 0x6f, 0x78, 0x4a, 0xf1, 0x08,
	0x21, 0x62, 0xf2, 0x89, 0x00, 0xd6, 0x77, 0x06, 0xed, 0xa1, 0x77, 0x7a, 0x1c, 0x6c, 0x9a, 0x34,
	0xb8, 0x63, 0xc4, 0x2e, 0x69, 0x42, 0xff, 0xbb, 0x83, 0xdc, 0xbb, 0x02, 0x3e, 0x43, 0x6e, 0xc6,
	0x05, 0x4d, 0x14, 0xcf, 0x17, 0xfd, 0xad, 0x81, 0x33, 0xf4, 0x4e, 0xfd, 0xcd, 0xc0, 0x31, 0x17,
	0x74, 0xc2, 0xf3, 0x45, 0xd4, 0x8a, 0xbb, 0x59, 0x1d, 0xe3, 0x73, 0xb4, 0xc3, 0x64, 0x91, 0x26,
	0xca, 0x0e, 0xd9, 0x6f, 0x1b, 0xca, 0xc9, 0x66, 0xca, 0x47, 0x59, 0xa4, 0xf5, 0x46, 0x51, 0x2b,
	0xf6, 0xd8, 0xfd, 0xe7, 0x68, 0x0f, 0xed, 0x7e, 0x2e, 0x75, 0x51, 0xea, 0x0f, 0x54, 0x69, 0x9e,
	0x93, 0xaa, 0xdb, 0xff, 0xea, 0xa0, 0x6e, 0xf3, 0x32, 0xc6, 0xa8, 0x53, 0x10, 0x3d, 0xef, 0x3b,
	0x03, 0x67, 0xe8, 0xc6, 0x26, 0xc6, 0x27, 0xa8, 0xa7, 0xb4, 0xe4, 0x39, 0x4b, 0x32, 0x90, 0x4b,
	0xa2, 0xcd, 0x22, 0x6e, 0xd4, 0x8a, 0x77, 0x6c, 0x7a, 0x6c, 0xb2, 0xf8, 0x2d, 0xf2, 0xae, 0x15,
	0xe4, 0x8d, 0xc8, 0xce, 0xf9, 0x38, 0xb0, 0x4e, 0x04, 0x8d, 0x13, 0xc1, 0xc4, 0x38, 0x11, 0xb5,
	0x62, 0x54, 0xa9, 0x6d, 0xef, 0xe8, 0x11, 0xea, 0x81, 0x19, 0xac, 0xee, 0xf6, 0x7f, 0x6e, 0x21,
	0x6f, 0x65, 0x11, 0x7c, 0x80, 0xba, 0x02, 0x58, 0x92, 0x93, 0x25, 0xad, 0x67, 0xfb, 0x5f, 0x00,
	0xfb, 0x44, 0x96, 0x14, 0xbf, 0x42, 0x7b, 0x4a, 0x13, 0xcd, 0xd3, 0x24, 0x15, 0xa5, 0xd2, 0x54,
	0x5a, 0x55, 0x33, 0xe4, 0xae, 0x2d, 0xbe, 0xb7, 0x35, 0xd3, 0x11, 0xa1, 0xe7, 0x64, 0x36, 0xe3,
	0xd5, 0xf6, 0x44, 0x24, 0x92, 0x7e, 0x29, 0xa9, 0xd2, 0xc9, 0x9c, 0x92, 0x19, 0x95, 0x2a, 0xd1,
	0x60, 0xec, 0xef, 0x0c, 0xda, 0x43, 0x37, 0x7e, 0x7a, 0x2f, 0x8c, 0xad, 0x2e, 0xb2, 0xb2, 0x4b,
	0xa8, 0xfc, 0x3d, 0x47, 0xfe, 0x1a, 0x49, 0x15, 0x90, 0x2b, 0xfa, 0x37, 0xea, 0x3f, 0x83, 0x3a,
	0x5a, 0x45, 0x59, 0xe1, 0x1a, 0xeb, 0x02, 0x1d, 0x3f, 0xc4, 0xd2, 0x92, 0x70, 0xb1, 0x02, 0xdb,
	0x36, 0xb0, 0x67, 0xff, 0xc2, 0x2e, 0x6b, 0xa1, 0xa1, 0x8d, 0x7a, 0xc8, 0xab, 0xff, 0x31, 0x89,
	0xa4, 0xd9, 0x68, 0xfc, 0xe3, 0x77, 0xc7, 0xf9, 0xf6, 0xeb, 0xc8, 0xb9, 0x7a, 0xc7, 0xb8, 0x9e,
	0x97, 0xd3, 0x20, 0x85, 0x65, 0x68, 0x0e, 0x88, 0x83, 0xbd, 0xd7, 0xf5, 0xeb, 0x2d, 0x16, 0xec,
	0x81, 0x0b, 0x9e, 0x6e, 0x1b, 0x1f, 0xdf, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x09, 0xdc,
	0xe9, 0xec, 0x03, 0x00, 0x00,
}

func (this *AccessLoggingService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessLoggingService)
	if !ok {
		that2, ok := that.(AccessLoggingService)
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
	if len(this.AccessLog) != len(that1.AccessLog) {
		return false
	}
	for i := range this.AccessLog {
		if !this.AccessLog[i].Equal(that1.AccessLog[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AccessLog) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessLog)
	if !ok {
		that2, ok := that.(AccessLog)
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
	if that1.OutputDestination == nil {
		if this.OutputDestination != nil {
			return false
		}
	} else if this.OutputDestination == nil {
		return false
	} else if !this.OutputDestination.Equal(that1.OutputDestination) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AccessLog_FileSink) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessLog_FileSink)
	if !ok {
		that2, ok := that.(AccessLog_FileSink)
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
	if !this.FileSink.Equal(that1.FileSink) {
		return false
	}
	return true
}
func (this *AccessLog_GrpcService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AccessLog_GrpcService)
	if !ok {
		that2, ok := that.(AccessLog_GrpcService)
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
	if !this.GrpcService.Equal(that1.GrpcService) {
		return false
	}
	return true
}
func (this *FileSink) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FileSink)
	if !ok {
		that2, ok := that.(FileSink)
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
	if this.Path != that1.Path {
		return false
	}
	if that1.OutputFormat == nil {
		if this.OutputFormat != nil {
			return false
		}
	} else if this.OutputFormat == nil {
		return false
	} else if !this.OutputFormat.Equal(that1.OutputFormat) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *FileSink_StringFormat) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FileSink_StringFormat)
	if !ok {
		that2, ok := that.(FileSink_StringFormat)
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
	if this.StringFormat != that1.StringFormat {
		return false
	}
	return true
}
func (this *FileSink_JsonFormat) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FileSink_JsonFormat)
	if !ok {
		that2, ok := that.(FileSink_JsonFormat)
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
	if !this.JsonFormat.Equal(that1.JsonFormat) {
		return false
	}
	return true
}
func (this *GrpcService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GrpcService)
	if !ok {
		that2, ok := that.(GrpcService)
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
	if this.LogName != that1.LogName {
		return false
	}
	if that1.ServiceRef == nil {
		if this.ServiceRef != nil {
			return false
		}
	} else if this.ServiceRef == nil {
		return false
	} else if !this.ServiceRef.Equal(that1.ServiceRef) {
		return false
	}
	if len(this.AdditionalRequestHeadersToLog) != len(that1.AdditionalRequestHeadersToLog) {
		return false
	}
	for i := range this.AdditionalRequestHeadersToLog {
		if this.AdditionalRequestHeadersToLog[i] != that1.AdditionalRequestHeadersToLog[i] {
			return false
		}
	}
	if len(this.AdditionalResponseHeadersToLog) != len(that1.AdditionalResponseHeadersToLog) {
		return false
	}
	for i := range this.AdditionalResponseHeadersToLog {
		if this.AdditionalResponseHeadersToLog[i] != that1.AdditionalResponseHeadersToLog[i] {
			return false
		}
	}
	if len(this.AdditionalResponseTrailersToLog) != len(that1.AdditionalResponseTrailersToLog) {
		return false
	}
	for i := range this.AdditionalResponseTrailersToLog {
		if this.AdditionalResponseTrailersToLog[i] != that1.AdditionalResponseTrailersToLog[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *GrpcService_StaticClusterName) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GrpcService_StaticClusterName)
	if !ok {
		that2, ok := that.(GrpcService_StaticClusterName)
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
	if this.StaticClusterName != that1.StaticClusterName {
		return false
	}
	return true
}
