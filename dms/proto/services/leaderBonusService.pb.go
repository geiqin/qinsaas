// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: leaderBonusService.proto

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

type LeaderBonus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	BonusSn        string  `protobuf:"bytes,2,opt,name=bonus_sn,json=bonusSn,proto3" json:"bonus_sn"`
	OrderId        int64   `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	OrderSn        string  `protobuf:"bytes,4,opt,name=order_sn,json=orderSn,proto3" json:"order_sn"`
	PlatformSource string  `protobuf:"bytes,5,opt,name=platform_source,json=platformSource,proto3" json:"platform_source"`
	OrderAmount    float32 `protobuf:"fixed32,6,opt,name=order_amount,json=orderAmount,proto3" json:"order_amount"`
	LeaderId       int64   `protobuf:"varint,7,opt,name=leader_id,json=leaderId,proto3" json:"leader_id"`
	Money          float32 `protobuf:"fixed32,8,opt,name=money,proto3" json:"money"`
	IncomeRate     float32 `protobuf:"fixed32,9,opt,name=income_rate,json=incomeRate,proto3" json:"income_rate"`
	Memo           string  `protobuf:"bytes,10,opt,name=memo,proto3" json:"memo"`
	OperatorId     int64   `protobuf:"varint,11,opt,name=operator_id,json=operatorId,proto3" json:"operator_id"`
	Status         int32   `protobuf:"varint,12,opt,name=status,proto3" json:"status"`
	CreatedAt      string  `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt      string  `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	// @inject_tag: gorm:"-"
	Leader *Leader `protobuf:"bytes,15,opt,name=leader,proto3" json:"leader"`
}

func (x *LeaderBonus) Reset() {
	*x = LeaderBonus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderBonusService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderBonus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderBonus) ProtoMessage() {}

func (x *LeaderBonus) ProtoReflect() protoreflect.Message {
	mi := &file_leaderBonusService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderBonus.ProtoReflect.Descriptor instead.
func (*LeaderBonus) Descriptor() ([]byte, []int) {
	return file_leaderBonusService_proto_rawDescGZIP(), []int{0}
}

func (x *LeaderBonus) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LeaderBonus) GetBonusSn() string {
	if x != nil {
		return x.BonusSn
	}
	return ""
}

func (x *LeaderBonus) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *LeaderBonus) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

func (x *LeaderBonus) GetPlatformSource() string {
	if x != nil {
		return x.PlatformSource
	}
	return ""
}

func (x *LeaderBonus) GetOrderAmount() float32 {
	if x != nil {
		return x.OrderAmount
	}
	return 0
}

func (x *LeaderBonus) GetLeaderId() int64 {
	if x != nil {
		return x.LeaderId
	}
	return 0
}

func (x *LeaderBonus) GetMoney() float32 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *LeaderBonus) GetIncomeRate() float32 {
	if x != nil {
		return x.IncomeRate
	}
	return 0
}

func (x *LeaderBonus) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *LeaderBonus) GetOperatorId() int64 {
	if x != nil {
		return x.OperatorId
	}
	return 0
}

func (x *LeaderBonus) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *LeaderBonus) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *LeaderBonus) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *LeaderBonus) GetLeader() *Leader {
	if x != nil {
		return x.Leader
	}
	return nil
}

type LeaderBonusWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged     int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Id        int64  `protobuf:"varint,3,opt,name=id,proto3" json:"id"`
	Mobile    string `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile"`
	OrderId   int64  `protobuf:"varint,5,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	OrderSn   string `protobuf:"bytes,6,opt,name=order_sn,json=orderSn,proto3" json:"order_sn"`
	Keywords  string `protobuf:"bytes,7,opt,name=keywords,proto3" json:"keywords"`
	Status    int32  `protobuf:"varint,8,opt,name=status,proto3" json:"status"`
	StartDate string `protobuf:"bytes,9,opt,name=start_date,json=startDate,proto3" json:"start_date"`
	EndDate   string `protobuf:"bytes,10,opt,name=end_date,json=endDate,proto3" json:"end_date"`
	// @inject_tag: gorm:"-"
	Ids      []int64 `protobuf:"varint,11,rep,packed,name=ids,proto3" json:"ids"`
	LeaderId int64   `protobuf:"varint,12,opt,name=leader_id,json=leaderId,proto3" json:"leader_id"`
}

func (x *LeaderBonusWhere) Reset() {
	*x = LeaderBonusWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderBonusService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderBonusWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderBonusWhere) ProtoMessage() {}

func (x *LeaderBonusWhere) ProtoReflect() protoreflect.Message {
	mi := &file_leaderBonusService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderBonusWhere.ProtoReflect.Descriptor instead.
func (*LeaderBonusWhere) Descriptor() ([]byte, []int) {
	return file_leaderBonusService_proto_rawDescGZIP(), []int{1}
}

func (x *LeaderBonusWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *LeaderBonusWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *LeaderBonusWhere) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LeaderBonusWhere) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *LeaderBonusWhere) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *LeaderBonusWhere) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

func (x *LeaderBonusWhere) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *LeaderBonusWhere) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *LeaderBonusWhere) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *LeaderBonusWhere) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *LeaderBonusWhere) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *LeaderBonusWhere) GetLeaderId() int64 {
	if x != nil {
		return x.LeaderId
	}
	return 0
}

type LeaderPerformance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	CustomerId    int64   `protobuf:"varint,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	DisplayName   string  `protobuf:"bytes,5,opt,name=display_name,json=displayName,proto3" json:"display_name"`
	RankId        int32   `protobuf:"varint,6,opt,name=rank_id,json=rankId,proto3" json:"rank_id"`
	BonusWithdrew float32 `protobuf:"fixed32,14,opt,name=bonus_withdrew,json=bonusWithdrew,proto3" json:"bonus_withdrew"`
	Mobile        string  `protobuf:"bytes,15,opt,name=mobile,proto3" json:"mobile"`
	RankName      string  `protobuf:"bytes,16,opt,name=rank_name,json=rankName,proto3" json:"rank_name"`
}

