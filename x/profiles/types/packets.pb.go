// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta1/packets.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// LinkChainAccountPacketData defines a struct for the packet payload
type LinkChainAccountPacketData struct {
	SourceAddress      string      `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty" yaml:"source_address"`
	SourceProof        Proof       `protobuf:"bytes,2,opt,name=source_proof,json=sourceProof,proto3" json:"source_proof" yaml:"source_proof"`
	SourceChainConfig  ChainConfig `protobuf:"bytes,3,opt,name=source_chain_config,json=sourceChainConfig,proto3" json:"source_chain_config" yaml:"source_chain_config"`
	DestinationAddress string      `protobuf:"bytes,4,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty" yaml:"destination_address"`
	DestinationProof   Proof       `protobuf:"bytes,5,opt,name=destination_proof,json=destinationProof,proto3" json:"destination_proof" yaml:"destination_proof"`
}

func (m *LinkChainAccountPacketData) Reset()         { *m = LinkChainAccountPacketData{} }
func (m *LinkChainAccountPacketData) String() string { return proto.CompactTextString(m) }
func (*LinkChainAccountPacketData) ProtoMessage()    {}
func (*LinkChainAccountPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0ff9a296e50f9ca, []int{0}
}
func (m *LinkChainAccountPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LinkChainAccountPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LinkChainAccountPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LinkChainAccountPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkChainAccountPacketData.Merge(m, src)
}
func (m *LinkChainAccountPacketData) XXX_Size() int {
	return m.Size()
}
func (m *LinkChainAccountPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkChainAccountPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_LinkChainAccountPacketData proto.InternalMessageInfo

func (m *LinkChainAccountPacketData) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *LinkChainAccountPacketData) GetSourceProof() Proof {
	if m != nil {
		return m.SourceProof
	}
	return Proof{}
}

func (m *LinkChainAccountPacketData) GetSourceChainConfig() ChainConfig {
	if m != nil {
		return m.SourceChainConfig
	}
	return ChainConfig{}
}

func (m *LinkChainAccountPacketData) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *LinkChainAccountPacketData) GetDestinationProof() Proof {
	if m != nil {
		return m.DestinationProof
	}
	return Proof{}
}

