// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: creditService.proto

package services

import (
	common "geiqin.saas.crm/app/proto/common"
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

type Credit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId    int64           `protobuf:"varint,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Status        bool            `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	UpperLimit    float32         `protobuf:"fixed32,4,opt,name=upper_limit,json=upperLimit,proto3" json:"upper_limit,omitempty"`
	Balance       float32         `protobuf:"fixed32,5,opt,name=balance,proto3" json:"balance,omitempty"`
	CreatedAt     string          `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string          `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Customer      *Customer       `protobuf:"bytes,8,opt,name=customer,proto3" json:"customer,omitempty"`
	Ids           []int64         `protobuf:"varint,9,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Name          string          `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName   string          `protobuf:"bytes,11,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Realname      string          `protobuf:"bytes,12,opt,name=realname,proto3" json:"realname,omitempty"`
	Mobile        string          `protobuf:"bytes,13,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Contacts      string          `protobuf:"bytes,14,opt,name=contacts,proto3" json:"contacts,omitempty"`
	Cid           int64           `protobuf:"varint,15,opt,name=cid,proto3" json:"cid,omitempty"` //客户表的客户id
	CreditRecords []*CreditRecord `protobuf:"bytes,16,rep,name=credit_records,json=creditRecords,proto3" json:"credit_records,omitempty"`
}

func (x *Credit) Reset() {
	*x = Credit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creditService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credit) ProtoMessage() {}

func (x *Credit) ProtoReflect() protoreflect.Message {
	mi := &file_creditService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credit.ProtoReflect.Descriptor instead.
func (*Credit) Descriptor() ([]byte, []int) {
	return file_creditService_proto_rawDescGZIP(), []int{0}
}

func (x *Credit) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Credit) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Credit) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *Credit) GetUpperLimit() float32 {
	if x != nil {
		return x.UpperLimit
	}
	return 0
}

func (x *Credit) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *Credit) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Credit) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Credit) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *Credit) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *Credit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Credit) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Credit) GetRealname() string {
	if x != nil {
		return x.Realname
	}
	return ""
}

func (x *Credit) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *Credit) GetContacts() string {
	if x != nil {
		return x.Contacts
	}
	return ""
}

func (x *Credit) GetCid() int64 {
	if x != nil {
		return x.Cid
	}
	return 0
}

func (x *Credit) GetCreditRecords() []*CreditRecord {
	if x != nil {
		return x.CreditRecords
	}
	return nil
}

//查询参数
type CreditWhere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int32 `protobuf:"varint,1,opt,name=paged,proto3" json:"paged,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	//以下为自定义参数
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Realname    string `protobuf:"bytes,5,opt,name=realname,proto3" json:"realname,omitempty"`
	Mobile      string `protobuf:"bytes,6,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Type        int32  `protobuf:"varint,7,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *CreditWhere) Reset() {
	*x = CreditWhere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creditService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreditWhere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditWhere) ProtoMessage() {}

func (x *CreditWhere) ProtoReflect() protoreflect.Message {
	mi := &file_creditService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditWhere.ProtoReflect.Descriptor instead.
func (*CreditWhere) Descriptor() ([]byte, []int) {
	return file_creditService_proto_rawDescGZIP(), []int{1}
}

func (x *CreditWhere) GetPaged() int32 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *CreditWhere) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *CreditWhere) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreditWhere) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *CreditWhere) GetRealname() string {
	if x != nil {
		return x.Realname
	}
	return ""
}

func (x *CreditWhere) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *CreditWhere) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

//
type CreditResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Credit       `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Items  []*Credit     `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *CreditResponse) Reset() {
	*x = CreditResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creditService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreditResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditResponse) ProtoMessage() {}

func (x *CreditResponse) ProtoReflect() protoreflect.Message {
	mi := &file_creditService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditResponse.ProtoReflect.Descriptor instead.
func (*CreditResponse) Descriptor() ([]byte, []int) {
	return file_creditService_proto_rawDescGZIP(), []int{2}
}

func (x *CreditResponse) GetEntity() *Credit {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *CreditResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *CreditResponse) GetItems() []*Credit {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CreditResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *CreditResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_creditService_proto protoreflect.FileDescriptor

var file_creditService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x03, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x70, 0x65,
	0x72, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x75,
	0x70, 0x70, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x2e, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x61, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x61, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69,
	0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x0e,
	0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x10,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0d, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0xbf, 0x01, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x57, 0x68, 0x65, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xce, 0x01,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x28, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12,
	0x26, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xb9,
	0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x33, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a,
	0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x57, 0x68, 0x65, 0x72, 0x65, 0x1a, 0x18,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_creditService_proto_rawDescOnce sync.Once
	file_creditService_proto_rawDescData = file_creditService_proto_rawDesc
)

func file_creditService_proto_rawDescGZIP() []byte {
	file_creditService_proto_rawDescOnce.Do(func() {
		file_creditService_proto_rawDescData = protoimpl.X.CompressGZIP(file_creditService_proto_rawDescData)
	})
	return file_creditService_proto_rawDescData
}

var file_creditService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_creditService_proto_goTypes = []interface{}{
	(*Credit)(nil),         // 0: services.Credit
	(*CreditWhere)(nil),    // 1: services.CreditWhere
	(*CreditResponse)(nil), // 2: services.CreditResponse
	(*Customer)(nil),       // 3: services.Customer
	(*CreditRecord)(nil),   // 4: services.CreditRecord
	(*common.Pager)(nil),   // 5: common.Pager
	(*common.Error)(nil),   // 6: common.Error
	(*common.Info)(nil),    // 7: common.Info
}
var file_creditService_proto_depIdxs = []int32{
	3,  // 0: services.Credit.customer:type_name -> services.Customer
	4,  // 1: services.Credit.credit_records:type_name -> services.CreditRecord
	0,  // 2: services.CreditResponse.entity:type_name -> services.Credit
	5,  // 3: services.CreditResponse.pager:type_name -> common.Pager
	0,  // 4: services.CreditResponse.items:type_name -> services.Credit
	6,  // 5: services.CreditResponse.error:type_name -> common.Error
	7,  // 6: services.CreditResponse.info:type_name -> common.Info
	0,  // 7: services.CreditService.Set:input_type -> services.Credit
	0,  // 8: services.CreditService.Detail:input_type -> services.Credit
	1,  // 9: services.CreditService.Search:input_type -> services.CreditWhere
	2,  // 10: services.CreditService.Set:output_type -> services.CreditResponse
	2,  // 11: services.CreditService.Detail:output_type -> services.CreditResponse
	2,  // 12: services.CreditService.Search:output_type -> services.CreditResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_creditService_proto_init() }
func file_creditService_proto_init() {
	if File_creditService_proto != nil {
		return
	}
	file_customerService_proto_init()
	file_creditRecordService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_creditService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credit); i {
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
		file_creditService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreditWhere); i {
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
		file_creditService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreditResponse); i {
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
			RawDescriptor: file_creditService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_creditService_proto_goTypes,
		DependencyIndexes: file_creditService_proto_depIdxs,
		MessageInfos:      file_creditService_proto_msgTypes,
	}.Build()
	File_creditService_proto = out.File
	file_creditService_proto_rawDesc = nil
	file_creditService_proto_goTypes = nil
	file_creditService_proto_depIdxs = nil
}
