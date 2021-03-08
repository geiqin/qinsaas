// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: groupService.proto

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

type GroupWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	Keywords string `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Id       int64  `protobuf:"varint,5,opt,name=id,proto3" json:"id"`
	// @inject_tag: gorm:"-"
	Ids     []int64 `protobuf:"varint,6,rep,packed,name=ids,proto3" json:"ids" gorm:"-"`
	Type    int32   `protobuf:"varint,7,opt,name=type,proto3" json:"type"`
	Status  int32   `protobuf:"varint,8,opt,name=status,proto3" json:"status"`
	ItemId  int64   `protobuf:"varint,9,opt,name=item_id,json=itemId,proto3" json:"item_id"`
	IsValid bool    `protobuf:"varint,10,opt,name=is_valid,json=isValid,proto3" json:"is_valid"`
	StartAt string  `protobuf:"bytes,11,opt,name=start_at,json=startAt,proto3" json:"start_at"`
	EndAt   string  `protobuf:"bytes,12,opt,name=end_at,json=endAt,proto3" json:"end_at"`
}

func (x *GroupWhere) Reset() {
	*x = GroupWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupWhere) ProtoMessage() {}

func (x *GroupWhere) ProtoReflect() protoreflect.Message {
	mi := &file_groupService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupWhere.ProtoReflect.Descriptor instead.
func (*GroupWhere) Descriptor() ([]byte, []int) {
	return file_groupService_proto_rawDescGZIP(), []int{0}
}

func (x *GroupWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *GroupWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GroupWhere) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *GroupWhere) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *GroupWhere) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GroupWhere) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *GroupWhere) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *GroupWhere) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GroupWhere) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *GroupWhere) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

func (x *GroupWhere) GetStartAt() string {
	if x != nil {
		return x.StartAt
	}
	return ""
}

func (x *GroupWhere) GetEndAt() string {
	if x != nil {
		return x.EndAt
	}
	return ""
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name             string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Type             int32          `protobuf:"varint,3,opt,name=type,proto3" json:"type"`
	StartAt          string         `protobuf:"bytes,4,opt,name=start_at,json=startAt,proto3" json:"start_at"`
	EndAt            string         `protobuf:"bytes,5,opt,name=end_at,json=endAt,proto3" json:"end_at"`
	ItemId           int64          `protobuf:"varint,6,opt,name=item_id,json=itemId,proto3" json:"item_id"`
	ExpireDay        int32          `protobuf:"varint,7,opt,name=expire_day,json=expireDay,proto3" json:"expire_day"`
	ExpireHour       int32          `protobuf:"varint,8,opt,name=expire_hour,json=expireHour,proto3" json:"expire_hour"`
	ExpireMinute     int32          `protobuf:"varint,9,opt,name=expire_minute,json=expireMinute,proto3" json:"expire_minute"`
	IsStack          bool           `protobuf:"varint,10,opt,name=is_stack,json=isStack,proto3" json:"is_stack"`
	Stack            int32          `protobuf:"varint,11,opt,name=stack,proto3" json:"stack"`
	IsLimit          bool           `protobuf:"varint,12,opt,name=is_limit,json=isLimit,proto3" json:"is_limit"`
	LimitTotalNum    int32          `protobuf:"varint,13,opt,name=limit_total_num,json=limitTotalNum,proto3" json:"limit_total_num"`
	LimitUnitNum     int32          `protobuf:"varint,14,opt,name=limit_unit_num,json=limitUnitNum,proto3" json:"limit_unit_num"`
	IsTogether       bool           `protobuf:"varint,15,opt,name=is_together,json=isTogether,proto3" json:"is_together"`
	IsSimulate       bool           `protobuf:"varint,16,opt,name=is_simulate,json=isSimulate,proto3" json:"is_simulate"`
	LowestNum        int32          `protobuf:"varint,17,opt,name=lowest_num,json=lowestNum,proto3" json:"lowest_num"`
	IsCollection     bool           `protobuf:"varint,18,opt,name=is_collection,json=isCollection,proto3" json:"is_collection"`
	CollectionType   int32          `protobuf:"varint,19,opt,name=collection_type,json=collectionType,proto3" json:"collection_type"`
	IsLeaderDiscount bool           `protobuf:"varint,20,opt,name=is_leader_discount,json=isLeaderDiscount,proto3" json:"is_leader_discount"`
	LeaderPrice      float32        `protobuf:"fixed32,21,opt,name=leader_price,json=leaderPrice,proto3" json:"leader_price"`
	IsFree           bool           `protobuf:"varint,22,opt,name=is_free,json=isFree,proto3" json:"is_free"`
	IsPreview        bool           `protobuf:"varint,23,opt,name=is_preview,json=isPreview,proto3" json:"is_preview"`
	Status           int32          `protobuf:"varint,24,opt,name=status,proto3" json:"status"`
	CreatedAt        string         `protobuf:"bytes,25,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt        string         `protobuf:"bytes,26,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Ladders          []*GroupLadder `protobuf:"bytes,27,rep,name=ladders,proto3" json:"ladders"`
	IsStackCoupon    bool           `protobuf:"varint,28,opt,name=is_stack_coupon,json=isStackCoupon,proto3" json:"is_stack_coupon"`       // 是否叠加优惠券
	IsStackIntegral  bool           `protobuf:"varint,29,opt,name=is_stack_integral,json=isStackIntegral,proto3" json:"is_stack_integral"` // 是否叠加积分抵现
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_groupService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_groupService_proto_rawDescGZIP(), []int{1}
}

