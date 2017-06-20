// Code generated by protoc-gen-go.
// source: ZooKeeper.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SplitLogTask_State int32

const (
	SplitLogTask_UNASSIGNED SplitLogTask_State = 0
	SplitLogTask_OWNED      SplitLogTask_State = 1
	SplitLogTask_RESIGNED   SplitLogTask_State = 2
	SplitLogTask_DONE       SplitLogTask_State = 3
	SplitLogTask_ERR        SplitLogTask_State = 4
)

var SplitLogTask_State_name = map[int32]string{
	0: "UNASSIGNED",
	1: "OWNED",
	2: "RESIGNED",
	3: "DONE",
	4: "ERR",
}
var SplitLogTask_State_value = map[string]int32{
	"UNASSIGNED": 0,
	"OWNED":      1,
	"RESIGNED":   2,
	"DONE":       3,
	"ERR":        4,
}

func (x SplitLogTask_State) Enum() *SplitLogTask_State {
	p := new(SplitLogTask_State)
	*p = x
	return p
}
func (x SplitLogTask_State) String() string {
	return proto.EnumName(SplitLogTask_State_name, int32(x))
}
func (x *SplitLogTask_State) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SplitLogTask_State_value, data, "SplitLogTask_State")
	if err != nil {
		return err
	}
	*x = SplitLogTask_State(value)
	return nil
}
func (SplitLogTask_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor26, []int{3, 0} }

type SplitLogTask_RecoveryMode int32

const (
	SplitLogTask_UNKNOWN       SplitLogTask_RecoveryMode = 0
	SplitLogTask_LOG_SPLITTING SplitLogTask_RecoveryMode = 1
	SplitLogTask_LOG_REPLAY    SplitLogTask_RecoveryMode = 2
)

var SplitLogTask_RecoveryMode_name = map[int32]string{
	0: "UNKNOWN",
	1: "LOG_SPLITTING",
	2: "LOG_REPLAY",
}
var SplitLogTask_RecoveryMode_value = map[string]int32{
	"UNKNOWN":       0,
	"LOG_SPLITTING": 1,
	"LOG_REPLAY":    2,
}

func (x SplitLogTask_RecoveryMode) Enum() *SplitLogTask_RecoveryMode {
	p := new(SplitLogTask_RecoveryMode)
	*p = x
	return p
}
func (x SplitLogTask_RecoveryMode) String() string {
	return proto.EnumName(SplitLogTask_RecoveryMode_name, int32(x))
}
func (x *SplitLogTask_RecoveryMode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SplitLogTask_RecoveryMode_value, data, "SplitLogTask_RecoveryMode")
	if err != nil {
		return err
	}
	*x = SplitLogTask_RecoveryMode(value)
	return nil
}
func (SplitLogTask_RecoveryMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor26, []int{3, 1}
}

// Table's current state
type DeprecatedTableState_State int32

const (
	DeprecatedTableState_ENABLED   DeprecatedTableState_State = 0
	DeprecatedTableState_DISABLED  DeprecatedTableState_State = 1
	DeprecatedTableState_DISABLING DeprecatedTableState_State = 2
	DeprecatedTableState_ENABLING  DeprecatedTableState_State = 3
)

var DeprecatedTableState_State_name = map[int32]string{
	0: "ENABLED",
	1: "DISABLED",
	2: "DISABLING",
	3: "ENABLING",
}
var DeprecatedTableState_State_value = map[string]int32{
	"ENABLED":   0,
	"DISABLED":  1,
	"DISABLING": 2,
	"ENABLING":  3,
}

func (x DeprecatedTableState_State) Enum() *DeprecatedTableState_State {
	p := new(DeprecatedTableState_State)
	*p = x
	return p
}
func (x DeprecatedTableState_State) String() string {
	return proto.EnumName(DeprecatedTableState_State_name, int32(x))
}
func (x *DeprecatedTableState_State) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DeprecatedTableState_State_value, data, "DeprecatedTableState_State")
	if err != nil {
		return err
	}
	*x = DeprecatedTableState_State(value)
	return nil
}
func (DeprecatedTableState_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor26, []int{4, 0}
}

