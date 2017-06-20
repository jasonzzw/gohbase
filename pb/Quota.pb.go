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
func (QuotaScope) EnumDescriptor() ([]byte, []int) { return fileDescriptor20, []int{0} }

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
func (ThrottleType) EnumDescriptor() ([]byte, []int) { return fileDescriptor20, []int{1} }

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
func (QuotaType) EnumDescriptor() ([]byte, []int) { return fileDescriptor20, []int{2} }

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
func (SpaceViolationPolicy) EnumDescriptor() ([]byte, []int) { return fileDescriptor20, []int{3} }

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
func (*TimedQuota) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{0} }

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
func (*Throttle) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{1} }

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
func (*ThrottleRequest) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{2} }

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
func (*Quotas) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{3} }

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
func (*QuotaUsage) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{4} }

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
func (*SpaceQuota) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{5} }

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
func (*SpaceLimitRequest) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{6} }

func (m *SpaceLimitRequest) GetQuota() *SpaceQuota {
	if m != nil {
		return m.Quota
	}
	return nil
}

// Represents the state of a quota on a table. Either the quota is not in violation
// or it is in violation there is a violation policy which should be in effect.
type SpaceQuotaStatus struct {
	ViolationPolicy  *SpaceViolationPolicy `protobuf:"varint,1,opt,name=violation_policy,json=violationPolicy,enum=pb.SpaceViolationPolicy" json:"violation_policy,omitempty"`
	InViolation      *bool                 `protobuf:"varint,2,opt,name=in_violation,json=inViolation" json:"in_violation,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *SpaceQuotaStatus) Reset()                    { *m = SpaceQuotaStatus{} }
func (m *SpaceQuotaStatus) String() string            { return proto.CompactTextString(m) }
func (*SpaceQuotaStatus) ProtoMessage()               {}
func (*SpaceQuotaStatus) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{7} }

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
func (*SpaceQuotaSnapshot) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{8} }

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

type GetSpaceQuotaRegionSizesRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetSpaceQuotaRegionSizesRequest) Reset()         { *m = GetSpaceQuotaRegionSizesRequest{} }
func (m *GetSpaceQuotaRegionSizesRequest) String() string { return proto.CompactTextString(m) }
func (*GetSpaceQuotaRegionSizesRequest) ProtoMessage()    {}
func (*GetSpaceQuotaRegionSizesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor20, []int{9}
}

type GetSpaceQuotaRegionSizesResponse struct {
	Sizes            []*GetSpaceQuotaRegionSizesResponse_RegionSizes `protobuf:"bytes,1,rep,name=sizes" json:"sizes,omitempty"`
	XXX_unrecognized []byte                                          `json:"-"`
}

func (m *GetSpaceQuotaRegionSizesResponse) Reset()         { *m = GetSpaceQuotaRegionSizesResponse{} }
func (m *GetSpaceQuotaRegionSizesResponse) String() string { return proto.CompactTextString(m) }
func (*GetSpaceQuotaRegionSizesResponse) ProtoMessage()    {}
func (*GetSpaceQuotaRegionSizesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor20, []int{10}
}

func (m *GetSpaceQuotaRegionSizesResponse) GetSizes() []*GetSpaceQuotaRegionSizesResponse_RegionSizes {
	if m != nil {
		return m.Sizes
	}
	return nil
}

type GetSpaceQuotaRegionSizesResponse_RegionSizes struct {
	TableName        *TableName `protobuf:"bytes,1,opt,name=table_name,json=tableName" json:"table_name,omitempty"`
	Size             *uint64    `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *GetSpaceQuotaRegionSizesResponse_RegionSizes) Reset() {
	*m = GetSpaceQuotaRegionSizesResponse_RegionSizes{}
}
func (m *GetSpaceQuotaRegionSizesResponse_RegionSizes) String() string {
	return proto.CompactTextString(m)
}
func (*GetSpaceQuotaRegionSizesResponse_RegionSizes) ProtoMessage() {}
func (*GetSpaceQuotaRegionSizesResponse_RegionSizes) Descriptor() ([]byte, []int) {
	return fileDescriptor20, []int{10, 0}
}

