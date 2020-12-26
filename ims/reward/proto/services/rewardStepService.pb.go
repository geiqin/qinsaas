// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: rewardStepService.proto

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

type RewardStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	RewardId       int64           `protobuf:"varint,2,opt,name=reward_id,json=rewardId,proto3" json:"reward_id"`
	UnitType       int32           `protobuf:"varint,3,opt,name=unit_type,json=unitType,proto3" json:"unit_type"`
	ConditionPrice float32         `protobuf:"fixed32,4,opt,name=condition_price,json=conditionPrice,proto3" json:"condition_price"`
	ConditionNum   int32           `protobuf:"varint,5,opt,name=condition_num,json=conditionNum,proto3" json:"condition_num"`
	Preferent      int32           `protobuf:"varint,6,opt,name=preferent,proto3" json:"preferent"`
	Money          float32         `protobuf:"fixed32,7,opt,name=money,proto3" json:"money"`
	Discount       float32         `protobuf:"fixed32,8,opt,name=discount,proto3" json:"discount"`
	Point          int32           `protobuf:"varint,9,opt,name=point,proto3" json:"point"`
	PresentId      int64           `protobuf:"varint,10,opt,name=present_id,json=presentId,proto3" json:"present_id"`
	PresentNum     int32           `protobuf:"varint,11,opt,name=present_num,json=presentNum,proto3" json:"present_num"`
	CreatedAt      string          `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt      string          `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Coupons        []*RewardCoupon `protobuf:"bytes,14,rep,name=coupons,proto3" json:"coupons"`
	FreePostage    bool            `protobuf:"varint,15,opt,name=free_postage,json=freePostage,proto3" json:"free_postage"` // 是否包邮
}

func (x *RewardStep) Reset() {
	*x = RewardStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewardStepService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardStep) ProtoMessage() {}

func (x *RewardStep) ProtoReflect() protoreflect.Message {
	mi := &file_rewardStepService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardStep.ProtoReflect.Descriptor instead.
func (*RewardStep) Descriptor() ([]byte, []int) {
	return file_rewardStepService_proto_rawDescGZIP(), []int{0}
}

func (x *RewardStep) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RewardStep) GetRewardId() int64 {
	if x != nil {
		return x.RewardId
	}
	return 0
}

func (x *RewardStep) GetUnitType() int32 {
	if x != nil {
		return x.UnitType
	}
	return 0
}

func (x *RewardStep) GetConditionPrice() float32 {
	if x != nil {
		return x.ConditionPrice
	}
	return 0
}

func (x *RewardStep) GetConditionNum() int32 {
	if x != nil {
		return x.ConditionNum
	}
	return 0
}

func (x *RewardStep) GetPreferent() int32 {
	if x != nil {
		return x.Preferent
	}
	return 0
}

func (x *RewardStep) GetMoney() float32 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *RewardStep) GetDiscount() float32 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *RewardStep) GetPoint() int32 {
	if x != nil {
		return x.Point
	}
	return 0
}

func (x *RewardStep) GetPresentId() int64 {
	if x != nil {
		return x.PresentId
	}
	return 0
}

func (x *RewardStep) GetPresentNum() int32 {
	if x != nil {
		return x.PresentNum
	}
	return 0
}

func (x *RewardStep) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *RewardStep) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *RewardStep) GetCoupons() []*RewardCoupon {
	if x != nil {
		return x.Coupons
	}
	return nil
}

func (x *RewardStep) GetFreePostage() bool {
	if x != nil {
		return x.FreePostage
	}
	return false
}

type RewardCoupon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	RewardStepId int64  `protobuf:"varint,2,opt,name=reward_step_id,json=rewardStepId,proto3" json:"reward_step_id"`
	CouponId     int64  `protobuf:"varint,3,opt,name=coupon_id,json=couponId,proto3" json:"coupon_id"`
	CouponNum    int32  `protobuf:"varint,4,opt,name=coupon_num,json=couponNum,proto3" json:"coupon_num"`
	CreatedAt    string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt    string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *RewardCoupon) Reset() {
	*x = RewardCoupon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewardStepService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardCoupon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardCoupon) ProtoMessage() {}

