// Code generated by protoc-gen-go. DO NOT EDIT.
// source: credentialmonitor.proto

package admin

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CertificateStatus int32

const (
	CertificateStatus_Single  CertificateStatus = 0
	CertificateStatus_Overlap CertificateStatus = 1
)

var CertificateStatus_name = map[int32]string{
	0: "Single",
	1: "Overlap",
}

var CertificateStatus_value = map[string]int32{
	"Single":  0,
	"Overlap": 1,
}

func (x CertificateStatus) String() string {
	return proto.EnumName(CertificateStatus_name, int32(x))
}

func (CertificateStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0e30998b3c2b899e, []int{0}
}

type CredentialMonitorRequest struct {
	CredentialMonitor    *CredentialMonitor `protobuf:"bytes,1,opt,name=CredentialMonitor,proto3" json:"CredentialMonitor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CredentialMonitorRequest) Reset()         { *m = CredentialMonitorRequest{} }
func (m *CredentialMonitorRequest) String() string { return proto.CompactTextString(m) }
func (*CredentialMonitorRequest) ProtoMessage()    {}
func (*CredentialMonitorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e30998b3c2b899e, []int{0}
}

func (m *CredentialMonitorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredentialMonitorRequest.Unmarshal(m, b)
}
func (m *CredentialMonitorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredentialMonitorRequest.Marshal(b, m, deterministic)
}
func (m *CredentialMonitorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialMonitorRequest.Merge(m, src)
}
func (m *CredentialMonitorRequest) XXX_Size() int {
	return xxx_messageInfo_CredentialMonitorRequest.Size(m)
}
func (m *CredentialMonitorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialMonitorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialMonitorRequest proto.InternalMessageInfo

func (m *CredentialMonitorRequest) GetCredentialMonitor() *CredentialMonitor {
	if m != nil {
		return m.CredentialMonitor
	}
	return nil
}

type CredentialMonitorResponse struct {
	CredentialMonitor    *CredentialMonitor `protobuf:"bytes,1,opt,name=CredentialMonitor,proto3" json:"CredentialMonitor,omitempty"`
	Error                string             `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CredentialMonitorResponse) Reset()         { *m = CredentialMonitorResponse{} }
func (m *CredentialMonitorResponse) String() string { return proto.CompactTextString(m) }
func (*CredentialMonitorResponse) ProtoMessage()    {}
func (*CredentialMonitorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e30998b3c2b899e, []int{1}
}

