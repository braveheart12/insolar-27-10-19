// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ledger/heavy/exporter/record_exporter.proto

package exporter

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_insolar_insolar_insolar "github.com/insolar/insolar/insolar"
	record "github.com/insolar/insolar/insolar/record"
	grpc "google.golang.org/grpc"
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

type GetRecords struct {
	Polymorph    uint32                                         `protobuf:"varint,16,opt,name=Polymorph,proto3" json:"Polymorph,omitempty"`
	PulseNumber  github_com_insolar_insolar_insolar.PulseNumber `protobuf:"bytes,20,opt,name=PulseNumber,proto3,customtype=github.com/insolar/insolar/insolar.PulseNumber" json:"PulseNumber"`
	RecordNumber uint32                                         `protobuf:"varint,21,opt,name=RecordNumber,proto3" json:"RecordNumber,omitempty"`
	Count        uint32                                         `protobuf:"varint,22,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (m *GetRecords) Reset()      { *m = GetRecords{} }
func (*GetRecords) ProtoMessage() {}
func (*GetRecords) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfb4fbd68f50939d, []int{0}
}
func (m *GetRecords) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetRecords) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetRecords.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetRecords) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRecords.Merge(m, src)
}
func (m *GetRecords) XXX_Size() int {
	return m.Size()
}
func (m *GetRecords) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRecords.DiscardUnknown(m)
}

var xxx_messageInfo_GetRecords proto.InternalMessageInfo

func (m *GetRecords) GetPolymorph() uint32 {
	if m != nil {
		return m.Polymorph
	}
	return 0
}

func (m *GetRecords) GetRecordNumber() uint32 {
	if m != nil {
		return m.RecordNumber
	}
	return 0
}

func (m *GetRecords) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Record struct {
	Polymorph         uint32                                          `protobuf:"varint,16,opt,name=Polymorph,proto3" json:"Polymorph,omitempty"`
	RecordNumber      uint32                                          `protobuf:"varint,20,opt,name=RecordNumber,proto3" json:"RecordNumber,omitempty"`
	Record            record.Material                                 `protobuf:"bytes,21,opt,name=Record,proto3" json:"Record"`
	ShouldIterateFrom *github_com_insolar_insolar_insolar.PulseNumber `protobuf:"bytes,22,opt,name=ShouldIterateFrom,proto3,customtype=github.com/insolar/insolar/insolar.PulseNumber" json:"ShouldIterateFrom,omitempty"`
}

func (m *Record) Reset()      { *m = Record{} }
func (*Record) ProtoMessage() {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfb4fbd68f50939d, []int{1}
}
func (m *Record) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Record.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return m.Size()
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetPolymorph() uint32 {
	if m != nil {
		return m.Polymorph
	}
	return 0
}

func (m *Record) GetRecordNumber() uint32 {
	if m != nil {
		return m.RecordNumber
	}
	return 0
}

func (m *Record) GetRecord() record.Material {
	if m != nil {
		return m.Record
	}
	return record.Material{}
}

func init() {
	proto.RegisterType((*GetRecords)(nil), "exporter.GetRecords")
	proto.RegisterType((*Record)(nil), "exporter.Record")
}

func init() {
	proto.RegisterFile("ledger/heavy/exporter/record_exporter.proto", fileDescriptor_dfb4fbd68f50939d)
}

var fileDescriptor_dfb4fbd68f50939d = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x9d, 0x49, 0xde, 0x23, 0x3a, 0xa0, 0xc1, 0xa6, 0x1a, 0x42, 0xcc, 0x40, 0xba, 0x22, 0x31,
	0xb6, 0x06, 0x09, 0x1f, 0x80, 0x11, 0xe3, 0x42, 0x43, 0xea, 0xc6, 0x9d, 0x69, 0xe9, 0xd8, 0x92,
	0xb4, 0x0c, 0x19, 0xa6, 0x46, 0x76, 0x7e, 0x82, 0x9f, 0xe1, 0x57, 0xb8, 0x66, 0xc9, 0x92, 0xb8,
	0x20, 0x30, 0x6c, 0x5c, 0xf2, 0x09, 0xc6, 0x99, 0x56, 0x50, 0x4c, 0x48, 0x5c, 0xcd, 0x9c, 0xd3,
	0x7b, 0xcf, 0x39, 0xbd, 0x73, 0xd1, 0x51, 0x48, 0x3c, 0x9f, 0x30, 0x2b, 0x20, 0xce, 0xc3, 0xc0,
	0x22, 0x8f, 0x3d, 0xca, 0x38, 0x61, 0x16, 0x23, 0x6d, 0xca, 0xbc, 0xbb, 0x14, 0x9b, 0x3d, 0x46,
	0x39, 0xd5, 0xb6, 0x52, 0x5c, 0x3c, 0xf6, 0x3b, 0x3c, 0x88, 0x5d, 0xb3, 0x4d, 0x23, 0xcb, 0xa7,
	0x3e, 0xb5, 0x64, 0x81, 0x1b, 0xdf, 0x4b, 0x24, 0x81, 0xbc, 0xa9, 0xc6, 0x62, 0x7d, 0xa5, 0xbc,
	0xd3, 0xed, 0xd3, 0xd0, 0x61, 0x6b, 0xa7, 0xb2, 0x4c, 0x0e, 0xd5, 0x67, 0xbc, 0x42, 0x84, 0x2e,
	0x08, 0xb7, 0x25, 0xd7, 0xd7, 0x0e, 0xd1, 0x76, 0x8b, 0x86, 0x83, 0x88, 0xb2, 0x5e, 0x50, 0xc8,
	0x97, 0x61, 0x65, 0xc7, 0x5e, 0x12, 0xda, 0x2d, 0xca, 0xb6, 0xe2, 0xb0, 0x4f, 0xae, 0xe3, 0xc8,
	0x25, 0xac, 0xa0, 0x97, 0x61, 0x25, 0xd7, 0xa8, 0x0f, 0x27, 0x25, 0xf0, 0x36, 0x29, 0x99, 0x9b,
	0x13, 0x98, 0x2b, 0xdd, 0xf6, 0xaa, 0x94, 0x66, 0xa0, 0x9c, 0x8a, 0x90, 0x48, 0xef, 0x4b, 0xeb,
	0x6f, 0x9c, 0xa6, 0xa3, 0xff, 0x67, 0x34, 0xee, 0xf2, 0xc2, 0x81, 0xfc, 0xa8, 0x80, 0x31, 0x85,
	0x28, 0xa3, 0xca, 0x36, 0x84, 0xff, 0x69, 0xa1, 0xff, 0x62, 0x61, 0xa6, 0x5a, 0x32, 0x40, 0xb6,
	0x9a, 0x37, 0x93, 0x61, 0x5d, 0x39, 0x9c, 0xb0, 0x8e, 0x13, 0x36, 0xfe, 0x7d, 0xfe, 0xad, 0x9d,
	0x3a, 0x7a, 0x68, 0xef, 0x26, 0xa0, 0x71, 0xe8, 0x5d, 0x72, 0xc2, 0x1c, 0x4e, 0x9a, 0x8c, 0x46,
	0x32, 0x9e, 0x1a, 0x0b, 0xfc, 0xc3, 0x58, 0xd6, 0x05, 0xab, 0x4d, 0xb4, 0xab, 0xfc, 0xce, 0x93,
	0xe5, 0xd0, 0x6a, 0x28, 0xa3, 0xee, 0x9a, 0x6e, 0x7e, 0x6d, 0xd0, 0xf2, 0x19, 0x8b, 0xf9, 0x25,
	0xab, 0x28, 0x03, 0x9c, 0xc0, 0x46, 0x6d, 0x34, 0xc3, 0x60, 0x3c, 0xc3, 0x60, 0x31, 0xc3, 0xf0,
	0x49, 0x60, 0xf8, 0x22, 0x30, 0x1c, 0x0a, 0x0c, 0x47, 0x02, 0xc3, 0xa9, 0xc0, 0xf0, 0x5d, 0x60,
	0xb0, 0x10, 0x18, 0x3e, 0xcf, 0x31, 0x18, 0xcd, 0x31, 0x18, 0xcf, 0x31, 0x70, 0x33, 0x72, 0x51,
	0x4e, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x29, 0x5d, 0xda, 0xd8, 0xc8, 0x02, 0x00, 0x00,
}

func (this *GetRecords) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetRecords)
	if !ok {
		that2, ok := that.(GetRecords)
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
	if this.Polymorph != that1.Polymorph {
		return false
	}
	if !this.PulseNumber.Equal(that1.PulseNumber) {
		return false
	}
	if this.RecordNumber != that1.RecordNumber {
		return false
	}
	if this.Count != that1.Count {
		return false
	}
	return true
}
func (this *Record) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Record)
	if !ok {
		that2, ok := that.(Record)
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
	if this.Polymorph != that1.Polymorph {
		return false
	}
	if this.RecordNumber != that1.RecordNumber {
		return false
	}
	if !this.Record.Equal(&that1.Record) {
		return false
	}
	if that1.ShouldIterateFrom == nil {
		if this.ShouldIterateFrom != nil {
			return false
		}
	} else if !this.ShouldIterateFrom.Equal(*that1.ShouldIterateFrom) {
		return false
	}
	return true
}
func (this *GetRecords) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&exporter.GetRecords{")
	s = append(s, "Polymorph: "+fmt.Sprintf("%#v", this.Polymorph)+",\n")
	s = append(s, "PulseNumber: "+fmt.Sprintf("%#v", this.PulseNumber)+",\n")
	s = append(s, "RecordNumber: "+fmt.Sprintf("%#v", this.RecordNumber)+",\n")
	s = append(s, "Count: "+fmt.Sprintf("%#v", this.Count)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Record) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&exporter.Record{")
	s = append(s, "Polymorph: "+fmt.Sprintf("%#v", this.Polymorph)+",\n")
	s = append(s, "RecordNumber: "+fmt.Sprintf("%#v", this.RecordNumber)+",\n")
	s = append(s, "Record: "+strings.Replace(this.Record.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "ShouldIterateFrom: "+fmt.Sprintf("%#v", this.ShouldIterateFrom)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringRecordExporter(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RecordExporterClient is the client API for RecordExporter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecordExporterClient interface {
	Export(ctx context.Context, in *GetRecords, opts ...grpc.CallOption) (RecordExporter_ExportClient, error)
}

type recordExporterClient struct {
	cc *grpc.ClientConn
}

func NewRecordExporterClient(cc *grpc.ClientConn) RecordExporterClient {
	return &recordExporterClient{cc}
}

func (c *recordExporterClient) Export(ctx context.Context, in *GetRecords, opts ...grpc.CallOption) (RecordExporter_ExportClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RecordExporter_serviceDesc.Streams[0], "/exporter.RecordExporter/Export", opts...)
	if err != nil {
		return nil, err
	}
	x := &recordExporterExportClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RecordExporter_ExportClient interface {
	Recv() (*Record, error)
	grpc.ClientStream
}

type recordExporterExportClient struct {
	grpc.ClientStream
}

func (x *recordExporterExportClient) Recv() (*Record, error) {
	m := new(Record)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RecordExporterServer is the server API for RecordExporter service.
type RecordExporterServer interface {
	Export(*GetRecords, RecordExporter_ExportServer) error
}

func RegisterRecordExporterServer(s *grpc.Server, srv RecordExporterServer) {
	s.RegisterService(&_RecordExporter_serviceDesc, srv)
}

func _RecordExporter_Export_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRecords)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RecordExporterServer).Export(m, &recordExporterExportServer{stream})
}

type RecordExporter_ExportServer interface {
	Send(*Record) error
	grpc.ServerStream
}

type recordExporterExportServer struct {
	grpc.ServerStream
}

func (x *recordExporterExportServer) Send(m *Record) error {
	return x.ServerStream.SendMsg(m)
}

var _RecordExporter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "exporter.RecordExporter",
	HandlerType: (*RecordExporterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Export",
			Handler:       _RecordExporter_Export_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ledger/heavy/exporter/record_exporter.proto",
}

func (m *GetRecords) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetRecords) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Polymorph != 0 {
		dAtA[i] = 0x80
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintRecordExporter(dAtA, i, uint64(m.Polymorph))
	}
	dAtA[i] = 0xa2
	i++
	dAtA[i] = 0x1
	i++
	i = encodeVarintRecordExporter(dAtA, i, uint64(m.PulseNumber.Size()))
	n1, err := m.PulseNumber.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.RecordNumber != 0 {
		dAtA[i] = 0xa8
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintRecordExporter(dAtA, i, uint64(m.RecordNumber))
	}
	if m.Count != 0 {
		dAtA[i] = 0xb0
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintRecordExporter(dAtA, i, uint64(m.Count))
	}
	return i, nil
}

func (m *Record) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Record) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Polymorph != 0 {
		dAtA[i] = 0x80
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintRecordExporter(dAtA, i, uint64(m.Polymorph))
	}
	if m.RecordNumber != 0 {
		dAtA[i] = 0xa0
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintRecordExporter(dAtA, i, uint64(m.RecordNumber))
	}
	dAtA[i] = 0xaa
	i++
	dAtA[i] = 0x1
	i++
	i = encodeVarintRecordExporter(dAtA, i, uint64(m.Record.Size()))
	n2, err := m.Record.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.ShouldIterateFrom != nil {
		dAtA[i] = 0xb2
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintRecordExporter(dAtA, i, uint64(m.ShouldIterateFrom.Size()))
		n3, err := m.ShouldIterateFrom.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func encodeVarintRecordExporter(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GetRecords) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Polymorph != 0 {
		n += 2 + sovRecordExporter(uint64(m.Polymorph))
	}
	l = m.PulseNumber.Size()
	n += 2 + l + sovRecordExporter(uint64(l))
	if m.RecordNumber != 0 {
		n += 2 + sovRecordExporter(uint64(m.RecordNumber))
	}
	if m.Count != 0 {
		n += 2 + sovRecordExporter(uint64(m.Count))
	}
	return n
}

func (m *Record) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Polymorph != 0 {
		n += 2 + sovRecordExporter(uint64(m.Polymorph))
	}
	if m.RecordNumber != 0 {
		n += 2 + sovRecordExporter(uint64(m.RecordNumber))
	}
	l = m.Record.Size()
	n += 2 + l + sovRecordExporter(uint64(l))
	if m.ShouldIterateFrom != nil {
		l = m.ShouldIterateFrom.Size()
		n += 2 + l + sovRecordExporter(uint64(l))
	}
	return n
}

func sovRecordExporter(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRecordExporter(x uint64) (n int) {
	return sovRecordExporter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GetRecords) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetRecords{`,
		`Polymorph:` + fmt.Sprintf("%v", this.Polymorph) + `,`,
		`PulseNumber:` + fmt.Sprintf("%v", this.PulseNumber) + `,`,
		`RecordNumber:` + fmt.Sprintf("%v", this.RecordNumber) + `,`,
		`Count:` + fmt.Sprintf("%v", this.Count) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Record) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Record{`,
		`Polymorph:` + fmt.Sprintf("%v", this.Polymorph) + `,`,
		`RecordNumber:` + fmt.Sprintf("%v", this.RecordNumber) + `,`,
		`Record:` + strings.Replace(strings.Replace(this.Record.String(), "Material", "record.Material", 1), `&`, ``, 1) + `,`,
		`ShouldIterateFrom:` + fmt.Sprintf("%v", this.ShouldIterateFrom) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringRecordExporter(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GetRecords) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecordExporter
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
			return fmt.Errorf("proto: GetRecords: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetRecords: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Polymorph", wireType)
			}
			m.Polymorph = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecordExporter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Polymorph |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PulseNumber", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecordExporter
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
				return ErrInvalidLengthRecordExporter
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRecordExporter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PulseNumber.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 21:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecordNumber", wireType)
			}
			m.RecordNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecordExporter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RecordNumber |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 22:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecordExporter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRecordExporter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRecordExporter
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRecordExporter
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
func (m *Record) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecordExporter
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
			return fmt.Errorf("proto: Record: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Record: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Polymorph", wireType)
			}
			m.Polymorph = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecordExporter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Polymorph |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 20:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecordNumber", wireType)
			}
			m.RecordNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecordExporter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RecordNumber |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Record", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecordExporter
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
				return ErrInvalidLengthRecordExporter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRecordExporter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Record.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 22:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShouldIterateFrom", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecordExporter
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
				return ErrInvalidLengthRecordExporter
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRecordExporter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_insolar_insolar_insolar.PulseNumber
			m.ShouldIterateFrom = &v
			if err := m.ShouldIterateFrom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRecordExporter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRecordExporter
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRecordExporter
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
func skipRecordExporter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRecordExporter
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
					return 0, ErrIntOverflowRecordExporter
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
					return 0, ErrIntOverflowRecordExporter
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
				return 0, ErrInvalidLengthRecordExporter
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthRecordExporter
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRecordExporter
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
				next, err := skipRecordExporter(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthRecordExporter
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
	ErrInvalidLengthRecordExporter = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRecordExporter   = fmt.Errorf("proto: integer overflow")
)
