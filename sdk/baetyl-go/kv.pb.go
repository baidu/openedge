// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kv.proto

package baetyl

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

// KV kv message
type KVMessage struct {
	// key is the key, in bytes, to put into the key-value store.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// value is the value, in bytes, to associate with the key in the key-value store.
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVMessage) Reset()         { *m = KVMessage{} }
func (m *KVMessage) String() string { return proto.CompactTextString(m) }
func (*KVMessage) ProtoMessage()    {}
func (*KVMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_2216fe83c9c12408, []int{0}
}

func (m *KVMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVMessage.Unmarshal(m, b)
}
func (m *KVMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVMessage.Marshal(b, m, deterministic)
}
func (m *KVMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVMessage.Merge(m, src)
}
func (m *KVMessage) XXX_Size() int {
	return xxx_messageInfo_KVMessage.Size(m)
}
func (m *KVMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_KVMessage.DiscardUnknown(m)
}

var xxx_messageInfo_KVMessage proto.InternalMessageInfo

func (m *KVMessage) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KVMessage) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type ArrayKVMessage struct {
	Kvs                  []*KVMessage `protobuf:"bytes,1,rep,name=kvs,proto3" json:"kvs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ArrayKVMessage) Reset()         { *m = ArrayKVMessage{} }
func (m *ArrayKVMessage) String() string { return proto.CompactTextString(m) }
func (*ArrayKVMessage) ProtoMessage()    {}
func (*ArrayKVMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_2216fe83c9c12408, []int{1}
}

func (m *ArrayKVMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrayKVMessage.Unmarshal(m, b)
}
func (m *ArrayKVMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrayKVMessage.Marshal(b, m, deterministic)
}
func (m *ArrayKVMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrayKVMessage.Merge(m, src)
}
func (m *ArrayKVMessage) XXX_Size() int {
	return xxx_messageInfo_ArrayKVMessage.Size(m)
}
func (m *ArrayKVMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrayKVMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ArrayKVMessage proto.InternalMessageInfo

func (m *ArrayKVMessage) GetKvs() []*KVMessage {
	if m != nil {
		return m.Kvs
	}
	return nil
}

func init() {
	proto.RegisterType((*KVMessage)(nil), "baetyl.KVMessage")
	proto.RegisterType((*ArrayKVMessage)(nil), "baetyl.ArrayKVMessage")
}

func init() { proto.RegisterFile("kv.proto", fileDescriptor_2216fe83c9c12408) }

var fileDescriptor_2216fe83c9c12408 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0x2e, 0xd3, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0x4a, 0x4c, 0x2d, 0xa9, 0xcc, 0x51, 0x32, 0xe6, 0xe2,
	0xf4, 0x0e, 0xf3, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad,
	0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0x31, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73,
	0x4a, 0x53, 0x25, 0x98, 0xc0, 0x62, 0x10, 0x8e, 0x92, 0x29, 0x17, 0x9f, 0x63, 0x51, 0x51, 0x62,
	0x25, 0x42, 0xa7, 0x32, 0x17, 0x73, 0x76, 0x59, 0xb1, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0xb7, 0x91,
	0xa0, 0x1e, 0xc4, 0x70, 0x3d, 0xb8, 0x7c, 0x10, 0x48, 0xd6, 0xe8, 0x1c, 0x23, 0x17, 0x93, 0x77,
	0x98, 0x90, 0x3e, 0x17, 0xab, 0x7b, 0x6a, 0x89, 0x77, 0x98, 0x10, 0xa6, 0x3a, 0x29, 0x4c, 0x21,
	0x25, 0x06, 0x21, 0x53, 0x2e, 0x36, 0x9f, 0xcc, 0x62, 0x1c, 0x3a, 0xc4, 0x60, 0x42, 0xa8, 0x2e,
	0x52, 0x62, 0x00, 0xd9, 0x13, 0x50, 0x4a, 0x8a, 0x3d, 0xfa, 0x5c, 0xac, 0x2e, 0xa9, 0x39, 0xc4,
	0x6b, 0x48, 0x62, 0x03, 0x87, 0xa5, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x46, 0x26, 0x41,
	0x57, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KVClient is the client API for KV service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KVClient interface {
	GetKV(ctx context.Context, in *KVMessage, opts ...grpc.CallOption) (*KVMessage, error)
	ListKV(ctx context.Context, in *KVMessage, opts ...grpc.CallOption) (*ArrayKVMessage, error)
	PutKV(ctx context.Context, in *KVMessage, opts ...grpc.CallOption) (*KVMessage, error)
	DelKV(ctx context.Context, in *KVMessage, opts ...grpc.CallOption) (*KVMessage, error)
}

type kVClient struct {
	cc *grpc.ClientConn
}

func NewKVClient(cc *grpc.ClientConn) KVClient {
	return &kVClient{cc}
}

func (c *kVClient) GetKV(ctx context.Context, in *KVMessage, opts ...grpc.CallOption) (*KVMessage, error) {
	out := new(KVMessage)
	err := c.cc.Invoke(ctx, "/baetyl.KV/GetKV", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) ListKV(ctx context.Context, in *KVMessage, opts ...grpc.CallOption) (*ArrayKVMessage, error) {
	out := new(ArrayKVMessage)
	err := c.cc.Invoke(ctx, "/baetyl.KV/ListKV", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) PutKV(ctx context.Context, in *KVMessage, opts ...grpc.CallOption) (*KVMessage, error) {
	out := new(KVMessage)
	err := c.cc.Invoke(ctx, "/baetyl.KV/PutKV", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) DelKV(ctx context.Context, in *KVMessage, opts ...grpc.CallOption) (*KVMessage, error) {
	out := new(KVMessage)
	err := c.cc.Invoke(ctx, "/baetyl.KV/DelKV", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KVServer is the server API for KV service.
type KVServer interface {
	GetKV(context.Context, *KVMessage) (*KVMessage, error)
	ListKV(context.Context, *KVMessage) (*ArrayKVMessage, error)
	PutKV(context.Context, *KVMessage) (*KVMessage, error)
	DelKV(context.Context, *KVMessage) (*KVMessage, error)
}

// UnimplementedKVServer can be embedded to have forward compatible implementations.
type UnimplementedKVServer struct {
}

func (*UnimplementedKVServer) GetKV(ctx context.Context, req *KVMessage) (*KVMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKV not implemented")
}
func (*UnimplementedKVServer) ListKV(ctx context.Context, req *KVMessage) (*ArrayKVMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKV not implemented")
}
func (*UnimplementedKVServer) PutKV(ctx context.Context, req *KVMessage) (*KVMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutKV not implemented")
}
func (*UnimplementedKVServer) DelKV(ctx context.Context, req *KVMessage) (*KVMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelKV not implemented")
}

func RegisterKVServer(s *grpc.Server, srv KVServer) {
	s.RegisterService(&_KV_serviceDesc, srv)
}

func _KV_GetKV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KVMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).GetKV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/baetyl.KV/GetKV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).GetKV(ctx, req.(*KVMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_ListKV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KVMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).ListKV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/baetyl.KV/ListKV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).ListKV(ctx, req.(*KVMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_PutKV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KVMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).PutKV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/baetyl.KV/PutKV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).PutKV(ctx, req.(*KVMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_DelKV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KVMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).DelKV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/baetyl.KV/DelKV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).DelKV(ctx, req.(*KVMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _KV_serviceDesc = grpc.ServiceDesc{
	ServiceName: "baetyl.KV",
	HandlerType: (*KVServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKV",
			Handler:    _KV_GetKV_Handler,
		},
		{
			MethodName: "ListKV",
			Handler:    _KV_ListKV_Handler,
		},
		{
			MethodName: "PutKV",
			Handler:    _KV_PutKV_Handler,
		},
		{
			MethodName: "DelKV",
			Handler:    _KV_DelKV_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kv.proto",
}
