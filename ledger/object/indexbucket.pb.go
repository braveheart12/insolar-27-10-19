// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ledger/object/indexbucket.proto

package object

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_insolar_insolar_insolar "github.com/insolar/insolar/insolar"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type IndexBucket struct {
	ObjID            github_com_insolar_insolar_insolar.ID          `protobuf:"bytes,1,opt,name=ObjID,proto3,customtype=github.com/insolar/insolar/insolar.ID" json:"ObjID"`
	Lifeline         *Lifeline                                      `protobuf:"bytes,2,opt,name=Lifeline,proto3" json:"Lifeline,omitempty"`
	LifelineLastUsed github_com_insolar_insolar_insolar.PulseNumber `protobuf:"bytes,3,opt,name=LifelineLastUsed,proto3,customtype=github.com/insolar/insolar/insolar.PulseNumber" json:"LifelineLastUsed"`
	Requests         []github_com_insolar_insolar_insolar.ID        `protobuf:"bytes,4,rep,name=Requests,proto3,customtype=github.com/insolar/insolar/insolar.ID" json:"Requests,omitempty"`
	Results          []github_com_insolar_insolar_insolar.ID        `protobuf:"bytes,5,rep,name=Results,proto3,customtype=github.com/insolar/insolar/insolar.ID" json:"Results,omitempty"`
}

func (m *IndexBucket) Reset()      { *m = IndexBucket{} }
func (*IndexBucket) ProtoMessage() {}
func (*IndexBucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_82c40bb7e64b245d, []int{0}
}
func (m *IndexBucket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IndexBucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IndexBucket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IndexBucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexBucket.Merge(m, src)
}
func (m *IndexBucket) XXX_Size() int {
	return m.Size()
}
func (m *IndexBucket) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexBucket.DiscardUnknown(m)
}

var xxx_messageInfo_IndexBucket proto.InternalMessageInfo

func (m *IndexBucket) GetLifeline() *Lifeline {
	if m != nil {
		return m.Lifeline
	}
	return nil
}

func init() {
	proto.RegisterType((*IndexBucket)(nil), "object.IndexBucket")
}

func init() { proto.RegisterFile("ledger/object/indexbucket.proto", fileDescriptor_82c40bb7e64b245d) }

var fileDescriptor_82c40bb7e64b245d = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcf, 0x49, 0x4d, 0x49,
	0x4f, 0x2d, 0xd2, 0xcf, 0x4f, 0xca, 0x4a, 0x4d, 0x2e, 0xd1, 0xcf, 0xcc, 0x4b, 0x49, 0xad, 0x48,
	0x2a, 0x4d, 0xce, 0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xc8, 0x48,
	0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7,
	0xeb, 0x83, 0xa5, 0x93, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0x68, 0x93, 0x92, 0x41,
	0x35, 0x37, 0x27, 0x33, 0x2d, 0x35, 0x27, 0x33, 0x2f, 0x15, 0x22, 0xab, 0xd4, 0xc2, 0xcc, 0xc5,
	0xed, 0x09, 0xb2, 0xca, 0x09, 0x6c, 0x95, 0x90, 0x33, 0x17, 0xab, 0x7f, 0x52, 0x96, 0xa7, 0x8b,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0x8f, 0x93, 0xee, 0x89, 0x7b, 0xf2, 0x0c, 0xb7, 0xee, 0xc9, 0xab,
	0x22, 0xd9, 0x99, 0x99, 0x57, 0x9c, 0x9f, 0x93, 0x58, 0x84, 0x4e, 0xeb, 0x79, 0xba, 0x04, 0x41,
	0xf4, 0x0a, 0x19, 0x71, 0x71, 0xf8, 0x40, 0xad, 0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x36, 0x12,
	0xd0, 0x83, 0x58, 0xaf, 0x07, 0x13, 0x77, 0x62, 0x39, 0x71, 0x4f, 0x9e, 0x31, 0x08, 0xae, 0x4e,
	0x28, 0x89, 0x4b, 0x00, 0xc6, 0xf6, 0x49, 0x2c, 0x2e, 0x09, 0x2d, 0x4e, 0x4d, 0x91, 0x60, 0x06,
	0xbb, 0xc1, 0x0c, 0xea, 0x06, 0x3d, 0x22, 0xdc, 0x10, 0x50, 0x9a, 0x53, 0x9c, 0xea, 0x57, 0x9a,
	0x9b, 0x94, 0x5a, 0x14, 0x84, 0x61, 0x9e, 0x90, 0x27, 0x17, 0x47, 0x50, 0x6a, 0x61, 0x69, 0x6a,
	0x71, 0x49, 0xb1, 0x04, 0x8b, 0x02, 0x33, 0xd4, 0x7f, 0x8c, 0xc4, 0xfb, 0x0f, 0xae, 0x5d, 0xc8,
	0x9d, 0x8b, 0x3d, 0x28, 0xb5, 0xb8, 0x34, 0xa7, 0xa4, 0x58, 0x82, 0x95, 0x1c, 0x93, 0x60, 0xba,
	0xad, 0x58, 0x5e, 0x2c, 0x90, 0x67, 0x70, 0x32, 0xb9, 0xf0, 0x50, 0x8e, 0xe1, 0xc6, 0x43, 0x39,
	0x86, 0x0f, 0x0f, 0xe5, 0x18, 0x1b, 0x1e, 0xc9, 0x31, 0xae, 0x78, 0x24, 0xc7, 0x78, 0xe2, 0x91,
	0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xbe, 0x78, 0x24, 0xc7, 0xf0, 0xe1,
	0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x90,
	0xc4, 0x06, 0x8e, 0x43, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x49, 0xd8, 0x5f, 0x3b,
	0x02, 0x00, 0x00,
}

