// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: buyingService.proto

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

// 购买清单
type Buying struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Solution    string  `protobuf:"bytes,1,opt,name=solution,proto3" json:"solution"`
	Changed     bool    `protobuf:"varint,2,opt,name=changed,proto3" json:"changed"`
	Count       int32   `protobuf:"varint,3,opt,name=count,proto3" json:"count"`
	Total       float32 `protobuf:"fixed32,4,opt,name=total,proto3" json:"total"`
	TotalWeight float32 `protobuf:"fixed32,25,opt,name=total_weight,json=totalWeight,proto3" json:"total_weight"`
	Discount    float32 `protobuf:"fixed32,5,opt,name=discount,proto3" json:"discount"`
	Freight     float32 `protobuf:"fixed32,6,opt,name=freight,proto3" json:"freight"`
	Amount      float32 `protobuf:"fixed32,7,opt,name=amount,proto3" json:"amount"`
	AddressId   int64   `protobuf:"varint,8,opt,name=address_id,json=addressId,proto3" json:"address_id"`
	CustomerId  int64   `protobuf:"varint,9,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	UseTicketId int64   `protobuf:"varint,10,opt,name=use_ticket_id,json=useTicketId,proto3" json:"use_ticket_id"` //正在使用的优惠劵ID
	Message     string  `protobuf:"bytes,11,opt,name=message,proto3" json:"message"`                               //买家留言(50字以内)
	// @inject_tag: gorm:"-"
	AvailableTicketIds []int64 `protobuf:"varint,12,rep,packed,name=available_ticket_ids,json=availableTicketIds,proto3" json:"available_ticket_ids" gorm:"-"` //可以使用的优惠劵
	// @inject_tag: gorm:"-"
	Items []*BuyingItem `protobuf:"bytes,13,rep,name=items,proto3" json:"items" gorm:"-"`
	// @inject_tag: gorm:"-"
	CartRowIds []string `protobuf:"bytes,14,rep,name=cart_row_ids,json=cartRowIds,proto3" json:"cart_row_ids" gorm:"-"`
	// @inject_tag: gorm:"-"
	Coupons []*OrderCoupon `protobuf:"bytes,15,rep,name=coupons,proto3" json:"coupons" gorm:"-"`
	// @inject_tag: gorm:"-"
	Promotions []*OrderPromotion `protobuf:"bytes,16,rep,name=promotions,proto3" json:"promotions" gorm:"-"`
	// @inject_tag: gorm:"-"
	Address *OrderAddress `protobuf:"bytes,17,opt,name=address,proto3" json:"address" gorm:"-"`
	// @inject_tag: gorm:"-"
	UseTicket          *CouponTicket `protobuf:"bytes,18,opt,name=use_ticket,json=useTicket,proto3" json:"use_ticket" gorm:"-"`
	PlatformSource     string        `protobuf:"bytes,19,opt,name=platform_source,json=platformSource,proto3" json:"platform_source"`                 //下单来源
	MemberMoney        float32       `protobuf:"fixed32,20,opt,name=member_money,json=memberMoney,proto3" json:"member_money"`                        // 会员优惠金额
	LimitDiscountMoney float32       `protobuf:"fixed32,21,opt,name=limit_discount_money,json=limitDiscountMoney,proto3" json:"limit_discount_money"` // 限时活动优惠金额
	RewardMoney        float32       `protobuf:"fixed32,22,opt,name=reward_money,json=rewardMoney,proto3" json:"reward_money"`                        // 满减送活动优惠金额
	CouponMoney        float32       `protobuf:"fixed32,23,opt,name=coupon_money,json=couponMoney,proto3" json:"coupon_money"`                        // 优惠券优惠金额
	// @inject_tag: gorm:"-"
	Shipment *OrderShipment `protobuf:"bytes,24,opt,name=shipment,proto3" json:"shipment" gorm:"-"` // 配送信息
	// @inject_tag: gorm:"-"
	Integral *Integral `protobuf:"bytes,26,opt,name=integral,proto3" json:"integral" gorm:"-"` // 积分信息
	// @inject_tag: gorm:"-"
	Food     *OrderFood     `protobuf:"bytes,27,opt,name=food,proto3" json:"food" gorm:"-"`                // 餐饮商品信息
	GroupReq *GroupOrderReq `protobuf:"bytes,28,opt,name=group_req,json=groupReq,proto3" json:"group_req"` // 拼团订单请求数据
	GroupRes *GroupOrderRes `protobuf:"bytes,29,opt,name=group_res,json=groupRes,proto3" json:"group_res"` // 拼团订单数据
}

func (x *Buying) Reset() {
	*x = Buying{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buyingService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Buying) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Buying) ProtoMessage() {}

func (x *Buying) ProtoReflect() protoreflect.Message {
	mi := &file_buyingService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Buying.ProtoReflect.Descriptor instead.
func (*Buying) Descriptor() ([]byte, []int) {
	return file_buyingService_proto_rawDescGZIP(), []int{0}
}

func (x *Buying) GetSolution() string {
	if x != nil {
		return x.Solution
	}
	return ""
}

func (x *Buying) GetChanged() bool {
	if x != nil {
		return x.Changed
	}
	return false
}

func (x *Buying) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Buying) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Buying) GetTotalWeight() float32 {
	if x != nil {
		return x.TotalWeight
	}
	return 0
}

func (x *Buying) GetDiscount() float32 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *Buying) GetFreight() float32 {
	if x != nil {
		return x.Freight
	}
	return 0
}

func (x *Buying) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Buying) GetAddressId() int64 {
	if x != nil {
		return x.AddressId
	}
	return 0
}

func (x *Buying) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Buying) GetUseTicketId() int64 {
	if x != nil {
		return x.UseTicketId
	}
	return 0
}

func (x *Buying) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Buying) GetAvailableTicketIds() []int64 {
	if x != nil {
		return x.AvailableTicketIds
	}
	return nil
}

func (x *Buying) GetItems() []*BuyingItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Buying) GetCartRowIds() []string {
	if x != nil {
		return x.CartRowIds
	}
	return nil
}

func (x *Buying) GetCoupons() []*OrderCoupon {
	if x != nil {
		return x.Coupons
	}
	return nil
}

func (x *Buying) GetPromotions() []*OrderPromotion {
	if x != nil {
		return x.Promotions
	}
	return nil
}

func (x *Buying) GetAddress() *OrderAddress {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Buying) GetUseTicket() *CouponTicket {
	if x != nil {
		return x.UseTicket
	}
	return nil
}

func (x *Buying) GetPlatformSource() string {
	if x != nil {
		return x.PlatformSource
	}
	return ""
}

func (x *Buying) GetMemberMoney() float32 {
	if x != nil {
		return x.MemberMoney
	}
	return 0
}

func (x *Buying) GetLimitDiscountMoney() float32 {
	if x != nil {
		return x.LimitDiscountMoney
	}
	return 0
}

func (x *Buying) GetRewardMoney() float32 {
	if x != nil {
		return x.RewardMoney
	}
	return 0
}

func (x *Buying) GetCouponMoney() float32 {
	if x != nil {
		return x.CouponMoney
	}
	return 0
}

func (x *Buying) GetShipment() *OrderShipment {
	if x != nil {
		return x.Shipment
	}
	return nil
}

func (x *Buying) GetIntegral() *Integral {
	if x != nil {
		return x.Integral
	}
	return nil
}

func (x *Buying) GetFood() *OrderFood {
	if x != nil {
		return x.Food
	}
	return nil
}

func (x *Buying) GetGroupReq() *GroupOrderReq {
	if x != nil {
		return x.GroupReq
	}
	return nil
}

func (x *Buying) GetGroupRes() *GroupOrderRes {
	if x != nil {
		return x.GroupRes
	}
	return nil
}

//购买清单明细
type BuyingItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId             int64   `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id"`
	SkuId              int64   `protobuf:"varint,2,opt,name=sku_id,json=skuId,proto3" json:"sku_id"`
	Num                int32   `protobuf:"varint,3,opt,name=num,proto3" json:"num"`
	Price              float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price"`
	OriginPrice        float32 `protobuf:"fixed32,5,opt,name=origin_price,json=originPrice,proto3" json:"origin_price"`
	SubTotal           float32 `protobuf:"fixed32,6,opt,name=sub_total,json=subTotal,proto3" json:"sub_total"`
	IsGift             bool    `protobuf:"varint,7,opt,name=is_gift,json=isGift,proto3" json:"is_gift"`
	PromotionId        int64   `protobuf:"varint,8,opt,name=promotion_id,json=promotionId,proto3" json:"promotion_id"`
	MemberPrice        float32 `protobuf:"fixed32,9,opt,name=member_price,json=memberPrice,proto3" json:"member_price"`                         // 会员价
	LimitDiscountPrice float32 `protobuf:"fixed32,10,opt,name=limit_discount_price,json=limitDiscountPrice,proto3" json:"limit_discount_price"` // 限时折扣价
	LimitDiscountNum   int32   `protobuf:"varint,11,opt,name=limit_discount_num,json=limitDiscountNum,proto3" json:"limit_discount_num"`        // 享受限时折扣的商品数量
	ExchangeNum        int32   `protobuf:"varint,12,opt,name=exchange_num,json=exchangeNum,proto3" json:"exchange_num"`                         // 兑换券兑换商品的数量
	// @inject_tag: gorm:"-"
	Goods     *GoodsInfo `protobuf:"bytes,13,opt,name=goods,proto3" json:"goods" gorm:"-"`
	CartRowId string     `protobuf:"bytes,14,opt,name=cart_row_id,json=cartRowId,proto3" json:"cart_row_id"`
}

