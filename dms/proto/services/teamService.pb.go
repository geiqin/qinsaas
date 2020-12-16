// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: teamService.proto

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

type TeamWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged          int32   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize       int32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Top            int32   `protobuf:"varint,3,opt,name=top,proto3" json:"top,omitempty"`
	Id             int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Mobile         string  `protobuf:"bytes,5,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Keywords       string  `protobuf:"bytes,6,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Ids            []int64 `protobuf:"varint,7,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	DistributorId  int64   `protobuf:"varint,8,opt,name=distributor_id,json=distributorId,proto3" json:"distributor_id,omitempty"`
	DistributorIds []int64 `protobuf:"varint,9,rep,packed,name=distributor_ids,json=distributorIds,proto3" json:"distributor_ids,omitempty"`
	LeaderId       int64   `protobuf:"varint,10,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	LeaderIds      []int64 `protobuf:"varint,11,rep,packed,name=leader_ids,json=leaderIds,proto3" json:"leader_ids,omitempty"`
}

func (x *TeamWhere) Reset() {
	*x = TeamWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teamService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamWhere) ProtoMessage() {}

func (x *TeamWhere) ProtoReflect() protoreflect.Message {
	mi := &file_teamService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamWhere.ProtoReflect.Descriptor instead.
func (*TeamWhere) Descriptor() ([]byte, []int) {
	return file_teamService_proto_rawDescGZIP(), []int{0}
}

func (x *TeamWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *TeamWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *TeamWhere) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *TeamWhere) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TeamWhere) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *TeamWhere) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *TeamWhere) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *TeamWhere) GetDistributorId() int64 {
	if x != nil {
		return x.DistributorId
	}
	return 0
}

func (x *TeamWhere) GetDistributorIds() []int64 {
	if x != nil {
		return x.DistributorIds
	}
	return nil
}

func (x *TeamWhere) GetLeaderId() int64 {
	if x != nil {
		return x.LeaderId
	}
	return 0
}

func (x *TeamWhere) GetLeaderIds() []int64 {
	if x != nil {
		return x.LeaderIds
	}
	return nil
}

type TeamBasic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DistributorId int64        `protobuf:"varint,2,opt,name=distributor_id,json=distributorId,proto3" json:"distributor_id,omitempty"`
	LeaderId      int64        `protobuf:"varint,3,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	Distributor   *Distributor `protobuf:"bytes,4,opt,name=distributor,proto3" json:"distributor,omitempty"`
	CreatedAt     string       `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string       `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *TeamBasic) Reset() {
	*x = TeamBasic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teamService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamBasic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamBasic) ProtoMessage() {}

func (x *TeamBasic) ProtoReflect() protoreflect.Message {
	mi := &file_teamService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamBasic.ProtoReflect.Descriptor instead.
func (*TeamBasic) Descriptor() ([]byte, []int) {
	return file_teamService_proto_rawDescGZIP(), []int{1}
}

func (x *TeamBasic) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TeamBasic) GetDistributorId() int64 {
	if x != nil {
		return x.DistributorId
	}
	return 0
}

func (x *TeamBasic) GetLeaderId() int64 {
	if x != nil {
		return x.LeaderId
	}
	return 0
}

func (x *TeamBasic) GetDistributor() *Distributor {
	if x != nil {
		return x.Distributor
	}
	return nil
}

func (x *TeamBasic) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *TeamBasic) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type MyTeam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DistributorId   int64   `protobuf:"varint,2,opt,name=distributor_id,json=distributorId,proto3" json:"distributor_id,omitempty"`
	DistributorName string  `protobuf:"bytes,3,opt,name=distributor_name,json=distributorName,proto3" json:"distributor_name,omitempty"` // 分销员昵称
	Mobile          string  `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile,omitempty"`                                          // 手机号
	AvatarUrl       string  `protobuf:"bytes,5,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`                   // 用户头像
	RecommendNum    int32   `protobuf:"varint,6,opt,name=recommend_num,json=recommendNum,proto3" json:"recommend_num,omitempty"`         // 邀请团员数
	OrderMoney      float32 `protobuf:"fixed32,8,opt,name=order_money,json=orderMoney,proto3" json:"order_money,omitempty"`              // 成交金额
	CustomerNum     int32   `protobuf:"varint,7,opt,name=customer_num,json=customerNum,proto3" json:"customer_num,omitempty"`            // 绑定客户数
	OrderNum        int32   `protobuf:"varint,9,opt,name=order_num,json=orderNum,proto3" json:"order_num,omitempty"`                     // 成交笔数
	CreatedAt       string  `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       string  `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *MyTeam) Reset() {
	*x = MyTeam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teamService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyTeam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyTeam) ProtoMessage() {}

