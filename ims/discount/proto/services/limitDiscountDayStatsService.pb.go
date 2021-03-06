// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: limitDiscountDayStatsService.proto

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

type LimitDiscountDayStatsWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged           int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize        int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Top             int32  `protobuf:"varint,3,opt,name=top,proto3" json:"top"`
	Id              int64  `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	LimitDiscountId int64  `protobuf:"varint,5,opt,name=limit_discount_id,json=limitDiscountId,proto3" json:"limit_discount_id"`
	StartDate       string `protobuf:"bytes,6,opt,name=start_date,json=startDate,proto3" json:"start_date"`
	EndDate         string `protobuf:"bytes,7,opt,name=end_date,json=endDate,proto3" json:"end_date"`
}

func (x *LimitDiscountDayStatsWhere) Reset() {
	*x = LimitDiscountDayStatsWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_limitDiscountDayStatsService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitDiscountDayStatsWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitDiscountDayStatsWhere) ProtoMessage() {}

func (x *LimitDiscountDayStatsWhere) ProtoReflect() protoreflect.Message {
	mi := &file_limitDiscountDayStatsService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitDiscountDayStatsWhere.ProtoReflect.Descriptor instead.
func (*LimitDiscountDayStatsWhere) Descriptor() ([]byte, []int) {
	return file_limitDiscountDayStatsService_proto_rawDescGZIP(), []int{0}
}

func (x *LimitDiscountDayStatsWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *LimitDiscountDayStatsWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *LimitDiscountDayStatsWhere) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *LimitDiscountDayStatsWhere) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LimitDiscountDayStatsWhere) GetLimitDiscountId() int64 {
	if x != nil {
		return x.LimitDiscountId
	}
	return 0
}

func (x *LimitDiscountDayStatsWhere) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *LimitDiscountDayStatsWhere) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type LimitDiscountDayStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	LimitDiscountId int64          `protobuf:"varint,2,opt,name=limit_discount_id,json=limitDiscountId,proto3" json:"limit_discount_id"`
	PayAmount       float32        `protobuf:"fixed32,3,opt,name=pay_amount,json=payAmount,proto3" json:"pay_amount"`
	DiscountAmount  float32        `protobuf:"fixed32,4,opt,name=discount_amount,json=discountAmount,proto3" json:"discount_amount"`
	PayNum          int32          `protobuf:"varint,5,opt,name=pay_num,json=payNum,proto3" json:"pay_num"`
	BuyerNum        int32          `protobuf:"varint,6,opt,name=buyer_num,json=buyerNum,proto3" json:"buyer_num"`
	GoodsNum        int32          `protobuf:"varint,7,opt,name=goods_num,json=goodsNum,proto3" json:"goods_num"`
	NewCustomerNum  int32          `protobuf:"varint,8,opt,name=new_customer_num,json=newCustomerNum,proto3" json:"new_customer_num"`
	OldCustomerNum  int32          `protobuf:"varint,9,opt,name=old_customer_num,json=oldCustomerNum,proto3" json:"old_customer_num"`
	StatsDate       string         `protobuf:"bytes,10,opt,name=stats_date,json=statsDate,proto3" json:"stats_date"`
	CreatedAt       string         `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt       string         `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	LimitDiscount   *LimitDiscount `protobuf:"bytes,13,opt,name=limit_discount,json=limitDiscount,proto3" json:"limit_discount"`
}

func (x *LimitDiscountDayStats) Reset() {
	*x = LimitDiscountDayStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_limitDiscountDayStatsService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitDiscountDayStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitDiscountDayStats) ProtoMessage() {}

func (x *LimitDiscountDayStats) ProtoReflect() protoreflect.Message {
	mi := &file_limitDiscountDayStatsService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitDiscountDayStats.ProtoReflect.Descriptor instead.
func (*LimitDiscountDayStats) Descriptor() ([]byte, []int) {
	return file_limitDiscountDayStatsService_proto_rawDescGZIP(), []int{1}
}

func (x *LimitDiscountDayStats) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LimitDiscountDayStats) GetLimitDiscountId() int64 {
	if x != nil {
		return x.LimitDiscountId
	}
	return 0
}

func (x *LimitDiscountDayStats) GetPayAmount() float32 {
	if x != nil {
		return x.PayAmount
	}
	return 0
}

func (x *LimitDiscountDayStats) GetDiscountAmount() float32 {
	if x != nil {
		return x.DiscountAmount
	}
	return 0
}

func (x *LimitDiscountDayStats) GetPayNum() int32 {
	if x != nil {
		return x.PayNum
	}
	return 0
}

func (x *LimitDiscountDayStats) GetBuyerNum() int32 {
	if x != nil {
		return x.BuyerNum
	}
	return 0
}

func (x *LimitDiscountDayStats) GetGoodsNum() int32 {
	if x != nil {
		return x.GoodsNum
	}
	return 0
}

func (x *LimitDiscountDayStats) GetNewCustomerNum() int32 {
	if x != nil {
		return x.NewCustomerNum
	}
	return 0
}

func (x *LimitDiscountDayStats) GetOldCustomerNum() int32 {
	if x != nil {
		return x.OldCustomerNum
	}
	return 0
}

func (x *LimitDiscountDayStats) GetStatsDate() string {
	if x != nil {
		return x.StatsDate
	}
	return ""
}

func (x *LimitDiscountDayStats) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *LimitDiscountDayStats) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *LimitDiscountDayStats) GetLimitDiscount() *LimitDiscount {
	if x != nil {
		return x.LimitDiscount
	}
	return nil
}

type LimitDiscountDayStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *LimitDiscountDayStats   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager            `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*LimitDiscountDayStats `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error            `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info             `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *LimitDiscountDayStatsResponse) Reset() {
	*x = LimitDiscountDayStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_limitDiscountDayStatsService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitDiscountDayStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitDiscountDayStatsResponse) ProtoMessage() {}

func (x *LimitDiscountDayStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_limitDiscountDayStatsService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitDiscountDayStatsResponse.ProtoReflect.Descriptor instead.
func (*LimitDiscountDayStatsResponse) Descriptor() ([]byte, []int) {
	return file_limitDiscountDayStatsService_proto_rawDescGZIP(), []int{2}
}

func (x *LimitDiscountDayStatsResponse) GetEntity() *LimitDiscountDayStats {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *LimitDiscountDayStatsResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *LimitDiscountDayStatsResponse) GetItems() []*LimitDiscountDayStats {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *LimitDiscountDayStatsResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *LimitDiscountDayStatsResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_limitDiscountDayStatsService_proto protoreflect.FileDescriptor

var file_limitDiscountDayStatsService_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44,
	0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1a, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x01,
	0x0a, 0x1a, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44,
	0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f,
	0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0xdf, 0x03, 0x0a, 0x15, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x79, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x09, 0x70, 0x61, 0x79, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x5f, 0x6e, 0x75, 0x6d,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x79, 0x4e, 0x75, 0x6d, 0x12, 0x1b,
	0x0a, 0x09, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x62, 0x75, 0x79, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x65, 0x77, 0x5f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0e, 0x6e, 0x65, 0x77, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e,
	0x75, 0x6d, 0x12, 0x28, 0x0a, 0x10, 0x6f, 0x6c, 0x64, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6f, 0x6c,
	0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3e, 0x0a, 0x0e, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0d, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xfb, 0x01, 0x0a, 0x1d, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xd1, 0x01, 0x0a, 0x1c, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x27, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x24, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x57, 0x68, 0x65, 0x72, 0x65,
	0x1a, 0x27, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_limitDiscountDayStatsService_proto_rawDescOnce sync.Once
	file_limitDiscountDayStatsService_proto_rawDescData = file_limitDiscountDayStatsService_proto_rawDesc
)

func file_limitDiscountDayStatsService_proto_rawDescGZIP() []byte {
	file_limitDiscountDayStatsService_proto_rawDescOnce.Do(func() {
		file_limitDiscountDayStatsService_proto_rawDescData = protoimpl.X.CompressGZIP(file_limitDiscountDayStatsService_proto_rawDescData)
	})
	return file_limitDiscountDayStatsService_proto_rawDescData
}

var file_limitDiscountDayStatsService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_limitDiscountDayStatsService_proto_goTypes = []interface{}{
	(*LimitDiscountDayStatsWhere)(nil),    // 0: services.LimitDiscountDayStatsWhere
	(*LimitDiscountDayStats)(nil),         // 1: services.LimitDiscountDayStats
	(*LimitDiscountDayStatsResponse)(nil), // 2: services.LimitDiscountDayStatsResponse
	(*LimitDiscount)(nil),                 // 3: services.LimitDiscount
	(*common.Pager)(nil),                  // 4: common.Pager
	(*common.Error)(nil),                  // 5: common.Error
	(*common.Info)(nil),                   // 6: common.Info
}
var file_limitDiscountDayStatsService_proto_depIdxs = []int32{
	3, // 0: services.LimitDiscountDayStats.limit_discount:type_name -> services.LimitDiscount
	1, // 1: services.LimitDiscountDayStatsResponse.entity:type_name -> services.LimitDiscountDayStats
	4, // 2: services.LimitDiscountDayStatsResponse.pager:type_name -> common.Pager
	1, // 3: services.LimitDiscountDayStatsResponse.items:type_name -> services.LimitDiscountDayStats
	5, // 4: services.LimitDiscountDayStatsResponse.error:type_name -> common.Error
	6, // 5: services.LimitDiscountDayStatsResponse.info:type_name -> common.Info
	0, // 6: services.LimitDiscountDayStatsService.Search:input_type -> services.LimitDiscountDayStatsWhere
	0, // 7: services.LimitDiscountDayStatsService.Get:input_type -> services.LimitDiscountDayStatsWhere
	2, // 8: services.LimitDiscountDayStatsService.Search:output_type -> services.LimitDiscountDayStatsResponse
	2, // 9: services.LimitDiscountDayStatsService.Get:output_type -> services.LimitDiscountDayStatsResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_limitDiscountDayStatsService_proto_init() }
func file_limitDiscountDayStatsService_proto_init() {
	if File_limitDiscountDayStatsService_proto != nil {
		return
	}
	file_limitDiscountService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_limitDiscountDayStatsService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitDiscountDayStatsWhere); i {
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
		file_limitDiscountDayStatsService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitDiscountDayStats); i {
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
		file_limitDiscountDayStatsService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitDiscountDayStatsResponse); i {
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
			RawDescriptor: file_limitDiscountDayStatsService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_limitDiscountDayStatsService_proto_goTypes,
		DependencyIndexes: file_limitDiscountDayStatsService_proto_depIdxs,
		MessageInfos:      file_limitDiscountDayStatsService_proto_msgTypes,
	}.Build()
	File_limitDiscountDayStatsService_proto = out.File
	file_limitDiscountDayStatsService_proto_rawDesc = nil
	file_limitDiscountDayStatsService_proto_goTypes = nil
	file_limitDiscountDayStatsService_proto_depIdxs = nil
}