func (x *BuyingItem) Reset() {
	*x = BuyingItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buyingService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyingItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyingItem) ProtoMessage() {}

func (x *BuyingItem) ProtoReflect() protoreflect.Message {
	mi := &file_buyingService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyingItem.ProtoReflect.Descriptor instead.
func (*BuyingItem) Descriptor() ([]byte, []int) {
	return file_buyingService_proto_rawDescGZIP(), []int{1}
}

func (x *BuyingItem) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *BuyingItem) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *BuyingItem) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *BuyingItem) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *BuyingItem) GetOriginPrice() float32 {
	if x != nil {
		return x.OriginPrice
	}
	return 0
}

func (x *BuyingItem) GetSubTotal() float32 {
	if x != nil {
		return x.SubTotal
	}
	return 0
}

func (x *BuyingItem) GetIsGift() bool {
	if x != nil {
		return x.IsGift
	}
	return false
}

func (x *BuyingItem) GetPromotionId() int64 {
	if x != nil {
		return x.PromotionId
	}
	return 0
}

func (x *BuyingItem) GetMemberPrice() float32 {
	if x != nil {
		return x.MemberPrice
	}
	return 0
}

func (x *BuyingItem) GetLimitDiscountPrice() float32 {
	if x != nil {
		return x.LimitDiscountPrice
	}
	return 0
}

