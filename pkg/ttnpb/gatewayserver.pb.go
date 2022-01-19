// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/gatewayserver.proto

package ttnpb

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GatewayUp may contain zero or more uplink messages and/or a status message for the gateway.
type GatewayUp struct {
	// Uplink messages received by the gateway.
	UplinkMessages []*UplinkMessage `protobuf:"bytes,1,rep,name=uplink_messages,json=uplinkMessages,proto3" json:"uplink_messages,omitempty"`
	// Gateway status produced by the gateway.
	GatewayStatus *GatewayStatus `protobuf:"bytes,2,opt,name=gateway_status,json=gatewayStatus,proto3" json:"gateway_status,omitempty"`
	// A Tx acknowledgment or error.
	TxAcknowledgment     *TxAcknowledgment `protobuf:"bytes,3,opt,name=tx_acknowledgment,json=txAcknowledgment,proto3" json:"tx_acknowledgment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GatewayUp) Reset()         { *m = GatewayUp{} }
func (m *GatewayUp) String() string { return proto.CompactTextString(m) }
func (*GatewayUp) ProtoMessage()    {}
func (*GatewayUp) Descriptor() ([]byte, []int) {
	return fileDescriptor_62b07a36420f2d6d, []int{0}
}
func (m *GatewayUp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayUp.Unmarshal(m, b)
}
func (m *GatewayUp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayUp.Marshal(b, m, deterministic)
}
func (m *GatewayUp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayUp.Merge(m, src)
}
func (m *GatewayUp) XXX_Size() int {
	return xxx_messageInfo_GatewayUp.Size(m)
}
func (m *GatewayUp) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayUp.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayUp proto.InternalMessageInfo

func (m *GatewayUp) GetUplinkMessages() []*UplinkMessage {
	if m != nil {
		return m.UplinkMessages
	}
	return nil
}

func (m *GatewayUp) GetGatewayStatus() *GatewayStatus {
	if m != nil {
		return m.GatewayStatus
	}
	return nil
}

func (m *GatewayUp) GetTxAcknowledgment() *TxAcknowledgment {
	if m != nil {
		return m.TxAcknowledgment
	}
	return nil
}

// GatewayDown contains downlink messages for the gateway.
type GatewayDown struct {
	// DownlinkMessage for the gateway.
	DownlinkMessage      *DownlinkMessage `protobuf:"bytes,1,opt,name=downlink_message,json=downlinkMessage,proto3" json:"downlink_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GatewayDown) Reset()         { *m = GatewayDown{} }
func (m *GatewayDown) String() string { return proto.CompactTextString(m) }
func (*GatewayDown) ProtoMessage()    {}
func (*GatewayDown) Descriptor() ([]byte, []int) {
	return fileDescriptor_62b07a36420f2d6d, []int{1}
}
func (m *GatewayDown) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayDown.Unmarshal(m, b)
}
func (m *GatewayDown) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayDown.Marshal(b, m, deterministic)
}
func (m *GatewayDown) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayDown.Merge(m, src)
}
func (m *GatewayDown) XXX_Size() int {
	return xxx_messageInfo_GatewayDown.Size(m)
}
func (m *GatewayDown) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayDown.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayDown proto.InternalMessageInfo

func (m *GatewayDown) GetDownlinkMessage() *DownlinkMessage {
	if m != nil {
		return m.DownlinkMessage
	}
	return nil
}

