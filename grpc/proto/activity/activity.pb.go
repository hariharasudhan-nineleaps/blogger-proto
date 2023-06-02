// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: grpc/proto/activity/activity.proto

package activity

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

type MostViewedBlogIdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MostViewedBlogIdsRequest) Reset() {
	*x = MostViewedBlogIdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_activity_activity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MostViewedBlogIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MostViewedBlogIdsRequest) ProtoMessage() {}

func (x *MostViewedBlogIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_activity_activity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MostViewedBlogIdsRequest.ProtoReflect.Descriptor instead.
func (*MostViewedBlogIdsRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_activity_activity_proto_rawDescGZIP(), []int{0}
}

type MostViewedBlogIdsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlogIds []string `protobuf:"bytes,1,rep,name=blogIds,proto3" json:"blogIds,omitempty"`
}

func (x *MostViewedBlogIdsResponse) Reset() {
	*x = MostViewedBlogIdsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_activity_activity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MostViewedBlogIdsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MostViewedBlogIdsResponse) ProtoMessage() {}

func (x *MostViewedBlogIdsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_activity_activity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MostViewedBlogIdsResponse.ProtoReflect.Descriptor instead.
func (*MostViewedBlogIdsResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_activity_activity_proto_rawDescGZIP(), []int{1}
}

func (x *MostViewedBlogIdsResponse) GetBlogIds() []string {
	if x != nil {
		return x.BlogIds
	}
	return nil
}

var File_grpc_proto_activity_activity_proto protoreflect.FileDescriptor

var file_grpc_proto_activity_activity_proto_rawDesc = []byte{
	0x0a, 0x22, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x22, 0x1a,
	0x0a, 0x18, 0x4d, 0x6f, 0x73, 0x74, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x67,
	0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x19, 0x4d, 0x6f,
	0x73, 0x74, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x67, 0x49,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64,
	0x73, 0x32, 0x71, 0x0a, 0x0f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x11, 0x4d, 0x6f, 0x73, 0x74, 0x56, 0x69, 0x65, 0x77,
	0x65, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x73, 0x12, 0x22, 0x2e, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x2e, 0x4d, 0x6f, 0x73, 0x74, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64, 0x42,
	0x6c, 0x6f, 0x67, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x4d, 0x6f, 0x73, 0x74, 0x56, 0x69, 0x65,
	0x77, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x72, 0x69, 0x68, 0x61, 0x72, 0x61, 0x73, 0x75, 0x64, 0x68, 0x61,
	0x6e, 0x2d, 0x6e, 0x69, 0x6e, 0x65, 0x6c, 0x65, 0x61, 0x70, 0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_activity_activity_proto_rawDescOnce sync.Once
	file_grpc_proto_activity_activity_proto_rawDescData = file_grpc_proto_activity_activity_proto_rawDesc
)

func file_grpc_proto_activity_activity_proto_rawDescGZIP() []byte {
	file_grpc_proto_activity_activity_proto_rawDescOnce.Do(func() {
		file_grpc_proto_activity_activity_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_activity_activity_proto_rawDescData)
	})
	return file_grpc_proto_activity_activity_proto_rawDescData
}

var file_grpc_proto_activity_activity_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_proto_activity_activity_proto_goTypes = []interface{}{
	(*MostViewedBlogIdsRequest)(nil),  // 0: activity.MostViewedBlogIdsRequest
	(*MostViewedBlogIdsResponse)(nil), // 1: activity.MostViewedBlogIdsResponse
}
var file_grpc_proto_activity_activity_proto_depIdxs = []int32{
	0, // 0: activity.ActivityService.MostViewedBlogIds:input_type -> activity.MostViewedBlogIdsRequest
	1, // 1: activity.ActivityService.MostViewedBlogIds:output_type -> activity.MostViewedBlogIdsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_proto_activity_activity_proto_init() }
func file_grpc_proto_activity_activity_proto_init() {
	if File_grpc_proto_activity_activity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_activity_activity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MostViewedBlogIdsRequest); i {
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
		file_grpc_proto_activity_activity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MostViewedBlogIdsResponse); i {
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
			RawDescriptor: file_grpc_proto_activity_activity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_activity_activity_proto_goTypes,
		DependencyIndexes: file_grpc_proto_activity_activity_proto_depIdxs,
		MessageInfos:      file_grpc_proto_activity_activity_proto_msgTypes,
	}.Build()
	File_grpc_proto_activity_activity_proto = out.File
	file_grpc_proto_activity_activity_proto_rawDesc = nil
	file_grpc_proto_activity_activity_proto_goTypes = nil
	file_grpc_proto_activity_activity_proto_depIdxs = nil
}