func (x *MyTeam) ProtoReflect() protoreflect.Message {
	mi := &file_teamService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyTeam.ProtoReflect.Descriptor instead.
func (*MyTeam) Descriptor() ([]byte, []int) {
	return file_teamService_proto_rawDescGZIP(), []int{2}
}

func (x *MyTeam) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MyTeam) GetDistributorId() int64 {
	if x != nil {
		return x.DistributorId
	}
	return 0
}

func (x *MyTeam) GetDistributorName() string {
	if x != nil {
		return x.DistributorName
	}
	return ""
}

func (x *MyTeam) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *MyTeam) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *MyTeam) GetRecommendNum() int32 {
	if x != nil {
		return x.RecommendNum
	}
	return 0
}

func (x *MyTeam) GetOrderMoney() float32 {
	if x != nil {
		return x.OrderMoney
	}
	return 0
}

func (x *MyTeam) GetCustomerNum() int32 {
	if x != nil {
		return x.CustomerNum
	}
	return 0
}

func (x *MyTeam) GetOrderNum() int32 {
	if x != nil {
		return x.OrderNum
	}
	return 0
}

func (x *MyTeam) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *MyTeam) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type Team struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                   // 团队名称
	LeaderId    int64   `protobuf:"varint,3,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`          // 团长ID
	LeaderName  string  `protobuf:"bytes,4,opt,name=leader_name,json=leaderName,proto3" json:"leader_name,omitempty"`     // 团长昵称
	Mobile      string  `protobuf:"bytes,5,opt,name=mobile,proto3" json:"mobile,omitempty"`                               // 手机号
	AvatarUrl   string  `protobuf:"bytes,6,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`        // 用户头像
	TeamNum     int32   `protobuf:"varint,7,opt,name=team_num,json=teamNum,proto3" json:"team_num,omitempty"`             // 团员数
	CustomerNum int32   `protobuf:"varint,8,opt,name=customer_num,json=customerNum,proto3" json:"customer_num,omitempty"` // 团队客户数
	OrderMoney  float32 `protobuf:"fixed32,9,opt,name=order_money,json=orderMoney,proto3" json:"order_money,omitempty"`   // 团队成交金额
	OrderNum    int32   `protobuf:"varint,10,opt,name=order_num,json=orderNum,proto3" json:"order_num,omitempty"`         // 团队成交笔数
	CreatedAt   string  `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string  `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Team) Reset() {
	*x = Team{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teamService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Team) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team) ProtoMessage() {}

func (x *Team) ProtoReflect() protoreflect.Message {
	mi := &file_teamService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Team.ProtoReflect.Descriptor instead.
func (*Team) Descriptor() ([]byte, []int) {
	return file_teamService_proto_rawDescGZIP(), []int{3}
}

func (x *Team) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Team) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Team) GetLeaderId() int64 {
	if x != nil {
		return x.LeaderId
	}
	return 0
}

func (x *Team) GetLeaderName() string {
	if x != nil {
		return x.LeaderName
	}
	return ""
}

func (x *Team) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *Team) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *Team) GetTeamNum() int32 {
	if x != nil {
		return x.TeamNum
	}
	return 0
}

func (x *Team) GetCustomerNum() int32 {
	if x != nil {
		return x.CustomerNum
	}
	return 0
}

func (x *Team) GetOrderMoney() float32 {
	if x != nil {
		return x.OrderMoney
	}
	return 0
}

func (x *Team) GetOrderNum() int32 {
	if x != nil {
		return x.OrderNum
	}
	return 0
}

func (x *Team) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Team) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type TeamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Team         `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*Team       `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *TeamResponse) Reset() {
	*x = TeamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teamService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamResponse) ProtoMessage() {}

