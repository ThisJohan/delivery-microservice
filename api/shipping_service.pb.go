// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: shipping_service.proto

package api

import (
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

type ShipmentStatus int32

const (
	Shipment_QUEUED    ShipmentStatus = 0
	Shipment_PENDING   ShipmentStatus = 1
	Shipment_ASSIGNED  ShipmentStatus = 2
	Shipment_DELIVERED ShipmentStatus = 3
	Shipment_CANCELED  ShipmentStatus = 4
)

// Enum value maps for ShipmentStatus.
var (
	ShipmentStatus_name = map[int32]string{
		0: "QUEUED",
		1: "PENDING",
		2: "ASSIGNED",
		3: "DELIVERED",
		4: "CANCELED",
	}
	ShipmentStatus_value = map[string]int32{
		"QUEUED":    0,
		"PENDING":   1,
		"ASSIGNED":  2,
		"DELIVERED": 3,
		"CANCELED":  4,
	}
)

func (x ShipmentStatus) Enum() *ShipmentStatus {
	p := new(ShipmentStatus)
	*p = x
	return p
}

func (x ShipmentStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ShipmentStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_shipping_service_proto_enumTypes[0].Descriptor()
}

func (ShipmentStatus) Type() protoreflect.EnumType {
	return &file_shipping_service_proto_enumTypes[0]
}

func (x ShipmentStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ShipmentStatus.Descriptor instead.
func (ShipmentStatus) EnumDescriptor() ([]byte, []int) {
	return file_shipping_service_proto_rawDescGZIP(), []int{2, 0}
}

type CreateShipmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID     string `protobuf:"bytes,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	UserAddress string `protobuf:"bytes,2,opt,name=UserAddress,proto3" json:"UserAddress,omitempty"`
	Origin      string `protobuf:"bytes,3,opt,name=Origin,proto3" json:"Origin,omitempty"`
	Destination string `protobuf:"bytes,4,opt,name=Destination,proto3" json:"Destination,omitempty"`
	TimeSlot    int64  `protobuf:"varint,5,opt,name=TimeSlot,proto3" json:"TimeSlot,omitempty"`
}

func (x *CreateShipmentRequest) Reset() {
	*x = CreateShipmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipping_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShipmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShipmentRequest) ProtoMessage() {}

func (x *CreateShipmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shipping_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShipmentRequest.ProtoReflect.Descriptor instead.
func (*CreateShipmentRequest) Descriptor() ([]byte, []int) {
	return file_shipping_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateShipmentRequest) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *CreateShipmentRequest) GetUserAddress() string {
	if x != nil {
		return x.UserAddress
	}
	return ""
}

func (x *CreateShipmentRequest) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *CreateShipmentRequest) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *CreateShipmentRequest) GetTimeSlot() int64 {
	if x != nil {
		return x.TimeSlot
	}
	return 0
}

type ShipmentStatusChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     uint32         `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Status ShipmentStatus `protobuf:"varint,2,opt,name=Status,proto3,enum=api.ShipmentStatus" json:"Status,omitempty"`
}

func (x *ShipmentStatusChange) Reset() {
	*x = ShipmentStatusChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipping_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipmentStatusChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipmentStatusChange) ProtoMessage() {}

func (x *ShipmentStatusChange) ProtoReflect() protoreflect.Message {
	mi := &file_shipping_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipmentStatusChange.ProtoReflect.Descriptor instead.
func (*ShipmentStatusChange) Descriptor() ([]byte, []int) {
	return file_shipping_service_proto_rawDescGZIP(), []int{1}
}

func (x *ShipmentStatusChange) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ShipmentStatusChange) GetStatus() ShipmentStatus {
	if x != nil {
		return x.Status
	}
	return Shipment_QUEUED
}

type Shipment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          uint32         `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	OrderID     string         `protobuf:"bytes,2,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	UserAddress string         `protobuf:"bytes,3,opt,name=UserAddress,proto3" json:"UserAddress,omitempty"`
	Origin      string         `protobuf:"bytes,4,opt,name=Origin,proto3" json:"Origin,omitempty"`
	Destination string         `protobuf:"bytes,5,opt,name=Destination,proto3" json:"Destination,omitempty"`
	TimeSlot    int64          `protobuf:"varint,6,opt,name=TimeSlot,proto3" json:"TimeSlot,omitempty"`
	Status      ShipmentStatus `protobuf:"varint,7,opt,name=Status,proto3,enum=api.ShipmentStatus" json:"Status,omitempty"`
}

