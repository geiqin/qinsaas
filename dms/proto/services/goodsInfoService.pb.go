// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: goodsInfoService.proto

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

type GoodsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                      int64              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ItemSn                  string             `protobuf:"bytes,2,opt,name=item_sn,json=itemSn,proto3" json:"item_sn,omitempty"`
	ModelType               string             `protobuf:"bytes,3,opt,name=model_type,json=modelType,proto3" json:"model_type,omitempty"`
	Name                    string             `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Unit                    string             `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit,omitempty"`
	BrandId                 int32              `protobuf:"varint,6,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	TaxonomyId              int64              `protobuf:"varint,7,opt,name=taxonomy_id,json=taxonomyId,proto3" json:"taxonomy_id,omitempty"`
	Quantity                int32              `protobuf:"varint,8,opt,name=quantity,proto3" json:"quantity,omitempty"`
	ThumbId                 int64              `protobuf:"varint,9,opt,name=thumb_id,json=thumbId,proto3" json:"thumb_id,omitempty"`
	ThumbUrl                string             `protobuf:"bytes,10,opt,name=thumb_url,json=thumbUrl,proto3" json:"thumb_url,omitempty"`
	Barcode                 string             `protobuf:"bytes,11,opt,name=barcode,proto3" json:"barcode,omitempty"`
	Price                   float32            `protobuf:"fixed32,12,opt,name=price,proto3" json:"price,omitempty"`
	OriginPrice             float32            `protobuf:"fixed32,13,opt,name=origin_price,json=originPrice,proto3" json:"origin_price,omitempty"`
	CostPrice               float32            `protobuf:"fixed32,14,opt,name=cost_price,json=costPrice,proto3" json:"cost_price,omitempty"`
	Weight                  float32            `protobuf:"fixed32,15,opt,name=weight,proto3" json:"weight,omitempty"`
	SkuId                   int64              `protobuf:"varint,16,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"` // 规格商品时使用
	SkuSn                   string             `protobuf:"bytes,17,opt,name=sku_sn,json=skuSn,proto3" json:"sku_sn,omitempty"`
	SkuName                 string             `protobuf:"bytes,18,opt,name=sku_name,json=skuName,proto3" json:"sku_name,omitempty"`
	SoldNum                 int32              `protobuf:"varint,19,opt,name=sold_num,json=soldNum,proto3" json:"sold_num,omitempty"`
	BuyCount                int32              `protobuf:"varint,20,opt,name=buy_count,json=buyCount,proto3" json:"buy_count,omitempty"`
	MinPrice                float32            `protobuf:"fixed32,21,opt,name=min_price,json=minPrice,proto3" json:"min_price,omitempty"`
	MaxPrice                float32            `protobuf:"fixed32,22,opt,name=max_price,json=maxPrice,proto3" json:"max_price,omitempty"`
	Disabled                bool               `protobuf:"varint,23,opt,name=disabled,proto3" json:"disabled,omitempty"`
	UpdatedAt               string             `protobuf:"bytes,24,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	IsCommissionDisplay     bool               `protobuf:"varint,25,opt,name=is_commission_display,json=isCommissionDisplay,proto3" json:"is_commission_display,omitempty"`
	IsCommissionRateDisplay bool               `protobuf:"varint,26,opt,name=is_commission_rate_display,json=isCommissionRateDisplay,proto3" json:"is_commission_rate_display,omitempty"`
	MinPrimaryRate          float32            `protobuf:"fixed32,27,opt,name=min_primary_rate,json=minPrimaryRate,proto3" json:"min_primary_rate,omitempty"`
	MaxPrimaryRate          float32            `protobuf:"fixed32,28,opt,name=max_primary_rate,json=maxPrimaryRate,proto3" json:"max_primary_rate,omitempty"`
	MinSecondRate           float32            `protobuf:"fixed32,29,opt,name=min_second_rate,json=minSecondRate,proto3" json:"min_second_rate,omitempty"`
	MaxSecondRate           float32            `protobuf:"fixed32,30,opt,name=max_second_rate,json=maxSecondRate,proto3" json:"max_second_rate,omitempty"`
	MinPrimary              float32            `protobuf:"fixed32,31,opt,name=min_primary,json=minPrimary,proto3" json:"min_primary,omitempty"`
	MaxPrimary              float32            `protobuf:"fixed32,32,opt,name=max_primary,json=maxPrimary,proto3" json:"max_primary,omitempty"`
	MinSecond               float32            `protobuf:"fixed32,33,opt,name=min_second,json=minSecond,proto3" json:"min_second,omitempty"`
	MaxSecond               float32            `protobuf:"fixed32,34,opt,name=max_second,json=maxSecond,proto3" json:"max_second,omitempty"`
	GoodsCommission         []*GoodsCommission `protobuf:"bytes,35,rep,name=goods_commission,json=goodsCommission,proto3" json:"goods_commission,omitempty"`
	Default                 bool               `protobuf:"varint,36,opt,name=default,proto3" json:"default,omitempty"`
}