func (x *LeaderPerformance) Reset() {
	*x = LeaderPerformance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderBonusService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderPerformance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderPerformance) ProtoMessage() {}

func (x *LeaderPerformance) ProtoReflect() protoreflect.Message {
	mi := &file_leaderBonusService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderPerformance.ProtoReflect.Descriptor instead.
func (*LeaderPerformance) Descriptor() ([]byte, []int) {
	return file_leaderBonusService_proto_rawDescGZIP(), []int{2}
}

func (x *LeaderPerformance) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LeaderPerformance) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *LeaderPerformance) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *LeaderPerformance) GetRankId() int32 {
	if x != nil {
		return x.RankId
	}
	return 0
}

func (x *LeaderPerformance) GetBonusWithdrew() float32 {
	if x != nil {
		return x.BonusWithdrew
	}
	return 0
}

func (x *LeaderPerformance) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *LeaderPerformance) GetRankName() string {
	if x != nil {
		return x.RankName
	}
	return ""
}

type LeaderPerformanceWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	Mobile   string `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile"`
	RankId   int32  `protobuf:"varint,5,opt,name=rank_id,json=rankId,proto3" json:"rank_id"`
	Keywords string `protobuf:"bytes,6,opt,name=keywords,proto3" json:"keywords"`
	Id       int64  `protobuf:"varint,7,opt,name=id,proto3" json:"id"`
	// @inject_tag: gorm:"-"
	Ids []int64 `protobuf:"varint,8,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *LeaderPerformanceWhere) Reset() {
	*x = LeaderPerformanceWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderBonusService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderPerformanceWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderPerformanceWhere) ProtoMessage() {}

func (x *LeaderPerformanceWhere) ProtoReflect() protoreflect.Message {
	mi := &file_leaderBonusService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderPerformanceWhere.ProtoReflect.Descriptor instead.
func (*LeaderPerformanceWhere) Descriptor() ([]byte, []int) {
	return file_leaderBonusService_proto_rawDescGZIP(), []int{3}
}

func (x *LeaderPerformanceWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *LeaderPerformanceWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *LeaderPerformanceWhere) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *LeaderPerformanceWhere) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *LeaderPerformanceWhere) GetRankId() int32 {
	if x != nil {
		return x.RankId
	}
	return 0
}

func (x *LeaderPerformanceWhere) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *LeaderPerformanceWhere) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LeaderPerformanceWhere) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type LeaderBonusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *LeaderBonus   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager  `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*LeaderBonus `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error  `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info   `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *LeaderBonusResponse) Reset() {
	*x = LeaderBonusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderBonusService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderBonusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderBonusResponse) ProtoMessage() {}

func (x *LeaderBonusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_leaderBonusService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderBonusResponse.ProtoReflect.Descriptor instead.
func (*LeaderBonusResponse) Descriptor() ([]byte, []int) {
	return file_leaderBonusService_proto_rawDescGZIP(), []int{4}
}

func (x *LeaderBonusResponse) GetEntity() *LeaderBonus {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *LeaderBonusResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *LeaderBonusResponse) GetItems() []*LeaderBonus {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *LeaderBonusResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *LeaderBonusResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type LeaderPerformanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *LeaderPerformance   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager        `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*LeaderPerformance `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error        `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info         `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *LeaderPerformanceResponse) Reset() {
	*x = LeaderPerformanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderBonusService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderPerformanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderPerformanceResponse) ProtoMessage() {}

func (x *LeaderPerformanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_leaderBonusService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderPerformanceResponse.ProtoReflect.Descriptor instead.
func (*LeaderPerformanceResponse) Descriptor() ([]byte, []int) {
	return file_leaderBonusService_proto_rawDescGZIP(), []int{5}
}

func (x *LeaderPerformanceResponse) GetEntity() *LeaderPerformance {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *LeaderPerformanceResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *LeaderPerformanceResponse) GetItems() []*LeaderPerformance {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *LeaderPerformanceResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *LeaderPerformanceResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_leaderBonusService_proto protoreflect.FileDescriptor

var file_leaderBonusService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x03, 0x0a,
	0x0b, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x73, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x62, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x12, 0x27, 0x0a,
	0x0f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b,
	0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0a, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d,
	0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x22, 0xc0, 0x02, 0x0a, 0x10, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6e,
	0x75, 0x73, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0xdc, 0x01, 0x0a, 0x11, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x72, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x6f, 0x6e, 0x75,
	0x73, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x65, 0x77, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0d, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x65, 0x77, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x61, 0x6e, 0x6b, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x61, 0x6e, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0xd4, 0x01, 0x0a, 0x16, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50,
	0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xdd, 0x01, 0x0a, 0x13,
	0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0xef, 0x01, 0x0a, 0x19,
	0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x65, 0x72, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23,
	0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xe5, 0x01,
	0x0a, 0x12, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6e,
	0x75, 0x73, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42,
	0x6f, 0x6e, 0x75, 0x73, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6e, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0a, 0x53, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x57,
	0x68, 0x65, 0x72, 0x65, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_leaderBonusService_proto_rawDescOnce sync.Once
	file_leaderBonusService_proto_rawDescData = file_leaderBonusService_proto_rawDesc
)

func file_leaderBonusService_proto_rawDescGZIP() []byte {
	file_leaderBonusService_proto_rawDescOnce.Do(func() {
		file_leaderBonusService_proto_rawDescData = protoimpl.X.CompressGZIP(file_leaderBonusService_proto_rawDescData)
	})
	return file_leaderBonusService_proto_rawDescData
}

var file_leaderBonusService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_leaderBonusService_proto_goTypes = []interface{}{
	(*LeaderBonus)(nil),               // 0: services.LeaderBonus
	(*LeaderBonusWhere)(nil),          // 1: services.LeaderBonusWhere
	(*LeaderPerformance)(nil),         // 2: services.LeaderPerformance
	(*LeaderPerformanceWhere)(nil),    // 3: services.LeaderPerformanceWhere
	(*LeaderBonusResponse)(nil),       // 4: services.LeaderBonusResponse
	(*LeaderPerformanceResponse)(nil), // 5: services.LeaderPerformanceResponse
	(*Leader)(nil),                    // 6: services.Leader
	(*common.Pager)(nil),              // 7: common.Pager
	(*common.Error)(nil),              // 8: common.Error
	(*common.Info)(nil),               // 9: common.Info
}
var file_leaderBonusService_proto_depIdxs = []int32{
	6,  // 0: services.LeaderBonus.leader:type_name -> services.Leader
	0,  // 1: services.LeaderBonusResponse.entity:type_name -> services.LeaderBonus
	7,  // 2: services.LeaderBonusResponse.pager:type_name -> common.Pager
	0,  // 3: services.LeaderBonusResponse.items:type_name -> services.LeaderBonus
	8,  // 4: services.LeaderBonusResponse.error:type_name -> common.Error
	9,  // 5: services.LeaderBonusResponse.info:type_name -> common.Info
	2,  // 6: services.LeaderPerformanceResponse.entity:type_name -> services.LeaderPerformance
	7,  // 7: services.LeaderPerformanceResponse.pager:type_name -> common.Pager
	2,  // 8: services.LeaderPerformanceResponse.items:type_name -> services.LeaderPerformance
	8,  // 9: services.LeaderPerformanceResponse.error:type_name -> common.Error
	9,  // 10: services.LeaderPerformanceResponse.info:type_name -> common.Info
	0,  // 11: services.LeaderBonusService.Get:input_type -> services.LeaderBonus
	1,  // 12: services.LeaderBonusService.Search:input_type -> services.LeaderBonusWhere
	1,  // 13: services.LeaderBonusService.Settlement:input_type -> services.LeaderBonusWhere
	4,  // 14: services.LeaderBonusService.Get:output_type -> services.LeaderBonusResponse
	4,  // 15: services.LeaderBonusService.Search:output_type -> services.LeaderBonusResponse
	4,  // 16: services.LeaderBonusService.Settlement:output_type -> services.LeaderBonusResponse
	14, // [14:17] is the sub-list for method output_type
	11, // [11:14] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_leaderBonusService_proto_init() }
func file_leaderBonusService_proto_init() {
	if File_leaderBonusService_proto != nil {
		return
	}
	file_leaderService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_leaderBonusService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderBonus); i {
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
		file_leaderBonusService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderBonusWhere); i {
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
		file_leaderBonusService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderPerformance); i {
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
		file_leaderBonusService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderPerformanceWhere); i {
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
		file_leaderBonusService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderBonusResponse); i {
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
		file_leaderBonusService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderPerformanceResponse); i {
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
			RawDescriptor: file_leaderBonusService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_leaderBonusService_proto_goTypes,
		DependencyIndexes: file_leaderBonusService_proto_depIdxs,
		MessageInfos:      file_leaderBonusService_proto_msgTypes,
	}.Build()
	File_leaderBonusService_proto = out.File
	file_leaderBonusService_proto_rawDesc = nil
	file_leaderBonusService_proto_goTypes = nil
	file_leaderBonusService_proto_depIdxs = nil
}
