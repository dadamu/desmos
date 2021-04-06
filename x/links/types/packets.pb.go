// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/links/v1beta1/packets.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Declare types of packet data
type LinksPacketData struct {
	// Switch one of packet type
	//
	// Types that are valid to be assigned to Packet:
	//	*LinksPacketData_NoData
	//	*LinksPacketData_IbcAccountConnectionPacket
	Packet isLinksPacketData_Packet `protobuf_oneof:"packet"`
}

func (m *LinksPacketData) Reset()         { *m = LinksPacketData{} }
func (m *LinksPacketData) String() string { return proto.CompactTextString(m) }
func (*LinksPacketData) ProtoMessage()    {}
func (*LinksPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_70e1574985ba68b0, []int{0}
}
func (m *LinksPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LinksPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LinksPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LinksPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinksPacketData.Merge(m, src)
}
func (m *LinksPacketData) XXX_Size() int {
	return m.Size()
}
func (m *LinksPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_LinksPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_LinksPacketData proto.InternalMessageInfo

type isLinksPacketData_Packet interface {
	isLinksPacketData_Packet()
	MarshalTo([]byte) (int, error)
	Size() int
}

type LinksPacketData_NoData struct {
	NoData *NoData `protobuf:"bytes,1,opt,name=no_data,json=noData,proto3,oneof" json:"no_data,omitempty"`
}
type LinksPacketData_IbcAccountConnectionPacket struct {
	IbcAccountConnectionPacket *IBCAccountConnectionPacketData `protobuf:"bytes,2,opt,name=ibc_account_connection_packet,json=ibcAccountConnectionPacket,proto3,oneof" json:"ibc_account_connection_packet,omitempty"`
}

func (*LinksPacketData_NoData) isLinksPacketData_Packet()                     {}
func (*LinksPacketData_IbcAccountConnectionPacket) isLinksPacketData_Packet() {}

func (m *LinksPacketData) GetPacket() isLinksPacketData_Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *LinksPacketData) GetNoData() *NoData {
	if x, ok := m.GetPacket().(*LinksPacketData_NoData); ok {
		return x.NoData
	}
	return nil
}

func (m *LinksPacketData) GetIbcAccountConnectionPacket() *IBCAccountConnectionPacketData {
	if x, ok := m.GetPacket().(*LinksPacketData_IbcAccountConnectionPacket); ok {
		return x.IbcAccountConnectionPacket
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LinksPacketData) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LinksPacketData_NoData)(nil),
		(*LinksPacketData_IbcAccountConnectionPacket)(nil),
	}
}

// No data
type NoData struct {
}

func (m *NoData) Reset()         { *m = NoData{} }
func (m *NoData) String() string { return proto.CompactTextString(m) }
func (*NoData) ProtoMessage()    {}
func (*NoData) Descriptor() ([]byte, []int) {
	return fileDescriptor_70e1574985ba68b0, []int{1}
}
func (m *NoData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NoData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NoData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NoData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoData.Merge(m, src)
}
func (m *NoData) XXX_Size() int {
	return m.Size()
}
func (m *NoData) XXX_DiscardUnknown() {
	xxx_messageInfo_NoData.DiscardUnknown(m)
}

var xxx_messageInfo_NoData proto.InternalMessageInfo

// IBCAccountConnectionPacketData defines a struct for the packet payload
type IBCAccountConnectionPacketData struct {
	SourceChainPrefix    string `protobuf:"bytes,1,opt,name=source_chain_prefix,json=sourceChainPrefix,proto3" json:"source_chain_prefix,omitempty"`
	SourceAddress        string `protobuf:"bytes,2,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	SourcePubKey         string `protobuf:"bytes,3,opt,name=source_pub_key,json=sourcePubKey,proto3" json:"source_pub_key,omitempty"`
	DestinationAddress   string `protobuf:"bytes,4,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	SourceSignature      string `protobuf:"bytes,5,opt,name=source_signature,json=sourceSignature,proto3" json:"source_signature,omitempty"`
	DestinationSignature string `protobuf:"bytes,6,opt,name=destination_signature,json=destinationSignature,proto3" json:"destination_signature,omitempty"`
}

