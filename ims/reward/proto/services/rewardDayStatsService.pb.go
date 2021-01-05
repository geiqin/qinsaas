// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: rewardDayStatsService.proto

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

type RewardDayStatsWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged     int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Top       int32  `protobuf:"varint,3,opt,name=top,proto3" json:"top"`
	Id        int64  `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	RewardId  int64  `protobuf:"varint,5,opt,name=reward_id,json=rewardId,proto3" json:"reward_id"`
	StartDate string `protobuf:"bytes,6,opt,name=start_date,json=startDate,proto3" json:"start_date"`
	EndDate   string `protobuf:"bytes,7,opt,name=end_date,json=endDate,proto3" json:"end_date"`
}

func (x *RewardDayStatsWhere) Reset() {
	*x = RewardDayStatsWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewardDayStatsService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardDayStatsWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardDayStatsWhere) ProtoMessage() {}

func (x *RewardDayStatsWhere) ProtoReflect() protoreflect.Message {
	mi := &file_rewardDayStatsService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardDayStatsWhere.ProtoReflect.Descriptor instead.
func (*RewardDayStatsWhere) Descriptor() ([]byte, []int) {
	return file_rewardDayStatsService_proto_rawDescGZIP(), []int{0}
}

func (x *RewardDayStatsWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *RewardDayStatsWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *RewardDayStatsWhere) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *RewardDayStatsWhere) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RewardDayStatsWhere) GetRewardId() int64 {
	if x != nil {
		return x.RewardId
	}
	return 0
}

func (x *RewardDayStatsWhere) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *RewardDayStatsWhere) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type RewardDayStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	RewardId       int64   `protobuf:"varint,2,opt,name=reward_id,json=rewardId,proto3" json:"reward_id"`
	PayAmount      float32 `protobuf:"fixed32,3,opt,name=pay_amount,json=payAmount,proto3" json:"pay_amount"`
	DiscountAmount float32 `protobuf:"fixed32,4,opt,name=discount_amount,json=discountAmount,proto3" json:"discount_amount"`
	PayNum         int32   `protobuf:"varint,5,opt,name=pay_num,json=payNum,proto3" json:"pay_num"`
	BuyerNum       int32   `protobuf:"varint,6,opt,name=buyer_num,json=buyerNum,proto3" json:"buyer_num"`
	GoodsNum       int32   `protobuf:"varint,7,opt,name=goods_num,json=goodsNum,proto3" json:"goods_num"`
	NewCustomerNum int32   `protobuf:"varint,8,opt,name=new_customer_num,json=newCustomerNum,proto3" json:"new_customer_num"`
	OldCustomerNum int32   `protobuf:"varint,9,opt,name=old_customer_num,json=oldCustomerNum,proto3" json:"old_customer_num"`
	StatsDate      string  `protobuf:"bytes,10,opt,name=stats_date,json=statsDate,proto3" json:"stats_date"`
	CreatedAt      string  `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt      string  `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Reward         *Reward `protobuf:"bytes,13,opt,name=reward,proto3" json:"reward"`
}

func (x *RewardDayStats) Reset() {
	*x = RewardDayStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewardDayStatsService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardDayStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardDayStats) ProtoMessage() {}

func (x *RewardDayStats) ProtoReflect() protoreflect.Message {
	mi := &file_rewardDayStatsService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardDayStats.ProtoReflect.Descriptor instead.
func (*RewardDayStats) Descriptor() ([]byte, []int) {
	return file_rewardDayStatsService_proto_rawDescGZIP(), []int{1}
}

func (x *RewardDayStats) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RewardDayStats) GetRewardId() int64 {
	if x != nil {
		return x.RewardId
	}
	return 0
}

func (x *RewardDayStats) GetPayAmount() float32 {
	if x != nil {
		return x.PayAmount
	}
	return 0
}

func (x *RewardDayStats) GetDiscountAmount() float32 {
	if x != nil {
		return x.DiscountAmount
	}
	return 0
}

func (x *RewardDayStats) GetPayNum() int32 {
	if x != nil {
		return x.PayNum
	}
	return 0
}

func (x *RewardDayStats) GetBuyerNum() int32 {
	if x != nil {
		return x.BuyerNum
	}
	return 0
}

func (x *RewardDayStats) GetGoodsNum() int32 {
	if x != nil {
		return x.GoodsNum
	}
	return 0
}

func (x *RewardDayStats) GetNewCustomerNum() int32 {
	if x != nil {
		return x.NewCustomerNum
	}
	return 0
}

func (x *RewardDayStats) GetOldCustomerNum() int32 {
	if x != nil {
		return x.OldCustomerNum
	}
	return 0
}

func (x *RewardDayStats) GetStatsDate() string {
	if x != nil {
		return x.StatsDate
	}
	return ""
}

func (x *RewardDayStats) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *RewardDayStats) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *RewardDayStats) GetReward() *Reward {
	if x != nil {
		return x.Reward
	}
	return nil
}

type RewardDayStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *RewardDayStats   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*RewardDayStats `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error     `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info      `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *RewardDayStatsResponse) Reset() {
	*x = RewardDayStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewardDayStatsService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardDayStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardDayStatsResponse) ProtoMessage() {}