type ScheduleDownlinkResponse struct {
	// The amount of time between the message has been scheduled and it will be transmitted by the gateway.
	Delay *types.Duration `protobuf:"bytes,1,opt,name=delay,proto3" json:"delay,omitempty"`
	// Downlink path chosen by the Gateway Server.
	DownlinkPath *DownlinkPath `protobuf:"bytes,2,opt,name=downlink_path,json=downlinkPath,proto3" json:"downlink_path,omitempty"`
	// Whether RX1 has been chosen for the downlink message.
	// Both RX1 and RX2 can be used for transmitting the same message by the same gateway.
	Rx1 bool `protobuf:"varint,3,opt,name=rx1,proto3" json:"rx1,omitempty"`
	// Whether RX2 has been chosen for the downlink message.
	// Both RX1 and RX2 can be used for transmitting the same message by the same gateway.
	Rx2                  bool     `protobuf:"varint,4,opt,name=rx2,proto3" json:"rx2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScheduleDownlinkResponse) Reset()         { *m = ScheduleDownlinkResponse{} }
func (m *ScheduleDownlinkResponse) String() string { return proto.CompactTextString(m) }
func (*ScheduleDownlinkResponse) ProtoMessage()    {}
func (*ScheduleDownlinkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62b07a36420f2d6d, []int{2}
}
func (m *ScheduleDownlinkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScheduleDownlinkResponse.Unmarshal(m, b)
}
func (m *ScheduleDownlinkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScheduleDownlinkResponse.Marshal(b, m, deterministic)
}
func (m *ScheduleDownlinkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleDownlinkResponse.Merge(m, src)
}
func (m *ScheduleDownlinkResponse) XXX_Size() int {
	return xxx_messageInfo_ScheduleDownlinkResponse.Size(m)
}
func (m *ScheduleDownlinkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleDownlinkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleDownlinkResponse proto.InternalMessageInfo

func (m *ScheduleDownlinkResponse) GetDelay() *types.Duration {
	if m != nil {
		return m.Delay
	}
	return nil
}

func (m *ScheduleDownlinkResponse) GetDownlinkPath() *DownlinkPath {
	if m != nil {
		return m.DownlinkPath
	}
	return nil
}

func (m *ScheduleDownlinkResponse) GetRx1() bool {
	if m != nil {
		return m.Rx1
	}
	return false
}

func (m *ScheduleDownlinkResponse) GetRx2() bool {
	if m != nil {
		return m.Rx2
	}
	return false
}

type ScheduleDownlinkErrorDetails struct {
	// Errors per path when downlink scheduling failed.
	PathErrors           []*ErrorDetails `protobuf:"bytes,1,rep,name=path_errors,json=pathErrors,proto3" json:"path_errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ScheduleDownlinkErrorDetails) Reset()         { *m = ScheduleDownlinkErrorDetails{} }
func (m *ScheduleDownlinkErrorDetails) String() string { return proto.CompactTextString(m) }
func (*ScheduleDownlinkErrorDetails) ProtoMessage()    {}
func (*ScheduleDownlinkErrorDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_62b07a36420f2d6d, []int{3}
}
func (m *ScheduleDownlinkErrorDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScheduleDownlinkErrorDetails.Unmarshal(m, b)
}
func (m *ScheduleDownlinkErrorDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScheduleDownlinkErrorDetails.Marshal(b, m, deterministic)
}
func (m *ScheduleDownlinkErrorDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleDownlinkErrorDetails.Merge(m, src)
}
func (m *ScheduleDownlinkErrorDetails) XXX_Size() int {
	return xxx_messageInfo_ScheduleDownlinkErrorDetails.Size(m)
}
func (m *ScheduleDownlinkErrorDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleDownlinkErrorDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleDownlinkErrorDetails proto.InternalMessageInfo

func (m *ScheduleDownlinkErrorDetails) GetPathErrors() []*ErrorDetails {
	if m != nil {
		return m.PathErrors
	}
	return nil
}

func init() {
	proto.RegisterType((*GatewayUp)(nil), "ttn.lorawan.v3.GatewayUp")
	golang_proto.RegisterType((*GatewayUp)(nil), "ttn.lorawan.v3.GatewayUp")
	proto.RegisterType((*GatewayDown)(nil), "ttn.lorawan.v3.GatewayDown")
	golang_proto.RegisterType((*GatewayDown)(nil), "ttn.lorawan.v3.GatewayDown")
	proto.RegisterType((*ScheduleDownlinkResponse)(nil), "ttn.lorawan.v3.ScheduleDownlinkResponse")
	golang_proto.RegisterType((*ScheduleDownlinkResponse)(nil), "ttn.lorawan.v3.ScheduleDownlinkResponse")
	proto.RegisterType((*ScheduleDownlinkErrorDetails)(nil), "ttn.lorawan.v3.ScheduleDownlinkErrorDetails")
	golang_proto.RegisterType((*ScheduleDownlinkErrorDetails)(nil), "ttn.lorawan.v3.ScheduleDownlinkErrorDetails")
}

func init() {
	proto.RegisterFile("lorawan-stack/api/gatewayserver.proto", fileDescriptor_62b07a36420f2d6d)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/gatewayserver.proto", fileDescriptor_62b07a36420f2d6d)
}

var fileDescriptor_62b07a36420f2d6d = []byte{
	// 783 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcf, 0x6e, 0xdb, 0x36,
	0x18, 0x87, 0xec, 0x64, 0xc8, 0xe8, 0x25, 0xf1, 0x08, 0x6c, 0x73, 0x1c, 0x27, 0x31, 0x34, 0x6c,
	0x08, 0x86, 0x59, 0xca, 0x14, 0x6c, 0xd8, 0x0e, 0x3b, 0x24, 0x71, 0x6a, 0xa4, 0x68, 0x8a, 0x56,
	0x49, 0x0a, 0xb4, 0x40, 0x60, 0xd0, 0x12, 0x4d, 0x0b, 0x96, 0x49, 0x55, 0xa4, 0xff, 0xa1, 0xe8,
	0xa5, 0xaf, 0xd0, 0x43, 0x0f, 0xbd, 0xf5, 0xd8, 0x67, 0xe8, 0xa1, 0xef, 0xd0, 0x07, 0xe8, 0xa1,
	0x3d, 0xf4, 0x19, 0x7a, 0x2a, 0x44, 0x51, 0x8e, 0x2d, 0x45, 0x6d, 0x2f, 0xbd, 0x91, 0xfc, 0x7e,
	0xdf, 0xef, 0xfb, 0xf1, 0xc7, 0xef, 0x23, 0xf8, 0xcd, 0x67, 0x21, 0x1a, 0x23, 0xda, 0xe0, 0x02,
	0x39, 0x7d, 0x13, 0x05, 0x9e, 0x49, 0x90, 0xc0, 0x63, 0x34, 0xe5, 0x38, 0x1c, 0xe1, 0xd0, 0x08,
	0x42, 0x26, 0x18, 0x5c, 0x13, 0x82, 0x1a, 0x0a, 0x6a, 0x8c, 0xf6, 0xab, 0x07, 0xc4, 0x13, 0xbd,
	0x61, 0xc7, 0x70, 0xd8, 0xc0, 0xc4, 0x74, 0xc4, 0xa6, 0x41, 0xc8, 0x26, 0x53, 0x53, 0x82, 0x9d,
	0x06, 0xc1, 0xb4, 0x31, 0x42, 0xbe, 0xe7, 0x22, 0x81, 0xcd, 0xcc, 0x22, 0xa6, 0xac, 0x36, 0xe6,
	0x28, 0x08, 0x23, 0x2c, 0x4e, 0xee, 0x0c, 0xbb, 0x72, 0x27, 0x37, 0x72, 0xa5, 0xe0, 0x35, 0xc2,
	0x18, 0xf1, 0xb1, 0x54, 0x88, 0x28, 0x65, 0x02, 0x09, 0x8f, 0x51, 0xae, 0xa2, 0xdb, 0x2a, 0x3a,
	0xe3, 0x70, 0x87, 0xa1, 0x04, 0xa8, 0xf8, 0x66, 0x3a, 0x8e, 0x07, 0x81, 0x98, 0xaa, 0xe0, 0x56,
	0xd6, 0x03, 0x1c, 0x86, 0x4c, 0xdd, 0xbd, 0xba, 0x93, 0x6b, 0x91, 0x02, 0xfc, 0x9a, 0x05, 0x78,
	0x2e, 0xa6, 0xc2, 0xeb, 0x7a, 0x38, 0xe4, 0xf9, 0x2c, 0x89, 0x9f, 0x31, 0xa0, 0x9e, 0x05, 0x0c,
	0x30, 0xe7, 0x88, 0xe0, 0x84, 0xa2, 0x76, 0x0d, 0xe2, 0xa1, 0x10, 0xf9, 0xf9, 0x21, 0x26, 0x1e,
	0xa3, 0xc8, 0x8f, 0x11, 0xfa, 0x07, 0x0d, 0x7c, 0xdf, 0x8a, 0x95, 0x5f, 0x04, 0xf0, 0x06, 0x58,
	0x1f, 0x06, 0xbe, 0x47, 0xfb, 0xed, 0xa4, 0x4c, 0x45, 0xab, 0x17, 0x77, 0x4b, 0xd6, 0x96, 0xb1,
	0xf8, 0xd8, 0xc6, 0x85, 0x84, 0x9d, 0xc6, 0x28, 0x7b, 0x6d, 0x38, 0xbf, 0xe5, 0xb0, 0x09, 0xd6,
	0x94, 0x1d, 0x6d, 0x2e, 0x90, 0x18, 0xf2, 0x4a, 0xa1, 0xae, 0x5d, 0x47, 0xa3, 0x4a, 0x9f, 0x49,
	0x90, 0xbd, 0x4a, 0xe6, 0xb7, 0xf0, 0x14, 0xfc, 0x28, 0x26, 0x6d, 0xe4, 0xf4, 0x29, 0x1b, 0xfb,
	0xd8, 0x25, 0x03, 0x4c, 0x45, 0xa5, 0x28, 0x89, 0xea, 0x69, 0xa2, 0xf3, 0xc9, 0xc1, 0x02, 0xce,
	0x2e, 0x8b, 0xd4, 0x89, 0x7e, 0x1f, 0x94, 0x54, 0xb9, 0x26, 0x1b, 0x53, 0x78, 0x13, 0x94, 0x5d,
	0x36, 0xa6, 0xf3, 0xb7, 0xad, 0x68, 0x92, 0x7c, 0x27, 0x4d, 0xde, 0x54, 0xb8, 0xe4, 0xba, 0xeb,
	0xee, 0xe2, 0x81, 0xfe, 0x4a, 0x03, 0x95, 0x33, 0xa7, 0x87, 0xdd, 0xa1, 0x8f, 0x13, 0xb0, 0x8d,
	0x79, 0xc0, 0x28, 0xc7, 0xf0, 0x3f, 0xb0, 0xec, 0x62, 0x1f, 0x4d, 0x15, 0xfb, 0x86, 0x11, 0xf7,
	0x9d, 0x91, 0xf4, 0x9d, 0xd1, 0x54, 0x7d, 0x79, 0xb8, 0xf2, 0xf1, 0x70, 0xf9, 0xa5, 0x56, 0x58,
	0xd1, 0xec, 0x38, 0x03, 0x1e, 0x80, 0xd5, 0x99, 0xc6, 0x00, 0x89, 0x9e, 0xb2, 0xb1, 0x96, 0x27,
	0xf0, 0x0e, 0x12, 0x3d, 0xfb, 0x07, 0x77, 0x6e, 0x07, 0xcb, 0xa0, 0x18, 0x4e, 0xfe, 0x92, 0xb6,
	0xad, 0xd8, 0xd1, 0x32, 0x3e, 0xb1, 0x2a, 0x4b, 0xc9, 0x89, 0xa5, 0x5f, 0x82, 0x5a, 0x5a, 0xfd,
	0x71, 0xd4, 0xec, 0x4d, 0x2c, 0x90, 0xe7, 0x73, 0xf8, 0x3f, 0x28, 0x45, 0xd5, 0xdb, 0x72, 0x02,
	0x92, 0x96, 0xc8, 0x88, 0x98, 0x4f, 0xb1, 0x41, 0x94, 0x20, 0x4f, 0xb8, 0xf5, 0xb6, 0x08, 0x96,
	0x5b, 0x62, 0xdc, 0xe2, 0xf0, 0x04, 0x94, 0x6e, 0x79, 0xb4, 0xaf, 0x9e, 0x01, 0x6e, 0xe4, 0xb4,
	0xc3, 0x45, 0x50, 0xdd, 0xcc, 0x09, 0x45, 0xfa, 0x76, 0xb5, 0x3d, 0x0d, 0x9e, 0x81, 0x9f, 0x5a,
	0x58, 0x1c, 0x31, 0xea, 0x60, 0x2a, 0x42, 0x24, 0x58, 0x78, 0xc4, 0x68, 0xd7, 0x23, 0xf0, 0xe7,
	0x8c, 0xbf, 0xc7, 0xd1, 0x5c, 0x57, 0xf5, 0x34, 0xe3, 0x35, 0xb9, 0xcf, 0x34, 0xc9, 0x7a, 0x7a,
	0xf7, 0xfc, 0xfc, 0x88, 0x51, 0x8a, 0x9d, 0xe8, 0x59, 0x4e, 0x68, 0x97, 0x41, 0x3d, 0x47, 0xcf,
	0xc9, 0xd5, 0x50, 0x67, 0x2b, 0x64, 0x79, 0xf4, 0x7f, 0x9e, 0xbc, 0x79, 0xff, 0xb4, 0xb0, 0x07,
	0x0d, 0x93, 0xf0, 0xd9, 0xaf, 0x6a, 0x3e, 0x4a, 0xa6, 0xc5, 0x73, 0x1f, 0xcb, 0xf9, 0x6d, 0x38,
	0xb3, 0xb4, 0x86, 0x17, 0xd5, 0x7f, 0xae, 0x81, 0x5f, 0x94, 0xb2, 0x7b, 0xd6, 0x37, 0xd2, 0xf6,
	0xaf, 0xd4, 0x66, 0xc1, 0xbd, 0xcf, 0x6b, 0x1b, 0x59, 0x69, 0x75, 0x16, 0x06, 0x4b, 0xb7, 0x79,
	0x8b, 0xc3, 0x4b, 0x50, 0x4e, 0x37, 0x12, 0xfc, 0xd2, 0x34, 0x55, 0x77, 0xd3, 0x80, 0xbc, 0x49,
	0xb2, 0x5e, 0x68, 0xa0, 0xd0, 0xe2, 0x91, 0x17, 0x1b, 0x2d, 0x2c, 0xd4, 0x2d, 0xaf, 0x2e, 0x11,
	0x7d, 0x1b, 0xfc, 0xab, 0xdc, 0xf8, 0x3d, 0x07, 0x93, 0xe2, 0xd2, 0x2d, 0xe9, 0xc8, 0x9f, 0xf0,
	0x8f, 0x7c, 0x47, 0xae, 0xac, 0x30, 0xa3, 0xaf, 0x8e, 0x1f, 0xfe, 0xfd, 0xfa, 0xdd, 0xb6, 0xf6,
	0xc0, 0x24, 0xcc, 0x10, 0x3d, 0x2c, 0x7a, 0x1e, 0x25, 0xdc, 0xa0, 0x58, 0x8c, 0x59, 0xd8, 0x37,
	0x17, 0xbf, 0xe3, 0xd1, 0xbe, 0x19, 0xf4, 0x89, 0x29, 0x04, 0x0d, 0x3a, 0x9d, 0xef, 0x64, 0xbb,
	0xee, 0x7f, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x63, 0x6c, 0x96, 0xa2, 0x7d, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GtwGsClient is the client API for GtwGs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GtwGsClient interface {
	// Link a gateway to the Gateway Server for streaming upstream messages and downstream messages.
	LinkGateway(ctx context.Context, opts ...grpc.CallOption) (GtwGs_LinkGatewayClient, error)
	// Get configuration for the concentrator.
	GetConcentratorConfig(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*ConcentratorConfig, error)
	// Get connection information to connect an MQTT gateway.
	GetMQTTConnectionInfo(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*MQTTConnectionInfo, error)
	// Get legacy connection information to connect a The Things Network Stack V2 MQTT gateway.
	GetMQTTV2ConnectionInfo(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*MQTTConnectionInfo, error)
}

type gtwGsClient struct {
	cc *grpc.ClientConn
}

func NewGtwGsClient(cc *grpc.ClientConn) GtwGsClient {
	return &gtwGsClient{cc}
}

func (c *gtwGsClient) LinkGateway(ctx context.Context, opts ...grpc.CallOption) (GtwGs_LinkGatewayClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GtwGs_serviceDesc.Streams[0], "/ttn.lorawan.v3.GtwGs/LinkGateway", opts...)
	if err != nil {
		return nil, err
	}
	x := &gtwGsLinkGatewayClient{stream}
	return x, nil
}

type GtwGs_LinkGatewayClient interface {
	Send(*GatewayUp) error
	Recv() (*GatewayDown, error)
	grpc.ClientStream
}

type gtwGsLinkGatewayClient struct {
	grpc.ClientStream
}

func (x *gtwGsLinkGatewayClient) Send(m *GatewayUp) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gtwGsLinkGatewayClient) Recv() (*GatewayDown, error) {
	m := new(GatewayDown)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gtwGsClient) GetConcentratorConfig(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*ConcentratorConfig, error) {
	out := new(ConcentratorConfig)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GtwGs/GetConcentratorConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gtwGsClient) GetMQTTConnectionInfo(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*MQTTConnectionInfo, error) {
	out := new(MQTTConnectionInfo)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GtwGs/GetMQTTConnectionInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gtwGsClient) GetMQTTV2ConnectionInfo(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*MQTTConnectionInfo, error) {
	out := new(MQTTConnectionInfo)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GtwGs/GetMQTTV2ConnectionInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GtwGsServer is the server API for GtwGs service.
type GtwGsServer interface {
	// Link a gateway to the Gateway Server for streaming upstream messages and downstream messages.
	LinkGateway(GtwGs_LinkGatewayServer) error
	// Get configuration for the concentrator.
	GetConcentratorConfig(context.Context, *types.Empty) (*ConcentratorConfig, error)
	// Get connection information to connect an MQTT gateway.
	GetMQTTConnectionInfo(context.Context, *GatewayIdentifiers) (*MQTTConnectionInfo, error)
	// Get legacy connection information to connect a The Things Network Stack V2 MQTT gateway.
	GetMQTTV2ConnectionInfo(context.Context, *GatewayIdentifiers) (*MQTTConnectionInfo, error)
}

// UnimplementedGtwGsServer can be embedded to have forward compatible implementations.
type UnimplementedGtwGsServer struct {
}

func (*UnimplementedGtwGsServer) LinkGateway(srv GtwGs_LinkGatewayServer) error {
	return status.Errorf(codes.Unimplemented, "method LinkGateway not implemented")
}
func (*UnimplementedGtwGsServer) GetConcentratorConfig(ctx context.Context, req *types.Empty) (*ConcentratorConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConcentratorConfig not implemented")
}
func (*UnimplementedGtwGsServer) GetMQTTConnectionInfo(ctx context.Context, req *GatewayIdentifiers) (*MQTTConnectionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMQTTConnectionInfo not implemented")
}
func (*UnimplementedGtwGsServer) GetMQTTV2ConnectionInfo(ctx context.Context, req *GatewayIdentifiers) (*MQTTConnectionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMQTTV2ConnectionInfo not implemented")
}

func RegisterGtwGsServer(s *grpc.Server, srv GtwGsServer) {
	s.RegisterService(&_GtwGs_serviceDesc, srv)
}

func _GtwGs_LinkGateway_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GtwGsServer).LinkGateway(&gtwGsLinkGatewayServer{stream})
}