func (m *IBCAccountConnectionPacketData) Reset()         { *m = IBCAccountConnectionPacketData{} }
func (m *IBCAccountConnectionPacketData) String() string { return proto.CompactTextString(m) }
func (*IBCAccountConnectionPacketData) ProtoMessage()    {}
func (*IBCAccountConnectionPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_70e1574985ba68b0, []int{2}
}
func (m *IBCAccountConnectionPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IBCAccountConnectionPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IBCAccountConnectionPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IBCAccountConnectionPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IBCAccountConnectionPacketData.Merge(m, src)
}
func (m *IBCAccountConnectionPacketData) XXX_Size() int {
	return m.Size()
}
func (m *IBCAccountConnectionPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_IBCAccountConnectionPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_IBCAccountConnectionPacketData proto.InternalMessageInfo

func (m *IBCAccountConnectionPacketData) GetSourceChainPrefix() string {
	if m != nil {
		return m.SourceChainPrefix
	}
	return ""
}

func (m *IBCAccountConnectionPacketData) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *IBCAccountConnectionPacketData) GetSourcePubKey() string {
	if m != nil {
		return m.SourcePubKey
	}
	return ""
}

func (m *IBCAccountConnectionPacketData) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *IBCAccountConnectionPacketData) GetSourceSignature() string {
	if m != nil {
		return m.SourceSignature
	}
	return ""
}

func (m *IBCAccountConnectionPacketData) GetDestinationSignature() string {
	if m != nil {
		return m.DestinationSignature
	}
	return ""
}

