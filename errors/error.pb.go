// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: errors/error.proto

package errors

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

type Error int32

const (
	// 该错误比较含糊，只有当未知异常时才能使用，大部分情况应使用更具体的错误码
	Error_INTERNAL_SERVER_ERROR Error = 0
	Error_INVALID_PARAM         Error = 1
	// 摄像机相关
	Error_EMPTY_STREAMING_PROTOCOL Error = 20
	Error_INVALID_RTSP_ADDR        Error = 21
	Error_STREAM_INFO_NOTFOUND     Error = 22
	Error_CAMERA_NOT_FOUND         Error = 23
	Error_CAMERA_ALREADY_EXIST     Error = 24
	Error_CAMERA_UPDATE_ERROR      Error = 25
	// 任务相关
	Error_TASK_STOP_ERROR Error = 40
	// 用户相关
	Error_USER_NAME_EXIST              Error = 60
	Error_USER_NOT_FOUND               Error = 61
	Error_USER_NAME_OR_PASSWORD_ERROR  Error = 62
	Error_USER_CAPTCHA_VERIFY_ERROR    Error = 63
	Error_USER_CAPTCHA_EXPIRED         Error = 64
	Error_USER_REFRESH_TOKEN_NOT_FOUND Error = 65
	// 订阅告警相关
	Error_SUBSCRIBE_NOT_FOUND Error = 70
	// 组织架构相关
	Error_ORGANIZATION_NOT_FOUND Error = 90
)

// Enum value maps for Error.
var (
	Error_name = map[int32]string{
		0:  "INTERNAL_SERVER_ERROR",
		1:  "INVALID_PARAM",
		20: "EMPTY_STREAMING_PROTOCOL",
		21: "INVALID_RTSP_ADDR",
		22: "STREAM_INFO_NOTFOUND",
		23: "CAMERA_NOT_FOUND",
		24: "CAMERA_ALREADY_EXIST",
		25: "CAMERA_UPDATE_ERROR",
		40: "TASK_STOP_ERROR",
		60: "USER_NAME_EXIST",
		61: "USER_NOT_FOUND",
		62: "USER_NAME_OR_PASSWORD_ERROR",
		63: "USER_CAPTCHA_VERIFY_ERROR",
		64: "USER_CAPTCHA_EXPIRED",
		65: "USER_REFRESH_TOKEN_NOT_FOUND",
		70: "SUBSCRIBE_NOT_FOUND",
		90: "ORGANIZATION_NOT_FOUND",
	}
	Error_value = map[string]int32{
		"INTERNAL_SERVER_ERROR":        0,
		"INVALID_PARAM":                1,
		"EMPTY_STREAMING_PROTOCOL":     20,
		"INVALID_RTSP_ADDR":            21,
		"STREAM_INFO_NOTFOUND":         22,
		"CAMERA_NOT_FOUND":             23,
		"CAMERA_ALREADY_EXIST":         24,
		"CAMERA_UPDATE_ERROR":          25,
		"TASK_STOP_ERROR":              40,
		"USER_NAME_EXIST":              60,
		"USER_NOT_FOUND":               61,
		"USER_NAME_OR_PASSWORD_ERROR":  62,
		"USER_CAPTCHA_VERIFY_ERROR":    63,
		"USER_CAPTCHA_EXPIRED":         64,
		"USER_REFRESH_TOKEN_NOT_FOUND": 65,
		"SUBSCRIBE_NOT_FOUND":          70,
		"ORGANIZATION_NOT_FOUND":       90,
	}
)

func (x Error) Enum() *Error {
	p := new(Error)
	*p = x
	return p
}

func (x Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_error_proto_enumTypes[0].Descriptor()
}

func (Error) Type() protoreflect.EnumType {
	return &file_errors_error_proto_enumTypes[0]
}

func (x Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error.Descriptor instead.
func (Error) EnumDescriptor() ([]byte, []int) {
	return file_errors_error_proto_rawDescGZIP(), []int{0}
}

var File_errors_error_proto protoreflect.FileDescriptor

var file_errors_error_proto_rawDesc = []byte{
	0x0a, 0x12, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x69, 0x61, 0x73, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xa8, 0x04, 0x0a, 0x05, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x15, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x1a, 0x04,
	0xa8, 0x45, 0xf4, 0x03, 0x12, 0x17, 0x0a, 0x0d, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x50, 0x41, 0x52, 0x41, 0x4d, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x22, 0x0a,
	0x18, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x49, 0x4e, 0x47,
	0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x10, 0x14, 0x1a, 0x04, 0xa8, 0x45, 0x90,
	0x03, 0x12, 0x1b, 0x0a, 0x11, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x52, 0x54, 0x53,
	0x50, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x10, 0x15, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1e,
	0x0a, 0x14, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x4e, 0x4f,
	0x54, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x16, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x1a,
	0x0a, 0x10, 0x43, 0x41, 0x4d, 0x45, 0x52, 0x41, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55,
	0x4e, 0x44, 0x10, 0x17, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x1e, 0x0a, 0x14, 0x43, 0x41,
	0x4d, 0x45, 0x52, 0x41, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49,
	0x53, 0x54, 0x10, 0x18, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1d, 0x0a, 0x13, 0x43, 0x41,
	0x4d, 0x45, 0x52, 0x41, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x19, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x54, 0x41, 0x53,
	0x4b, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x28, 0x1a, 0x04,
	0xa8, 0x45, 0x90, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x41, 0x4d,
	0x45, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x3c, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12,
	0x18, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0x3d, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x25, 0x0a, 0x1b, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x52, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f,
	0x52, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x3e, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03,
	0x12, 0x23, 0x0a, 0x19, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x43, 0x41, 0x50, 0x54, 0x43, 0x48, 0x41,
	0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x3f, 0x1a,
	0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1e, 0x0a, 0x14, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x43, 0x41,
	0x50, 0x54, 0x43, 0x48, 0x41, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x40, 0x1a,
	0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x26, 0x0a, 0x1c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x52, 0x45,
	0x46, 0x52, 0x45, 0x53, 0x48, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x41, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1d, 0x0a,
	0x13, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46,
	0x4f, 0x55, 0x4e, 0x44, 0x10, 0x46, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x20, 0x0a, 0x16,
	0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x5a, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x1a, 0x04,
	0xa0, 0x45, 0xf4, 0x03, 0x42, 0x31, 0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x63, 0x74, 0x79, 0x75, 0x6e, 0x63, 0x64, 0x6e, 0x2e, 0x63, 0x6e, 0x2f, 0x69, 0x61, 0x73,
	0x2f, 0x69, 0x61, 0x73, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errors_error_proto_rawDescOnce sync.Once
	file_errors_error_proto_rawDescData = file_errors_error_proto_rawDesc
)

func file_errors_error_proto_rawDescGZIP() []byte {
	file_errors_error_proto_rawDescOnce.Do(func() {
		file_errors_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_error_proto_rawDescData)
	})
	return file_errors_error_proto_rawDescData
}

var file_errors_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_error_proto_goTypes = []interface{}{
	(Error)(0), // 0: ias_core.error.Error
}
var file_errors_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_error_proto_init() }
func file_errors_error_proto_init() {
	if File_errors_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_errors_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_error_proto_goTypes,
		DependencyIndexes: file_errors_error_proto_depIdxs,
		EnumInfos:         file_errors_error_proto_enumTypes,
	}.Build()
	File_errors_error_proto = out.File
	file_errors_error_proto_rawDesc = nil
	file_errors_error_proto_goTypes = nil
	file_errors_error_proto_depIdxs = nil
}