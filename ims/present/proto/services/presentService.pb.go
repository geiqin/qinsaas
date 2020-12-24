// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: presentService.proto

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

type Present struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	ItemId        int64  `protobuf:"varint,3,opt,name=item_id,json=itemId,proto3" json:"item_id"`
	SkuId         int64  `protobuf:"varint,4,opt,name=sku_id,json=skuId,proto3" json:"sku_id"`
	FetchLimit    int32  `protobuf:"varint,5,opt,name=fetch_limit,json=fetchLimit,proto3" json:"fetch_limit"`
	LongEffective bool   `protobuf:"varint,6,opt,name=long_effective,json=longEffective,proto3" json:"long_effective"`
	StartAt       string `protobuf:"bytes,7,opt,name=start_at,json=startAt,proto3" json:"start_at"`
	EndAt         string `protobuf:"bytes,8,opt,name=end_at,json=endAt,proto3" json:"end_at"`
	IssuedNum     int32  `protobuf:"varint,9,opt,name=issued_num,json=issuedNum,proto3" json:"issued_num"`
	ExchangedNum  int32  `protobuf:"varint,10,opt,name=exchanged_num,json=exchangedNum,proto3" json:"exchanged_num"`
	Status        int32  `protobuf:"varint,11,opt,name=status,proto3" json:"status"`
	CreatedAt     string `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt     string `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	// @inject_tag: gorm:"-"
	Ids []int64 `protobuf:"varint,14,rep,packed,name=ids,proto3" json:"ids" gorm:"-"`
}

func (x *Present) Reset() {
	*x = Present{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presentService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Present) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Present) ProtoMessage() {}

func (x *Present) ProtoReflect() protoreflect.Message {
	mi := &file_presentService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Present.ProtoReflect.Descriptor instead.
func (*Present) Descriptor() ([]byte, []int) {
	return file_presentService_proto_rawDescGZIP(), []int{0}
}

func (x *Present) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Present) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Present) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *Present) GetSkuId() int64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *Present) GetFetchLimit() int32 {
	if x != nil {
		return x.FetchLimit
	}
	return 0
}

func (x *Present) GetLongEffective() bool {
	if x != nil {
		return x.LongEffective
	}
	return false
}

func (x *Present) GetStartAt() string {
	if x != nil {
		return x.StartAt
	}
	return ""
}

func (x *Present) GetEndAt() string {
	if x != nil {
		return x.EndAt
	}
	return ""
}

func (x *Present) GetIssuedNum() int32 {
	if x != nil {
		return x.IssuedNum
	}
	return 0
}

func (x *Present) GetExchangedNum() int32 {
	if x != nil {
		return x.ExchangedNum
	}
	return 0
}

func (x *Present) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Present) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Present) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Present) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

//
type PresentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Present      `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Present    `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *PresentResponse) Reset() {
	*x = PresentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presentService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PresentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PresentResponse) ProtoMessage() {}

func (x *PresentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_presentService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PresentResponse.ProtoReflect.Descriptor instead.
func (*PresentResponse) Descriptor() ([]byte, []int) {
	return file_presentService_proto_rawDescGZIP(), []int{1}
}

func (x *PresentResponse) GetEntity() *Present {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *PresentResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *PresentResponse) GetItems() []*Present {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *PresentResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *PresentResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_presentService_proto protoreflect.FileDescriptor

var file_presentService_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x83, 0x03, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06,
	0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x6b,
	0x75, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x65, 0x74, 0x63, 0x68, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x65, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x6c, 0x6f,
	0x6e, 0x67, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x23, 0x0a, 0x0d,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x4e, 0x75,
	0x6d, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x0e,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xd1, 0x01, 0x0a, 0x0f, 0x50, 0x72,
	0x65, 0x73, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74,
	0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x27, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xe7, 0x02,
	0x0a, 0x0e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x38, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x1a, 0x19, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74,
	0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_presentService_proto_rawDescOnce sync.Once
	file_presentService_proto_rawDescData = file_presentService_proto_rawDesc
)

func file_presentService_proto_rawDescGZIP() []byte {
	file_presentService_proto_rawDescOnce.Do(func() {
		file_presentService_proto_rawDescData = protoimpl.X.CompressGZIP(file_presentService_proto_rawDescData)
	})
	return file_presentService_proto_rawDescData
}

var file_presentService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_presentService_proto_goTypes = []interface{}{
	(*Present)(nil),          // 0: services.Present
	(*PresentResponse)(nil),  // 1: services.PresentResponse
	(*common.Pager)(nil),     // 2: common.Pager
	(*common.Error)(nil),     // 3: common.Error
	(*common.Info)(nil),      // 4: common.Info
	(*common.BaseWhere)(nil), // 5: common.BaseWhere
}
var file_presentService_proto_depIdxs = []int32{
	0,  // 0: services.PresentResponse.entity:type_name -> services.Present
	2,  // 1: services.PresentResponse.pager:type_name -> common.Pager
	0,  // 2: services.PresentResponse.items:type_name -> services.Present
	3,  // 3: services.PresentResponse.error:type_name -> common.Error
	4,  // 4: services.PresentResponse.info:type_name -> common.Info
	0,  // 5: services.PresentService.Create:input_type -> services.Present
	0,  // 6: services.PresentService.Update:input_type -> services.Present
	0,  // 7: services.PresentService.Delete:input_type -> services.Present
	0,  // 8: services.PresentService.Get:input_type -> services.Present
	0,  // 9: services.PresentService.List:input_type -> services.Present
	5,  // 10: services.PresentService.Search:input_type -> common.BaseWhere
	1,  // 11: services.PresentService.Create:output_type -> services.PresentResponse
	1,  // 12: services.PresentService.Update:output_type -> services.PresentResponse
	1,  // 13: services.PresentService.Delete:output_type -> services.PresentResponse
	1,  // 14: services.PresentService.Get:output_type -> services.PresentResponse
	1,  // 15: services.PresentService.List:output_type -> services.PresentResponse
	1,  // 16: services.PresentService.Search:output_type -> services.PresentResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_presentService_proto_init() }
func file_presentService_proto_init() {
	if File_presentService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_presentService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Present); i {
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
		file_presentService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PresentResponse); i {
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
			RawDescriptor: file_presentService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_presentService_proto_goTypes,
		DependencyIndexes: file_presentService_proto_depIdxs,
		MessageInfos:      file_presentService_proto_msgTypes,
	}.Build()
	File_presentService_proto = out.File
	file_presentService_proto_rawDesc = nil
	file_presentService_proto_goTypes = nil
	file_presentService_proto_depIdxs = nil
}