type GtwGs_LinkGatewayServer interface {
	Send(*GatewayDown) error
	Recv() (*GatewayUp, error)
	grpc.ServerStream
}

type gtwGsLinkGatewayServer struct {
	grpc.ServerStream
}

func (x *gtwGsLinkGatewayServer) Send(m *GatewayDown) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gtwGsLinkGatewayServer) Recv() (*GatewayUp, error) {
	m := new(GatewayUp)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GtwGs_GetConcentratorConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GtwGsServer).GetConcentratorConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GtwGs/GetConcentratorConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GtwGsServer).GetConcentratorConfig(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GtwGs_GetMQTTConnectionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GtwGsServer).GetMQTTConnectionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GtwGs/GetMQTTConnectionInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GtwGsServer).GetMQTTConnectionInfo(ctx, req.(*GatewayIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _GtwGs_GetMQTTV2ConnectionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GtwGsServer).GetMQTTV2ConnectionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GtwGs/GetMQTTV2ConnectionInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GtwGsServer).GetMQTTV2ConnectionInfo(ctx, req.(*GatewayIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _GtwGs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.GtwGs",
	HandlerType: (*GtwGsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConcentratorConfig",
			Handler:    _GtwGs_GetConcentratorConfig_Handler,
		},
		{
			MethodName: "GetMQTTConnectionInfo",
			Handler:    _GtwGs_GetMQTTConnectionInfo_Handler,
		},
		{
			MethodName: "GetMQTTV2ConnectionInfo",
			Handler:    _GtwGs_GetMQTTV2ConnectionInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LinkGateway",
			Handler:       _GtwGs_LinkGateway_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "lorawan-stack/api/gatewayserver.proto",
}

// NsGsClient is the client API for NsGs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NsGsClient interface {
	// Instructs the Gateway Server to schedule a downlink message.
	// The Gateway Server may refuse if there are any conflicts in the schedule or
	// if a duty cycle prevents the gateway from transmitting.
	ScheduleDownlink(ctx context.Context, in *DownlinkMessage, opts ...grpc.CallOption) (*ScheduleDownlinkResponse, error)
}

type nsGsClient struct {
	cc *grpc.ClientConn
}

func NewNsGsClient(cc *grpc.ClientConn) NsGsClient {
	return &nsGsClient{cc}
}

func (c *nsGsClient) ScheduleDownlink(ctx context.Context, in *DownlinkMessage, opts ...grpc.CallOption) (*ScheduleDownlinkResponse, error) {
	out := new(ScheduleDownlinkResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.NsGs/ScheduleDownlink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NsGsServer is the server API for NsGs service.
type NsGsServer interface {
	// Instructs the Gateway Server to schedule a downlink message.
	// The Gateway Server may refuse if there are any conflicts in the schedule or
	// if a duty cycle prevents the gateway from transmitting.
	ScheduleDownlink(context.Context, *DownlinkMessage) (*ScheduleDownlinkResponse, error)
}

// UnimplementedNsGsServer can be embedded to have forward compatible implementations.
type UnimplementedNsGsServer struct {
}

func (*UnimplementedNsGsServer) ScheduleDownlink(ctx context.Context, req *DownlinkMessage) (*ScheduleDownlinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleDownlink not implemented")
}

func RegisterNsGsServer(s *grpc.Server, srv NsGsServer) {
	s.RegisterService(&_NsGs_serviceDesc, srv)
}

func _NsGs_ScheduleDownlink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownlinkMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsGsServer).ScheduleDownlink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.NsGs/ScheduleDownlink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsGsServer).ScheduleDownlink(ctx, req.(*DownlinkMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _NsGs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.NsGs",
	HandlerType: (*NsGsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScheduleDownlink",
			Handler:    _NsGs_ScheduleDownlink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/gatewayserver.proto",
}

// GsClient is the client API for Gs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GsClient interface {
	// Get statistics about the current gateway connection to the Gateway Server.
	// This is not persisted between reconnects.
	GetGatewayConnectionStats(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*GatewayConnectionStats, error)
}

type gsClient struct {
	cc *grpc.ClientConn
}

func NewGsClient(cc *grpc.ClientConn) GsClient {
	return &gsClient{cc}
}

func (c *gsClient) GetGatewayConnectionStats(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*GatewayConnectionStats, error) {
	out := new(GatewayConnectionStats)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.Gs/GetGatewayConnectionStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GsServer is the server API for Gs service.
type GsServer interface {
	// Get statistics about the current gateway connection to the Gateway Server.
	// This is not persisted between reconnects.
	GetGatewayConnectionStats(context.Context, *GatewayIdentifiers) (*GatewayConnectionStats, error)
}

// UnimplementedGsServer can be embedded to have forward compatible implementations.
type UnimplementedGsServer struct {
}

func (*UnimplementedGsServer) GetGatewayConnectionStats(ctx context.Context, req *GatewayIdentifiers) (*GatewayConnectionStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGatewayConnectionStats not implemented")
}

func RegisterGsServer(s *grpc.Server, srv GsServer) {
	s.RegisterService(&_Gs_serviceDesc, srv)
}

func _Gs_GetGatewayConnectionStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GsServer).GetGatewayConnectionStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.Gs/GetGatewayConnectionStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GsServer).GetGatewayConnectionStats(ctx, req.(*GatewayIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.Gs",
	HandlerType: (*GsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGatewayConnectionStats",
			Handler:    _Gs_GetGatewayConnectionStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/gatewayserver.proto",
}