func (m *GetSpaceQuotaRegionSizesResponse_RegionSizes) GetTableName() *TableName {
	if m != nil {
		return m.TableName
	}
	return nil
}

func (m *GetSpaceQuotaRegionSizesResponse_RegionSizes) GetSize() uint64 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

type GetSpaceQuotaSnapshotsRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetSpaceQuotaSnapshotsRequest) Reset()                    { *m = GetSpaceQuotaSnapshotsRequest{} }
func (m *GetSpaceQuotaSnapshotsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSpaceQuotaSnapshotsRequest) ProtoMessage()               {}
func (*GetSpaceQuotaSnapshotsRequest) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{11} }

type GetSpaceQuotaSnapshotsResponse struct {
	Snapshots        []*GetSpaceQuotaSnapshotsResponse_TableQuotaSnapshot `protobuf:"bytes,1,rep,name=snapshots" json:"snapshots,omitempty"`
	XXX_unrecognized []byte                                               `json:"-"`
}

func (m *GetSpaceQuotaSnapshotsResponse) Reset()         { *m = GetSpaceQuotaSnapshotsResponse{} }
func (m *GetSpaceQuotaSnapshotsResponse) String() string { return proto.CompactTextString(m) }
func (*GetSpaceQuotaSnapshotsResponse) ProtoMessage()    {}
func (*GetSpaceQuotaSnapshotsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor20, []int{12}
}

