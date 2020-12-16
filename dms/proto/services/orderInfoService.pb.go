// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: orderInfoService.proto

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

type OrderInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId        int64              `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	OrderSn        string             `protobuf:"bytes,3,opt,name=order_sn,json=orderSn,proto3" json:"order_sn,omitempty"`
	PlatformSource string             `protobuf:"bytes,4,opt,name=platform_source,json=platformSource,proto3" json:"platform_source,omitempty"`
	OrderAmount    float32            `protobuf:"fixed32,5,opt,name=order_amount,json=orderAmount,proto3" json:"order_amount,omitempty"`
	Money          float32            `protobuf:"fixed32,7,opt,name=money,proto3" json:"money,omitempty"`
	IncomeType     string             `protobuf:"bytes,8,opt,name=income_type,json=incomeType,proto3" json:"income_type,omitempty"`
	IncomeRate     float32            `protobuf:"fixed32,9,opt,name=income_rate,json=incomeRate,proto3" json:"income_rate,omitempty"`
	Status         int32              `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`                                        // 结算状态
	BuyerName      string             `protobuf:"bytes,12,opt,name=buyer_name,json=buyerName,proto3" json:"buyer_name,omitempty"`                  // 下单用户名称
	Mobile         string             `protobuf:"bytes,13,opt,name=mobile,proto3" json:"mobile,omitempty"`                                         // 下单用户手机号
	OrderStatus    string             `protobuf:"bytes,14,opt,name=order_status,json=orderStatus,proto3" json:"order_status,omitempty"`            // 订单状态
	TotalNum       int32              `protobuf:"varint,15,opt,name=total_num,json=totalNum,proto3" json:"total_num,omitempty"`                    // 商品数量
	CanDelivery    bool               `protobuf:"varint,16,opt,name=can_delivery,json=canDelivery,proto3" json:"can_delivery,omitempty"`           // 是否需要发货
	CanLogistics   bool               `protobuf:"varint,17,opt,name=can_logistics,json=canLogistics,proto3" json:"can_logistics,omitempty"`        // 是否可查看物流
	OrderCreatedAt string             `protobuf:"bytes,18,opt,name=order_created_at,json=orderCreatedAt,proto3" json:"order_created_at,omitempty"` // 下单时间
	OrderDetails   []*OrderDetailInfo `protobuf:"bytes,19,rep,name=order_details,json=orderDetails,proto3" json:"order_details,omitempty"`
	LeaderId       int64              `protobuf:"varint,20,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`            // 团长ID
	LeaderName     string             `protobuf:"bytes,21,opt,name=leader_name,json=leaderName,proto3" json:"leader_name,omitempty"`       // 团长昵称
	LeaderMobile   string             `protobuf:"bytes,22,opt,name=leader_mobile,json=leaderMobile,proto3" json:"leader_mobile,omitempty"` // 团长手机号码
}

func (x *OrderInfo) Reset() {
	*x = OrderInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderInfoService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderInfo) ProtoMessage() {}

func (x *OrderInfo) ProtoReflect() protoreflect.Message {
	mi := &file_orderInfoService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderInfo.ProtoReflect.Descriptor instead.
func (*OrderInfo) Descriptor() ([]byte, []int) {
	return file_orderInfoService_proto_rawDescGZIP(), []int{0}
}

func (x *OrderInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderInfo) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderInfo) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

func (x *OrderInfo) GetPlatformSource() string {
	if x != nil {
		return x.PlatformSource
	}
	return ""
}

func (x *OrderInfo) GetOrderAmount() float32 {
	if x != nil {
		return x.OrderAmount
	}
	return 0
}

func (x *OrderInfo) GetMoney() float32 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *OrderInfo) GetIncomeType() string {
	if x != nil {
		return x.IncomeType
	}
	return ""
}

func (x *OrderInfo) GetIncomeRate() float32 {
	if x != nil {
		return x.IncomeRate
	}
	return 0
}

func (x *OrderInfo) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *OrderInfo) GetBuyerName() string {
	if x != nil {
		return x.BuyerName
	}
	return ""
}

func (x *OrderInfo) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *OrderInfo) GetOrderStatus() string {
	if x != nil {
		return x.OrderStatus
	}
	return ""
}

func (x *OrderInfo) GetTotalNum() int32 {
	if x != nil {
		return x.TotalNum
	}
	return 0
}

func (x *OrderInfo) GetCanDelivery() bool {
	if x != nil {
		return x.CanDelivery
	}
	return false
}

func (x *OrderInfo) GetCanLogistics() bool {
	if x != nil {
		return x.CanLogistics
	}
	return false
}

func (x *OrderInfo) GetOrderCreatedAt() string {
	if x != nil {
		return x.OrderCreatedAt
	}
	return ""
}

func (x *OrderInfo) GetOrderDetails() []*OrderDetailInfo {
	if x != nil {
		return x.OrderDetails
	}
	return nil
}