func (x *Shipment) Reset() {
	*x = Shipment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipping_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shipment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shipment) ProtoMessage() {}

func (x *Shipment) ProtoReflect() protoreflect.Message {
	mi := &file_shipping_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shipment.ProtoReflect.Descriptor instead.
func (*Shipment) Descriptor() ([]byte, []int) {
	return file_shipping_service_proto_rawDescGZIP(), []int{2}
}

func (x *Shipment) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Shipment) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *Shipment) GetUserAddress() string {
	if x != nil {
		return x.UserAddress
	}
	return ""
}

func (x *Shipment) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *Shipment) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (x *Shipment) GetTimeSlot() int64 {
	if x != nil {
		return x.TimeSlot
	}
	return 0
}

func (x *Shipment) GetStatus() ShipmentStatus {
	if x != nil {
		return x.Status
	}
	return Shipment_QUEUED
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipping_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_shipping_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_shipping_service_proto_rawDescGZIP(), []int{3}
}

var File_shipping_service_proto protoreflect.FileDescriptor

var file_shipping_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0xa9, 0x01,
	0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x20, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x22, 0x54, 0x0a, 0x14, 0x53, 0x68, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x2c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0xa8, 0x02, 0x0a, 0x08, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x2c,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4c, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x51, 0x55, 0x45, 0x55, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x41, 0x53, 0x53, 0x49, 0x47, 0x4e, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a,
	0x09, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08,
	0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f,
	0x69, 0x64, 0x32, 0x80, 0x01, 0x0a, 0x0f, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x36, 0x0a,
	0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x19, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x1a, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x56,
	0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x68, 0x69, 0x73, 0x4a, 0x6f, 0x68, 0x61, 0x6e, 0x2f, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shipping_service_proto_rawDescOnce sync.Once
	file_shipping_service_proto_rawDescData = file_shipping_service_proto_rawDesc
)

func file_shipping_service_proto_rawDescGZIP() []byte {
	file_shipping_service_proto_rawDescOnce.Do(func() {
		file_shipping_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_shipping_service_proto_rawDescData)
	})
	return file_shipping_service_proto_rawDescData
}

var file_shipping_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_shipping_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_shipping_service_proto_goTypes = []any{
	(ShipmentStatus)(0),           // 0: api.Shipment.status
	(*CreateShipmentRequest)(nil), // 1: api.CreateShipmentRequest
	(*ShipmentStatusChange)(nil),  // 2: api.ShipmentStatusChange
	(*Shipment)(nil),              // 3: api.Shipment
	(*Void)(nil),                  // 4: api.Void
}
var file_shipping_service_proto_depIdxs = []int32{
	0, // 0: api.ShipmentStatusChange.Status:type_name -> api.Shipment.status
	0, // 1: api.Shipment.Status:type_name -> api.Shipment.status
	1, // 2: api.ShippingService.Create:input_type -> api.CreateShipmentRequest
	2, // 3: api.ShippingService.StatusChange:input_type -> api.ShipmentStatusChange
	3, // 4: api.ShippingService.Create:output_type -> api.Shipment
	4, // 5: api.ShippingService.StatusChange:output_type -> api.Void
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_shipping_service_proto_init() }
func file_shipping_service_proto_init() {
	if File_shipping_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shipping_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateShipmentRequest); i {
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
		file_shipping_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ShipmentStatusChange); i {
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
		file_shipping_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Shipment); i {
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
		file_shipping_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Void); i {
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
			RawDescriptor: file_shipping_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shipping_service_proto_goTypes,
		DependencyIndexes: file_shipping_service_proto_depIdxs,
		EnumInfos:         file_shipping_service_proto_enumTypes,
		MessageInfos:      file_shipping_service_proto_msgTypes,
	}.Build()
	File_shipping_service_proto = out.File
	file_shipping_service_proto_rawDesc = nil
	file_shipping_service_proto_goTypes = nil
	file_shipping_service_proto_depIdxs = nil
}