func (m *GetSpaceQuotaSnapshotsResponse) GetSnapshots() []*GetSpaceQuotaSnapshotsResponse_TableQuotaSnapshot {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

// Cannot use TableName as a map key, do the repeated nested message by hand.
type GetSpaceQuotaSnapshotsResponse_TableQuotaSnapshot struct {
	TableName        *TableName          `protobuf:"bytes,1,opt,name=table_name,json=tableName" json:"table_name,omitempty"`
	Snapshot         *SpaceQuotaSnapshot `protobuf:"bytes,2,opt,name=snapshot" json:"snapshot,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *GetSpaceQuotaSnapshotsResponse_TableQuotaSnapshot) Reset() {
	*m = GetSpaceQuotaSnapshotsResponse_TableQuotaSnapshot{}
}
func (m *GetSpaceQuotaSnapshotsResponse_TableQuotaSnapshot) String() string {
	return proto.CompactTextString(m)
}
func (*GetSpaceQuotaSnapshotsResponse_TableQuotaSnapshot) ProtoMessage() {}
func (*GetSpaceQuotaSnapshotsResponse_TableQuotaSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor20, []int{12, 0}
}

func (m *GetSpaceQuotaSnapshotsResponse_TableQuotaSnapshot) GetTableName() *TableName {
	if m != nil {
		return m.TableName
	}
	return nil
}

func (m *GetSpaceQuotaSnapshotsResponse_TableQuotaSnapshot) GetSnapshot() *SpaceQuotaSnapshot {
	if m != nil {
		return m.Snapshot
	}
	return nil
}

type GetQuotaStatesRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetQuotaStatesRequest) Reset()                    { *m = GetQuotaStatesRequest{} }
func (m *GetQuotaStatesRequest) String() string            { return proto.CompactTextString(m) }
func (*GetQuotaStatesRequest) ProtoMessage()               {}
func (*GetQuotaStatesRequest) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{13} }

type GetQuotaStatesResponse struct {
	TableSnapshots   []*GetQuotaStatesResponse_TableQuotaSnapshot     `protobuf:"bytes,1,rep,name=table_snapshots,json=tableSnapshots" json:"table_snapshots,omitempty"`
	NsSnapshots      []*GetQuotaStatesResponse_NamespaceQuotaSnapshot `protobuf:"bytes,2,rep,name=ns_snapshots,json=nsSnapshots" json:"ns_snapshots,omitempty"`
	XXX_unrecognized []byte                                           `json:"-"`
}

func (m *GetQuotaStatesResponse) Reset()                    { *m = GetQuotaStatesResponse{} }
func (m *GetQuotaStatesResponse) String() string            { return proto.CompactTextString(m) }
func (*GetQuotaStatesResponse) ProtoMessage()               {}
func (*GetQuotaStatesResponse) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{14} }

func (m *GetQuotaStatesResponse) GetTableSnapshots() []*GetQuotaStatesResponse_TableQuotaSnapshot {
	if m != nil {
		return m.TableSnapshots
	}
	return nil
}

func (m *GetQuotaStatesResponse) GetNsSnapshots() []*GetQuotaStatesResponse_NamespaceQuotaSnapshot {
	if m != nil {
		return m.NsSnapshots
	}
	return nil
}

type GetQuotaStatesResponse_TableQuotaSnapshot struct {
	TableName        *TableName          `protobuf:"bytes,1,opt,name=table_name,json=tableName" json:"table_name,omitempty"`
	Snapshot         *SpaceQuotaSnapshot `protobuf:"bytes,2,opt,name=snapshot" json:"snapshot,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *GetQuotaStatesResponse_TableQuotaSnapshot) Reset() {
	*m = GetQuotaStatesResponse_TableQuotaSnapshot{}
}
func (m *GetQuotaStatesResponse_TableQuotaSnapshot) String() string { return proto.CompactTextString(m) }
func (*GetQuotaStatesResponse_TableQuotaSnapshot) ProtoMessage()    {}
func (*GetQuotaStatesResponse_TableQuotaSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor20, []int{14, 0}
}

func (m *GetQuotaStatesResponse_TableQuotaSnapshot) GetTableName() *TableName {
	if m != nil {
		return m.TableName
	}
	return nil
}

func (m *GetQuotaStatesResponse_TableQuotaSnapshot) GetSnapshot() *SpaceQuotaSnapshot {
	if m != nil {
		return m.Snapshot
	}
	return nil
}

type GetQuotaStatesResponse_NamespaceQuotaSnapshot struct {
	Namespace        *string             `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Snapshot         *SpaceQuotaSnapshot `protobuf:"bytes,2,opt,name=snapshot" json:"snapshot,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *GetQuotaStatesResponse_NamespaceQuotaSnapshot) Reset() {
	*m = GetQuotaStatesResponse_NamespaceQuotaSnapshot{}
}
func (m *GetQuotaStatesResponse_NamespaceQuotaSnapshot) String() string {
	return proto.CompactTextString(m)
}
func (*GetQuotaStatesResponse_NamespaceQuotaSnapshot) ProtoMessage() {}
func (*GetQuotaStatesResponse_NamespaceQuotaSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor20, []int{14, 1}
}

func (m *GetQuotaStatesResponse_NamespaceQuotaSnapshot) GetNamespace() string {
	if m != nil && m.Namespace != nil {
		return *m.Namespace
	}
	return ""
}

func (m *GetQuotaStatesResponse_NamespaceQuotaSnapshot) GetSnapshot() *SpaceQuotaSnapshot {
	if m != nil {
		return m.Snapshot
	}
	return nil
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
	proto.RegisterType((*GetSpaceQuotaRegionSizesRequest)(nil), "pb.GetSpaceQuotaRegionSizesRequest")
	proto.RegisterType((*GetSpaceQuotaRegionSizesResponse)(nil), "pb.GetSpaceQuotaRegionSizesResponse")
	proto.RegisterType((*GetSpaceQuotaRegionSizesResponse_RegionSizes)(nil), "pb.GetSpaceQuotaRegionSizesResponse.RegionSizes")
	proto.RegisterType((*GetSpaceQuotaSnapshotsRequest)(nil), "pb.GetSpaceQuotaSnapshotsRequest")
	proto.RegisterType((*GetSpaceQuotaSnapshotsResponse)(nil), "pb.GetSpaceQuotaSnapshotsResponse")
	proto.RegisterType((*GetSpaceQuotaSnapshotsResponse_TableQuotaSnapshot)(nil), "pb.GetSpaceQuotaSnapshotsResponse.TableQuotaSnapshot")
	proto.RegisterType((*GetQuotaStatesRequest)(nil), "pb.GetQuotaStatesRequest")
	proto.RegisterType((*GetQuotaStatesResponse)(nil), "pb.GetQuotaStatesResponse")
	proto.RegisterType((*GetQuotaStatesResponse_TableQuotaSnapshot)(nil), "pb.GetQuotaStatesResponse.TableQuotaSnapshot")
	proto.RegisterType((*GetQuotaStatesResponse_NamespaceQuotaSnapshot)(nil), "pb.GetQuotaStatesResponse.NamespaceQuotaSnapshot")
	proto.RegisterEnum("pb.QuotaScope", QuotaScope_name, QuotaScope_value)
	proto.RegisterEnum("pb.ThrottleType", ThrottleType_name, ThrottleType_value)
	proto.RegisterEnum("pb.QuotaType", QuotaType_name, QuotaType_value)
	proto.RegisterEnum("pb.SpaceViolationPolicy", SpaceViolationPolicy_name, SpaceViolationPolicy_value)
}

var fileDescriptor20 = []byte{
	// 1010 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x56, 0x5f, 0x6f, 0xe3, 0x44,
	0x10, 0x97, 0xf3, 0xa7, 0x8d, 0x27, 0x69, 0x1a, 0x56, 0xbd, 0x52, 0x2a, 0x8e, 0xf6, 0xac, 0x13,
	0x94, 0x42, 0x03, 0x57, 0x09, 0x21, 0xee, 0xad, 0xed, 0x85, 0x6b, 0xa4, 0x36, 0xed, 0xad, 0xdd,
	0x43, 0x42, 0x42, 0x96, 0xd3, 0x6c, 0x1b, 0xa3, 0xc4, 0x76, 0xb2, 0x4e, 0x50, 0xef, 0x03, 0x20,
	0xde, 0x90, 0x78, 0xe5, 0x85, 0x67, 0xbe, 0x02, 0x9f, 0x89, 0xef, 0xc0, 0xec, 0xac, 0xed, 0xb8,
	0x49, 0x0a, 0x3a, 0x5e, 0xee, 0x6d, 0xf7, 0x37, 0xbf, 0x99, 0xf9, 0xcd, 0xec, 0x78, 0xbd, 0x50,
	0x7d, 0x35, 0x09, 0x63, 0xaf, 0x19, 0x8d, 0xc3, 0x38, 0x64, 0x85, 0xa8, 0xbb, 0x5d, 0x3d, 0x3d,
	0xf6, 0xa4, 0xd0, 0x80, 0xf5, 0xbb, 0x01, 0xe0, 0xf8, 0x43, 0xd1, 0x23, 0x16, 0xfb, 0x14, 0xcc,
	0x18, 0x77, 0xee, 0x24, 0xf0, 0xe3, 0x2d, 0x63, 0xb7, 0xb0, 0x57, 0x3f, 0xac, 0x35, 0xa3, 0x6e,
	0x53, 0x51, 0xae, 0x10, 0xe3, 0x95, 0x38, 0x59, 0xb1, 0xc7, 0x00, 0x32, 0xbc, 0x89, 0xdd, 0x81,
	0x3f, 0x44, 0x6e, 0x61, 0xd7, 0xd8, 0x2b, 0x71, 0x53, 0x21, 0x67, 0x0a, 0x60, 0x1b, 0x50, 0x96,
	0x7d, 0x6f, 0x2c, 0xb6, 0x8a, 0x68, 0x29, 0x70, 0xbd, 0x61, 0x07, 0x88, 0x5e, 0x87, 0x91, 0xd8,
	0x2a, 0x21, 0x5a, 0x3f, 0xac, 0xab, 0xd8, 0x94, 0xd9, 0x56, 0xe8, 0xf3, 0xd5, 0xf3, 0xa3, 0x93,
	0xd3, 0x76, 0xa7, 0xc5, 0x35, 0xcb, 0xfa, 0xad, 0x00, 0x15, 0xa7, 0x8f, 0x42, 0xe3, 0x81, 0x60,
	0x9f, 0xc0, 0xea, 0x58, 0x8c, 0xdc, 0x60, 0x32, 0x44, 0x65, 0xc6, 0x5e, 0x55, 0x7b, 0xcf, 0xc4,
	0xf3, 0x15, 0x34, 0x77, 0x26, 0x43, 0x2c, 0xa2, 0xa2, 0x88, 0xd2, 0x7f, 0x23, 0x48, 0xd7, 0x22,
	0x53, 0x05, 0xb2, 0xd1, 0xcc, 0x3e, 0x03, 0xf3, 0xa7, 0xb1, 0x1f, 0x0b, 0x8a, 0x5a, 0x5c, 0xca,
	0xad, 0x10, 0x41, 0xc5, 0x3d, 0x00, 0xd0, 0x64, 0x8a, 0x5c, 0x5a, 0xca, 0xd6, 0xe1, 0x28, 0x36,
	0xc9, 0xf0, 0x7a, 0x14, 0xba, 0xfc, 0x90, 0x0c, 0xaf, 0xa7, 0x22, 0xa3, 0x0c, 0xa2, 0x52, 0xe0,
	0x95, 0xe5, 0x32, 0x14, 0x41, 0xc5, 0xb5, 0xfa, 0xb0, 0x9e, 0xf6, 0x84, 0x8b, 0xd1, 0x44, 0xc8,
	0x98, 0x3d, 0x85, 0x52, 0x7c, 0x87, 0x5d, 0x35, 0xa8, 0xab, 0x0d, 0x72, 0x4d, 0x28, 0x0e, 0xe2,
	0x9c, 0xac, 0xec, 0x0b, 0xa8, 0xaa, 0xd3, 0xeb, 0xb9, 0x23, 0x15, 0xf1, 0x81, 0xd6, 0x40, 0x9c,
	0xad, 0xad, 0x9f, 0x0d, 0x58, 0xa1, 0x95, 0x64, 0x9f, 0x43, 0xbd, 0x7b, 0x17, 0x79, 0x52, 0xba,
	0xb7, 0x83, 0xb0, 0xeb, 0x0d, 0x24, 0xe5, 0xaa, 0x3c, 0x2f, 0xdf, 0xe0, 0x5a, 0xf0, 0x35, 0x6d,
	0x7c, 0xa9, 0x6d, 0x6c, 0x0f, 0x2a, 0x71, 0x92, 0x3f, 0x49, 0x53, 0xcb, 0x6b, 0xe2, 0x99, 0x15,
	0x95, 0x97, 0x65, 0xe4, 0x5d, 0x8b, 0x7c, 0xf3, 0x6d, 0x05, 0x68, 0x35, 0xda, 0x68, 0xd5, 0x00,
	0x68, 0x7f, 0x25, 0xbd, 0x5b, 0x61, 0xfd, 0x8a, 0x33, 0x3b, 0xe3, 0xcc, 0x0d, 0xa2, 0x31, 0x3f,
	0x88, 0x27, 0xd0, 0x98, 0xfa, 0xe1, 0xc0, 0x8b, 0xfd, 0x30, 0x70, 0xa3, 0x70, 0xe0, 0x5f, 0xdf,
	0x91, 0xa6, 0xfa, 0xe1, 0x56, 0x96, 0xec, 0x75, 0x4a, 0xb8, 0x24, 0x3b, 0x5f, 0x9f, 0xde, 0x07,
	0x30, 0x07, 0x0e, 0xd7, 0x30, 0x9c, 0x6a, 0x9d, 0x59, 0xd9, 0x09, 0x68, 0x7d, 0x03, 0xef, 0x51,
	0x1c, 0xca, 0x38, 0x3b, 0x94, 0xb2, 0x6e, 0xb4, 0xb1, 0xbc, 0x34, 0x32, 0x5a, 0x6f, 0xa0, 0x31,
	0x03, 0xed, 0xd8, 0x8b, 0x27, 0x72, 0xa9, 0x64, 0xe3, 0x6d, 0x25, 0x3f, 0x81, 0x9a, 0x1f, 0xb8,
	0x19, 0x4a, 0x35, 0x57, 0x78, 0xd5, 0x0f, 0x32, 0x4f, 0xd5, 0x48, 0x96, 0x4b, 0x1e, 0x78, 0x91,
	0xec, 0x87, 0x31, 0xfb, 0x1a, 0x6a, 0xa4, 0xcd, 0x95, 0x24, 0x27, 0xd1, 0xbf, 0x71, 0x5f, 0xbf,
	0x96, 0xca, 0xab, 0xa3, 0x9c, 0xee, 0x1d, 0xd0, 0x5b, 0x77, 0xa2, 0xce, 0x29, 0xb9, 0x13, 0x60,
	0x94, 0x9d, 0xdc, 0x8c, 0xa0, 0xcf, 0xaa, 0x98, 0x23, 0x50, 0xeb, 0xac, 0x27, 0xb0, 0xf3, 0x52,
	0xc4, 0xb9, 0x2e, 0x89, 0x5b, 0x14, 0xaa, 0xc6, 0x5e, 0x26, 0x6d, 0xb5, 0xfe, 0x32, 0x60, 0xf7,
	0x61, 0x8e, 0x8c, 0xc2, 0x40, 0x0a, 0xf6, 0x2d, 0x8e, 0x95, 0x02, 0x50, 0x7b, 0x11, 0xb5, 0x7f,
	0xa9, 0xb4, 0xff, 0x97, 0x53, 0x33, 0x8f, 0x69, 0xf7, 0xed, 0x0b, 0xa8, 0xe6, 0x50, 0xfc, 0x0a,
	0x20, 0xf6, 0xba, 0x03, 0xbc, 0x2e, 0xbc, 0xa1, 0x48, 0xfa, 0xb2, 0x46, 0x93, 0xad, 0xd0, 0x0e,
	0x82, 0xdc, 0x8c, 0xd3, 0x25, 0x63, 0x50, 0xca, 0xee, 0xa0, 0x12, 0xa7, 0xb5, 0xb5, 0x03, 0x8f,
	0xef, 0xe9, 0x48, 0x9b, 0x9e, 0x95, 0xf7, 0xb7, 0x01, 0x1f, 0x3d, 0xc4, 0x48, 0x8a, 0xb3, 0xc1,
	0x94, 0x29, 0x98, 0x14, 0xf8, 0xd5, 0x42, 0x81, 0x0b, 0x6e, 0x5a, 0xe3, 0x3d, 0x1b, 0x9f, 0xc5,
	0xd9, 0x9e, 0x02, 0x5b, 0x24, 0xbc, 0x65, 0xc1, 0x87, 0x50, 0x49, 0x03, 0x26, 0x9f, 0xfd, 0xe6,
	0xdc, 0xd0, 0xa4, 0x89, 0x33, 0x9e, 0xf5, 0x3e, 0x3c, 0x42, 0xdd, 0xd9, 0x48, 0xcd, 0xce, 0xf9,
	0xcf, 0x22, 0x6c, 0xce, 0x5b, 0x92, 0x06, 0xbc, 0x86, 0x75, 0xad, 0x6a, 0xbe, 0x0d, 0x07, 0x49,
	0x1b, 0x96, 0x38, 0x2d, 0x2b, 0xbf, 0x4e, 0x51, 0xb2, 0x4e, 0x31, 0x07, 0x6a, 0x81, 0xcc, 0x05,
	0x2d, 0x50, 0xd0, 0x67, 0xff, 0x12, 0x54, 0x95, 0x2d, 0x17, 0xcb, 0xab, 0x06, 0xd2, 0x7e, 0xd7,
	0x9d, 0xdd, 0xfe, 0x11, 0x36, 0x97, 0xcb, 0x63, 0x1f, 0x82, 0x19, 0xa4, 0x16, 0x4a, 0x6d, 0xf2,
	0x19, 0xf0, 0x7f, 0x72, 0xed, 0x7f, 0x9c, 0x5c, 0xd0, 0xf4, 0x1b, 0x67, 0x55, 0x58, 0x3d, 0x39,
	0xbb, 0xb2, 0x9d, 0x16, 0x6f, 0x18, 0x6a, 0x93, 0xfc, 0xd5, 0x1b, 0x85, 0xfd, 0x29, 0xd4, 0xf2,
	0x3f, 0x26, 0xfc, 0x44, 0xea, 0xbc, 0xf5, 0xea, 0xaa, 0x65, 0x3b, 0x6e, 0xe7, 0xea, 0xfc, 0x98,
	0x1c, 0x1a, 0x50, 0x4b, 0x31, 0xbb, 0xfd, 0x3d, 0x7a, 0x29, 0xe4, 0x3b, 0xde, 0x76, 0x5a, 0x29,
	0xa7, 0xc8, 0xea, 0x00, 0x1a, 0x21, 0x46, 0x89, 0xad, 0xe3, 0x77, 0xda, 0x3a, 0x7a, 0x91, 0x12,
	0xca, 0x6c, 0x0d, 0x4c, 0x02, 0xc8, 0xbe, 0xb2, 0xff, 0x14, 0x4c, 0xd2, 0x47, 0x49, 0x6b, 0xf8,
	0xa8, 0x38, 0xe5, 0x17, 0x8e, 0x73, 0xd6, 0xc2, 0x74, 0x26, 0x94, 0xed, 0xcb, 0xa3, 0x13, 0xa5,
	0xee, 0x07, 0xd8, 0x58, 0x76, 0xb7, 0xaa, 0x12, 0x5e, 0xb4, 0xed, 0xa3, 0x63, 0xe2, 0x7f, 0x00,
	0x8f, 0x3a, 0x17, 0x2e, 0x65, 0xb7, 0xdd, 0x93, 0x8b, 0x73, 0xf4, 0x75, 0xda, 0x17, 0x1d, 0x1b,
	0x75, 0x62, 0xd2, 0xcc, 0xa4, 0x45, 0xe2, 0xb6, 0xdd, 0xb1, 0x5b, 0xdc, 0xb1, 0x1b, 0xa5, 0xe3,
	0x53, 0x78, 0x16, 0x8e, 0x6f, 0x9b, 0x1e, 0xa6, 0xe8, 0x8b, 0x66, 0xdf, 0xeb, 0x85, 0x61, 0xd4,
	0xec, 0x77, 0xd5, 0x63, 0x0c, 0x5f, 0x47, 0x3d, 0xd1, 0xd3, 0x6f, 0xb2, 0xee, 0xe4, 0xa6, 0x79,
	0x2b, 0x02, 0x31, 0xc6, 0x49, 0xeb, 0x1d, 0xeb, 0xe7, 0xdb, 0xa5, 0x32, 0xc8, 0x53, 0xe3, 0x17,
	0xc3, 0xf8, 0xc3, 0x30, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xca, 0x54, 0x68, 0xf8, 0xd4, 0x09,
	0x00, 0x00,
}
