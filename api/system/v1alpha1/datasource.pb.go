// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: system/v1alpha1/datasource.proto

package v1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// DataSource defines the source of bytes to use.
type DataSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*DataSource_Secret
	//	*DataSource_File
	//	*DataSource_Inline
	//	*DataSource_InlineString
	Type isDataSource_Type `protobuf_oneof:"type"`
}

func (x *DataSource) Reset() {
	*x = DataSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1alpha1_datasource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSource) ProtoMessage() {}

func (x *DataSource) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1alpha1_datasource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSource.ProtoReflect.Descriptor instead.
func (*DataSource) Descriptor() ([]byte, []int) {
	return file_system_v1alpha1_datasource_proto_rawDescGZIP(), []int{0}
}

func (m *DataSource) GetType() isDataSource_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *DataSource) GetSecret() string {
	if x, ok := x.GetType().(*DataSource_Secret); ok {
		return x.Secret
	}
	return ""
}

func (x *DataSource) GetFile() string {
	if x, ok := x.GetType().(*DataSource_File); ok {
		return x.File
	}
	return ""
}

func (x *DataSource) GetInline() *wrapperspb.BytesValue {
	if x, ok := x.GetType().(*DataSource_Inline); ok {
		return x.Inline
	}
	return nil
}

func (x *DataSource) GetInlineString() string {
	if x, ok := x.GetType().(*DataSource_InlineString); ok {
		return x.InlineString
	}
	return ""
}

type isDataSource_Type interface {
	isDataSource_Type()
}

type DataSource_Secret struct {
	// Data source is a secret with given Secret key.
	Secret string `protobuf:"bytes,1,opt,name=secret,proto3,oneof"`
}

type DataSource_File struct {
	// Data source is a path to a file.
	File string `protobuf:"bytes,2,opt,name=file,proto3,oneof"`
}

type DataSource_Inline struct {
	// Data source is inline bytes.
	Inline *wrapperspb.BytesValue `protobuf:"bytes,3,opt,name=inline,proto3,oneof"`
}

type DataSource_InlineString struct {
	// Data source is inline string
	InlineString string `protobuf:"bytes,4,opt,name=inlineString,proto3,oneof"`
}

func (*DataSource_Secret) isDataSource_Type() {}

func (*DataSource_File) isDataSource_Type() {}

func (*DataSource_Inline) isDataSource_Type() {}

func (*DataSource_InlineString) isDataSource_Type() {}

var File_system_v1alpha1_datasource_proto protoreflect.FileDescriptor

var file_system_v1alpha1_datasource_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x0a, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x69, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x06, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x24,
	0x0a, 0x0c, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x2c, 0x5a, 0x2a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x6d, 0x61, 0x68,
	0x71, 0x2f, 0x6b, 0x75, 0x6d, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_system_v1alpha1_datasource_proto_rawDescOnce sync.Once
	file_system_v1alpha1_datasource_proto_rawDescData = file_system_v1alpha1_datasource_proto_rawDesc
)

func file_system_v1alpha1_datasource_proto_rawDescGZIP() []byte {
	file_system_v1alpha1_datasource_proto_rawDescOnce.Do(func() {
		file_system_v1alpha1_datasource_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_v1alpha1_datasource_proto_rawDescData)
	})
	return file_system_v1alpha1_datasource_proto_rawDescData
}

var file_system_v1alpha1_datasource_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_system_v1alpha1_datasource_proto_goTypes = []interface{}{
	(*DataSource)(nil),            // 0: kuma.system.v1alpha1.DataSource
	(*wrapperspb.BytesValue)(nil), // 1: google.protobuf.BytesValue
}
var file_system_v1alpha1_datasource_proto_depIdxs = []int32{
	1, // 0: kuma.system.v1alpha1.DataSource.inline:type_name -> google.protobuf.BytesValue
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_system_v1alpha1_datasource_proto_init() }
func file_system_v1alpha1_datasource_proto_init() {
	if File_system_v1alpha1_datasource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_system_v1alpha1_datasource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSource); i {
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
	file_system_v1alpha1_datasource_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DataSource_Secret)(nil),
		(*DataSource_File)(nil),
		(*DataSource_Inline)(nil),
		(*DataSource_InlineString)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_system_v1alpha1_datasource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_system_v1alpha1_datasource_proto_goTypes,
		DependencyIndexes: file_system_v1alpha1_datasource_proto_depIdxs,
		MessageInfos:      file_system_v1alpha1_datasource_proto_msgTypes,
	}.Build()
	File_system_v1alpha1_datasource_proto = out.File
	file_system_v1alpha1_datasource_proto_rawDesc = nil
	file_system_v1alpha1_datasource_proto_goTypes = nil
	file_system_v1alpha1_datasource_proto_depIdxs = nil
}
