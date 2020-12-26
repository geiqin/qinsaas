// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: rankService.proto

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

type RankWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Top      int32 `protobuf:"varint,3,opt,name=top,proto3" json:"top"`
	Id       int64 `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	// @inject_tag: gorm:"-"
	Ids  []int64 `protobuf:"varint,5,rep,packed,name=ids,proto3" json:"ids" gorm:"-"`
	Type string  `protobuf:"bytes,6,opt,name=type,proto3" json:"type"`
}

func (x *RankWhere) Reset() {
	*x = RankWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rankService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankWhere) ProtoMessage() {}

func (x *RankWhere) ProtoReflect() protoreflect.Message {
	mi := &file_rankService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankWhere.ProtoReflect.Descriptor instead.
func (*RankWhere) Descriptor() ([]byte, []int) {
	return file_rankService_proto_rawDescGZIP(), []int{0}
}

func (x *RankWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *RankWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *RankWhere) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *RankWhere) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RankWhere) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *RankWhere) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Rank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32             `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name            string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	PrimaryRate     float32           `protobuf:"fixed32,3,opt,name=primary_rate,json=primaryRate,proto3" json:"primary_rate"`
	SecondRate      float32           `protobuf:"fixed32,4,opt,name=second_rate,json=secondRate,proto3" json:"second_rate"`
	ThreeRate       float32           `protobuf:"fixed32,5,opt,name=three_rate,json=threeRate,proto3" json:"three_rate"`
	RewardMoney     float32           `protobuf:"fixed32,6,opt,name=reward_money,json=rewardMoney,proto3" json:"reward_money"`
	MonthDrawMoney  float32           `protobuf:"fixed32,7,opt,name=month_draw_money,json=monthDrawMoney,proto3" json:"month_draw_money"`
	MonthDrawNum    int32             `protobuf:"varint,8,opt,name=month_draw_num,json=monthDrawNum,proto3" json:"month_draw_num"`
	LevelId         int32             `protobuf:"varint,9,opt,name=level_id,json=levelId,proto3" json:"level_id"`
	IsCondition     bool              `protobuf:"varint,10,opt,name=is_condition,json=isCondition,proto3" json:"is_condition"`
	PromotionAmount float32           `protobuf:"fixed32,11,opt,name=promotion_amount,json=promotionAmount,proto3" json:"promotion_amount"`
	PrimaryNum      int32             `protobuf:"varint,12,opt,name=primary_num,json=primaryNum,proto3" json:"primary_num"`
	SecondNum       int32             `protobuf:"varint,13,opt,name=second_num,json=secondNum,proto3" json:"second_num"`
	IsBought        bool              `protobuf:"varint,14,opt,name=is_bought,json=isBought,proto3" json:"is_bought"`
	ConsumeAmount   float32           `protobuf:"fixed32,15,opt,name=consume_amount,json=consumeAmount,proto3" json:"consume_amount"`
	ConsumeNum      int32             `protobuf:"varint,16,opt,name=consume_num,json=consumeNum,proto3" json:"consume_num"`
	IsBindIdcard    bool              `protobuf:"varint,17,opt,name=is_bind_idcard,json=isBindIdcard,proto3" json:"is_bind_idcard"`
	IsBindMobile    bool              `protobuf:"varint,18,opt,name=is_bind_mobile,json=isBindMobile,proto3" json:"is_bind_mobile"`
	JoinFee         float32           `protobuf:"fixed32,19,opt,name=join_fee,json=joinFee,proto3" json:"join_fee"`
	Defaulted       bool              `protobuf:"varint,20,opt,name=defaulted,proto3" json:"defaulted"`
	Goodses         []*GoodsCondition `protobuf:"bytes,21,rep,name=goodses,proto3" json:"goodses"`
	// @inject_tag: gorm:"-"
	Ids         []int32 `protobuf:"varint,22,rep,packed,name=ids,proto3" json:"ids" gorm:"-"`
	Type        string  `protobuf:"bytes,23,opt,name=type,proto3" json:"type"`
	TeamRate    float32 `protobuf:"fixed32,24,opt,name=team_rate,json=teamRate,proto3" json:"team_rate"`
	IsChecked   bool    `protobuf:"varint,25,opt,name=is_checked,json=isChecked,proto3" json:"is_checked"`
	AgreementId int64   `protobuf:"varint,26,opt,name=agreement_id,json=agreementId,proto3" json:"agreement_id"`
	IntroduceId int64   `protobuf:"varint,27,opt,name=introduce_id,json=introduceId,proto3" json:"introduce_id"`
	Weight      int32   `protobuf:"varint,28,opt,name=weight,proto3" json:"weight"`
	IsBonusSame bool    `protobuf:"varint,29,opt,name=is_bonus_same,json=isBonusSame,proto3" json:"is_bonus_same"`
}

