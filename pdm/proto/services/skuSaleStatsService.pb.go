// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: skuSaleStatsService.proto

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

type SkuSaleStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ItemId      int64   `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SkuId       int64   `protobuf:"varint,3,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	SaleNum     int32   `protobuf:"varint,4,opt,name=sale_num,json=saleNum,proto3" json:"sale_num,omitempty"`
	UnitPrice   float32 `protobuf:"fixed32,5,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
	SaleAmount  float32 `protobuf:"fixed32,6,opt,name=sale_amount,json=saleAmount,proto3" json:"sale_amount,omitempty"`
	CostAmount  float32 `protobuf:"fixed32,7,opt,name=cost_amount,json=costAmount,proto3" json:"cost_amount,omitempty"`
	Profit      float32 `protobuf:"fixed32,8,opt,name=profit,proto3" json:"profit,omitempty"`
	ProfitRate  float32 `protobuf:"fixed32,9,opt,name=profit_rate,json=profitRate,proto3" json:"profit_rate,omitempty"`
	LastStatsAt string  `protobuf:"bytes,10,opt,name=last_stats_at,json=lastStatsAt,proto3" json:"last_stats_at,omitempty"`
	CreatedAt   string  `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string  `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *SkuSaleStats) Reset() {
	*x = SkuSaleStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skuSaleStatsService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkuSaleStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkuSaleStats) ProtoMessage() {}

func (x *SkuSaleStats) ProtoReflect() protoreflect.Message {
	mi := &file_skuSaleStatsService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkuSaleStats.ProtoReflect.Descriptor instead.
func (*SkuSaleStats) Descriptor() ([]byte, []int) {
	return file_skuSaleStatsService_proto_rawDescGZIP(), []int{0}
}

func (x *SkuSaleStats) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SkuSaleStats) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *SkuSaleStats) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *SkuSaleStats) GetSaleNum() int32 {
	if x != nil {
		return x.SaleNum
	}
	return 0
}

func (x *SkuSaleStats) GetUnitPrice() float32 {
	if x != nil {
		return x.UnitPrice
	}
	return 0
}

func (x *SkuSaleStats) GetSaleAmount() float32 {
	if x != nil {
		return x.SaleAmount
	}
	return 0
}

func (x *SkuSaleStats) GetCostAmount() float32 {
	if x != nil {
		return x.CostAmount
	}
	return 0
}

func (x *SkuSaleStats) GetProfit() float32 {
	if x != nil {
		return x.Profit
	}
	return 0
}

func (x *SkuSaleStats) GetProfitRate() float32 {
	if x != nil {
		return x.ProfitRate
	}
	return 0
}

func (x *SkuSaleStats) GetLastStatsAt() string {
	if x != nil {
		return x.LastStatsAt
	}
	return ""
}

func (x *SkuSaleStats) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SkuSaleStats) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type SkuSaleStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *SkuSaleStats   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*SkuSaleStats `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *common.Error   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info    `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SkuSaleStatsResponse) Reset() {
	*x = SkuSaleStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skuSaleStatsService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkuSaleStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkuSaleStatsResponse) ProtoMessage() {}

func (x *SkuSaleStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_skuSaleStatsService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkuSaleStatsResponse.ProtoReflect.Descriptor instead.
func (*SkuSaleStatsResponse) Descriptor() ([]byte, []int) {
	return file_skuSaleStatsService_proto_rawDescGZIP(), []int{1}
}

func (x *SkuSaleStatsResponse) GetEntity() *SkuSaleStats {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *SkuSaleStatsResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *SkuSaleStatsResponse) GetItems() []*SkuSaleStats {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *SkuSaleStatsResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SkuSaleStatsResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type SkuSaleDayStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ItemId        int64   `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	SkuId         int64   `protobuf:"varint,3,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	SaleNum       int32   `protobuf:"varint,4,opt,name=sale_num,json=saleNum,proto3" json:"sale_num,omitempty"`
	UnitPrice     float32 `protobuf:"fixed32,5,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
	SaleAmount    float32 `protobuf:"fixed32,6,opt,name=sale_amount,json=saleAmount,proto3" json:"sale_amount,omitempty"`
	CostAmount    float32 `protobuf:"fixed32,7,opt,name=cost_amount,json=costAmount,proto3" json:"cost_amount,omitempty"`
	Profit        float32 `protobuf:"fixed32,8,opt,name=profit,proto3" json:"profit,omitempty"`
	ProfitRate    float32 `protobuf:"fixed32,9,opt,name=profit_rate,json=profitRate,proto3" json:"profit_rate,omitempty"`
	StatisticDate string  `protobuf:"bytes,10,opt,name=statistic_date,json=statisticDate,proto3" json:"statistic_date,omitempty"`
	CreatedAt     string  `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string  `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *SkuSaleDayStats) Reset() {
	*x = SkuSaleDayStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skuSaleStatsService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkuSaleDayStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkuSaleDayStats) ProtoMessage() {}

