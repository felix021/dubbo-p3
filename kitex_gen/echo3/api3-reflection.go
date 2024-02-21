// Code generated by thriftgo (0.3.6). DO NOT EDIT.

package echo3

import (
	"reflect"

	"github.com/cloudwego/thriftgo/thrift_reflection"
)

// IDL Name: api3
// IDL Path: api3.thrift

var file_api3_thrift_go_types = []interface{}{
	(*EchoRequest)(nil),  // Struct 0: echo3.EchoRequest
	(*EchoResponse)(nil), // Struct 1: echo3.EchoResponse
}
var file_api3_thrift *thrift_reflection.FileDescriptor
var file_idl_api3_rawDesc = []byte{
	0x1f, 0x8b, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xac, 0x92, 0xc1, 0x4e, 0xc3, 0x30,
	0x10, 0x44, 0x7, 0x27, 0x69, 0x2, 0x4e, 0x94, 0x23, 0x7c, 0x45, 0x2f, 0xfd, 0x4, 0xe0, 0xc2,
	0x81, 0x4a, 0xc0, 0xf, 0xb8, 0x61, 0xdb, 0xba, 0xb4, 0x75, 0x1a, 0xbb, 0x11, 0x9f, 0x8f, 0xd6,
	0x4d, 0x4a, 0x24, 0x48, 0x52, 0x24, 0x4e, 0x23, 0x79, 0x56, 0xb3, 0xeb, 0xa7, 0x91, 0xb8, 0x2,
	0x20, 0x55, 0xa9, 0x67, 0x53, 0xb7, 0xae, 0xf4, 0xd2, 0x65, 0x10, 0x52, 0x2, 0xfe, 0x3d, 0xdc,
	0xa8, 0x5a, 0xb1, 0xcf, 0x7a, 0xf6, 0x83, 0xb3, 0x2f, 0x56, 0x6, 0x40, 0x44, 0xc5, 0xda, 0xcc,
	0x72, 0x84, 0x29, 0x3f, 0xff, 0x4c, 0x94, 0x10, 0x0, 0xd2, 0x37, 0xb2, 0xee, 0x95, 0xaa, 0x5a,
	0x17, 0x34, 0xcb, 0x11, 0xc, 0xe, 0x87, 0x8f, 0xc5, 0xda, 0xa4, 0x8, 0x7a, 0xc3, 0xd8, 0x7f,
	0x21, 0x5b, 0x9a, 0xbd, 0x25, 0x8c, 0xad, 0xe, 0x2a, 0x3a, 0xc, 0x84, 0xc9, 0x53, 0xd8, 0xe1,
	0x48, 0xd6, 0x41, 0x22, 0x4, 0x10, 0x3f, 0xd0, 0x52, 0x1d, 0xb7, 0x2e, 0x41, 0xc4, 0xc1, 0x19,
	0x62, 0x99, 0x83, 0x47, 0x91, 0xb0, 0x20, 0x43, 0xd4, 0x3e, 0x4c, 0x58, 0x72, 0xc4, 0x7c, 0x1,
	0x4, 0xfb, 0x19, 0x42, 0x6f, 0xf2, 0xb6, 0xec, 0x49, 0xd5, 0xea, 0x7e, 0xab, 0xac, 0x7d, 0x56,
	0x3b, 0x6a, 0xb9, 0xdd, 0x7d, 0x68, 0x47, 0x9f, 0x53, 0xcf, 0x6d, 0xda, 0xe5, 0x22, 0xfd, 0x3e,
	0xce, 0x8b, 0x38, 0x4f, 0x5c, 0x72, 0x72, 0xc3, 0xb2, 0x77, 0x36, 0xde, 0x91, 0xb5, 0x6a, 0x45,
	0x3, 0x4, 0x26, 0xd6, 0x55, 0x7a, 0xbf, 0x6a, 0x3e, 0x9f, 0x70, 0xb0, 0xae, 0xe8, 0xbd, 0xef,
	0xf7, 0xbd, 0x9c, 0xcd, 0x62, 0x33, 0xc4, 0xd9, 0xb7, 0x68, 0xbe, 0xd8, 0x50, 0xd1, 0x72, 0x4e,
	0xe6, 0xa5, 0xd3, 0x66, 0xaf, 0xb6, 0xa7, 0x55, 0xe2, 0x17, 0xd0, 0x63, 0x2c, 0x6f, 0xbb, 0x2c,
	0x3b, 0x58, 0x5a, 0x94, 0x17, 0x55, 0x68, 0xac, 0x8f, 0xff, 0xcf, 0xf0, 0x8f, 0x25, 0xe9, 0x1e,
	0xfb, 0x5d, 0x92, 0x49, 0xda, 0x2d, 0x5f, 0x8e, 0xa4, 0xd1, 0xeb, 0x46, 0x6f, 0xbc, 0xe2, 0x2b,
	0x0, 0x0, 0xff, 0xff, 0x19, 0xf6, 0x2c, 0xc7, 0xe5, 0x3, 0x0, 0x0,
}

func init() {
	if file_api3_thrift != nil {
		return
	}
	type x struct{}
	builder := &thrift_reflection.FileDescriptorBuilder{
		Bytes:         file_idl_api3_rawDesc,
		GoTypes:       file_api3_thrift_go_types,
		GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
	}
	file_api3_thrift = thrift_reflection.BuildFileDescriptor(builder)
}

func GetFileDescriptorForApi3() *thrift_reflection.FileDescriptor {
	return file_api3_thrift
}
func (p *EchoRequest) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api3_thrift.GetStructDescriptor("EchoRequest")
}

func (p *EchoRequest) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api3_thrift.Filepath
	ret.Name = "EchoRequest"
	return ret
}
func (p *EchoResponse) GetDescriptor() *thrift_reflection.StructDescriptor {
	return file_api3_thrift.GetStructDescriptor("EchoResponse")
}

func (p *EchoResponse) GetTypeDescriptor() *thrift_reflection.TypeDescriptor {
	ret := thrift_reflection.NewTypeDescriptor()
	ret.Filepath = file_api3_thrift.Filepath
	ret.Name = "EchoResponse"
	return ret
}