type ReplicationState_State int32

const (
	ReplicationState_ENABLED  ReplicationState_State = 0
	ReplicationState_DISABLED ReplicationState_State = 1
)

var ReplicationState_State_name = map[int32]string{
	0: "ENABLED",
	1: "DISABLED",
}
var ReplicationState_State_value = map[string]int32{
	"ENABLED":  0,
	"DISABLED": 1,
}

func (x ReplicationState_State) Enum() *ReplicationState_State {
	p := new(ReplicationState_State)
	*p = x
	return p
}
func (x ReplicationState_State) String() string {
	return proto.EnumName(ReplicationState_State_name, int32(x))
}
func (x *ReplicationState_State) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ReplicationState_State_value, data, "ReplicationState_State")
	if err != nil {
		return err
	}
	*x = ReplicationState_State(value)
	return nil
}
func (ReplicationState_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor26, []int{7, 0} }

// *
// Content of the meta-region-server znode.
type MetaRegionServer struct {
	// The ServerName hosting the meta region currently, or destination server,
	// if meta region is in transition.
	Server *ServerName `protobuf:"bytes,1,req,name=server" json:"server,omitempty"`
	// The major version of the rpc the server speaks.  This is used so that
	// clients connecting to the cluster can have prior knowledge of what version
	// to send to a RegionServer.  AsyncHBase will use this to detect versions.
	RpcVersion *uint32 `protobuf:"varint,2,opt,name=rpc_version,json=rpcVersion" json:"rpc_version,omitempty"`
	// State of the region transition. OPEN means fully operational 'hbase:meta'
	State            *RegionState_State `protobuf:"varint,3,opt,name=state,enum=pb.RegionState_State" json:"state,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *MetaRegionServer) Reset()                    { *m = MetaRegionServer{} }
func (m *MetaRegionServer) String() string            { return proto.CompactTextString(m) }
func (*MetaRegionServer) ProtoMessage()               {}
func (*MetaRegionServer) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{0} }

func (m *MetaRegionServer) GetServer() *ServerName {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *MetaRegionServer) GetRpcVersion() uint32 {
	if m != nil && m.RpcVersion != nil {
		return *m.RpcVersion
	}
	return 0
}

func (m *MetaRegionServer) GetState() RegionState_State {
	if m != nil && m.State != nil {
		return *m.State
	}
	return RegionState_OFFLINE
}

// *
// Content of the master znode.
type Master struct {
	// The ServerName of the current Master
	Master *ServerName `protobuf:"bytes,1,req,name=master" json:"master,omitempty"`
	// Major RPC version so that clients can know what version the master can accept.
	RpcVersion       *uint32 `protobuf:"varint,2,opt,name=rpc_version,json=rpcVersion" json:"rpc_version,omitempty"`
	InfoPort         *uint32 `protobuf:"varint,3,opt,name=info_port,json=infoPort" json:"info_port,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Master) Reset()                    { *m = Master{} }
func (m *Master) String() string            { return proto.CompactTextString(m) }
func (*Master) ProtoMessage()               {}
func (*Master) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{1} }

func (m *Master) GetMaster() *ServerName {
	if m != nil {
		return m.Master
	}
	return nil
}

func (m *Master) GetRpcVersion() uint32 {
	if m != nil && m.RpcVersion != nil {
		return *m.RpcVersion
	}
	return 0
}

func (m *Master) GetInfoPort() uint32 {
	if m != nil && m.InfoPort != nil {
		return *m.InfoPort
	}
	return 0
}