func (x *RewardCoupon) ProtoReflect() protoreflect.Message {
	mi := &file_rewardStepService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardCoupon.ProtoReflect.Descriptor instead.
func (*RewardCoupon) Descriptor() ([]byte, []int) {
	return file_rewardStepService_proto_rawDescGZIP(), []int{1}
}

func (x *RewardCoupon) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RewardCoupon) GetRewardStepId() int64 {
	if x != nil {
		return x.RewardStepId
	}
	return 0
}

func (x *RewardCoupon) GetCouponId() int64 {
	if x != nil {
		return x.CouponId
	}
	return 0
}

func (x *RewardCoupon) GetCouponNum() int32 {
	if x != nil {
		return x.CouponNum
	}
	return 0
}

func (x *RewardCoupon) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *RewardCoupon) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

//
type RewardStepResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *RewardStep   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*RewardStep `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *RewardStepResponse) Reset() {
	*x = RewardStepResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewardStepService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardStepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardStepResponse) ProtoMessage() {}

func (x *RewardStepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rewardStepService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardStepResponse.ProtoReflect.Descriptor instead.
func (*RewardStepResponse) Descriptor() ([]byte, []int) {
	return file_rewardStepService_proto_rawDescGZIP(), []int{2}
}

func (x *RewardStepResponse) GetEntity() *RewardStep {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *RewardStepResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *RewardStepResponse) GetItems() []*RewardStep {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *RewardStepResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *RewardStepResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_rewardStepService_proto protoreflect.FileDescriptor

var file_rewardStepService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x74, 0x65, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x03, 0x0a, 0x0a, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x53, 0x74, 0x65, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x6f, 0x6e, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74,
	0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x30, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x18, 0x0e, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x66, 0x72, 0x65, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x22, 0xbe, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x74, 0x65, 0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xda, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x53, 0x74, 0x65, 0x70, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x53, 0x74, 0x65, 0x70, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rewardStepService_proto_rawDescOnce sync.Once
	file_rewardStepService_proto_rawDescData = file_rewardStepService_proto_rawDesc
)

func file_rewardStepService_proto_rawDescGZIP() []byte {
	file_rewardStepService_proto_rawDescOnce.Do(func() {
		file_rewardStepService_proto_rawDescData = protoimpl.X.CompressGZIP(file_rewardStepService_proto_rawDescData)
	})
	return file_rewardStepService_proto_rawDescData
}

var file_rewardStepService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rewardStepService_proto_goTypes = []interface{}{
	(*RewardStep)(nil),         // 0: services.RewardStep
	(*RewardCoupon)(nil),       // 1: services.RewardCoupon
	(*RewardStepResponse)(nil), // 2: services.RewardStepResponse
	(*common.Pager)(nil),       // 3: common.Pager
	(*common.Error)(nil),       // 4: common.Error
	(*common.Info)(nil),        // 5: common.Info
}
var file_rewardStepService_proto_depIdxs = []int32{
	1, // 0: services.RewardStep.coupons:type_name -> services.RewardCoupon
	0, // 1: services.RewardStepResponse.entity:type_name -> services.RewardStep
	3, // 2: services.RewardStepResponse.pager:type_name -> common.Pager
	0, // 3: services.RewardStepResponse.items:type_name -> services.RewardStep
	4, // 4: services.RewardStepResponse.error:type_name -> common.Error
	5, // 5: services.RewardStepResponse.info:type_name -> common.Info
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_rewardStepService_proto_init() }
func file_rewardStepService_proto_init() {
	if File_rewardStepService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rewardStepService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardStep); i {
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
		file_rewardStepService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardCoupon); i {
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
		file_rewardStepService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardStepResponse); i {
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
			RawDescriptor: file_rewardStepService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rewardStepService_proto_goTypes,
		DependencyIndexes: file_rewardStepService_proto_depIdxs,
		MessageInfos:      file_rewardStepService_proto_msgTypes,
	}.Build()
	File_rewardStepService_proto = out.File
	file_rewardStepService_proto_rawDesc = nil
	file_rewardStepService_proto_goTypes = nil
	file_rewardStepService_proto_depIdxs = nil
}
