// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: insolar/payload/payload.proto

package payload

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

type GetObject struct {
	XPolymorph      int32                                 `protobuf:"varint,16,opt,name=_polymorph,json=Polymorph,proto3" json:"_polymorph,omitempty"`
	ObjectID        github_com_insolar_insolar_insolar.ID `protobuf:"bytes,20,opt,name=ObjectID,proto3,customtype=github.com/insolar/insolar/insolar.ID" json:"ObjectID"`
	ObjectRequestID github_com_insolar_insolar_insolar.ID `protobuf:"bytes,21,opt,name=ObjectRequestID,proto3,customtype=github.com/insolar/insolar/insolar.ID" json:"ObjectRequestID"`
}

func (m *GetObject) Reset()      { *m = GetObject{} }
func (*GetObject) ProtoMessage() {}
func (*GetObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_33334fec96407f54, []int{0}
}
func (m *GetObject) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetObject.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetObject.Merge(m, src)
}
func (m *GetObject) XXX_Size() int {
	return m.Size()
}
func (m *GetObject) XXX_DiscardUnknown() {
	xxx_messageInfo_GetObject.DiscardUnknown(m)
}

var xxx_messageInfo_GetObject proto.InternalMessageInfo

func (m *GetObject) GetXPolymorph() int32 {
	if m != nil {
		return m.XPolymorph
	}
	return 0
}

type Jet struct {
	XPolymorph int32                                          `protobuf:"varint,16,opt,name=_polymorph,json=Polymorph,proto3" json:"_polymorph,omitempty"`
	JetID      github_com_insolar_insolar_insolar.JetID       `protobuf:"bytes,20,opt,name=JetID,proto3,customtype=github.com/insolar/insolar/insolar.JetID" json:"JetID"`
	Pulse      github_com_insolar_insolar_insolar.PulseNumber `protobuf:"bytes,21,opt,name=Pulse,proto3,customtype=github.com/insolar/insolar/insolar.PulseNumber" json:"Pulse"`
}

func (m *Jet) Reset()      { *m = Jet{} }
func (*Jet) ProtoMessage() {}
func (*Jet) Descriptor() ([]byte, []int) {
	return fileDescriptor_33334fec96407f54, []int{1}
}
func (m *Jet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Jet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Jet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Jet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Jet.Merge(m, src)
}
func (m *Jet) XXX_Size() int {
	return m.Size()
}
func (m *Jet) XXX_DiscardUnknown() {
	xxx_messageInfo_Jet.DiscardUnknown(m)
}

var xxx_messageInfo_Jet proto.InternalMessageInfo

func (m *Jet) GetXPolymorph() int32 {
	if m != nil {
		return m.XPolymorph
	}
	return 0
}

func init() {
	proto.RegisterType((*GetObject)(nil), "payload.GetObject")
	proto.RegisterType((*Jet)(nil), "payload.Jet")
}

func init() { proto.RegisterFile("insolar/payload/payload.proto", fileDescriptor_33334fec96407f54) }

var fileDescriptor_33334fec96407f54 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0xcc, 0x2b, 0xce,
	0xcf, 0x49, 0x2c, 0xd2, 0x2f, 0x48, 0xac, 0xcc, 0xc9, 0x4f, 0x4c, 0x81, 0xd1, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae, 0x94, 0x6e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e,
	0x72, 0x7e, 0xae, 0x7e, 0x7a, 0x7e, 0x7a, 0xbe, 0x3e, 0x58, 0x3e, 0xa9, 0x34, 0x0d, 0xcc, 0x03,
	0x73, 0xc0, 0x2c, 0x88, 0x3e, 0xa5, 0x73, 0x8c, 0x5c, 0x9c, 0xee, 0xa9, 0x25, 0xfe, 0x49, 0x59,
	0xa9, 0xc9, 0x25, 0x42, 0xb2, 0x5c, 0x5c, 0xf1, 0x05, 0xf9, 0x39, 0x95, 0xb9, 0xf9, 0x45, 0x05,
	0x19, 0x12, 0x02, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x9c, 0x01, 0x30, 0x01, 0x21, 0x4f, 0x2e, 0x0e,
	0x88, 0x42, 0x4f, 0x17, 0x09, 0x11, 0x05, 0x46, 0x0d, 0x1e, 0x27, 0xdd, 0x13, 0xf7, 0xe4, 0x19,
	0x6e, 0xdd, 0x93, 0x57, 0x45, 0xb2, 0x15, 0xe6, 0x50, 0x34, 0x5a, 0xcf, 0xd3, 0x25, 0x08, 0xae,
	0x5d, 0x28, 0x9c, 0x8b, 0x1f, 0xc2, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x06, 0x99, 0x28, 0x4a,
	0x8e, 0x89, 0xe8, 0xa6, 0x28, 0xed, 0x62, 0xe4, 0x62, 0xf6, 0x4a, 0x25, 0xe8, 0x15, 0x37, 0x2e,
	0x56, 0xaf, 0x54, 0x84, 0x3f, 0x0c, 0xa0, 0xb6, 0x6a, 0x10, 0x61, 0x2b, 0x58, 0x5f, 0x10, 0x44,
	0xbb, 0x90, 0x0f, 0x17, 0x6b, 0x40, 0x69, 0x4e, 0x71, 0x2a, 0xd4, 0xf5, 0x66, 0x50, 0x73, 0xf4,
	0x88, 0x30, 0x07, 0xac, 0xcf, 0xaf, 0x34, 0x37, 0x29, 0xb5, 0x28, 0x08, 0x62, 0x88, 0x93, 0xc9,
	0x85, 0x87, 0x72, 0x0c, 0x37, 0x1e, 0xca, 0x31, 0x7c, 0x78, 0x28, 0xc7, 0xd8, 0xf0, 0x48, 0x8e,
	0x71, 0xc5, 0x23, 0x39, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48,
	0x8e, 0xf1, 0xc5, 0x23, 0x39, 0x86, 0x0f, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0,
	0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x24, 0x36, 0x70, 0x54, 0x1a, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xa4, 0xfb, 0x88, 0xbf, 0x23, 0x02, 0x00, 0x00,
}

