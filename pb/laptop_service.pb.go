// Code generated by protoc-gen-go. DO NOT EDIT.
// source: laptop_service.proto

package pb

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

type CreateLaptopRequest struct {
	Laptop               *Laptop  `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLaptopRequest) Reset()         { *m = CreateLaptopRequest{} }
func (m *CreateLaptopRequest) String() string { return proto.CompactTextString(m) }
func (*CreateLaptopRequest) ProtoMessage()    {}
func (*CreateLaptopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{0}
}

func (m *CreateLaptopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLaptopRequest.Unmarshal(m, b)
}
func (m *CreateLaptopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLaptopRequest.Marshal(b, m, deterministic)
}
func (m *CreateLaptopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLaptopRequest.Merge(m, src)
}
func (m *CreateLaptopRequest) XXX_Size() int {
	return xxx_messageInfo_CreateLaptopRequest.Size(m)
}
func (m *CreateLaptopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLaptopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLaptopRequest proto.InternalMessageInfo

func (m *CreateLaptopRequest) GetLaptop() *Laptop {
	if m != nil {
		return m.Laptop
	}
	return nil
}

type CreateLapTopResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLapTopResponse) Reset()         { *m = CreateLapTopResponse{} }
func (m *CreateLapTopResponse) String() string { return proto.CompactTextString(m) }
func (*CreateLapTopResponse) ProtoMessage()    {}
func (*CreateLapTopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{1}
}

func (m *CreateLapTopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLapTopResponse.Unmarshal(m, b)
}
func (m *CreateLapTopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLapTopResponse.Marshal(b, m, deterministic)
}
func (m *CreateLapTopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLapTopResponse.Merge(m, src)
}
func (m *CreateLapTopResponse) XXX_Size() int {
	return xxx_messageInfo_CreateLapTopResponse.Size(m)
}
func (m *CreateLapTopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLapTopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLapTopResponse proto.InternalMessageInfo

func (m *CreateLapTopResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateLaptopRequest)(nil), "techschool.pcbook.go.CreateLaptopRequest")
	proto.RegisterType((*CreateLapTopResponse)(nil), "techschool.pcbook.go.CreateLapTopResponse")
}

func init() {
	proto.RegisterFile("laptop_service.proto", fileDescriptor_240c60d9fb227e71)
}

var fileDescriptor_240c60d9fb227e71 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0x49, 0x2c, 0x28,
	0xc9, 0x2f, 0x88, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0x29, 0x49, 0x4d, 0xce, 0x28, 0x4e, 0xce, 0xc8, 0xcf, 0xcf, 0xd1, 0x2b, 0x48, 0x4e,
	0xca, 0xcf, 0xcf, 0xd6, 0x4b, 0xcf, 0x97, 0x82, 0xa9, 0xcd, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x87,
	0xaa, 0x55, 0xf2, 0xe6, 0x12, 0x76, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0xf5, 0x01, 0xcb, 0x06, 0xa5,
	0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x99, 0x70, 0xb1, 0x41, 0x94, 0x4b, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x1b, 0xc9, 0xe8, 0x61, 0x33, 0x53, 0x0f, 0xaa, 0x09, 0xaa, 0x56, 0x49, 0x8d, 0x4b, 0x04,
	0x6e, 0x58, 0x08, 0xc8, 0xb0, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x3e, 0x2e, 0xa6, 0xcc,
	0x14, 0xb0, 0x49, 0x9c, 0x41, 0x4c, 0x99, 0x29, 0x46, 0x15, 0x5c, 0xbc, 0x10, 0x15, 0xc1, 0x10,
	0x77, 0x0b, 0xa5, 0x73, 0xf1, 0x20, 0xbb, 0x42, 0x48, 0x13, 0xbb, 0x75, 0x58, 0x5c, 0x2a, 0xa5,
	0x45, 0x40, 0x29, 0x92, 0x3b, 0x94, 0x18, 0x9c, 0xb4, 0xb9, 0x94, 0x93, 0xf3, 0x73, 0xf5, 0xd2,
	0x33, 0x4b, 0x72, 0x12, 0x93, 0x30, 0x75, 0x66, 0x25, 0x96, 0x25, 0xea, 0x15, 0x24, 0x05, 0x30,
	0x46, 0x31, 0x15, 0x24, 0x25, 0xb1, 0x81, 0x83, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf2,
	0x95, 0x13, 0x34, 0x66, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LapTopServiceClient is the client API for LapTopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LapTopServiceClient interface {
	CreateLaptop(ctx context.Context, in *CreateLaptopRequest, opts ...grpc.CallOption) (*CreateLapTopResponse, error)
}

type lapTopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLapTopServiceClient(cc grpc.ClientConnInterface) LapTopServiceClient {
	return &lapTopServiceClient{cc}
}

func (c *lapTopServiceClient) CreateLaptop(ctx context.Context, in *CreateLaptopRequest, opts ...grpc.CallOption) (*CreateLapTopResponse, error) {
	out := new(CreateLapTopResponse)
	err := c.cc.Invoke(ctx, "/techschool.pcbook.go.LapTopService/CreateLaptop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LapTopServiceServer is the server API for LapTopService service.
type LapTopServiceServer interface {
	CreateLaptop(context.Context, *CreateLaptopRequest) (*CreateLapTopResponse, error)
}

// UnimplementedLapTopServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLapTopServiceServer struct {
}

func (*UnimplementedLapTopServiceServer) CreateLaptop(ctx context.Context, req *CreateLaptopRequest) (*CreateLapTopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLaptop not implemented")
}

func RegisterLapTopServiceServer(s *grpc.Server, srv LapTopServiceServer) {
	s.RegisterService(&_LapTopService_serviceDesc, srv)
}

func _LapTopService_CreateLaptop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLaptopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LapTopServiceServer).CreateLaptop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/techschool.pcbook.go.LapTopService/CreateLaptop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LapTopServiceServer).CreateLaptop(ctx, req.(*CreateLaptopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LapTopService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "techschool.pcbook.go.LapTopService",
	HandlerType: (*LapTopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLaptop",
			Handler:    _LapTopService_CreateLaptop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "laptop_service.proto",
}