// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: deliveryService.proto

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

type Delivery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Type           string  `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`
	OrderId        int64   `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	DeliveryType   string  `protobuf:"bytes,4,opt,name=delivery_type,json=deliveryType,proto3" json:"delivery_type"`
	DeliverySn     string  `protobuf:"bytes,5,opt,name=delivery_sn,json=deliverySn,proto3" json:"delivery_sn"`
	Freight        float32 `protobuf:"fixed32,6,opt,name=freight,proto3" json:"freight"`
	Protected      bool    `protobuf:"varint,7,opt,name=protected,proto3" json:"protected"`
	ExpressId      int32   `protobuf:"varint,8,opt,name=express_id,json=expressId,proto3" json:"express_id"`
	ExpressName    string  `protobuf:"bytes,9,opt,name=express_name,json=expressName,proto3" json:"express_name"`
	ExpressCode    string  `protobuf:"bytes,10,opt,name=express_code,json=expressCode,proto3" json:"express_code"`
	ExpressNo      string  `protobuf:"bytes,11,opt,name=express_no,json=expressNo,proto3" json:"express_no"`
	CustomerId     int64   `protobuf:"varint,12,opt,name=customer_id,json=customerId,proto3" json:"customer_id"`
	ReceiverName   string  `protobuf:"bytes,13,opt,name=receiver_name,json=receiverName,proto3" json:"receiver_name"`
	ReceiverAreaId int64   `protobuf:"varint,14,opt,name=receiver_area_id,json=receiverAreaId,proto3" json:"receiver_area_id"`
	ReceiverAddr   string  `protobuf:"bytes,15,opt,name=receiver_addr,json=receiverAddr,proto3" json:"receiver_addr"`
	ReceiverZip    string  `protobuf:"bytes,16,opt,name=receiver_zip,json=receiverZip,proto3" json:"receiver_zip"`
	ReceiverTel    string  `protobuf:"bytes,17,opt,name=receiver_tel,json=receiverTel,proto3" json:"receiver_tel"`
	ReceiverMobile string  `protobuf:"bytes,18,opt,name=receiver_mobile,json=receiverMobile,proto3" json:"receiver_mobile"`
	ReceiverEmail  string  `protobuf:"bytes,19,opt,name=receiver_email,json=receiverEmail,proto3" json:"receiver_email"`
	OpId           int64   `protobuf:"varint,20,opt,name=op_id,json=opId,proto3" json:"op_id"`
	Status         string  `protobuf:"bytes,21,opt,name=status,proto3" json:"status"`
	Memo           string  `protobuf:"bytes,22,opt,name=memo,proto3" json:"memo"`
	ArrivedAt      string  `protobuf:"bytes,23,opt,name=arrived_at,json=arrivedAt,proto3" json:"arrived_at"`
	CreatedAt      string  `protobuf:"bytes,24,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt      string  `protobuf:"bytes,25,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	// @inject_tag: gorm:"foreignKey:DeliveryId"
	Details     []*DeliveryDetail `protobuf:"bytes,26,rep,name=details,proto3" json:"details" gorm:"foreignKey:DeliveryId"`
	SafeguardId int64             `protobuf:"varint,27,opt,name=safeguard_id,json=safeguardId,proto3" json:"safeguard_id"`
}

func (x *Delivery) Reset() {
	*x = Delivery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deliveryService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Delivery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Delivery) ProtoMessage() {}

func (x *Delivery) ProtoReflect() protoreflect.Message {
	mi := &file_deliveryService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Delivery.ProtoReflect.Descriptor instead.
func (*Delivery) Descriptor() ([]byte, []int) {
	return file_deliveryService_proto_rawDescGZIP(), []int{0}
}

func (x *Delivery) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Delivery) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Delivery) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *Delivery) GetDeliveryType() string {
	if x != nil {
		return x.DeliveryType
	}
	return ""
}

func (x *Delivery) GetDeliverySn() string {
	if x != nil {
		return x.DeliverySn
	}
	return ""
}

