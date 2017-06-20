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
func (QuotaScope) EnumDescriptor() ([]byte, []int) { return fileDescriptor18, []int{0} }

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
func (ThrottleType) EnumDescriptor() ([]byte, []int) { return fileDescriptor18, []int{1} }

type QuotaType int32

const (
	QuotaType_THROTTLE QuotaType = 1
	QuotaType_SPACE    QuotaType = 2
)

var QuotaType_name = map[int32]string{
	1: "THROTTLE",
	2: "SPACE",
}
var QuotaType_value = map[string]int32{
	"THROTTLE": 1,
	"SPACE":    2,
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
func (QuotaType) EnumDescriptor() ([]byte, []int) { return fileDescriptor18, []int{2} }

// Defines what action should be taken when the SpaceQuota is violated
type SpaceViolationPolicy int32

const (
	SpaceViolationPolicy_DISABLE               SpaceViolationPolicy = 1
	SpaceViolationPolicy_NO_WRITES_COMPACTIONS SpaceViolationPolicy = 2
	SpaceViolationPolicy_NO_WRITES             SpaceViolationPolicy = 3
	SpaceViolationPolicy_NO_INSERTS            SpaceViolationPolicy = 4
)

var SpaceViolationPolicy_name = map[int32]string{
	1: "DISABLE",
	2: "NO_WRITES_COMPACTIONS",
	3: "NO_WRITES",
	4: "NO_INSERTS",
}
var SpaceViolationPolicy_value = map[string]int32{
	"DISABLE":               1,
	"NO_WRITES_COMPACTIONS": 2,
	"NO_WRITES":             3,
	"NO_INSERTS":            4,
}

func (x SpaceViolationPolicy) Enum() *SpaceViolationPolicy {
	p := new(SpaceViolationPolicy)
	*p = x
	return p
}
func (x SpaceViolationPolicy) String() string {
	return proto.EnumName(SpaceViolationPolicy_name, int32(x))
}
func (x *SpaceViolationPolicy) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SpaceViolationPolicy_value, data, "SpaceViolationPolicy")
	if err != nil {
		return err
	}
	*x = SpaceViolationPolicy(value)
	return nil
}
func (SpaceViolationPolicy) EnumDescriptor() ([]byte, []int) { return fileDescriptor18, []int{3} }

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
func (*TimedQuota) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{0} }

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
func (*Throttle) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{1} }

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
func (*ThrottleRequest) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{2} }

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
	BypassGlobals    *bool       `protobuf:"varint,1,opt,name=bypass_globals,json=bypassGlobals,def=0" json:"bypass_globals,omitempty"`
	Throttle         *Throttle   `protobuf:"bytes,2,opt,name=throttle" json:"throttle,omitempty"`
	Space            *SpaceQuota `protobuf:"bytes,3,opt,name=space" json:"space,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Quotas) Reset()                    { *m = Quotas{} }
func (m *Quotas) String() string            { return proto.CompactTextString(m) }
func (*Quotas) ProtoMessage()               {}
func (*Quotas) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{3} }

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

func (m *Quotas) GetSpace() *SpaceQuota {
	if m != nil {
		return m.Space
	}
	return nil
}

type QuotaUsage struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *QuotaUsage) Reset()                    { *m = QuotaUsage{} }
func (m *QuotaUsage) String() string            { return proto.CompactTextString(m) }
func (*QuotaUsage) ProtoMessage()               {}
func (*QuotaUsage) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{4} }

// Defines a limit on the amount of filesystem space used by a table/namespace
type SpaceQuota struct {
	SoftLimit        *uint64               `protobuf:"varint,1,opt,name=soft_limit,json=softLimit" json:"soft_limit,omitempty"`
	ViolationPolicy  *SpaceViolationPolicy `protobuf:"varint,2,opt,name=violation_policy,json=violationPolicy,enum=pb.SpaceViolationPolicy" json:"violation_policy,omitempty"`
	Remove           *bool                 `protobuf:"varint,3,opt,name=remove,def=0" json:"remove,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *SpaceQuota) Reset()                    { *m = SpaceQuota{} }
