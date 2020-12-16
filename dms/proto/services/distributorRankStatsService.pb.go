// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: distributorRankStatsService.proto

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

type DistributorRankStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DistributorId     int64   `protobuf:"varint,2,opt,name=distributor_id,json=distributorId,proto3" json:"distributor_id,omitempty"`
	RankId            int32   `protobuf:"varint,3,opt,name=rank_id,json=rankId,proto3" json:"rank_id,omitempty"`
	RecommendNum      int32   `protobuf:"varint,4,opt,name=recommend_num,json=recommendNum,proto3" json:"recommend_num,omitempty"`
	CustomerNum       int32   `protobuf:"varint,5,opt,name=customer_num,json=customerNum,proto3" json:"customer_num,omitempty"`
	PromotionNum      int32   `protobuf:"varint,6,opt,name=promotion_num,json=promotionNum,proto3" json:"promotion_num,omitempty"`
	ConsumptionNum    int32   `protobuf:"varint,7,opt,name=consumption_num,json=consumptionNum,proto3" json:"consumption_num,omitempty"`
	PromotionAmount   float32 `protobuf:"fixed32,8,opt,name=promotion_amount,json=promotionAmount,proto3" json:"promotion_amount,omitempty"`
	ConsumptionAmount float32 `protobuf:"fixed32,9,opt,name=consumption_amount,json=consumptionAmount,proto3" json:"consumption_amount,omitempty"`
	CreatedAt         string  `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt         string  `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *DistributorRankStats) Reset() {
	*x = DistributorRankStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_distributorRankStatsService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributorRankStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributorRankStats) ProtoMessage() {}

func (x *DistributorRankStats) ProtoReflect() protoreflect.Message {
	mi := &file_distributorRankStatsService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributorRankStats.ProtoReflect.Descriptor instead.
func (*DistributorRankStats) Descriptor() ([]byte, []int) {
	return file_distributorRankStatsService_proto_rawDescGZIP(), []int{0}
}

func (x *DistributorRankStats) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DistributorRankStats) GetDistributorId() int64 {
	if x != nil {
		return x.DistributorId
	}
	return 0
}

func (x *DistributorRankStats) GetRankId() int32 {
	if x != nil {
		return x.RankId
	}
	return 0
}

func (x *DistributorRankStats) GetRecommendNum() int32 {
	if x != nil {
		return x.RecommendNum
	}
	return 0
}

func (x *DistributorRankStats) GetCustomerNum() int32 {
	if x != nil {
		return x.CustomerNum
	}
	return 0
}

func (x *DistributorRankStats) GetPromotionNum() int32 {
	if x != nil {
		return x.PromotionNum
	}
	return 0
}

func (x *DistributorRankStats) GetConsumptionNum() int32 {
	if x != nil {
		return x.ConsumptionNum
	}
	return 0
}

func (x *DistributorRankStats) GetPromotionAmount() float32 {
	if x != nil {
		return x.PromotionAmount
	}
	return 0
}

func (x *DistributorRankStats) GetConsumptionAmount() float32 {
	if x != nil {
		return x.ConsumptionAmount
	}
	return 0
}

func (x *DistributorRankStats) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *DistributorRankStats) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_distributorRankStatsService_proto protoreflect.FileDescriptor

var file_distributorRankStatsService_proto_rawDesc = []byte{
	0x0a, 0x21, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x52, 0x61, 0x6e,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x94, 0x03,
	0x0a, 0x14, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x52, 0x61, 0x6e,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x72, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x23,
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x75, 0x6d, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x12, 0x29, 0x0a, 0x10,
	0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_distributorRankStatsService_proto_rawDescOnce sync.Once
	file_distributorRankStatsService_proto_rawDescData = file_distributorRankStatsService_proto_rawDesc
)

func file_distributorRankStatsService_proto_rawDescGZIP() []byte {
	file_distributorRankStatsService_proto_rawDescOnce.Do(func() {
		file_distributorRankStatsService_proto_rawDescData = protoimpl.X.CompressGZIP(file_distributorRankStatsService_proto_rawDescData)
	})
	return file_distributorRankStatsService_proto_rawDescData
}

var file_distributorRankStatsService_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_distributorRankStatsService_proto_goTypes = []interface{}{
	(*DistributorRankStats)(nil), // 0: services.DistributorRankStats
}
var file_distributorRankStatsService_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_distributorRankStatsService_proto_init() }
func file_distributorRankStatsService_proto_init() {
	if File_distributorRankStatsService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_distributorRankStatsService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributorRankStats); i {
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
			RawDescriptor: file_distributorRankStatsService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_distributorRankStatsService_proto_goTypes,
		DependencyIndexes: file_distributorRankStatsService_proto_depIdxs,
		MessageInfos:      file_distributorRankStatsService_proto_msgTypes,
	}.Build()
	File_distributorRankStatsService_proto = out.File
	file_distributorRankStatsService_proto_rawDesc = nil
	file_distributorRankStatsService_proto_goTypes = nil
	file_distributorRankStatsService_proto_depIdxs = nil
}
