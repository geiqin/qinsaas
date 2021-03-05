// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: groupGoodsStatsService.proto

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

type GroupGoodsStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	GroupId    int64      `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id"`
	CustomerId int64      `protobuf:"varint,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	ItemId     int64      `protobuf:"varint,4,opt,name=item_id,json=itemId,proto3" json:"item_id"`
	GoodsNum   int32      `protobuf:"varint,5,opt,name=goods_num,json=goodsNum,proto3" json:"goods_num"`
	CreatedAt  string     `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt  string     `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Group      *Group     `protobuf:"bytes,8,opt,name=group,proto3" json:"group"`
	Goods      *GoodsInfo `protobuf:"bytes,9,opt,name=goods,proto3" json:"goods"`
}

func (x *GroupGoodsStats) Reset() {
	*x = GroupGoodsStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupGoodsStatsService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupGoodsStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupGoodsStats) ProtoMessage() {}

func (x *GroupGoodsStats) ProtoReflect() protoreflect.Message {
	mi := &file_groupGoodsStatsService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupGoodsStats.ProtoReflect.Descriptor instead.
func (*GroupGoodsStats) Descriptor() ([]byte, []int) {
	return file_groupGoodsStatsService_proto_rawDescGZIP(), []int{0}
}

func (x *GroupGoodsStats) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GroupGoodsStats) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GroupGoodsStats) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *GroupGoodsStats) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *GroupGoodsStats) GetGoodsNum() int32 {
	if x != nil {
		return x.GoodsNum
	}
	return 0
}

func (x *GroupGoodsStats) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GroupGoodsStats) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *GroupGoodsStats) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *GroupGoodsStats) GetGoods() *GoodsInfo {
	if x != nil {
		return x.Goods
	}
	return nil
}

type GroupGoodsStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *GroupGoodsStats   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager      `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*GroupGoodsStats `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error      `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info       `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *GroupGoodsStatsResponse) Reset() {
	*x = GroupGoodsStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupGoodsStatsService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupGoodsStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupGoodsStatsResponse) ProtoMessage() {}

func (x *GroupGoodsStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_groupGoodsStatsService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupGoodsStatsResponse.ProtoReflect.Descriptor instead.
func (*GroupGoodsStatsResponse) Descriptor() ([]byte, []int) {
	return file_groupGoodsStatsService_proto_rawDescGZIP(), []int{1}
}

func (x *GroupGoodsStatsResponse) GetEntity() *GroupGoodsStats {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *GroupGoodsStatsResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *GroupGoodsStatsResponse) GetItems() []*GroupGoodsStats {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *GroupGoodsStatsResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *GroupGoodsStatsResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_groupGoodsStatsService_proto protoreflect.FileDescriptor

var file_groupGoodsStatsService_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa3, 0x02, 0x0a, 0x0f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x29, 0x0a, 0x05, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x22, 0xe9, 0x01, 0x0a, 0x17, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x32, 0xf1, 0x01, 0x0a, 0x16, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x73, 0x57, 0x68, 0x65,
	0x72, 0x65, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x73, 0x57, 0x68, 0x65, 0x72,
	0x65, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_groupGoodsStatsService_proto_rawDescOnce sync.Once
	file_groupGoodsStatsService_proto_rawDescData = file_groupGoodsStatsService_proto_rawDesc
)

func file_groupGoodsStatsService_proto_rawDescGZIP() []byte {
	file_groupGoodsStatsService_proto_rawDescOnce.Do(func() {
		file_groupGoodsStatsService_proto_rawDescData = protoimpl.X.CompressGZIP(file_groupGoodsStatsService_proto_rawDescData)
	})
	return file_groupGoodsStatsService_proto_rawDescData
}

var file_groupGoodsStatsService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_groupGoodsStatsService_proto_goTypes = []interface{}{
	(*GroupGoodsStats)(nil),         // 0: services.GroupGoodsStats
	(*GroupGoodsStatsResponse)(nil), // 1: services.GroupGoodsStatsResponse
	(*Group)(nil),                   // 2: services.Group
	(*GoodsInfo)(nil),               // 3: services.GoodsInfo
	(*common.Pager)(nil),            // 4: common.Pager
	(*common.Error)(nil),            // 5: common.Error
	(*common.Info)(nil),             // 6: common.Info
	(*GroupStatsWhere)(nil),         // 7: services.GroupStatsWhere
}
var file_groupGoodsStatsService_proto_depIdxs = []int32{
	2,  // 0: services.GroupGoodsStats.group:type_name -> services.Group
	3,  // 1: services.GroupGoodsStats.goods:type_name -> services.GoodsInfo
	0,  // 2: services.GroupGoodsStatsResponse.entity:type_name -> services.GroupGoodsStats
	4,  // 3: services.GroupGoodsStatsResponse.pager:type_name -> common.Pager
	0,  // 4: services.GroupGoodsStatsResponse.items:type_name -> services.GroupGoodsStats
	5,  // 5: services.GroupGoodsStatsResponse.error:type_name -> common.Error
	6,  // 6: services.GroupGoodsStatsResponse.info:type_name -> common.Info
	7,  // 7: services.GroupGoodsStatsService.Search:input_type -> services.GroupStatsWhere
	7,  // 8: services.GroupGoodsStatsService.Get:input_type -> services.GroupStatsWhere
	7,  // 9: services.GroupGoodsStatsService.List:input_type -> services.GroupStatsWhere
	1,  // 10: services.GroupGoodsStatsService.Search:output_type -> services.GroupGoodsStatsResponse
	1,  // 11: services.GroupGoodsStatsService.Get:output_type -> services.GroupGoodsStatsResponse
	1,  // 12: services.GroupGoodsStatsService.List:output_type -> services.GroupGoodsStatsResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_groupGoodsStatsService_proto_init() }
func file_groupGoodsStatsService_proto_init() {
	if File_groupGoodsStatsService_proto != nil {
		return
	}
	file_groupService_proto_init()
	file_goodsInfoService_proto_init()
	file_groupStatsService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_groupGoodsStatsService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupGoodsStats); i {
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
		file_groupGoodsStatsService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupGoodsStatsResponse); i {
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
			RawDescriptor: file_groupGoodsStatsService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_groupGoodsStatsService_proto_goTypes,
		DependencyIndexes: file_groupGoodsStatsService_proto_depIdxs,
		MessageInfos:      file_groupGoodsStatsService_proto_msgTypes,
	}.Build()
	File_groupGoodsStatsService_proto = out.File
	file_groupGoodsStatsService_proto_rawDesc = nil
	file_groupGoodsStatsService_proto_goTypes = nil
	file_groupGoodsStatsService_proto_depIdxs = nil
}
