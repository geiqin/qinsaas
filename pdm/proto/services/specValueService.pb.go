// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: specValueService.proto

package services

import (
	common "github.com/geiqin/microkit/protobuf/common"
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

type SpecValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	SpecId     int32  `protobuf:"varint,2,opt,name=spec_id,json=specId,proto3" json:"spec_id"`
	Alias      string `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias"`
	SpecValue  string `protobuf:"bytes,4,opt,name=spec_value,json=specValue,proto3" json:"spec_value"`
	SpecImgId  int64  `protobuf:"varint,5,opt,name=spec_img_id,json=specImgId,proto3" json:"spec_img_id"`
	SpecImgUrl string `protobuf:"bytes,6,opt,name=spec_img_url,json=specImgUrl,proto3" json:"spec_img_url"`
	Sorting    int32  `protobuf:"varint,7,opt,name=sorting,proto3" json:"sorting"`
}

func (x *SpecValue) Reset() {
	*x = SpecValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_specValueService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecValue) ProtoMessage() {}

func (x *SpecValue) ProtoReflect() protoreflect.Message {
	mi := &file_specValueService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecValue.ProtoReflect.Descriptor instead.
func (*SpecValue) Descriptor() ([]byte, []int) {
	return file_specValueService_proto_rawDescGZIP(), []int{0}
}

func (x *SpecValue) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SpecValue) GetSpecId() int32 {
	if x != nil {
		return x.SpecId
	}
	return 0
}

func (x *SpecValue) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *SpecValue) GetSpecValue() string {
	if x != nil {
		return x.SpecValue
	}
	return ""
}

func (x *SpecValue) GetSpecImgId() int64 {
	if x != nil {
		return x.SpecImgId
	}
	return 0
}

func (x *SpecValue) GetSpecImgUrl() string {
	if x != nil {
		return x.SpecImgUrl
	}
	return ""
}

func (x *SpecValue) GetSorting() int32 {
	if x != nil {
		return x.Sorting
	}
	return 0
}

type SpecValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *SpecValue    `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*SpecValue  `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *SpecValueResponse) Reset() {
	*x = SpecValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_specValueService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecValueResponse) ProtoMessage() {}

func (x *SpecValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_specValueService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecValueResponse.ProtoReflect.Descriptor instead.
func (*SpecValueResponse) Descriptor() ([]byte, []int) {
	return file_specValueService_proto_rawDescGZIP(), []int{1}
}

func (x *SpecValueResponse) GetEntity() *SpecValue {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *SpecValueResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *SpecValueResponse) GetItems() []*SpecValue {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *SpecValueResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SpecValueResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_specValueService_proto protoreflect.FileDescriptor

var file_specValueService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x70, 0x65, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x09, 0x53, 0x70, 0x65, 0x63, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x70, 0x65, 0x63, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x70, 0x65, 0x63, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x69, 0x6d, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x70, 0x65, 0x63, 0x49, 0x6d, 0x67, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x69, 0x6d, 0x67, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x70, 0x65, 0x63, 0x49, 0x6d, 0x67,
	0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x22, 0xd7, 0x01,
	0x0a, 0x11, 0x53, 0x70, 0x65, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53,
	0x70, 0x65, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x70, 0x65, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xc3, 0x02, 0x0a, 0x10, 0x53, 0x70, 0x65, 0x63,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x70, 0x65, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x70, 0x65, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x65, 0x63,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3c, 0x0a, 0x06, 0x49, 0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_specValueService_proto_rawDescOnce sync.Once
	file_specValueService_proto_rawDescData = file_specValueService_proto_rawDesc
)

func file_specValueService_proto_rawDescGZIP() []byte {
	file_specValueService_proto_rawDescOnce.Do(func() {
		file_specValueService_proto_rawDescData = protoimpl.X.CompressGZIP(file_specValueService_proto_rawDescData)
	})
	return file_specValueService_proto_rawDescData
}

var file_specValueService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_specValueService_proto_goTypes = []interface{}{
	(*SpecValue)(nil),         // 0: services.SpecValue
	(*SpecValueResponse)(nil), // 1: services.SpecValueResponse
	(*common.Pager)(nil),      // 2: common.Pager
	(*common.Error)(nil),      // 3: common.Error
	(*common.Info)(nil),       // 4: common.Info
}
var file_specValueService_proto_depIdxs = []int32{
	0,  // 0: services.SpecValueResponse.entity:type_name -> services.SpecValue
	2,  // 1: services.SpecValueResponse.pager:type_name -> common.Pager
	0,  // 2: services.SpecValueResponse.items:type_name -> services.SpecValue
	3,  // 3: services.SpecValueResponse.error:type_name -> common.Error
	4,  // 4: services.SpecValueResponse.info:type_name -> common.Info
	0,  // 5: services.SpecValueService.Create:input_type -> services.SpecValue
	0,  // 6: services.SpecValueService.Delete:input_type -> services.SpecValue
	0,  // 7: services.SpecValueService.Get:input_type -> services.SpecValue
	0,  // 8: services.SpecValueService.List:input_type -> services.SpecValue
	0,  // 9: services.SpecValueService.IsUsed:input_type -> services.SpecValue
	1,  // 10: services.SpecValueService.Create:output_type -> services.SpecValueResponse
	1,  // 11: services.SpecValueService.Delete:output_type -> services.SpecValueResponse
	1,  // 12: services.SpecValueService.Get:output_type -> services.SpecValueResponse
	1,  // 13: services.SpecValueService.List:output_type -> services.SpecValueResponse
	1,  // 14: services.SpecValueService.IsUsed:output_type -> services.SpecValueResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_specValueService_proto_init() }
func file_specValueService_proto_init() {
	if File_specValueService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_specValueService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecValue); i {
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
		file_specValueService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecValueResponse); i {
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
			RawDescriptor: file_specValueService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_specValueService_proto_goTypes,
		DependencyIndexes: file_specValueService_proto_depIdxs,
		MessageInfos:      file_specValueService_proto_msgTypes,
	}.Build()
	File_specValueService_proto = out.File
	file_specValueService_proto_rawDesc = nil
	file_specValueService_proto_goTypes = nil
	file_specValueService_proto_depIdxs = nil
}
