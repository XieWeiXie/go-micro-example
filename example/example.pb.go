// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package example

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

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}

var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}

func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0, 0}
}

type Person struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email                string                `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	LastUpdated          *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
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

func (m *Person) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type Person_PhoneNumber struct {
	Number               string           `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 Person_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=example.Person_PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0, 0}
}

func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(m, src)
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
	return Person_MOBILE
}

// Our address book file is just one of these.
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
	return fileDescriptor_15a1dc8d40dadaa6, []int{1}
}

func (m *AddressBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBook.Unmarshal(m, b)
}
func (m *AddressBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBook.Marshal(b, m, deterministic)
}
func (m *AddressBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBook.Merge(m, src)
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
	proto.RegisterEnum("example.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
	proto.RegisterType((*Person)(nil), "example.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "example.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "example.AddressBook")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0xbf, 0xa4, 0x69, 0x3e, 0x3b, 0xd1, 0x5a, 0x06, 0x91, 0x58, 0x0f, 0x86, 0x5e, 0x0c,
	0x88, 0x29, 0xb4, 0xe0, 0xcd, 0x83, 0x85, 0x82, 0xa2, 0xb5, 0x65, 0xa9, 0x78, 0x94, 0x94, 0x1d,
	0x6b, 0x30, 0x9b, 0x5d, 0xb2, 0x5b, 0xb0, 0x3f, 0xcd, 0x7f, 0x27, 0xd9, 0xa4, 0x22, 0xe2, 0x6d,
	0xde, 0x9d, 0x87, 0xd9, 0x67, 0x06, 0x0e, 0xe8, 0x23, 0x15, 0x2a, 0xa7, 0x44, 0x95, 0xd2, 0x48,
	0xfc, 0xdf, 0xc4, 0xfe, 0xd9, 0x5a, 0xca, 0x75, 0x4e, 0x43, 0xfb, 0xbc, 0xda, 0xbc, 0x0e, 0x4d,
	0x26, 0x48, 0x9b, 0x54, 0xa8, 0x9a, 0x1c, 0x7c, 0xba, 0xe0, 0x2f, 0xa8, 0xd4, 0xb2, 0x40, 0x04,
	0xaf, 0x48, 0x05, 0x85, 0x4e, 0xe4, 0xc4, 0x1d, 0x66, 0x6b, 0xec, 0x82, 0x9b, 0xf1, 0xd0, 0x8d,
	0x9c, 0xb8, 0xcd, 0xdc, 0x8c, 0xe3, 0x11, 0xb4, 0x49, 0xa4, 0x59, 0x1e, 0xb6, 0x2c, 0x54, 0x07,
	0x1c, 0x83, 0xaf, 0xde, 0x64, 0x41, 0x3a, 0xf4, 0xa2, 0x56, 0x1c, 0x8c, 0x4e, 0x93, 0x9d, 0x4e,
	0x3d, 0x3a, 0x59, 0x54, 0xdd, 0xc7, 0x8d, 0x58, 0x51, 0xc9, 0x1a, 0x14, 0xaf, 0x61, 0x3f, 0x4f,
	0xb5, 0x79, 0xd9, 0x28, 0x9e, 0x1a, 0xe2, 0x61, 0x3b, 0x72, 0xe2, 0x60, 0xd4, 0x4f, 0x6a, 0xe3,
	0x64, 0x67, 0x9c, 0x2c, 0x77, 0xc6, 0x2c, 0xa8, 0xf8, 0xa7, 0x1a, 0xef, 0x2f, 0x21, 0xf8, 0x31,
	0x15, 0x8f, 0xc1, 0x2f, 0x6c, 0xd5, 0xe8, 0x37, 0x09, 0x2f, 0xc1, 0x33, 0x5b, 0x45, 0x76, 0x85,
	0xee, 0xe8, 0xe4, 0x4f, 0xb1, 0xe5, 0x56, 0x11, 0xb3, 0xd8, 0xe0, 0x02, 0x3a, 0xdf, 0x4f, 0x08,
	0xe0, 0xcf, 0xe6, 0x93, 0xbb, 0x87, 0x69, 0xef, 0x1f, 0xee, 0x81, 0x77, 0x3b, 0x9f, 0x4d, 0x7b,
	0x4e, 0x55, 0x3d, 0xcf, 0xd9, 0x7d, 0xcf, 0x1d, 0x5c, 0x41, 0x70, 0xc3, 0x79, 0x49, 0x5a, 0x4f,
	0xa4, 0x7c, 0xc7, 0x73, 0xf0, 0x15, 0x49, 0x95, 0x57, 0x17, 0xac, 0xae, 0x70, 0xf8, 0xeb, 0x33,
	0xd6, 0xb4, 0x57, 0xbe, 0xdd, 0x6d, 0xfc, 0x15, 0x00, 0x00, 0xff, 0xff, 0x99, 0x32, 0x54, 0xa1,
	0xb5, 0x01, 0x00, 0x00,
}
