// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: fetchGalleryService.proto

package services

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FetchGallery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FetchId   int64   `protobuf:"varint,2,opt,name=fetch_id,json=fetchId,proto3" json:"fetch_id,omitempty"`
	MediaId   int64   `protobuf:"varint,3,opt,name=media_id,json=mediaId,proto3" json:"media_id,omitempty"`
	MediaUrl  string  `protobuf:"bytes,4,opt,name=media_url,json=mediaUrl,proto3" json:"media_url,omitempty"`
	Defaulted bool    `protobuf:"varint,5,opt,name=defaulted,proto3" json:"defaulted,omitempty"`
	CreatedAt string  `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string  `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Ids       []int64 `protobuf:"varint,8,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *FetchGallery) Reset() {
	*x = FetchGallery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fetchGalleryService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchGallery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchGallery) ProtoMessage() {}

func (x *FetchGallery) ProtoReflect() protoreflect.Message {
	mi := &file_fetchGalleryService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchGallery.ProtoReflect.Descriptor instead.
func (*FetchGallery) Descriptor() ([]byte, []int) {
	return file_fetchGalleryService_proto_rawDescGZIP(), []int{0}
}

func (x *FetchGallery) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FetchGallery) GetFetchId() int64 {
	if x != nil {
		return x.FetchId
	}
	return 0
}

func (x *FetchGallery) GetMediaId() int64 {
	if x != nil {
		return x.MediaId
	}
	return 0
}

func (x *FetchGallery) GetMediaUrl() string {
	if x != nil {
		return x.MediaUrl
	}
	return ""
}

func (x *FetchGallery) GetDefaulted() bool {
	if x != nil {
		return x.Defaulted
	}
	return false
}

func (x *FetchGallery) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *FetchGallery) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *FetchGallery) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

var File_fetchGalleryService_proto protoreflect.FileDescriptor

var file_fetchGalleryService_proto_rawDesc = []byte{
	0x0a, 0x19, 0x66, 0x65, 0x74, 0x63, 0x68, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0xdf, 0x01, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x47,
	0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x66, 0x65, 0x74, 0x63, 0x68, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x55, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fetchGalleryService_proto_rawDescOnce sync.Once
	file_fetchGalleryService_proto_rawDescData = file_fetchGalleryService_proto_rawDesc
)

func file_fetchGalleryService_proto_rawDescGZIP() []byte {
	file_fetchGalleryService_proto_rawDescOnce.Do(func() {
		file_fetchGalleryService_proto_rawDescData = protoimpl.X.CompressGZIP(file_fetchGalleryService_proto_rawDescData)
	})
	return file_fetchGalleryService_proto_rawDescData
}

var file_fetchGalleryService_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fetchGalleryService_proto_goTypes = []interface{}{
	(*FetchGallery)(nil), // 0: services.FetchGallery
}
var file_fetchGalleryService_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fetchGalleryService_proto_init() }
func file_fetchGalleryService_proto_init() {
	if File_fetchGalleryService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fetchGalleryService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchGallery); i {
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
			RawDescriptor: file_fetchGalleryService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fetchGalleryService_proto_goTypes,
		DependencyIndexes: file_fetchGalleryService_proto_depIdxs,
		MessageInfos:      file_fetchGalleryService_proto_msgTypes,
	}.Build()
	File_fetchGalleryService_proto = out.File
	file_fetchGalleryService_proto_rawDesc = nil
	file_fetchGalleryService_proto_goTypes = nil
	file_fetchGalleryService_proto_depIdxs = nil
}
