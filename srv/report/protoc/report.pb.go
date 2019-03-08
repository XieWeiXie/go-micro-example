// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/wuxiaoxiaoshen/go-micro-example/srv/report/protoc/report.proto

package report

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protoc "github.com/wuxiaoxiaoshen/go-micro-example/srv/auth/protoc"
	grpc "google.golang.org/grpc"
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

type Request struct {
	ID                   int32             `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Auth                 []*protoc.Request `protobuf:"bytes,3,rep,name=auth,proto3" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9569cd25f2840ef3, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetAuth() []*protoc.Request {
	if m != nil {
		return m.Auth
	}
	return nil
}

type Response struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	School               string   `protobuf:"bytes,3,opt,name=school,proto3" json:"school,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9569cd25f2840ef3, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Response) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Response) GetSchool() string {
	if m != nil {
		return m.School
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "report.Request")
	proto.RegisterType((*Response)(nil), "report.Response")
}

func init() {
	proto.RegisterFile("github.com/wuxiaoxiaoshen/go-micro-example/srv/report/protoc/report.proto", fileDescriptor_9569cd25f2840ef3)
}

var fileDescriptor_9569cd25f2840ef3 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xcf, 0x4e, 0x84, 0x30,
	0x10, 0x87, 0x05, 0x56, 0x74, 0xc7, 0xbf, 0xe9, 0xc1, 0x90, 0x3d, 0x21, 0x27, 0x2e, 0x80, 0x59,
	0x5f, 0x01, 0x35, 0x9c, 0x34, 0x4d, 0x7c, 0x80, 0x2e, 0x4e, 0x16, 0x12, 0x60, 0x6a, 0x5b, 0x74,
	0x5f, 0xc7, 0x37, 0x35, 0x94, 0x1a, 0xbd, 0x98, 0xec, 0x1e, 0x9a, 0xcc, 0x37, 0xcd, 0xef, 0x9b,
	0xcc, 0x40, 0xb5, 0x6d, 0x4d, 0x33, 0x6e, 0xf2, 0x9a, 0xfa, 0xe2, 0x73, 0xdc, 0xb5, 0x82, 0xa6,
	0xa7, 0x1b, 0x1c, 0x8a, 0x2d, 0x65, 0x7d, 0x5b, 0x2b, 0xca, 0x70, 0x27, 0x7a, 0xd9, 0x61, 0xa1,
	0xd5, 0x47, 0xa1, 0x50, 0x92, 0x32, 0x85, 0x54, 0x64, 0xa8, 0x76, 0x94, 0x5b, 0x62, 0xe1, 0x4c,
	0xab, 0x87, 0x03, 0x95, 0x62, 0x34, 0xcd, 0x8f, 0x70, 0xaa, 0x67, 0x5d, 0xf2, 0x02, 0x27, 0x1c,
	0xdf, 0x47, 0xd4, 0x86, 0x5d, 0x82, 0x5f, 0x95, 0x91, 0x17, 0x7b, 0xe9, 0x31, 0xf7, 0xab, 0x92,
	0x31, 0x58, 0x0c, 0xa2, 0xc7, 0xc8, 0x8f, 0xbd, 0x74, 0xc9, 0x6d, 0xcd, 0x6e, 0x61, 0x31, 0x85,
	0xa3, 0x20, 0x0e, 0xd2, 0xb3, 0xf5, 0x45, 0x6e, 0x4d, 0x4e, 0xc0, 0xed, 0x57, 0xf2, 0x08, 0xa7,
	0x1c, 0xb5, 0xa4, 0x41, 0xe3, 0x1f, 0x65, 0xf0, 0xaf, 0xf2, 0x06, 0x42, 0x5d, 0x37, 0x44, 0x5d,
	0x14, 0xd8, 0xae, 0xa3, 0xf5, 0x97, 0x07, 0xe7, 0x4f, 0x68, 0x9e, 0x07, 0xe4, 0x76, 0x63, 0x96,
	0x41, 0x38, 0x33, 0xbb, 0xca, 0xdd, 0x49, 0xdc, 0xe4, 0xd5, 0xf5, 0x6f, 0x63, 0x9e, 0x9c, 0x1c,
	0xb1, 0x3b, 0x58, 0xbe, 0xca, 0x37, 0x61, 0xf0, 0x90, 0x44, 0x89, 0x1d, 0xee, 0x9f, 0xd8, 0x84,
	0xf6, 0x88, 0xf7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x97, 0xf0, 0x9f, 0x5b, 0xe0, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GetOneReportClient is the client API for GetOneReport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GetOneReportClient interface {
	GetOne(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	UpdateOne(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	DeleteOne(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type getOneReportClient struct {
	cc *grpc.ClientConn
}

func NewGetOneReportClient(cc *grpc.ClientConn) GetOneReportClient {
	return &getOneReportClient{cc}
}

func (c *getOneReportClient) GetOne(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/report.GetOneReport/GetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getOneReportClient) UpdateOne(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/report.GetOneReport/UpdateOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getOneReportClient) DeleteOne(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/report.GetOneReport/DeleteOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetOneReportServer is the server API for GetOneReport service.
type GetOneReportServer interface {
	GetOne(context.Context, *Request) (*Response, error)
	UpdateOne(context.Context, *Request) (*Response, error)
	DeleteOne(context.Context, *Request) (*Response, error)
}

func RegisterGetOneReportServer(s *grpc.Server, srv GetOneReportServer) {
	s.RegisterService(&_GetOneReport_serviceDesc, srv)
}

func _GetOneReport_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetOneReportServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/report.GetOneReport/GetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetOneReportServer).GetOne(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetOneReport_UpdateOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetOneReportServer).UpdateOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/report.GetOneReport/UpdateOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetOneReportServer).UpdateOne(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetOneReport_DeleteOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetOneReportServer).DeleteOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/report.GetOneReport/DeleteOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetOneReportServer).DeleteOne(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _GetOneReport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "report.GetOneReport",
	HandlerType: (*GetOneReportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOne",
			Handler:    _GetOneReport_GetOne_Handler,
		},
		{
			MethodName: "UpdateOne",
			Handler:    _GetOneReport_UpdateOne_Handler,
		},
		{
			MethodName: "DeleteOne",
			Handler:    _GetOneReport_DeleteOne_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/wuxiaoxiaoshen/go-micro-example/srv/report/protoc/report.proto",
}
