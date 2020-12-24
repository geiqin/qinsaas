// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: levelService.proto

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

type Level struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name             string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Type             int32   `protobuf:"varint,3,opt,name=type,proto3" json:"type"`
	Discount         float32 `protobuf:"fixed32,4,opt,name=discount,proto3" json:"discount"`
	PostageFree      bool    `protobuf:"varint,5,opt,name=postage_free,json=postageFree,proto3" json:"postage_free"`
	PointRate        float32 `protobuf:"fixed32,6,opt,name=point_rate,json=pointRate,proto3" json:"point_rate"`
	UplevelId        int32   `protobuf:"varint,7,opt,name=uplevel_id,json=uplevelId,proto3" json:"uplevel_id"`
	OrderLimitType   int32   `protobuf:"varint,8,opt,name=order_limit_type,json=orderLimitType,proto3" json:"order_limit_type"`
	OrderLimitMoney  float32 `protobuf:"fixed32,9,opt,name=order_limit_money,json=orderLimitMoney,proto3" json:"order_limit_money"`
	DepositLimit     int32   `protobuf:"varint,10,opt,name=deposit_limit,json=depositLimit,proto3" json:"deposit_limit"`
	GrowthLimit      int32   `protobuf:"varint,11,opt,name=growth_limit,json=growthLimit,proto3" json:"growth_limit"`
	PointExpiredDays int32   `protobuf:"varint,12,opt,name=point_expired_days,json=pointExpiredDays,proto3" json:"point_expired_days"`
	IconId           int64   `protobuf:"varint,13,opt,name=icon_id,json=iconId,proto3" json:"icon_id"`
	IconUrl          string  `protobuf:"bytes,14,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url"`
	CardBg           string  `protobuf:"bytes,15,opt,name=card_bg,json=cardBg,proto3" json:"card_bg"`
	Defaulted        bool    `protobuf:"varint,16,opt,name=defaulted,proto3" json:"defaulted"`
	Disabled         bool    `protobuf:"varint,17,opt,name=disabled,proto3" json:"disabled"`
	Locked           bool    `protobuf:"varint,18,opt,name=locked,proto3" json:"locked"`
	Description      string  `protobuf:"bytes,19,opt,name=description,proto3" json:"description"`
	CreatedAt        string  `protobuf:"bytes,20,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt        string  `protobuf:"bytes,21,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	// @inject_tag: gorm:"-"
	Ids []int32 `protobuf:"varint,22,rep,packed,name=ids,proto3" json:"ids" gorm:"-"`
}

func (x *Level) Reset() {
	*x = Level{}
	if protoimpl.UnsafeEnabled {
		mi := &file_levelService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Level) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Level) ProtoMessage() {}

func (x *Level) ProtoReflect() protoreflect.Message {
	mi := &file_levelService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Level.ProtoReflect.Descriptor instead.
func (*Level) Descriptor() ([]byte, []int) {
	return file_levelService_proto_rawDescGZIP(), []int{0}
}

func (x *Level) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Level) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Level) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Level) GetDiscount() float32 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *Level) GetPostageFree() bool {
	if x != nil {
		return x.PostageFree
	}
	return false
}

func (x *Level) GetPointRate() float32 {
	if x != nil {
		return x.PointRate
	}
	return 0
}

func (x *Level) GetUplevelId() int32 {
	if x != nil {
		return x.UplevelId
	}
	return 0
}

func (x *Level) GetOrderLimitType() int32 {
	if x != nil {
		return x.OrderLimitType
	}
	return 0
}

func (x *Level) GetOrderLimitMoney() float32 {
	if x != nil {
		return x.OrderLimitMoney
	}
	return 0
}

func (x *Level) GetDepositLimit() int32 {
	if x != nil {
		return x.DepositLimit
	}
	return 0
}

func (x *Level) GetGrowthLimit() int32 {
	if x != nil {
		return x.GrowthLimit
	}
	return 0
}

func (x *Level) GetPointExpiredDays() int32 {
	if x != nil {
		return x.PointExpiredDays
	}
	return 0
}

func (x *Level) GetIconId() int64 {
	if x != nil {
		return x.IconId
	}
	return 0
}

func (x *Level) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

func (x *Level) GetCardBg() string {
	if x != nil {
		return x.CardBg
	}
	return ""
}

func (x *Level) GetDefaulted() bool {
	if x != nil {
		return x.Defaulted
	}
	return false
}