// IBCAccountConnectionPacketAck defines a struct for the packet acknowledgment
type IBCAccountConnectionPacketAck struct {
	SourceAddress string `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
}

func (m *IBCAccountConnectionPacketAck) Reset()         { *m = IBCAccountConnectionPacketAck{} }
func (m *IBCAccountConnectionPacketAck) String() string { return proto.CompactTextString(m) }
func (*IBCAccountConnectionPacketAck) ProtoMessage()    {}
func (*IBCAccountConnectionPacketAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_70e1574985ba68b0, []int{3}
}
func (m *IBCAccountConnectionPacketAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IBCAccountConnectionPacketAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IBCAccountConnectionPacketAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IBCAccountConnectionPacketAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IBCAccountConnectionPacketAck.Merge(m, src)
}
func (m *IBCAccountConnectionPacketAck) XXX_Size() int {
	return m.Size()
}
func (m *IBCAccountConnectionPacketAck) XXX_DiscardUnknown() {
	xxx_messageInfo_IBCAccountConnectionPacketAck.DiscardUnknown(m)
}

var xxx_messageInfo_IBCAccountConnectionPacketAck proto.InternalMessageInfo

func (m *IBCAccountConnectionPacketAck) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*LinksPacketData)(nil), "desmos.links.v1beta1.LinksPacketData")
	proto.RegisterType((*NoData)(nil), "desmos.links.v1beta1.NoData")
	proto.RegisterType((*IBCAccountConnectionPacketData)(nil), "desmos.links.v1beta1.IBCAccountConnectionPacketData")
	proto.RegisterType((*IBCAccountConnectionPacketAck)(nil), "desmos.links.v1beta1.IBCAccountConnectionPacketAck")
}

func init() {
	proto.RegisterFile("desmos/links/v1beta1/packets.proto", fileDescriptor_70e1574985ba68b0)
}

var fileDescriptor_70e1574985ba68b0 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x93, 0xaa, 0xb1, 0x1d, 0xff, 0x54, 0xa7, 0x15, 0x82, 0xd8, 0x20, 0x41, 0x41, 0x11,
	0x13, 0x6a, 0x05, 0xd7, 0x6d, 0x55, 0x2a, 0x8a, 0x94, 0xb8, 0x73, 0x13, 0x66, 0x26, 0x63, 0x3b,
	0xa4, 0x9d, 0x09, 0x99, 0x89, 0x34, 0x6f, 0xe1, 0x93, 0xf8, 0x1c, 0x77, 0x77, 0xbb, 0xbc, 0xcb,
	0x4b, 0xfb, 0x22, 0x97, 0xce, 0xa4, 0x7f, 0x16, 0xe9, 0xdd, 0x0d, 0xdf, 0xf7, 0x3b, 0xdf, 0x9c,
	0x73, 0x38, 0xc0, 0x4f, 0xa8, 0x5c, 0x0a, 0x19, 0x2e, 0x18, 0x4f, 0x65, 0xf8, 0xb7, 0x8f, 0xa9,
	0x42, 0xfd, 0x30, 0x43, 0x24, 0xa5, 0x4a, 0x06, 0x59, 0x2e, 0x94, 0x80, 0x5d, 0xc3, 0x04, 0x9a,
	0x09, 0x2a, 0xc6, 0xbf, 0xb4, 0x41, 0xfb, 0xc7, 0x4e, 0x99, 0x6a, 0xf8, 0x33, 0x52, 0x08, 0x7e,
	0x02, 0xf7, 0xb9, 0x88, 0x13, 0xa4, 0x90, 0x6b, 0xbf, 0xb4, 0xdf, 0x3c, 0xf8, 0xf0, 0x22, 0xa8,
	0xab, 0x0d, 0x7e, 0x8a, 0x1d, 0x3e, 0xb1, 0x22, 0x87, 0xeb, 0x17, 0x2c, 0x41, 0x8f, 0x61, 0x12,
	0x23, 0x42, 0x44, 0xc1, 0x55, 0x4c, 0x04, 0xe7, 0x94, 0x28, 0x26, 0x78, 0x6c, 0x5a, 0x71, 0x1b,
	0x3a, 0xee, 0x63, 0x7d, 0xdc, 0xb7, 0xd1, 0x78, 0x68, 0x2a, 0xc7, 0x87, 0xc2, 0x63, 0x57, 0x13,
	0x2b, 0x7a, 0xce, 0x30, 0x39, 0x43, 0x8c, 0x9a, 0xc0, 0x31, 0x7f, 0xf8, 0x4d, 0xe0, 0x98, 0xc6,
	0xfc, 0xff, 0x0d, 0xe0, 0xdd, 0x1e, 0x0a, 0x03, 0xd0, 0x91, 0xa2, 0xc8, 0x09, 0x8d, 0xc9, 0x1c,
	0x31, 0x1e, 0x67, 0x39, 0xfd, 0xc3, 0x56, 0x7a, 0xec, 0x56, 0xf4, 0xd4, 0x58, 0xe3, 0x9d, 0x33,
	0xd5, 0x06, 0x7c, 0x0d, 0x1e, 0x57, 0x3c, 0x4a, 0x92, 0x9c, 0x4a, 0xa9, 0x47, 0x6a, 0x45, 0x8f,
	0x8c, 0x3a, 0x34, 0x22, 0x7c, 0x75, 0xc0, 0xb2, 0x02, 0xc7, 0x29, 0x2d, 0xdd, 0x3b, 0x1a, 0x7b,
	0x68, 0xd4, 0x69, 0x81, 0xbf, 0xd3, 0x12, 0x86, 0xa0, 0x93, 0x50, 0xa9, 0x18, 0x47, 0x7a, 0x47,
	0xfb, 0xc4, 0xbb, 0x1a, 0x85, 0x27, 0xd6, 0x3e, 0xf6, 0x2d, 0x78, 0x52, 0xc5, 0x4a, 0x36, 0xe3,
	0x48, 0x15, 0x39, 0x75, 0xef, 0x69, 0xba, 0x6d, 0xf4, 0x5f, 0x7b, 0x19, 0x0e, 0xc0, 0xb3, 0xd3,
	0xec, 0x23, 0xef, 0x68, 0xbe, 0x7b, 0x62, 0x1e, 0x8a, 0xfc, 0xaf, 0xa0, 0x77, 0x7e, 0x5f, 0x43,
	0x92, 0xd6, 0x8c, 0x6f, 0xd7, 0x8c, 0x3f, 0xfa, 0x72, 0xb1, 0xf1, 0xec, 0xf5, 0xc6, 0xb3, 0xaf,
	0x37, 0x9e, 0xfd, 0x6f, 0xeb, 0x59, 0xeb, 0xad, 0x67, 0x5d, 0x6d, 0x3d, 0xeb, 0xf7, 0xbb, 0x19,
	0x53, 0xf3, 0x02, 0x07, 0x44, 0x2c, 0x43, 0x73, 0x04, 0xef, 0x17, 0x08, 0xcb, 0xea, 0x1d, 0xae,
	0xaa, 0x0b, 0x56, 0x65, 0x46, 0x25, 0x76, 0xf4, 0xe1, 0x0e, 0x6e, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x8e, 0x1f, 0xc4, 0x5d, 0xde, 0x02, 0x00, 0x00,
}

func (m *LinksPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LinksPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LinksPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Packet != nil {
		{
			size := m.Packet.Size()
			i -= size
			if _, err := m.Packet.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *LinksPacketData_NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LinksPacketData_NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NoData != nil {
		{
			size, err := m.NoData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPackets(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *LinksPacketData_IbcAccountConnectionPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LinksPacketData_IbcAccountConnectionPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.IbcAccountConnectionPacket != nil {
		{
			size, err := m.IbcAccountConnectionPacket.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPackets(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *NoData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *IBCAccountConnectionPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IBCAccountConnectionPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IBCAccountConnectionPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DestinationSignature) > 0 {
		i -= len(m.DestinationSignature)
		copy(dAtA[i:], m.DestinationSignature)
		i = encodeVarintPackets(dAtA, i, uint64(len(m.DestinationSignature)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.SourceSignature) > 0 {
		i -= len(m.SourceSignature)
		copy(dAtA[i:], m.SourceSignature)
		i = encodeVarintPackets(dAtA, i, uint64(len(m.SourceSignature)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DestinationAddress) > 0 {
		i -= len(m.DestinationAddress)
		copy(dAtA[i:], m.DestinationAddress)
		i = encodeVarintPackets(dAtA, i, uint64(len(m.DestinationAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SourcePubKey) > 0 {
		i -= len(m.SourcePubKey)
		copy(dAtA[i:], m.SourcePubKey)
		i = encodeVarintPackets(dAtA, i, uint64(len(m.SourcePubKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SourceAddress) > 0 {
		i -= len(m.SourceAddress)
		copy(dAtA[i:], m.SourceAddress)
		i = encodeVarintPackets(dAtA, i, uint64(len(m.SourceAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SourceChainPrefix) > 0 {
		i -= len(m.SourceChainPrefix)
		copy(dAtA[i:], m.SourceChainPrefix)
		i = encodeVarintPackets(dAtA, i, uint64(len(m.SourceChainPrefix)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IBCAccountConnectionPacketAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IBCAccountConnectionPacketAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IBCAccountConnectionPacketAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SourceAddress) > 0 {
		i -= len(m.SourceAddress)
		copy(dAtA[i:], m.SourceAddress)
		i = encodeVarintPackets(dAtA, i, uint64(len(m.SourceAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPackets(dAtA []byte, offset int, v uint64) int {
	offset -= sovPackets(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LinksPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Packet != nil {
		n += m.Packet.Size()
	}
	return n
}

func (m *LinksPacketData_NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NoData != nil {
		l = m.NoData.Size()
		n += 1 + l + sovPackets(uint64(l))
	}
	return n
}
func (m *LinksPacketData_IbcAccountConnectionPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IbcAccountConnectionPacket != nil {
		l = m.IbcAccountConnectionPacket.Size()
		n += 1 + l + sovPackets(uint64(l))
	}
	return n
}
func (m *NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *IBCAccountConnectionPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourceChainPrefix)
	if l > 0 {
		n += 1 + l + sovPackets(uint64(l))
	}
	l = len(m.SourceAddress)
	if l > 0 {
		n += 1 + l + sovPackets(uint64(l))
	}
	l = len(m.SourcePubKey)
	if l > 0 {
		n += 1 + l + sovPackets(uint64(l))
	}
	l = len(m.DestinationAddress)
	if l > 0 {
		n += 1 + l + sovPackets(uint64(l))
	}
	l = len(m.SourceSignature)
	if l > 0 {
		n += 1 + l + sovPackets(uint64(l))
	}
	l = len(m.DestinationSignature)
	if l > 0 {
		n += 1 + l + sovPackets(uint64(l))
	}
	return n
}

func (m *IBCAccountConnectionPacketAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourceAddress)
	if l > 0 {
		n += 1 + l + sovPackets(uint64(l))
	}
	return n
}

func sovPackets(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPackets(x uint64) (n int) {
	return sovPackets(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LinksPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPackets
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LinksPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LinksPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPackets
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &NoData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &LinksPacketData_NoData{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcAccountConnectionPacket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPackets
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &IBCAccountConnectionPacketData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &LinksPacketData_IbcAccountConnectionPacket{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPackets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPackets
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NoData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPackets
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NoData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NoData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPackets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPackets
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IBCAccountConnectionPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPackets
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IBCAccountConnectionPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IBCAccountConnectionPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChainPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPackets
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChainPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPackets
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPackets
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourcePubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPackets
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceSignature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPackets
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceSignature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationSignature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPackets
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationSignature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPackets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPackets
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IBCAccountConnectionPacketAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPackets
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IBCAccountConnectionPacketAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IBCAccountConnectionPacketAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPackets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPackets
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPackets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPackets
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPackets(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPackets
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPackets
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPackets
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPackets
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPackets
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPackets
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPackets        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPackets          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPackets = fmt.Errorf("proto: unexpected end of group")
)