func (x *OrderInfo) GetLeaderId() int64 {
	if x != nil {
		return x.LeaderId
	}
	return 0
}

func (x *OrderInfo) GetLeaderName() string {
	if x != nil {
		return x.LeaderName
	}
	return ""
}

func (x *OrderInfo) GetLeaderMobile() string {
	if x != nil {
		return x.LeaderMobile
	}
	return ""
}

type OrderDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId    int64   `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	ItemId     int64   `protobuf:"varint,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	ThumbUrl   string  `protobuf:"bytes,5,opt,name=thumb_url,json=thumbUrl,proto3" json:"thumb_url,omitempty"`
	SkuId      int64   `protobuf:"varint,6,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	ItemSn     string  `protobuf:"bytes,7,opt,name=item_sn,json=itemSn,proto3" json:"item_sn,omitempty"`
	SkuSn      string  `protobuf:"bytes,8,opt,name=sku_sn,json=skuSn,proto3" json:"sku_sn,omitempty"`
	ModelType  string  `protobuf:"bytes,9,opt,name=model_type,json=modelType,proto3" json:"model_type,omitempty"`
	Name       string  `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	Num        int32   `protobuf:"varint,12,opt,name=num,proto3" json:"num,omitempty"`
	Price      float32 `protobuf:"fixed32,13,opt,name=price,proto3" json:"price,omitempty"`
	TotalPrice float32 `protobuf:"fixed32,14,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
}

func (x *OrderDetailInfo) Reset() {
	*x = OrderDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderInfoService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderDetailInfo) ProtoMessage() {}

func (x *OrderDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_orderInfoService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderDetailInfo.ProtoReflect.Descriptor instead.
func (*OrderDetailInfo) Descriptor() ([]byte, []int) {
	return file_orderInfoService_proto_rawDescGZIP(), []int{1}
}

func (x *OrderDetailInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderDetailInfo) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderDetailInfo) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *OrderDetailInfo) GetThumbUrl() string {
	if x != nil {
		return x.ThumbUrl
	}
	return ""
}

func (x *OrderDetailInfo) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *OrderDetailInfo) GetItemSn() string {
	if x != nil {
		return x.ItemSn
	}
	return ""
}

func (x *OrderDetailInfo) GetSkuSn() string {
	if x != nil {
		return x.SkuSn
	}
	return ""
}

func (x *OrderDetailInfo) GetModelType() string {
	if x != nil {
		return x.ModelType
	}
	return ""
}

func (x *OrderDetailInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrderDetailInfo) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *OrderDetailInfo) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OrderDetailInfo) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

type OrderInfoWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged         int32   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize      int32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Id            int64   `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Mobile        string  `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile,omitempty"`
	OrderId       int64   `protobuf:"varint,5,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	OrderSn       string  `protobuf:"bytes,6,opt,name=order_sn,json=orderSn,proto3" json:"order_sn,omitempty"`
	Keywords      string  `protobuf:"bytes,7,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Status        string  `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Ids           []int64 `protobuf:"varint,9,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	DistributorId int64   `protobuf:"varint,10,opt,name=distributor_id,json=distributorId,proto3" json:"distributor_id,omitempty"`
	CustomerId    int64   `protobuf:"varint,11,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	LeaderId      int64   `protobuf:"varint,12,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
}

func (x *OrderInfoWhere) Reset() {
	*x = OrderInfoWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderInfoService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderInfoWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderInfoWhere) ProtoMessage() {}

func (x *OrderInfoWhere) ProtoReflect() protoreflect.Message {
	mi := &file_orderInfoService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderInfoWhere.ProtoReflect.Descriptor instead.
func (*OrderInfoWhere) Descriptor() ([]byte, []int) {
	return file_orderInfoService_proto_rawDescGZIP(), []int{2}
}

func (x *OrderInfoWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *OrderInfoWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *OrderInfoWhere) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderInfoWhere) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *OrderInfoWhere) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderInfoWhere) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

func (x *OrderInfoWhere) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *OrderInfoWhere) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderInfoWhere) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *OrderInfoWhere) GetDistributorId() int64 {
	if x != nil {
		return x.DistributorId
	}
	return 0
}

func (x *OrderInfoWhere) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *OrderInfoWhere) GetLeaderId() int64 {
	if x != nil {
		return x.LeaderId
	}
	return 0
}

type OrderInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *OrderInfo    `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*OrderInfo  `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *OrderInfoResponse) Reset() {
	*x = OrderInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderInfoService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderInfoResponse) ProtoMessage() {}