func (m *SpaceQuota) String() string            { return proto.CompactTextString(m) }
func (*SpaceQuota) ProtoMessage()               {}
func (*SpaceQuota) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{5} }

const Default_SpaceQuota_Remove bool = false

func (m *SpaceQuota) GetSoftLimit() uint64 {
	if m != nil && m.SoftLimit != nil {
		return *m.SoftLimit
	}
	return 0
}

func (m *SpaceQuota) GetViolationPolicy() SpaceViolationPolicy {
	if m != nil && m.ViolationPolicy != nil {
		return *m.ViolationPolicy
	}
	return SpaceViolationPolicy_DISABLE
}

func (m *SpaceQuota) GetRemove() bool {
	if m != nil && m.Remove != nil {
		return *m.Remove
	}
	return Default_SpaceQuota_Remove
}

// The Request to limit space usage (to allow for schema evolution not tied to SpaceQuota).
type SpaceLimitRequest struct {
	Quota            *SpaceQuota `protobuf:"bytes,1,opt,name=quota" json:"quota,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SpaceLimitRequest) Reset()                    { *m = SpaceLimitRequest{} }
func (m *SpaceLimitRequest) String() string            { return proto.CompactTextString(m) }
func (*SpaceLimitRequest) ProtoMessage()               {}
func (*SpaceLimitRequest) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{6} }

func (m *SpaceLimitRequest) GetQuota() *SpaceQuota {
	if m != nil {
		return m.Quota
	}
	return nil
}

// Represents the state of a quota on a table. Either the quota is not in violation
// or it is in violatino there is a violation policy which should be in effect.
type SpaceQuotaStatus struct {
	ViolationPolicy  *SpaceViolationPolicy `protobuf:"varint,1,opt,name=violation_policy,json=violationPolicy,enum=pb.SpaceViolationPolicy" json:"violation_policy,omitempty"`
	InViolation      *bool                 `protobuf:"varint,2,opt,name=in_violation,json=inViolation" json:"in_violation,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *SpaceQuotaStatus) Reset()                    { *m = SpaceQuotaStatus{} }
func (m *SpaceQuotaStatus) String() string            { return proto.CompactTextString(m) }
func (*SpaceQuotaStatus) ProtoMessage()               {}
func (*SpaceQuotaStatus) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{7} }

func (m *SpaceQuotaStatus) GetViolationPolicy() SpaceViolationPolicy {
	if m != nil && m.ViolationPolicy != nil {
		return *m.ViolationPolicy
	}
	return SpaceViolationPolicy_DISABLE
}

func (m *SpaceQuotaStatus) GetInViolation() bool {
	if m != nil && m.InViolation != nil {
		return *m.InViolation
	}
	return false
}

