// Code generated by protoc-gen-go. DO NOT EDIT.
// source: meetup.proto

/*
Package gen is a generated protocol buffer package.

It is generated from these files:
	meetup.proto

It has these top-level messages:
	WelcomMessages
	WelcomMessage
	Person
	ThankYouResponse
*/
package gen

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Gender int32

const (
	Gender_MALE   Gender = 0
	Gender_FEMALE Gender = 1
)

var Gender_name = map[int32]string{
	0: "MALE",
	1: "FEMALE",
}
var Gender_value = map[string]int32{
	"MALE":   0,
	"FEMALE": 1,
}

func (x Gender) String() string {
	return proto.EnumName(Gender_name, int32(x))
}
func (Gender) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type WelcomMessages struct {
	Messages []*WelcomMessage `protobuf:"bytes,1,rep,name=messages" json:"messages,omitempty"`
}

func (m *WelcomMessages) Reset()                    { *m = WelcomMessages{} }
func (m *WelcomMessages) String() string            { return proto.CompactTextString(m) }
func (*WelcomMessages) ProtoMessage()               {}
func (*WelcomMessages) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WelcomMessages) GetMessages() []*WelcomMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

type WelcomMessage struct {
	Person  *Person `protobuf:"bytes,1,opt,name=person" json:"person,omitempty"`
	Message string  `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *WelcomMessage) Reset()                    { *m = WelcomMessage{} }
func (m *WelcomMessage) String() string            { return proto.CompactTextString(m) }
func (*WelcomMessage) ProtoMessage()               {}
func (*WelcomMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *WelcomMessage) GetPerson() *Person {
	if m != nil {
		return m.Person
	}
	return nil
}

func (m *WelcomMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Person struct {
	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Age       uint32 `protobuf:"varint,3,opt,name=age" json:"age,omitempty"`
	Gender    Gender `protobuf:"varint,4,opt,name=gender,enum=Gender" json:"gender,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Person) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Person) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Person) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Person) GetGender() Gender {
	if m != nil {
		return m.Gender
	}
	return Gender_MALE
}

type ThankYouResponse struct {
	MessageResponse string `protobuf:"bytes,1,opt,name=message_response,json=messageResponse" json:"message_response,omitempty"`
}

func (m *ThankYouResponse) Reset()                    { *m = ThankYouResponse{} }
func (m *ThankYouResponse) String() string            { return proto.CompactTextString(m) }
func (*ThankYouResponse) ProtoMessage()               {}
func (*ThankYouResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ThankYouResponse) GetMessageResponse() string {
	if m != nil {
		return m.MessageResponse
	}
	return ""
}

