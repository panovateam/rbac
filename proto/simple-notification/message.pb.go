// Code generated by protoc-gen-go. DO NOT EDIT.
// source: simple-notification/message.proto

package message

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	calling "github.com/onskycloud/rbac/proto/calling"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	Thing          *calling.Thing `protobuf:"bytes,1,opt,name=thing,proto3" json:"thing,omitempty"`
	Gateway        *calling.Thing `protobuf:"bytes,2,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Template       *calling.Thing `protobuf:"bytes,3,opt,name=template,proto3" json:"template,omitempty"`
	Zone           *calling.Thing `protobuf:"bytes,4,opt,name=zone,proto3" json:"zone,omitempty"`
	CustomerNumber string         `protobuf:"bytes,5,opt,name=customerNumber,proto3" json:"customerNumber,omitempty"`
	MacAddress     string         `protobuf:"bytes,8,opt,name=macAddress,proto3" json:"macAddress,omitempty"`
	// notification type are below
	// DoorSensor
	// MotionSensor
	// SecurityBreach
	// SafetyBreachCO
	// SafetyBreachSOS
	// SafetyBreachSmoke
	// SecurityAlarm
	// SafetyAlarm
	Type                 int32    `protobuf:"varint,6,opt,name=type,proto3" json:"type,omitempty"`
	Value                string   `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_94158a4fda103638, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetThing() *calling.Thing {
	if m != nil {
		return m.Thing
	}
	return nil
}

func (m *Request) GetGateway() *calling.Thing {
	if m != nil {
		return m.Gateway
	}
	return nil
}

func (m *Request) GetTemplate() *calling.Thing {
	if m != nil {
		return m.Template
	}
	return nil
}

func (m *Request) GetZone() *calling.Thing {
	if m != nil {
		return m.Zone
	}
	return nil
}

func (m *Request) GetCustomerNumber() string {
	if m != nil {
		return m.CustomerNumber
	}
	return ""
}

func (m *Request) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *Request) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Request) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GetNotificationRequest struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Serial               string   `protobuf:"bytes,2,opt,name=serial,proto3" json:"serial,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNotificationRequest) Reset()         { *m = GetNotificationRequest{} }
func (m *GetNotificationRequest) String() string { return proto.CompactTextString(m) }
func (*GetNotificationRequest) ProtoMessage()    {}
func (*GetNotificationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94158a4fda103638, []int{1}
}

func (m *GetNotificationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNotificationRequest.Unmarshal(m, b)
}
func (m *GetNotificationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNotificationRequest.Marshal(b, m, deterministic)
}
func (m *GetNotificationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNotificationRequest.Merge(m, src)
}
func (m *GetNotificationRequest) XXX_Size() int {
	return xxx_messageInfo_GetNotificationRequest.Size(m)
}
func (m *GetNotificationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNotificationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNotificationRequest proto.InternalMessageInfo

func (m *GetNotificationRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GetNotificationRequest) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

type GetNotificationResponse struct {
	Success              bool                  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrorMessage         string                `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	Data                 []*NotificationResult `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetNotificationResponse) Reset()         { *m = GetNotificationResponse{} }
func (m *GetNotificationResponse) String() string { return proto.CompactTextString(m) }
func (*GetNotificationResponse) ProtoMessage()    {}
func (*GetNotificationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94158a4fda103638, []int{2}
}

func (m *GetNotificationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNotificationResponse.Unmarshal(m, b)
}
func (m *GetNotificationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNotificationResponse.Marshal(b, m, deterministic)
}
func (m *GetNotificationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNotificationResponse.Merge(m, src)
}
func (m *GetNotificationResponse) XXX_Size() int {
	return xxx_messageInfo_GetNotificationResponse.Size(m)
}
func (m *GetNotificationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNotificationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNotificationResponse proto.InternalMessageInfo

func (m *GetNotificationResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *GetNotificationResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *GetNotificationResponse) GetData() []*NotificationResult {
	if m != nil {
		return m.Data
	}
	return nil
}

type NotificationResult struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ThingName            string   `protobuf:"bytes,2,opt,name=thingName,proto3" json:"thingName,omitempty"`
	ThingSerial          string   `protobuf:"bytes,3,opt,name=thingSerial,proto3" json:"thingSerial,omitempty"`
	ThingDisplayName     string   `protobuf:"bytes,4,opt,name=thingDisplayName,proto3" json:"thingDisplayName,omitempty"`
	GatewayName          string   `protobuf:"bytes,5,opt,name=gatewayName,proto3" json:"gatewayName,omitempty"`
	GatewayDisplayName   string   `protobuf:"bytes,6,opt,name=gatewayDisplayName,proto3" json:"gatewayDisplayName,omitempty"`
	GatewayMacAddress    string   `protobuf:"bytes,7,opt,name=gatewayMacAddress,proto3" json:"gatewayMacAddress,omitempty"`
	ZoneName             string   `protobuf:"bytes,8,opt,name=zoneName,proto3" json:"zoneName,omitempty"`
	ZoneDisplayName      string   `protobuf:"bytes,9,opt,name=zoneDisplayName,proto3" json:"zoneDisplayName,omitempty"`
	Status               int32    `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`
	Type                 int32    `protobuf:"varint,11,opt,name=type,proto3" json:"type,omitempty"`
	Template             string   `protobuf:"bytes,12,opt,name=template,proto3" json:"template,omitempty"`
	State                int32    `protobuf:"varint,13,opt,name=state,proto3" json:"state,omitempty"`
	Description          string   `protobuf:"bytes,14,opt,name=Description,proto3" json:"Description,omitempty"`
	DescriptionVN        string   `protobuf:"bytes,15,opt,name=descriptionVN,proto3" json:"descriptionVN,omitempty"`
	AlertType            int32    `protobuf:"varint,16,opt,name=AlertType,proto3" json:"AlertType,omitempty"`
	DateTime             string   `protobuf:"bytes,17,opt,name=DateTime,proto3" json:"DateTime,omitempty"`
	Acknowledged         int32    `protobuf:"varint,18,opt,name=Acknowledged,proto3" json:"Acknowledged,omitempty"`
	AlertStatus          int32    `protobuf:"varint,19,opt,name=AlertStatus,proto3" json:"AlertStatus,omitempty"`
	DeviceId             int32    `protobuf:"varint,20,opt,name=DeviceId,proto3" json:"DeviceId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotificationResult) Reset()         { *m = NotificationResult{} }