func (x *Delivery) GetFreight() float32 {
	if x != nil {
		return x.Freight
	}
	return 0
}

func (x *Delivery) GetProtected() bool {
	if x != nil {
		return x.Protected
	}
	return false
}

func (x *Delivery) GetExpressId() int32 {
	if x != nil {
		return x.ExpressId
	}
	return 0
}

func (x *Delivery) GetExpressName() string {
	if x != nil {
		return x.ExpressName
	}
	return ""
}

func (x *Delivery) GetExpressCode() string {
	if x != nil {
		return x.ExpressCode
	}
	return ""
}

func (x *Delivery) GetExpressNo() string {
	if x != nil {
		return x.ExpressNo
	}
	return ""
}

func (x *Delivery) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Delivery) GetReceiverName() string {
	if x != nil {
		return x.ReceiverName
	}
	return ""
}

func (x *Delivery) GetReceiverAreaId() int64 {
	if x != nil {
		return x.ReceiverAreaId
	}
	return 0
}

func (x *Delivery) GetReceiverAddr() string {
	if x != nil {
		return x.ReceiverAddr
	}
	return ""
}

func (x *Delivery) GetReceiverZip() string {
	if x != nil {
		return x.ReceiverZip
	}
	return ""
}

func (x *Delivery) GetReceiverTel() string {
	if x != nil {
		return x.ReceiverTel
	}
	return ""
}

func (x *Delivery) GetReceiverMobile() string {
	if x != nil {
		return x.ReceiverMobile
	}
	return ""
}

func (x *Delivery) GetReceiverEmail() string {
	if x != nil {
		return x.ReceiverEmail
	}
	return ""
}

func (x *Delivery) GetOpId() int64 {
	if x != nil {
		return x.OpId
	}
	return 0
}

func (x *Delivery) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Delivery) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Delivery) GetArrivedAt() string {
	if x != nil {
		return x.ArrivedAt
	}
	return ""
}

func (x *Delivery) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Delivery) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Delivery) GetDetails() []*DeliveryDetail {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *Delivery) GetSafeguardId() int64 {
	if x != nil {
		return x.SafeguardId
	}
	return 0
}

type DeliveryDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	DeliveryId    int64 `protobuf:"varint,2,opt,name=delivery_id,json=deliveryId,proto3" json:"delivery_id"`
	OrderDetailId int64 `protobuf:"varint,3,opt,name=order_detail_id,json=orderDetailId,proto3" json:"order_detail_id"`
	ItemId        int64 `protobuf:"varint,4,opt,name=item_id,json=itemId,proto3" json:"item_id"`
	SkuId         int64 `protobuf:"varint,5,opt,name=sku_id,json=skuId,proto3" json:"sku_id"`
	Num           int32 `protobuf:"varint,6,opt,name=num,proto3" json:"num"`
	// @inject_tag: gorm:"-"
	Goods *GoodsInfo `protobuf:"bytes,7,opt,name=goods,proto3" json:"goods" gorm:"-"`
}

func (x *DeliveryDetail) Reset() {
	*x = DeliveryDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deliveryService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryDetail) ProtoMessage() {}

func (x *DeliveryDetail) ProtoReflect() protoreflect.Message {
	mi := &file_deliveryService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryDetail.ProtoReflect.Descriptor instead.
func (*DeliveryDetail) Descriptor() ([]byte, []int) {
	return file_deliveryService_proto_rawDescGZIP(), []int{1}
}

func (x *DeliveryDetail) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeliveryDetail) GetDeliveryId() int64 {
	if x != nil {
		return x.DeliveryId
	}
	return 0
}

func (x *DeliveryDetail) GetOrderDetailId() int64 {
	if x != nil {
		return x.OrderDetailId
	}
	return 0
}

func (x *DeliveryDetail) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *DeliveryDetail) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *DeliveryDetail) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *DeliveryDetail) GetGoods() *GoodsInfo {
	if x != nil {
		return x.Goods
	}
	return nil
}

type DeliveryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Delivery     `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Delivery   `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *DeliveryResponse) Reset() {
	*x = DeliveryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deliveryService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryResponse) ProtoMessage() {}

func (x *DeliveryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deliveryService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryResponse.ProtoReflect.Descriptor instead.
func (*DeliveryResponse) Descriptor() ([]byte, []int) {
	return file_deliveryService_proto_rawDescGZIP(), []int{2}
}

func (x *DeliveryResponse) GetEntity() *Delivery {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *DeliveryResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *DeliveryResponse) GetItems() []*Delivery {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *DeliveryResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *DeliveryResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_deliveryService_proto protoreflect.FileDescriptor

var file_deliveryService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x06, 0x0a,
	0x08, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x07, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6e, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x4e, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x28, 0x0a, 0x10, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x72, 0x65,
	0x61, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x72, 0x41, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x7a, 0x69, 0x70, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5a,
	0x69, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x74,
	0x65, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x54, 0x65, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x13, 0x0a, 0x05, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x72, 0x72, 0x69, 0x76, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x72, 0x69,
	0x76, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x1a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x07,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x61, 0x66, 0x65, 0x67,
	0x75, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73,
	0x61, 0x66, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x49, 0x64, 0x22, 0xd6, 0x01, 0x0a, 0x0e, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x49, 0x64, 0x12, 0x26,
	0x0a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12,
	0x15, 0x0a, 0x06, 0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x29, 0x0a, 0x05, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x22, 0xd4, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x4b, 0x0a, 0x0f, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deliveryService_proto_rawDescOnce sync.Once
	file_deliveryService_proto_rawDescData = file_deliveryService_proto_rawDesc
)

func file_deliveryService_proto_rawDescGZIP() []byte {
	file_deliveryService_proto_rawDescOnce.Do(func() {
		file_deliveryService_proto_rawDescData = protoimpl.X.CompressGZIP(file_deliveryService_proto_rawDescData)
	})
	return file_deliveryService_proto_rawDescData
}

var file_deliveryService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_deliveryService_proto_goTypes = []interface{}{
	(*Delivery)(nil),         // 0: services.Delivery
	(*DeliveryDetail)(nil),   // 1: services.DeliveryDetail
	(*DeliveryResponse)(nil), // 2: services.DeliveryResponse
	(*GoodsInfo)(nil),        // 3: services.GoodsInfo
	(*common.Pager)(nil),     // 4: common.Pager
	(*common.Error)(nil),     // 5: common.Error
	(*common.Info)(nil),      // 6: common.Info
}
var file_deliveryService_proto_depIdxs = []int32{
	1, // 0: services.Delivery.details:type_name -> services.DeliveryDetail
	3, // 1: services.DeliveryDetail.goods:type_name -> services.GoodsInfo
	0, // 2: services.DeliveryResponse.entity:type_name -> services.Delivery
	4, // 3: services.DeliveryResponse.pager:type_name -> common.Pager
	0, // 4: services.DeliveryResponse.items:type_name -> services.Delivery
	5, // 5: services.DeliveryResponse.error:type_name -> common.Error
	6, // 6: services.DeliveryResponse.info:type_name -> common.Info
	0, // 7: services.DeliveryService.List:input_type -> services.Delivery
	2, // 8: services.DeliveryService.List:output_type -> services.DeliveryResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_deliveryService_proto_init() }
func file_deliveryService_proto_init() {
	if File_deliveryService_proto != nil {
		return
	}
	file_goodsInfoService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_deliveryService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Delivery); i {
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
		file_deliveryService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryDetail); i {
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
		file_deliveryService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryResponse); i {
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
			RawDescriptor: file_deliveryService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_deliveryService_proto_goTypes,
		DependencyIndexes: file_deliveryService_proto_depIdxs,
		MessageInfos:      file_deliveryService_proto_msgTypes,
	}.Build()
	File_deliveryService_proto = out.File
	file_deliveryService_proto_rawDesc = nil
	file_deliveryService_proto_goTypes = nil
	file_deliveryService_proto_depIdxs = nil
}
