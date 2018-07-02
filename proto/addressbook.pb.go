// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto_test.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Person_PhoneType int32

const (
	Person_Mobile Person_PhoneType = 0
	Person_Home   Person_PhoneType = 1
	Person_Work   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "Mobile",
	1: "Home",
	2: "Work",
}
var Person_PhoneType_value = map[string]int32{
	"Mobile": 0,
	"Home":   1,
	"Work":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}
func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_proto_test_02d3b75d474a8004, []int{0, 0}
}

type Person struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email                string                `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto_test_02d3b75d474a8004, []int{0}
}
func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (dst *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(dst, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type Person_PhoneNumber struct {
	Number               string           `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 Person_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=tutorial.Person_PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto_test_02d3b75d474a8004, []int{0, 0}
}
func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (dst *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(dst, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_Mobile
}

type AddressBook struct {
	People               []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}
func (*AddressBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto_test_02d3b75d474a8004, []int{1}
}
func (m *AddressBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBook.Unmarshal(m, b)
}
func (m *AddressBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBook.Marshal(b, m, deterministic)
}
func (dst *AddressBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBook.Merge(dst, src)
}
func (m *AddressBook) XXX_Size() int {
	return xxx_messageInfo_AddressBook.Size(m)
}
func (m *AddressBook) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBook.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBook proto.InternalMessageInfo

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "tutorial.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "tutorial.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "tutorial.AddressBook")
	proto.RegisterEnum("tutorial.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
}

func init() { proto.RegisterFile("proto_test.proto", fileDescriptor_proto_test_02d3b75d474a8004) }

var fileDescriptor_proto_test_02d3b75d474a8004 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4b, 0xf3, 0x40,
	0x10, 0xc5, 0xbf, 0xa4, 0xe9, 0xd2, 0x4e, 0xa0, 0x2c, 0xc3, 0x87, 0x84, 0xe2, 0x21, 0xe4, 0x14,
	0x10, 0xf6, 0x50, 0x05, 0xcf, 0x7a, 0xf2, 0xa2, 0x94, 0x45, 0xf1, 0x28, 0x09, 0x19, 0x70, 0x69,
	0x92, 0x59, 0x76, 0xb7, 0x87, 0xfe, 0xef, 0x1e, 0x24, 0xdb, 0x28, 0x22, 0xde, 0x7e, 0x33, 0xf3,
	0x78, 0xfb, 0xf6, 0x81, 0xb4, 0x8e, 0x03, 0xbf, 0x05, 0xf2, 0x41, 0x45, 0xc4, 0x55, 0x38, 0x06,
	0x76, 0xa6, 0xe9, 0xab, 0x8f, 0x04, 0xc4, 0x9e, 0x9c, 0xe7, 0x11, 0x11, 0xb2, 0xb1, 0x19, 0xa8,
	0x48, 0xca, 0xa4, 0x5e, 0xeb, 0xc8, 0xb8, 0x81, 0xd4, 0x74, 0x45, 0x5a, 0x26, 0xf5, 0x52, 0xa7,
	0xa6, 0xc3, 0xff, 0xb0, 0xa4, 0xa1, 0x31, 0x7d, 0xb1, 0x88, 0xa2, 0xf3, 0x80, 0x37, 0x20, 0xec,
	0x3b, 0x8f, 0xe4, 0x8b, 0xac, 0x5c, 0xd4, 0xf9, 0xee, 0x52, 0x7d, 0xf9, 0xab, 0xb3, 0xb7, 0xda,
	0x4f, 0xe7, 0xa7, 0xe3, 0xd0, 0x92, 0xd3, 0xb3, 0x76, 0xfb, 0x02, 0xf9, 0x8f, 0x35, 0x5e, 0x80,
	0x18, 0x23, 0xcd, 0x01, 0xe6, 0x09, 0x15, 0x64, 0xe1, 0x64, 0x29, 0x86, 0xd8, 0xec, 0xb6, 0x7f,
	0x5b, 0x3f, 0x9f, 0x2c, 0xe9, 0xa8, 0xab, 0xae, 0x60, 0xfd, 0xbd, 0x42, 0x00, 0xf1, 0xc8, 0xad,
	0xe9, 0x49, 0xfe, 0xc3, 0x15, 0x64, 0x0f, 0x3c, 0x90, 0x4c, 0x26, 0x7a, 0x65, 0x77, 0x90, 0x69,
	0x75, 0x0b, 0xf9, 0x5d, 0xd7, 0x39, 0xf2, 0xfe, 0x9e, 0xf9, 0x80, 0x35, 0x08, 0x4b, 0x6c, 0xfb,
	0xa9, 0x84, 0xe9, 0x23, 0xf2, 0xf7, 0x6b, 0x7a, 0xbe, 0xb7, 0x22, 0x16, 0x79, 0xfd, 0x19, 0x00,
	0x00, 0xff, 0xff, 0x7c, 0xc0, 0x43, 0x6c, 0x5c, 0x01, 0x00, 0x00,
}
