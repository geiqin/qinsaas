// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: rewardStatsService.proto

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

type RewardStatsWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Top      int32 `protobuf:"varint,3,opt,name=top,proto3" json:"top"`
	Id       int64 `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	RewardId int64 `protobuf:"varint,5,opt,name=reward_id,json=rewardId,proto3" json:"reward_id"`
}

func (x *RewardStatsWhere) Reset() {
	*x = RewardStatsWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewardStatsService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardStatsWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardStatsWhere) ProtoMessage() {}

func (x *RewardStatsWhere) ProtoReflect() protoreflect.Message {
	mi := &file_rewardStatsService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardStatsWhere.ProtoReflect.Descriptor instead.
func (*RewardStatsWhere) Descriptor() ([]byte, []int) {
	return file_rewardStatsService_proto_rawDescGZIP(), []int{0}
}

func (x *RewardStatsWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *RewardStatsWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *RewardStatsWhere) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *RewardStatsWhere) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RewardStatsWhere) GetRewardId() int64 {
	if x != nil {
		return x.RewardId
	}
	return 0
}

type RewardStats struct {
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
	CreatedAt      string  `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt      string  `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Reward         *Reward `protobuf:"bytes,12,opt,name=reward,proto3" json:"reward"`
}

func (x *RewardStats) Reset() {
	*x = RewardStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewardStatsService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardStats) ProtoMessage() {}

func (x *RewardStats) ProtoReflect() protoreflect.Message {
	mi := &file_rewardStatsService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardStats.ProtoReflect.Descriptor instead.
func (*RewardStats) Descriptor() ([]byte, []int) {
	return file_rewardStatsService_proto_rawDescGZIP(), []int{1}
}

func (x *RewardStats) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RewardStats) GetRewardId() int64 {
	if x != nil {
		return x.RewardId
	}
	return 0
}

func (x *RewardStats) GetPayAmount() float32 {
	if x != nil {
		return x.PayAmount
	}
	return 0
}

func (x *RewardStats) GetDiscountAmount() float32 {
	if x != nil {
		return x.DiscountAmount
	}
	return 0
}

func (x *RewardStats) GetPayNum() int32 {
	if x != nil {
		return x.PayNum
	}
	return 0
}

func (x *RewardStats) GetBuyerNum() int32 {
	if x != nil {
		return x.BuyerNum
	}
	return 0
}

func (x *RewardStats) GetGoodsNum() int32 {
	if x != nil {
		return x.GoodsNum
	}
	return 0
}

func (x *RewardStats) GetNewCustomerNum() int32 {
	if x != nil {
		return x.NewCustomerNum
	}
	return 0
}

func (x *RewardStats) GetOldCustomerNum() int32 {
	if x != nil {
		return x.OldCustomerNum
	}
	return 0
}

func (x *RewardStats) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *RewardStats) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *RewardStats) GetReward() *Reward {
	if x != nil {
		return x.Reward
	}
	return nil
}

type RewardStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *RewardStats   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager  `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*RewardStats `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error  `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info   `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *RewardStatsResponse) Reset() {
	*x = RewardStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewardStatsService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardStatsResponse) ProtoMessage() {}

func (x *RewardStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rewardStatsService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardStatsResponse.ProtoReflect.Descriptor instead.
func (*RewardStatsResponse) Descriptor() ([]byte, []int) {
	return file_rewardStatsService_proto_rawDescGZIP(), []int{2}
}

func (x *RewardStatsResponse) GetEntity() *RewardStats {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *RewardStatsResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *RewardStatsResponse) GetItems() []*RewardStats {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *RewardStatsResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *RewardStatsResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_rewardStatsService_proto protoreflect.FileDescriptor

var file_rewardStatsService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a,
	0x10, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x57, 0x68, 0x65, 0x72,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x49, 0x64, 0x22, 0x91, 0x03, 0x0a, 0x0b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x61, 0x79, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x5f,
	0x6e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x79, 0x4e, 0x75,
	0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x62, 0x75, 0x79, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x1b,
	0x0a, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x28, 0x0a, 0x10, 0x6e,
	0x65, 0x77, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6e, 0x65, 0x77, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x28, 0x0a, 0x10, 0x6f, 0x6c, 0x64, 0x5f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0e, 0x6f, 0x6c, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a,
	0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52,
	0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x22, 0xdd, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23,
	0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xe5, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45,
	0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x57,
	0x68, 0x65, 0x72, 0x65, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x05, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x73, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rewardStatsService_proto_rawDescOnce sync.Once
	file_rewardStatsService_proto_rawDescData = file_rewardStatsService_proto_rawDesc
)

func file_rewardStatsService_proto_rawDescGZIP() []byte {
	file_rewardStatsService_proto_rawDescOnce.Do(func() {
		file_rewardStatsService_proto_rawDescData = protoimpl.X.CompressGZIP(file_rewardStatsService_proto_rawDescData)
	})
	return file_rewardStatsService_proto_rawDescData
}

var file_rewardStatsService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rewardStatsService_proto_goTypes = []interface{}{
	(*RewardStatsWhere)(nil),    // 0: services.RewardStatsWhere
	(*RewardStats)(nil),         // 1: services.RewardStats
	(*RewardStatsResponse)(nil), // 2: services.RewardStatsResponse
	(*Reward)(nil),              // 3: services.Reward
	(*common.Pager)(nil),        // 4: common.Pager
	(*common.Error)(nil),        // 5: common.Error
	(*common.Info)(nil),         // 6: common.Info
}
var file_rewardStatsService_proto_depIdxs = []int32{
	3, // 0: services.RewardStats.reward:type_name -> services.Reward
	1, // 1: services.RewardStatsResponse.entity:type_name -> services.RewardStats
	4, // 2: services.RewardStatsResponse.pager:type_name -> common.Pager
	1, // 3: services.RewardStatsResponse.items:type_name -> services.RewardStats
	5, // 4: services.RewardStatsResponse.error:type_name -> common.Error
	6, // 5: services.RewardStatsResponse.info:type_name -> common.Info
	0, // 6: services.RewardStatsService.Search:input_type -> services.RewardStatsWhere
	0, // 7: services.RewardStatsService.Get:input_type -> services.RewardStatsWhere
	0, // 8: services.RewardStatsService.Reset:input_type -> services.RewardStatsWhere
	2, // 9: services.RewardStatsService.Search:output_type -> services.RewardStatsResponse
	2, // 10: services.RewardStatsService.Get:output_type -> services.RewardStatsResponse
	2, // 11: services.RewardStatsService.Reset:output_type -> services.RewardStatsResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_rewardStatsService_proto_init() }
func file_rewardStatsService_proto_init() {
	if File_rewardStatsService_proto != nil {
		return
	}
	file_rewardService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rewardStatsService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardStatsWhere); i {
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
		file_rewardStatsService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardStats); i {
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
		file_rewardStatsService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardStatsResponse); i {
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
			RawDescriptor: file_rewardStatsService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rewardStatsService_proto_goTypes,
		DependencyIndexes: file_rewardStatsService_proto_depIdxs,
		MessageInfos:      file_rewardStatsService_proto_msgTypes,
	}.Build()
	File_rewardStatsService_proto = out.File
	file_rewardStatsService_proto_rawDesc = nil
	file_rewardStatsService_proto_goTypes = nil
	file_rewardStatsService_proto_depIdxs = nil
}