// *
// Content of the '/hbase/running', cluster state, znode.
type ClusterUp struct {
	// If this znode is present, cluster is up.  Currently
	// the data is cluster start_date.
	StartDate        *string `protobuf:"bytes,1,req,name=start_date,json=startDate" json:"start_date,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ClusterUp) Reset()                    { *m = ClusterUp{} }
func (m *ClusterUp) String() string            { return proto.CompactTextString(m) }
func (*ClusterUp) ProtoMessage()               {}
func (*ClusterUp) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{2} }

func (m *ClusterUp) GetStartDate() string {
	if m != nil && m.StartDate != nil {
		return *m.StartDate
	}
	return ""
}

// *
// WAL SplitLog directory znodes have this for content.  Used doing distributed
// WAL splitting.  Holds current state and name of server that originated split.
type SplitLogTask struct {
	State            *SplitLogTask_State        `protobuf:"varint,1,req,name=state,enum=pb.SplitLogTask_State" json:"state,omitempty"`
	ServerName       *ServerName                `protobuf:"bytes,2,req,name=server_name,json=serverName" json:"server_name,omitempty"`
	Mode             *SplitLogTask_RecoveryMode `protobuf:"varint,3,opt,name=mode,enum=pb.SplitLogTask_RecoveryMode,def=0" json:"mode,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *SplitLogTask) Reset()                    { *m = SplitLogTask{} }
func (m *SplitLogTask) String() string            { return proto.CompactTextString(m) }
func (*SplitLogTask) ProtoMessage()               {}
func (*SplitLogTask) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{3} }

const Default_SplitLogTask_Mode SplitLogTask_RecoveryMode = SplitLogTask_UNKNOWN

func (m *SplitLogTask) GetState() SplitLogTask_State {
	if m != nil && m.State != nil {
		return *m.State
	}
	return SplitLogTask_UNASSIGNED
}

func (m *SplitLogTask) GetServerName() *ServerName {
	if m != nil {
		return m.ServerName
	}
	return nil
}

func (m *SplitLogTask) GetMode() SplitLogTask_RecoveryMode {
	if m != nil && m.Mode != nil {
		return *m.Mode
	}
	return Default_SplitLogTask_Mode
}

// *
// The znode that holds state of table.
// Deprected, table state is stored in table descriptor on HDFS.
type DeprecatedTableState struct {
	// This is the table's state.  If no znode for a table,
	// its state is presumed enabled.  See o.a.h.h.zookeeper.ZKTable class
	// for more.
	State            *DeprecatedTableState_State `protobuf:"varint,1,req,name=state,enum=pb.DeprecatedTableState_State,def=0" json:"state,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *DeprecatedTableState) Reset()                    { *m = DeprecatedTableState{} }
func (m *DeprecatedTableState) String() string            { return proto.CompactTextString(m) }
func (*DeprecatedTableState) ProtoMessage()               {}
func (*DeprecatedTableState) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{4} }

const Default_DeprecatedTableState_State DeprecatedTableState_State = DeprecatedTableState_ENABLED

func (m *DeprecatedTableState) GetState() DeprecatedTableState_State {
	if m != nil && m.State != nil {
		return *m.State
	}
	return Default_DeprecatedTableState_State
}

type TableCF struct {
	TableName        *TableName `protobuf:"bytes,1,opt,name=table_name,json=tableName" json:"table_name,omitempty"`
	Families         [][]byte   `protobuf:"bytes,2,rep,name=families" json:"families,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *TableCF) Reset()                    { *m = TableCF{} }
func (m *TableCF) String() string            { return proto.CompactTextString(m) }
func (*TableCF) ProtoMessage()               {}
func (*TableCF) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{5} }

func (m *TableCF) GetTableName() *TableName {
	if m != nil {
		return m.TableName
	}
	return nil
}

func (m *TableCF) GetFamilies() [][]byte {
	if m != nil {
		return m.Families
	}
	return nil
}

