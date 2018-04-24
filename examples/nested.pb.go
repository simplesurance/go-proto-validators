// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/nested.proto

/*
Package validator_examples is a generated protocol buffer package.

It is generated from these files:
	examples/nested.proto

It has these top-level messages:
	InnerMessage
	OuterMessage
*/
package validator_examples

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/simplesurance/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type InnerMessage struct {
	// some_integer can only be in range (1, 100).
	SomeInteger int32 `protobuf:"varint,1,opt,name=some_integer,json=someInteger" json:"some_integer,omitempty"`
}

func (m *InnerMessage) Reset()                    { *m = InnerMessage{} }
func (m *InnerMessage) String() string            { return proto.CompactTextString(m) }
func (*InnerMessage) ProtoMessage()               {}
func (*InnerMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InnerMessage) GetSomeInteger() int32 {
	if m != nil {
		return m.SomeInteger
	}
	return 0
}

type OuterMessage struct {
	// important_string must be a lowercase alpha-numeric of 5 to 30 characters (RE2 syntax).
	ImportantString string `protobuf:"bytes,1,opt,name=important_string,json=importantString" json:"important_string,omitempty"`
	// proto3 doesn't have `required`, the `msg_exist` enforces presence of InnerMessage.
	Inner *InnerMessage `protobuf:"bytes,2,opt,name=inner" json:"inner,omitempty"`
	// user_id must be valid uuid4
	UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// user_id must be valid uuid version 1-5
	UserIdd string `protobuf:"bytes,4,opt,name=user_idd,json=userIdd" json:"user_idd,omitempty"`
}

func (m *OuterMessage) Reset()                    { *m = OuterMessage{} }
func (m *OuterMessage) String() string            { return proto.CompactTextString(m) }
func (*OuterMessage) ProtoMessage()               {}
func (*OuterMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OuterMessage) GetImportantString() string {
	if m != nil {
		return m.ImportantString
	}
	return ""
}

func (m *OuterMessage) GetInner() *InnerMessage {
	if m != nil {
		return m.Inner
	}
	return nil
}

func (m *OuterMessage) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *OuterMessage) GetUserIdd() string {
	if m != nil {
		return m.UserIdd
	}
	return ""
}

func init() {
	proto.RegisterType((*InnerMessage)(nil), "validator.examples.InnerMessage")
	proto.RegisterType((*OuterMessage)(nil), "validator.examples.OuterMessage")
}

func init() { proto.RegisterFile("examples/nested.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x4d, 0x4b, 0xc3, 0x30,
	0x18, 0xc7, 0x97, 0xbd, 0x9b, 0x15, 0x1d, 0x01, 0xa1, 0x78, 0x69, 0x19, 0x1e, 0x06, 0xba, 0x16,
	0x26, 0x9e, 0x7c, 0x39, 0xf4, 0xd6, 0x83, 0x08, 0xf5, 0x28, 0x5a, 0xb2, 0xe5, 0xa1, 0x06, 0xd6,
	0xa4, 0x24, 0xa9, 0x88, 0xe2, 0xdd, 0xef, 0xe7, 0x07, 0x28, 0xf4, 0x93, 0x48, 0x53, 0xdd, 0x06,
	0x5e, 0x9f, 0xff, 0xef, 0xf7, 0xe4, 0x9f, 0x07, 0x1f, 0xc3, 0x1b, 0xcd, 0x8b, 0x0d, 0xe8, 0x50,
	0x80, 0x36, 0xc0, 0x82, 0x42, 0x49, 0x23, 0x09, 0x79, 0xa5, 0x1b, 0xce, 0xa8, 0x91, 0x2a, 0xf8,
	0x03, 0x4e, 0xae, 0x33, 0x6e, 0x5e, 0xca, 0x55, 0xb0, 0x96, 0x79, 0xa8, 0xb9, 0x9d, 0x95, 0x8a,
	0x8a, 0x35, 0x84, 0x99, 0x5c, 0x58, 0x6d, 0xb1, 0xb5, 0x74, 0xb8, 0x5b, 0x60, 0xa3, 0xd9, 0x15,
	0x76, 0x62, 0x21, 0x40, 0xdd, 0x81, 0xd6, 0x34, 0x03, 0x72, 0x86, 0x1d, 0x2d, 0x73, 0x48, 0xb9,
	0x30, 0x90, 0x81, 0x72, 0x91, 0x8f, 0xe6, 0x83, 0x68, 0x5c, 0x57, 0x5e, 0x7f, 0xda, 0x71, 0x59,
	0x32, 0x69, 0xd2, 0xb8, 0x0d, 0x67, 0xdf, 0x08, 0x3b, 0xf7, 0xa5, 0xd9, 0xd9, 0x37, 0x78, 0xca,
	0xf3, 0x42, 0x2a, 0x43, 0x85, 0x49, 0xb5, 0x51, 0x5c, 0x64, 0x76, 0xc3, 0x41, 0x44, 0xea, 0xca,
	0x3b, 0xc4, 0xce, 0xf3, 0x23, 0x5d, 0xbc, 0x3f, 0x7d, 0x2c, 0xcf, 0x2f, 0x3f, 0x4f, 0x93, 0xa3,
	0x2d, 0xfb, 0x60, 0x51, 0x72, 0x8b, 0x07, 0xbc, 0x29, 0xe3, 0x76, 0x7d, 0x34, 0x9f, 0x2c, 0xfd,
	0xe0, 0xff, 0x77, 0x83, 0xfd, 0xb6, 0xd1, 0xb0, 0xae, 0xbc, 0xae, 0x8f, 0x92, 0x56, 0x23, 0x3e,
	0x1e, 0x95, 0x1a, 0x54, 0xca, 0x99, 0xdb, 0xb3, 0xaf, 0x8e, 0xea, 0xca, 0xeb, 0x7d, 0xa1, 0x7e,
	0x32, 0x6c, 0xe6, 0x31, 0x23, 0x33, 0x3c, 0xfe, 0x25, 0x98, 0xdb, 0xdf, 0x47, 0x3a, 0xc9, 0xa8,
	0x45, 0xd8, 0x6a, 0x68, 0x2f, 0x73, 0xf1, 0x13, 0x00, 0x00, 0xff, 0xff, 0x84, 0xa6, 0x42, 0xa8,
	0x84, 0x01, 0x00, 0x00,
}
