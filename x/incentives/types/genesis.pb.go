// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/incentives/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the incentives module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module
	Params            Params          `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Gauges            []Gauge         `protobuf:"bytes,2,rep,name=gauges,proto3" json:"gauges"`
	LockableDurations []time.Duration `protobuf:"bytes,3,rep,name=lockable_durations,json=lockableDurations,proto3,stdduration" json:"lockable_durations" yaml:"lockable_durations"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a288ccc95d977d2d, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetGauges() []Gauge {
	if m != nil {
		return m.Gauges
	}
	return nil
}

func (m *GenesisState) GetLockableDurations() []time.Duration {
	if m != nil {
		return m.LockableDurations
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "osmosis.incentives.GenesisState")
}

func init() { proto.RegisterFile("osmosis/incentives/genesis.proto", fileDescriptor_a288ccc95d977d2d) }

var fileDescriptor_a288ccc95d977d2d = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0xe3, 0xdb, 0xab, 0x0e, 0x29, 0x0b, 0x16, 0x43, 0xdb, 0xc1, 0xa9, 0x2a, 0x21, 0x75,
	0xc1, 0x96, 0xca, 0x00, 0x62, 0xac, 0x90, 0xba, 0x30, 0xa0, 0xb2, 0xb1, 0x20, 0xa7, 0x18, 0x63,
	0x91, 0xe4, 0x54, 0x3d, 0x0e, 0xa2, 0x6f, 0xc1, 0xc8, 0x23, 0x75, 0xec, 0xc8, 0x54, 0x50, 0xfb,
	0x06, 0x7d, 0x02, 0xd4, 0xd8, 0x16, 0x48, 0xcd, 0x96, 0xa3, 0xff, 0x3b, 0x5f, 0x7e, 0x9f, 0xb8,
	0x07, 0x98, 0x03, 0x1a, 0x14, 0xa6, 0x98, 0xaa, 0xc2, 0x9a, 0x57, 0x85, 0x42, 0xab, 0x42, 0xa1,
	0x41, 0x3e, 0x9b, 0x83, 0x05, 0x4a, 0x3d, 0xc1, 0x7f, 0x89, 0xee, 0x89, 0x06, 0x0d, 0x55, 0x2c,
	0xf6, 0x5f, 0x8e, 0xec, 0x32, 0x0d, 0xa0, 0x33, 0x25, 0xaa, 0x29, 0x2d, 0x9f, 0xc4, 0x63, 0x39,
	0x97, 0xd6, 0x40, 0xe1, 0xf3, 0xa4, 0xe6, 0x5f, 0x33, 0x39, 0x97, 0x39, 0x06, 0x41, 0x5d, 0x19,
	0x59, 0x6a, 0xe5, 0xf2, 0xfe, 0x8e, 0xc4, 0x47, 0x63, 0x57, 0xee, 0xce, 0x4a, 0xab, 0xe8, 0x65,
	0xdc, 0x74, 0x82, 0x36, 0xe9, 0x91, 0x41, 0x6b, 0xd8, 0xe5, 0x87, 0x65, 0xf9, 0x6d, 0x45, 0x8c,
	0xfe, 0x2f, 0xd7, 0x49, 0x34, 0xf1, 0x3c, 0xbd, 0x88, 0x9b, 0x95, 0x19, 0xdb, 0xff, 0x7a, 0x8d,
	0x41, 0x6b, 0xd8, 0xa9, 0xdb, 0x1c, 0xef, 0x89, 0xb0, 0xe8, 0x70, 0x0a, 0x31, 0xcd, 0x60, 0xfa,
	0x22, 0xd3, 0x4c, 0x3d, 0x84, 0xf7, 0x61, 0xbb, 0xe1, 0x25, 0xee, 0x02, 0x3c, 0x5c, 0x80, 0x5f,
	0x7b, 0x62, 0x74, 0xba, 0x97, 0xec, 0xd6, 0x49, 0x67, 0x21, 0xf3, 0xec, 0xaa, 0x7f, 0xa8, 0xe8,
	0x7f, 0x7c, 0x25, 0x64, 0x72, 0x1c, 0x82, 0xb0, 0x88, 0xa3, 0x9b, 0xe5, 0x86, 0x91, 0xd5, 0x86,
	0x91, 0xef, 0x0d, 0x23, 0xef, 0x5b, 0x16, 0xad, 0xb6, 0x2c, 0xfa, 0xdc, 0xb2, 0xe8, 0x7e, 0xa8,
	0x8d, 0x7d, 0x2e, 0x53, 0x3e, 0x85, 0x5c, 0xf8, 0xf6, 0x67, 0x99, 0x4c, 0x31, 0x0c, 0xe2, 0xed,
	0xef, 0x21, 0xed, 0x62, 0xa6, 0x30, 0x6d, 0x56, 0xd5, 0xce, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x8b, 0x2e, 0x6a, 0x36, 0xf8, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LockableDurations) > 0 {
		for iNdEx := len(m.LockableDurations) - 1; iNdEx >= 0; iNdEx-- {
			n, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.LockableDurations[iNdEx], dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.LockableDurations[iNdEx]):])
			if err != nil {
				return 0, err
			}
			i -= n
			i = encodeVarintGenesis(dAtA, i, uint64(n))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Gauges) > 0 {
		for iNdEx := len(m.Gauges) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Gauges[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Gauges) > 0 {
		for _, e := range m.Gauges {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LockableDurations) > 0 {
		for _, e := range m.LockableDurations {
			l = github_com_gogo_protobuf_types.SizeOfStdDuration(e)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gauges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Gauges = append(m.Gauges, Gauge{})
			if err := m.Gauges[len(m.Gauges)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockableDurations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LockableDurations = append(m.LockableDurations, time.Duration(0))
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&(m.LockableDurations[len(m.LockableDurations)-1]), dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
