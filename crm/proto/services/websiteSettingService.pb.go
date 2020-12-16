// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: websiteSettingService.proto

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

type WebsiteSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RootUrl   string `protobuf:"bytes,3,opt,name=root_url,json=rootUrl,proto3" json:"root_url,omitempty"`
	Keywords  string `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Describe  string `protobuf:"bytes,5,opt,name=describe,proto3" json:"describe,omitempty"`
	Tel       string `protobuf:"bytes,6,opt,name=tel,proto3" json:"tel,omitempty"`
	Mobile    string `protobuf:"bytes,7,opt,name=mobile,proto3" json:"mobile,omitempty"`
	AreaId    int64  `protobuf:"varint,8,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Address   string `protobuf:"bytes,9,opt,name=address,proto3" json:"address,omitempty"`
	Lng       string `protobuf:"bytes,10,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat       string `protobuf:"bytes,11,opt,name=lat,proto3" json:"lat,omitempty"`
	Email     string `protobuf:"bytes,12,opt,name=email,proto3" json:"email,omitempty"`
	Qq        string `protobuf:"bytes,13,opt,name=qq,proto3" json:"qq,omitempty"`
	RecordNo  string `protobuf:"bytes,14,opt,name=record_no,json=recordNo,proto3" json:"record_no,omitempty"`
	Copyright string `protobuf:"bytes,15,opt,name=copyright,proto3" json:"copyright,omitempty"`
	CreatedAt string `protobuf:"bytes,16,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,17,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *WebsiteSetting) Reset() {
	*x = WebsiteSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_websiteSettingService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebsiteSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebsiteSetting) ProtoMessage() {}

func (x *WebsiteSetting) ProtoReflect() protoreflect.Message {
	mi := &file_websiteSettingService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebsiteSetting.ProtoReflect.Descriptor instead.
func (*WebsiteSetting) Descriptor() ([]byte, []int) {
	return file_websiteSettingService_proto_rawDescGZIP(), []int{0}
}

func (x *WebsiteSetting) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WebsiteSetting) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WebsiteSetting) GetRootUrl() string {
	if x != nil {
		return x.RootUrl
	}
	return ""
}

func (x *WebsiteSetting) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *WebsiteSetting) GetDescribe() string {
	if x != nil {
		return x.Describe
	}
	return ""
}

func (x *WebsiteSetting) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *WebsiteSetting) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *WebsiteSetting) GetAreaId() int64 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *WebsiteSetting) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *WebsiteSetting) GetLng() string {
	if x != nil {
		return x.Lng
	}
	return ""
}

func (x *WebsiteSetting) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

func (x *WebsiteSetting) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *WebsiteSetting) GetQq() string {
	if x != nil {
		return x.Qq
	}
	return ""
}

func (x *WebsiteSetting) GetRecordNo() string {
	if x != nil {
		return x.RecordNo
	}
	return ""
}

func (x *WebsiteSetting) GetCopyright() string {
	if x != nil {
		return x.Copyright
	}
	return ""
}

func (x *WebsiteSetting) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *WebsiteSetting) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type WebsiteSettingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  *common.Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info    `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	Entity *WebsiteSetting `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *WebsiteSettingResponse) Reset() {
	*x = WebsiteSettingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_websiteSettingService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebsiteSettingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebsiteSettingResponse) ProtoMessage() {}

func (x *WebsiteSettingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_websiteSettingService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebsiteSettingResponse.ProtoReflect.Descriptor instead.
func (*WebsiteSettingResponse) Descriptor() ([]byte, []int) {
	return file_websiteSettingService_proto_rawDescGZIP(), []int{1}
}

func (x *WebsiteSettingResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *WebsiteSettingResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *WebsiteSettingResponse) GetEntity() *WebsiteSetting {
	if x != nil {
		return x.Entity
	}
	return nil
}

var File_websiteSettingService_proto protoreflect.FileDescriptor

var file_websiteSettingService_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x03, 0x0a, 0x0e, 0x57,
	0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x74, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6c, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x71,
	0x71, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x71, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x6e, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x70, 0x79,
	0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x70,
	0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x91, 0x01, 0x0a, 0x16, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x30, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x32, 0xa2, 0x01, 0x0a, 0x15, 0x57, 0x65, 0x62,
	0x73, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x43, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57,
	0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12,
	0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_websiteSettingService_proto_rawDescOnce sync.Once
	file_websiteSettingService_proto_rawDescData = file_websiteSettingService_proto_rawDesc
)

func file_websiteSettingService_proto_rawDescGZIP() []byte {
	file_websiteSettingService_proto_rawDescOnce.Do(func() {
		file_websiteSettingService_proto_rawDescData = protoimpl.X.CompressGZIP(file_websiteSettingService_proto_rawDescData)
	})
	return file_websiteSettingService_proto_rawDescData
}

var file_websiteSettingService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_websiteSettingService_proto_goTypes = []interface{}{
	(*WebsiteSetting)(nil),         // 0: services.WebsiteSetting
	(*WebsiteSettingResponse)(nil), // 1: services.WebsiteSettingResponse
	(*common.Error)(nil),           // 2: common.Error
	(*common.Info)(nil),            // 3: common.Info
}
var file_websiteSettingService_proto_depIdxs = []int32{
	2, // 0: services.WebsiteSettingResponse.error:type_name -> common.Error
	3, // 1: services.WebsiteSettingResponse.info:type_name -> common.Info
	0, // 2: services.WebsiteSettingResponse.entity:type_name -> services.WebsiteSetting
	0, // 3: services.WebsiteSettingService.Get:input_type -> services.WebsiteSetting
	0, // 4: services.WebsiteSettingService.Save:input_type -> services.WebsiteSetting
	1, // 5: services.WebsiteSettingService.Get:output_type -> services.WebsiteSettingResponse
	1, // 6: services.WebsiteSettingService.Save:output_type -> services.WebsiteSettingResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_websiteSettingService_proto_init() }
func file_websiteSettingService_proto_init() {
	if File_websiteSettingService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_websiteSettingService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebsiteSetting); i {
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
		file_websiteSettingService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebsiteSettingResponse); i {
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
			RawDescriptor: file_websiteSettingService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_websiteSettingService_proto_goTypes,
		DependencyIndexes: file_websiteSettingService_proto_depIdxs,
		MessageInfos:      file_websiteSettingService_proto_msgTypes,
	}.Build()
	File_websiteSettingService_proto = out.File
	file_websiteSettingService_proto_rawDesc = nil
	file_websiteSettingService_proto_goTypes = nil
	file_websiteSettingService_proto_depIdxs = nil
}