func (x *RewardDayStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rewardDayStatsService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardDayStatsResponse.ProtoReflect.Descriptor instead.
func (*RewardDayStatsResponse) Descriptor() ([]byte, []int) {
	return file_rewardDayStatsService_proto_rawDescGZIP(), []int{2}
}

func (x *RewardDayStatsResponse) GetEntity() *RewardDayStats {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *RewardDayStatsResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *RewardDayStatsResponse) GetItems() []*RewardDayStats {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *RewardDayStatsResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *RewardDayStatsResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_rewardDayStatsService_proto protoreflect.FileDescriptor

var file_rewardDayStatsService_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x72, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc1, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x22, 0xb3, 0x03, 0x0a, 0x0e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x44, 0x61,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x61, 0x79, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61,
	0x79, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x62, 0x75, 0x79, 0x65, 0x72, 0x4e, 0x75,
	0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x28,
	0x0a, 0x10, 0x6e, 0x65, 0x77, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6e, 0x65, 0x77, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x28, 0x0a, 0x10, 0x6f, 0x6c, 0x64, 0x5f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0e, 0x6f, 0x6c, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e,
	0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x28, 0x0a, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x52, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x22, 0xe6, 0x01, 0x0a, 0x16, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x44, 0x61, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x32, 0xae, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x44, 0x61, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x06,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a,
	0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rewardDayStatsService_proto_rawDescOnce sync.Once
	file_rewardDayStatsService_proto_rawDescData = file_rewardDayStatsService_proto_rawDesc
)

func file_rewardDayStatsService_proto_rawDescGZIP() []byte {
	file_rewardDayStatsService_proto_rawDescOnce.Do(func() {
		file_rewardDayStatsService_proto_rawDescData = protoimpl.X.CompressGZIP(file_rewardDayStatsService_proto_rawDescData)
	})
	return file_rewardDayStatsService_proto_rawDescData
}

var file_rewardDayStatsService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rewardDayStatsService_proto_goTypes = []interface{}{
	(*RewardDayStatsWhere)(nil),    // 0: services.RewardDayStatsWhere
	(*RewardDayStats)(nil),         // 1: services.RewardDayStats
	(*RewardDayStatsResponse)(nil), // 2: services.RewardDayStatsResponse
	(*Reward)(nil),                 // 3: services.Reward
	(*common.Pager)(nil),           // 4: common.Pager
	(*common.Error)(nil),           // 5: common.Error
	(*common.Info)(nil),            // 6: common.Info
}
var file_rewardDayStatsService_proto_depIdxs = []int32{
	3, // 0: services.RewardDayStats.reward:type_name -> services.Reward
	1, // 1: services.RewardDayStatsResponse.entity:type_name -> services.RewardDayStats
	4, // 2: services.RewardDayStatsResponse.pager:type_name -> common.Pager
	1, // 3: services.RewardDayStatsResponse.items:type_name -> services.RewardDayStats
	5, // 4: services.RewardDayStatsResponse.error:type_name -> common.Error
	6, // 5: services.RewardDayStatsResponse.info:type_name -> common.Info
	0, // 6: services.RewardDayStatsService.Search:input_type -> services.RewardDayStatsWhere
	0, // 7: services.RewardDayStatsService.Get:input_type -> services.RewardDayStatsWhere
	2, // 8: services.RewardDayStatsService.Search:output_type -> services.RewardDayStatsResponse
	2, // 9: services.RewardDayStatsService.Get:output_type -> services.RewardDayStatsResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_rewardDayStatsService_proto_init() }
func file_rewardDayStatsService_proto_init() {
	if File_rewardDayStatsService_proto != nil {
		return
	}
	file_rewardService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rewardDayStatsService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardDayStatsWhere); i {
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
		file_rewardDayStatsService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardDayStats); i {
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
		file_rewardDayStatsService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardDayStatsResponse); i {
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
			RawDescriptor: file_rewardDayStatsService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rewardDayStatsService_proto_goTypes,
		DependencyIndexes: file_rewardDayStatsService_proto_depIdxs,
		MessageInfos:      file_rewardDayStatsService_proto_msgTypes,
	}.Build()
	File_rewardDayStatsService_proto = out.File
	file_rewardDayStatsService_proto_rawDesc = nil
	file_rewardDayStatsService_proto_goTypes = nil
	file_rewardDayStatsService_proto_depIdxs = nil
}
