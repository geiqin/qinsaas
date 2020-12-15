// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: itemPresaleService.proto

package services

import (
	common "geiqin.saas.pdm/app/proto/common"
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

type ItemPresale struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                      int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ItemId                  int64   `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	PresaleType             int32   `protobuf:"varint,3,opt,name=presale_type,json=presaleType,proto3" json:"presale_type,omitempty"`
	PayedDate               string  `protobuf:"bytes,4,opt,name=payed_date,json=payedDate,proto3" json:"payed_date,omitempty"`
	DepositPaymentStartDate string  `protobuf:"bytes,5,opt,name=deposit_payment_start_date,json=depositPaymentStartDate,proto3" json:"deposit_payment_start_date,omitempty"`
	DepositPaymentEndDate   string  `protobuf:"bytes,6,opt,name=deposit_payment_end_date,json=depositPaymentEndDate,proto3" json:"deposit_payment_end_date,omitempty"`
	PaymentStartDate        string  `protobuf:"bytes,7,opt,name=payment_start_date,json=paymentStartDate,proto3" json:"payment_start_date,omitempty"`
	PaymentEndDate          string  `protobuf:"bytes,8,opt,name=payment_end_date,json=paymentEndDate,proto3" json:"payment_end_date,omitempty"`
	DepositType             int32   `protobuf:"varint,9,opt,name=deposit_type,json=depositType,proto3" json:"deposit_type,omitempty"`
	DepositRatio            int32   `protobuf:"varint,10,opt,name=deposit_ratio,json=depositRatio,proto3" json:"deposit_ratio,omitempty"`
	DepositAmount           float32 `protobuf:"fixed32,11,opt,name=deposit_amount,json=depositAmount,proto3" json:"deposit_amount,omitempty"`
	RestPayment             float32 `protobuf:"fixed32,12,opt,name=rest_payment,json=restPayment,proto3" json:"rest_payment,omitempty"`
	DeliveryType            int32   `protobuf:"varint,13,opt,name=delivery_type,json=deliveryType,proto3" json:"delivery_type,omitempty"`
	DeliveryDate            string  `protobuf:"bytes,14,opt,name=delivery_date,json=deliveryDate,proto3" json:"delivery_date,omitempty"`
	DeliveryDays            int32   `protobuf:"varint,15,opt,name=delivery_days,json=deliveryDays,proto3" json:"delivery_days,omitempty"`
}

func (x *ItemPresale) Reset() {
	*x = ItemPresale{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itemPresaleService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemPresale) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemPresale) ProtoMessage() {}

func (x *ItemPresale) ProtoReflect() protoreflect.Message {
	mi := &file_itemPresaleService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemPresale.ProtoReflect.Descriptor instead.
func (*ItemPresale) Descriptor() ([]byte, []int) {
	return file_itemPresaleService_proto_rawDescGZIP(), []int{0}
}

func (x *ItemPresale) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ItemPresale) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *ItemPresale) GetPresaleType() int32 {
	if x != nil {
		return x.PresaleType
	}
	return 0
}

func (x *ItemPresale) GetPayedDate() string {
	if x != nil {
		return x.PayedDate
	}
	return ""
}

func (x *ItemPresale) GetDepositPaymentStartDate() string {
	if x != nil {
		return x.DepositPaymentStartDate
	}
	return ""
}

func (x *ItemPresale) GetDepositPaymentEndDate() string {
	if x != nil {
		return x.DepositPaymentEndDate
	}
	return ""
}

func (x *ItemPresale) GetPaymentStartDate() string {
	if x != nil {
		return x.PaymentStartDate
	}
	return ""
}

func (x *ItemPresale) GetPaymentEndDate() string {
	if x != nil {
		return x.PaymentEndDate
	}
	return ""
}

func (x *ItemPresale) GetDepositType() int32 {
	if x != nil {
		return x.DepositType
	}
	return 0
}

func (x *ItemPresale) GetDepositRatio() int32 {
	if x != nil {
		return x.DepositRatio
	}
	return 0
}

func (x *ItemPresale) GetDepositAmount() float32 {
	if x != nil {
		return x.DepositAmount
	}
	return 0
}

func (x *ItemPresale) GetRestPayment() float32 {
	if x != nil {
		return x.RestPayment
	}
	return 0
}

func (x *ItemPresale) GetDeliveryType() int32 {
	if x != nil {
		return x.DeliveryType
	}
	return 0
}

func (x *ItemPresale) GetDeliveryDate() string {
	if x != nil {
		return x.DeliveryDate
	}
	return ""
}

func (x *ItemPresale) GetDeliveryDays() int32 {
	if x != nil {
		return x.DeliveryDays
	}
	return 0
}

type ItemPresaleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *ItemPresale  `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Error  *common.Error `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info  `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *ItemPresaleResponse) Reset() {
	*x = ItemPresaleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itemPresaleService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemPresaleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemPresaleResponse) ProtoMessage() {}

