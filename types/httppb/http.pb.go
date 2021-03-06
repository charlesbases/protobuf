// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        v3.20.1
// source: google/protobuf/plugin/http.proto

package httppb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Http struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Pattern:
	//	*Http_Get
	//	*Http_Put
	//	*Http_Post
	//	*Http_Delete
	Pattern isHttp_Pattern `protobuf_oneof:"pattern"`
	// content-type for request
	Consume string `protobuf:"bytes,5,opt,name=consume,proto3" json:"consume,omitempty"`
	// content-type for response
	Produce string `protobuf:"bytes,6,opt,name=produce,proto3" json:"produce,omitempty"`
}

func (x *Http) Reset() {
	*x = Http{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_protobuf_plugin_http_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http) ProtoMessage() {}

func (x *Http) ProtoReflect() protoreflect.Message {
	mi := &file_google_protobuf_plugin_http_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http.ProtoReflect.Descriptor instead.
func (*Http) Descriptor() ([]byte, []int) {
	return file_google_protobuf_plugin_http_proto_rawDescGZIP(), []int{0}
}

func (m *Http) GetPattern() isHttp_Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (x *Http) GetGet() string {
	if x, ok := x.GetPattern().(*Http_Get); ok {
		return x.Get
	}
	return ""
}

func (x *Http) GetPut() string {
	if x, ok := x.GetPattern().(*Http_Put); ok {
		return x.Put
	}
	return ""
}

func (x *Http) GetPost() string {
	if x, ok := x.GetPattern().(*Http_Post); ok {
		return x.Post
	}
	return ""
}

func (x *Http) GetDelete() string {
	if x, ok := x.GetPattern().(*Http_Delete); ok {
		return x.Delete
	}
	return ""
}

func (x *Http) GetConsume() string {
	if x != nil {
		return x.Consume
	}
	return ""
}

func (x *Http) GetProduce() string {
	if x != nil {
		return x.Produce
	}
	return ""
}

type isHttp_Pattern interface {
	isHttp_Pattern()
}

type Http_Get struct {
	Get string `protobuf:"bytes,1,opt,name=get,proto3,oneof"`
}

type Http_Put struct {
	Put string `protobuf:"bytes,2,opt,name=put,proto3,oneof"`
}

type Http_Post struct {
	Post string `protobuf:"bytes,3,opt,name=post,proto3,oneof"`
}

type Http_Delete struct {
	Delete string `protobuf:"bytes,4,opt,name=delete,proto3,oneof"`
}

func (*Http_Get) isHttp_Pattern() {}

func (*Http_Put) isHttp_Pattern() {}

func (*Http_Post) isHttp_Pattern() {}

func (*Http_Delete) isHttp_Pattern() {}

var file_google_protobuf_plugin_http_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Http)(nil),
		Field:         199522,
		Name:          "google.protobuf.plugin.http",
		Tag:           "bytes,199522,opt,name=http",
		Filename:      "google/protobuf/plugin/http.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional google.protobuf.plugin.Http http = 199522;
	E_Http = &file_google_protobuf_plugin_http_proto_extTypes[0]
)

var File_google_protobuf_plugin_http_proto protoreflect.FileDescriptor

var file_google_protobuf_plugin_http_proto_rawDesc = []byte{
	0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01,
	0x0a, 0x04, 0x48, 0x74, 0x74, 0x70, 0x12, 0x12, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x67, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x03, 0x70, 0x75,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x70, 0x75, 0x74, 0x12, 0x14,
	0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04,
	0x70, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x3a, 0x52, 0x0a,
	0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe2, 0x96, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x04, 0x68, 0x74, 0x74,
	0x70, 0x42, 0x69, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x42,
	0x09, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x6c, 0x65, 0x73,
	0x62, 0x61, 0x73, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x70, 0x62, 0x3b, 0x68, 0x74, 0x74, 0x70,
	0x70, 0x62, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x04, 0x47, 0x41, 0x50, 0x49, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_protobuf_plugin_http_proto_rawDescOnce sync.Once
	file_google_protobuf_plugin_http_proto_rawDescData = file_google_protobuf_plugin_http_proto_rawDesc
)

func file_google_protobuf_plugin_http_proto_rawDescGZIP() []byte {
	file_google_protobuf_plugin_http_proto_rawDescOnce.Do(func() {
		file_google_protobuf_plugin_http_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_protobuf_plugin_http_proto_rawDescData)
	})
	return file_google_protobuf_plugin_http_proto_rawDescData
}

var file_google_protobuf_plugin_http_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_protobuf_plugin_http_proto_goTypes = []interface{}{
	(*Http)(nil),                       // 0: google.protobuf.plugin.Http
	(*descriptorpb.MethodOptions)(nil), // 1: google.protobuf.MethodOptions
}
var file_google_protobuf_plugin_http_proto_depIdxs = []int32{
	1, // 0: google.protobuf.plugin.http:extendee -> google.protobuf.MethodOptions
	0, // 1: google.protobuf.plugin.http:type_name -> google.protobuf.plugin.Http
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_protobuf_plugin_http_proto_init() }
func file_google_protobuf_plugin_http_proto_init() {
	if File_google_protobuf_plugin_http_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_protobuf_plugin_http_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http); i {
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
	file_google_protobuf_plugin_http_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Http_Get)(nil),
		(*Http_Put)(nil),
		(*Http_Post)(nil),
		(*Http_Delete)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_protobuf_plugin_http_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_google_protobuf_plugin_http_proto_goTypes,
		DependencyIndexes: file_google_protobuf_plugin_http_proto_depIdxs,
		MessageInfos:      file_google_protobuf_plugin_http_proto_msgTypes,
		ExtensionInfos:    file_google_protobuf_plugin_http_proto_extTypes,
	}.Build()
	File_google_protobuf_plugin_http_proto = out.File
	file_google_protobuf_plugin_http_proto_rawDesc = nil
	file_google_protobuf_plugin_http_proto_goTypes = nil
	file_google_protobuf_plugin_http_proto_depIdxs = nil
}
