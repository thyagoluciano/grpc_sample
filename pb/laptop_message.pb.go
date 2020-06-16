// Code generated by protoc-gen-go. DO NOT EDIT.
// source: laptop_message.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Laptop struct {
	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Brand    string     `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Name     string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Cpu      *CPU       `protobuf:"bytes,4,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Ram      *Memory    `protobuf:"bytes,5,opt,name=ram,proto3" json:"ram,omitempty"`
	Gpus     []*GPU     `protobuf:"bytes,6,rep,name=gpus,proto3" json:"gpus,omitempty"`
	Storages []*Storage `protobuf:"bytes,7,rep,name=storages,proto3" json:"storages,omitempty"`
	Screen   *Screen    `protobuf:"bytes,8,opt,name=screen,proto3" json:"screen,omitempty"`
	Keyboard *Keyboard  `protobuf:"bytes,9,opt,name=keyboard,proto3" json:"keyboard,omitempty"`
	// Types that are valid to be assigned to Weight:
	//	*Laptop_WeightKg
	//	*Laptop_WeightLb
	Weight               isLaptop_Weight      `protobuf_oneof:"weight"`
	PriceUsd             float64              `protobuf:"fixed64,12,opt,name=price_usd,json=priceUsd,proto3" json:"price_usd,omitempty"`
	ReleaseYear          uint32               `protobuf:"varint,13,opt,name=release_year,json=releaseYear,proto3" json:"release_year,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Laptop) Reset()         { *m = Laptop{} }
func (m *Laptop) String() string { return proto.CompactTextString(m) }
func (*Laptop) ProtoMessage()    {}
func (*Laptop) Descriptor() ([]byte, []int) {
	return fileDescriptor_07a3824d46f4b731, []int{0}
}

func (m *Laptop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Laptop.Unmarshal(m, b)
}
func (m *Laptop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Laptop.Marshal(b, m, deterministic)
}
func (m *Laptop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Laptop.Merge(m, src)
}
func (m *Laptop) XXX_Size() int {
	return xxx_messageInfo_Laptop.Size(m)
}
func (m *Laptop) XXX_DiscardUnknown() {
	xxx_messageInfo_Laptop.DiscardUnknown(m)
}

var xxx_messageInfo_Laptop proto.InternalMessageInfo

func (m *Laptop) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Laptop) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *Laptop) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Laptop) GetCpu() *CPU {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *Laptop) GetRam() *Memory {
	if m != nil {
		return m.Ram
	}
	return nil
}

func (m *Laptop) GetGpus() []*GPU {
	if m != nil {
		return m.Gpus
	}
	return nil
}

func (m *Laptop) GetStorages() []*Storage {
	if m != nil {
		return m.Storages
	}
	return nil
}

func (m *Laptop) GetScreen() *Screen {
	if m != nil {
		return m.Screen
	}
	return nil
}

func (m *Laptop) GetKeyboard() *Keyboard {
	if m != nil {
		return m.Keyboard
	}
	return nil
}

type isLaptop_Weight interface {
	isLaptop_Weight()
}

type Laptop_WeightKg struct {
	WeightKg float64 `protobuf:"fixed64,10,opt,name=weight_kg,json=weightKg,proto3,oneof"`
}

type Laptop_WeightLb struct {
	WeightLb float64 `protobuf:"fixed64,11,opt,name=weight_lb,json=weightLb,proto3,oneof"`
}

func (*Laptop_WeightKg) isLaptop_Weight() {}

func (*Laptop_WeightLb) isLaptop_Weight() {}

func (m *Laptop) GetWeight() isLaptop_Weight {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *Laptop) GetWeightKg() float64 {
	if x, ok := m.GetWeight().(*Laptop_WeightKg); ok {
		return x.WeightKg
	}
	return 0
}

func (m *Laptop) GetWeightLb() float64 {
	if x, ok := m.GetWeight().(*Laptop_WeightLb); ok {
		return x.WeightLb
	}
	return 0
}

func (m *Laptop) GetPriceUsd() float64 {
	if m != nil {
		return m.PriceUsd
	}
	return 0
}

func (m *Laptop) GetReleaseYear() uint32 {
	if m != nil {
		return m.ReleaseYear
	}
	return 0
}

func (m *Laptop) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Laptop) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Laptop_WeightKg)(nil),
		(*Laptop_WeightLb)(nil),
	}
}

func init() {
	proto.RegisterType((*Laptop)(nil), "techschool.pcbook.go.Laptop")
}

func init() {
	proto.RegisterFile("laptop_message.proto", fileDescriptor_07a3824d46f4b731)
}

var fileDescriptor_07a3824d46f4b731 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x51, 0x6f, 0xd3, 0x3e,
	0x10, 0xc0, 0xff, 0x69, 0xbb, 0xfe, 0x93, 0xeb, 0xb6, 0x07, 0xab, 0x80, 0x29, 0x8c, 0x05, 0x5e,
	0xa8, 0x84, 0xf0, 0x24, 0xe0, 0x65, 0xbc, 0x31, 0x1e, 0x40, 0xda, 0x90, 0x2a, 0x8f, 0x3e, 0xc0,
	0x4b, 0x64, 0x27, 0x87, 0x1b, 0x35, 0xa9, 0x2d, 0xdb, 0x11, 0xea, 0x97, 0xe3, 0xb3, 0xa1, 0x3a,
	0xe9, 0xc4, 0xaa, 0xc0, 0x9b, 0xfd, 0xbb, 0xdf, 0x9d, 0xee, 0x4e, 0x07, 0xd3, 0x4a, 0x18, 0xaf,
	0x4d, 0x56, 0xa3, 0x73, 0x42, 0x21, 0x33, 0x56, 0x7b, 0x4d, 0xa6, 0x1e, 0xf3, 0x95, 0xcb, 0x57,
	0x5a, 0x57, 0xcc, 0xe4, 0x52, 0xeb, 0x35, 0x53, 0x7a, 0xf6, 0xc8, 0x58, 0x9d, 0xa3, 0x73, 0xda,
	0xde, 0xd7, 0x67, 0xd3, 0x1a, 0x6b, 0x6d, 0xb7, 0x07, 0xf4, 0x81, 0xf3, 0xda, 0x0a, 0x85, 0x87,
	0xb2, 0xcb, 0x2d, 0xe2, 0xe6, 0x80, 0x3e, 0x5c, 0xe3, 0x56, 0x6a, 0x61, 0x8b, 0x03, 0x7e, 0xae,
	0xb4, 0x56, 0x15, 0x5e, 0x84, 0x9f, 0x6c, 0x7e, 0x5c, 0xf8, 0xb2, 0x46, 0xe7, 0x45, 0x6d, 0x5a,
	0xe1, 0xc5, 0xaf, 0x11, 0x8c, 0x6f, 0xc2, 0x0c, 0xe4, 0x14, 0x06, 0x65, 0x41, 0xa3, 0x34, 0x9a,
	0x27, 0x7c, 0x50, 0x16, 0x64, 0x0a, 0x47, 0xd2, 0x8a, 0x4d, 0x41, 0x07, 0x01, 0xb5, 0x1f, 0x42,
	0x60, 0xb4, 0x11, 0x35, 0xd2, 0x61, 0x80, 0xe1, 0x4d, 0x5e, 0xc1, 0x30, 0x37, 0x0d, 0x1d, 0xa5,
	0xd1, 0x7c, 0xf2, 0xe6, 0x31, 0xeb, 0x9b, 0x9e, 0x7d, 0x5c, 0x2c, 0xf9, 0xce, 0x22, 0x0c, 0x86,
	0x56, 0xd4, 0xf4, 0x28, 0xc8, 0x4f, 0xfb, 0xe5, 0x2f, 0x61, 0x21, 0x7c, 0x27, 0x92, 0xd7, 0x30,
	0x52, 0xa6, 0x71, 0x74, 0x9c, 0x0e, 0xff, 0x5e, 0xfd, 0xd3, 0x62, 0xc9, 0x83, 0x46, 0x2e, 0x21,
	0xee, 0x16, 0xe7, 0xe8, 0xff, 0x21, 0xe5, 0xac, 0x3f, 0xe5, 0xb6, 0xb5, 0xf8, 0x9d, 0x4e, 0xde,
	0xc1, 0xb8, 0x5d, 0x2e, 0x8d, 0xff, 0xd5, 0xdc, 0x6d, 0x70, 0x78, 0xe7, 0x92, 0xf7, 0x10, 0xef,
	0x97, 0x4f, 0x93, 0x90, 0xf7, 0xac, 0x3f, 0xef, 0xba, 0xb3, 0xf8, 0x9d, 0x4f, 0xce, 0x20, 0xf9,
	0x89, 0xa5, 0x5a, 0xf9, 0x6c, 0xad, 0x28, 0xa4, 0xd1, 0x3c, 0xfa, 0xfc, 0x1f, 0x8f, 0x5b, 0x74,
	0xad, 0xfe, 0x08, 0x57, 0x92, 0x4e, 0xee, 0x87, 0x6f, 0x24, 0x79, 0x02, 0x89, 0xb1, 0x65, 0x8e,
	0x59, 0xe3, 0x0a, 0x7a, 0xbc, 0x0b, 0xf3, 0x38, 0x80, 0xa5, 0x2b, 0xc8, 0x73, 0x38, 0xb6, 0x58,
	0xa1, 0x70, 0x98, 0x6d, 0x51, 0x58, 0x7a, 0x92, 0x46, 0xf3, 0x13, 0x3e, 0xe9, 0xd8, 0x37, 0x14,
	0x96, 0x5c, 0x02, 0x34, 0xa6, 0x10, 0x1e, 0x8b, 0x4c, 0x78, 0x7a, 0x1a, 0x7a, 0x9f, 0xb1, 0xf6,
	0x62, 0xd8, 0xfe, 0x62, 0xd8, 0xd7, 0xfd, 0xc5, 0xf0, 0xa4, 0xb3, 0x3f, 0xf8, 0xab, 0x18, 0xc6,
	0x6d, 0x1b, 0x57, 0x2f, 0xe1, 0x3c, 0xd7, 0x35, 0x53, 0xa5, 0xaf, 0x84, 0xec, 0x19, 0xdc, 0xc8,
	0x45, 0xf4, 0x7d, 0x60, 0xa4, 0x1c, 0x87, 0x8a, 0x6f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x22,
	0xc6, 0x39, 0xa4, 0x33, 0x03, 0x00, 0x00,
}
