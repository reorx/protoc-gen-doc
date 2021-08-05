// Code generated by protoc-gen-go. DO NOT EDIT.
// source: page.proto

package hugo

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
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

type Page struct {
	Weight               int32    `protobuf:"varint,1,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Page) Reset()         { *m = Page{} }
func (m *Page) String() string { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()    {}
func (*Page) Descriptor() ([]byte, []int) {
	return fileDescriptor_f14a105a5ef2e917, []int{0}
}

func (m *Page) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Page.Unmarshal(m, b)
}
func (m *Page) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Page.Marshal(b, m, deterministic)
}
func (m *Page) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Page.Merge(m, src)
}
func (m *Page) XXX_Size() int {
	return xxx_messageInfo_Page.Size(m)
}
func (m *Page) XXX_DiscardUnknown() {
	xxx_messageInfo_Page.DiscardUnknown(m)
}

var xxx_messageInfo_Page proto.InternalMessageInfo

func (m *Page) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

var E_Page = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.FileOptions)(nil),
	ExtensionType: (*Page)(nil),
	Field:         17736974,
	Name:          "hugo.page",
	Tag:           "bytes,17736974,opt,name=page",
	Filename:      "page.proto",
}

func init() {
	proto.RegisterType((*Page)(nil), "hugo.Page")
	proto.RegisterExtension(E_Page)
}

func init() { proto.RegisterFile("page.proto", fileDescriptor_f14a105a5ef2e917) }

var fileDescriptor_f14a105a5ef2e917 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x48, 0x4c, 0x4f,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x28, 0x4d, 0xcf, 0x97, 0x52, 0x48, 0xcf,
	0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x25, 0x95, 0xa6, 0xe9, 0xa7, 0xa4, 0x16, 0x27, 0x17,
	0x65, 0x16, 0x94, 0xe4, 0x17, 0x41, 0xd4, 0x29, 0xc9, 0x71, 0xb1, 0x04, 0x24, 0xa6, 0xa7, 0x0a,
	0x89, 0x71, 0xb1, 0x95, 0xa7, 0x66, 0xa6, 0x67, 0x94, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06,
	0x41, 0x79, 0x56, 0xf6, 0x5c, 0x2c, 0x20, 0x53, 0x85, 0x64, 0xf4, 0x20, 0x46, 0xe9, 0xc1, 0x8c,
	0xd2, 0x73, 0xcb, 0xcc, 0x49, 0xf5, 0x2f, 0x28, 0xc9, 0xcc, 0xcf, 0x2b, 0x96, 0xe8, 0x3b, 0xb5,
	0x8b, 0x43, 0x81, 0x51, 0x83, 0xdb, 0x88, 0x4b, 0x0f, 0x64, 0xaf, 0x1e, 0xc8, 0xc8, 0x20, 0xb0,
	0xc6, 0x24, 0x36, 0xb0, 0x06, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xb3, 0x6d, 0x97,
	0x9d, 0x00, 0x00, 0x00,
}
