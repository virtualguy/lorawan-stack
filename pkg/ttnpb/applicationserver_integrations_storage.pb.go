// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/applicationserver_integrations_storage.proto

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

type GetStoredApplicationUpRequest struct {
	// Query upstream messages from all end devices of an application. Cannot be used in conjunction with end_device_ids.
	ApplicationIds *ApplicationIdentifiers `protobuf:"bytes,1,opt,name=application_ids,json=applicationIds,proto3" json:"application_ids,omitempty"`
	// Query upstream messages from a single end device. Cannot be used in conjunction with application_ids.
	EndDeviceIds *EndDeviceIdentifiers `protobuf:"bytes,2,opt,name=end_device_ids,json=endDeviceIds,proto3" json:"end_device_ids,omitempty"`
	// Query upstream messages of a specific type. If not set, then all upstream messages are returned.
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// Limit number of results.
	Limit *types.UInt32Value `protobuf:"bytes,4,opt,name=limit,proto3" json:"limit,omitempty"`
	// Query upstream messages after this timestamp only. Cannot be used in conjunction with last.
	After *types.Timestamp `protobuf:"bytes,5,opt,name=after,proto3" json:"after,omitempty"`
	// Query upstream messages before this timestamp only. Cannot be used in conjunction with last.
	Before *types.Timestamp `protobuf:"bytes,6,opt,name=before,proto3" json:"before,omitempty"`
	// Query uplinks on a specific FPort only.
	FPort *types.UInt32Value `protobuf:"bytes,7,opt,name=f_port,json=fPort,proto3" json:"f_port,omitempty"`
	// Order results.
	Order string `protobuf:"bytes,8,opt,name=order,proto3" json:"order,omitempty"`
	// The names of the upstream message fields that should be returned. See the API reference
	// for allowed field names for each type of upstream message.
	FieldMask *types.FieldMask `protobuf:"bytes,9,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	// Query upstream messages that have arrived in the last minutes or hours. Cannot be used in conjunction with after and before.
	Last                 *types.Duration `protobuf:"bytes,10,opt,name=last,proto3" json:"last,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetStoredApplicationUpRequest) Reset()         { *m = GetStoredApplicationUpRequest{} }
func (m *GetStoredApplicationUpRequest) String() string { return proto.CompactTextString(m) }
func (*GetStoredApplicationUpRequest) ProtoMessage()    {}
func (*GetStoredApplicationUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff0e9f52f73d254, []int{0}
}
func (m *GetStoredApplicationUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStoredApplicationUpRequest.Unmarshal(m, b)
}
func (m *GetStoredApplicationUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStoredApplicationUpRequest.Marshal(b, m, deterministic)
}
func (m *GetStoredApplicationUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoredApplicationUpRequest.Merge(m, src)
}
func (m *GetStoredApplicationUpRequest) XXX_Size() int {
	return xxx_messageInfo_GetStoredApplicationUpRequest.Size(m)
}
func (m *GetStoredApplicationUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoredApplicationUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoredApplicationUpRequest proto.InternalMessageInfo

func (m *GetStoredApplicationUpRequest) GetApplicationIds() *ApplicationIdentifiers {
	if m != nil {
		return m.ApplicationIds
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetEndDeviceIds() *EndDeviceIdentifiers {
	if m != nil {
		return m.EndDeviceIds
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetStoredApplicationUpRequest) GetLimit() *types.UInt32Value {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetAfter() *types.Timestamp {
	if m != nil {
		return m.After
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetBefore() *types.Timestamp {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetFPort() *types.UInt32Value {
	if m != nil {
		return m.FPort
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

func (m *GetStoredApplicationUpRequest) GetFieldMask() *types.FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetLast() *types.Duration {
	if m != nil {
		return m.Last
	}
	return nil
}

type GetStoredApplicationUpCountRequest struct {
	// Count upstream messages from all end devices of an application. Cannot be used in conjunction with end_device_ids.
	ApplicationIds *ApplicationIdentifiers `protobuf:"bytes,1,opt,name=application_ids,json=applicationIds,proto3" json:"application_ids,omitempty"`
	// Count upstream messages from a single end device. Cannot be used in conjunction with application_ids.
	EndDeviceIds *EndDeviceIdentifiers `protobuf:"bytes,2,opt,name=end_device_ids,json=endDeviceIds,proto3" json:"end_device_ids,omitempty"`
	// Count upstream messages of a specific type. If not set, then all upstream messages are returned.
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// Count upstream messages after this timestamp only. Cannot be used in conjunction with last.
	After *types.Timestamp `protobuf:"bytes,4,opt,name=after,proto3" json:"after,omitempty"`
	// Count upstream messages before this timestamp only. Cannot be used in conjunction with last.
	Before *types.Timestamp `protobuf:"bytes,5,opt,name=before,proto3" json:"before,omitempty"`
	// Count uplinks on a specific FPort only.
	FPort *types.UInt32Value `protobuf:"bytes,6,opt,name=f_port,json=fPort,proto3" json:"f_port,omitempty"`
	// Count upstream messages that have arrived in the last minutes or hours. Cannot be used in conjunction with after and before.
	Last                 *types.Duration `protobuf:"bytes,7,opt,name=last,proto3" json:"last,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetStoredApplicationUpCountRequest) Reset()         { *m = GetStoredApplicationUpCountRequest{} }
func (m *GetStoredApplicationUpCountRequest) String() string { return proto.CompactTextString(m) }
func (*GetStoredApplicationUpCountRequest) ProtoMessage()    {}
func (*GetStoredApplicationUpCountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff0e9f52f73d254, []int{1}
}
func (m *GetStoredApplicationUpCountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStoredApplicationUpCountRequest.Unmarshal(m, b)
}
func (m *GetStoredApplicationUpCountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStoredApplicationUpCountRequest.Marshal(b, m, deterministic)
}
func (m *GetStoredApplicationUpCountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoredApplicationUpCountRequest.Merge(m, src)
}
func (m *GetStoredApplicationUpCountRequest) XXX_Size() int {
	return xxx_messageInfo_GetStoredApplicationUpCountRequest.Size(m)
}
func (m *GetStoredApplicationUpCountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoredApplicationUpCountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoredApplicationUpCountRequest proto.InternalMessageInfo

func (m *GetStoredApplicationUpCountRequest) GetApplicationIds() *ApplicationIdentifiers {
	if m != nil {
		return m.ApplicationIds
	}
	return nil
}

func (m *GetStoredApplicationUpCountRequest) GetEndDeviceIds() *EndDeviceIdentifiers {
	if m != nil {
		return m.EndDeviceIds
	}
	return nil
}

func (m *GetStoredApplicationUpCountRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetStoredApplicationUpCountRequest) GetAfter() *types.Timestamp {
	if m != nil {
		return m.After
	}
	return nil
}

func (m *GetStoredApplicationUpCountRequest) GetBefore() *types.Timestamp {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *GetStoredApplicationUpCountRequest) GetFPort() *types.UInt32Value {
	if m != nil {
		return m.FPort
	}
	return nil
}

func (m *GetStoredApplicationUpCountRequest) GetLast() *types.Duration {
	if m != nil {
		return m.Last
	}
	return nil
}

type GetStoredApplicationUpCountResponse struct {
	// Number of stored messages by end device ID.
	Count                map[string]uint32 `protobuf:"bytes,1,rep,name=count,proto3" json:"count,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetStoredApplicationUpCountResponse) Reset()         { *m = GetStoredApplicationUpCountResponse{} }
func (m *GetStoredApplicationUpCountResponse) String() string { return proto.CompactTextString(m) }
func (*GetStoredApplicationUpCountResponse) ProtoMessage()    {}
func (*GetStoredApplicationUpCountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff0e9f52f73d254, []int{2}
}
func (m *GetStoredApplicationUpCountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStoredApplicationUpCountResponse.Unmarshal(m, b)
}
func (m *GetStoredApplicationUpCountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStoredApplicationUpCountResponse.Marshal(b, m, deterministic)
}
func (m *GetStoredApplicationUpCountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoredApplicationUpCountResponse.Merge(m, src)
}
func (m *GetStoredApplicationUpCountResponse) XXX_Size() int {
	return xxx_messageInfo_GetStoredApplicationUpCountResponse.Size(m)
}
func (m *GetStoredApplicationUpCountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoredApplicationUpCountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoredApplicationUpCountResponse proto.InternalMessageInfo

func (m *GetStoredApplicationUpCountResponse) GetCount() map[string]uint32 {
	if m != nil {
		return m.Count
	}
	return nil
}

func init() {
	proto.RegisterType((*GetStoredApplicationUpRequest)(nil), "ttn.lorawan.v3.GetStoredApplicationUpRequest")
	golang_proto.RegisterType((*GetStoredApplicationUpRequest)(nil), "ttn.lorawan.v3.GetStoredApplicationUpRequest")
	proto.RegisterType((*GetStoredApplicationUpCountRequest)(nil), "ttn.lorawan.v3.GetStoredApplicationUpCountRequest")
	golang_proto.RegisterType((*GetStoredApplicationUpCountRequest)(nil), "ttn.lorawan.v3.GetStoredApplicationUpCountRequest")
	proto.RegisterType((*GetStoredApplicationUpCountResponse)(nil), "ttn.lorawan.v3.GetStoredApplicationUpCountResponse")
	golang_proto.RegisterType((*GetStoredApplicationUpCountResponse)(nil), "ttn.lorawan.v3.GetStoredApplicationUpCountResponse")
	proto.RegisterMapType((map[string]uint32)(nil), "ttn.lorawan.v3.GetStoredApplicationUpCountResponse.CountEntry")
	golang_proto.RegisterMapType((map[string]uint32)(nil), "ttn.lorawan.v3.GetStoredApplicationUpCountResponse.CountEntry")
}

func init() {
	proto.RegisterFile("lorawan-stack/api/applicationserver_integrations_storage.proto", fileDescriptor_6ff0e9f52f73d254)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/applicationserver_integrations_storage.proto", fileDescriptor_6ff0e9f52f73d254)
}

var fileDescriptor_6ff0e9f52f73d254 = []byte{
	// 894 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0x3f, 0x6f, 0x1c, 0x45,
	0x14, 0xcf, 0x9e, 0xef, 0x2e, 0x78, 0xec, 0xd8, 0x68, 0x14, 0xa1, 0xe5, 0x88, 0x93, 0xd3, 0x05,
	0x21, 0x37, 0xb7, 0x1b, 0xdd, 0x09, 0xc9, 0x50, 0x44, 0x8a, 0x49, 0x40, 0x06, 0xa1, 0xc0, 0x24,
	0xa6, 0x70, 0xb3, 0x9a, 0xdb, 0x7d, 0xb7, 0x1e, 0x6e, 0x6f, 0x66, 0x32, 0x33, 0xbb, 0xe6, 0x64,
	0xb9, 0x80, 0xaf, 0xc0, 0x97, 0x40, 0xa2, 0xa0, 0xa2, 0x43, 0x82, 0x82, 0x8a, 0x0e, 0x28, 0xe8,
	0x68, 0x40, 0xe2, 0x1b, 0x50, 0xa4, 0x42, 0x3b, 0xbb, 0xf7, 0x3f, 0x76, 0x2e, 0x11, 0x82, 0x26,
	0xdd, 0xbc, 0x37, 0xbf, 0xf7, 0x9b, 0x37, 0xef, 0xfd, 0x9e, 0x66, 0xd0, 0xed, 0x44, 0x28, 0x7a,
	0x42, 0x79, 0x5b, 0x1b, 0x1a, 0x0e, 0x7c, 0x2a, 0x99, 0x4f, 0xa5, 0x4c, 0x58, 0x48, 0x0d, 0x13,
	0x5c, 0x83, 0xca, 0x40, 0x05, 0x8c, 0x1b, 0x88, 0x55, 0xe1, 0x09, 0xb4, 0x11, 0x8a, 0xc6, 0xe0,
	0x49, 0x25, 0x8c, 0xc0, 0x5b, 0xc6, 0x70, 0xaf, 0xe4, 0xf0, 0xb2, 0x6e, 0xe3, 0x4e, 0xcc, 0xcc,
	0x71, 0xda, 0xf3, 0x42, 0x31, 0xf4, 0x81, 0x67, 0x62, 0x24, 0x95, 0xf8, 0x6c, 0xe4, 0x5b, 0x70,
	0xd8, 0x8e, 0x81, 0xb7, 0x33, 0x9a, 0xb0, 0x88, 0x1a, 0xf0, 0x97, 0x16, 0x05, 0x65, 0xa3, 0x3d,
	0x43, 0x11, 0x8b, 0x58, 0x14, 0xc1, 0xbd, 0xb4, 0x6f, 0x2d, 0x6b, 0xd8, 0x55, 0x09, 0xbf, 0x16,
	0x0b, 0x11, 0x27, 0x50, 0xa4, 0xce, 0xb9, 0x30, 0x45, 0x9e, 0xe5, 0x6e, 0xb3, 0xdc, 0x9d, 0x70,
	0xf4, 0x19, 0x24, 0x51, 0x30, 0xa4, 0x7a, 0x50, 0x22, 0x6e, 0x2c, 0x22, 0x0c, 0x1b, 0x82, 0x36,
	0x74, 0x28, 0x4b, 0xc0, 0xf5, 0x45, 0x40, 0x94, 0x16, 0xb5, 0x38, 0x6f, 0xff, 0x44, 0x51, 0x29,
	0x41, 0x8d, 0x53, 0xb8, 0xb9, 0x5c, 0x62, 0x16, 0x01, 0x37, 0xac, 0xcf, 0xa6, 0xa0, 0xe6, 0x32,
	0x68, 0x08, 0x5a, 0xd3, 0x18, 0x4a, 0x44, 0xeb, 0xef, 0x1a, 0xda, 0x79, 0x0f, 0xcc, 0x03, 0x23,
	0x14, 0x44, 0x77, 0xa6, 0x3d, 0x3a, 0x94, 0x04, 0x1e, 0xa5, 0xa0, 0x0d, 0xbe, 0x8f, 0xb6, 0x67,
	0x7a, 0x17, 0xb0, 0x48, 0xbb, 0x4e, 0xd3, 0xd9, 0xdd, 0xe8, 0xbc, 0xe1, 0xcd, 0x77, 0xc9, 0x9b,
	0x09, 0x3f, 0x98, 0xa6, 0x42, 0xb6, 0xe8, 0xac, 0x5f, 0xe3, 0xf7, 0xd1, 0x16, 0xf0, 0x28, 0x88,
	0x20, 0x63, 0x21, 0x58, 0xbe, 0x8a, 0xe5, 0x7b, 0x7d, 0x91, 0xef, 0x1e, 0x8f, 0xee, 0x5a, 0xd0,
	0x2c, 0xdb, 0x26, 0x4c, 0xbd, 0x1a, 0xff, 0xe8, 0xa0, 0xaa, 0x19, 0x49, 0x70, 0xd7, 0x9a, 0xce,
	0xee, 0xfa, 0xfe, 0x37, 0xce, 0xe3, 0xfd, 0xaf, 0x1d, 0xf5, 0x95, 0x43, 0x2e, 0x91, 0xad, 0x54,
	0x26, 0x8c, 0x0f, 0x82, 0xf2, 0xc2, 0x64, 0xe3, 0x53, 0xc1, 0x78, 0x40, 0xc3, 0x10, 0xa4, 0x21,
	0x9b, 0x91, 0x38, 0xe1, 0x76, 0x9b, 0x86, 0x03, 0x72, 0x65, 0x62, 0xf1, 0x79, 0x53, 0x03, 0x37,
	0x64, 0x7b, 0x62, 0xf6, 0x29, 0x4b, 0x20, 0x9a, 0x71, 0x3c, 0x4a, 0x21, 0x85, 0x88, 0x34, 0xe6,
	0x1d, 0x01, 0xe3, 0x63, 0xf1, 0x45, 0x64, 0x3b, 0x11, 0x65, 0xe5, 0xb4, 0x48, 0x32, 0x88, 0xc8,
	0x66, 0x2e, 0xff, 0xfc, 0xe6, 0x11, 0x35, 0x94, 0xd8, 0xec, 0x71, 0x07, 0xd5, 0x12, 0x36, 0x64,
	0xc6, 0xad, 0xda, 0x4a, 0x5c, 0xf3, 0x8a, 0xe6, 0x7b, 0xe3, 0xe6, 0x7b, 0x87, 0x07, 0xdc, 0x74,
	0x3b, 0x9f, 0xd0, 0x24, 0x05, 0x52, 0x40, 0xf1, 0x2d, 0x54, 0xa3, 0x7d, 0x03, 0xca, 0xad, 0xd9,
	0x98, 0xc6, 0x52, 0xcc, 0xc3, 0xb1, 0xe2, 0x48, 0x01, 0xc4, 0x1d, 0x54, 0xef, 0x41, 0x5f, 0x28,
	0x70, 0xeb, 0x4f, 0x0d, 0x29, 0x91, 0xb8, 0x8b, 0xea, 0xfd, 0x40, 0x0a, 0x65, 0xdc, 0xcb, 0xab,
	0xa4, 0xd6, 0xff, 0x48, 0x28, 0x83, 0xf7, 0x50, 0x4d, 0xa8, 0x08, 0x94, 0xfb, 0x92, 0xed, 0x4a,
	0xeb, 0xf1, 0xfe, 0x0d, 0xb5, 0x43, 0x2e, 0x91, 0xcd, 0xb6, 0x82, 0x10, 0x58, 0x06, 0x51, 0x40,
	0x0d, 0xd9, 0x98, 0x35, 0x8a, 0x00, 0xfc, 0x16, 0x42, 0xd3, 0x51, 0x72, 0xd7, 0xcf, 0x49, 0xf3,
	0xdd, 0x1c, 0xf2, 0x21, 0xd5, 0x03, 0xb2, 0xde, 0x1f, 0x2f, 0x71, 0x1b, 0x55, 0x13, 0xaa, 0x8d,
	0x8b, 0x6c, 0xd0, 0xab, 0x4b, 0x41, 0x77, 0xcb, 0xf9, 0x22, 0x16, 0xd6, 0xfa, 0xa5, 0x8a, 0x5a,
	0x4f, 0x16, 0xfe, 0x3b, 0x22, 0xe5, 0xe6, 0x85, 0xfa, 0xff, 0x3b, 0xf5, 0x4f, 0x94, 0x5c, 0x7d,
	0x76, 0x25, 0xd7, 0x9e, 0x43, 0xc9, 0xf5, 0xd5, 0x95, 0x3c, 0x16, 0xd5, 0xe5, 0xd5, 0x44, 0xf5,
	0xad, 0x83, 0x6e, 0x5e, 0x28, 0x2a, 0x2d, 0xf3, 0xe7, 0x0f, 0x3f, 0x44, 0xb5, 0x30, 0x77, 0xb8,
	0x4e, 0x73, 0x6d, 0x77, 0xa3, 0x73, 0x7b, 0xb1, 0xf7, 0x2b, 0x70, 0x78, 0xd6, 0xba, 0xc7, 0x8d,
	0x1a, 0x91, 0x82, 0xac, 0xb1, 0x87, 0xd0, 0xd4, 0x89, 0x5f, 0x46, 0x6b, 0x03, 0x18, 0x59, 0xb5,
	0xae, 0x93, 0x7c, 0x89, 0xaf, 0xa2, 0x5a, 0x96, 0x5f, 0xce, 0x2a, 0xee, 0x0a, 0x29, 0x8c, 0xb7,
	0x2b, 0x7b, 0x4e, 0xe7, 0xfb, 0x1a, 0xba, 0x3a, 0x77, 0xd4, 0x83, 0xe2, 0x39, 0xc6, 0xdf, 0x55,
	0xd0, 0x2b, 0x4f, 0x4e, 0x06, 0xb7, 0x57, 0x4b, 0xba, 0x1c, 0xa4, 0xc6, 0xce, 0x05, 0xf3, 0x72,
	0x28, 0x5b, 0x3f, 0x3b, 0x5f, 0xfc, 0xfa, 0xe7, 0x97, 0x95, 0x9f, 0x1c, 0x7c, 0xea, 0x53, 0x3d,
	0xf7, 0x5b, 0xf0, 0x4f, 0xe7, 0x07, 0xc6, 0x5b, 0x18, 0xc8, 0x05, 0xfb, 0xcc, 0x2f, 0xa0, 0xcb,
	0x71, 0x93, 0xe5, 0x99, 0x2f, 0x69, 0x38, 0xc8, 0x5f, 0x41, 0xbf, 0xfc, 0x77, 0xf8, 0xa7, 0xb9,
	0x12, 0xcf, 0x8e, 0x3e, 0xc0, 0x07, 0xcb, 0xc7, 0x3f, 0xed, 0xbc, 0x73, 0xc8, 0x6e, 0x39, 0xf8,
	0xaf, 0x0a, 0x7a, 0xed, 0x82, 0x5e, 0xe2, 0xce, 0x33, 0x35, 0xbe, 0x28, 0x64, 0xf7, 0x39, 0xc4,
	0xd2, 0xfa, 0xbd, 0x28, 0xef, 0x6f, 0x0e, 0xfe, 0xdc, 0xf9, 0x1f, 0xeb, 0xeb, 0x5b, 0xa1, 0x1e,
	0x7d, 0x8c, 0xef, 0xff, 0x6b, 0x55, 0x2e, 0x28, 0xf7, 0xdf, 0xfc, 0xe1, 0x8f, 0xeb, 0xce, 0x91,
	0x1f, 0x0b, 0xcf, 0x1c, 0x83, 0x39, 0x66, 0x3c, 0xd6, 0x1e, 0x07, 0x73, 0x22, 0xd4, 0xc0, 0x9f,
	0xff, 0x04, 0x65, 0x5d, 0x5f, 0x0e, 0x62, 0xdf, 0x18, 0x2e, 0x7b, 0xbd, 0xba, 0x9d, 0xe4, 0xee,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xe3, 0x5a, 0xd4, 0xb1, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApplicationUpStorageClient is the client API for ApplicationUpStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApplicationUpStorageClient interface {
	// Returns a stream of application messages that have been stored in the database.
	GetStoredApplicationUp(ctx context.Context, in *GetStoredApplicationUpRequest, opts ...grpc.CallOption) (ApplicationUpStorage_GetStoredApplicationUpClient, error)
	// Returns how many application messages have been stored in the database for an application or end device.
	GetStoredApplicationUpCount(ctx context.Context, in *GetStoredApplicationUpCountRequest, opts ...grpc.CallOption) (*GetStoredApplicationUpCountResponse, error)
}

type applicationUpStorageClient struct {
	cc *grpc.ClientConn
}

func NewApplicationUpStorageClient(cc *grpc.ClientConn) ApplicationUpStorageClient {
	return &applicationUpStorageClient{cc}
}

func (c *applicationUpStorageClient) GetStoredApplicationUp(ctx context.Context, in *GetStoredApplicationUpRequest, opts ...grpc.CallOption) (ApplicationUpStorage_GetStoredApplicationUpClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ApplicationUpStorage_serviceDesc.Streams[0], "/ttn.lorawan.v3.ApplicationUpStorage/GetStoredApplicationUp", opts...)
	if err != nil {
		return nil, err
	}
	x := &applicationUpStorageGetStoredApplicationUpClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ApplicationUpStorage_GetStoredApplicationUpClient interface {
	Recv() (*ApplicationUp, error)
	grpc.ClientStream
}

type applicationUpStorageGetStoredApplicationUpClient struct {
	grpc.ClientStream
}

func (x *applicationUpStorageGetStoredApplicationUpClient) Recv() (*ApplicationUp, error) {
	m := new(ApplicationUp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *applicationUpStorageClient) GetStoredApplicationUpCount(ctx context.Context, in *GetStoredApplicationUpCountRequest, opts ...grpc.CallOption) (*GetStoredApplicationUpCountResponse, error) {
	out := new(GetStoredApplicationUpCountResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ApplicationUpStorage/GetStoredApplicationUpCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationUpStorageServer is the server API for ApplicationUpStorage service.
type ApplicationUpStorageServer interface {
	// Returns a stream of application messages that have been stored in the database.
	GetStoredApplicationUp(*GetStoredApplicationUpRequest, ApplicationUpStorage_GetStoredApplicationUpServer) error
	// Returns how many application messages have been stored in the database for an application or end device.
	GetStoredApplicationUpCount(context.Context, *GetStoredApplicationUpCountRequest) (*GetStoredApplicationUpCountResponse, error)
}

// UnimplementedApplicationUpStorageServer can be embedded to have forward compatible implementations.
type UnimplementedApplicationUpStorageServer struct {
}

func (*UnimplementedApplicationUpStorageServer) GetStoredApplicationUp(req *GetStoredApplicationUpRequest, srv ApplicationUpStorage_GetStoredApplicationUpServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStoredApplicationUp not implemented")
}
func (*UnimplementedApplicationUpStorageServer) GetStoredApplicationUpCount(ctx context.Context, req *GetStoredApplicationUpCountRequest) (*GetStoredApplicationUpCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStoredApplicationUpCount not implemented")
}

func RegisterApplicationUpStorageServer(s *grpc.Server, srv ApplicationUpStorageServer) {
	s.RegisterService(&_ApplicationUpStorage_serviceDesc, srv)
}

func _ApplicationUpStorage_GetStoredApplicationUp_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetStoredApplicationUpRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ApplicationUpStorageServer).GetStoredApplicationUp(m, &applicationUpStorageGetStoredApplicationUpServer{stream})
}

type ApplicationUpStorage_GetStoredApplicationUpServer interface {
	Send(*ApplicationUp) error
	grpc.ServerStream
}

type applicationUpStorageGetStoredApplicationUpServer struct {
	grpc.ServerStream
}

func (x *applicationUpStorageGetStoredApplicationUpServer) Send(m *ApplicationUp) error {
	return x.ServerStream.SendMsg(m)
}

func _ApplicationUpStorage_GetStoredApplicationUpCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoredApplicationUpCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationUpStorageServer).GetStoredApplicationUpCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ApplicationUpStorage/GetStoredApplicationUpCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationUpStorageServer).GetStoredApplicationUpCount(ctx, req.(*GetStoredApplicationUpCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationUpStorage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.ApplicationUpStorage",
	HandlerType: (*ApplicationUpStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStoredApplicationUpCount",
			Handler:    _ApplicationUpStorage_GetStoredApplicationUpCount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStoredApplicationUp",
			Handler:       _ApplicationUpStorage_GetStoredApplicationUp_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "lorawan-stack/api/applicationserver_integrations_storage.proto",
}