func (x *Rank) Reset() {
	*x = Rank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rankService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rank) ProtoMessage() {}

func (x *Rank) ProtoReflect() protoreflect.Message {
	mi := &file_rankService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rank.ProtoReflect.Descriptor instead.
func (*Rank) Descriptor() ([]byte, []int) {
	return file_rankService_proto_rawDescGZIP(), []int{1}
}

func (x *Rank) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Rank) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Rank) GetPrimaryRate() float32 {
	if x != nil {
		return x.PrimaryRate
	}
	return 0
}

func (x *Rank) GetSecondRate() float32 {
	if x != nil {
		return x.SecondRate
	}
	return 0
}

func (x *Rank) GetThreeRate() float32 {
	if x != nil {
		return x.ThreeRate
	}
	return 0
}

func (x *Rank) GetRewardMoney() float32 {
	if x != nil {
		return x.RewardMoney
	}
	return 0
}

func (x *Rank) GetMonthDrawMoney() float32 {
	if x != nil {
		return x.MonthDrawMoney
	}
	return 0
}

func (x *Rank) GetMonthDrawNum() int32 {
	if x != nil {
		return x.MonthDrawNum
	}
	return 0
}

func (x *Rank) GetLevelId() int32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

func (x *Rank) GetIsCondition() bool {
	if x != nil {
		return x.IsCondition
	}
	return false
}

func (x *Rank) GetPromotionAmount() float32 {
	if x != nil {
		return x.PromotionAmount
	}
	return 0
}

func (x *Rank) GetPrimaryNum() int32 {
	if x != nil {
		return x.PrimaryNum
	}
	return 0
}

func (x *Rank) GetSecondNum() int32 {
	if x != nil {
		return x.SecondNum
	}
	return 0
}

func (x *Rank) GetIsBought() bool {
	if x != nil {
		return x.IsBought
	}
	return false
}

func (x *Rank) GetConsumeAmount() float32 {
	if x != nil {
		return x.ConsumeAmount
	}
	return 0
}

func (x *Rank) GetConsumeNum() int32 {
	if x != nil {
		return x.ConsumeNum
	}
	return 0
}

func (x *Rank) GetIsBindIdcard() bool {
	if x != nil {
		return x.IsBindIdcard
	}
	return false
}

func (x *Rank) GetIsBindMobile() bool {
	if x != nil {
		return x.IsBindMobile
	}
	return false
}

func (x *Rank) GetJoinFee() float32 {
	if x != nil {
		return x.JoinFee
	}
	return 0
}

func (x *Rank) GetDefaulted() bool {
	if x != nil {
		return x.Defaulted
	}
	return false
}

func (x *Rank) GetGoodses() []*GoodsCondition {
	if x != nil {
		return x.Goodses
	}
	return nil
}

func (x *Rank) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *Rank) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Rank) GetTeamRate() float32 {
	if x != nil {
		return x.TeamRate
	}
	return 0
}

func (x *Rank) GetIsChecked() bool {
	if x != nil {
		return x.IsChecked
	}
	return false
}

func (x *Rank) GetAgreementId() int64 {
	if x != nil {
		return x.AgreementId
	}
	return 0
}

func (x *Rank) GetIntroduceId() int64 {
	if x != nil {
		return x.IntroduceId
	}
	return 0
}

func (x *Rank) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Rank) GetIsBonusSame() bool {
	if x != nil {
		return x.IsBonusSame
	}
	return false
}

type RankResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Rank         `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Rank       `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *RankResponse) Reset() {
	*x = RankResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rankService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankResponse) ProtoMessage() {}

