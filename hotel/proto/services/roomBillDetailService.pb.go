// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: roomBillDetailService.proto

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

type RoomBillDetailWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged      int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize   int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	RoomBillId int64 `protobuf:"varint,3,opt,name=room_bill_id,json=roomBillId,proto3" json:"room_bill_id,omitempty"`
	Status     int32 `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	RoomBookId int64 `protobuf:"varint,5,opt,name=room_book_id,json=roomBookId,proto3" json:"room_book_id,omitempty"`
}

func (x *RoomBillDetailWhere) Reset() {
	*x = RoomBillDetailWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomBillDetailService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomBillDetailWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomBillDetailWhere) ProtoMessage() {}

func (x *RoomBillDetailWhere) ProtoReflect() protoreflect.Message {
	mi := &file_roomBillDetailService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomBillDetailWhere.ProtoReflect.Descriptor instead.
func (*RoomBillDetailWhere) Descriptor() ([]byte, []int) {
	return file_roomBillDetailService_proto_rawDescGZIP(), []int{0}
}

func (x *RoomBillDetailWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *RoomBillDetailWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *RoomBillDetailWhere) GetRoomBillId() int64 {
	if x != nil {
		return x.RoomBillId
	}
	return 0
}

func (x *RoomBillDetailWhere) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RoomBillDetailWhere) GetRoomBookId() int64 {
	if x != nil {
		return x.RoomBookId
	}
	return 0
}

type RoomBillDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RoomBillId    int64   `protobuf:"varint,2,opt,name=room_bill_id,json=roomBillId,proto3" json:"room_bill_id,omitempty"`
	BillItemCode  string  `protobuf:"bytes,3,opt,name=bill_item_code,json=billItemCode,proto3" json:"bill_item_code,omitempty"`
	ConsumeAmount float32 `protobuf:"fixed32,4,opt,name=consume_amount,json=consumeAmount,proto3" json:"consume_amount,omitempty"`
	PayAmount     float32 `protobuf:"fixed32,5,opt,name=pay_amount,json=payAmount,proto3" json:"pay_amount,omitempty"`
	Memo          string  `protobuf:"bytes,6,opt,name=memo,proto3" json:"memo,omitempty"`
	Status        int32   `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
	OptId         int64   `protobuf:"varint,8,opt,name=opt_id,json=optId,proto3" json:"opt_id,omitempty"`
	CreatedAt     string  `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string  `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *RoomBillDetail) Reset() {
	*x = RoomBillDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomBillDetailService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomBillDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomBillDetail) ProtoMessage() {}

func (x *RoomBillDetail) ProtoReflect() protoreflect.Message {
	mi := &file_roomBillDetailService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomBillDetail.ProtoReflect.Descriptor instead.
func (*RoomBillDetail) Descriptor() ([]byte, []int) {
	return file_roomBillDetailService_proto_rawDescGZIP(), []int{1}
}

func (x *RoomBillDetail) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RoomBillDetail) GetRoomBillId() int64 {
	if x != nil {
		return x.RoomBillId
	}
	return 0
}

func (x *RoomBillDetail) GetBillItemCode() string {
	if x != nil {
		return x.BillItemCode
	}
	return ""
}

func (x *RoomBillDetail) GetConsumeAmount() float32 {
	if x != nil {
		return x.ConsumeAmount
	}
	return 0
}

func (x *RoomBillDetail) GetPayAmount() float32 {
	if x != nil {
		return x.PayAmount
	}
	return 0
}

func (x *RoomBillDetail) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *RoomBillDetail) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RoomBillDetail) GetOptId() int64 {
	if x != nil {
		return x.OptId
	}
	return 0
}