func (m *NotificationResult) String() string { return proto.CompactTextString(m) }
func (*NotificationResult) ProtoMessage()    {}
func (*NotificationResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_94158a4fda103638, []int{3}
}

func (m *NotificationResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationResult.Unmarshal(m, b)
}
func (m *NotificationResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationResult.Marshal(b, m, deterministic)
}
func (m *NotificationResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationResult.Merge(m, src)
}
func (m *NotificationResult) XXX_Size() int {
	return xxx_messageInfo_NotificationResult.Size(m)
}
func (m *NotificationResult) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationResult.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationResult proto.InternalMessageInfo

func (m *NotificationResult) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NotificationResult) GetThingName() string {
	if m != nil {
		return m.ThingName
	}
	return ""
}

func (m *NotificationResult) GetThingSerial() string {
	if m != nil {
		return m.ThingSerial
	}
	return ""
}

func (m *NotificationResult) GetThingDisplayName() string {
	if m != nil {
		return m.ThingDisplayName
	}
	return ""
}

func (m *NotificationResult) GetGatewayName() string {
	if m != nil {
		return m.GatewayName
	}
	return ""
}

func (m *NotificationResult) GetGatewayDisplayName() string {
	if m != nil {
		return m.GatewayDisplayName
	}
	return ""
}

func (m *NotificationResult) GetGatewayMacAddress() string {
	if m != nil {
		return m.GatewayMacAddress
	}
	return ""
}

func (m *NotificationResult) GetZoneName() string {
	if m != nil {
		return m.ZoneName
	}
	return ""
}

func (m *NotificationResult) GetZoneDisplayName() string {
	if m != nil {
		return m.ZoneDisplayName
	}
	return ""
}

func (m *NotificationResult) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *NotificationResult) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *NotificationResult) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *NotificationResult) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *NotificationResult) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *NotificationResult) GetDescriptionVN() string {
	if m != nil {
		return m.DescriptionVN
	}
	return ""
}

func (m *NotificationResult) GetAlertType() int32 {
	if m != nil {
		return m.AlertType
	}
	return 0
}

func (m *NotificationResult) GetDateTime() string {
	if m != nil {
		return m.DateTime
	}
	return ""
}

func (m *NotificationResult) GetAcknowledged() int32 {
	if m != nil {
		return m.Acknowledged
	}
	return 0
}

func (m *NotificationResult) GetAlertStatus() int32 {
	if m != nil {
		return m.AlertStatus
	}
	return 0
}