func (x *OrderInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderInfoService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderInfoResponse.ProtoReflect.Descriptor instead.
func (*OrderInfoResponse) Descriptor() ([]byte, []int) {
	return file_orderInfoService_proto_rawDescGZIP(), []int{3}
}

func (x *OrderInfoResponse) GetEntity() *OrderInfo {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *OrderInfoResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *OrderInfoResponse) GetItems() []*OrderInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *OrderInfoResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *OrderInfoResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_orderInfoService_proto protoreflect.FileDescriptor

var file_orderInfoService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x05, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0a, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x79, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x61, 0x6e, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x61, 0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x63, 0x61, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3e,
	0x0a, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x22, 0xb5, 0x02, 0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x68, 0x75,
	0x6d, 0x62, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x68,
	0x75, 0x6d, 0x62, 0x55, 0x72, 0x6c, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x69, 0x74, 0x65, 0x6d, 0x53, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x73, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x53, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e,
	0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0xcc, 0x02, 0x0a, 0x0e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0xd7, 0x01, 0x0a, 0x11, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72,
	0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x32, 0x93, 0x01, 0x0a, 0x0e, 0x4d, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x91, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x96, 0x01, 0x0a,
	0x11, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x57, 0x68,
	0x65, 0x72, 0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x98, 0x01, 0x0a, 0x13, 0x4d, 0x79, 0x42, 0x6f, 0x6e, 0x75,
	0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x57, 0x68, 0x65, 0x72,
	0x65, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orderInfoService_proto_rawDescOnce sync.Once
	file_orderInfoService_proto_rawDescData = file_orderInfoService_proto_rawDesc
)

func file_orderInfoService_proto_rawDescGZIP() []byte {
	file_orderInfoService_proto_rawDescOnce.Do(func() {
		file_orderInfoService_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderInfoService_proto_rawDescData)
	})
	return file_orderInfoService_proto_rawDescData
}

var file_orderInfoService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_orderInfoService_proto_goTypes = []interface{}{
	(*OrderInfo)(nil),         // 0: services.OrderInfo
	(*OrderDetailInfo)(nil),   // 1: services.OrderDetailInfo
	(*OrderInfoWhere)(nil),    // 2: services.OrderInfoWhere
	(*OrderInfoResponse)(nil), // 3: services.OrderInfoResponse
	(*common.Pager)(nil),      // 4: common.Pager
	(*common.Error)(nil),      // 5: common.Error
	(*common.Info)(nil),       // 6: common.Info
}
var file_orderInfoService_proto_depIdxs = []int32{
	1,  // 0: services.OrderInfo.order_details:type_name -> services.OrderDetailInfo
	0,  // 1: services.OrderInfoResponse.entity:type_name -> services.OrderInfo
	4,  // 2: services.OrderInfoResponse.pager:type_name -> common.Pager
	0,  // 3: services.OrderInfoResponse.items:type_name -> services.OrderInfo
	5,  // 4: services.OrderInfoResponse.error:type_name -> common.Error
	6,  // 5: services.OrderInfoResponse.info:type_name -> common.Info
	2,  // 6: services.MyOrderService.Get:input_type -> services.OrderInfoWhere
	2,  // 7: services.MyOrderService.Search:input_type -> services.OrderInfoWhere
	2,  // 8: services.OrderService.Get:input_type -> services.OrderInfoWhere
	2,  // 9: services.OrderService.Search:input_type -> services.OrderInfoWhere
	2,  // 10: services.BonusOrderService.Get:input_type -> services.OrderInfoWhere
	2,  // 11: services.BonusOrderService.Search:input_type -> services.OrderInfoWhere
	2,  // 12: services.MyBonusOrderService.Get:input_type -> services.OrderInfoWhere
	2,  // 13: services.MyBonusOrderService.Search:input_type -> services.OrderInfoWhere
	3,  // 14: services.MyOrderService.Get:output_type -> services.OrderInfoResponse
	3,  // 15: services.MyOrderService.Search:output_type -> services.OrderInfoResponse
	3,  // 16: services.OrderService.Get:output_type -> services.OrderInfoResponse
	3,  // 17: services.OrderService.Search:output_type -> services.OrderInfoResponse
	3,  // 18: services.BonusOrderService.Get:output_type -> services.OrderInfoResponse
	3,  // 19: services.BonusOrderService.Search:output_type -> services.OrderInfoResponse
	3,  // 20: services.MyBonusOrderService.Get:output_type -> services.OrderInfoResponse
	3,  // 21: services.MyBonusOrderService.Search:output_type -> services.OrderInfoResponse
	14, // [14:22] is the sub-list for method output_type
	6,  // [6:14] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_orderInfoService_proto_init() }
func file_orderInfoService_proto_init() {
	if File_orderInfoService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orderInfoService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderInfo); i {
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
		file_orderInfoService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderDetailInfo); i {
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
		file_orderInfoService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderInfoWhere); i {
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
		file_orderInfoService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderInfoResponse); i {
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
			RawDescriptor: file_orderInfoService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_orderInfoService_proto_goTypes,
		DependencyIndexes: file_orderInfoService_proto_depIdxs,
		MessageInfos:      file_orderInfoService_proto_msgTypes,
	}.Build()
	File_orderInfoService_proto = out.File
	file_orderInfoService_proto_rawDesc = nil
	file_orderInfoService_proto_goTypes = nil
	file_orderInfoService_proto_depIdxs = nil
}