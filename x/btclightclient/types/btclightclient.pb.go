// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/btclightclient/v1/btclightclient.proto

package types

import (
	fmt "fmt"
	github_com_babylonchain_babylon_types "github.com/babylonchain/babylon/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// BTCHeaderInfo is a structure that contains all relevant information about a
// BTC header
//   - Full header bytes
//   - Header hash for easy retrieval
//   - Height of the header in the BTC chain
//   - Total work spent on the header. This is the sum of the work corresponding
//     to the header Bits field
//     and the total work of the header.
type BTCHeaderInfo struct {
	Header *github_com_babylonchain_babylon_types.BTCHeaderBytes     `protobuf:"bytes,1,opt,name=header,proto3,customtype=github.com/babylonchain/babylon/types.BTCHeaderBytes" json:"header,omitempty"`
	Hash   *github_com_babylonchain_babylon_types.BTCHeaderHashBytes `protobuf:"bytes,2,opt,name=hash,proto3,customtype=github.com/babylonchain/babylon/types.BTCHeaderHashBytes" json:"hash,omitempty"`
	Height uint64                                                    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Work   *github_com_cosmos_cosmos_sdk_types.Uint                  `protobuf:"bytes,4,opt,name=work,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"work,omitempty"`
}

func (m *BTCHeaderInfo) Reset()         { *m = BTCHeaderInfo{} }
func (m *BTCHeaderInfo) String() string { return proto.CompactTextString(m) }
func (*BTCHeaderInfo) ProtoMessage()    {}
func (*BTCHeaderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_84bf438d909b681d, []int{0}
}
func (m *BTCHeaderInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BTCHeaderInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BTCHeaderInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BTCHeaderInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BTCHeaderInfo.Merge(m, src)
}
func (m *BTCHeaderInfo) XXX_Size() int {
	return m.Size()
}
func (m *BTCHeaderInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BTCHeaderInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BTCHeaderInfo proto.InternalMessageInfo

func (m *BTCHeaderInfo) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*BTCHeaderInfo)(nil), "babylon.btclightclient.v1.BTCHeaderInfo")
}

func init() {
	proto.RegisterFile("babylon/btclightclient/v1/btclightclient.proto", fileDescriptor_84bf438d909b681d)
}

var fileDescriptor_84bf438d909b681d = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0x4a, 0x4c, 0xaa,
	0xcc, 0xc9, 0xcf, 0xd3, 0x4f, 0x2a, 0x49, 0xce, 0xc9, 0x4c, 0xcf, 0x00, 0x91, 0xa9, 0x79, 0x25,
	0xfa, 0x65, 0x86, 0x68, 0x22, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x92, 0x50, 0xf5, 0x7a,
	0x68, 0xb2, 0x65, 0x86, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x55, 0xfa, 0x20, 0x16, 0x44,
	0x83, 0x52, 0x0f, 0x13, 0x17, 0xaf, 0x53, 0x88, 0xb3, 0x47, 0x6a, 0x62, 0x4a, 0x6a, 0x91, 0x67,
	0x5e, 0x5a, 0xbe, 0x50, 0x00, 0x17, 0x5b, 0x06, 0x98, 0x27, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe3,
	0x64, 0x71, 0xeb, 0x9e, 0xbc, 0x49, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae,
	0x3e, 0xd4, 0x86, 0xe4, 0x8c, 0xc4, 0xcc, 0x3c, 0x18, 0x47, 0xbf, 0xa4, 0xb2, 0x20, 0xb5, 0x58,
	0x0f, 0x6e, 0x90, 0x53, 0x65, 0x49, 0x6a, 0x71, 0x10, 0xd4, 0x1c, 0xa1, 0x00, 0x2e, 0x96, 0x8c,
	0xc4, 0xe2, 0x0c, 0x09, 0x26, 0xb0, 0x79, 0x36, 0xb7, 0xee, 0xc9, 0x5b, 0x90, 0x68, 0x9e, 0x47,
	0x62, 0x71, 0x06, 0xc4, 0x4c, 0xb0, 0x49, 0x42, 0x62, 0x20, 0x37, 0x82, 0xbc, 0x27, 0xc1, 0xac,
	0xc0, 0xa8, 0xc1, 0x12, 0x04, 0xe5, 0x09, 0xd9, 0x73, 0xb1, 0x94, 0xe7, 0x17, 0x65, 0x4b, 0xb0,
	0x80, 0x6d, 0xd2, 0xbe, 0x75, 0x4f, 0x5e, 0x1d, 0xc9, 0xa6, 0xe4, 0xfc, 0xe2, 0xdc, 0xfc, 0x62,
	0x28, 0xa5, 0x5b, 0x9c, 0x92, 0x0d, 0xb5, 0x26, 0x34, 0x33, 0xaf, 0x24, 0x08, 0xac, 0xd1, 0x29,
	0xe0, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58,
	0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xcc, 0x08, 0x39, 0xb9, 0x02,
	0x3d, 0x8e, 0xc0, 0x86, 0x27, 0xb1, 0x81, 0xc3, 0xd9, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x8a,
	0x8d, 0x9c, 0x25, 0xca, 0x01, 0x00, 0x00,
}

func (m *BTCHeaderInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BTCHeaderInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BTCHeaderInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Work != nil {
		{
			size := m.Work.Size()
			i -= size
			if _, err := m.Work.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintBtclightclient(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Height != 0 {
		i = encodeVarintBtclightclient(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	if m.Hash != nil {
		{
			size := m.Hash.Size()
			i -= size
			if _, err := m.Hash.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintBtclightclient(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		{
			size := m.Header.Size()
			i -= size
			if _, err := m.Header.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintBtclightclient(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBtclightclient(dAtA []byte, offset int, v uint64) int {
	offset -= sovBtclightclient(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BTCHeaderInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovBtclightclient(uint64(l))
	}
	if m.Hash != nil {
		l = m.Hash.Size()
		n += 1 + l + sovBtclightclient(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovBtclightclient(uint64(m.Height))
	}
	if m.Work != nil {
		l = m.Work.Size()
		n += 1 + l + sovBtclightclient(uint64(l))
	}
	return n
}

func sovBtclightclient(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBtclightclient(x uint64) (n int) {
	return sovBtclightclient(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BTCHeaderInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBtclightclient
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
			return fmt.Errorf("proto: BTCHeaderInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BTCHeaderInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBtclightclient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBtclightclient
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBtclightclient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_types.BTCHeaderBytes
			m.Header = &v
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBtclightclient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBtclightclient
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBtclightclient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_types.BTCHeaderHashBytes
			m.Hash = &v
			if err := m.Hash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBtclightclient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Work", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBtclightclient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBtclightclient
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBtclightclient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Uint
			m.Work = &v
			if err := m.Work.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBtclightclient(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBtclightclient
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
func skipBtclightclient(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBtclightclient
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
					return 0, ErrIntOverflowBtclightclient
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
					return 0, ErrIntOverflowBtclightclient
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
				return 0, ErrInvalidLengthBtclightclient
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBtclightclient
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBtclightclient
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBtclightclient        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBtclightclient          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBtclightclient = fmt.Errorf("proto: unexpected end of group")
)
