// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/validate/validate.proto

package validate

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

var E_Disabled = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         53201,
	Name:          "validate.disabled",
	Tag:           "varint,53201,opt,name=disabled",
	Filename:      "pkg/validate/validate.proto",
}

var E_AdditionalHandler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         53202,
	Name:          "validate.additional_handler",
	Tag:           "bytes,53202,opt,name=additional_handler",
	Filename:      "pkg/validate/validate.proto",
}

var E_Required = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         53301,
	Name:          "validate.required",
	Tag:           "varint,53301,opt,name=required",
	Filename:      "pkg/validate/validate.proto",
}

var E_Skip = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         53302,
	Name:          "validate.skip",
	Tag:           "varint,53302,opt,name=skip",
	Filename:      "pkg/validate/validate.proto",
}

var E_MinLen = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*int64)(nil),
	Field:         53303,
	Name:          "validate.min_len",
	Tag:           "varint,53303,opt,name=min_len",
	Filename:      "pkg/validate/validate.proto",
}

var E_MaxLen = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*int64)(nil),
	Field:         53304,
	Name:          "validate.max_len",
	Tag:           "varint,53304,opt,name=max_len",
	Filename:      "pkg/validate/validate.proto",
}

var E_IsContactKey = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         53310,
	Name:          "validate.is_contact_key",
	Tag:           "varint,53310,opt,name=is_contact_key",
	Filename:      "pkg/validate/validate.proto",
}

var E_MinItems = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*int64)(nil),
	Field:         53305,
	Name:          "validate.min_items",
	Tag:           "varint,53305,opt,name=min_items",
	Filename:      "pkg/validate/validate.proto",
}

var E_MaxItems = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*int64)(nil),
	Field:         53306,
	Name:          "validate.max_items",
	Tag:           "varint,53306,opt,name=max_items",
	Filename:      "pkg/validate/validate.proto",
}

var E_IntGt = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*int64)(nil),
	Field:         53307,
	Name:          "validate.int_gt",
	Tag:           "varint,53307,opt,name=int_gt",
	Filename:      "pkg/validate/validate.proto",
}

var E_IntLt = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*int64)(nil),
	Field:         53308,
	Name:          "validate.int_lt",
	Tag:           "varint,53308,opt,name=int_lt",
	Filename:      "pkg/validate/validate.proto",
}

var E_DefinedOnly = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         53309,
	Name:          "validate.defined_only",
	Tag:           "varint,53309,opt,name=defined_only",
	Filename:      "pkg/validate/validate.proto",
}

func init() {
	proto.RegisterExtension(E_Disabled)
	proto.RegisterExtension(E_AdditionalHandler)
	proto.RegisterExtension(E_Required)
	proto.RegisterExtension(E_Skip)
	proto.RegisterExtension(E_MinLen)
	proto.RegisterExtension(E_MaxLen)
	proto.RegisterExtension(E_IsContactKey)
	proto.RegisterExtension(E_MinItems)
	proto.RegisterExtension(E_MaxItems)
	proto.RegisterExtension(E_IntGt)
	proto.RegisterExtension(E_IntLt)
	proto.RegisterExtension(E_DefinedOnly)
}

func init() { proto.RegisterFile("pkg/validate/validate.proto", fileDescriptor_536d18386cedcbe8) }