func (x *GoodsInfo) Reset() {
	*x = GoodsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goodsInfoService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsInfo) ProtoMessage() {}

func (x *GoodsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_goodsInfoService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsInfo.ProtoReflect.Descriptor instead.
func (*GoodsInfo) Descriptor() ([]byte, []int) {
	return file_goodsInfoService_proto_rawDescGZIP(), []int{0}
}

func (x *GoodsInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GoodsInfo) GetItemSn() string {
	if x != nil {
		return x.ItemSn
	}
	return ""
}

func (x *GoodsInfo) GetModelType() string {
	if x != nil {
		return x.ModelType
	}
	return ""
}

func (x *GoodsInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoodsInfo) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *GoodsInfo) GetBrandId() int32 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

func (x *GoodsInfo) GetTaxonomyId() int64 {
	if x != nil {
		return x.TaxonomyId
	}
	return 0
}

func (x *GoodsInfo) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *GoodsInfo) GetThumbId() int64 {
	if x != nil {
		return x.ThumbId
	}
	return 0
}

func (x *GoodsInfo) GetThumbUrl() string {
	if x != nil {
		return x.ThumbUrl
	}
	return ""
}

func (x *GoodsInfo) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

func (x *GoodsInfo) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GoodsInfo) GetOriginPrice() float32 {
	if x != nil {
		return x.OriginPrice
	}
	return 0
}

func (x *GoodsInfo) GetCostPrice() float32 {
	if x != nil {
		return x.CostPrice
	}
	return 0
}

func (x *GoodsInfo) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *GoodsInfo) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *GoodsInfo) GetSkuSn() string {
	if x != nil {
		return x.SkuSn
	}
	return ""
}

func (x *GoodsInfo) GetSkuName() string {
	if x != nil {
		return x.SkuName
	}
	return ""
}

func (x *GoodsInfo) GetSoldNum() int32 {
	if x != nil {
		return x.SoldNum
	}
	return 0
}

func (x *GoodsInfo) GetBuyCount() int32 {
	if x != nil {
		return x.BuyCount
	}
	return 0
}

func (x *GoodsInfo) GetMinPrice() float32 {
	if x != nil {
		return x.MinPrice
	}
	return 0
}

func (x *GoodsInfo) GetMaxPrice() float32 {
	if x != nil {
		return x.MaxPrice
	}
	return 0
}

func (x *GoodsInfo) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *GoodsInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *GoodsInfo) GetIsCommissionDisplay() bool {
	if x != nil {
		return x.IsCommissionDisplay
	}
	return false
}

func (x *GoodsInfo) GetIsCommissionRateDisplay() bool {
	if x != nil {
		return x.IsCommissionRateDisplay
	}
	return false
}

func (x *GoodsInfo) GetMinPrimaryRate() float32 {
	if x != nil {
		return x.MinPrimaryRate
	}
	return 0
}

func (x *GoodsInfo) GetMaxPrimaryRate() float32 {
	if x != nil {
		return x.MaxPrimaryRate
	}
	return 0
}

func (x *GoodsInfo) GetMinSecondRate() float32 {
	if x != nil {
		return x.MinSecondRate
	}
	return 0
}

func (x *GoodsInfo) GetMaxSecondRate() float32 {
	if x != nil {
		return x.MaxSecondRate
	}
	return 0
}

func (x *GoodsInfo) GetMinPrimary() float32 {
	if x != nil {
		return x.MinPrimary
	}
	return 0
}

func (x *GoodsInfo) GetMaxPrimary() float32 {
	if x != nil {
		return x.MaxPrimary
	}
	return 0
}

func (x *GoodsInfo) GetMinSecond() float32 {
	if x != nil {
		return x.MinSecond
	}
	return 0
}

func (x *GoodsInfo) GetMaxSecond() float32 {
	if x != nil {
		return x.MaxSecond
	}
	return 0
}

func (x *GoodsInfo) GetGoodsCommission() []*GoodsCommission {
	if x != nil {
		return x.GoodsCommission
	}
	return nil
}

