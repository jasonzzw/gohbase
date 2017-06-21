// Code generated by protoc-gen-go.
// source: Quota.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type QuotaScope int32

const (
	QuotaScope_CLUSTER QuotaScope = 1
	QuotaScope_MACHINE QuotaScope = 2
)

var QuotaScope_name = map[int32]string{
	1: "CLUSTER",
	2: "MACHINE",
}
var QuotaScope_value = map[string]int32{
	"CLUSTER": 1,
	"MACHINE": 2,
}

func (x QuotaScope) Enum() *QuotaScope {
	p := new(QuotaScope)
	*p = x
	return p
}
func (x QuotaScope) String() string {
	return proto.EnumName(QuotaScope_name, int32(x))
}
func (x *QuotaScope) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(QuotaScope_value, data, "QuotaScope")
	if err != nil {
		return err
	}
	*x = QuotaScope(value)
	return nil
}
func (QuotaScope) EnumDescriptor() ([]byte, []int) { return fileDescriptor21, []int{0} }

type ThrottleType int32

const (
	ThrottleType_REQUEST_NUMBER ThrottleType = 1
	ThrottleType_REQUEST_SIZE   ThrottleType = 2
	ThrottleType_WRITE_NUMBER   ThrottleType = 3
	ThrottleType_WRITE_SIZE     ThrottleType = 4
	ThrottleType_READ_NUMBER    ThrottleType = 5
	ThrottleType_READ_SIZE      ThrottleType = 6
)

var ThrottleType_name = map[int32]string{
	1: "REQUEST_NUMBER",
	2: "REQUEST_SIZE",
	3: "WRITE_NUMBER",
	4: "WRITE_SIZE",
	5: "READ_NUMBER",
	6: "READ_SIZE",
}
var ThrottleType_value = map[string]int32{
	"REQUEST_NUMBER": 1,
	"REQUEST_SIZE":   2,
	"WRITE_NUMBER":   3,
	"WRITE_SIZE":     4,
	"READ_NUMBER":    5,
	"READ_SIZE":      6,
}

func (x ThrottleType) Enum() *ThrottleType {
	p := new(ThrottleType)
	*p = x
	return p
}
func (x ThrottleType) String() string {
	return proto.EnumName(ThrottleType_name, int32(x))
}
func (x *ThrottleType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ThrottleType_value, data, "ThrottleType")
	if err != nil {
		return err
	}
	*x = ThrottleType(value)
	return nil
}
func (ThrottleType) EnumDescriptor() ([]byte, []int) { return fileDescriptor21, []int{1} }

type QuotaType int32

const (
	QuotaType_THROTTLE QuotaType = 1
)

var QuotaType_name = map[int32]string{
	1: "THROTTLE",
}
var QuotaType_value = map[string]int32{
	"THROTTLE": 1,
}

func (x QuotaType) Enum() *QuotaType {
	p := new(QuotaType)
	*p = x
	return p
}
func (x QuotaType) String() string {
	return proto.EnumName(QuotaType_name, int32(x))
}
func (x *QuotaType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(QuotaType_value, data, "QuotaType")
	if err != nil {
		return err
	}
	*x = QuotaType(value)
	return nil
}
func (QuotaType) EnumDescriptor() ([]byte, []int) { return fileDescriptor21, []int{2} }