func (x *Level) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *Level) GetLocked() bool {
	if x != nil {
		return x.Locked
	}
	return false
}

func (x *Level) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Level) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Level) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Level) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

//
type LevelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Level        `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Level      `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *LevelResponse) Reset() {
	*x = LevelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_levelService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LevelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LevelResponse) ProtoMessage() {}

func (x *LevelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_levelService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LevelResponse.ProtoReflect.Descriptor instead.
func (*LevelResponse) Descriptor() ([]byte, []int) {
	return file_levelService_proto_rawDescGZIP(), []int{1}
}

func (x *LevelResponse) GetEntity() *Level {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *LevelResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *LevelResponse) GetItems() []*Level {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *LevelResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *LevelResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_levelService_proto protoreflect.FileDescriptor

var file_levelService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x99, 0x05, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x46, 0x72,
	0x65, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x75, 0x70, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64,
	0x12, 0x28, 0x0a, 0x10, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x64,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x67,
	0x72, 0x6f, 0x77, 0x74, 0x68, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x67, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x2c,
	0x0a, 0x12, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f,
	0x64, 0x61, 0x79, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x44, 0x61, 0x79, 0x73, 0x12, 0x17, 0x0a, 0x07,
	0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69,
	0x63, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x62, 0x67, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x61, 0x72, 0x64, 0x42, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x16, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xcb, 0x01,
	0x0a, 0x0d, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x27, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x25, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xd1, 0x02, 0x0a, 0x0c,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x1a, 0x17, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x36, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x11, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x57,
	0x68, 0x65, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_levelService_proto_rawDescOnce sync.Once
	file_levelService_proto_rawDescData = file_levelService_proto_rawDesc
)

func file_levelService_proto_rawDescGZIP() []byte {
	file_levelService_proto_rawDescOnce.Do(func() {
		file_levelService_proto_rawDescData = protoimpl.X.CompressGZIP(file_levelService_proto_rawDescData)
	})
	return file_levelService_proto_rawDescData
}

var file_levelService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_levelService_proto_goTypes = []interface{}{
	(*Level)(nil),            // 0: services.Level
	(*LevelResponse)(nil),    // 1: services.LevelResponse
	(*common.Pager)(nil),     // 2: common.Pager
	(*common.Error)(nil),     // 3: common.Error
	(*common.Info)(nil),      // 4: common.Info
	(*common.BaseWhere)(nil), // 5: common.BaseWhere
}
var file_levelService_proto_depIdxs = []int32{
	0,  // 0: services.LevelResponse.entity:type_name -> services.Level
	2,  // 1: services.LevelResponse.pager:type_name -> common.Pager
	0,  // 2: services.LevelResponse.items:type_name -> services.Level
	3,  // 3: services.LevelResponse.error:type_name -> common.Error
	4,  // 4: services.LevelResponse.info:type_name -> common.Info
	0,  // 5: services.LevelService.Create:input_type -> services.Level
	0,  // 6: services.LevelService.Update:input_type -> services.Level
	0,  // 7: services.LevelService.Delete:input_type -> services.Level
	0,  // 8: services.LevelService.Get:input_type -> services.Level
	5,  // 9: services.LevelService.Search:input_type -> common.BaseWhere
	5,  // 10: services.LevelService.List:input_type -> common.BaseWhere
	1,  // 11: services.LevelService.Create:output_type -> services.LevelResponse
	1,  // 12: services.LevelService.Update:output_type -> services.LevelResponse
	1,  // 13: services.LevelService.Delete:output_type -> services.LevelResponse
	1,  // 14: services.LevelService.Get:output_type -> services.LevelResponse
	1,  // 15: services.LevelService.Search:output_type -> services.LevelResponse
	1,  // 16: services.LevelService.List:output_type -> services.LevelResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_levelService_proto_init() }
func file_levelService_proto_init() {
	if File_levelService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_levelService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Level); i {
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
		file_levelService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LevelResponse); i {
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
			RawDescriptor: file_levelService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_levelService_proto_goTypes,
		DependencyIndexes: file_levelService_proto_depIdxs,
		MessageInfos:      file_levelService_proto_msgTypes,
	}.Build()
	File_levelService_proto = out.File
	file_levelService_proto_rawDesc = nil
	file_levelService_proto_goTypes = nil
	file_levelService_proto_depIdxs = nil
}
