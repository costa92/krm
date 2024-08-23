// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0--rc1
// source: usercenter/v1/errors.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorReason int32

const (
	// 用户登录失败，身份验证未通过
	ErrorReason_UserLoginFailed ErrorReason = 0
	// 用户已存在，无法创建用户
	ErrorReason_UserAlreadyExists ErrorReason = 1
	// 用户未找到，可能是用户不存在或输入的用户标识有误
	ErrorReason_UserNotFound ErrorReason = 2
	// 创建用户失败，可能是由于服务器或其他问题导致的创建过程中的错误
	ErrorReason_UserCreateFailed ErrorReason = 3
	// 用户操作被禁止，可能是由于权限不足或其他安全限制导致的
	ErrorReason_UserOperationForbidden ErrorReason = 4
	// 密钥达到最大数量限制，无法继续创建新密钥
	ErrorReason_SecretReachMaxCount ErrorReason = 5
	// 密钥未找到，可能是由于密钥不存在或输入的密钥标识有误
	ErrorReason_SecretNotFound ErrorReason = 6
	// 创建密钥失败，可能是由于服务器或其他问题导致的创建过程中的错误
	ErrorReason_SecretCreateFailed ErrorReason = 7
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "UserLoginFailed",
		1: "UserAlreadyExists",
		2: "UserNotFound",
		3: "UserCreateFailed",
		4: "UserOperationForbidden",
		5: "SecretReachMaxCount",
		6: "SecretNotFound",
		7: "SecretCreateFailed",
	}
	ErrorReason_value = map[string]int32{
		"UserLoginFailed":        0,
		"UserAlreadyExists":      1,
		"UserNotFound":           2,
		"UserCreateFailed":       3,
		"UserOperationForbidden": 4,
		"SecretReachMaxCount":    5,
		"SecretNotFound":         6,
		"SecretCreateFailed":     7,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_usercenter_v1_errors_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_usercenter_v1_errors_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_usercenter_v1_errors_proto_rawDescGZIP(), []int{0}
}

var File_usercenter_v1_errors_proto protoreflect.FileDescriptor

var file_usercenter_v1_errors_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x75, 0x73,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2a, 0xf8, 0x01, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x19, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x46, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x1b, 0x0a, 0x11, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73,
	0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x99, 0x03, 0x12, 0x16, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03,
	0x12, 0x1a, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x10, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x9d, 0x04, 0x12, 0x20, 0x0a, 0x16,
	0x55, 0x73, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72,
	0x62, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x10, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0x93, 0x03, 0x12, 0x1d,
	0x0a, 0x13, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x61, 0x63, 0x68, 0x4d, 0x61, 0x78,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0x05, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x18, 0x0a,
	0x0e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10,
	0x06, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x1c, 0x0a, 0x12, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x07, 0x1a,
	0x04, 0xa8, 0x45, 0x9d, 0x04, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x31, 0x5a, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x61, 0x39,
	0x32, 0x2f, 0x6b, 0x72, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_usercenter_v1_errors_proto_rawDescOnce sync.Once
	file_usercenter_v1_errors_proto_rawDescData = file_usercenter_v1_errors_proto_rawDesc
)

func file_usercenter_v1_errors_proto_rawDescGZIP() []byte {
	file_usercenter_v1_errors_proto_rawDescOnce.Do(func() {
		file_usercenter_v1_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_usercenter_v1_errors_proto_rawDescData)
	})
	return file_usercenter_v1_errors_proto_rawDescData
}

var file_usercenter_v1_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_usercenter_v1_errors_proto_goTypes = []any{
	(ErrorReason)(0), // 0: usercenter.v1.ErrorReason
}
var file_usercenter_v1_errors_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_usercenter_v1_errors_proto_init() }
func file_usercenter_v1_errors_proto_init() {
	if File_usercenter_v1_errors_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_usercenter_v1_errors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_usercenter_v1_errors_proto_goTypes,
		DependencyIndexes: file_usercenter_v1_errors_proto_depIdxs,
		EnumInfos:         file_usercenter_v1_errors_proto_enumTypes,
	}.Build()
	File_usercenter_v1_errors_proto = out.File
	file_usercenter_v1_errors_proto_rawDesc = nil
	file_usercenter_v1_errors_proto_goTypes = nil
	file_usercenter_v1_errors_proto_depIdxs = nil
}