// *
// Used by replication. Holds a replication peer key.
type ReplicationPeer struct {
	// clusterkey is the concatenation of the slave cluster's
	// hbase.zookeeper.quorum:hbase.zookeeper.property.clientPort:zookeeper.znode.parent
	Clusterkey              *string           `protobuf:"bytes,1,req,name=clusterkey" json:"clusterkey,omitempty"`
	ReplicationEndpointImpl *string           `protobuf:"bytes,2,opt,name=replicationEndpointImpl" json:"replicationEndpointImpl,omitempty"`
	Data                    []*BytesBytesPair `protobuf:"bytes,3,rep,name=data" json:"data,omitempty"`
	Configuration           []*NameStringPair `protobuf:"bytes,4,rep,name=configuration" json:"configuration,omitempty"`
	TableCfs                []*TableCF        `protobuf:"bytes,5,rep,name=table_cfs,json=tableCfs" json:"table_cfs,omitempty"`
	Namespaces              [][]byte          `protobuf:"bytes,6,rep,name=namespaces" json:"namespaces,omitempty"`
	Bandwidth               *int64            `protobuf:"varint,7,opt,name=bandwidth" json:"bandwidth,omitempty"`
	XXX_unrecognized        []byte            `json:"-"`
}

func (m *ReplicationPeer) Reset()                    { *m = ReplicationPeer{} }
func (m *ReplicationPeer) String() string            { return proto.CompactTextString(m) }
func (*ReplicationPeer) ProtoMessage()               {}
func (*ReplicationPeer) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{6} }

func (m *ReplicationPeer) GetClusterkey() string {
	if m != nil && m.Clusterkey != nil {
		return *m.Clusterkey
	}
	return ""
}

func (m *ReplicationPeer) GetReplicationEndpointImpl() string {
	if m != nil && m.ReplicationEndpointImpl != nil {
		return *m.ReplicationEndpointImpl
	}
	return ""
}

func (m *ReplicationPeer) GetData() []*BytesBytesPair {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ReplicationPeer) GetConfiguration() []*NameStringPair {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *ReplicationPeer) GetTableCfs() []*TableCF {
	if m != nil {
		return m.TableCfs
	}
	return nil
}

