// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opentelemetry/proto/collector/metrics/v1/metrics_service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1 "github.com/open-telemetry/opentelemetry-proto/gen/go/metrics/v1"
	v11 "github.com/open-telemetry/opentelemetry-proto/gen/go/resource/v1"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ExportMetricsServiceRequest struct {
	// Metric data. An array of ResourceMetrics.
	// For data coming from a single resource this array will typically contain one
	// element. Intermediary nodes (such as OpenTelemetry Collector) that receive
	// data from multiple origins typically batch the data before forwarding further and
	// in that case this array will contain multiple elements.
	ResourceMetrics      []*ResourceMetrics `protobuf:"bytes,1,rep,name=resource_metrics,json=resourceMetrics,proto3" json:"resource_metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ExportMetricsServiceRequest) Reset()         { *m = ExportMetricsServiceRequest{} }
func (m *ExportMetricsServiceRequest) String() string { return proto.CompactTextString(m) }
func (*ExportMetricsServiceRequest) ProtoMessage()    {}
func (*ExportMetricsServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_75fb6015e6e64798, []int{0}
}

func (m *ExportMetricsServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportMetricsServiceRequest.Unmarshal(m, b)
}
func (m *ExportMetricsServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportMetricsServiceRequest.Marshal(b, m, deterministic)
}
func (m *ExportMetricsServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportMetricsServiceRequest.Merge(m, src)
}
func (m *ExportMetricsServiceRequest) XXX_Size() int {
	return xxx_messageInfo_ExportMetricsServiceRequest.Size(m)
}
func (m *ExportMetricsServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportMetricsServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportMetricsServiceRequest proto.InternalMessageInfo

func (m *ExportMetricsServiceRequest) GetResourceMetrics() []*ResourceMetrics {
	if m != nil {
		return m.ResourceMetrics
	}
	return nil
}

// A collection of metrics from a Resource.
type ResourceMetrics struct {
	// A list of metrics that originate from a resource.
	Metrics []*v1.Metric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	// The resource for the metrics in this message that do not have an explicit
	// Metric.resource field set. If neither this field nor Metric.resource are set then no
	// resource info is known.
	Resource             *v11.Resource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ResourceMetrics) Reset()         { *m = ResourceMetrics{} }
func (m *ResourceMetrics) String() string { return proto.CompactTextString(m) }
func (*ResourceMetrics) ProtoMessage()    {}
func (*ResourceMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_75fb6015e6e64798, []int{1}
}

func (m *ResourceMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceMetrics.Unmarshal(m, b)
}
func (m *ResourceMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceMetrics.Marshal(b, m, deterministic)
}
func (m *ResourceMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceMetrics.Merge(m, src)
}
func (m *ResourceMetrics) XXX_Size() int {
	return xxx_messageInfo_ResourceMetrics.Size(m)
}
func (m *ResourceMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceMetrics proto.InternalMessageInfo

func (m *ResourceMetrics) GetMetrics() []*v1.Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *ResourceMetrics) GetResource() *v11.Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

type ExportMetricsServiceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportMetricsServiceResponse) Reset()         { *m = ExportMetricsServiceResponse{} }
func (m *ExportMetricsServiceResponse) String() string { return proto.CompactTextString(m) }
func (*ExportMetricsServiceResponse) ProtoMessage()    {}
func (*ExportMetricsServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_75fb6015e6e64798, []int{2}
}

func (m *ExportMetricsServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportMetricsServiceResponse.Unmarshal(m, b)
}
func (m *ExportMetricsServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportMetricsServiceResponse.Marshal(b, m, deterministic)
}
func (m *ExportMetricsServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportMetricsServiceResponse.Merge(m, src)
}
func (m *ExportMetricsServiceResponse) XXX_Size() int {
	return xxx_messageInfo_ExportMetricsServiceResponse.Size(m)
}
func (m *ExportMetricsServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportMetricsServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportMetricsServiceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExportMetricsServiceRequest)(nil), "opentelemetry.proto.collector.metrics.v1.ExportMetricsServiceRequest")
	proto.RegisterType((*ResourceMetrics)(nil), "opentelemetry.proto.collector.metrics.v1.ResourceMetrics")
	proto.RegisterType((*ExportMetricsServiceResponse)(nil), "opentelemetry.proto.collector.metrics.v1.ExportMetricsServiceResponse")
}

func init() {
	proto.RegisterFile("opentelemetry/proto/collector/metrics/v1/metrics_service.proto", fileDescriptor_75fb6015e6e64798)
}

var fileDescriptor_75fb6015e6e64798 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xcb, 0x2f, 0x48, 0xcd,
	0x2b, 0x49, 0xcd, 0x49, 0xcd, 0x4d, 0x2d, 0x29, 0xaa, 0xd4, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7,
	0x4f, 0xce, 0xcf, 0xc9, 0x49, 0x4d, 0x2e, 0xc9, 0x2f, 0xd2, 0x07, 0x89, 0x66, 0x26, 0x17, 0xeb,
	0x97, 0x19, 0xc2, 0x98, 0xf1, 0xc5, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x7a, 0x60, 0xa5, 0x42,
	0x1a, 0x28, 0xfa, 0x21, 0x82, 0x7a, 0x70, 0xfd, 0x7a, 0x50, 0x4d, 0x7a, 0x65, 0x86, 0x52, 0x3a,
	0xd8, 0x6c, 0xc2, 0x34, 0x1f, 0x62, 0x84, 0x94, 0x1e, 0x36, 0xd5, 0x45, 0xa9, 0xc5, 0xf9, 0xa5,
	0x45, 0xc9, 0xa9, 0x20, 0xe5, 0x30, 0x36, 0x44, 0xbd, 0x52, 0x33, 0x23, 0x97, 0xb4, 0x6b, 0x45,
	0x41, 0x7e, 0x51, 0x89, 0x2f, 0xc4, 0x9c, 0x60, 0x88, 0x33, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b,
	0x4b, 0x84, 0x52, 0xb8, 0x04, 0x60, 0x3a, 0xe2, 0xa1, 0x36, 0x49, 0x30, 0x2a, 0x30, 0x6b, 0x70,
	0x1b, 0x59, 0xea, 0x11, 0xeb, 0x05, 0xbd, 0x20, 0xa8, 0x09, 0x50, 0x2b, 0x82, 0xf8, 0x8b, 0x50,
	0x05, 0x94, 0x66, 0x31, 0x72, 0xf1, 0xa3, 0x29, 0x12, 0x72, 0xe0, 0x62, 0x47, 0xb5, 0x50, 0x0d,
	0xab, 0x85, 0x48, 0xd6, 0x40, 0x74, 0x06, 0xc1, 0xb4, 0x09, 0xb9, 0x72, 0x71, 0xc0, 0x2c, 0x92,
	0x60, 0x52, 0x60, 0xd4, 0xe0, 0x36, 0xd2, 0xc4, 0x6a, 0x04, 0x3c, 0x48, 0x90, 0x9c, 0x1a, 0x04,
	0xd7, 0xaa, 0x24, 0xc7, 0x25, 0x83, 0x3d, 0x84, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x8d, 0xd6,
	0x30, 0x72, 0xf1, 0xa1, 0x4a, 0x09, 0xcd, 0x64, 0xe4, 0x62, 0x83, 0xe8, 0x11, 0x72, 0x25, 0x3e,
	0x98, 0xf0, 0xc4, 0x83, 0x94, 0x1b, 0xa5, 0xc6, 0x40, 0x1c, 0xab, 0xc4, 0xe0, 0xd4, 0xcf, 0xc8,
	0xa5, 0x9d, 0x99, 0x4f, 0xb4, 0x71, 0x4e, 0xc2, 0xa8, 0x26, 0x05, 0x80, 0x54, 0x06, 0x30, 0x46,
	0x79, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x83, 0xcc, 0xd2, 0x45,
	0x24, 0x3a, 0x14, 0xa3, 0x75, 0x21, 0x49, 0x30, 0x3d, 0x35, 0x4f, 0x3f, 0x1d, 0x7b, 0x0e, 0x49,
	0x62, 0x03, 0x2b, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xdb, 0xf8, 0xb2, 0x54, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetricsServiceClient is the client API for MetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricsServiceClient interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(ctx context.Context, in *ExportMetricsServiceRequest, opts ...grpc.CallOption) (*ExportMetricsServiceResponse, error)
}

type metricsServiceClient struct {
	cc *grpc.ClientConn
}

func NewMetricsServiceClient(cc *grpc.ClientConn) MetricsServiceClient {
	return &metricsServiceClient{cc}
}

func (c *metricsServiceClient) Export(ctx context.Context, in *ExportMetricsServiceRequest, opts ...grpc.CallOption) (*ExportMetricsServiceResponse, error) {
	out := new(ExportMetricsServiceResponse)
	err := c.cc.Invoke(ctx, "/opentelemetry.proto.collector.metrics.v1.MetricsService/Export", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsServiceServer is the server API for MetricsService service.
type MetricsServiceServer interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(context.Context, *ExportMetricsServiceRequest) (*ExportMetricsServiceResponse, error)
}

// UnimplementedMetricsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMetricsServiceServer struct {
}

func (*UnimplementedMetricsServiceServer) Export(ctx context.Context, req *ExportMetricsServiceRequest) (*ExportMetricsServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Export not implemented")
}

func RegisterMetricsServiceServer(s *grpc.Server, srv MetricsServiceServer) {
	s.RegisterService(&_MetricsService_serviceDesc, srv)
}

func _MetricsService_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportMetricsServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opentelemetry.proto.collector.metrics.v1.MetricsService/Export",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).Export(ctx, req.(*ExportMetricsServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "opentelemetry.proto.collector.metrics.v1.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Export",
			Handler:    _MetricsService_Export_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "opentelemetry/proto/collector/metrics/v1/metrics_service.proto",
}