func (x *GoodsInfo) GetDefault() bool {
	if x != nil {
		return x.Default
	}
	return false
}

type GoodsInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *GoodsInfo    `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*GoodsInfo  `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *GoodsInfoResponse) Reset() {
	*x = GoodsInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goodsInfoService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsInfoResponse) ProtoMessage() {}

func (x *GoodsInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goodsInfoService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsInfoResponse.ProtoReflect.Descriptor instead.
func (*GoodsInfoResponse) Descriptor() ([]byte, []int) {
	return file_goodsInfoService_proto_rawDescGZIP(), []int{1}
}

func (x *GoodsInfoResponse) GetEntity() *GoodsInfo {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *GoodsInfoResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *GoodsInfoResponse) GetItems() []*GoodsInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *GoodsInfoResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *GoodsInfoResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_goodsInfoService_proto protoreflect.FileDescriptor

var file_goodsInfoService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x80, 0x09, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x53, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x68, 0x75,
	0x6d, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x68, 0x75,
	0x6d, 0x62, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x55, 0x72,
	0x6c, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x73,
	0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x6b, 0x75,
	0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x73, 0x6e, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x53, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x75,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6b, 0x75,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x75, 0x6d,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x6f, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x75, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x62, 0x75, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x08, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6d, 0x61,
	0x78, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x32, 0x0a, 0x15, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x19, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x13, 0x69, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x3b, 0x0a, 0x1a, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x69, 0x73, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x6d, 0x69,
	0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x10,
	0x6d, 0x61, 0x78, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x72, 0x61, 0x74, 0x65,
	0x18, 0x1c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x52, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0d, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x52, 0x61, 0x74, 0x65, 0x12, 0x26,
	0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x6d, 0x69, 0x6e,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x70,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x20, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x6d, 0x61,
	0x78, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x69, 0x6e, 0x5f,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x21, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6d, 0x69,
	0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x22, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6d, 0x61, 0x78,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x44, 0x0a, 0x10, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x23, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x24, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22, 0xd7, 0x01, 0x0a, 0x11, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x29,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goodsInfoService_proto_rawDescOnce sync.Once
	file_goodsInfoService_proto_rawDescData = file_goodsInfoService_proto_rawDesc
)

func file_goodsInfoService_proto_rawDescGZIP() []byte {
	file_goodsInfoService_proto_rawDescOnce.Do(func() {
		file_goodsInfoService_proto_rawDescData = protoimpl.X.CompressGZIP(file_goodsInfoService_proto_rawDescData)
	})
	return file_goodsInfoService_proto_rawDescData
}

var file_goodsInfoService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_goodsInfoService_proto_goTypes = []interface{}{
	(*GoodsInfo)(nil),         // 0: services.GoodsInfo
	(*GoodsInfoResponse)(nil), // 1: services.GoodsInfoResponse
	(*GoodsCommission)(nil),   // 2: services.GoodsCommission
	(*common.Pager)(nil),      // 3: common.Pager
	(*common.Error)(nil),      // 4: common.Error
	(*common.Info)(nil),       // 5: common.Info
}
var file_goodsInfoService_proto_depIdxs = []int32{
	2, // 0: services.GoodsInfo.goods_commission:type_name -> services.GoodsCommission
	0, // 1: services.GoodsInfoResponse.entity:type_name -> services.GoodsInfo
	3, // 2: services.GoodsInfoResponse.pager:type_name -> common.Pager
	0, // 3: services.GoodsInfoResponse.items:type_name -> services.GoodsInfo
	4, // 4: services.GoodsInfoResponse.error:type_name -> common.Error
	5, // 5: services.GoodsInfoResponse.info:type_name -> common.Info
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_goodsInfoService_proto_init() }
func file_goodsInfoService_proto_init() {
	if File_goodsInfoService_proto != nil {
		return
	}
	file_goodsCommissionService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_goodsInfoService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsInfo); i {
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
		file_goodsInfoService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsInfoResponse); i {
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
			RawDescriptor: file_goodsInfoService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_goodsInfoService_proto_goTypes,
		DependencyIndexes: file_goodsInfoService_proto_depIdxs,
		MessageInfos:      file_goodsInfoService_proto_msgTypes,
	}.Build()
	File_goodsInfoService_proto = out.File
	file_goodsInfoService_proto_rawDesc = nil
	file_goodsInfoService_proto_goTypes = nil
	file_goodsInfoService_proto_depIdxs = nil
}