func (this *GetObject) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetObject)
	if !ok {
		that2, ok := that.(GetObject)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.XPolymorph != that1.XPolymorph {
		return false
	}
	if !this.ObjectID.Equal(that1.ObjectID) {
		return false
	}
	if !this.ObjectRequestID.Equal(that1.ObjectRequestID) {
		return false
	}
	return true
}
func (this *Jet) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Jet)
	if !ok {
		that2, ok := that.(Jet)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.XPolymorph != that1.XPolymorph {
		return false
	}
	if !this.JetID.Equal(that1.JetID) {
		return false
	}
	if !this.Pulse.Equal(that1.Pulse) {
		return false
	}
	return true
}
func (this *GetObject) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&payload.GetObject{")
	s = append(s, "XPolymorph: "+fmt.Sprintf("%#v", this.XPolymorph)+",\n")
	s = append(s, "ObjectID: "+fmt.Sprintf("%#v", this.ObjectID)+",\n")
	s = append(s, "ObjectRequestID: "+fmt.Sprintf("%#v", this.ObjectRequestID)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Jet) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&payload.Jet{")
	s = append(s, "XPolymorph: "+fmt.Sprintf("%#v", this.XPolymorph)+",\n")
	s = append(s, "JetID: "+fmt.Sprintf("%#v", this.JetID)+",\n")
	s = append(s, "Pulse: "+fmt.Sprintf("%#v", this.Pulse)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPayload(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *GetObject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetObject) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XPolymorph != 0 {
		dAtA[i] = 0x80
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintPayload(dAtA, i, uint64(m.XPolymorph))
	}
	dAtA[i] = 0xa2
	i++
	dAtA[i] = 0x1
	i++
	i = encodeVarintPayload(dAtA, i, uint64(m.ObjectID.Size()))
	n1, err := m.ObjectID.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0xaa
	i++
	dAtA[i] = 0x1
	i++
	i = encodeVarintPayload(dAtA, i, uint64(m.ObjectRequestID.Size()))
	n2, err := m.ObjectRequestID.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func (m *Jet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Jet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XPolymorph != 0 {
		dAtA[i] = 0x80
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintPayload(dAtA, i, uint64(m.XPolymorph))
	}
	dAtA[i] = 0xa2
	i++
	dAtA[i] = 0x1
	i++
	i = encodeVarintPayload(dAtA, i, uint64(m.JetID.Size()))
	n3, err := m.JetID.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0xaa
	i++
	dAtA[i] = 0x1
	i++
	i = encodeVarintPayload(dAtA, i, uint64(m.Pulse.Size()))
	n4, err := m.Pulse.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func encodeVarintPayload(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GetObject) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XPolymorph != 0 {
		n += 2 + sovPayload(uint64(m.XPolymorph))
	}
	l = m.ObjectID.Size()
	n += 2 + l + sovPayload(uint64(l))
	l = m.ObjectRequestID.Size()
	n += 2 + l + sovPayload(uint64(l))
	return n
}

func (m *Jet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XPolymorph != 0 {
		n += 2 + sovPayload(uint64(m.XPolymorph))
	}
	l = m.JetID.Size()
	n += 2 + l + sovPayload(uint64(l))
	l = m.Pulse.Size()
	n += 2 + l + sovPayload(uint64(l))
	return n
}

func sovPayload(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPayload(x uint64) (n int) {
	return sovPayload(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GetObject) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetObject{`,
		`XPolymorph:` + fmt.Sprintf("%v", this.XPolymorph) + `,`,
		`ObjectID:` + fmt.Sprintf("%v", this.ObjectID) + `,`,
		`ObjectRequestID:` + fmt.Sprintf("%v", this.ObjectRequestID) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Jet) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Jet{`,
		`XPolymorph:` + fmt.Sprintf("%v", this.XPolymorph) + `,`,
		`JetID:` + fmt.Sprintf("%v", this.JetID) + `,`,
		`Pulse:` + fmt.Sprintf("%v", this.Pulse) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPayload(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GetObject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayload
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
			return fmt.Errorf("proto: GetObject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetObject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field XPolymorph", wireType)
			}
			m.XPolymorph = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.XPolymorph |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
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
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectRequestID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
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
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectRequestID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPayload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPayload
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPayload
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
func (m *Jet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayload
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
			return fmt.Errorf("proto: Jet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Jet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field XPolymorph", wireType)
			}
			m.XPolymorph = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.XPolymorph |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JetID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
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
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.JetID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pulse", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
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
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPayload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pulse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPayload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPayload
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPayload
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
func skipPayload(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPayload
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
					return 0, ErrIntOverflowPayload
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
					return 0, ErrIntOverflowPayload
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
				return 0, ErrInvalidLengthPayload
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthPayload
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPayload
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
				next, err := skipPayload(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthPayload
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
	ErrInvalidLengthPayload = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPayload   = fmt.Errorf("proto: integer overflow")
)