type TimedQuota struct {
	TimeUnit         *TimeUnit   `protobuf:"varint,1,req,name=time_unit,json=timeUnit,enum=pb.TimeUnit" json:"time_unit,omitempty"`
	SoftLimit        *uint64     `protobuf:"varint,2,opt,name=soft_limit,json=softLimit" json:"soft_limit,omitempty"`
	Share            *float32    `protobuf:"fixed32,3,opt,name=share" json:"share,omitempty"`
	Scope            *QuotaScope `protobuf:"varint,4,opt,name=scope,enum=pb.QuotaScope,def=2" json:"scope,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *TimedQuota) Reset()                    { *m = TimedQuota{} }
func (m *TimedQuota) String() string            { return proto.CompactTextString(m) }
func (*TimedQuota) ProtoMessage()               {}
func (*TimedQuota) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{0} }

const Default_TimedQuota_Scope QuotaScope = QuotaScope_MACHINE

func (m *TimedQuota) GetTimeUnit() TimeUnit {
	if m != nil && m.TimeUnit != nil {
		return *m.TimeUnit
	}
	return TimeUnit_NANOSECONDS
}

func (m *TimedQuota) GetSoftLimit() uint64 {
	if m != nil && m.SoftLimit != nil {
		return *m.SoftLimit
	}
	return 0
}

func (m *TimedQuota) GetShare() float32 {
	if m != nil && m.Share != nil {
		return *m.Share
	}
	return 0
}

func (m *TimedQuota) GetScope() QuotaScope {
	if m != nil && m.Scope != nil {
		return *m.Scope
	}
	return Default_TimedQuota_Scope
}

type Throttle struct {
	ReqNum           *TimedQuota `protobuf:"bytes,1,opt,name=req_num,json=reqNum" json:"req_num,omitempty"`
	ReqSize          *TimedQuota `protobuf:"bytes,2,opt,name=req_size,json=reqSize" json:"req_size,omitempty"`
	WriteNum         *TimedQuota `protobuf:"bytes,3,opt,name=write_num,json=writeNum" json:"write_num,omitempty"`
	WriteSize        *TimedQuota `protobuf:"bytes,4,opt,name=write_size,json=writeSize" json:"write_size,omitempty"`
	ReadNum          *TimedQuota `protobuf:"bytes,5,opt,name=read_num,json=readNum" json:"read_num,omitempty"`
	ReadSize         *TimedQuota `protobuf:"bytes,6,opt,name=read_size,json=readSize" json:"read_size,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Throttle) Reset()                    { *m = Throttle{} }
func (m *Throttle) String() string            { return proto.CompactTextString(m) }
func (*Throttle) ProtoMessage()               {}
func (*Throttle) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{1} }

func (m *Throttle) GetReqNum() *TimedQuota {
	if m != nil {
		return m.ReqNum
	}
	return nil
}

func (m *Throttle) GetReqSize() *TimedQuota {
	if m != nil {
		return m.ReqSize
	}
	return nil
}

func (m *Throttle) GetWriteNum() *TimedQuota {
	if m != nil {
		return m.WriteNum
	}
	return nil
}

func (m *Throttle) GetWriteSize() *TimedQuota {
	if m != nil {
		return m.WriteSize
	}
	return nil
}

func (m *Throttle) GetReadNum() *TimedQuota {
	if m != nil {
		return m.ReadNum
	}
	return nil
}

func (m *Throttle) GetReadSize() *TimedQuota {
	if m != nil {
		return m.ReadSize
	}
	return nil
}

type ThrottleRequest struct {
	Type             *ThrottleType `protobuf:"varint,1,opt,name=type,enum=pb.ThrottleType" json:"type,omitempty"`
	TimedQuota       *TimedQuota   `protobuf:"bytes,2,opt,name=timed_quota,json=timedQuota" json:"timed_quota,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *ThrottleRequest) Reset()                    { *m = ThrottleRequest{} }
func (m *ThrottleRequest) String() string            { return proto.CompactTextString(m) }
func (*ThrottleRequest) ProtoMessage()               {}
func (*ThrottleRequest) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{2} }

func (m *ThrottleRequest) GetType() ThrottleType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ThrottleType_REQUEST_NUMBER
}

func (m *ThrottleRequest) GetTimedQuota() *TimedQuota {
	if m != nil {
		return m.TimedQuota
	}
	return nil
}

type Quotas struct {
	BypassGlobals    *bool     `protobuf:"varint,1,opt,name=bypass_globals,json=bypassGlobals,def=0" json:"bypass_globals,omitempty"`
	Throttle         *Throttle `protobuf:"bytes,2,opt,name=throttle" json:"throttle,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Quotas) Reset()                    { *m = Quotas{} }
func (m *Quotas) String() string            { return proto.CompactTextString(m) }
func (*Quotas) ProtoMessage()               {}
func (*Quotas) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{3} }

const Default_Quotas_BypassGlobals bool = false

func (m *Quotas) GetBypassGlobals() bool {
	if m != nil && m.BypassGlobals != nil {
		return *m.BypassGlobals
	}
	return Default_Quotas_BypassGlobals
}