var fileDescriptor_536d18386cedcbe8 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd3, 0xcb, 0xaa, 0xd3, 0x40,
	0x1c, 0x06, 0x70, 0x43, 0xb5, 0xa6, 0x63, 0x11, 0xcc, 0x4a, 0xbc, 0xc4, 0x2e, 0x5d, 0x25, 0x0b,
	0x41, 0x24, 0xea, 0xa6, 0xe2, 0x0d, 0x2b, 0x95, 0x2e, 0xdd, 0x84, 0x69, 0xe6, 0xdf, 0xf4, 0x4f,
	0x27, 0x33, 0x71, 0x66, 0x2a, 0xc9, 0x9b, 0xb8, 0xd3, 0x97, 0xf0, 0x7e, 0x59, 0xbb, 0xd4, 0xf3,
	0x04, 0x87, 0x9e, 0x17, 0x39, 0xe4, 0x7a, 0x0a, 0x67, 0x31, 0xdd, 0xe5, 0xf2, 0xfd, 0x3e, 0x3e,
	0x06, 0x86, 0xdc, 0xcc, 0x37, 0x69, 0xf8, 0x9e, 0x72, 0x64, 0xd4, 0x40, 0xff, 0x10, 0xe4, 0x4a,
	0x1a, 0xe9, 0xb9, 0xdd, 0xfb, 0x8d, 0x49, 0x2a, 0x65, 0xca, 0x21, 0xac, 0xbf, 0x2f, 0xb7, 0xab,
	0x90, 0x81, 0x4e, 0x14, 0xe6, 0x46, 0xaa, 0x26, 0x1b, 0x3d, 0x26, 0x2e, 0x43, 0x4d, 0x97, 0x1c,
	0x98, 0x77, 0x27, 0x68, 0xe2, 0x41, 0x17, 0x0f, 0x5e, 0x83, 0xd6, 0x34, 0x85, 0x79, 0x6e, 0x50,
	0x0a, 0x7d, 0xfd, 0xff, 0xc7, 0xc1, 0xc4, 0xb9, 0xeb, 0x2e, 0x7a, 0x12, 0xbd, 0x21, 0x1e, 0x65,
	0x0c, 0xab, 0xdf, 0x94, 0xc7, 0x6b, 0x2a, 0x18, 0x07, 0x65, 0x2f, 0x3a, 0xaa, 0x8b, 0x46, 0x8b,
	0x6b, 0x67, 0xf8, 0x45, 0x63, 0xa3, 0x87, 0xc4, 0x55, 0xf0, 0x6e, 0x8b, 0x0a, 0x98, 0x77, 0xfb,
	0x5c, 0xcf, 0x33, 0x04, 0xce, 0xba, 0x96, 0xcf, 0x9f, 0xda, 0x39, 0x1d, 0x88, 0xee, 0x91, 0x8b,
	0x7a, 0x83, 0xb9, 0x0d, 0x7e, 0x69, 0x61, 0x1d, 0x8e, 0x1e, 0x90, 0xcb, 0x19, 0x8a, 0x98, 0x83,
	0xb0, 0xb9, 0xaf, 0xb5, 0x1b, 0x2c, 0x86, 0x19, 0x8a, 0x19, 0x88, 0x5a, 0xd2, 0xe2, 0x10, 0xf9,
	0xad, 0x97, 0xb4, 0xa8, 0xe4, 0x53, 0x72, 0x15, 0x75, 0x9c, 0x48, 0x61, 0x68, 0x62, 0xe2, 0x0d,
	0x94, 0xb6, 0x82, 0x3f, 0xed, 0xe4, 0x31, 0xea, 0x27, 0x8d, 0x7a, 0x05, 0x65, 0xf4, 0x88, 0x8c,
	0xaa, 0xe9, 0x68, 0x20, 0xd3, 0xb6, 0x86, 0xef, 0xed, 0x04, 0x37, 0x43, 0xf1, 0xb2, 0x02, 0xb5,
	0xa6, 0xc5, 0x61, 0xfa, 0x47, 0xaf, 0x69, 0xd1, 0xe8, 0xfb, 0x64, 0x88, 0xc2, 0xc4, 0xa9, 0xb1,
	0xd1, 0x9f, 0x2d, 0xbd, 0x84, 0xc2, 0x3c, 0x37, 0x9d, 0xe3, 0x56, 0xf7, 0x6b, 0xcf, 0xcd, 0x4c,
	0x34, 0x25, 0x63, 0x06, 0x2b, 0x14, 0xc0, 0x62, 0x29, 0xb8, 0xf5, 0xc0, 0x7e, 0xb7, 0x07, 0x76,
	0xa5, 0x45, 0x73, 0xc1, 0xcb, 0x69, 0xf0, 0x77, 0xe7, 0x3b, 0xff, 0x76, 0xbe, 0x73, 0xbc, 0xf3,
	0x9d, 0x0f, 0x27, 0xfe, 0x85, 0xb7, 0xb7, 0x96, 0xa0, 0x4c, 0x19, 0x18, 0x48, 0xd6, 0x61, 0x22,
	0x15, 0x84, 0xfb, 0x17, 0xeb, 0x34, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x27, 0xf8, 0x81, 0x67, 0x03,
	0x00, 0x00,
}
