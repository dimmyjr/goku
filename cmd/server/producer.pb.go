package server

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

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

type Request struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Reply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Reply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "server.Request")
	proto.RegisterType((*Reply)(nil), "server.Reply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Producer service

type ProducerClient interface {
	Publish(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
}

type producerClient struct {
	cc *grpc.ClientConn
}

func NewProducerClient(cc *grpc.ClientConn) ProducerClient {
	return &producerClient{cc}
}

func (c *producerClient) Publish(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/server.Producer/publish", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Producer service

type ProducerServer interface {
	Publish(context.Context, *Request) (*Reply, error)
}

func RegisterProducerServer(s *grpc.Server, srv ProducerServer) {
	s.RegisterService(&_Producer_serviceDesc, srv)
}

func _Producer_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProducerServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Producer/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProducerServer).Publish(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Producer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.Producer",
	HandlerType: (*ProducerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "publish",
			Handler:    _Producer_Publish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "producer.proto",
}

func init() { proto.RegisterFile("producer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x28, 0xca, 0x4f,
	0x29, 0x4d, 0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2b, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0x52, 0x52, 0xe6, 0x62, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe0,
	0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82,
	0x71, 0x95, 0x14, 0xb9, 0x58, 0x83, 0x52, 0x0b, 0x72, 0x2a, 0x71, 0x2b, 0x31, 0x32, 0xe7, 0xe2,
	0x08, 0x80, 0xda, 0x20, 0xa4, 0xcd, 0xc5, 0x5e, 0x50, 0x9a, 0x94, 0x93, 0x59, 0x9c, 0x21, 0xc4,
	0xaf, 0x07, 0xb1, 0x47, 0x0f, 0x6a, 0x89, 0x14, 0x2f, 0x42, 0xa0, 0x20, 0xa7, 0x52, 0x89, 0xc1,
	0x49, 0x87, 0x4b, 0x2c, 0x25, 0x33, 0x37, 0xb7, 0x32, 0xab, 0x48, 0x2f, 0x3d, 0x3f, 0x3b, 0x31,
	0x2d, 0x3b, 0x11, 0xaa, 0xc2, 0x09, 0x6e, 0x60, 0x00, 0x63, 0x14, 0xd4, 0xb9, 0x49, 0x6c, 0x60,
	0xd7, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x10, 0x9e, 0x59, 0xcf, 0x00, 0x00, 0x00,
}