func (x *SkuSaleDayStats) ProtoReflect() protoreflect.Message {
	mi := &file_skuSaleStatsService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkuSaleDayStats.ProtoReflect.Descriptor instead.
func (*SkuSaleDayStats) Descriptor() ([]byte, []int) {
	return file_skuSaleStatsService_proto_rawDescGZIP(), []int{2}
}

func (x *SkuSaleDayStats) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SkuSaleDayStats) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *SkuSaleDayStats) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *SkuSaleDayStats) GetSaleNum() int32 {
	if x != nil {
		return x.SaleNum
	}
	return 0
}

func (x *SkuSaleDayStats) GetUnitPrice() float32 {
	if x != nil {
		return x.UnitPrice
	}
	return 0
}

func (x *SkuSaleDayStats) GetSaleAmount() float32 {
	if x != nil {
		return x.SaleAmount
	}
	return 0
}

func (x *SkuSaleDayStats) GetCostAmount() float32 {
	if x != nil {
		return x.CostAmount
	}
	return 0
}

func (x *SkuSaleDayStats) GetProfit() float32 {
	if x != nil {
		return x.Profit
	}
	return 0
}

func (x *SkuSaleDayStats) GetProfitRate() float32 {
	if x != nil {
		return x.ProfitRate
	}
	return 0
}

func (x *SkuSaleDayStats) GetStatisticDate() string {
	if x != nil {
		return x.StatisticDate
	}
	return ""
}

func (x *SkuSaleDayStats) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SkuSaleDayStats) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type SkuSaleDayStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *SkuSaleDayStats   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager      `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*SkuSaleDayStats `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *common.Error      `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info       `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SkuSaleDayStatsResponse) Reset() {
	*x = SkuSaleDayStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skuSaleStatsService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkuSaleDayStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkuSaleDayStatsResponse) ProtoMessage() {}