func (m *NotificationResult) GetDeviceId() int32 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*GetNotificationRequest)(nil), "GetNotificationRequest")
	proto.RegisterType((*GetNotificationResponse)(nil), "GetNotificationResponse")
	proto.RegisterType((*NotificationResult)(nil), "NotificationResult")
}

func init() { proto.RegisterFile("simple-notification/message.proto", fileDescriptor_94158a4fda103638) }

var fileDescriptor_94158a4fda103638 = []byte{
	// 629 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x54, 0xcd, 0x4f, 0xdb, 0x4e,
	0x10, 0x25, 0x90, 0xcf, 0xe1, 0x23, 0xb0, 0x20, 0x58, 0x45, 0x3f, 0xfd, 0x4a, 0xad, 0xaa, 0x45,
	0x55, 0xeb, 0x48, 0x54, 0xea, 0x1d, 0x15, 0x51, 0xf5, 0x40, 0x0e, 0x06, 0xf5, 0xbe, 0x59, 0x4f,
	0x83, 0x85, 0x3f, 0x52, 0xef, 0x3a, 0x28, 0x9c, 0xfa, 0x67, 0xf4, 0x7f, 0xed, 0xa5, 0xbb, 0xb3,
	0x4e, 0x62, 0x27, 0xf4, 0xe6, 0xf7, 0xe6, 0xed, 0xcc, 0xee, 0xcc, 0x1b, 0xc3, 0x6b, 0x15, 0x25,
	0xd3, 0x18, 0x3f, 0xa6, 0x99, 0x8e, 0x7e, 0x44, 0x52, 0xe8, 0x28, 0x4b, 0x87, 0x09, 0x2a, 0x25,
	0x26, 0xe8, 0x4f, 0xf3, 0x4c, 0x67, 0x83, 0xcf, 0x93, 0x48, 0x3f, 0x14, 0x63, 0x5f, 0x66, 0xc9,
	0x30, 0x4b, 0xd5, 0xe3, 0x5c, 0xc6, 0x59, 0x11, 0x0e, 0xf3, 0xb1, 0x90, 0x43, 0x52, 0x0c, 0xa5,
	0x88, 0xe3, 0x28, 0x9d, 0xd4, 0xcf, 0x79, 0x7f, 0x1a, 0xd0, 0x09, 0xf0, 0x67, 0x81, 0x4a, 0xb3,
	0xff, 0xa0, 0xa5, 0x1f, 0x8c, 0x84, 0x37, 0xce, 0x1b, 0x17, 0xbb, 0x97, 0x6d, 0xff, 0xde, 0xa2,
	0xc0, 0x91, 0xec, 0x1c, 0x3a, 0x13, 0xa1, 0xf1, 0x49, 0xcc, 0xf9, 0x76, 0x2d, 0xbe, 0xa0, 0x99,
	0x07, 0x5d, 0x8d, 0xe6, 0xa2, 0x06, 0xf2, 0x9d, 0x9a, 0x64, 0xc9, 0xb3, 0x01, 0x34, 0x9f, 0xb3,
	0x14, 0x79, 0xb3, 0x16, 0x27, 0x8e, 0xbd, 0x85, 0x03, 0x59, 0x28, 0x9d, 0x25, 0x98, 0x8f, 0x8a,
	0x64, 0x8c, 0x39, 0x6f, 0x19, 0x55, 0x2f, 0x58, 0x63, 0xd9, 0xff, 0x00, 0x89, 0x90, 0x57, 0x61,
	0x98, 0x9b, 0xa7, 0xf0, 0x2e, 0x69, 0x2a, 0x0c, 0x63, 0xd0, 0xd4, 0xf3, 0x29, 0xf2, 0xb6, 0x89,
	0xb4, 0x02, 0xfa, 0x66, 0x27, 0xd0, 0x9a, 0x89, 0xb8, 0x40, 0xde, 0x21, 0xb9, 0x03, 0xde, 0x0d,
	0x9c, 0x7e, 0x45, 0x3d, 0xaa, 0xb4, 0x75, 0xd1, 0x0b, 0xa3, 0x97, 0x59, 0x91, 0x6a, 0xea, 0x45,
	0x2b, 0x70, 0x80, 0x9d, 0x42, 0x5b, 0x61, 0x1e, 0x89, 0x98, 0x5a, 0xd0, 0x0b, 0x4a, 0xe4, 0xfd,
	0x6a, 0xc0, 0xd9, 0x46, 0x22, 0x35, 0x35, 0x83, 0x40, 0xc6, 0xa1, 0xa3, 0x0a, 0x29, 0xed, 0x55,
	0x6d, 0xae, 0x6e, 0xb0, 0x80, 0xa6, 0x5f, 0x7b, 0x98, 0xe7, 0x59, 0x7e, 0xeb, 0x26, 0x52, 0xe6,
	0xac, 0x71, 0xec, 0x1d, 0x34, 0x43, 0xa1, 0x85, 0xe9, 0xe7, 0x8e, 0xe9, 0xd7, 0xb1, 0xbf, 0x56,
	0xa2, 0x88, 0x75, 0x40, 0x02, 0xef, 0x77, 0x0b, 0xd8, 0x66, 0x90, 0x1d, 0xc0, 0x76, 0x14, 0x96,
	0x8f, 0x30, 0x5f, 0x66, 0xc6, 0x3d, 0x1a, 0xe7, 0x48, 0x24, 0x8b, 0x82, 0x2b, 0xc2, 0xcc, 0x78,
	0x97, 0xc0, 0x9d, 0x7b, 0xe4, 0x0e, 0xc5, 0xab, 0x14, 0x7b, 0x0f, 0x87, 0x04, 0xaf, 0x23, 0x65,
	0x06, 0x3a, 0xa7, 0x34, 0x4d, 0x92, 0x6d, 0xf0, 0x36, 0x5b, 0x69, 0x0d, 0x92, 0xb9, 0x61, 0x56,
	0x29, 0xe6, 0x03, 0x2b, 0x61, 0x35, 0x5f, 0x9b, 0x84, 0x2f, 0x44, 0xd8, 0x07, 0x38, 0x2a, 0xd9,
	0xdb, 0x95, 0x01, 0xdc, 0x44, 0x37, 0x03, 0xc6, 0x6b, 0x5d, 0xeb, 0x2b, 0xca, 0xe9, 0x5c, 0xb2,
	0xc4, 0xec, 0x02, 0xfa, 0xf6, 0xbb, 0x5a, 0xb6, 0x47, 0x92, 0x75, 0x9a, 0x66, 0xae, 0x85, 0x2e,
	0x14, 0x07, 0xea, 0x62, 0x89, 0x96, 0x2e, 0xdb, 0xad, 0xb8, 0x6c, 0x50, 0xd9, 0x80, 0x3d, 0x57,
	0x71, 0xe9, 0x7c, 0xe3, 0x28, 0x7b, 0x12, 0xf9, 0xbe, 0x73, 0x14, 0x01, 0xdb, 0xa3, 0x6b, 0x54,
	0x32, 0x8f, 0xa6, 0x76, 0x68, 0xfc, 0xc0, 0xf5, 0xa8, 0x42, 0xb1, 0x37, 0xb0, 0x1f, 0xae, 0xe0,
	0xf7, 0x11, 0xef, 0x93, 0xa6, 0x4e, 0xda, 0xb9, 0x5e, 0xc5, 0x98, 0xeb, 0x7b, 0x7b, 0xa5, 0x43,
	0xaa, 0xb0, 0x22, 0xec, 0xbd, 0xae, 0x4d, 0xb5, 0xfb, 0xc8, 0x3c, 0xf3, 0xc8, 0xdd, 0x6b, 0x81,
	0xad, 0x0b, 0xaf, 0xe4, 0x63, 0x9a, 0x3d, 0xc5, 0x18, 0x4e, 0x30, 0xe4, 0x8c, 0x0e, 0xd7, 0x38,
	0x7b, 0x4b, 0x4a, 0x76, 0xe7, 0x1a, 0x71, 0x4c, 0x92, 0x2a, 0x45, 0x15, 0x70, 0x16, 0x49, 0xfc,
	0x16, 0xf2, 0x13, 0x0a, 0x2f, 0xf1, 0xe5, 0x33, 0xf4, 0xab, 0xce, 0xbc, 0x9b, 0x49, 0xf6, 0x0a,
	0xda, 0x5f, 0x72, 0xb4, 0x0d, 0xe8, 0xfa, 0xe5, 0xca, 0x0d, 0x7a, 0xfe, 0x62, 0x67, 0xbc, 0x2d,
	0x76, 0x03, 0xfd, 0xb5, 0x85, 0x62, 0x67, 0xfe, 0xcb, 0xbb, 0x3a, 0xe0, 0xfe, 0x3f, 0x76, 0xcf,
	0xdb, 0x1a, 0xb7, 0xe9, 0x37, 0xf7, 0xe9, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x66, 0x8d,
	0xf1, 0x43, 0x05, 0x00, 0x00,
}
