// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/m3db/m3db/src/dbnode/generated/proto/pagetoken/pagetoken.proto

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

/*
	Package pagetoken is a generated protocol buffer package.

	It is generated from these files:
		github.com/m3db/m3db/src/dbnode/generated/proto/pagetoken/pagetoken.proto

	It has these top-level messages:
		PageToken
*/
package pagetoken

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PageToken struct {
	ActiveSeriesPhase  *PageToken_ActiveSeriesPhase  `protobuf:"bytes,1,opt,name=active_series_phase,json=activeSeriesPhase" json:"active_series_phase,omitempty"`
	FlushedSeriesPhase *PageToken_FlushedSeriesPhase `protobuf:"bytes,2,opt,name=flushed_series_phase,json=flushedSeriesPhase" json:"flushed_series_phase,omitempty"`
}

func (m *PageToken) Reset()                    { *m = PageToken{} }
func (m *PageToken) String() string            { return proto.CompactTextString(m) }
func (*PageToken) ProtoMessage()               {}
func (*PageToken) Descriptor() ([]byte, []int) { return fileDescriptorPagetoken, []int{0} }

func (m *PageToken) GetActiveSeriesPhase() *PageToken_ActiveSeriesPhase {
	if m != nil {
		return m.ActiveSeriesPhase
	}
	return nil
}

func (m *PageToken) GetFlushedSeriesPhase() *PageToken_FlushedSeriesPhase {
	if m != nil {
		return m.FlushedSeriesPhase
	}
	return nil
}

type PageToken_ActiveSeriesPhase struct {
	IndexCursor int64 `protobuf:"varint,1,opt,name=indexCursor,proto3" json:"indexCursor,omitempty"`
}

func (m *PageToken_ActiveSeriesPhase) Reset()         { *m = PageToken_ActiveSeriesPhase{} }
func (m *PageToken_ActiveSeriesPhase) String() string { return proto.CompactTextString(m) }
func (*PageToken_ActiveSeriesPhase) ProtoMessage()    {}
func (*PageToken_ActiveSeriesPhase) Descriptor() ([]byte, []int) {
	return fileDescriptorPagetoken, []int{0, 0}
}

func (m *PageToken_ActiveSeriesPhase) GetIndexCursor() int64 {
	if m != nil {
		return m.IndexCursor
	}
	return 0
}

type PageToken_FlushedSeriesPhase struct {
	CurrBlockStartUnixNanos int64 `protobuf:"varint,1,opt,name=currBlockStartUnixNanos,proto3" json:"currBlockStartUnixNanos,omitempty"`
	CurrBlockEntryIdx       int64 `protobuf:"varint,2,opt,name=currBlockEntryIdx,proto3" json:"currBlockEntryIdx,omitempty"`
}

func (m *PageToken_FlushedSeriesPhase) Reset()         { *m = PageToken_FlushedSeriesPhase{} }
func (m *PageToken_FlushedSeriesPhase) String() string { return proto.CompactTextString(m) }
func (*PageToken_FlushedSeriesPhase) ProtoMessage()    {}
func (*PageToken_FlushedSeriesPhase) Descriptor() ([]byte, []int) {
	return fileDescriptorPagetoken, []int{0, 1}
}

func (m *PageToken_FlushedSeriesPhase) GetCurrBlockStartUnixNanos() int64 {
	if m != nil {
		return m.CurrBlockStartUnixNanos
	}
	return 0
}

func (m *PageToken_FlushedSeriesPhase) GetCurrBlockEntryIdx() int64 {
	if m != nil {
		return m.CurrBlockEntryIdx
	}
	return 0
}

func init() {
	proto.RegisterType((*PageToken)(nil), "pagetoken.PageToken")
	proto.RegisterType((*PageToken_ActiveSeriesPhase)(nil), "pagetoken.PageToken.ActiveSeriesPhase")
	proto.RegisterType((*PageToken_FlushedSeriesPhase)(nil), "pagetoken.PageToken.FlushedSeriesPhase")
}
func (m *PageToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PageToken) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ActiveSeriesPhase != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPagetoken(dAtA, i, uint64(m.ActiveSeriesPhase.Size()))
		n1, err := m.ActiveSeriesPhase.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.FlushedSeriesPhase != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPagetoken(dAtA, i, uint64(m.FlushedSeriesPhase.Size()))
		n2, err := m.FlushedSeriesPhase.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *PageToken_ActiveSeriesPhase) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PageToken_ActiveSeriesPhase) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.IndexCursor != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPagetoken(dAtA, i, uint64(m.IndexCursor))
	}
	return i, nil
}

func (m *PageToken_FlushedSeriesPhase) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PageToken_FlushedSeriesPhase) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.CurrBlockStartUnixNanos != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPagetoken(dAtA, i, uint64(m.CurrBlockStartUnixNanos))
	}
	if m.CurrBlockEntryIdx != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintPagetoken(dAtA, i, uint64(m.CurrBlockEntryIdx))
	}
	return i, nil
}

func encodeVarintPagetoken(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PageToken) Size() (n int) {
	var l int
	_ = l
	if m.ActiveSeriesPhase != nil {
		l = m.ActiveSeriesPhase.Size()
		n += 1 + l + sovPagetoken(uint64(l))
	}
	if m.FlushedSeriesPhase != nil {
		l = m.FlushedSeriesPhase.Size()
		n += 1 + l + sovPagetoken(uint64(l))
	}
	return n
}

func (m *PageToken_ActiveSeriesPhase) Size() (n int) {
	var l int
	_ = l
	if m.IndexCursor != 0 {
		n += 1 + sovPagetoken(uint64(m.IndexCursor))
	}
	return n
}