func init() {
	proto.RegisterType((*WelcomMessages)(nil), "WelcomMessages")
	proto.RegisterType((*WelcomMessage)(nil), "WelcomMessage")
	proto.RegisterType((*Person)(nil), "Person")
	proto.RegisterType((*ThankYouResponse)(nil), "ThankYouResponse")
	proto.RegisterEnum("Gender", Gender_name, Gender_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Hello service

type HelloClient interface {
	SayHello(ctx context.Context, in *WelcomMessages, opts ...grpc.CallOption) (*ThankYouResponse, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) SayHello(ctx context.Context, in *WelcomMessages, opts ...grpc.CallOption) (*ThankYouResponse, error) {
	out := new(ThankYouResponse)
	err := grpc.Invoke(ctx, "/Hello/sayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hello service

type HelloServer interface {
	SayHello(context.Context, *WelcomMessages) (*ThankYouResponse, error)
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WelcomMessages)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Hello/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).SayHello(ctx, req.(*WelcomMessages))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sayHello",
			Handler:    _Hello_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "meetup.proto",
}

func init() { proto.RegisterFile("meetup.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xc1, 0x4b, 0xf4, 0x30,
	0x10, 0xc5, 0xbf, 0x7c, 0x5d, 0xbb, 0xed, 0xac, 0xbb, 0x5b, 0x73, 0x0a, 0x8a, 0x5a, 0x7a, 0xaa,
	0x7b, 0x88, 0x50, 0x0f, 0x5e, 0xf4, 0xa0, 0xb0, 0x2a, 0xe2, 0x8a, 0x04, 0x41, 0xf4, 0xb2, 0x44,
	0x1d, 0xab, 0xd8, 0x26, 0xa5, 0xe9, 0x22, 0xfe, 0xf7, 0xb2, 0x69, 0x2a, 0xd4, 0xdb, 0x9b, 0xdf,
	0x9b, 0x3c, 0x1e, 0x19, 0xd8, 0x2c, 0x11, 0x9b, 0x55, 0xc5, 0xab, 0x5a, 0x37, 0x3a, 0x39, 0x81,
	0xc9, 0x03, 0x16, 0x2f, 0xba, 0x5c, 0xa0, 0x31, 0x32, 0x47, 0x43, 0x67, 0x10, 0x94, 0x4e, 0x33,
	0x12, 0x7b, 0xe9, 0x28, 0x9b, 0xf0, 0xde, 0x8a, 0xf8, 0xf5, 0x93, 0x6b, 0x18, 0xf7, 0x2c, 0xba,
	0x0f, 0x7e, 0x85, 0xb5, 0xd1, 0x8a, 0x91, 0x98, 0xa4, 0xa3, 0x6c, 0xc8, 0xef, 0xec, 0x28, 0x1c,
	0xa6, 0x0c, 0x86, 0xee, 0x35, 0xfb, 0x1f, 0x93, 0x34, 0x14, 0xdd, 0x98, 0x7c, 0x81, 0xdf, 0xee,
	0xd2, 0x5d, 0x80, 0xb7, 0x8f, 0xda, 0x34, 0x4b, 0x25, 0x4b, 0xb4, 0x41, 0xa1, 0x08, 0x2d, 0xb9,
	0x95, 0x25, 0xd2, 0x1d, 0x08, 0x0b, 0xd9, 0xb9, 0x6d, 0x48, 0xb0, 0x06, 0xd6, 0x8c, 0xc0, 0x5b,
	0x67, 0x7b, 0x31, 0x49, 0xc7, 0xc2, 0x73, 0x95, 0x72, 0x54, 0xaf, 0x58, 0xb3, 0x41, 0x4c, 0xd2,
	0x49, 0x36, 0xe4, 0x97, 0x76, 0x14, 0x0e, 0x27, 0xa7, 0x10, 0xdd, 0xbf, 0x4b, 0xf5, 0xf9, 0xa8,
	0x57, 0x02, 0x4d, 0xa5, 0x95, 0x41, 0x7a, 0x00, 0x91, 0xeb, 0xb5, 0xac, 0x1d, 0x73, 0x45, 0xa6,
	0x8e, 0x77, 0xab, 0xb3, 0x3d, 0xf0, 0xdb, 0x40, 0x1a, 0xc0, 0x60, 0x71, 0x76, 0x33, 0x8f, 0xfe,
	0x51, 0x00, 0xff, 0x62, 0x6e, 0x35, 0xc9, 0x8e, 0x61, 0xe3, 0x0a, 0x8b, 0x42, 0x53, 0x0e, 0x81,
	0x91, 0xdf, 0xad, 0x9e, 0xf6, 0xbf, 0xd4, 0x6c, 0x6f, 0xf1, 0xbf, 0x1d, 0xce, 0x19, 0x80, 0x3b,
	0x55, 0x8e, 0xea, 0xc9, 0xe9, 0xc3, 0x1c, 0xd5, 0xb3, 0x6f, 0x6f, 0x77, 0xf4, 0x13, 0x00, 0x00,
	0xff, 0xff, 0xae, 0xb9, 0x48, 0x10, 0xcb, 0x01, 0x00, 0x00,
}