func (x *Group) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Group) GetStartAt() string {
	if x != nil {
		return x.StartAt
	}
	return ""
}

func (x *Group) GetEndAt() string {
	if x != nil {
		return x.EndAt
	}
	return ""
}

func (x *Group) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *Group) GetExpireDay() int32 {
	if x != nil {
		return x.ExpireDay
	}
	return 0
}

func (x *Group) GetExpireHour() int32 {
	if x != nil {
		return x.ExpireHour
	}
	return 0
}

func (x *Group) GetExpireMinute() int32 {
	if x != nil {
		return x.ExpireMinute
	}
	return 0
}

func (x *Group) GetIsStack() bool {
	if x != nil {
		return x.IsStack
	}
	return false
}

func (x *Group) GetStack() int32 {
	if x != nil {
		return x.Stack
	}
	return 0
}

func (x *Group) GetIsLimit() bool {
	if x != nil {
		return x.IsLimit
	}
	return false
}

func (x *Group) GetLimitTotalNum() int32 {
	if x != nil {
		return x.LimitTotalNum
	}
	return 0
}

func (x *Group) GetLimitUnitNum() int32 {
	if x != nil {
		return x.LimitUnitNum
	}
	return 0
}

func (x *Group) GetIsTogether() bool {
	if x != nil {
		return x.IsTogether
	}
	return false
}

func (x *Group) GetIsSimulate() bool {
	if x != nil {
		return x.IsSimulate
	}
	return false
}

func (x *Group) GetLowestNum() int32 {
	if x != nil {
		return x.LowestNum
	}
	return 0
}

func (x *Group) GetIsCollection() bool {
	if x != nil {
		return x.IsCollection
	}
	return false
}

func (x *Group) GetCollectionType() int32 {
	if x != nil {
		return x.CollectionType
	}
	return 0
}

func (x *Group) GetIsLeaderDiscount() bool {
	if x != nil {
		return x.IsLeaderDiscount
	}
	return false
}

func (x *Group) GetLeaderPrice() float32 {
	if x != nil {
		return x.LeaderPrice
	}
	return 0
}

func (x *Group) GetIsFree() bool {
	if x != nil {
		return x.IsFree
	}
	return false
}

func (x *Group) GetIsPreview() bool {
	if x != nil {
		return x.IsPreview
	}
	return false
}

func (x *Group) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Group) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Group) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Group) GetLadders() []*GroupLadder {
	if x != nil {
		return x.Ladders
	}
	return nil
}

func (x *Group) GetIsStackCoupon() bool {
	if x != nil {
		return x.IsStackCoupon
	}
	return false
}

func (x *Group) GetIsStackIntegral() bool {
	if x != nil {
		return x.IsStackIntegral
	}
	return false
}

type GroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Group        `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Group      `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *GroupResponse) Reset() {
	*x = GroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupResponse) ProtoMessage() {}

