// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: deliveryRangeService.proto

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

type DeliveryRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	DeliveryId    int64   `protobuf:"varint,2,opt,name=delivery_id,json=deliveryId,proto3" json:"delivery_id"`
	AreaName      string  `protobuf:"bytes,3,opt,name=area_name,json=areaName,proto3" json:"area_name"`
	StartPrice    float32 `protobuf:"fixed32,4,opt,name=start_price,json=startPrice,proto3" json:"start_price"`
	DeliveryPrice float32 `protobuf:"fixed32,5,opt,name=delivery_price,json=deliveryPrice,proto3" json:"delivery_price"`
	Method        int32   `protobuf:"varint,6,opt,name=method,proto3" json:"method"`
	RegionRadius  float32 `protobuf:"fixed32,7,opt,name=region_radius,json=regionRadius,proto3" json:"region_radius"`
	RegionData    string  `protobuf:"bytes,8,opt,name=region_data,json=regionData,proto3" json:"region_data"`
	CreatedAt     string  `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt     string  `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	// @inject_tag: gorm:"foreignKey:DeliveryRangeId"
	Coordinate []*DeliveryMap `protobuf:"bytes,11,rep,name=coordinate,proto3" json:"coordinate" gorm:"foreignKey:DeliveryRangeId"`
}

func (x *DeliveryRange) Reset() {
	*x = DeliveryRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deliveryRangeService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryRange) ProtoMessage() {}

func (x *DeliveryRange) ProtoReflect() protoreflect.Message {
	mi := &file_deliveryRangeService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryRange.ProtoReflect.Descriptor instead.
func (*DeliveryRange) Descriptor() ([]byte, []int) {
	return file_deliveryRangeService_proto_rawDescGZIP(), []int{0}
}

func (x *DeliveryRange) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeliveryRange) GetDeliveryId() int64 {
	if x != nil {
		return x.DeliveryId
	}
	return 0
}

func (x *DeliveryRange) GetAreaName() string {
	if x != nil {
		return x.AreaName
	}
	return ""
}

func (x *DeliveryRange) GetStartPrice() float32 {
	if x != nil {
		return x.StartPrice
	}
	return 0
}

func (x *DeliveryRange) GetDeliveryPrice() float32 {
	if x != nil {
		return x.DeliveryPrice
	}
	return 0
}

func (x *DeliveryRange) GetMethod() int32 {
	if x != nil {
		return x.Method
	}
	return 0
}

func (x *DeliveryRange) GetRegionRadius() float32 {
	if x != nil {
		return x.RegionRadius
	}
	return 0
}

func (x *DeliveryRange) GetRegionData() string {
	if x != nil {
		return x.RegionData
	}
	return ""
}

func (x *DeliveryRange) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *DeliveryRange) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *DeliveryRange) GetCoordinate() []*DeliveryMap {
	if x != nil {
		return x.Coordinate
	}
	return nil
}

type DeliveryMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	DeliveryRangeId int32  `protobuf:"varint,2,opt,name=delivery_range_id,json=deliveryRangeId,proto3" json:"delivery_range_id"`
	Lng             string `protobuf:"bytes,3,opt,name=lng,proto3" json:"lng"`
	Lat             string `protobuf:"bytes,4,opt,name=lat,proto3" json:"lat"`
	CreatedAt       string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt       string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *DeliveryMap) Reset() {
	*x = DeliveryMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deliveryRangeService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryMap) ProtoMessage() {}

func (x *DeliveryMap) ProtoReflect() protoreflect.Message {
	mi := &file_deliveryRangeService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryMap.ProtoReflect.Descriptor instead.
func (*DeliveryMap) Descriptor() ([]byte, []int) {
	return file_deliveryRangeService_proto_rawDescGZIP(), []int{1}
}

func (x *DeliveryMap) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeliveryMap) GetDeliveryRangeId() int32 {
	if x != nil {
		return x.DeliveryRangeId
	}
	return 0
}

func (x *DeliveryMap) GetLng() string {
	if x != nil {
		return x.Lng
	}
	return ""
}

func (x *DeliveryMap) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

func (x *DeliveryMap) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *DeliveryMap) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_deliveryRangeService_proto protoreflect.FileDescriptor

var file_deliveryRangeService_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0xf8, 0x02, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x72, 0x65,
	0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72,
	0x65, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x5f, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x64, 0x69, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x35, 0x0a, 0x0a, 0x63, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x4d, 0x61, 0x70, 0x52, 0x0a, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x22, 0xab, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x61,
	0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x6c, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x12,
	0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x61,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deliveryRangeService_proto_rawDescOnce sync.Once
	file_deliveryRangeService_proto_rawDescData = file_deliveryRangeService_proto_rawDesc
)

func file_deliveryRangeService_proto_rawDescGZIP() []byte {
	file_deliveryRangeService_proto_rawDescOnce.Do(func() {
		file_deliveryRangeService_proto_rawDescData = protoimpl.X.CompressGZIP(file_deliveryRangeService_proto_rawDescData)
	})
	return file_deliveryRangeService_proto_rawDescData
}

var file_deliveryRangeService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_deliveryRangeService_proto_goTypes = []interface{}{
	(*DeliveryRange)(nil), // 0: services.DeliveryRange
	(*DeliveryMap)(nil),   // 1: services.DeliveryMap
}
var file_deliveryRangeService_proto_depIdxs = []int32{
	1, // 0: services.DeliveryRange.coordinate:type_name -> services.DeliveryMap
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_deliveryRangeService_proto_init() }
func file_deliveryRangeService_proto_init() {
	if File_deliveryRangeService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deliveryRangeService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryRange); i {
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
		file_deliveryRangeService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryMap); i {
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
			RawDescriptor: file_deliveryRangeService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_deliveryRangeService_proto_goTypes,
		DependencyIndexes: file_deliveryRangeService_proto_depIdxs,
		MessageInfos:      file_deliveryRangeService_proto_msgTypes,
	}.Build()
	File_deliveryRangeService_proto = out.File
	file_deliveryRangeService_proto_rawDesc = nil
	file_deliveryRangeService_proto_goTypes = nil
	file_deliveryRangeService_proto_depIdxs = nil
}