func (x *TeamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teamService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamResponse.ProtoReflect.Descriptor instead.
func (*TeamResponse) Descriptor() ([]byte, []int) {
	return file_teamService_proto_rawDescGZIP(), []int{4}
}

func (x *TeamResponse) GetEntity() *Team {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *TeamResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *TeamResponse) GetItems() []*Team {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *TeamResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *TeamResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type MyTeamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *MyTeam       `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*MyTeam     `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *MyTeamResponse) Reset() {
	*x = MyTeamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teamService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyTeamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyTeamResponse) ProtoMessage() {}

func (x *MyTeamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teamService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyTeamResponse.ProtoReflect.Descriptor instead.
func (*MyTeamResponse) Descriptor() ([]byte, []int) {
	return file_teamService_proto_rawDescGZIP(), []int{5}
}

func (x *MyTeamResponse) GetEntity() *MyTeam {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *MyTeamResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *MyTeamResponse) GetItems() []*MyTeam {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *MyTeamResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *MyTeamResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_teamService_proto protoreflect.FileDescriptor

var file_teamService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x02, 0x0a, 0x09, 0x54,
	0x65, 0x61, 0x6d, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72,
	0x49, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0b,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22,
	0xd6, 0x01, 0x0a, 0x09, 0x54, 0x65, 0x61, 0x6d, 0x42, 0x61, 0x73, 0x69, 0x63, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a,
	0x0e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x6f, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x37, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xe5, 0x02, 0x0a, 0x06, 0x4d, 0x79, 0x54,
	0x65, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x4e, 0x75,
	0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x6f, 0x6e,
	0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e,
	0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0xd9, 0x02, 0x0a, 0x04, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55,
	0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x4e, 0x75, 0x6d, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x75, 0x6d,
	0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x65,
	0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xc8, 0x01, 0x0a,
	0x0c, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0xce, 0x01, 0x0a, 0x0e, 0x4d, 0x79, 0x54, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4d, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x4a, 0x0a, 0x0d, 0x4d, 0x79, 0x54, 0x65,
	0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54,
	0x65, 0x61, 0x6d, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4d, 0x79, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x32, 0x88, 0x01, 0x0a, 0x0b, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x13,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x57, 0x68,
	0x65, 0x72, 0x65, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a,
	0x0d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x13,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x57, 0x68,
	0x65, 0x72, 0x65, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x79, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teamService_proto_rawDescOnce sync.Once
	file_teamService_proto_rawDescData = file_teamService_proto_rawDesc
)

func file_teamService_proto_rawDescGZIP() []byte {
	file_teamService_proto_rawDescOnce.Do(func() {
		file_teamService_proto_rawDescData = protoimpl.X.CompressGZIP(file_teamService_proto_rawDescData)
	})
	return file_teamService_proto_rawDescData
}

var file_teamService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_teamService_proto_goTypes = []interface{}{
	(*TeamWhere)(nil),      // 0: services.TeamWhere
	(*TeamBasic)(nil),      // 1: services.TeamBasic
	(*MyTeam)(nil),         // 2: services.MyTeam
	(*Team)(nil),           // 3: services.Team
	(*TeamResponse)(nil),   // 4: services.TeamResponse
	(*MyTeamResponse)(nil), // 5: services.MyTeamResponse
	(*Distributor)(nil),    // 6: services.Distributor
	(*common.Pager)(nil),   // 7: common.Pager
	(*common.Error)(nil),   // 8: common.Error
	(*common.Info)(nil),    // 9: common.Info
}
var file_teamService_proto_depIdxs = []int32{
	6,  // 0: services.TeamBasic.distributor:type_name -> services.Distributor
	3,  // 1: services.TeamResponse.entity:type_name -> services.Team
	7,  // 2: services.TeamResponse.pager:type_name -> common.Pager
	3,  // 3: services.TeamResponse.items:type_name -> services.Team
	8,  // 4: services.TeamResponse.error:type_name -> common.Error
	9,  // 5: services.TeamResponse.info:type_name -> common.Info
	2,  // 6: services.MyTeamResponse.entity:type_name -> services.MyTeam
	7,  // 7: services.MyTeamResponse.pager:type_name -> common.Pager
	2,  // 8: services.MyTeamResponse.items:type_name -> services.MyTeam
	8,  // 9: services.MyTeamResponse.error:type_name -> common.Error
	9,  // 10: services.MyTeamResponse.info:type_name -> common.Info
	0,  // 11: services.MyTeamService.Search:input_type -> services.TeamWhere
	0,  // 12: services.TeamService.Search:input_type -> services.TeamWhere
	0,  // 13: services.TeamService.MembersSearch:input_type -> services.TeamWhere
	5,  // 14: services.MyTeamService.Search:output_type -> services.MyTeamResponse
	4,  // 15: services.TeamService.Search:output_type -> services.TeamResponse
	5,  // 16: services.TeamService.MembersSearch:output_type -> services.MyTeamResponse
	14, // [14:17] is the sub-list for method output_type
	11, // [11:14] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_teamService_proto_init() }
func file_teamService_proto_init() {
	if File_teamService_proto != nil {
		return
	}
	file_distributorService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_teamService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamWhere); i {
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
		file_teamService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamBasic); i {
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
		file_teamService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyTeam); i {
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
		file_teamService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Team); i {
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
		file_teamService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamResponse); i {
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
		file_teamService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyTeamResponse); i {
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
			RawDescriptor: file_teamService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_teamService_proto_goTypes,
		DependencyIndexes: file_teamService_proto_depIdxs,
		MessageInfos:      file_teamService_proto_msgTypes,
	}.Build()
	File_teamService_proto = out.File
	file_teamService_proto_rawDesc = nil
	file_teamService_proto_goTypes = nil
	file_teamService_proto_depIdxs = nil
}