func (x *GroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_groupService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupResponse.ProtoReflect.Descriptor instead.
func (*GroupResponse) Descriptor() ([]byte, []int) {
	return file_groupService_proto_rawDescGZIP(), []int{2}
}

func (x *GroupResponse) GetEntity() *Group {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *GroupResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *GroupResponse) GetItems() []*Group {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *GroupResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *GroupResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_groupService_proto protoreflect.FileDescriptor

var file_groupService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x61, 0x64, 0x64, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x02, 0x0a, 0x0a,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73,
	0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74,
	0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x22, 0x9c, 0x07, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x41, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74,
	0x65, 0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x64,
	0x61, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x44, 0x61, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x68, 0x6f,
	0x75, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x48, 0x6f, 0x75, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x6d,
	0x69, 0x6e, 0x75, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x53,
	0x74, 0x61, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x24, 0x0a,
	0x0e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x55, 0x6e, 0x69, 0x74,
	0x4e, 0x75, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x74, 0x6f, 0x67, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x54, 0x6f, 0x67, 0x65,
	0x74, 0x68, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x53, 0x69, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x5f,
	0x6e, 0x75, 0x6d, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x6f, 0x77, 0x65, 0x73,
	0x74, 0x4e, 0x75, 0x6d, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x73, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10,
	0x69, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x46, 0x72, 0x65, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x17, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x18, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x2f, 0x0a, 0x07, 0x6c, 0x61, 0x64, 0x64, 0x65, 0x72, 0x73, 0x18, 0x1b, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4c, 0x61, 0x64, 0x64, 0x65, 0x72, 0x52, 0x07, 0x6c, 0x61, 0x64, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x63,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x53,
	0x74, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73,
	0x5f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x18,
	0x1d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x22, 0xcb, 0x01, 0x0a, 0x0d, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x32, 0x9f, 0x03, 0x0a, 0x0c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x36, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x09, 0x54, 0x6f, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x57, 0x68, 0x65,
	0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x4b, 0x0a, 0x11, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_groupService_proto_rawDescOnce sync.Once
	file_groupService_proto_rawDescData = file_groupService_proto_rawDesc
)

func file_groupService_proto_rawDescGZIP() []byte {
	file_groupService_proto_rawDescOnce.Do(func() {
		file_groupService_proto_rawDescData = protoimpl.X.CompressGZIP(file_groupService_proto_rawDescData)
	})
	return file_groupService_proto_rawDescData
}

var file_groupService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_groupService_proto_goTypes = []interface{}{
	(*GroupWhere)(nil),    // 0: services.GroupWhere
	(*Group)(nil),         // 1: services.Group
	(*GroupResponse)(nil), // 2: services.GroupResponse
	(*GroupLadder)(nil),   // 3: services.GroupLadder
	(*common.Pager)(nil),  // 4: common.Pager
	(*common.Error)(nil),  // 5: common.Error
	(*common.Info)(nil),   // 6: common.Info
}
var file_groupService_proto_depIdxs = []int32{
	3,  // 0: services.Group.ladders:type_name -> services.GroupLadder
	1,  // 1: services.GroupResponse.entity:type_name -> services.Group
	4,  // 2: services.GroupResponse.pager:type_name -> common.Pager
	1,  // 3: services.GroupResponse.items:type_name -> services.Group
	5,  // 4: services.GroupResponse.error:type_name -> common.Error
	6,  // 5: services.GroupResponse.info:type_name -> common.Info
	1,  // 6: services.GroupService.Create:input_type -> services.Group
	1,  // 7: services.GroupService.Update:input_type -> services.Group
	0,  // 8: services.GroupService.Get:input_type -> services.GroupWhere
	0,  // 9: services.GroupService.Delete:input_type -> services.GroupWhere
	0,  // 10: services.GroupService.ToInvalid:input_type -> services.GroupWhere
	0,  // 11: services.GroupService.Search:input_type -> services.GroupWhere
	0,  // 12: services.GroupService.List:input_type -> services.GroupWhere
	0,  // 13: services.FrontGroupService.Get:input_type -> services.GroupWhere
	2,  // 14: services.GroupService.Create:output_type -> services.GroupResponse
	2,  // 15: services.GroupService.Update:output_type -> services.GroupResponse
	2,  // 16: services.GroupService.Get:output_type -> services.GroupResponse
	2,  // 17: services.GroupService.Delete:output_type -> services.GroupResponse
	2,  // 18: services.GroupService.ToInvalid:output_type -> services.GroupResponse
	2,  // 19: services.GroupService.Search:output_type -> services.GroupResponse
	2,  // 20: services.GroupService.List:output_type -> services.GroupResponse
	2,  // 21: services.FrontGroupService.Get:output_type -> services.GroupResponse
	14, // [14:22] is the sub-list for method output_type
	6,  // [6:14] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_groupService_proto_init() }
func file_groupService_proto_init() {
	if File_groupService_proto != nil {
		return
	}
	file_groupLadderService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_groupService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupWhere); i {
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
		file_groupService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_groupService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupResponse); i {
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
			RawDescriptor: file_groupService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_groupService_proto_goTypes,
		DependencyIndexes: file_groupService_proto_depIdxs,
		MessageInfos:      file_groupService_proto_msgTypes,
	}.Build()
	File_groupService_proto = out.File
	file_groupService_proto_rawDesc = nil
	file_groupService_proto_goTypes = nil
	file_groupService_proto_depIdxs = nil
}