func (x *RoomBillDetail) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *RoomBillDetail) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type RoomBillDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  *common.Error     `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info      `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	Pager  *common.Pager     `protobuf:"bytes,3,opt,name=pager,proto3" json:"pager,omitempty"`
	Entity *RoomBillDetail   `protobuf:"bytes,4,opt,name=entity,proto3" json:"entity,omitempty"`
	Items  []*RoomBillDetail `protobuf:"bytes,5,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *RoomBillDetailResponse) Reset() {
	*x = RoomBillDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roomBillDetailService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomBillDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomBillDetailResponse) ProtoMessage() {}

func (x *RoomBillDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_roomBillDetailService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomBillDetailResponse.ProtoReflect.Descriptor instead.
func (*RoomBillDetailResponse) Descriptor() ([]byte, []int) {
	return file_roomBillDetailService_proto_rawDescGZIP(), []int{2}
}

func (x *RoomBillDetailResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *RoomBillDetailResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *RoomBillDetailResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *RoomBillDetailResponse) GetEntity() *RoomBillDetail {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *RoomBillDetailResponse) GetItems() []*RoomBillDetail {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_roomBillDetailService_proto protoreflect.FileDescriptor

var file_roomBillDetailService_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x72, 0x6f, 0x6f, 0x6d, 0x42, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x01, 0x0a, 0x13, 0x52,
	0x6f, 0x6f, 0x6d, 0x42, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x57, 0x68, 0x65,
	0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x62, 0x69,
	0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x6f, 0x6f,
	0x6d, 0x42, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x20, 0x0a, 0x0c, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x42, 0x6f, 0x6f, 0x6b, 0x49,
	0x64, 0x22, 0xaf, 0x02, 0x0a, 0x0e, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x69, 0x6c, 0x6c, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x62, 0x69, 0x6c,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x6f, 0x6f, 0x6d,
	0x42, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x69, 0x6c, 0x6c, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x62, 0x69, 0x6c, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x61, 0x79, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x15,
	0x0a, 0x06, 0x6f, 0x70, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x6f, 0x70, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xe6, 0x01, 0x0a, 0x16, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x69, 0x6c, 0x6c,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x69, 0x6c, 0x6c, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x69, 0x6c, 0x6c, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x62, 0x0a, 0x15,
	0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x69, 0x6c, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x42,
	0x69, 0x6c, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x20,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x69,
	0x6c, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_roomBillDetailService_proto_rawDescOnce sync.Once
	file_roomBillDetailService_proto_rawDescData = file_roomBillDetailService_proto_rawDesc
)

func file_roomBillDetailService_proto_rawDescGZIP() []byte {
	file_roomBillDetailService_proto_rawDescOnce.Do(func() {
		file_roomBillDetailService_proto_rawDescData = protoimpl.X.CompressGZIP(file_roomBillDetailService_proto_rawDescData)
	})
	return file_roomBillDetailService_proto_rawDescData
}

var file_roomBillDetailService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_roomBillDetailService_proto_goTypes = []interface{}{
	(*RoomBillDetailWhere)(nil),    // 0: services.RoomBillDetailWhere
	(*RoomBillDetail)(nil),         // 1: services.RoomBillDetail
	(*RoomBillDetailResponse)(nil), // 2: services.RoomBillDetailResponse
	(*common.Error)(nil),           // 3: common.Error
	(*common.Info)(nil),            // 4: common.Info
	(*common.Pager)(nil),           // 5: common.Pager
}
var file_roomBillDetailService_proto_depIdxs = []int32{
	3, // 0: services.RoomBillDetailResponse.error:type_name -> common.Error
	4, // 1: services.RoomBillDetailResponse.info:type_name -> common.Info
	5, // 2: services.RoomBillDetailResponse.pager:type_name -> common.Pager
	1, // 3: services.RoomBillDetailResponse.entity:type_name -> services.RoomBillDetail
	1, // 4: services.RoomBillDetailResponse.items:type_name -> services.RoomBillDetail
	0, // 5: services.RoomBillDetailService.Search:input_type -> services.RoomBillDetailWhere
	2, // 6: services.RoomBillDetailService.Search:output_type -> services.RoomBillDetailResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_roomBillDetailService_proto_init() }
func file_roomBillDetailService_proto_init() {
	if File_roomBillDetailService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_roomBillDetailService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomBillDetailWhere); i {
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
		file_roomBillDetailService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomBillDetail); i {
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
		file_roomBillDetailService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomBillDetailResponse); i {
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
			RawDescriptor: file_roomBillDetailService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_roomBillDetailService_proto_goTypes,
		DependencyIndexes: file_roomBillDetailService_proto_depIdxs,
		MessageInfos:      file_roomBillDetailService_proto_msgTypes,
	}.Build()
	File_roomBillDetailService_proto = out.File
	file_roomBillDetailService_proto_rawDesc = nil
	file_roomBillDetailService_proto_goTypes = nil
	file_roomBillDetailService_proto_depIdxs = nil
}
