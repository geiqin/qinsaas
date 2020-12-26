// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: cashAccountService.proto

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

type CashAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	DistributorId   int64        `protobuf:"varint,2,opt,name=distributor_id,json=distributorId,proto3" json:"distributor_id"`
	Type            string       `protobuf:"bytes,3,opt,name=type,proto3" json:"type"`
	RealName        string       `protobuf:"bytes,4,opt,name=real_name,json=realName,proto3" json:"real_name"`
	PayeeName       string       `protobuf:"bytes,5,opt,name=payee_name,json=payeeName,proto3" json:"payee_name"`
	PayeeAccount    string       `protobuf:"bytes,6,opt,name=payee_account,json=payeeAccount,proto3" json:"payee_account"`
	PayeeBank       string       `protobuf:"bytes,7,opt,name=payee_bank,json=payeeBank,proto3" json:"payee_bank"`
	PlatformAccount string       `protobuf:"bytes,8,opt,name=platform_account,json=platformAccount,proto3" json:"platform_account"`
	Status          string       `protobuf:"bytes,9,opt,name=status,proto3" json:"status"`
	CreatedAt       string       `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt       string       `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Distributor     *Distributor `protobuf:"bytes,12,opt,name=distributor,proto3" json:"distributor"`
}

func (x *CashAccount) Reset() {
	*x = CashAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cashAccountService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CashAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CashAccount) ProtoMessage() {}

func (x *CashAccount) ProtoReflect() protoreflect.Message {
	mi := &file_cashAccountService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CashAccount.ProtoReflect.Descriptor instead.
func (*CashAccount) Descriptor() ([]byte, []int) {
	return file_cashAccountService_proto_rawDescGZIP(), []int{0}
}

func (x *CashAccount) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CashAccount) GetDistributorId() int64 {
	if x != nil {
		return x.DistributorId
	}
	return 0
}

func (x *CashAccount) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CashAccount) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *CashAccount) GetPayeeName() string {
	if x != nil {
		return x.PayeeName
	}
	return ""
}

func (x *CashAccount) GetPayeeAccount() string {
	if x != nil {
		return x.PayeeAccount
	}
	return ""
}

func (x *CashAccount) GetPayeeBank() string {
	if x != nil {
		return x.PayeeBank
	}
	return ""
}

func (x *CashAccount) GetPlatformAccount() string {
	if x != nil {
		return x.PlatformAccount
	}
	return ""
}

func (x *CashAccount) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CashAccount) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CashAccount) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *CashAccount) GetDistributor() *Distributor {
	if x != nil {
		return x.Distributor
	}
	return nil
}

type CashAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *CashAccount   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager  `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*CashAccount `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error  `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info   `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *CashAccountResponse) Reset() {
	*x = CashAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cashAccountService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CashAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CashAccountResponse) ProtoMessage() {}

func (x *CashAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cashAccountService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CashAccountResponse.ProtoReflect.Descriptor instead.
func (*CashAccountResponse) Descriptor() ([]byte, []int) {
	return file_cashAccountService_proto_rawDescGZIP(), []int{1}
}

func (x *CashAccountResponse) GetEntity() *CashAccount {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *CashAccountResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *CashAccountResponse) GetItems() []*CashAccount {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CashAccountResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *CashAccountResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_cashAccountService_proto protoreflect.FileDescriptor

var file_cashAccountService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x61, 0x73, 0x68, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x92, 0x03, 0x0a, 0x0b, 0x43, 0x61, 0x73, 0x68, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79,
	0x65, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x61, 0x79, 0x65, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x65,
	0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x70, 0x61, 0x79, 0x65, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x79, 0x65, 0x65, 0x5f, 0x62, 0x61, 0x6e, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x79, 0x65, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x29, 0x0a, 0x10,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x37, 0x0a,
	0x0b, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x22, 0xdd, 0x01, 0x0a, 0x13, 0x43, 0x61, 0x73, 0x68, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x73, 0x68, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x72, 0x12, 0x2b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x73,
	0x68, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0x95, 0x02, 0x0a, 0x12, 0x43, 0x61, 0x73, 0x68, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x61, 0x73, 0x68, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x1d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x73, 0x68, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x40, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x73, 0x68, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x73, 0x68,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x73, 0x68, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a,
	0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x73, 0x68, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3c, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x1d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x73, 0x68, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cashAccountService_proto_rawDescOnce sync.Once
	file_cashAccountService_proto_rawDescData = file_cashAccountService_proto_rawDesc
)

func file_cashAccountService_proto_rawDescGZIP() []byte {
	file_cashAccountService_proto_rawDescOnce.Do(func() {
		file_cashAccountService_proto_rawDescData = protoimpl.X.CompressGZIP(file_cashAccountService_proto_rawDescData)
	})
	return file_cashAccountService_proto_rawDescData
}

var file_cashAccountService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cashAccountService_proto_goTypes = []interface{}{
	(*CashAccount)(nil),         // 0: services.CashAccount
	(*CashAccountResponse)(nil), // 1: services.CashAccountResponse
	(*Distributor)(nil),         // 2: services.Distributor
	(*common.Pager)(nil),        // 3: common.Pager
	(*common.Error)(nil),        // 4: common.Error
	(*common.Info)(nil),         // 5: common.Info
	(*common.BaseWhere)(nil),    // 6: common.BaseWhere
}
var file_cashAccountService_proto_depIdxs = []int32{
	2,  // 0: services.CashAccount.distributor:type_name -> services.Distributor
	0,  // 1: services.CashAccountResponse.entity:type_name -> services.CashAccount
	3,  // 2: services.CashAccountResponse.pager:type_name -> common.Pager
	0,  // 3: services.CashAccountResponse.items:type_name -> services.CashAccount
	4,  // 4: services.CashAccountResponse.error:type_name -> common.Error
	5,  // 5: services.CashAccountResponse.info:type_name -> common.Info
	0,  // 6: services.CashAccountService.Create:input_type -> services.CashAccount
	0,  // 7: services.CashAccountService.Delete:input_type -> services.CashAccount
	0,  // 8: services.CashAccountService.Get:input_type -> services.CashAccount
	6,  // 9: services.CashAccountService.Search:input_type -> common.BaseWhere
	1,  // 10: services.CashAccountService.Create:output_type -> services.CashAccountResponse
	1,  // 11: services.CashAccountService.Delete:output_type -> services.CashAccountResponse
	1,  // 12: services.CashAccountService.Get:output_type -> services.CashAccountResponse
	1,  // 13: services.CashAccountService.Search:output_type -> services.CashAccountResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_cashAccountService_proto_init() }
func file_cashAccountService_proto_init() {
	if File_cashAccountService_proto != nil {
		return
	}
	file_distributorService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cashAccountService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CashAccount); i {
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
		file_cashAccountService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CashAccountResponse); i {
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
			RawDescriptor: file_cashAccountService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cashAccountService_proto_goTypes,
		DependencyIndexes: file_cashAccountService_proto_depIdxs,
		MessageInfos:      file_cashAccountService_proto_msgTypes,
	}.Build()
	File_cashAccountService_proto = out.File
	file_cashAccountService_proto_rawDesc = nil
	file_cashAccountService_proto_goTypes = nil
	file_cashAccountService_proto_depIdxs = nil
}
