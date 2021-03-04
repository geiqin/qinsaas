// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: cardService.proto

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

type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name          string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	CardBg        string         `protobuf:"bytes,3,opt,name=card_bg,json=cardBg,proto3" json:"card_bg"`
	Term          int32          `protobuf:"varint,4,opt,name=term,proto3" json:"term"`
	TermDays      int32          `protobuf:"varint,5,opt,name=term_days,json=termDays,proto3" json:"term_days"`
	TermStartDate string         `protobuf:"bytes,6,opt,name=term_start_date,json=termStartDate,proto3" json:"term_start_date"`
	TermEndDate   string         `protobuf:"bytes,7,opt,name=term_end_date,json=termEndDate,proto3" json:"term_end_date"`
	Method        int32          `protobuf:"varint,8,opt,name=method,proto3" json:"method"`
	Fee           float32        `protobuf:"fixed32,9,opt,name=fee,proto3" json:"fee"`
	Stock         int32          `protobuf:"varint,10,opt,name=stock,proto3" json:"stock"`
	Code          string         `protobuf:"bytes,11,opt,name=code,proto3" json:"code"`
	TotalPayment  int32          `protobuf:"varint,12,opt,name=total_payment,json=totalPayment,proto3" json:"total_payment"`
	TotalConsume  float32        `protobuf:"fixed32,13,opt,name=total_consume,json=totalConsume,proto3" json:"total_consume"`
	TotalIntegral int32          `protobuf:"varint,14,opt,name=total_integral,json=totalIntegral,proto3" json:"total_integral"`
	Description   string         `protobuf:"bytes,15,opt,name=description,proto3" json:"description"`
	Phone         string         `protobuf:"bytes,16,opt,name=phone,proto3" json:"phone"`
	Disabled      bool           `protobuf:"varint,17,opt,name=disabled,proto3" json:"disabled"`
	CreatedAt     string         `protobuf:"bytes,19,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt     string         `protobuf:"bytes,20,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	CardBenefits  []*CardBenefit `protobuf:"bytes,21,rep,name=card_benefits,json=cardBenefits,proto3" json:"card_benefits"`
	// @inject_tag: gorm:"-"
	Ids    []int32 `protobuf:"varint,22,rep,packed,name=ids,proto3" json:"ids" gorm:"-"`
	Status int32   `protobuf:"varint,23,opt,name=status,proto3" json:"status"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cardService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_cardService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_cardService_proto_rawDescGZIP(), []int{0}
}

func (x *Card) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Card) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Card) GetCardBg() string {
	if x != nil {
		return x.CardBg
	}
	return ""
}

func (x *Card) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *Card) GetTermDays() int32 {
	if x != nil {
		return x.TermDays
	}
	return 0
}

func (x *Card) GetTermStartDate() string {
	if x != nil {
		return x.TermStartDate
	}
	return ""
}

func (x *Card) GetTermEndDate() string {
	if x != nil {
		return x.TermEndDate
	}
	return ""
}

func (x *Card) GetMethod() int32 {
	if x != nil {
		return x.Method
	}
	return 0
}

func (x *Card) GetFee() float32 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *Card) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *Card) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Card) GetTotalPayment() int32 {
	if x != nil {
		return x.TotalPayment
	}
	return 0
}

func (x *Card) GetTotalConsume() float32 {
	if x != nil {
		return x.TotalConsume
	}
	return 0
}

func (x *Card) GetTotalIntegral() int32 {
	if x != nil {
		return x.TotalIntegral
	}
	return 0
}

func (x *Card) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Card) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Card) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *Card) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Card) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Card) GetCardBenefits() []*CardBenefit {
	if x != nil {
		return x.CardBenefits
	}
	return nil
}

func (x *Card) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *Card) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

//查询参数
type CardWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	//以下为自定义参数
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	Status int32  `protobuf:"varint,4,opt,name=status,proto3" json:"status"` //状态 1使用中   2 已禁用  3已过期
	Method int32  `protobuf:"varint,5,opt,name=method,proto3" json:"method"` //领取方式：1 直接领取   2 付费购买
}

func (x *CardWhere) Reset() {
	*x = CardWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cardService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardWhere) ProtoMessage() {}

func (x *CardWhere) ProtoReflect() protoreflect.Message {
	mi := &file_cardService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardWhere.ProtoReflect.Descriptor instead.
func (*CardWhere) Descriptor() ([]byte, []int) {
	return file_cardService_proto_rawDescGZIP(), []int{1}
}

func (x *CardWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *CardWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *CardWhere) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CardWhere) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CardWhere) GetMethod() int32 {
	if x != nil {
		return x.Method
	}
	return 0
}

//
type CardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Card         `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Card       `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *CardResponse) Reset() {
	*x = CardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cardService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardResponse) ProtoMessage() {}

func (x *CardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cardService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardResponse.ProtoReflect.Descriptor instead.
func (*CardResponse) Descriptor() ([]byte, []int) {
	return file_cardService_proto_rawDescGZIP(), []int{2}
}

func (x *CardResponse) GetEntity() *Card {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *CardResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *CardResponse) GetItems() []*Card {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CardResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *CardResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_cardService_proto protoreflect.FileDescriptor

var file_cardService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x63, 0x61, 0x72, 0x64, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x04, 0x0a, 0x04, 0x43,
	0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x64, 0x5f,
	0x62, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x72, 0x64, 0x42, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x74, 0x65, 0x72, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x72, 0x6d, 0x5f, 0x64, 0x61, 0x79,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x65, 0x72, 0x6d, 0x44, 0x61, 0x79,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x65, 0x72, 0x6d, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x65, 0x72, 0x6d,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x65, 0x72,
	0x6d, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x74, 0x65, 0x72, 0x6d, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x3a, 0x0a, 0x0d, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x62, 0x65, 0x6e,
	0x65, 0x66, 0x69, 0x74, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x42, 0x65, 0x6e, 0x65, 0x66,
	0x69, 0x74, 0x52, 0x0c, 0x63, 0x61, 0x72, 0x64, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x16, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x09, 0x43,
	0x61, 0x72, 0x64, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22,
	0xc8, 0x01, 0x0a, 0x0c, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x64,
	0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x24, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xec, 0x03, 0x0a, 0x0b, 0x43,
	0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x61, 0x72, 0x64, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x1a, 0x16, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61,
	0x72, 0x64, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x07,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x2f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x35, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a,
	0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x09, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a,
	0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cardService_proto_rawDescOnce sync.Once
	file_cardService_proto_rawDescData = file_cardService_proto_rawDesc
)

func file_cardService_proto_rawDescGZIP() []byte {
	file_cardService_proto_rawDescOnce.Do(func() {
		file_cardService_proto_rawDescData = protoimpl.X.CompressGZIP(file_cardService_proto_rawDescData)
	})
	return file_cardService_proto_rawDescData
}

var file_cardService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cardService_proto_goTypes = []interface{}{
	(*Card)(nil),         // 0: services.Card
	(*CardWhere)(nil),    // 1: services.CardWhere
	(*CardResponse)(nil), // 2: services.CardResponse
	(*CardBenefit)(nil),  // 3: services.CardBenefit
	(*common.Pager)(nil), // 4: common.Pager
	(*common.Error)(nil), // 5: common.Error
	(*common.Info)(nil),  // 6: common.Info
}
var file_cardService_proto_depIdxs = []int32{
	3,  // 0: services.Card.card_benefits:type_name -> services.CardBenefit
	0,  // 1: services.CardResponse.entity:type_name -> services.Card
	4,  // 2: services.CardResponse.pager:type_name -> common.Pager
	0,  // 3: services.CardResponse.items:type_name -> services.Card
	5,  // 4: services.CardResponse.error:type_name -> common.Error
	6,  // 5: services.CardResponse.info:type_name -> common.Info
	0,  // 6: services.CardService.Create:input_type -> services.Card
	0,  // 7: services.CardService.Update:input_type -> services.Card
	0,  // 8: services.CardService.Delete:input_type -> services.Card
	0,  // 9: services.CardService.Disabled:input_type -> services.Card
	0,  // 10: services.CardService.Enabled:input_type -> services.Card
	0,  // 11: services.CardService.Get:input_type -> services.Card
	1,  // 12: services.CardService.List:input_type -> services.CardWhere
	0,  // 13: services.CardService.ValidList:input_type -> services.Card
	1,  // 14: services.CardService.Search:input_type -> services.CardWhere
	2,  // 15: services.CardService.Create:output_type -> services.CardResponse
	2,  // 16: services.CardService.Update:output_type -> services.CardResponse
	2,  // 17: services.CardService.Delete:output_type -> services.CardResponse
	2,  // 18: services.CardService.Disabled:output_type -> services.CardResponse
	2,  // 19: services.CardService.Enabled:output_type -> services.CardResponse
	2,  // 20: services.CardService.Get:output_type -> services.CardResponse
	2,  // 21: services.CardService.List:output_type -> services.CardResponse
	2,  // 22: services.CardService.ValidList:output_type -> services.CardResponse
	2,  // 23: services.CardService.Search:output_type -> services.CardResponse
	15, // [15:24] is the sub-list for method output_type
	6,  // [6:15] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_cardService_proto_init() }
func file_cardService_proto_init() {
	if File_cardService_proto != nil {
		return
	}
	file_cardBenefitService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cardService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
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
		file_cardService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardWhere); i {
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
		file_cardService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardResponse); i {
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
			RawDescriptor: file_cardService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cardService_proto_goTypes,
		DependencyIndexes: file_cardService_proto_depIdxs,
		MessageInfos:      file_cardService_proto_msgTypes,
	}.Build()
	File_cardService_proto = out.File
	file_cardService_proto_rawDesc = nil
	file_cardService_proto_goTypes = nil
	file_cardService_proto_depIdxs = nil
}