func (m *Quotas) GetThrottle() *Throttle {
	if m != nil {
		return m.Throttle
	}
	return nil
}

type QuotaUsage struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *QuotaUsage) Reset()                    { *m = QuotaUsage{} }
func (m *QuotaUsage) String() string            { return proto.CompactTextString(m) }
func (*QuotaUsage) ProtoMessage()               {}
func (*QuotaUsage) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{4} }

func init() {
	proto.RegisterType((*TimedQuota)(nil), "pb.TimedQuota")
	proto.RegisterType((*Throttle)(nil), "pb.Throttle")
	proto.RegisterType((*ThrottleRequest)(nil), "pb.ThrottleRequest")
	proto.RegisterType((*Quotas)(nil), "pb.Quotas")
	proto.RegisterType((*QuotaUsage)(nil), "pb.QuotaUsage")
	proto.RegisterEnum("pb.QuotaScope", QuotaScope_name, QuotaScope_value)
	proto.RegisterEnum("pb.ThrottleType", ThrottleType_name, ThrottleType_value)
	proto.RegisterEnum("pb.QuotaType", QuotaType_name, QuotaType_value)
}

var fileDescriptor21 = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xd5, 0xba, 0xf9, 0x70, 0xc6, 0xae, 0x1b, 0xad, 0x38, 0x04, 0x24, 0xa4, 0x2a, 0x42, 0x10,
	0x02, 0x35, 0x52, 0x8f, 0xbd, 0x25, 0xc5, 0x22, 0x91, 0xd2, 0x40, 0x37, 0x8e, 0x90, 0xb8, 0x98,
	0x75, 0xbd, 0x89, 0x23, 0x25, 0xb1, 0x63, 0xaf, 0x41, 0xe5, 0x17, 0x70, 0xe6, 0xca, 0x85, 0x9f,
	0xca, 0xee, 0xd8, 0x09, 0x95, 0x48, 0x6f, 0x3b, 0xef, 0xbd, 0x79, 0xf3, 0x3c, 0xde, 0x05, 0xeb,
	0xb6, 0x48, 0x24, 0x77, 0xd3, 0x2c, 0x91, 0x09, 0x35, 0xd2, 0xf0, 0x99, 0x35, 0x1a, 0xf2, 0x5c,
	0x94, 0x40, 0xf7, 0x37, 0x01, 0xf0, 0x57, 0x1b, 0x11, 0xa1, 0x8a, 0xbe, 0x86, 0x96, 0x54, 0x55,
	0x50, 0x6c, 0x57, 0xb2, 0x43, 0xce, 0x8d, 0x9e, 0x73, 0x69, 0xbb, 0x69, 0xe8, 0x6a, 0xc9, 0x5c,
	0x61, 0xcc, 0x94, 0xd5, 0x89, 0x3e, 0x07, 0xc8, 0x93, 0x85, 0x0c, 0xd6, 0xab, 0x8d, 0xd2, 0x1a,
	0xe7, 0xa4, 0x57, 0x63, 0x2d, 0x8d, 0x4c, 0x34, 0x40, 0x9f, 0x40, 0x3d, 0x8f, 0x79, 0x26, 0x3a,
	0x27, 0x8a, 0x31, 0x58, 0x59, 0xd0, 0x0b, 0x85, 0xde, 0x25, 0xa9, 0xe8, 0xd4, 0x14, 0xea, 0x5c,
	0x3a, 0xda, 0x1b, 0x27, 0xcf, 0x34, 0x7a, 0xd5, 0xbc, 0x19, 0x5c, 0x8f, 0xc6, 0x53, 0x8f, 0x95,
	0xaa, 0xee, 0x2f, 0x03, 0x4c, 0x3f, 0x56, 0x41, 0xe5, 0x5a, 0xd0, 0x57, 0xd0, 0xcc, 0xc4, 0x2e,
	0xd8, 0x16, 0x1b, 0x95, 0x8c, 0xf4, 0xac, 0xb2, 0xfb, 0x5f, 0x78, 0xd6, 0x50, 0xf4, 0xb4, 0xd8,
	0xa8, 0x8f, 0x30, 0xb5, 0x30, 0x5f, 0xfd, 0x10, 0x98, 0xeb, 0x7f, 0xa5, 0x36, 0x9a, 0x29, 0x9a,
	0xbe, 0x81, 0xd6, 0xf7, 0x6c, 0x25, 0x05, 0xba, 0x9e, 0x1c, 0xd5, 0x9a, 0x28, 0xd0, 0xbe, 0x17,
	0x00, 0xa5, 0x18, 0x9d, 0x6b, 0x47, 0xd5, 0xa5, 0x1d, 0x7a, 0x63, 0x0c, 0x1e, 0xa1, 0x75, 0xfd,
	0xb1, 0x18, 0x3c, 0xd2, 0xce, 0x2a, 0x06, 0x4a, 0xd1, 0xb8, 0x71, 0x3c, 0x86, 0x16, 0x68, 0xdf,
	0x6e, 0x0c, 0x67, 0xfb, 0x9d, 0x30, 0xb1, 0x2b, 0x44, 0x2e, 0xe9, 0x0b, 0xa8, 0xc9, 0x7b, 0xb5,
	0x55, 0x82, 0x5b, 0x6d, 0x63, 0x6b, 0x25, 0xf1, 0x15, 0xce, 0x90, 0xa5, 0xef, 0xc0, 0xd2, 0x7f,
	0x2f, 0x0a, 0x76, 0xda, 0xf1, 0x91, 0xd5, 0x80, 0x3c, 0x9c, 0xbb, 0x5f, 0xa1, 0x81, 0x87, 0x9c,
	0xbe, 0x05, 0x27, 0xbc, 0x4f, 0x79, 0x9e, 0x07, 0xcb, 0x75, 0x12, 0xf2, 0x75, 0x8e, 0xa3, 0xcc,
	0xab, 0xfa, 0x42, 0x9d, 0x05, 0x3b, 0x2d, 0xc9, 0x0f, 0x25, 0x47, 0x7b, 0x60, 0xca, 0x6a, 0x7c,
	0x35, 0xc5, 0x7e, 0x18, 0x89, 0x1d, 0xd8, 0xae, 0x0d, 0x80, 0x13, 0xe6, 0x39, 0x5f, 0x8a, 0xfe,
	0xcb, 0xaa, 0xc2, 0xcb, 0x40, 0x2d, 0x68, 0x5e, 0x4f, 0xe6, 0x33, 0xdf, 0x63, 0x6d, 0xa2, 0x8b,
	0xea, 0x6e, 0xb4, 0x8d, 0xfe, 0x37, 0xb0, 0x1f, 0x7e, 0x1e, 0xa5, 0xe0, 0x30, 0xef, 0x76, 0xee,
	0xcd, 0xfc, 0x60, 0x3a, 0xbf, 0x19, 0x62, 0x43, 0x1b, 0xec, 0x3d, 0x36, 0x1b, 0x7f, 0x51, 0x5d,
	0x1a, 0xf9, 0xcc, 0xc6, 0xbe, 0xb7, 0xd7, 0x9c, 0x50, 0x07, 0xa0, 0x44, 0x50, 0x51, 0xa3, 0x67,
	0x60, 0x31, 0x6f, 0xf0, 0x7e, 0x2f, 0xa8, 0xd3, 0x53, 0x68, 0x21, 0x80, 0x7c, 0xa3, 0xff, 0x14,
	0x5a, 0x98, 0x0f, 0x87, 0xda, 0xea, 0x6a, 0x8e, 0xd8, 0x47, 0xdf, 0x9f, 0x78, 0x6d, 0x32, 0x1c,
	0x40, 0x3f, 0xc9, 0x96, 0x2e, 0x4f, 0xf9, 0x5d, 0x2c, 0xdc, 0x98, 0x47, 0x49, 0x92, 0xba, 0x71,
	0x78, 0x78, 0x68, 0x61, 0xb1, 0x70, 0x97, 0x62, 0x2b, 0x32, 0x2e, 0x45, 0x34, 0x2c, 0xdf, 0xe4,
	0x27, 0x4d, 0xe4, 0x23, 0xf2, 0x93, 0x90, 0x3f, 0x84, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x31,
	0x6b, 0x5c, 0x31, 0xa9, 0x03, 0x00, 0x00,
}
