// Code generated by protoc-gen-go.
// source: examples/nested.proto
// DO NOT EDIT!

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
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type InnerMessage struct {
	// some_integer can only be in range (1, 100).
	SomeInteger int32 `protobuf:"varint,1,opt,name=some_integer,json=someInteger" json:"some_integer,omitempty"`
}

func (m *InnerMessage) Reset()                    { *m = InnerMessage{} }
func (m *InnerMessage) String() string            { return proto.CompactTextString(m) }
func (*InnerMessage) ProtoMessage()               {}
func (*InnerMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type OuterMessage struct {
	// important_string must be a lowercase alpha-numeric of 5 to 30 characters (RE2 syntax).
	ImportantString string `protobuf:"bytes,1,opt,name=important_string,json=importantString" json:"important_string,omitempty"`
	// proto3 doesn't have `required`, the `msg_exist` enforces presence of InnerMessage.
	Inner *InnerMessage `protobuf:"bytes,2,opt,name=inner" json:"inner,omitempty"`
}

func (m *OuterMessage) Reset()                    { *m = OuterMessage{} }
func (m *OuterMessage) String() string            { return proto.CompactTextString(m) }
func (*OuterMessage) ProtoMessage()               {}
func (*OuterMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OuterMessage) GetInner() *InnerMessage {
	if m != nil {
		return m.Inner
	}
	return nil
}

func init() {
	proto.RegisterType((*InnerMessage)(nil), "validator.examples.InnerMessage")
	proto.RegisterType((*OuterMessage)(nil), "validator.examples.OuterMessage")
}

var fileDescriptor0 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0xcf, 0x4b, 0x2d, 0x2e, 0x49, 0x4d, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x2a, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x83, 0x29, 0x90,
	0x32, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x2d, 0xcf, 0x2c,
	0xc9, 0xce, 0x2f, 0xd7, 0x4f, 0xcf, 0xd7, 0x05, 0x6b, 0xd0, 0x85, 0xab, 0x2f, 0xd6, 0x47, 0x68,
	0x05, 0x4b, 0x29, 0x59, 0x73, 0xf1, 0x78, 0xe6, 0xe5, 0xa5, 0x16, 0xf9, 0xa6, 0x16, 0x17, 0x27,
	0xa6, 0xa7, 0x0a, 0x69, 0x73, 0xf1, 0x14, 0xe7, 0xe7, 0xa6, 0xc6, 0x67, 0xe6, 0x95, 0xa4, 0xa6,
	0xa7, 0x16, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x3a, 0x71, 0x3c, 0xba, 0x2f, 0xcf, 0x22, 0xc0,
	0x20, 0x91, 0x12, 0xc4, 0x0d, 0x92, 0xf5, 0x84, 0x48, 0x2a, 0xf5, 0x32, 0x72, 0xf1, 0xf8, 0x97,
	0x96, 0x20, 0x74, 0xdb, 0x72, 0x09, 0x64, 0xe6, 0x16, 0xe4, 0x17, 0x95, 0x24, 0xe6, 0x95, 0xc4,
	0x17, 0x97, 0x14, 0x65, 0xe6, 0xa5, 0x83, 0x4d, 0xe0, 0x74, 0x12, 0x02, 0x9a, 0xc0, 0xc7, 0xc5,
	0x13, 0x17, 0x9d, 0xa8, 0x5b, 0x15, 0x5b, 0x6d, 0xa4, 0x63, 0x5a, 0xab, 0x12, 0xc4, 0x0f, 0x57,
	0x1b, 0x0c, 0x56, 0x2a, 0x64, 0xc7, 0xc5, 0x9a, 0x09, 0x72, 0x8c, 0x04, 0x13, 0x50, 0x0f, 0xb7,
	0x91, 0x82, 0x1e, 0xa6, 0x47, 0xf5, 0x90, 0x5d, 0xeb, 0xc4, 0x06, 0x34, 0x15, 0xa8, 0x36, 0x08,
	0xa2, 0x2d, 0x89, 0x0d, 0xec, 0x27, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x08, 0x68, 0xd9,
	0xd5, 0x38, 0x01, 0x00, 0x00,
}