func (m *CredentialMonitorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredentialMonitorResponse.Unmarshal(m, b)
}
func (m *CredentialMonitorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredentialMonitorResponse.Marshal(b, m, deterministic)
}
func (m *CredentialMonitorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialMonitorResponse.Merge(m, src)
}
func (m *CredentialMonitorResponse) XXX_Size() int {
	return xxx_messageInfo_CredentialMonitorResponse.Size(m)
}
func (m *CredentialMonitorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialMonitorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialMonitorResponse proto.InternalMessageInfo

func (m *CredentialMonitorResponse) GetCredentialMonitor() *CredentialMonitor {
	if m != nil {
		return m.CredentialMonitor
	}
	return nil
}

func (m *CredentialMonitorResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type CredentialMonitor struct {
	Certificate          string            `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
	Status               CertificateStatus `protobuf:"varint,2,opt,name=status,proto3,enum=moc.cloudagent.admin.CertificateStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CredentialMonitor) Reset()         { *m = CredentialMonitor{} }
func (m *CredentialMonitor) String() string { return proto.CompactTextString(m) }
func (*CredentialMonitor) ProtoMessage()    {}
func (*CredentialMonitor) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e30998b3c2b899e, []int{2}
}

func (m *CredentialMonitor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredentialMonitor.Unmarshal(m, b)
}
func (m *CredentialMonitor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredentialMonitor.Marshal(b, m, deterministic)
}
func (m *CredentialMonitor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialMonitor.Merge(m, src)
}
func (m *CredentialMonitor) XXX_Size() int {
	return xxx_messageInfo_CredentialMonitor.Size(m)
}
func (m *CredentialMonitor) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialMonitor.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialMonitor proto.InternalMessageInfo

func (m *CredentialMonitor) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

func (m *CredentialMonitor) GetStatus() CertificateStatus {
	if m != nil {
		return m.Status
	}
	return CertificateStatus_Single
}

func init() {
	proto.RegisterEnum("moc.cloudagent.admin.CertificateStatus", CertificateStatus_name, CertificateStatus_value)
	proto.RegisterType((*CredentialMonitorRequest)(nil), "moc.cloudagent.admin.CredentialMonitorRequest")
	proto.RegisterType((*CredentialMonitorResponse)(nil), "moc.cloudagent.admin.CredentialMonitorResponse")
	proto.RegisterType((*CredentialMonitor)(nil), "moc.cloudagent.admin.CredentialMonitor")
}

func init() { proto.RegisterFile("credentialmonitor.proto", fileDescriptor_0e30998b3c2b899e) }

var fileDescriptor_0e30998b3c2b899e = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xc1, 0x4a, 0x33, 0x31,
	0x10, 0xc7, 0x9b, 0xef, 0xc3, 0x4a, 0xa7, 0x20, 0x6d, 0x28, 0x5a, 0x3d, 0x95, 0xbd, 0x58, 0x44,
	0x13, 0xa8, 0x0f, 0x20, 0x2a, 0xe2, 0x49, 0x84, 0x2d, 0x5e, 0xbc, 0xa5, 0x69, 0xba, 0x0d, 0x6c,
	0x32, 0xdb, 0x64, 0xb6, 0x0f, 0xe0, 0xc9, 0xc7, 0x16, 0x77, 0xc5, 0x0a, 0xbb, 0x42, 0x2f, 0x1e,
	0x93, 0x99, 0xff, 0xef, 0x37, 0x64, 0x02, 0x27, 0x3a, 0x98, 0xa5, 0xf1, 0x64, 0x55, 0xee, 0xd0,
	0x5b, 0xc2, 0x20, 0x8a, 0x80, 0x84, 0x7c, 0xe4, 0x50, 0x0b, 0x9d, 0x63, 0xb9, 0x54, 0x99, 0xf1,
	0x24, 0xd4, 0xd2, 0x59, 0x9f, 0x6c, 0x60, 0x7c, 0xff, 0x1d, 0x78, 0xaa, 0x03, 0xa9, 0xd9, 0x94,
	0x26, 0x12, 0x7f, 0x81, 0x61, 0xa3, 0x36, 0x66, 0x13, 0x36, 0xed, 0xcf, 0xce, 0x45, 0x1b, 0x4d,
	0x34, 0x51, 0x4d, 0x42, 0xf2, 0xce, 0xe0, 0xb4, 0xc5, 0x19, 0x0b, 0xf4, 0xd1, 0xfc, 0x91, 0x94,
	0x8f, 0xe0, 0xe0, 0x21, 0x04, 0x0c, 0xe3, 0x7f, 0x13, 0x36, 0xed, 0xa5, 0xf5, 0x21, 0xd9, 0xb6,
	0xc8, 0xf8, 0x04, 0xfa, 0xda, 0x04, 0xb2, 0x2b, 0xab, 0x15, 0x99, 0xca, 0xdd, 0x4b, 0x7f, 0x5e,
	0xf1, 0x1b, 0xe8, 0x46, 0x52, 0x54, 0xc6, 0x8a, 0x76, 0xf4, 0xeb, 0x60, 0xbb, 0xc8, 0xbc, 0x6a,
	0x4f, 0xbf, 0x62, 0x17, 0x97, 0x30, 0x6c, 0x14, 0x39, 0x40, 0x77, 0x6e, 0x7d, 0x96, 0x9b, 0x41,
	0x87, 0xf7, 0xe1, 0xf0, 0x79, 0x6b, 0x42, 0xae, 0x8a, 0x01, 0x9b, 0xbd, 0x31, 0x38, 0x6e, 0x8c,
	0x79, 0xfb, 0x29, 0xe2, 0x6b, 0xf8, 0xff, 0x68, 0x88, 0x8b, 0x7d, 0x5f, 0xa6, 0xde, 0xec, 0x99,
	0xdc, 0xbb, 0xbf, 0xde, 0x4a, 0xd2, 0xb9, 0x93, 0xaf, 0x57, 0x99, 0xa5, 0x75, 0xb9, 0x10, 0x1a,
	0x9d, 0x74, 0x56, 0x07, 0x8c, 0xb8, 0x22, 0xe9, 0x50, 0xcb, 0x50, 0x68, 0xb9, 0x83, 0xc9, 0x0a,
	0xb6, 0xe8, 0x56, 0xdf, 0xee, 0xfa, 0x23, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x5a, 0xc6, 0xde, 0x91,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CredentialMonitorAgentClient is the client API for CredentialMonitorAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CredentialMonitorAgentClient interface {
	Get(ctx context.Context, in *CredentialMonitorRequest, opts ...grpc.CallOption) (*CredentialMonitorResponse, error)
}

type credentialMonitorAgentClient struct {
	cc *grpc.ClientConn
}

func NewCredentialMonitorAgentClient(cc *grpc.ClientConn) CredentialMonitorAgentClient {
	return &credentialMonitorAgentClient{cc}
}

func (c *credentialMonitorAgentClient) Get(ctx context.Context, in *CredentialMonitorRequest, opts ...grpc.CallOption) (*CredentialMonitorResponse, error) {
	out := new(CredentialMonitorResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.admin.CredentialMonitorAgent/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CredentialMonitorAgentServer is the server API for CredentialMonitorAgent service.
type CredentialMonitorAgentServer interface {
	Get(context.Context, *CredentialMonitorRequest) (*CredentialMonitorResponse, error)
}

// UnimplementedCredentialMonitorAgentServer can be embedded to have forward compatible implementations.
type UnimplementedCredentialMonitorAgentServer struct {
}

func (*UnimplementedCredentialMonitorAgentServer) Get(ctx context.Context, req *CredentialMonitorRequest) (*CredentialMonitorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterCredentialMonitorAgentServer(s *grpc.Server, srv CredentialMonitorAgentServer) {
	s.RegisterService(&_CredentialMonitorAgent_serviceDesc, srv)
}

func _CredentialMonitorAgent_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialMonitorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialMonitorAgentServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.admin.CredentialMonitorAgent/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialMonitorAgentServer).Get(ctx, req.(*CredentialMonitorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CredentialMonitorAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.admin.CredentialMonitorAgent",
	HandlerType: (*CredentialMonitorAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CredentialMonitorAgent_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "credentialmonitor.proto",
}