func (this *IndexBucket) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&object.IndexBucket{")
	s = append(s, "ObjID: "+fmt.Sprintf("%#v", this.ObjID)+",\n")
	if this.Lifeline != nil {
		s = append(s, "Lifeline: "+fmt.Sprintf("%#v", this.Lifeline)+",\n")
	}
	s = append(s, "LifelineLastUsed: "+fmt.Sprintf("%#v", this.LifelineLastUsed)+",\n")
	s = append(s, "Requests: "+fmt.Sprintf("%#v", this.Requests)+",\n")
	s = append(s, "Results: "+fmt.Sprintf("%#v", this.Results)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringIndexbucket(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *IndexBucket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexBucket) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintIndexbucket(dAtA, i, uint64(m.ObjID.Size()))
	n1, err := m.ObjID.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.Lifeline != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintIndexbucket(dAtA, i, uint64(m.Lifeline.Size()))
		n2, err := m.Lifeline.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintIndexbucket(dAtA, i, uint64(m.LifelineLastUsed.Size()))
	n3, err := m.LifelineLastUsed.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	if len(m.Requests) > 0 {
		for _, msg := range m.Requests {
			dAtA[i] = 0x22
			i++
			i = encodeVarintIndexbucket(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Results) > 0 {
		for _, msg := range m.Results {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintIndexbucket(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintIndexbucket(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *IndexBucket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjID.Size()
	n += 1 + l + sovIndexbucket(uint64(l))
	if m.Lifeline != nil {
		l = m.Lifeline.Size()
		n += 1 + l + sovIndexbucket(uint64(l))
	}
	l = m.LifelineLastUsed.Size()
	n += 1 + l + sovIndexbucket(uint64(l))
	if len(m.Requests) > 0 {
		for _, e := range m.Requests {
			l = e.Size()
			n += 1 + l + sovIndexbucket(uint64(l))
		}
	}
	if len(m.Results) > 0 {
		for _, e := range m.Results {
			l = e.Size()
			n += 1 + l + sovIndexbucket(uint64(l))
		}
	}
	return n
}

func sovIndexbucket(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozIndexbucket(x uint64) (n int) {
	return sovIndexbucket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *IndexBucket) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&IndexBucket{`,
		`ObjID:` + fmt.Sprintf("%v", this.ObjID) + `,`,
		`Lifeline:` + strings.Replace(fmt.Sprintf("%v", this.Lifeline), "Lifeline", "Lifeline", 1) + `,`,
		`LifelineLastUsed:` + fmt.Sprintf("%v", this.LifelineLastUsed) + `,`,
		`Requests:` + fmt.Sprintf("%v", this.Requests) + `,`,
		`Results:` + fmt.Sprintf("%v", this.Results) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringIndexbucket(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *IndexBucket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndexbucket
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
			return fmt.Errorf("proto: IndexBucket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexBucket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexbucket
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
				return ErrInvalidLengthIndexbucket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIndexbucket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lifeline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexbucket
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
				return ErrInvalidLengthIndexbucket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIndexbucket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Lifeline == nil {
				m.Lifeline = &Lifeline{}
			}
			if err := m.Lifeline.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LifelineLastUsed", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexbucket
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
				return ErrInvalidLengthIndexbucket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIndexbucket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LifelineLastUsed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requests", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexbucket
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
				return ErrInvalidLengthIndexbucket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIndexbucket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_insolar_insolar_insolar.ID
			m.Requests = append(m.Requests, v)
			if err := m.Requests[len(m.Requests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexbucket
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
				return ErrInvalidLengthIndexbucket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIndexbucket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_insolar_insolar_insolar.ID
			m.Results = append(m.Results, v)
			if err := m.Results[len(m.Results)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIndexbucket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIndexbucket
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIndexbucket
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
func skipIndexbucket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIndexbucket
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
					return 0, ErrIntOverflowIndexbucket
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
					return 0, ErrIntOverflowIndexbucket
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
				return 0, ErrInvalidLengthIndexbucket
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthIndexbucket
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowIndexbucket
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
				next, err := skipIndexbucket(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthIndexbucket
				}
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
	ErrInvalidLengthIndexbucket = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIndexbucket   = fmt.Errorf("proto: integer overflow")
)
