// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: opentelemetry/proto/collector/logs/v1/logs_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/influxdata/influxdb-observability/otlp/logs/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ExportLogsServiceRequest struct {
	// An array of ResourceLogs.
	// For data coming from a single resource this array will typically contain one
	// element. Intermediary nodes (such as OpenTelemetry Collector) that receive
	// data from multiple origins typically batch the data before forwarding further and
	// in that case this array will contain multiple elements.
	ResourceLogs         []*v1.ResourceLogs `protobuf:"bytes,1,rep,name=resource_logs,json=resourceLogs,proto3" json:"resource_logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ExportLogsServiceRequest) Reset()         { *m = ExportLogsServiceRequest{} }
func (m *ExportLogsServiceRequest) String() string { return proto.CompactTextString(m) }
func (*ExportLogsServiceRequest) ProtoMessage()    {}
func (*ExportLogsServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e3bf87aaa43acd4, []int{0}
}
func (m *ExportLogsServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportLogsServiceRequest.Unmarshal(m, b)
}
func (m *ExportLogsServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportLogsServiceRequest.Marshal(b, m, deterministic)
}
func (m *ExportLogsServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportLogsServiceRequest.Merge(m, src)
}
func (m *ExportLogsServiceRequest) XXX_Size() int {
	return xxx_messageInfo_ExportLogsServiceRequest.Size(m)
}
func (m *ExportLogsServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportLogsServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportLogsServiceRequest proto.InternalMessageInfo

func (m *ExportLogsServiceRequest) GetResourceLogs() []*v1.ResourceLogs {
	if m != nil {
		return m.ResourceLogs
	}
	return nil
}

type ExportLogsServiceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportLogsServiceResponse) Reset()         { *m = ExportLogsServiceResponse{} }
func (m *ExportLogsServiceResponse) String() string { return proto.CompactTextString(m) }
func (*ExportLogsServiceResponse) ProtoMessage()    {}
func (*ExportLogsServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e3bf87aaa43acd4, []int{1}
}
func (m *ExportLogsServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportLogsServiceResponse.Unmarshal(m, b)
}
func (m *ExportLogsServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportLogsServiceResponse.Marshal(b, m, deterministic)
}
func (m *ExportLogsServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportLogsServiceResponse.Merge(m, src)
}
func (m *ExportLogsServiceResponse) XXX_Size() int {
	return xxx_messageInfo_ExportLogsServiceResponse.Size(m)
}
func (m *ExportLogsServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportLogsServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportLogsServiceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExportLogsServiceRequest)(nil), "opentelemetry.proto.collector.logs.v1.ExportLogsServiceRequest")
	proto.RegisterType((*ExportLogsServiceResponse)(nil), "opentelemetry.proto.collector.logs.v1.ExportLogsServiceResponse")
}

func init() {
	proto.RegisterFile("opentelemetry/proto/collector/logs/v1/logs_service.proto", fileDescriptor_8e3bf87aaa43acd4)
}

var fileDescriptor_8e3bf87aaa43acd4 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x2d, 0xc2, 0x0e, 0x99, 0x82, 0xf4, 0x34, 0xe7, 0x45, 0x0a, 0xca, 0x3c, 0x98, 0xb0,
	0x79, 0xf1, 0xa6, 0x4c, 0xbc, 0x89, 0x8c, 0x7a, 0xf3, 0x32, 0xda, 0xf8, 0xac, 0x91, 0xac, 0x2f,
	0x26, 0xaf, 0x65, 0x3b, 0x7b, 0xf6, 0x4f, 0xf0, 0x7f, 0x95, 0x34, 0x53, 0x2a, 0x56, 0x18, 0x3b,
	0x25, 0x79, 0x7c, 0xbf, 0xef, 0x7b, 0x8f, 0x3c, 0x76, 0x89, 0x06, 0x4a, 0x02, 0x0d, 0x0b, 0x20,
	0xbb, 0x12, 0xc6, 0x22, 0xa1, 0x90, 0xa8, 0x35, 0x48, 0x42, 0x2b, 0x34, 0x16, 0x4e, 0xd4, 0xe3,
	0xe6, 0x9c, 0x3b, 0xb0, 0xb5, 0x92, 0xc0, 0x1b, 0x51, 0x7c, 0xf2, 0x8b, 0x0c, 0x45, 0xfe, 0x43,
	0x72, 0x4f, 0xf0, 0x7a, 0x3c, 0x3c, 0xed, 0x0a, 0x68, 0xdb, 0x06, 0x32, 0x79, 0x65, 0x83, 0xdb,
	0xa5, 0x41, 0x4b, 0x77, 0x58, 0xb8, 0x87, 0x90, 0x94, 0xc2, 0x5b, 0x05, 0x8e, 0xe2, 0x7b, 0xb6,
	0x6f, 0xc1, 0x61, 0x65, 0x25, 0xcc, 0x3d, 0x32, 0x88, 0x8e, 0x77, 0x47, 0xfd, 0xc9, 0x19, 0xef,
	0x6a, 0x61, 0x1d, 0xcc, 0xd3, 0x35, 0xe1, 0xfd, 0xd2, 0x3d, 0xdb, 0x7a, 0x25, 0x47, 0xec, 0xb0,
	0x23, 0xcb, 0x19, 0x2c, 0x1d, 0x4c, 0x3e, 0x23, 0xd6, 0x6f, 0xd5, 0xe3, 0x8f, 0x88, 0xf5, 0x82,
	0x3a, 0xbe, 0xe2, 0x1b, 0xcd, 0xcc, 0xff, 0x1b, 0x64, 0x78, 0xbd, 0xbd, 0x41, 0xe8, 0x2e, 0xd9,
	0x99, 0xbe, 0x47, 0x6c, 0xa4, 0x70, 0x33, 0xa3, 0xe9, 0x41, 0xcb, 0x63, 0xe6, 0x35, 0xb3, 0xe8,
	0xf1, 0xa6, 0x50, 0xf4, 0x52, 0xe5, 0x5c, 0xe2, 0x42, 0xa8, 0xf2, 0x59, 0x57, 0xcb, 0xa7, 0x8c,
	0xb2, 0xef, 0x6b, 0x7e, 0x8e, 0xb9, 0xff, 0xe8, 0x2c, 0x57, 0x5a, 0xd1, 0x4a, 0x20, 0x69, 0xf3,
	0x77, 0x21, 0xf2, 0x5e, 0x93, 0x78, 0xf1, 0x15, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x5f, 0x54, 0x67,
	0x40, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogsServiceClient is the client API for LogsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogsServiceClient interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(ctx context.Context, in *ExportLogsServiceRequest, opts ...grpc.CallOption) (*ExportLogsServiceResponse, error)
}

type logsServiceClient struct {
	cc *grpc.ClientConn
}

func NewLogsServiceClient(cc *grpc.ClientConn) LogsServiceClient {
	return &logsServiceClient{cc}
}

func (c *logsServiceClient) Export(ctx context.Context, in *ExportLogsServiceRequest, opts ...grpc.CallOption) (*ExportLogsServiceResponse, error) {
	out := new(ExportLogsServiceResponse)
	err := c.cc.Invoke(ctx, "/opentelemetry.proto.collector.logs.v1.LogsService/Export", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogsServiceServer is the server API for LogsService service.
type LogsServiceServer interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(context.Context, *ExportLogsServiceRequest) (*ExportLogsServiceResponse, error)
}

// UnimplementedLogsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLogsServiceServer struct {
}

func (*UnimplementedLogsServiceServer) Export(ctx context.Context, req *ExportLogsServiceRequest) (*ExportLogsServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Export not implemented")
}

func RegisterLogsServiceServer(s *grpc.Server, srv LogsServiceServer) {
	s.RegisterService(&_LogsService_serviceDesc, srv)
}

func _LogsService_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportLogsServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogsServiceServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opentelemetry.proto.collector.logs.v1.LogsService/Export",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogsServiceServer).Export(ctx, req.(*ExportLogsServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "opentelemetry.proto.collector.logs.v1.LogsService",
	HandlerType: (*LogsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Export",
			Handler:    _LogsService_Export_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "opentelemetry/proto/collector/logs/v1/logs_service.proto",
}