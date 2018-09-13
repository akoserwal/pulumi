// Code generated by protoc-gen-go. DO NOT EDIT.
// source: engine.proto

package pulumirpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/golang/protobuf/ptypes/struct"

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

// LogSeverity is the severity level of a log message.  Errors are fatal; all others are informational.
type LogSeverity int32

const (
	LogSeverity_DEBUG   LogSeverity = 0
	LogSeverity_INFO    LogSeverity = 1
	LogSeverity_WARNING LogSeverity = 2
	LogSeverity_ERROR   LogSeverity = 3
)

var LogSeverity_name = map[int32]string{
	0: "DEBUG",
	1: "INFO",
	2: "WARNING",
	3: "ERROR",
}
var LogSeverity_value = map[string]int32{
	"DEBUG":   0,
	"INFO":    1,
	"WARNING": 2,
	"ERROR":   3,
}

func (x LogSeverity) String() string {
	return proto.EnumName(LogSeverity_name, int32(x))
}
func (LogSeverity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_engine_f09f44fe55773ec3, []int{0}
}

type LogRequest struct {
	// the logging level of this message.
	Severity LogSeverity `protobuf:"varint,1,opt,name=severity,enum=pulumirpc.LogSeverity" json:"severity,omitempty"`
	// the contents of the logged message.
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	// the (optional) resource urn this log is associated with.
	Urn string `protobuf:"bytes,3,opt,name=urn" json:"urn,omitempty"`
	// the (optional) stream id that a stream of log messages can be associated with. This allows
	// clients to not have to buffer a large set of log messages that they all want to be
	// conceptually connected.  Instead the messages can be sent as chunks (with the same stream id)
	// and the end display can show the messages as they arrive, while still stitching them together
	// into one total log message.
	//
	// 0/not-given means: do not associate with any stream.
	StreamId int32 `protobuf:"varint,4,opt,name=streamId" json:"streamId,omitempty"`
	// Optional value indicating whether this is a status message.
	Ephemeral            bool     `protobuf:"varint,5,opt,name=ephemeral" json:"ephemeral,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogRequest) Reset()         { *m = LogRequest{} }
func (m *LogRequest) String() string { return proto.CompactTextString(m) }
func (*LogRequest) ProtoMessage()    {}
func (*LogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_engine_f09f44fe55773ec3, []int{0}
}
func (m *LogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRequest.Unmarshal(m, b)
}
func (m *LogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRequest.Marshal(b, m, deterministic)
}
func (dst *LogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRequest.Merge(dst, src)
}
func (m *LogRequest) XXX_Size() int {
	return xxx_messageInfo_LogRequest.Size(m)
}
func (m *LogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogRequest proto.InternalMessageInfo

func (m *LogRequest) GetSeverity() LogSeverity {
	if m != nil {
		return m.Severity
	}
	return LogSeverity_DEBUG
}

func (m *LogRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *LogRequest) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *LogRequest) GetStreamId() int32 {
	if m != nil {
		return m.StreamId
	}
	return 0
}

func (m *LogRequest) GetEphemeral() bool {
	if m != nil {
		return m.Ephemeral
	}
	return false
}

type GetRootResourceRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRootResourceRequest) Reset()         { *m = GetRootResourceRequest{} }
func (m *GetRootResourceRequest) String() string { return proto.CompactTextString(m) }
func (*GetRootResourceRequest) ProtoMessage()    {}
func (*GetRootResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_engine_f09f44fe55773ec3, []int{1}
}
func (m *GetRootResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRootResourceRequest.Unmarshal(m, b)
}
func (m *GetRootResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRootResourceRequest.Marshal(b, m, deterministic)
}
func (dst *GetRootResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRootResourceRequest.Merge(dst, src)
}
func (m *GetRootResourceRequest) XXX_Size() int {
	return xxx_messageInfo_GetRootResourceRequest.Size(m)
}
func (m *GetRootResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRootResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRootResourceRequest proto.InternalMessageInfo

type GetRootResourceResponse struct {
	Urn                  string   `protobuf:"bytes,1,opt,name=urn" json:"urn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRootResourceResponse) Reset()         { *m = GetRootResourceResponse{} }
func (m *GetRootResourceResponse) String() string { return proto.CompactTextString(m) }
func (*GetRootResourceResponse) ProtoMessage()    {}
func (*GetRootResourceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_engine_f09f44fe55773ec3, []int{2}
}
func (m *GetRootResourceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRootResourceResponse.Unmarshal(m, b)
}
func (m *GetRootResourceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRootResourceResponse.Marshal(b, m, deterministic)
}
func (dst *GetRootResourceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRootResourceResponse.Merge(dst, src)
}
func (m *GetRootResourceResponse) XXX_Size() int {
	return xxx_messageInfo_GetRootResourceResponse.Size(m)
}
func (m *GetRootResourceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRootResourceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRootResourceResponse proto.InternalMessageInfo

func (m *GetRootResourceResponse) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

type SetRootResourceRequest struct {
	Urn                  string   `protobuf:"bytes,1,opt,name=urn" json:"urn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRootResourceRequest) Reset()         { *m = SetRootResourceRequest{} }
func (m *SetRootResourceRequest) String() string { return proto.CompactTextString(m) }
func (*SetRootResourceRequest) ProtoMessage()    {}
func (*SetRootResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_engine_f09f44fe55773ec3, []int{3}
}
func (m *SetRootResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRootResourceRequest.Unmarshal(m, b)
}
func (m *SetRootResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRootResourceRequest.Marshal(b, m, deterministic)
}
func (dst *SetRootResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRootResourceRequest.Merge(dst, src)
}
func (m *SetRootResourceRequest) XXX_Size() int {
	return xxx_messageInfo_SetRootResourceRequest.Size(m)
}
func (m *SetRootResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRootResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRootResourceRequest proto.InternalMessageInfo

func (m *SetRootResourceRequest) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

type SetRootResourceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRootResourceResponse) Reset()         { *m = SetRootResourceResponse{} }
func (m *SetRootResourceResponse) String() string { return proto.CompactTextString(m) }
func (*SetRootResourceResponse) ProtoMessage()    {}
func (*SetRootResourceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_engine_f09f44fe55773ec3, []int{4}
}
func (m *SetRootResourceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRootResourceResponse.Unmarshal(m, b)
}
func (m *SetRootResourceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRootResourceResponse.Marshal(b, m, deterministic)
}
func (dst *SetRootResourceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRootResourceResponse.Merge(dst, src)
}
func (m *SetRootResourceResponse) XXX_Size() int {
	return xxx_messageInfo_SetRootResourceResponse.Size(m)
}
func (m *SetRootResourceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRootResourceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetRootResourceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LogRequest)(nil), "pulumirpc.LogRequest")
	proto.RegisterType((*GetRootResourceRequest)(nil), "pulumirpc.GetRootResourceRequest")
	proto.RegisterType((*GetRootResourceResponse)(nil), "pulumirpc.GetRootResourceResponse")
	proto.RegisterType((*SetRootResourceRequest)(nil), "pulumirpc.SetRootResourceRequest")
	proto.RegisterType((*SetRootResourceResponse)(nil), "pulumirpc.SetRootResourceResponse")
	proto.RegisterEnum("pulumirpc.LogSeverity", LogSeverity_name, LogSeverity_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Engine service

type EngineClient interface {
	// Log logs a global message in the engine, including errors and warnings.
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetRootResource(ctx context.Context, in *GetRootResourceRequest, opts ...grpc.CallOption) (*GetRootResourceResponse, error)
	SetRootResource(ctx context.Context, in *SetRootResourceRequest, opts ...grpc.CallOption) (*SetRootResourceResponse, error)
}

type engineClient struct {
	cc *grpc.ClientConn
}

func NewEngineClient(cc *grpc.ClientConn) EngineClient {
	return &engineClient{cc}
}

func (c *engineClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/pulumirpc.Engine/Log", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) GetRootResource(ctx context.Context, in *GetRootResourceRequest, opts ...grpc.CallOption) (*GetRootResourceResponse, error) {
	out := new(GetRootResourceResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.Engine/GetRootResource", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) SetRootResource(ctx context.Context, in *SetRootResourceRequest, opts ...grpc.CallOption) (*SetRootResourceResponse, error) {
	out := new(SetRootResourceResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.Engine/SetRootResource", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Engine service

type EngineServer interface {
	// Log logs a global message in the engine, including errors and warnings.
	Log(context.Context, *LogRequest) (*empty.Empty, error)
	GetRootResource(context.Context, *GetRootResourceRequest) (*GetRootResourceResponse, error)
	SetRootResource(context.Context, *SetRootResourceRequest) (*SetRootResourceResponse, error)
}

func RegisterEngineServer(s *grpc.Server, srv EngineServer) {
	s.RegisterService(&_Engine_serviceDesc, srv)
}

func _Engine_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.Engine/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).Log(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_GetRootResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRootResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).GetRootResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.Engine/GetRootResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).GetRootResource(ctx, req.(*GetRootResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_SetRootResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRootResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).SetRootResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.Engine/SetRootResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).SetRootResource(ctx, req.(*SetRootResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Engine_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pulumirpc.Engine",
	HandlerType: (*EngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Log",
			Handler:    _Engine_Log_Handler,
		},
		{
			MethodName: "GetRootResource",
			Handler:    _Engine_GetRootResource_Handler,
		},
		{
			MethodName: "SetRootResource",
			Handler:    _Engine_SetRootResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "engine.proto",
}

func init() { proto.RegisterFile("engine.proto", fileDescriptor_engine_f09f44fe55773ec3) }

var fileDescriptor_engine_f09f44fe55773ec3 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x9b, 0xa6, 0xff, 0x32, 0x15, 0x0d, 0x0b, 0xa6, 0x31, 0xf6, 0xa0, 0x39, 0x49, 0x85,
	0x14, 0x2a, 0x78, 0xf0, 0xa6, 0x18, 0x4b, 0x41, 0x5a, 0xd8, 0x20, 0x82, 0xb7, 0xb6, 0x8e, 0xb1,
	0xd0, 0x64, 0xe3, 0xee, 0x46, 0xe8, 0x17, 0xf2, 0x33, 0x7a, 0x34, 0x4d, 0xda, 0x58, 0x6b, 0xac,
	0xb7, 0xcc, 0xbc, 0x97, 0xdf, 0xce, 0xcc, 0x83, 0x3d, 0x0c, 0xfd, 0x59, 0x88, 0x4e, 0xc4, 0x99,
	0x64, 0x44, 0x8b, 0xe2, 0x79, 0x1c, 0xcc, 0x78, 0x34, 0xb5, 0x8e, 0x7d, 0xc6, 0xfc, 0x39, 0x76,
	0x53, 0x61, 0x12, 0xbf, 0x74, 0x31, 0x88, 0xe4, 0x22, 0xf3, 0x59, 0xed, 0x6d, 0x51, 0x48, 0x1e,
	0x4f, 0x65, 0xa6, 0xda, 0x1f, 0x0a, 0xc0, 0x3d, 0xf3, 0x29, 0xbe, 0xc5, 0x28, 0x24, 0xe9, 0x41,
	0x43, 0xe0, 0x3b, 0xf2, 0x99, 0x5c, 0x98, 0xca, 0x89, 0x72, 0xb6, 0xdf, 0x33, 0x9c, 0xfc, 0x1d,
	0x27, 0x31, 0x7a, 0x2b, 0x95, 0xe6, 0x3e, 0x62, 0x42, 0x3d, 0x40, 0x21, 0xc6, 0x3e, 0x9a, 0xe5,
	0xe4, 0x17, 0x8d, 0xae, 0x4b, 0xa2, 0x83, 0x1a, 0xf3, 0xd0, 0x54, 0xd3, 0xee, 0xf2, 0x93, 0x58,
	0x09, 0x5f, 0x72, 0x1c, 0x07, 0x83, 0x67, 0xb3, 0x92, 0xb4, 0xab, 0x34, 0xaf, 0x49, 0x1b, 0x34,
	0x8c, 0x5e, 0x31, 0x40, 0x3e, 0x9e, 0x9b, 0xd5, 0x44, 0x6c, 0xd0, 0xef, 0x86, 0x6d, 0x82, 0xd1,
	0x47, 0x49, 0x19, 0x93, 0x14, 0x05, 0x8b, 0xf9, 0x14, 0x57, 0x33, 0xdb, 0xe7, 0xd0, 0xfa, 0xa5,
	0x88, 0x88, 0x85, 0x22, 0x1f, 0x40, 0xc9, 0x07, 0xb0, 0x3b, 0x60, 0x78, 0x85, 0x98, 0x02, 0xef,
	0x11, 0xb4, 0xbc, 0x62, 0x70, 0xe7, 0x0a, 0x9a, 0x1b, 0xc7, 0x20, 0x1a, 0x54, 0x6f, 0xdd, 0x9b,
	0x87, 0xbe, 0x5e, 0x22, 0x0d, 0xa8, 0x0c, 0x86, 0x77, 0x23, 0x5d, 0x21, 0x4d, 0xa8, 0x3f, 0x5e,
	0xd3, 0xe1, 0x60, 0xd8, 0xd7, 0xcb, 0x4b, 0x87, 0x4b, 0xe9, 0x88, 0xea, 0x6a, 0xef, 0x53, 0x81,
	0x9a, 0x9b, 0x26, 0x49, 0x2e, 0x41, 0x4d, 0x30, 0xe4, 0xf0, 0xe7, 0x8d, 0x57, 0x13, 0x59, 0x86,
	0x93, 0x45, 0xe7, 0xac, 0xa3, 0x73, 0xdc, 0x65, 0xae, 0x76, 0x89, 0x3c, 0xc1, 0xc1, 0xd6, 0xca,
	0xe4, 0x74, 0x83, 0x51, 0x7c, 0x28, 0xcb, 0xde, 0x65, 0xc9, 0x16, 0xcb, 0xd8, 0xde, 0x0e, 0xb6,
	0xf7, 0x3f, 0xdb, 0xfb, 0x8b, 0x3d, 0xa9, 0xa5, 0x9b, 0x5c, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x00, 0xdb, 0x4a, 0xa9, 0xca, 0x02, 0x00, 0x00,
}