func (x *SkuSaleDayStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_skuSaleStatsService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkuSaleDayStatsResponse.ProtoReflect.Descriptor instead.
func (*SkuSaleDayStatsResponse) Descriptor() ([]byte, []int) {
	return file_skuSaleStatsService_proto_rawDescGZIP(), []int{3}
}

func (x *SkuSaleDayStatsResponse) GetEntity() *SkuSaleDayStats {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *SkuSaleDayStatsResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *SkuSaleDayStatsResponse) GetItems() []*SkuSaleDayStats {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *SkuSaleDayStatsResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SkuSaleDayStatsResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_skuSaleStatsService_proto protoreflect.FileDescriptor

var file_skuSaleStatsService_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x6b, 0x75, 0x53, 0x61, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x02, 0x0a, 0x0c, 0x53, 0x6b, 0x75,
	0x53, 0x61, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d,
	0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x61, 0x6c,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x61, 0x6c,
	0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x22,
	0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x61, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0xe0, 0x01, 0x0a, 0x14, 0x53, 0x6b, 0x75, 0x53, 0x61, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6b, 0x75, 0x53, 0x61, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2c,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6b, 0x75, 0x53, 0x61, 0x6c, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x22, 0xeb, 0x02, 0x0a, 0x0f, 0x53, 0x6b, 0x75, 0x53, 0x61, 0x6c, 0x65, 0x44,
	0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64,
	0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x61, 0x6c, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x61, 0x6c, 0x65, 0x4e,
	0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xe9, 0x01, 0x0a, 0x17, 0x53, 0x6b, 0x75, 0x53, 0x61, 0x6c, 0x65, 0x44, 0x61, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x6b, 0x75, 0x53, 0x61, 0x6c, 0x65,
	0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x53, 0x6b, 0x75, 0x53, 0x61, 0x6c, 0x65, 0x44, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_skuSaleStatsService_proto_rawDescOnce sync.Once
	file_skuSaleStatsService_proto_rawDescData = file_skuSaleStatsService_proto_rawDesc
)

func file_skuSaleStatsService_proto_rawDescGZIP() []byte {
	file_skuSaleStatsService_proto_rawDescOnce.Do(func() {
		file_skuSaleStatsService_proto_rawDescData = protoimpl.X.CompressGZIP(file_skuSaleStatsService_proto_rawDescData)
	})
	return file_skuSaleStatsService_proto_rawDescData
}

var file_skuSaleStatsService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_skuSaleStatsService_proto_goTypes = []interface{}{
	(*SkuSaleStats)(nil),            // 0: services.SkuSaleStats
	(*SkuSaleStatsResponse)(nil),    // 1: services.SkuSaleStatsResponse
	(*SkuSaleDayStats)(nil),         // 2: services.SkuSaleDayStats
	(*SkuSaleDayStatsResponse)(nil), // 3: services.SkuSaleDayStatsResponse
	(*common.Pager)(nil),            // 4: common.Pager
	(*common.Error)(nil),            // 5: common.Error
	(*common.Info)(nil),             // 6: common.Info
}
var file_skuSaleStatsService_proto_depIdxs = []int32{
	0,  // 0: services.SkuSaleStatsResponse.entity:type_name -> services.SkuSaleStats
	4,  // 1: services.SkuSaleStatsResponse.pager:type_name -> common.Pager
	0,  // 2: services.SkuSaleStatsResponse.items:type_name -> services.SkuSaleStats
	5,  // 3: services.SkuSaleStatsResponse.error:type_name -> common.Error
	6,  // 4: services.SkuSaleStatsResponse.info:type_name -> common.Info
	2,  // 5: services.SkuSaleDayStatsResponse.entity:type_name -> services.SkuSaleDayStats
	4,  // 6: services.SkuSaleDayStatsResponse.pager:type_name -> common.Pager
	2,  // 7: services.SkuSaleDayStatsResponse.items:type_name -> services.SkuSaleDayStats
	5,  // 8: services.SkuSaleDayStatsResponse.error:type_name -> common.Error
	6,  // 9: services.SkuSaleDayStatsResponse.info:type_name -> common.Info
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_skuSaleStatsService_proto_init() }
func file_skuSaleStatsService_proto_init() {
	if File_skuSaleStatsService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_skuSaleStatsService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkuSaleStats); i {
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
		file_skuSaleStatsService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkuSaleStatsResponse); i {
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
		file_skuSaleStatsService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkuSaleDayStats); i {
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
		file_skuSaleStatsService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkuSaleDayStatsResponse); i {
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
			RawDescriptor: file_skuSaleStatsService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_skuSaleStatsService_proto_goTypes,
		DependencyIndexes: file_skuSaleStatsService_proto_depIdxs,
		MessageInfos:      file_skuSaleStatsService_proto_msgTypes,
	}.Build()
	File_skuSaleStatsService_proto = out.File
	file_skuSaleStatsService_proto_rawDesc = nil
	file_skuSaleStatsService_proto_goTypes = nil
	file_skuSaleStatsService_proto_depIdxs = nil
}
