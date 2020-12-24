// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: trackService.proto

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

type Track struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	ShipperCode  string `protobuf:"bytes,2,opt,name=shipper_code,json=shipperCode,proto3" json:"shipper_code"`
	OrderCode    string `protobuf:"bytes,3,opt,name=order_code,json=orderCode,proto3" json:"order_code"`
	LogisticCode string `protobuf:"bytes,4,opt,name=logistic_code,json=logisticCode,proto3" json:"logistic_code"`
	Success      bool   `protobuf:"varint,5,opt,name=success,proto3" json:"success"`
	Reason       string `protobuf:"bytes,6,opt,name=reason,proto3" json:"reason"`
	State        string `protobuf:"bytes,7,opt,name=state,proto3" json:"state"`
	StateEx      string `protobuf:"bytes,8,opt,name=stateEx,proto3" json:"stateEx"`
	Location     string `protobuf:"bytes,9,opt,name=location,proto3" json:"location"`
	// @inject_tag: gorm:"foreignKey:TrackId"
	Details      []*TrackDetail `protobuf:"bytes,10,rep,name=details,proto3" json:"details" gorm:"foreignKey:TrackId"`
	CustomerName string         `protobuf:"bytes,11,opt,name=customer_name,json=customerName,proto3" json:"customer_name"`
}

func (x *Track) Reset() {
	*x = Track{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trackService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Track) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Track) ProtoMessage() {}

func (x *Track) ProtoReflect() protoreflect.Message {
	mi := &file_trackService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Track.ProtoReflect.Descriptor instead.
func (*Track) Descriptor() ([]byte, []int) {
	return file_trackService_proto_rawDescGZIP(), []int{0}
}

func (x *Track) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Track) GetShipperCode() string {
	if x != nil {
		return x.ShipperCode
	}
	return ""
}

func (x *Track) GetOrderCode() string {
	if x != nil {
		return x.OrderCode
	}
	return ""
}

func (x *Track) GetLogisticCode() string {
	if x != nil {
		return x.LogisticCode
	}
	return ""
}

func (x *Track) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Track) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Track) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Track) GetStateEx() string {
	if x != nil {
		return x.StateEx
	}
	return ""
}

func (x *Track) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Track) GetDetails() []*TrackDetail {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *Track) GetCustomerName() string {
	if x != nil {
		return x.CustomerName
	}
	return ""
}

type TrackDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	TrackId       int32  `protobuf:"varint,2,opt,name=track_id,json=trackId,proto3" json:"track_id"`
	AcceptTime    string `protobuf:"bytes,3,opt,name=accept_time,json=acceptTime,proto3" json:"accept_time"`
	AcceptStation string `protobuf:"bytes,4,opt,name=accept_station,json=acceptStation,proto3" json:"accept_station"`
	Location      string `protobuf:"bytes,5,opt,name=location,proto3" json:"location"`
	Action        string `protobuf:"bytes,6,opt,name=action,proto3" json:"action"`
	Remark        string `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark"`
}

func (x *TrackDetail) Reset() {
	*x = TrackDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trackService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackDetail) ProtoMessage() {}

func (x *TrackDetail) ProtoReflect() protoreflect.Message {
	mi := &file_trackService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackDetail.ProtoReflect.Descriptor instead.
func (*TrackDetail) Descriptor() ([]byte, []int) {
	return file_trackService_proto_rawDescGZIP(), []int{1}
}

func (x *TrackDetail) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TrackDetail) GetTrackId() int32 {
	if x != nil {
		return x.TrackId
	}
	return 0
}

func (x *TrackDetail) GetAcceptTime() string {
	if x != nil {
		return x.AcceptTime
	}
	return ""
}

func (x *TrackDetail) GetAcceptStation() string {
	if x != nil {
		return x.AcceptStation
	}
	return ""
}

func (x *TrackDetail) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *TrackDetail) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *TrackDetail) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type TrackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Track        `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Track      `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Error  *common.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error"`
	Info   *common.Info  `protobuf:"bytes,5,opt,name=info,proto3" json:"info"`
}

func (x *TrackResponse) Reset() {
	*x = TrackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trackService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackResponse) ProtoMessage() {}

func (x *TrackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trackService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackResponse.ProtoReflect.Descriptor instead.
func (*TrackResponse) Descriptor() ([]byte, []int) {
	return file_trackService_proto_rawDescGZIP(), []int{2}
}

func (x *TrackResponse) GetEntity() *Track {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *TrackResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *TrackResponse) GetItems() []*Track {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *TrackResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *TrackResponse) GetInfo() *common.Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_trackService_proto protoreflect.FileDescriptor

var file_trackService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd2, 0x02, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x45, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x45, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2f, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xcc, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0xcb, 0x01, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x32, 0x43, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x0f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x1a, 0x17, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trackService_proto_rawDescOnce sync.Once
	file_trackService_proto_rawDescData = file_trackService_proto_rawDesc
)

func file_trackService_proto_rawDescGZIP() []byte {
	file_trackService_proto_rawDescOnce.Do(func() {
		file_trackService_proto_rawDescData = protoimpl.X.CompressGZIP(file_trackService_proto_rawDescData)
	})
	return file_trackService_proto_rawDescData
}

var file_trackService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_trackService_proto_goTypes = []interface{}{
	(*Track)(nil),         // 0: services.Track
	(*TrackDetail)(nil),   // 1: services.TrackDetail
	(*TrackResponse)(nil), // 2: services.TrackResponse
	(*common.Pager)(nil),  // 3: common.Pager
	(*common.Error)(nil),  // 4: common.Error
	(*common.Info)(nil),   // 5: common.Info
}
var file_trackService_proto_depIdxs = []int32{
	1, // 0: services.Track.details:type_name -> services.TrackDetail
	0, // 1: services.TrackResponse.entity:type_name -> services.Track
	3, // 2: services.TrackResponse.pager:type_name -> common.Pager
	0, // 3: services.TrackResponse.items:type_name -> services.Track
	4, // 4: services.TrackResponse.error:type_name -> common.Error
	5, // 5: services.TrackResponse.info:type_name -> common.Info
	0, // 6: services.TrackService.Query:input_type -> services.Track
	2, // 7: services.TrackService.Query:output_type -> services.TrackResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_trackService_proto_init() }
func file_trackService_proto_init() {
	if File_trackService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trackService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Track); i {
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
		file_trackService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackDetail); i {
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
		file_trackService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackResponse); i {
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
			RawDescriptor: file_trackService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trackService_proto_goTypes,
		DependencyIndexes: file_trackService_proto_depIdxs,
		MessageInfos:      file_trackService_proto_msgTypes,
	}.Build()
	File_trackService_proto = out.File
	file_trackService_proto_rawDesc = nil
	file_trackService_proto_goTypes = nil
	file_trackService_proto_depIdxs = nil
}