func (m *PageToken_FlushedSeriesPhase) Size() (n int) {
	var l int
	_ = l
	if m.CurrBlockStartUnixNanos != 0 {
		n += 1 + sovPagetoken(uint64(m.CurrBlockStartUnixNanos))
	}
	if m.CurrBlockEntryIdx != 0 {
		n += 1 + sovPagetoken(uint64(m.CurrBlockEntryIdx))
	}
	return n
}

func sovPagetoken(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPagetoken(x uint64) (n int) {
	return sovPagetoken(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PageToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPagetoken
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PageToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PageToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveSeriesPhase", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagetoken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPagetoken
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ActiveSeriesPhase == nil {
				m.ActiveSeriesPhase = &PageToken_ActiveSeriesPhase{}
			}
			if err := m.ActiveSeriesPhase.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FlushedSeriesPhase", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagetoken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPagetoken
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FlushedSeriesPhase == nil {
				m.FlushedSeriesPhase = &PageToken_FlushedSeriesPhase{}
			}
			if err := m.FlushedSeriesPhase.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPagetoken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPagetoken
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
func (m *PageToken_ActiveSeriesPhase) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPagetoken
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ActiveSeriesPhase: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActiveSeriesPhase: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexCursor", wireType)
			}
			m.IndexCursor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagetoken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IndexCursor |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPagetoken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPagetoken
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
func (m *PageToken_FlushedSeriesPhase) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPagetoken
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FlushedSeriesPhase: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FlushedSeriesPhase: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrBlockStartUnixNanos", wireType)
			}
			m.CurrBlockStartUnixNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagetoken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrBlockStartUnixNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrBlockEntryIdx", wireType)
			}
			m.CurrBlockEntryIdx = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPagetoken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrBlockEntryIdx |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPagetoken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPagetoken
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
func skipPagetoken(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPagetoken
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
					return 0, ErrIntOverflowPagetoken
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPagetoken
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPagetoken
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPagetoken
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPagetoken(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPagetoken = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPagetoken   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/m3db/m3db/src/dbnode/generated/proto/pagetoken/pagetoken.proto", fileDescriptorPagetoken)
}

var fileDescriptorPagetoken = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x35, 0x4e, 0x49, 0x82, 0x10, 0xc5, 0x45, 0xc9,
	0xfa, 0x29, 0x49, 0x79, 0xf9, 0x29, 0xa9, 0xfa, 0xe9, 0xa9, 0x79, 0xa9, 0x45, 0x89, 0x25, 0xa9,
	0x29, 0xfa, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0xfa, 0x05, 0x89, 0xe9, 0xa9, 0x25, 0xf9, 0xd9, 0xa9,
	0x79, 0x08, 0x96, 0x1e, 0x58, 0x46, 0x88, 0x13, 0x2e, 0xa0, 0xf4, 0x99, 0x89, 0x8b, 0x33, 0x20,
	0x31, 0x3d, 0x35, 0x04, 0xc4, 0x13, 0x0a, 0xe3, 0x12, 0x4e, 0x4c, 0x2e, 0xc9, 0x2c, 0x4b, 0x8d,
	0x2f, 0x4e, 0x2d, 0xca, 0x4c, 0x2d, 0x8e, 0x2f, 0xc8, 0x48, 0x2c, 0x4e, 0x95, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x36, 0x52, 0xd3, 0x43, 0x98, 0x03, 0xd7, 0xa2, 0xe7, 0x08, 0x56, 0x1f, 0x0c, 0x56,
	0x1e, 0x00, 0x52, 0x1d, 0x24, 0x98, 0x88, 0x2e, 0x24, 0x14, 0xc9, 0x25, 0x92, 0x96, 0x53, 0x5a,
	0x9c, 0x91, 0x9a, 0x82, 0x6a, 0x30, 0x13, 0xd8, 0x60, 0x75, 0xac, 0x06, 0xbb, 0x41, 0x34, 0x20,
	0x9b, 0x2c, 0x94, 0x86, 0x21, 0x26, 0x65, 0xca, 0x25, 0x88, 0xe1, 0x04, 0x21, 0x05, 0x2e, 0xee,
	0xcc, 0xbc, 0x94, 0xd4, 0x0a, 0xe7, 0xd2, 0xa2, 0xe2, 0xfc, 0x22, 0xb0, 0xfb, 0x99, 0x83, 0x90,
	0x85, 0xa4, 0x6a, 0xb8, 0x84, 0x30, 0x2d, 0x10, 0xb2, 0xe0, 0x12, 0x4f, 0x2e, 0x2d, 0x2a, 0x72,
	0xca, 0xc9, 0x4f, 0xce, 0x0e, 0x2e, 0x49, 0x2c, 0x2a, 0x09, 0xcd, 0xcb, 0xac, 0xf0, 0x4b, 0xcc,
	0xcb, 0x2f, 0x86, 0x9a, 0x81, 0x4b, 0x5a, 0x48, 0x87, 0x4b, 0x10, 0x2e, 0xe5, 0x9a, 0x57, 0x52,
	0x54, 0xe9, 0x99, 0x52, 0x01, 0xf6, 0x1e, 0x73, 0x10, 0xa6, 0x84, 0x93, 0xc0, 0x89, 0x47, 0x72,
	0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0x43, 0x12, 0x1b,
	0x38, 0x66, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xf7, 0x73, 0xf3, 0xe6, 0x01, 0x00,
	0x00,
}