// Message stored in the value of hbase:quota table to denote the status of a table WRT
// the quota applicable to it.
type SpaceQuotaSnapshot struct {
	QuotaStatus      *SpaceQuotaStatus `protobuf:"bytes,1,opt,name=quota_status,json=quotaStatus" json:"quota_status,omitempty"`
	QuotaUsage       *uint64           `protobuf:"varint,2,opt,name=quota_usage,json=quotaUsage" json:"quota_usage,omitempty"`
	QuotaLimit       *uint64           `protobuf:"varint,3,opt,name=quota_limit,json=quotaLimit" json:"quota_limit,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *SpaceQuotaSnapshot) Reset()                    { *m = SpaceQuotaSnapshot{} }
func (m *SpaceQuotaSnapshot) String() string            { return proto.CompactTextString(m) }
func (*SpaceQuotaSnapshot) ProtoMessage()               {}
func (*SpaceQuotaSnapshot) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{8} }

func (m *SpaceQuotaSnapshot) GetQuotaStatus() *SpaceQuotaStatus {
	if m != nil {
		return m.QuotaStatus
	}
	return nil
}

func (m *SpaceQuotaSnapshot) GetQuotaUsage() uint64 {
	if m != nil && m.QuotaUsage != nil {
		return *m.QuotaUsage
	}
	return 0
}

func (m *SpaceQuotaSnapshot) GetQuotaLimit() uint64 {
	if m != nil && m.QuotaLimit != nil {
		return *m.QuotaLimit
	}
	return 0
}

func init() {
	proto.RegisterType((*TimedQuota)(nil), "pb.TimedQuota")
	proto.RegisterType((*Throttle)(nil), "pb.Throttle")
	proto.RegisterType((*ThrottleRequest)(nil), "pb.ThrottleRequest")
	proto.RegisterType((*Quotas)(nil), "pb.Quotas")
	proto.RegisterType((*QuotaUsage)(nil), "pb.QuotaUsage")
	proto.RegisterType((*SpaceQuota)(nil), "pb.SpaceQuota")
	proto.RegisterType((*SpaceLimitRequest)(nil), "pb.SpaceLimitRequest")
	proto.RegisterType((*SpaceQuotaStatus)(nil), "pb.SpaceQuotaStatus")
	proto.RegisterType((*SpaceQuotaSnapshot)(nil), "pb.SpaceQuotaSnapshot")
	proto.RegisterEnum("pb.QuotaScope", QuotaScope_name, QuotaScope_value)
	proto.RegisterEnum("pb.ThrottleType", ThrottleType_name, ThrottleType_value)
	proto.RegisterEnum("pb.QuotaType", QuotaType_name, QuotaType_value)
	proto.RegisterEnum("pb.SpaceViolationPolicy", SpaceViolationPolicy_name, SpaceViolationPolicy_value)
}

var fileDescriptor18 = []byte{
	// 766 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0x5d, 0x6f, 0xda, 0x48,
	0x14, 0x95, 0xf9, 0x0a, 0x5c, 0x3b, 0x84, 0x1d, 0x65, 0x25, 0x76, 0xa5, 0xd5, 0x66, 0xad, 0x68,
	0x4b, 0x69, 0x43, 0xa5, 0xbc, 0x54, 0xcd, 0x1b, 0x10, 0xab, 0x20, 0x25, 0x90, 0x8c, 0x4d, 0x2b,
	0x55, 0xaa, 0x2c, 0x13, 0x26, 0x60, 0x09, 0xb0, 0xc1, 0x63, 0xaa, 0xe4, 0x07, 0x54, 0x7d, 0xab,
	0xd4, 0xd7, 0xbe, 0xf4, 0xa7, 0x76, 0xe6, 0x8e, 0x31, 0x24, 0x25, 0x0f, 0x7d, 0x9b, 0x39, 0xf7,
	0xcc, 0x39, 0x67, 0xee, 0x1d, 0x0d, 0xe8, 0xd7, 0x71, 0xc0, 0xbd, 0x46, 0xb8, 0x0c, 0x78, 0x40,
	0x32, 0xe1, 0xf0, 0x6f, 0xbd, 0xd3, 0xf2, 0x22, 0xa6, 0x00, 0xf3, 0xbb, 0x06, 0xe0, 0xf8, 0x33,
	0x36, 0x42, 0x16, 0x79, 0x0e, 0x25, 0x2e, 0x76, 0x6e, 0x3c, 0xf7, 0x79, 0x55, 0x3b, 0xca, 0xd4,
	0xca, 0xa7, 0x46, 0x23, 0x1c, 0x36, 0x24, 0x65, 0x20, 0x30, 0x5a, 0xe4, 0xc9, 0x8a, 0xfc, 0x03,
	0x10, 0x05, 0xb7, 0xdc, 0x9d, 0xfa, 0x33, 0xc1, 0xcd, 0x1c, 0x69, 0xb5, 0x1c, 0x2d, 0x49, 0xe4,
	0x42, 0x02, 0xe4, 0x10, 0xf2, 0xd1, 0xc4, 0x5b, 0xb2, 0x6a, 0x56, 0x54, 0x32, 0x54, 0x6d, 0xc8,
	0x89, 0x40, 0x6f, 0x82, 0x90, 0x55, 0x73, 0x02, 0x2d, 0x9f, 0x96, 0xa5, 0x36, 0x3a, 0xdb, 0x12,
	0x3d, 0xdb, 0xbb, 0x6c, 0xb6, 0x3b, 0xdd, 0x9e, 0x45, 0x15, 0xcb, 0xfc, 0x96, 0x81, 0xa2, 0x33,
	0x11, 0x41, 0xf9, 0x94, 0x91, 0x67, 0xb0, 0xb7, 0x64, 0x0b, 0x77, 0x1e, 0xcf, 0x44, 0x32, 0xad,
	0xa6, 0xab, 0xd3, 0x9b, 0xf0, 0xb4, 0x20, 0xca, 0xbd, 0x78, 0x26, 0x2e, 0x51, 0x94, 0xc4, 0xc8,
	0xbf, 0x67, 0x98, 0xeb, 0x57, 0xa6, 0x14, 0xb2, 0x45, 0x99, 0xbc, 0x80, 0xd2, 0xa7, 0xa5, 0xcf,
	0x19, 0xaa, 0x66, 0x77, 0x72, 0x8b, 0x48, 0x90, 0xba, 0x27, 0x00, 0x8a, 0x8c, 0xca, 0xb9, 0x9d,
	0x6c, 0x25, 0x87, 0xda, 0x18, 0xc3, 0x1b, 0xa1, 0x74, 0xfe, 0xa9, 0x18, 0xde, 0x48, 0x2a, 0x8b,
	0x18, 0x48, 0x45, 0xe1, 0xc2, 0xee, 0x18, 0x92, 0x20, 0x75, 0xcd, 0x09, 0x1c, 0xac, 0x7b, 0x42,
	0xd9, 0x22, 0x66, 0x11, 0x27, 0xc7, 0x90, 0xe3, 0x77, 0xa2, 0xab, 0x1a, 0x76, 0xb5, 0x82, 0x47,
	0x13, 0x8a, 0x23, 0x70, 0x8a, 0x55, 0xf2, 0x0a, 0x74, 0x39, 0xbd, 0x91, 0xbb, 0x90, 0x8a, 0x4f,
	0xb4, 0x06, 0x78, 0xba, 0x36, 0x3f, 0x6b, 0x50, 0xc0, 0x55, 0x44, 0x5e, 0x42, 0x79, 0x78, 0x17,
	0x7a, 0x51, 0xe4, 0x8e, 0xa7, 0xc1, 0xd0, 0x9b, 0x46, 0xe8, 0x55, 0x3c, 0xcb, 0xdf, 0x8a, 0x35,
	0xa3, 0xfb, 0xaa, 0xf8, 0x56, 0xd5, 0x48, 0x0d, 0x8a, 0x3c, 0xf1, 0x4f, 0x6c, 0x8c, 0xed, 0x4c,
	0x34, 0xad, 0x8a, 0xe4, 0xf9, 0x28, 0xf4, 0x6e, 0xd8, 0x76, 0xf3, 0x6d, 0x09, 0xa8, 0x34, 0xaa,
	0x68, 0x1a, 0x00, 0xb8, 0x1f, 0x44, 0xde, 0x98, 0x99, 0x5f, 0xc5, 0x9b, 0xdd, 0x70, 0x1e, 0x3d,
	0x44, 0xed, 0xf1, 0x43, 0x6c, 0x43, 0x65, 0xe5, 0x07, 0x53, 0x8f, 0xfb, 0xc1, 0xdc, 0x0d, 0x83,
	0xa9, 0x7f, 0x73, 0x87, 0x99, 0xca, 0xa7, 0xd5, 0xd4, 0xec, 0xdd, 0x9a, 0x70, 0x85, 0x75, 0x7a,
	0xb0, 0x7a, 0x08, 0x08, 0x0f, 0xf1, 0xb8, 0x66, 0xc1, 0x4a, 0xe5, 0x4c, 0xaf, 0x9d, 0x80, 0xe6,
	0x1b, 0xf8, 0x03, 0x75, 0xd0, 0x71, 0x33, 0x94, 0xbc, 0x6a, 0xb4, 0xb6, 0xfb, 0x6a, 0x58, 0x34,
	0xef, 0xa1, 0xb2, 0x01, 0x6d, 0xee, 0xf1, 0x38, 0xda, 0x19, 0x59, 0xfb, 0xdd, 0xc8, 0xff, 0x81,
	0xe1, 0xcf, 0xdd, 0x14, 0xc5, 0x3b, 0x17, 0xa9, 0xee, 0xcf, 0xd3, 0x93, 0xb2, 0x91, 0x64, 0xcb,
	0x7c, 0xee, 0x85, 0xd1, 0x24, 0xe0, 0xe4, 0x35, 0x18, 0x98, 0xcd, 0x8d, 0x30, 0x4e, 0x92, 0xff,
	0xf0, 0x61, 0x7e, 0x15, 0x95, 0xea, 0x8b, 0xad, 0xdc, 0xff, 0x82, 0xda, 0xba, 0xb1, 0x9c, 0x53,
	0xf2, 0x27, 0xc0, 0x22, 0x9d, 0xdc, 0x86, 0xa0, 0x66, 0x95, 0xdd, 0x22, 0x60, 0xeb, 0xea, 0xff,
	0x27, 0x83, 0xc6, 0xef, 0x80, 0xe8, 0xb0, 0xd7, 0xbe, 0x18, 0xd8, 0x8e, 0x45, 0x2b, 0x9a, 0xdc,
	0x24, 0xbf, 0x43, 0x25, 0x53, 0x5f, 0x81, 0xb1, 0xfd, 0xc0, 0x09, 0x81, 0x32, 0xb5, 0xae, 0x07,
	0x96, 0xed, 0xb8, 0xbd, 0xc1, 0x65, 0x0b, 0x0f, 0x54, 0xc0, 0x58, 0x63, 0x76, 0xf7, 0x83, 0x38,
	0x25, 0x91, 0xf7, 0xb4, 0xeb, 0x58, 0x6b, 0x4e, 0x96, 0x94, 0x01, 0x14, 0x82, 0x8c, 0x1c, 0x39,
	0x00, 0x9d, 0x5a, 0xcd, 0xf3, 0x35, 0x21, 0x4f, 0xf6, 0xa1, 0x84, 0x00, 0xd6, 0x0b, 0xf5, 0x63,
	0x28, 0x61, 0x3e, 0x34, 0x35, 0xc4, 0xe7, 0xd4, 0xa1, 0x7d, 0xc7, 0xb9, 0xb0, 0x84, 0x5d, 0x09,
	0xf2, 0xf6, 0x55, 0xb3, 0x2d, 0xd3, 0x7d, 0x84, 0xc3, 0x5d, 0x33, 0x92, 0x57, 0x38, 0xef, 0xda,
	0xcd, 0x16, 0xf2, 0xff, 0x82, 0x3f, 0x7b, 0x7d, 0x17, 0xdd, 0x6d, 0xb7, 0xdd, 0xbf, 0x14, 0x67,
	0x9d, 0x6e, 0xbf, 0x67, 0x8b, 0x9c, 0xc2, 0x34, 0x2d, 0xa9, 0x90, 0x62, 0xdb, 0xed, 0xd9, 0x16,
	0x75, 0xec, 0x4a, 0xae, 0xd5, 0x84, 0x7a, 0xb0, 0x1c, 0x37, 0x3c, 0x61, 0x31, 0x61, 0x8d, 0x89,
	0x37, 0x0a, 0x82, 0xb0, 0x31, 0x19, 0xa6, 0x9f, 0xfa, 0x30, 0xbe, 0x6d, 0x8c, 0xd9, 0x9c, 0x2d,
	0x3d, 0xce, 0x46, 0x2d, 0xf5, 0xff, 0x5f, 0xc9, 0x42, 0xd4, 0xd1, 0xbe, 0x68, 0xda, 0x0f, 0x4d,
	0xfb, 0x19, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xbe, 0xa6, 0xeb, 0x15, 0x06, 0x00, 0x00,
}