func (x *BuyingItem) GetLimitDiscountNum() int32 {
	if x != nil {
		return x.LimitDiscountNum
	}
	return 0
}

func (x *BuyingItem) GetExchangeNum() int32 {
	if x != nil {
		return x.ExchangeNum
	}
	return 0
}

func (x *BuyingItem) GetGoods() *GoodsInfo {
	if x != nil {
		return x.Goods
	}
	return nil
}

func (x *BuyingItem) GetCartRowId() string {
	if x != nil {
		return x.CartRowId
	}
	return ""
}

//购买清单请求(订单下单使用)
type BuyingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId     int64  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`            //后台下单专用
	Source         int32  `protobuf:"varint,2,opt,name=source,proto3" json:"source"`                                      //请求来源 1是直接购买，2是 购物车下单，3，确认订单重复计算
	AddressId      int64  `protobuf:"varint,3,opt,name=address_id,json=addressId,proto3" json:"address_id"`               //收货地址
	Message        string `protobuf:"bytes,4,opt,name=message,proto3" json:"message"`                                     //买家留言(50字以内)
	VipcardId      int64  `protobuf:"varint,5,opt,name=vipcard_id,json=vipcardId,proto3" json:"vipcard_id"`               //选中的会员卡
	PayMethod      int32  `protobuf:"varint,6,opt,name=pay_method,json=payMethod,proto3" json:"pay_method"`               //选中的支付方式
	PlatformSource string `protobuf:"bytes,8,opt,name=platform_source,json=platformSource,proto3" json:"platform_source"` //下单来源
	UseTicketId    int64  `protobuf:"varint,7,opt,name=use_ticket_id,json=useTicketId,proto3" json:"use_ticket_id"`       //选中的优惠劵凭证ID
	// @inject_tag: gorm:"-"
	Items []*BuyingItem `protobuf:"bytes,9,rep,name=items,proto3" json:"items" gorm:"-"`
	// @inject_tag: gorm:"-"
	Shipment *OrderShipment `protobuf:"bytes,10,opt,name=shipment,proto3" json:"shipment" gorm:"-"` // 配送信息
	// @inject_tag: gorm:"-"
	Integral *Integral `protobuf:"bytes,11,opt,name=integral,proto3" json:"integral" gorm:"-"` // 积分信息
	// @inject_tag: gorm:"-"
	Food *OrderFood `protobuf:"bytes,12,opt,name=food,proto3" json:"food" gorm:"-"` // 餐饮商品信息
}

func (x *BuyingRequest) Reset() {
	*x = BuyingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buyingService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyingRequest) ProtoMessage() {}

func (x *BuyingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buyingService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyingRequest.ProtoReflect.Descriptor instead.
func (*BuyingRequest) Descriptor() ([]byte, []int) {
	return file_buyingService_proto_rawDescGZIP(), []int{2}
}

func (x *BuyingRequest) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *BuyingRequest) GetSource() int32 {
	if x != nil {
		return x.Source
	}
	return 0
}

func (x *BuyingRequest) GetAddressId() int64 {
	if x != nil {
		return x.AddressId
	}
	return 0
}

func (x *BuyingRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BuyingRequest) GetVipcardId() int64 {
	if x != nil {
		return x.VipcardId
	}
	return 0
}

func (x *BuyingRequest) GetPayMethod() int32 {
	if x != nil {
		return x.PayMethod
	}
	return 0
}

func (x *BuyingRequest) GetPlatformSource() string {
	if x != nil {
		return x.PlatformSource
	}
	return ""
}

func (x *BuyingRequest) GetUseTicketId() int64 {
	if x != nil {
		return x.UseTicketId
	}
	return 0
}

func (x *BuyingRequest) GetItems() []*BuyingItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *BuyingRequest) GetShipment() *OrderShipment {
	if x != nil {
		return x.Shipment
	}
	return nil
}

func (x *BuyingRequest) GetIntegral() *Integral {
	if x != nil {
		return x.Integral
	}
	return nil
}

func (x *BuyingRequest) GetFood() *OrderFood {
	if x != nil {
		return x.Food
	}
	return nil
}

type BuyingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Buying       `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Buying     `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *BuyingResponse) Reset() {
	*x = BuyingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buyingService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyingResponse) ProtoMessage() {}

func (x *BuyingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buyingService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyingResponse.ProtoReflect.Descriptor instead.
func (*BuyingResponse) Descriptor() ([]byte, []int) {
	return file_buyingService_proto_rawDescGZIP(), []int{3}
}

func (x *BuyingResponse) GetEntity() *Buying {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *BuyingResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *BuyingResponse) GetItems() []*Buying {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *BuyingResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *BuyingResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_buyingService_proto protoreflect.FileDescriptor

var file_buyingService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x62, 0x75, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x73, 0x68, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x46, 0x6f, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x75, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x08, 0x0a, 0x06,
	0x42, 0x75, 0x79, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x72, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x75, 0x73, 0x65,
	0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x75, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x0c, 0x20, 0x03, 0x28, 0x03, 0x52, 0x12, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x42, 0x75, 0x79, 0x69, 0x6e, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x72, 0x6f,
	0x77, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72,
	0x74, 0x52, 0x6f, 0x77, 0x49, 0x64, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x30, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x35, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x09, 0x75, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6d,
	0x6f, 0x6e, 0x65, 0x79, 0x18, 0x14, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x14, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x02, 0x52, 0x12, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x16, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0b, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12,
	0x33, 0x0a, 0x08, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x18, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x73, 0x68, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c,
	0x18, 0x1a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x04, 0x66, 0x6f, 0x6f, 0x64, 0x18, 0x1b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x04, 0x66, 0x6f, 0x6f, 0x64, 0x12, 0x34, 0x0a,
	0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x72, 0x65, 0x71, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x52, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x12, 0x34, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x72, 0x65, 0x73,
	0x18, 0x1d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x52,
	0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x22, 0xd1, 0x03, 0x0a, 0x0a, 0x42, 0x75,
	0x79, 0x69, 0x6e, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x73, 0x75, 0x62, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x67, 0x69, 0x66, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x69, 0x73, 0x47, 0x69, 0x66, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x30, 0x0a, 0x14, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x12, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x12,
	0x21, 0x0a, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x12, 0x29, 0x0a, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x1e, 0x0a,
	0x0b, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x72, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x72, 0x74, 0x52, 0x6f, 0x77, 0x49, 0x64, 0x22, 0xc6, 0x03,
	0x0a, 0x0d, 0x42, 0x75, 0x79, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x70, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x76, 0x69, 0x70, 0x63, 0x61, 0x72, 0x64, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x27, 0x0a, 0x0f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x5f,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x75, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x75, 0x79, 0x69, 0x6e, 0x67, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x73, 0x68, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x68, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x08, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a,
	0x08, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x6c, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x12, 0x27, 0x0a,
	0x04, 0x66, 0x6f, 0x6f, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x6f, 0x6f, 0x64,
	0x52, 0x04, 0x66, 0x6f, 0x6f, 0x64, 0x22, 0xce, 0x01, 0x0a, 0x0e, 0x42, 0x75, 0x79, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x75, 0x79, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x42, 0x75, 0x79, 0x69, 0x6e, 0x67, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x4a, 0x0a, 0x0d, 0x42, 0x75, 0x79, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x43, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x42, 0x75, 0x79, 0x69, 0x6e, 0x67, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x42, 0x75, 0x79, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buyingService_proto_rawDescOnce sync.Once
	file_buyingService_proto_rawDescData = file_buyingService_proto_rawDesc
)

func file_buyingService_proto_rawDescGZIP() []byte {
	file_buyingService_proto_rawDescOnce.Do(func() {
		file_buyingService_proto_rawDescData = protoimpl.X.CompressGZIP(file_buyingService_proto_rawDescData)
	})
	return file_buyingService_proto_rawDescData
}

var file_buyingService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_buyingService_proto_goTypes = []interface{}{
	(*Buying)(nil),         // 0: services.Buying
	(*BuyingItem)(nil),     // 1: services.BuyingItem
	(*BuyingRequest)(nil),  // 2: services.BuyingRequest
	(*BuyingResponse)(nil), // 3: services.BuyingResponse
	(*OrderCoupon)(nil),    // 4: services.OrderCoupon
	(*OrderPromotion)(nil), // 5: services.OrderPromotion
	(*OrderAddress)(nil),   // 6: services.OrderAddress
	(*CouponTicket)(nil),   // 7: services.CouponTicket
	(*OrderShipment)(nil),  // 8: services.OrderShipment
	(*Integral)(nil),       // 9: services.Integral
	(*OrderFood)(nil),      // 10: services.OrderFood
	(*GroupOrderReq)(nil),  // 11: services.GroupOrderReq
	(*GroupOrderRes)(nil),  // 12: services.GroupOrderRes
	(*GoodsInfo)(nil),      // 13: services.GoodsInfo
	(*common.Pager)(nil),   // 14: common.Pager
	(*common.Error)(nil),   // 15: common.Error
	(*common.Info)(nil),    // 16: common.Info
}
var file_buyingService_proto_depIdxs = []int32{
	1,  // 0: services.Buying.items:type_name -> services.BuyingItem
	4,  // 1: services.Buying.coupons:type_name -> services.OrderCoupon
	5,  // 2: services.Buying.promotions:type_name -> services.OrderPromotion
	6,  // 3: services.Buying.address:type_name -> services.OrderAddress
	7,  // 4: services.Buying.use_ticket:type_name -> services.CouponTicket
	8,  // 5: services.Buying.shipment:type_name -> services.OrderShipment
	9,  // 6: services.Buying.integral:type_name -> services.Integral
	10, // 7: services.Buying.food:type_name -> services.OrderFood
	11, // 8: services.Buying.group_req:type_name -> services.GroupOrderReq
	12, // 9: services.Buying.group_res:type_name -> services.GroupOrderRes
	13, // 10: services.BuyingItem.goods:type_name -> services.GoodsInfo
	1,  // 11: services.BuyingRequest.items:type_name -> services.BuyingItem
	8,  // 12: services.BuyingRequest.shipment:type_name -> services.OrderShipment
	9,  // 13: services.BuyingRequest.integral:type_name -> services.Integral
	10, // 14: services.BuyingRequest.food:type_name -> services.OrderFood
	0,  // 15: services.BuyingResponse.entity:type_name -> services.Buying
	14, // 16: services.BuyingResponse.pager:type_name -> common.Pager
	0,  // 17: services.BuyingResponse.items:type_name -> services.Buying
	15, // 18: services.BuyingResponse.error:type_name -> common.Error
	16, // 19: services.BuyingResponse.info:type_name -> common.Info
	0,  // 20: services.BuyingService.Calculate:input_type -> services.Buying
	3,  // 21: services.BuyingService.Calculate:output_type -> services.BuyingResponse
	21, // [21:22] is the sub-list for method output_type
	20, // [20:21] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_buyingService_proto_init() }
func file_buyingService_proto_init() {
	if File_buyingService_proto != nil {
		return
	}
	file_goodsInfoService_proto_init()
	file_orderAddressService_proto_init()
	file_orderCouponService_proto_init()
	file_orderPromotionService_proto_init()
	file_couponTicketInfoService_proto_init()
	file_shipmentService_proto_init()
	file_integralService_proto_init()
	file_orderFoodService_proto_init()
	file_groupBuyingService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_buyingService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Buying); i {
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
		file_buyingService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyingItem); i {
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
		file_buyingService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyingRequest); i {
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
		file_buyingService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyingResponse); i {
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
			RawDescriptor: file_buyingService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buyingService_proto_goTypes,
		DependencyIndexes: file_buyingService_proto_depIdxs,
		MessageInfos:      file_buyingService_proto_msgTypes,
	}.Build()
	File_buyingService_proto = out.File
	file_buyingService_proto_rawDesc = nil
	file_buyingService_proto_goTypes = nil
	file_buyingService_proto_depIdxs = nil
}