// LinkChainAccountPacketAck defines a struct for the packet acknowledgment
type LinkChainAccountPacketAck struct {
	SourceAddress string `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
}

func (m *LinkChainAccountPacketAck) Reset()         { *m = LinkChainAccountPacketAck{} }
func (m *LinkChainAccountPacketAck) String() string { return proto.CompactTextString(m) }
func (*LinkChainAccountPacketAck) ProtoMessage()    {}
func (*LinkChainAccountPacketAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0ff9a296e50f9ca, []int{1}
}
func (m *LinkChainAccountPacketAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LinkChainAccountPacketAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LinkChainAccountPacketAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LinkChainAccountPacketAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkChainAccountPacketAck.Merge(m, src)
}
func (m *LinkChainAccountPacketAck) XXX_Size() int {
	return m.Size()
}
func (m *LinkChainAccountPacketAck) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkChainAccountPacketAck.DiscardUnknown(m)
}

var xxx_messageInfo_LinkChainAccountPacketAck proto.InternalMessageInfo

func (m *LinkChainAccountPacketAck) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*LinkChainAccountPacketData)(nil), "desmos.profiles.v1beta1.LinkChainAccountPacketData")
	proto.RegisterType((*LinkChainAccountPacketAck)(nil), "desmos.profiles.v1beta1.LinkChainAccountPacketAck")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta1/packets.proto", fileDescriptor_e0ff9a296e50f9ca)
}

var fileDescriptor_e0ff9a296e50f9ca = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0xef, 0xd2, 0x30,
	0x18, 0xc6, 0x37, 0x51, 0x13, 0x8b, 0x1a, 0x19, 0x1a, 0xc7, 0x4c, 0x0a, 0x69, 0x20, 0xe1, 0xe2,
	0x16, 0xf4, 0xe6, 0x49, 0x86, 0x07, 0x0f, 0x26, 0x92, 0x1d, 0x3d, 0x48, 0xba, 0xae, 0x8c, 0x85,
	0x6d, 0x5d, 0xd6, 0x62, 0xe0, 0x5b, 0xf8, 0xb1, 0x38, 0x72, 0xf4, 0x44, 0x0c, 0x7c, 0x03, 0x4e,
	0x1e, 0x0d, 0x6d, 0x91, 0xe1, 0x1f, 0x92, 0xff, 0xad, 0x7d, 0xfb, 0xeb, 0xf3, 0xf4, 0x7d, 0xde,
	0x82, 0x5e, 0x44, 0x79, 0xc6, 0xb8, 0x57, 0x94, 0x6c, 0x9a, 0xa4, 0x94, 0x7b, 0x3f, 0x06, 0x21,
	0x15, 0x78, 0xe0, 0x15, 0x98, 0xcc, 0xa9, 0xe0, 0x6e, 0x51, 0x32, 0xc1, 0xac, 0xd7, 0x0a, 0x73,
	0x4f, 0x98, 0xab, 0x31, 0xe7, 0x65, 0xcc, 0x62, 0x26, 0x19, 0xef, 0xb8, 0x52, 0xb8, 0xd3, 0xbd,
	0xa5, 0x9a, 0xb1, 0x88, 0xa6, 0x5a, 0x14, 0xfd, 0xa9, 0x01, 0xe7, 0x4b, 0x92, 0xcf, 0x47, 0x33,
	0x9c, 0xe4, 0x43, 0x42, 0xd8, 0x22, 0x17, 0x63, 0x69, 0xfb, 0x09, 0x0b, 0x6c, 0x7d, 0x04, 0xcf,
	0x39, 0x5b, 0x94, 0x84, 0x4e, 0x70, 0x14, 0x95, 0x94, 0x73, 0xdb, 0xec, 0x98, 0xfd, 0x27, 0x7e,
	0xeb, 0xb0, 0x6d, 0xbf, 0x5a, 0xe1, 0x2c, 0xfd, 0x80, 0x2e, 0xcf, 0x51, 0xf0, 0x4c, 0x15, 0x86,
	0x6a, 0x6f, 0x7d, 0x07, 0x4f, 0x35, 0x51, 0x94, 0x8c, 0x4d, 0xed, 0x07, 0x1d, 0xb3, 0x5f, 0x7f,
	0x07, 0xdd, 0x1b, 0xcd, 0xb8, 0xe3, 0x23, 0xe5, 0xbf, 0x59, 0x6f, 0xdb, 0xc6, 0x61, 0xdb, 0x6e,
	0x5e, 0x78, 0x48, 0x05, 0x14, 0xd4, 0xd5, 0x56, 0x92, 0xd6, 0x12, 0x34, 0xf5, 0x29, 0x39, 0xb6,
	0x30, 0x21, 0x2c, 0x9f, 0x26, 0xb1, 0x5d, 0x93, 0x36, 0xdd, 0x9b, 0x36, 0xb2, 0xdf, 0x91, 0x64,
	0x7d, 0xa4, 0xcd, 0x9c, 0x0b, 0xb3, 0xaa, 0x1c, 0x0a, 0x1a, 0xaa, 0x5a, 0xb9, 0x66, 0x7d, 0x05,
	0xcd, 0x88, 0x72, 0x91, 0xe4, 0x58, 0x24, 0x2c, 0xff, 0x17, 0xd0, 0x43, 0x19, 0x10, 0x3c, 0xeb,
	0x5d, 0x81, 0x50, 0x60, 0x55, 0xaa, 0xa7, 0xa8, 0x32, 0xd0, 0xa8, 0xb2, 0x2a, 0xaf, 0x47, 0xf7,
	0xca, 0xab, 0xa3, 0x5b, 0xb0, 0xef, 0x5a, 0xea, 0xd0, 0x5e, 0x54, 0x6a, 0xf2, 0x0e, 0xf2, 0x41,
	0xeb, 0xfa, 0xe4, 0x87, 0x64, 0x6e, 0xf5, 0xae, 0x0f, 0xfe, 0xbf, 0xe9, 0xfa, 0x9f, 0xd7, 0x3b,
	0x68, 0x6e, 0x76, 0xd0, 0xfc, 0xbd, 0x83, 0xe6, 0xcf, 0x3d, 0x34, 0x36, 0x7b, 0x68, 0xfc, 0xda,
	0x43, 0xe3, 0x9b, 0x1b, 0x27, 0x62, 0xb6, 0x08, 0x5d, 0xc2, 0x32, 0x4f, 0xbd, 0xfd, 0x6d, 0x8a,
	0x43, 0xae, 0xd7, 0xde, 0xf2, 0xfc, 0x2f, 0xc5, 0xaa, 0xa0, 0x3c, 0x7c, 0x2c, 0xff, 0xe3, 0xfb,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x08, 0x55, 0x46, 0x74, 0x0d, 0x03, 0x00, 0x00,
}

func (m *LinkChainAccountPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LinkChainAccountPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LinkChainAccountPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.DestinationProof.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPackets(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.DestinationAddress) > 0 {
		i -= len(m.DestinationAddress)
		copy(dAtA[i:], m.DestinationAddress)
		i = encodeVarintPackets(dAtA, i, uint64(len(m.DestinationAddress)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.SourceChainConfig.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPackets(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.SourceProof.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPackets(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.SourceAddress) > 0 {
		i -= len(m.SourceAddress)
		copy(dAtA[i:], m.SourceAddress)
		i = encodeVarintPackets(dAtA, i, uint64(len(m.SourceAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LinkChainAccountPacketAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LinkChainAccountPacketAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LinkChainAccountPacketAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *LinkChainAccountPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourceAddress)
	if l > 0 {
		n += 1 + l + sovPackets(uint64(l))
	}
	l = m.SourceProof.Size()
	n += 1 + l + sovPackets(uint64(l))
	l = m.SourceChainConfig.Size()
	n += 1 + l + sovPackets(uint64(l))
	l = len(m.DestinationAddress)
	if l > 0 {
		n += 1 + l + sovPackets(uint64(l))
	}
	l = m.DestinationProof.Size()
	n += 1 + l + sovPackets(uint64(l))
	return n
}

func (m *LinkChainAccountPacketAck) Size() (n int) {
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
func (m *LinkChainAccountPacketData) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: LinkChainAccountPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LinkChainAccountPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceProof", wireType)
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
			if err := m.SourceProof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChainConfig", wireType)
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
			if err := m.SourceChainConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationProof", wireType)
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
			if err := m.DestinationProof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *LinkChainAccountPacketAck) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: LinkChainAccountPacketAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LinkChainAccountPacketAck: illegal tag %d (wire type %d)", fieldNum, wire)
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