func (x *RankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rankService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankResponse.ProtoReflect.Descriptor instead.
func (*RankResponse) Descriptor() ([]byte, []int) {
	return file_rankService_proto_rawDescGZIP(), []int{2}
}

func (x *RankResponse) GetEntity() *Rank {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *RankResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *RankResponse) GetItems() []*Rank {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *RankResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *RankResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_rankService_proto protoreflect.FileDescriptor

var file_rankService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01,
	0x0a, 0x09, 0x52, 0x61, 0x6e, 0x6b, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xab, 0x07, 0x0a, 0x04, 0x52, 0x61, 0x6e, 0x6b, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x68, 0x72, 0x65, 0x65,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x74, 0x68, 0x72,
	0x65, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x6f, 0x6e,
	0x74, 0x68, 0x5f, 0x64, 0x72, 0x61, 0x77, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0e, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x44, 0x72, 0x61, 0x77, 0x4d, 0x6f,
	0x6e, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x5f, 0x64, 0x72, 0x61,
	0x77, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x6f, 0x6e,
	0x74, 0x68, 0x44, 0x72, 0x61, 0x77, 0x4e, 0x75, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e,
	0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x62, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x42, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x62, 0x69,
	0x6e, 0x64, 0x5f, 0x69, 0x64, 0x63, 0x61, 0x72, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x69, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x49, 0x64, 0x63, 0x61, 0x72, 0x64, 0x12, 0x24, 0x0a,
	0x0e, 0x69, 0x73, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x42, 0x69, 0x6e, 0x64, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x66, 0x65, 0x65, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x6a, 0x6f, 0x69, 0x6e, 0x46, 0x65, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x65, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x65, 0x64, 0x12, 0x32, 0x0a, 0x07,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x65, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x65, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x16, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65,
	0x64, 0x18, 0x19, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x67, 0x72, 0x65, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x69, 0x6e, 0x74,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x73, 0x61, 0x6d,
	0x65, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x42, 0x6f, 0x6e, 0x75, 0x73,
	0x53, 0x61, 0x6d, 0x65, 0x22, 0xc8, 0x01, 0x0a, 0x0c, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x72, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x61, 0x6e,
	0x6b, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32,
	0xc5, 0x02, 0x0a, 0x0b, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x32, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x1a, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x61, 0x6e,
	0x6b, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x61, 0x6e,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x61,
	0x6e, 0x6b, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x61,
	0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x52, 0x61, 0x6e, 0x6b, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37,
	0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rankService_proto_rawDescOnce sync.Once
	file_rankService_proto_rawDescData = file_rankService_proto_rawDesc
)

func file_rankService_proto_rawDescGZIP() []byte {
	file_rankService_proto_rawDescOnce.Do(func() {
		file_rankService_proto_rawDescData = protoimpl.X.CompressGZIP(file_rankService_proto_rawDescData)
	})
	return file_rankService_proto_rawDescData
}

var file_rankService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rankService_proto_goTypes = []interface{}{
	(*RankWhere)(nil),      // 0: services.RankWhere
	(*Rank)(nil),           // 1: services.Rank
	(*RankResponse)(nil),   // 2: services.RankResponse
	(*GoodsCondition)(nil), // 3: services.GoodsCondition
	(*common.Pager)(nil),   // 4: common.Pager
	(*common.Error)(nil),   // 5: common.Error
	(*common.Info)(nil),    // 6: common.Info
}
var file_rankService_proto_depIdxs = []int32{
	3,  // 0: services.Rank.goodses:type_name -> services.GoodsCondition
	1,  // 1: services.RankResponse.entity:type_name -> services.Rank
	4,  // 2: services.RankResponse.pager:type_name -> common.Pager
	1,  // 3: services.RankResponse.items:type_name -> services.Rank
	5,  // 4: services.RankResponse.error:type_name -> common.Error
	6,  // 5: services.RankResponse.info:type_name -> common.Info
	1,  // 6: services.RankService.Create:input_type -> services.Rank
	1,  // 7: services.RankService.Update:input_type -> services.Rank
	1,  // 8: services.RankService.Delete:input_type -> services.Rank
	1,  // 9: services.RankService.Get:input_type -> services.Rank
	1,  // 10: services.RankService.List:input_type -> services.Rank
	0,  // 11: services.RankService.Search:input_type -> services.RankWhere
	2,  // 12: services.RankService.Create:output_type -> services.RankResponse
	2,  // 13: services.RankService.Update:output_type -> services.RankResponse
	2,  // 14: services.RankService.Delete:output_type -> services.RankResponse
	2,  // 15: services.RankService.Get:output_type -> services.RankResponse
	2,  // 16: services.RankService.List:output_type -> services.RankResponse
	2,  // 17: services.RankService.Search:output_type -> services.RankResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_rankService_proto_init() }
func file_rankService_proto_init() {
	if File_rankService_proto != nil {
		return
	}
	file_goodsConditionService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rankService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankWhere); i {
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
		file_rankService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rank); i {
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
		file_rankService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankResponse); i {
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
			RawDescriptor: file_rankService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rankService_proto_goTypes,
		DependencyIndexes: file_rankService_proto_depIdxs,
		MessageInfos:      file_rankService_proto_msgTypes,
	}.Build()
	File_rankService_proto = out.File
	file_rankService_proto_rawDesc = nil
	file_rankService_proto_goTypes = nil
	file_rankService_proto_depIdxs = nil
}
