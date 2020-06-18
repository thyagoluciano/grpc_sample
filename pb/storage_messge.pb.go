// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage_messge.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Storage_Driver int32

const (
	Storage_UNKNOWN Storage_Driver = 0
	Storage_HDD     Storage_Driver = 1
	Storage_SSD     Storage_Driver = 2
)

var Storage_Driver_name = map[int32]string{
	0: "UNKNOWN",
	1: "HDD",
	2: "SSD",
}

var Storage_Driver_value = map[string]int32{
	"UNKNOWN": 0,
	"HDD":     1,
	"SSD":     2,
}

func (x Storage_Driver) String() string {
	return proto.EnumName(Storage_Driver_name, int32(x))
}

func (Storage_Driver) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4cf460078884b782, []int{0, 0}
}

type Storage struct {
	Driver               Storage_Driver `protobuf:"varint,1,opt,name=driver,proto3,enum=techschool.pcbook.go.Storage_Driver" json:"driver,omitempty"`
	Memory               *Memory        `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Storage) Reset()         { *m = Storage{} }
func (m *Storage) String() string { return proto.CompactTextString(m) }
func (*Storage) ProtoMessage()    {}
func (*Storage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cf460078884b782, []int{0}
}

func (m *Storage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Storage.Unmarshal(m, b)
}
func (m *Storage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Storage.Marshal(b, m, deterministic)
}
func (m *Storage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storage.Merge(m, src)
}
func (m *Storage) XXX_Size() int {
	return xxx_messageInfo_Storage.Size(m)
}
func (m *Storage) XXX_DiscardUnknown() {
	xxx_messageInfo_Storage.DiscardUnknown(m)
}

var xxx_messageInfo_Storage proto.InternalMessageInfo

func (m *Storage) GetDriver() Storage_Driver {
	if m != nil {
		return m.Driver
	}
	return Storage_UNKNOWN
}

func (m *Storage) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func init() {
	proto.RegisterEnum("techschool.pcbook.go.Storage_Driver", Storage_Driver_name, Storage_Driver_value)
	proto.RegisterType((*Storage)(nil), "techschool.pcbook.go.Storage")
}

func init() {
	proto.RegisterFile("storage_messge.proto", fileDescriptor_4cf460078884b782)
}

var fileDescriptor_4cf460078884b782 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0x8d, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0x29, 0x49, 0x4d, 0xce, 0x28, 0x4e, 0xce, 0xc8, 0xcf, 0xcf, 0xd1, 0x2b, 0x48, 0x4e,
	0xca, 0xcf, 0xcf, 0xd6, 0x4b, 0xcf, 0x97, 0x12, 0xc9, 0x4d, 0xcd, 0xcd, 0x2f, 0xaa, 0x04, 0x2b,
	0x4d, 0x84, 0xa9, 0x55, 0x5a, 0xc6, 0xc8, 0xc5, 0x1e, 0x0c, 0x31, 0x44, 0xc8, 0x86, 0x8b, 0x2d,
	0xa5, 0x28, 0xb3, 0x2c, 0xb5, 0x48, 0x82, 0x51, 0x81, 0x51, 0x83, 0xcf, 0x48, 0x45, 0x0f, 0x9b,
	0x41, 0x7a, 0x50, 0xe5, 0x7a, 0x2e, 0x60, 0xb5, 0x41, 0x50, 0x3d, 0x42, 0x26, 0x5c, 0x6c, 0x10,
	0x1b, 0x24, 0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0x64, 0xb0, 0xeb, 0xf6, 0x05, 0xab, 0x09, 0x82,
	0xaa, 0x55, 0x52, 0xe7, 0x62, 0x83, 0x98, 0x23, 0xc4, 0xcd, 0xc5, 0x1e, 0xea, 0xe7, 0xed, 0xe7,
	0x1f, 0xee, 0x27, 0xc0, 0x20, 0xc4, 0xce, 0xc5, 0xec, 0xe1, 0xe2, 0x22, 0xc0, 0x08, 0x62, 0x04,
	0x07, 0xbb, 0x08, 0x30, 0x39, 0x69, 0x73, 0x29, 0x27, 0xe7, 0xe7, 0xea, 0xa5, 0x67, 0x96, 0xe4,
	0x24, 0x26, 0x61, 0x1a, 0x9d, 0x95, 0x58, 0x96, 0xa8, 0x57, 0x90, 0x14, 0xc0, 0x18, 0xc5, 0x54,
	0x90, 0x94, 0xc4, 0x06, 0xf6, 0x9c, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x83, 0xdf, 0xa3, 0xe2,
	0x20, 0x01, 0x00, 0x00,
}