func (x *ItemPresaleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_itemPresaleService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemPresaleResponse.ProtoReflect.Descriptor instead.
func (*ItemPresaleResponse) Descriptor() ([]byte, []int) {
	return file_itemPresaleService_proto_rawDescGZIP(), []int{1}
}

func (x *ItemPresaleResponse) GetEntity() *ItemPresale {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *ItemPresaleResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ItemPresaleResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ItemPresaleResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_itemPresaleService_proto protoreflect.FileDescriptor

var file_itemPresaleService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x04, 0x0a, 0x0b, 0x49, 0x74, 0x65, 0x6d,
	0x50, 0x72, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x79, 0x65, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x3b, 0x0a, 0x1a, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x5f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x37, 0x0a, 0x18, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x15, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x5f, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0d, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x61, 0x79,
	0x73, 0x22, 0xb0, 0x01, 0x0a, 0x13, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x65, 0x73, 0x61, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x65, 0x73, 0x61, 0x6c, 0x65,
	0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x23, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x32, 0x56, 0x0a, 0x12, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x65, 0x73,
	0x61, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x1a, 0x1d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x65, 0x73, 0x61,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_itemPresaleService_proto_rawDescOnce sync.Once
	file_itemPresaleService_proto_rawDescData = file_itemPresaleService_proto_rawDesc
)

func file_itemPresaleService_proto_rawDescGZIP() []byte {
	file_itemPresaleService_proto_rawDescOnce.Do(func() {
		file_itemPresaleService_proto_rawDescData = protoimpl.X.CompressGZIP(file_itemPresaleService_proto_rawDescData)
	})
	return file_itemPresaleService_proto_rawDescData
}

var file_itemPresaleService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_itemPresaleService_proto_goTypes = []interface{}{
	(*ItemPresale)(nil),         // 0: services.ItemPresale
	(*ItemPresaleResponse)(nil), // 1: services.ItemPresaleResponse
	(*common.Pager)(nil),        // 2: common.Pager
	(*common.Error)(nil),        // 3: common.Error
	(*common.Info)(nil),         // 4: common.Info
}
var file_itemPresaleService_proto_depIdxs = []int32{
	0, // 0: services.ItemPresaleResponse.entity:type_name -> services.ItemPresale
	2, // 1: services.ItemPresaleResponse.pager:type_name -> common.Pager
	3, // 2: services.ItemPresaleResponse.error:type_name -> common.Error
	4, // 3: services.ItemPresaleResponse.info:type_name -> common.Info
	0, // 4: services.ItemPresaleService.Create:input_type -> services.ItemPresale
	1, // 5: services.ItemPresaleService.Create:output_type -> services.ItemPresaleResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_itemPresaleService_proto_init() }
func file_itemPresaleService_proto_init() {
	if File_itemPresaleService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_itemPresaleService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemPresale); i {
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
		file_itemPresaleService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemPresaleResponse); i {
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
			RawDescriptor: file_itemPresaleService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_itemPresaleService_proto_goTypes,
		DependencyIndexes: file_itemPresaleService_proto_depIdxs,
		MessageInfos:      file_itemPresaleService_proto_msgTypes,
	}.Build()
	File_itemPresaleService_proto = out.File
	file_itemPresaleService_proto_rawDesc = nil
	file_itemPresaleService_proto_goTypes = nil
	file_itemPresaleService_proto_depIdxs = nil
}