// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: groupCustomerStatsDayService.proto

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

type GroupCustomerStatsDay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	GroupId        int64         `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id"`
	CustomerId     int64         `protobuf:"varint,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	PayAmount      float32       `protobuf:"fixed32,4,opt,name=pay_amount,json=payAmount,proto3" json:"pay_amount"`
	DiscountAmount float32       `protobuf:"fixed32,5,opt,name=discount_amount,json=discountAmount,proto3" json:"discount_amount"`
	PayNum         int32         `protobuf:"varint,6,opt,name=pay_num,json=payNum,proto3" json:"pay_num"`
	GoodsNum       int32         `protobuf:"varint,7,opt,name=goods_num,json=goodsNum,proto3" json:"goods_num"`
	StatsDate      string        `protobuf:"bytes,8,opt,name=stats_date,json=statsDate,proto3" json:"stats_date"`
	CreatedAt      string        `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt      string        `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Group          *Group        `protobuf:"bytes,11,opt,name=group,proto3" json:"group"`
	Customer       *CustomerInfo `protobuf:"bytes,12,opt,name=customer,proto3" json:"customer"`
}

func (x *GroupCustomerStatsDay) Reset() {
	*x = GroupCustomerStatsDay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupCustomerStatsDayService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupCustomerStatsDay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupCustomerStatsDay) ProtoMessage() {}

func (x *GroupCustomerStatsDay) ProtoReflect() protoreflect.Message {
	mi := &file_groupCustomerStatsDayService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupCustomerStatsDay.ProtoReflect.Descriptor instead.
func (*GroupCustomerStatsDay) Descriptor() ([]byte, []int) {
	return file_groupCustomerStatsDayService_proto_rawDescGZIP(), []int{0}
}

func (x *GroupCustomerStatsDay) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GroupCustomerStatsDay) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GroupCustomerStatsDay) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *GroupCustomerStatsDay) GetPayAmount() float32 {
	if x != nil {
		return x.PayAmount
	}
	return 0
}

func (x *GroupCustomerStatsDay) GetDiscountAmount() float32 {
	if x != nil {
		return x.DiscountAmount
	}
	return 0
}

func (x *GroupCustomerStatsDay) GetPayNum() int32 {
	if x != nil {
		return x.PayNum
	}
	return 0
}

func (x *GroupCustomerStatsDay) GetGoodsNum() int32 {
	if x != nil {
		return x.GoodsNum
	}
	return 0
}

func (x *GroupCustomerStatsDay) GetStatsDate() string {
	if x != nil {
		return x.StatsDate
	}
	return ""
}

func (x *GroupCustomerStatsDay) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GroupCustomerStatsDay) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *GroupCustomerStatsDay) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *GroupCustomerStatsDay) GetCustomer() *CustomerInfo {
	if x != nil {
		return x.Customer
	}
	return nil
}

type GroupCustomerStatsDayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *GroupCustomerStatsDay   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager            `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*GroupCustomerStatsDay `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error            `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info             `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *GroupCustomerStatsDayResponse) Reset() {
	*x = GroupCustomerStatsDayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupCustomerStatsDayService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupCustomerStatsDayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupCustomerStatsDayResponse) ProtoMessage() {}

func (x *GroupCustomerStatsDayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_groupCustomerStatsDayService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupCustomerStatsDayResponse.ProtoReflect.Descriptor instead.
func (*GroupCustomerStatsDayResponse) Descriptor() ([]byte, []int) {
	return file_groupCustomerStatsDayService_proto_rawDescGZIP(), []int{1}
}

func (x *GroupCustomerStatsDayResponse) GetEntity() *GroupCustomerStatsDay {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *GroupCustomerStatsDayResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *GroupCustomerStatsDayResponse) GetItems() []*GroupCustomerStatsDay {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *GroupCustomerStatsDayResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *GroupCustomerStatsDayResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_groupCustomerStatsDayService_proto protoreflect.FileDescriptor

var file_groupCustomerStatsDayService_proto_rawDesc = []byte{
	0x0a, 0x22, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x03, 0x0a, 0x15, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x44, 0x61, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x61, 0x79, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27,
	0x0a, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x79, 0x4e, 0x75, 0x6d,
	0x12, 0x1b, 0x0a, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x32, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0xfb, 0x01, 0x0a, 0x1d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x44, 0x61, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x32, 0x89, 0x02, 0x0a, 0x1c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x27, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x27, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4c, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x73, 0x57,
	0x68, 0x65, 0x72, 0x65, 0x1a, 0x27, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x44, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_groupCustomerStatsDayService_proto_rawDescOnce sync.Once
	file_groupCustomerStatsDayService_proto_rawDescData = file_groupCustomerStatsDayService_proto_rawDesc
)

func file_groupCustomerStatsDayService_proto_rawDescGZIP() []byte {
	file_groupCustomerStatsDayService_proto_rawDescOnce.Do(func() {
		file_groupCustomerStatsDayService_proto_rawDescData = protoimpl.X.CompressGZIP(file_groupCustomerStatsDayService_proto_rawDescData)
	})
	return file_groupCustomerStatsDayService_proto_rawDescData
}

var file_groupCustomerStatsDayService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_groupCustomerStatsDayService_proto_goTypes = []interface{}{
	(*GroupCustomerStatsDay)(nil),         // 0: services.GroupCustomerStatsDay
	(*GroupCustomerStatsDayResponse)(nil), // 1: services.GroupCustomerStatsDayResponse
	(*Group)(nil),                         // 2: services.Group
	(*CustomerInfo)(nil),                  // 3: services.CustomerInfo
	(*common.Pager)(nil),                  // 4: common.Pager
	(*common.Error)(nil),                  // 5: common.Error
	(*common.Info)(nil),                   // 6: common.Info
	(*GroupStatsWhere)(nil),               // 7: services.GroupStatsWhere
}
var file_groupCustomerStatsDayService_proto_depIdxs = []int32{
	2,  // 0: services.GroupCustomerStatsDay.group:type_name -> services.Group
	3,  // 1: services.GroupCustomerStatsDay.customer:type_name -> services.CustomerInfo
	0,  // 2: services.GroupCustomerStatsDayResponse.entity:type_name -> services.GroupCustomerStatsDay
	4,  // 3: services.GroupCustomerStatsDayResponse.pager:type_name -> common.Pager
	0,  // 4: services.GroupCustomerStatsDayResponse.items:type_name -> services.GroupCustomerStatsDay
	5,  // 5: services.GroupCustomerStatsDayResponse.error:type_name -> common.Error
	6,  // 6: services.GroupCustomerStatsDayResponse.info:type_name -> common.Info
	7,  // 7: services.GroupCustomerStatsDayService.Search:input_type -> services.GroupStatsWhere
	7,  // 8: services.GroupCustomerStatsDayService.Get:input_type -> services.GroupStatsWhere
	7,  // 9: services.GroupCustomerStatsDayService.List:input_type -> services.GroupStatsWhere
	1,  // 10: services.GroupCustomerStatsDayService.Search:output_type -> services.GroupCustomerStatsDayResponse
	1,  // 11: services.GroupCustomerStatsDayService.Get:output_type -> services.GroupCustomerStatsDayResponse
	1,  // 12: services.GroupCustomerStatsDayService.List:output_type -> services.GroupCustomerStatsDayResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_groupCustomerStatsDayService_proto_init() }
func file_groupCustomerStatsDayService_proto_init() {
	if File_groupCustomerStatsDayService_proto != nil {
		return
	}
	file_groupService_proto_init()
	file_customerInfoService_proto_init()
	file_groupStatsService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_groupCustomerStatsDayService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupCustomerStatsDay); i {
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
		file_groupCustomerStatsDayService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupCustomerStatsDayResponse); i {
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
			RawDescriptor: file_groupCustomerStatsDayService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_groupCustomerStatsDayService_proto_goTypes,
		DependencyIndexes: file_groupCustomerStatsDayService_proto_depIdxs,
		MessageInfos:      file_groupCustomerStatsDayService_proto_msgTypes,
	}.Build()
	File_groupCustomerStatsDayService_proto = out.File
	file_groupCustomerStatsDayService_proto_rawDesc = nil
	file_groupCustomerStatsDayService_proto_goTypes = nil
	file_groupCustomerStatsDayService_proto_depIdxs = nil
}