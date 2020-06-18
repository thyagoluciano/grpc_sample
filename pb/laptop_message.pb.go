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
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0x80, 0x49, 0xdb, 0x95, 0xe4, 0xba, 0xed, 0xc1, 0xaa, 0xc0, 0x14, 0x06, 0x01, 0x5e, 0x2a,
	0x4d, 0x78, 0x12, 0xf0, 0x32, 0xde, 0x18, 0x0f, 0x20, 0x6d, 0x48, 0x95, 0x47, 0x1f, 0xe0, 0x25,
	0xb2, 0x93, 0xc3, 0x0d, 0x4d, 0x6a, 0xcb, 0x76, 0x40, 0xfd, 0x73, 0xfc, 0x36, 0x54, 0x27, 0x9d,
	0xb6, 0x52, 0x78, 0xb3, 0xbf, 0xfb, 0xee, 0x74, 0x77, 0x3a, 0x18, 0x57, 0xc2, 0x78, 0x6d, 0xb2,
	0x1a, 0x9d, 0x13, 0x0a, 0x99, 0xb1, 0xda, 0x6b, 0x32, 0xf6, 0x98, 0x2f, 0x5c, 0xbe, 0xd0, 0xba,
	0x62, 0x26, 0x97, 0x5a, 0x2f, 0x99, 0xd2, 0x93, 0x87, 0xc6, 0xea, 0x1c, 0x9d, 0xd3, 0xf6, 0xae,
	0x3e, 0x19, 0xd7, 0x58, 0x6b, 0xbb, 0xde, 0xa5, 0xce, 0x6b, 0x2b, 0x14, 0x06, 0x7c, 0x8b, 0xe6,
	0x16, 0x71, 0xb5, 0xe3, 0x3e, 0x58, 0xe2, 0x5a, 0x6a, 0x61, 0x8b, 0x1d, 0xfe, 0x4c, 0x69, 0xad,
	0x2a, 0x3c, 0x0b, 0x3f, 0xd9, 0x7c, 0x3f, 0xf3, 0x65, 0x8d, 0xce, 0x8b, 0xda, 0xb4, 0xc2, 0x8b,
	0xdf, 0x03, 0x18, 0x5e, 0x85, 0x11, 0xc8, 0x31, 0xf4, 0xca, 0x82, 0x46, 0x69, 0x34, 0x4d, 0x78,
	0xaf, 0x2c, 0xc8, 0x18, 0x0e, 0xa4, 0x15, 0xab, 0x82, 0xf6, 0x02, 0x6a, 0x3f, 0x84, 0xc0, 0x60,
	0x25, 0x6a, 0xa4, 0xfd, 0x00, 0xc3, 0x9b, 0x9c, 0x42, 0x3f, 0x37, 0x0d, 0x1d, 0xa4, 0xd1, 0x74,
	0xf4, 0xfa, 0x11, 0xdb, 0x37, 0x3c, 0xfb, 0x30, 0x9b, 0xf3, 0x8d, 0x45, 0x18, 0xf4, 0xad, 0xa8,
	0xe9, 0x41, 0x90, 0x9f, 0xec, 0x97, 0x3f, 0x87, 0x7d, 0xf0, 0x8d, 0x48, 0x5e, 0xc1, 0x40, 0x99,
	0xc6, 0xd1, 0x61, 0xda, 0xff, 0x77, 0xf5, 0x8f, 0xb3, 0x39, 0x0f, 0x1a, 0x39, 0x87, 0xb8, 0xdb,
	0x9b, 0xa3, 0xf7, 0x43, 0xca, 0xc9, 0xfe, 0x94, 0xeb, 0xd6, 0xe2, 0x37, 0x3a, 0x79, 0x0b, 0xc3,
	0x76, 0xb9, 0x34, 0xfe, 0x5f, 0x73, 0xd7, 0xc1, 0xe1, 0x9d, 0x4b, 0xde, 0x41, 0xbc, 0x5d, 0x3e,
	0x4d, 0x42, 0xde, 0xd3, 0xfd, 0x79, 0x97, 0x9d, 0xc5, 0x6f, 0x7c, 0x72, 0x02, 0xc9, 0x2f, 0x2c,
	0xd5, 0xc2, 0x67, 0x4b, 0x45, 0x21, 0x8d, 0xa6, 0xd1, 0xa7, 0x7b, 0x3c, 0x6e, 0xd1, 0xa5, 0xba,
	0x15, 0xae, 0x24, 0x1d, 0xdd, 0x0d, 0x5f, 0x49, 0xf2, 0x18, 0x12, 0x63, 0xcb, 0x1c, 0xb3, 0xc6,
	0x15, 0xf4, 0x70, 0x13, 0xe6, 0x71, 0x00, 0x73, 0x57, 0x90, 0xe7, 0x70, 0x68, 0xb1, 0x42, 0xe1,
	0x30, 0x5b, 0xa3, 0xb0, 0xf4, 0x28, 0x8d, 0xa6, 0x47, 0x7c, 0xd4, 0xb1, 0xaf, 0x28, 0x2c, 0x39,
	0x07, 0x68, 0x4c, 0x21, 0x3c, 0x16, 0x99, 0xf0, 0xf4, 0x38, 0xf4, 0x3e, 0x61, 0xed, 0xc5, 0xb0,
	0xed, 0xc5, 0xb0, 0x2f, 0xdb, 0x8b, 0xe1, 0x49, 0x67, 0xbf, 0xf7, 0x17, 0x31, 0x0c, 0xdb, 0x36,
	0x2e, 0x4e, 0xe1, 0x65, 0xae, 0x6b, 0xa6, 0x4a, 0x5f, 0x09, 0xf9, 0xf7, 0xe0, 0x3f, 0xc4, 0x4f,
	0xc1, 0x8c, 0x9c, 0x45, 0xdf, 0x7a, 0x46, 0xca, 0x61, 0xa8, 0xfa, 0xe6, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x12, 0x7b, 0x65, 0xf2, 0x36, 0x03, 0x00, 0x00,
}
