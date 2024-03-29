// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: params.proto

package params

import (
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

//Define response body
type RequestParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Param        string `protobuf:"bytes,1,opt,name=param,proto3" json:"param,omitempty"`
	AuthToken    string `protobuf:"bytes,2,opt,name=authToken,proto3" json:"authToken,omitempty"`
	CustomValues string `protobuf:"bytes,3,opt,name=customValues,proto3" json:"customValues,omitempty"`
}

func (x *RequestParam) Reset() {
	*x = RequestParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_params_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestParam) ProtoMessage() {}

func (x *RequestParam) ProtoReflect() protoreflect.Message {
	mi := &file_params_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestParam.ProtoReflect.Descriptor instead.
func (*RequestParam) Descriptor() ([]byte, []int) {
	return file_params_proto_rawDescGZIP(), []int{0}
}

func (x *RequestParam) GetParam() string {
	if x != nil {
		return x.Param
	}
	return ""
}

func (x *RequestParam) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

func (x *RequestParam) GetCustomValues() string {
	if x != nil {
		return x.CustomValues
	}
	return ""
}

type ResponseParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId           string `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	ParamValue          string `protobuf:"bytes,2,opt,name=paramValue,proto3" json:"paramValue,omitempty"`
	FromAppliedPolicyId string `protobuf:"bytes,3,opt,name=fromAppliedPolicyId,proto3" json:"fromAppliedPolicyId,omitempty"`
}

func (x *ResponseParam) Reset() {
	*x = ResponseParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_params_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseParam) ProtoMessage() {}

func (x *ResponseParam) ProtoReflect() protoreflect.Message {
	mi := &file_params_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseParam.ProtoReflect.Descriptor instead.
func (*ResponseParam) Descriptor() ([]byte, []int) {
	return file_params_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseParam) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ResponseParam) GetParamValue() string {
	if x != nil {
		return x.ParamValue
	}
	return ""
}

func (x *ResponseParam) GetFromAppliedPolicyId() string {
	if x != nil {
		return x.FromAppliedPolicyId
	}
	return ""
}

type RequestParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Param string `protobuf:"bytes,1,opt,name=param,proto3" json:"param,omitempty"`
}

func (x *RequestParams) Reset() {
	*x = RequestParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_params_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestParams) ProtoMessage() {}

func (x *RequestParams) ProtoReflect() protoreflect.Message {
	mi := &file_params_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestParams.ProtoReflect.Descriptor instead.
func (*RequestParams) Descriptor() ([]byte, []int) {
	return file_params_proto_rawDescGZIP(), []int{2}
}

func (x *RequestParams) GetParam() string {
	if x != nil {
		return x.Param
	}
	return ""
}

type ResponseParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId           int32  `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	ParamValue          string `protobuf:"bytes,2,opt,name=paramValue,proto3" json:"paramValue,omitempty"`
	FromAppliedPolicyId string `protobuf:"bytes,3,opt,name=fromAppliedPolicyId,proto3" json:"fromAppliedPolicyId,omitempty"`
}

func (x *ResponseParams) Reset() {
	*x = ResponseParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_params_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseParams) ProtoMessage() {}

func (x *ResponseParams) ProtoReflect() protoreflect.Message {
	mi := &file_params_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseParams.ProtoReflect.Descriptor instead.
func (*ResponseParams) Descriptor() ([]byte, []int) {
	return file_params_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseParams) GetRequestId() int32 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *ResponseParams) GetParamValue() string {
	if x != nil {
		return x.ParamValue
	}
	return ""
}

func (x *ResponseParams) GetFromAppliedPolicyId() string {
	if x != nil {
		return x.FromAppliedPolicyId
	}
	return ""
}

var File_params_proto protoreflect.FileDescriptor

var file_params_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x66, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x7f,
	0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12,
	0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x30, 0x0a,
	0x13, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x66, 0x72, 0x6f, 0x6d,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x22,
	0x25, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x80, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x66, 0x72, 0x6f, 0x6d, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x32, 0x48, 0x0a, 0x0b, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x52, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x08, 0x67, 0x65, 0x74, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x12, 0x14, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x15, 0x2e, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x6c, 0x79, 0x61, 0x74, 0x62, 0x6e, 0x2f, 0x6b, 0x65, 0x79, 0x6d, 0x76, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_params_proto_rawDescOnce sync.Once
	file_params_proto_rawDescData = file_params_proto_rawDesc
)

func file_params_proto_rawDescGZIP() []byte {
	file_params_proto_rawDescOnce.Do(func() {
		file_params_proto_rawDescData = protoimpl.X.CompressGZIP(file_params_proto_rawDescData)
	})
	return file_params_proto_rawDescData
}

var file_params_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_params_proto_goTypes = []interface{}{
	(*RequestParam)(nil),   // 0: params.RequestParam
	(*ResponseParam)(nil),  // 1: params.ResponseParam
	(*RequestParams)(nil),  // 2: params.RequestParams
	(*ResponseParams)(nil), // 3: params.ResponseParams
}
var file_params_proto_depIdxs = []int32{
	0, // 0: params.ParamReader.getParam:input_type -> params.RequestParam
	1, // 1: params.ParamReader.getParam:output_type -> params.ResponseParam
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_params_proto_init() }
func file_params_proto_init() {
	if File_params_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_params_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestParam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_params_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseParam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_params_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_params_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_params_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_params_proto_goTypes,
		DependencyIndexes: file_params_proto_depIdxs,
		MessageInfos:      file_params_proto_msgTypes,
	}.Build()
	File_params_proto = out.File
	file_params_proto_rawDesc = nil
	file_params_proto_goTypes = nil
	file_params_proto_depIdxs = nil
}