func (m *ReplicationPeer) GetNamespaces() [][]byte {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func (m *ReplicationPeer) GetBandwidth() int64 {
	if m != nil && m.Bandwidth != nil {
		return *m.Bandwidth
	}
	return 0
}

// *
// Used by replication. Holds whether enabled or disabled
type ReplicationState struct {
	State            *ReplicationState_State `protobuf:"varint,1,req,name=state,enum=pb.ReplicationState_State" json:"state,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ReplicationState) Reset()                    { *m = ReplicationState{} }
func (m *ReplicationState) String() string            { return proto.CompactTextString(m) }
func (*ReplicationState) ProtoMessage()               {}
func (*ReplicationState) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{7} }

func (m *ReplicationState) GetState() ReplicationState_State {
	if m != nil && m.State != nil {
		return *m.State
	}
	return ReplicationState_ENABLED
}

// *
// Used by replication. Holds the current position in an WAL file.
type ReplicationHLogPosition struct {
	Position         *int64 `protobuf:"varint,1,req,name=position" json:"position,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ReplicationHLogPosition) Reset()                    { *m = ReplicationHLogPosition{} }
func (m *ReplicationHLogPosition) String() string            { return proto.CompactTextString(m) }
func (*ReplicationHLogPosition) ProtoMessage()               {}
func (*ReplicationHLogPosition) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{8} }

func (m *ReplicationHLogPosition) GetPosition() int64 {
	if m != nil && m.Position != nil {
		return *m.Position
	}
	return 0
}

// *
// State of the switch.
type SwitchState struct {
	Enabled          *bool  `protobuf:"varint,1,opt,name=enabled" json:"enabled,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SwitchState) Reset()                    { *m = SwitchState{} }
func (m *SwitchState) String() string            { return proto.CompactTextString(m) }
func (*SwitchState) ProtoMessage()               {}
func (*SwitchState) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{9} }

func (m *SwitchState) GetEnabled() bool {
	if m != nil && m.Enabled != nil {
		return *m.Enabled
	}
	return false
}

func init() {
	proto.RegisterType((*MetaRegionServer)(nil), "pb.MetaRegionServer")
	proto.RegisterType((*Master)(nil), "pb.Master")
	proto.RegisterType((*ClusterUp)(nil), "pb.ClusterUp")
	proto.RegisterType((*SplitLogTask)(nil), "pb.SplitLogTask")
	proto.RegisterType((*DeprecatedTableState)(nil), "pb.DeprecatedTableState")
	proto.RegisterType((*TableCF)(nil), "pb.TableCF")
	proto.RegisterType((*ReplicationPeer)(nil), "pb.ReplicationPeer")
	proto.RegisterType((*ReplicationState)(nil), "pb.ReplicationState")
	proto.RegisterType((*ReplicationHLogPosition)(nil), "pb.ReplicationHLogPosition")
	proto.RegisterType((*SwitchState)(nil), "pb.SwitchState")
	proto.RegisterEnum("pb.SplitLogTask_State", SplitLogTask_State_name, SplitLogTask_State_value)
	proto.RegisterEnum("pb.SplitLogTask_RecoveryMode", SplitLogTask_RecoveryMode_name, SplitLogTask_RecoveryMode_value)
	proto.RegisterEnum("pb.DeprecatedTableState_State", DeprecatedTableState_State_name, DeprecatedTableState_State_value)
	proto.RegisterEnum("pb.ReplicationState_State", ReplicationState_State_name, ReplicationState_State_value)
}

var fileDescriptor26 = []byte{
	// 793 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x54, 0xdd, 0x6e, 0xe2, 0x46,
	0x14, 0xae, 0x6d, 0x08, 0xf8, 0x18, 0x58, 0x3a, 0xfd, 0x59, 0x44, 0xbb, 0xdb, 0x95, 0x2f, 0x5a,
	0xb4, 0xad, 0xdc, 0x0a, 0xa9, 0x52, 0x15, 0xa9, 0xad, 0x02, 0x78, 0x13, 0xb4, 0xc4, 0x41, 0x63,
	0xd2, 0xaa, 0xbd, 0x41, 0xc6, 0x1e, 0xc0, 0x5a, 0xf0, 0x58, 0xb6, 0xe9, 0x2a, 0x6f, 0x90, 0x27,
	0xe8, 0x75, 0x9f, 0xa0, 0x77, 0x7d, 0xbf, 0x9e, 0x99, 0xe1, 0x57, 0x9b, 0x28, 0x37, 0xd6, 0xf9,
	0x9b, 0xf9, 0xbe, 0xef, 0x9c, 0xe3, 0x81, 0x67, 0x7f, 0x72, 0xfe, 0x96, 0xb1, 0x94, 0x65, 0x4e,
	0x9a, 0xf1, 0x82, 0x13, 0x3d, 0x9d, 0xb5, 0xad, 0xab, 0x5e, 0x90, 0x33, 0x15, 0x68, 0x7f, 0xd2,
	0x5f, 0x6d, 0xf2, 0x82, 0x65, 0x7e, 0x11, 0x14, 0x9b, 0x5c, 0x05, 0xed, 0x7b, 0x0d, 0x9a, 0xd7,
	0xac, 0x08, 0x28, 0x5b, 0xc4, 0x3c, 0xf1, 0x59, 0xf6, 0x17, 0xcb, 0xc8, 0xd7, 0x70, 0x96, 0x4b,
	0xab, 0xa5, 0xbd, 0xd2, 0x3b, 0x56, 0xb7, 0xe1, 0xa4, 0x33, 0x47, 0xe5, 0xbc, 0x60, 0xcd, 0xe8,
	0x36, 0x4b, 0xbe, 0x02, 0x2b, 0x4b, 0xc3, 0x29, 0x9a, 0x39, 0x1e, 0x6e, 0xe9, 0xaf, 0xb4, 0x4e,
	0x9d, 0x02, 0x86, 0x7e, 0x53, 0x11, 0xf2, 0x2d, 0x94, 0x73, 0x44, 0x63, 0x2d, 0x03, 0x53, 0x8d,
	0xee, 0x67, 0xe2, 0x9e, 0x2d, 0x92, 0x08, 0x3b, 0xf2, 0x4b, 0x55, 0x8d, 0x9d, 0xc0, 0xd9, 0x75,
	0x20, 0x08, 0x0a, 0xfc, 0xb5, 0xb4, 0x1e, 0xc3, 0x57, 0xd9, 0xa7, 0xf1, 0xbf, 0x00, 0x33, 0x4e,
	0xe6, 0x7c, 0x9a, 0xf2, 0xac, 0x90, 0x1c, 0xea, 0xb4, 0x2a, 0x02, 0x63, 0xf4, 0xed, 0xd7, 0x60,
	0x6e, 0x3b, 0x72, 0x9b, 0x92, 0x17, 0x00, 0xc8, 0x22, 0x2b, 0xa6, 0x91, 0xa0, 0x2b, 0x60, 0x4d,
	0x6a, 0xca, 0xc8, 0x40, 0x70, 0xfb, 0x57, 0x87, 0x9a, 0x9f, 0xae, 0xe2, 0x62, 0xc4, 0x17, 0x93,
	0x20, 0x7f, 0x47, 0xbe, 0xdb, 0x29, 0x13, 0xa5, 0x8d, 0xee, 0xe7, 0x92, 0xe1, 0x51, 0xc1, 0x89,
	0x34, 0xf2, 0x3d, 0x58, 0xaa, 0x65, 0xd3, 0x04, 0xf9, 0x23, 0xd1, 0x87, 0x54, 0x41, 0xbe, 0xb7,
	0xc9, 0x39, 0x94, 0xd6, 0x3c, 0xda, 0xf5, 0xed, 0xc5, 0x07, 0xb7, 0x53, 0x16, 0x72, 0xac, 0xbd,
	0xbb, 0xc6, 0xa2, 0xf3, 0xca, 0xad, 0xf7, 0xd6, 0xbb, 0xf9, 0xdd, 0xa3, 0xf2, 0x8c, 0xdd, 0x87,
	0xb2, 0x04, 0x27, 0x0d, 0x80, 0x5b, 0xef, 0xc2, 0xf7, 0x87, 0x97, 0x9e, 0x3b, 0x68, 0x7e, 0x44,
	0x4c, 0x28, 0x63, 0x15, 0x9a, 0x1a, 0xa9, 0x41, 0x95, 0xba, 0xdb, 0x84, 0x4e, 0xaa, 0x50, 0x1a,
	0xdc, 0x78, 0x6e, 0xd3, 0x20, 0x15, 0x30, 0x5c, 0x4a, 0x9b, 0x25, 0xfb, 0x17, 0xa8, 0x1d, 0x63,
	0x10, 0x0b, 0x76, 0x28, 0x78, 0xd1, 0xc7, 0x50, 0x1f, 0xdd, 0x5c, 0x4e, 0xfd, 0xf1, 0x68, 0x38,
	0x99, 0x0c, 0xbd, 0x4b, 0xbc, 0x10, 0xb1, 0x44, 0x88, 0xba, 0xe3, 0xd1, 0xc5, 0x1f, 0x4d, 0xdd,
	0xfe, 0x5b, 0x83, 0x4f, 0x07, 0x2c, 0xcd, 0x58, 0x88, 0x4c, 0xa2, 0x49, 0x30, 0x5b, 0x31, 0x45,
	0xea, 0xe7, 0xd3, 0xc6, 0xbd, 0x14, 0xd2, 0x1e, 0x2a, 0x54, 0x0d, 0x3c, 0xaf, 0xb8, 0xde, 0x45,
	0x6f, 0xe4, 0x0e, 0x76, 0x4b, 0xf2, 0xeb, 0x4e, 0x1c, 0x12, 0xda, 0xa6, 0x90, 0x10, 0xca, 0x19,
	0x0c, 0x7d, 0xe5, 0x69, 0xa4, 0x0e, 0xa6, 0xf2, 0x04, 0x35, 0x5d, 0x24, 0x65, 0xa5, 0xf0, 0x0c,
	0xdb, 0x87, 0x8a, 0x04, 0xe9, 0xbf, 0xc1, 0x19, 0x42, 0x21, 0x4c, 0x35, 0x14, 0x0d, 0x5b, 0x6d,
	0x75, 0xeb, 0x82, 0x8f, 0x2c, 0x90, 0x33, 0x31, 0x8b, 0x9d, 0x49, 0xda, 0x50, 0x9d, 0x07, 0xeb,
	0x78, 0x15, 0xb3, 0x1c, 0x07, 0x68, 0x74, 0x6a, 0x74, 0xef, 0xdb, 0xff, 0xe9, 0xf0, 0x8c, 0x32,
	0x1c, 0x10, 0x8a, 0xc0, 0xbd, 0x1b, 0x33, 0x5c, 0xce, 0x97, 0x00, 0xa1, 0x5a, 0xaf, 0x77, 0xec,
	0x6e, 0xbb, 0x51, 0x47, 0x11, 0xf2, 0x13, 0x3c, 0xcf, 0x0e, 0x47, 0xdc, 0x24, 0x4a, 0x79, 0x9c,
	0x14, 0xc3, 0x75, 0xba, 0x92, 0x8b, 0x6c, 0xd2, 0xc7, 0xd2, 0xf8, 0x7b, 0x94, 0x70, 0x4b, 0x03,
	0x5c, 0x0e, 0x03, 0x19, 0x13, 0xc1, 0xb8, 0x77, 0x57, 0xb0, 0x5c, 0x7e, 0xc6, 0x41, 0x9c, 0x51,
	0x99, 0x47, 0x84, 0x7a, 0xc8, 0x93, 0x79, 0xbc, 0xd8, 0x64, 0xf2, 0x92, 0x56, 0xe9, 0x70, 0x40,
	0x48, 0xf2, 0x8b, 0x2c, 0x4e, 0x16, 0xf2, 0xc0, 0x69, 0x21, 0xe9, 0x80, 0x12, 0x3e, 0x0d, 0xe7,
	0x79, 0xab, 0x2c, 0x4f, 0x59, 0xfb, 0xc6, 0xf4, 0xdf, 0xd0, 0xaa, 0xcc, 0xf6, 0xe7, 0xb9, 0x50,
	0x29, 0xba, 0x97, 0xa7, 0x41, 0x88, 0x7d, 0x39, 0x93, 0x7d, 0x39, 0x8a, 0x90, 0x2f, 0xc1, 0x9c,
	0x05, 0x49, 0xf4, 0x3e, 0x8e, 0x8a, 0x65, 0xab, 0x82, 0xba, 0x0c, 0x7a, 0x08, 0xd8, 0x4b, 0x68,
	0x1e, 0xb5, 0x4d, 0x0d, 0xf6, 0x87, 0xd3, 0x05, 0x69, 0xab, 0x37, 0xe3, 0xb4, 0xe8, 0xf4, 0xe1,
	0xb0, 0x9f, 0xde, 0x09, 0xfb, 0x47, 0x78, 0x7e, 0x74, 0xc9, 0x15, 0xfe, 0x47, 0x63, 0x9e, 0xc7,
	0x52, 0x2c, 0x0e, 0x36, 0xdd, 0xda, 0x12, 0xd3, 0xa0, 0x7b, 0xdf, 0xfe, 0x06, 0x2c, 0xff, 0x7d,
	0x5c, 0x84, 0x4b, 0x05, 0xd0, 0x82, 0x0a, 0x4b, 0x84, 0xf4, 0x48, 0xae, 0x4b, 0x95, 0xee, 0xdc,
	0x9e, 0x0b, 0xaf, 0x79, 0xb6, 0x70, 0x02, 0x54, 0xbd, 0x64, 0xce, 0x32, 0x88, 0x38, 0x4f, 0x9d,
	0xe5, 0x6c, 0xff, 0xfa, 0xce, 0x36, 0x73, 0x67, 0xc1, 0x12, 0x96, 0x89, 0x2d, 0xef, 0x1d, 0x1e,
	0xeb, 0xb1, 0x48, 0xe6, 0x57, 0xda, 0xbd, 0xa6, 0xfd, 0xa3, 0x69, 0xff, 0x07, 0x00, 0x00, 0xff,
	0xff, 0x76, 0x48, 0x5f, 0xc7, 0xc6, 0x05, 0x00, 0x00,
}
