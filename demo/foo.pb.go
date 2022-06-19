// Code generated by protoc-gen-go. DO NOT EDIT.
// source: foo.proto

package main

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

type GreetReq struct {
	MyName               string   `protobuf:"bytes,1,opt,name=my_name,json=myName,proto3" json:"my_name,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetReq) Reset()         { *m = GreetReq{} }
func (m *GreetReq) String() string { return proto.CompactTextString(m) }
func (*GreetReq) ProtoMessage()    {}
func (*GreetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce1e2eec643ca48, []int{0}
}

func (m *GreetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetReq.Unmarshal(m, b)
}
func (m *GreetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetReq.Marshal(b, m, deterministic)
}
func (m *GreetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetReq.Merge(m, src)
}
func (m *GreetReq) XXX_Size() int {
	return xxx_messageInfo_GreetReq.Size(m)
}
func (m *GreetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetReq.DiscardUnknown(m)
}

var xxx_messageInfo_GreetReq proto.InternalMessageInfo

func (m *GreetReq) GetMyName() string {
	if m != nil {
		return m.MyName
	}
	return ""
}

func (m *GreetReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type GreetResp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetResp) Reset()         { *m = GreetResp{} }
func (m *GreetResp) String() string { return proto.CompactTextString(m) }
func (*GreetResp) ProtoMessage()    {}
func (*GreetResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce1e2eec643ca48, []int{1}
}

func (m *GreetResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetResp.Unmarshal(m, b)
}
func (m *GreetResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetResp.Marshal(b, m, deterministic)
}
func (m *GreetResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResp.Merge(m, src)
}
func (m *GreetResp) XXX_Size() int {
	return xxx_messageInfo_GreetResp.Size(m)
}
func (m *GreetResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResp.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResp proto.InternalMessageInfo

func (m *GreetResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*GreetReq)(nil), "pb.GreetReq")
	proto.RegisterType((*GreetResp)(nil), "pb.GreetResp")
}

func init() { proto.RegisterFile("foo.proto", fileDescriptor_7ce1e2eec643ca48) }

var fileDescriptor_7ce1e2eec643ca48 = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0xcb, 0xcf, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x32, 0xe5, 0xe2, 0x70, 0x2f, 0x4a,
	0x4d, 0x2d, 0x09, 0x4a, 0x2d, 0x14, 0x12, 0xe7, 0x62, 0xcf, 0xad, 0x8c, 0xcf, 0x4b, 0xcc, 0x4d,
	0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xcb, 0xad, 0xf4, 0x4b, 0xcc, 0x4d, 0x15, 0x12,
	0xe0, 0x62, 0xce, 0x2d, 0x4e, 0x97, 0x60, 0x02, 0x0b, 0x82, 0x98, 0x4a, 0xb2, 0x5c, 0x9c, 0x50,
	0x6d, 0xc5, 0x05, 0x30, 0x69, 0x46, 0xb8, 0xb4, 0x91, 0x36, 0x17, 0xb3, 0x5b, 0x7e, 0xbe, 0x90,
	0x0a, 0x17, 0x2b, 0x58, 0x95, 0x10, 0x8f, 0x5e, 0x41, 0x92, 0x1e, 0xcc, 0x1e, 0x29, 0x5e, 0x24,
	0x5e, 0x71, 0x41, 0x12, 0x1b, 0xd8, 0x35, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x8d,
	0x70, 0xf6, 0x9a, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FooClient is the client API for Foo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FooClient interface {
	Greet(ctx context.Context, in *GreetReq, opts ...grpc.CallOption) (*GreetResp, error)
}

type fooClient struct {
	cc *grpc.ClientConn
}

func NewFooClient(cc *grpc.ClientConn) FooClient {
	return &fooClient{cc}
}

func (c *fooClient) Greet(ctx context.Context, in *GreetReq, opts ...grpc.CallOption) (*GreetResp, error) {
	out := new(GreetResp)
	err := c.cc.Invoke(ctx, "/pb.Foo/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FooServer is the server API for Foo service.
type FooServer interface {
	Greet(context.Context, *GreetReq) (*GreetResp, error)
}

// UnimplementedFooServer can be embedded to have forward compatible implementations.
type UnimplementedFooServer struct {
}

func (*UnimplementedFooServer) Greet(ctx context.Context, req *GreetReq) (*GreetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}

func RegisterFooServer(s *grpc.Server, srv FooServer) {
	s.RegisterService(&_Foo_serviceDesc, srv)
}

func _Foo_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FooServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Foo/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FooServer).Greet(ctx, req.(*GreetReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Foo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Foo",
	HandlerType: (*FooServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _Foo_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "foo.proto",
}
