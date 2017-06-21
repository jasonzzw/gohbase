// Code generated by protoc-gen-go.
// source: ClusterStatus.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RegionState_State int32

const (
	RegionState_OFFLINE       RegionState_State = 0
	RegionState_PENDING_OPEN  RegionState_State = 1
	RegionState_OPENING       RegionState_State = 2
	RegionState_OPEN          RegionState_State = 3
	RegionState_PENDING_CLOSE RegionState_State = 4
	RegionState_CLOSING       RegionState_State = 5
	RegionState_CLOSED        RegionState_State = 6
	RegionState_SPLITTING     RegionState_State = 7
	RegionState_SPLIT         RegionState_State = 8
	RegionState_FAILED_OPEN   RegionState_State = 9
	RegionState_FAILED_CLOSE  RegionState_State = 10
	RegionState_MERGING       RegionState_State = 11
	RegionState_MERGED        RegionState_State = 12
	RegionState_SPLITTING_NEW RegionState_State = 13
	// region but hasn't be created yet, or master doesn't
	// know it's already created
	RegionState_MERGING_NEW RegionState_State = 14
)

var RegionState_State_name = map[int32]string{
	0:  "OFFLINE",
	1:  "PENDING_OPEN",
	2:  "OPENING",
	3:  "OPEN",
	4:  "PENDING_CLOSE",
	5:  "CLOSING",
	6:  "CLOSED",
	7:  "SPLITTING",
	8:  "SPLIT",
	9:  "FAILED_OPEN",
	10: "FAILED_CLOSE",
	11: "MERGING",
	12: "MERGED",
	13: "SPLITTING_NEW",
	14: "MERGING_NEW",
}
var RegionState_State_value = map[string]int32{
	"OFFLINE":       0,
	"PENDING_OPEN":  1,
	"OPENING":       2,
	"OPEN":          3,
	"PENDING_CLOSE": 4,
	"CLOSING":       5,
	"CLOSED":        6,
	"SPLITTING":     7,
	"SPLIT":         8,
	"FAILED_OPEN":   9,
	"FAILED_CLOSE":  10,
	"MERGING":       11,
	"MERGED":        12,
	"SPLITTING_NEW": 13,
	"MERGING_NEW":   14,
}

func (x RegionState_State) Enum() *RegionState_State {
	p := new(RegionState_State)
	*p = x
	return p
}
func (x RegionState_State) String() string {
	return proto.EnumName(RegionState_State_name, int32(x))
}
func (x *RegionState_State) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RegionState_State_value, data, "RegionState_State")
	if err != nil {
		return err
	}
	*x = RegionState_State(value)
	return nil
}
func (RegionState_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{0, 0} }

type RegionState struct {
	RegionInfo       *RegionInfo        `protobuf:"bytes,1,req,name=region_info,json=regionInfo" json:"region_info,omitempty"`
	State            *RegionState_State `protobuf:"varint,2,req,name=state,enum=pb.RegionState_State" json:"state,omitempty"`
	Stamp            *uint64            `protobuf:"varint,3,opt,name=stamp" json:"stamp,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *RegionState) Reset()                    { *m = RegionState{} }
func (m *RegionState) String() string            { return proto.CompactTextString(m) }
func (*RegionState) ProtoMessage()               {}
func (*RegionState) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *RegionState) GetRegionInfo() *RegionInfo {
	if m != nil {
		return m.RegionInfo
	}
	return nil
}

func (m *RegionState) GetState() RegionState_State {
	if m != nil && m.State != nil {
		return *m.State
	}
	return RegionState_OFFLINE
}

func (m *RegionState) GetStamp() uint64 {
	if m != nil && m.Stamp != nil {
		return *m.Stamp
	}
	return 0
}

type RegionInTransition struct {
	Spec             *RegionSpecifier `protobuf:"bytes,1,req,name=spec" json:"spec,omitempty"`
	RegionState      *RegionState     `protobuf:"bytes,2,req,name=region_state,json=regionState" json:"region_state,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *RegionInTransition) Reset()                    { *m = RegionInTransition{} }
func (m *RegionInTransition) String() string            { return proto.CompactTextString(m) }
func (*RegionInTransition) ProtoMessage()               {}
func (*RegionInTransition) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *RegionInTransition) GetSpec() *RegionSpecifier {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *RegionInTransition) GetRegionState() *RegionState {
	if m != nil {
		return m.RegionState
	}
	return nil
}

// *
// sequence Id of a store
type StoreSequenceId struct {
	FamilyName       []byte  `protobuf:"bytes,1,req,name=family_name,json=familyName" json:"family_name,omitempty"`
	SequenceId       *uint64 `protobuf:"varint,2,req,name=sequence_id,json=sequenceId" json:"sequence_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StoreSequenceId) Reset()                    { *m = StoreSequenceId{} }
func (m *StoreSequenceId) String() string            { return proto.CompactTextString(m) }
func (*StoreSequenceId) ProtoMessage()               {}
func (*StoreSequenceId) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *StoreSequenceId) GetFamilyName() []byte {
	if m != nil {
		return m.FamilyName
	}
	return nil
}

func (m *StoreSequenceId) GetSequenceId() uint64 {
	if m != nil && m.SequenceId != nil {
		return *m.SequenceId
	}
	return 0
}

// *
// contains a sequence id of a region which should be the minimum of its store sequence ids and
// list of sequence ids of the region's stores
type RegionStoreSequenceIds struct {
	LastFlushedSequenceId *uint64            `protobuf:"varint,1,req,name=last_flushed_sequence_id,json=lastFlushedSequenceId" json:"last_flushed_sequence_id,omitempty"`
	StoreSequenceId       []*StoreSequenceId `protobuf:"bytes,2,rep,name=store_sequence_id,json=storeSequenceId" json:"store_sequence_id,omitempty"`
	XXX_unrecognized      []byte             `json:"-"`
}

func (m *RegionStoreSequenceIds) Reset()                    { *m = RegionStoreSequenceIds{} }
func (m *RegionStoreSequenceIds) String() string            { return proto.CompactTextString(m) }
func (*RegionStoreSequenceIds) ProtoMessage()               {}
func (*RegionStoreSequenceIds) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *RegionStoreSequenceIds) GetLastFlushedSequenceId() uint64 {
	if m != nil && m.LastFlushedSequenceId != nil {
		return *m.LastFlushedSequenceId
	}
	return 0
}

func (m *RegionStoreSequenceIds) GetStoreSequenceId() []*StoreSequenceId {
	if m != nil {
		return m.StoreSequenceId
	}
	return nil
}

type RegionLoad struct {
	// * the region specifier
	RegionSpecifier *RegionSpecifier `protobuf:"bytes,1,req,name=region_specifier,json=regionSpecifier" json:"region_specifier,omitempty"`
	// * the number of stores for the region
	Stores *uint32 `protobuf:"varint,2,opt,name=stores" json:"stores,omitempty"`
	// * the number of storefiles for the region
	Storefiles *uint32 `protobuf:"varint,3,opt,name=storefiles" json:"storefiles,omitempty"`
	// * the total size of the store files for the region, uncompressed, in MB
	StoreUncompressedSize_MB *uint32 `protobuf:"varint,4,opt,name=store_uncompressed_size_MB,json=storeUncompressedSizeMB" json:"store_uncompressed_size_MB,omitempty"`
	// * the current total size of the store files for the region, in MB
	StorefileSize_MB *uint32 `protobuf:"varint,5,opt,name=storefile_size_MB,json=storefileSizeMB" json:"storefile_size_MB,omitempty"`
	// * the current size of the memstore for the region, in MB
	MemstoreSize_MB *uint32 `protobuf:"varint,6,opt,name=memstore_size_MB,json=memstoreSizeMB" json:"memstore_size_MB,omitempty"`
	// *
	// The current total size of root-level store file indexes for the region,
	// in MB. The same as {@link #rootIndexSizeKB} but in MB.
	StorefileIndexSize_MB *uint32 `protobuf:"varint,7,opt,name=storefile_index_size_MB,json=storefileIndexSizeMB" json:"storefile_index_size_MB,omitempty"`
	// * the current total read requests made to region
	ReadRequestsCount *uint64 `protobuf:"varint,8,opt,name=read_requests_count,json=readRequestsCount" json:"read_requests_count,omitempty"`
	// * the current total write requests made to region
	WriteRequestsCount *uint64 `protobuf:"varint,9,opt,name=write_requests_count,json=writeRequestsCount" json:"write_requests_count,omitempty"`
	// * the total compacting key values in currently running compaction
	TotalCompacting_KVs *uint64 `protobuf:"varint,10,opt,name=total_compacting_KVs,json=totalCompactingKVs" json:"total_compacting_KVs,omitempty"`
	// * the completed count of key values in currently running compaction
	CurrentCompacted_KVs *uint64 `protobuf:"varint,11,opt,name=current_compacted_KVs,json=currentCompactedKVs" json:"current_compacted_KVs,omitempty"`
	// * The current total size of root-level indexes for the region, in KB.
	RootIndexSize_KB *uint32 `protobuf:"varint,12,opt,name=root_index_size_KB,json=rootIndexSizeKB" json:"root_index_size_KB,omitempty"`
	// * The total size of all index blocks, not just the root level, in KB.
	TotalStaticIndexSize_KB *uint32 `protobuf:"varint,13,opt,name=total_static_index_size_KB,json=totalStaticIndexSizeKB" json:"total_static_index_size_KB,omitempty"`
	// *
	// The total size of all Bloom filter blocks, not just loaded into the
	// block cache, in KB.
	TotalStaticBloomSize_KB *uint32 `protobuf:"varint,14,opt,name=total_static_bloom_size_KB,json=totalStaticBloomSizeKB" json:"total_static_bloom_size_KB,omitempty"`
	// * the most recent sequence Id from cache flush
	CompleteSequenceId *uint64 `protobuf:"varint,15,opt,name=complete_sequence_id,json=completeSequenceId" json:"complete_sequence_id,omitempty"`
	// * The current data locality for region in the regionserver
	DataLocality          *float32 `protobuf:"fixed32,16,opt,name=data_locality,json=dataLocality" json:"data_locality,omitempty"`
	LastMajorCompactionTs *uint64  `protobuf:"varint,17,opt,name=last_major_compaction_ts,json=lastMajorCompactionTs,def=0" json:"last_major_compaction_ts,omitempty"`
	// * the most recent sequence Id of store from cache flush
	StoreCompleteSequenceId []*StoreSequenceId `protobuf:"bytes,18,rep,name=store_complete_sequence_id,json=storeCompleteSequenceId" json:"store_complete_sequence_id,omitempty"`
	XXX_unrecognized        []byte             `json:"-"`
}

func (m *RegionLoad) Reset()                    { *m = RegionLoad{} }
func (m *RegionLoad) String() string            { return proto.CompactTextString(m) }
func (*RegionLoad) ProtoMessage()               {}
func (*RegionLoad) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

const Default_RegionLoad_LastMajorCompactionTs uint64 = 0

func (m *RegionLoad) GetRegionSpecifier() *RegionSpecifier {
	if m != nil {
		return m.RegionSpecifier
	}
	return nil
}

func (m *RegionLoad) GetStores() uint32 {
	if m != nil && m.Stores != nil {
		return *m.Stores
	}
	return 0
}

func (m *RegionLoad) GetStorefiles() uint32 {
	if m != nil && m.Storefiles != nil {
		return *m.Storefiles
	}
	return 0
}

func (m *RegionLoad) GetStoreUncompressedSize_MB() uint32 {
	if m != nil && m.StoreUncompressedSize_MB != nil {
		return *m.StoreUncompressedSize_MB
	}
	return 0
}

func (m *RegionLoad) GetStorefileSize_MB() uint32 {
	if m != nil && m.StorefileSize_MB != nil {
		return *m.StorefileSize_MB
	}
	return 0
}

func (m *RegionLoad) GetMemstoreSize_MB() uint32 {
	if m != nil && m.MemstoreSize_MB != nil {
		return *m.MemstoreSize_MB
	}
	return 0
}

func (m *RegionLoad) GetStorefileIndexSize_MB() uint32 {
	if m != nil && m.StorefileIndexSize_MB != nil {
		return *m.StorefileIndexSize_MB
	}
	return 0
}

func (m *RegionLoad) GetReadRequestsCount() uint64 {
	if m != nil && m.ReadRequestsCount != nil {
		return *m.ReadRequestsCount
	}
	return 0
}

func (m *RegionLoad) GetWriteRequestsCount() uint64 {
	if m != nil && m.WriteRequestsCount != nil {
		return *m.WriteRequestsCount
	}
	return 0
}

func (m *RegionLoad) GetTotalCompacting_KVs() uint64 {
	if m != nil && m.TotalCompacting_KVs != nil {
		return *m.TotalCompacting_KVs
	}
	return 0
}

func (m *RegionLoad) GetCurrentCompacted_KVs() uint64 {
	if m != nil && m.CurrentCompacted_KVs != nil {
		return *m.CurrentCompacted_KVs
	}
	return 0
}

func (m *RegionLoad) GetRootIndexSize_KB() uint32 {
	if m != nil && m.RootIndexSize_KB != nil {
		return *m.RootIndexSize_KB
	}
	return 0
}

func (m *RegionLoad) GetTotalStaticIndexSize_KB() uint32 {
	if m != nil && m.TotalStaticIndexSize_KB != nil {
		return *m.TotalStaticIndexSize_KB
	}
	return 0
}

func (m *RegionLoad) GetTotalStaticBloomSize_KB() uint32 {
	if m != nil && m.TotalStaticBloomSize_KB != nil {
		return *m.TotalStaticBloomSize_KB
	}
	return 0
}

func (m *RegionLoad) GetCompleteSequenceId() uint64 {
	if m != nil && m.CompleteSequenceId != nil {
		return *m.CompleteSequenceId
	}
	return 0
}

func (m *RegionLoad) GetDataLocality() float32 {
	if m != nil && m.DataLocality != nil {
		return *m.DataLocality
	}
	return 0
}

func (m *RegionLoad) GetLastMajorCompactionTs() uint64 {
	if m != nil && m.LastMajorCompactionTs != nil {
		return *m.LastMajorCompactionTs
	}
	return Default_RegionLoad_LastMajorCompactionTs
}

func (m *RegionLoad) GetStoreCompleteSequenceId() []*StoreSequenceId {
	if m != nil {
		return m.StoreCompleteSequenceId
	}
	return nil
}

type ReplicationLoadSink struct {
	AgeOfLastAppliedOp        *uint64 `protobuf:"varint,1,req,name=ageOfLastAppliedOp" json:"ageOfLastAppliedOp,omitempty"`
	TimeStampsOfLastAppliedOp *uint64 `protobuf:"varint,2,req,name=timeStampsOfLastAppliedOp" json:"timeStampsOfLastAppliedOp,omitempty"`
	XXX_unrecognized          []byte  `json:"-"`
}

func (m *ReplicationLoadSink) Reset()                    { *m = ReplicationLoadSink{} }
func (m *ReplicationLoadSink) String() string            { return proto.CompactTextString(m) }
func (*ReplicationLoadSink) ProtoMessage()               {}
func (*ReplicationLoadSink) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{5} }

func (m *ReplicationLoadSink) GetAgeOfLastAppliedOp() uint64 {
	if m != nil && m.AgeOfLastAppliedOp != nil {
		return *m.AgeOfLastAppliedOp
	}
	return 0
}

func (m *ReplicationLoadSink) GetTimeStampsOfLastAppliedOp() uint64 {
	if m != nil && m.TimeStampsOfLastAppliedOp != nil {
		return *m.TimeStampsOfLastAppliedOp
	}
	return 0
}

type ReplicationLoadSource struct {
	PeerID                   *string `protobuf:"bytes,1,req,name=peerID" json:"peerID,omitempty"`
	AgeOfLastShippedOp       *uint64 `protobuf:"varint,2,req,name=ageOfLastShippedOp" json:"ageOfLastShippedOp,omitempty"`
	SizeOfLogQueue           *uint32 `protobuf:"varint,3,req,name=sizeOfLogQueue" json:"sizeOfLogQueue,omitempty"`
	TimeStampOfLastShippedOp *uint64 `protobuf:"varint,4,req,name=timeStampOfLastShippedOp" json:"timeStampOfLastShippedOp,omitempty"`
	ReplicationLag           *uint64 `protobuf:"varint,5,req,name=replicationLag" json:"replicationLag,omitempty"`
	XXX_unrecognized         []byte  `json:"-"`
}

func (m *ReplicationLoadSource) Reset()                    { *m = ReplicationLoadSource{} }
func (m *ReplicationLoadSource) String() string            { return proto.CompactTextString(m) }
func (*ReplicationLoadSource) ProtoMessage()               {}
func (*ReplicationLoadSource) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{6} }

func (m *ReplicationLoadSource) GetPeerID() string {
	if m != nil && m.PeerID != nil {
		return *m.PeerID
	}
	return ""
}

func (m *ReplicationLoadSource) GetAgeOfLastShippedOp() uint64 {
	if m != nil && m.AgeOfLastShippedOp != nil {
		return *m.AgeOfLastShippedOp
	}
	return 0
}

func (m *ReplicationLoadSource) GetSizeOfLogQueue() uint32 {
	if m != nil && m.SizeOfLogQueue != nil {
		return *m.SizeOfLogQueue
	}
	return 0
}

func (m *ReplicationLoadSource) GetTimeStampOfLastShippedOp() uint64 {
	if m != nil && m.TimeStampOfLastShippedOp != nil {
		return *m.TimeStampOfLastShippedOp
	}
	return 0
}

func (m *ReplicationLoadSource) GetReplicationLag() uint64 {
	if m != nil && m.ReplicationLag != nil {
		return *m.ReplicationLag
	}
	return 0
}

type ServerLoad struct {
	// * Number of requests since last report.
	NumberOfRequests *uint64 `protobuf:"varint,1,opt,name=number_of_requests,json=numberOfRequests" json:"number_of_requests,omitempty"`
	// * Total Number of requests from the start of the region server.
	TotalNumberOfRequests *uint64 `protobuf:"varint,2,opt,name=total_number_of_requests,json=totalNumberOfRequests" json:"total_number_of_requests,omitempty"`
	// * the amount of used heap, in MB.
	UsedHeap_MB *uint32 `protobuf:"varint,3,opt,name=used_heap_MB,json=usedHeapMB" json:"used_heap_MB,omitempty"`
	// * the maximum allowable size of the heap, in MB.
	MaxHeap_MB *uint32 `protobuf:"varint,4,opt,name=max_heap_MB,json=maxHeapMB" json:"max_heap_MB,omitempty"`
	// * Information on the load of individual regions.
	RegionLoads []*RegionLoad `protobuf:"bytes,5,rep,name=region_loads,json=regionLoads" json:"region_loads,omitempty"`
	// *
	// Regionserver-level coprocessors, e.g., WALObserver implementations.
	// Region-level coprocessors, on the other hand, are stored inside RegionLoad
	// objects.
	Coprocessors []*Coprocessor `protobuf:"bytes,6,rep,name=coprocessors" json:"coprocessors,omitempty"`
	// *
	// Time when incremental (non-total) counts began being calculated (e.g. number_of_requests)
	// time is measured as the difference, measured in milliseconds, between the current time
	// and midnight, January 1, 1970 UTC.
	ReportStartTime *uint64 `protobuf:"varint,7,opt,name=report_start_time,json=reportStartTime" json:"report_start_time,omitempty"`
	// *
	// Time when report was generated.
	// time is measured as the difference, measured in milliseconds, between the current time
	// and midnight, January 1, 1970 UTC.
	ReportEndTime *uint64 `protobuf:"varint,8,opt,name=report_end_time,json=reportEndTime" json:"report_end_time,omitempty"`
	// *
	// The port number that this region server is hosing an info server on.
	InfoServerPort *uint32 `protobuf:"varint,9,opt,name=info_server_port,json=infoServerPort" json:"info_server_port,omitempty"`
	// *
	// The replicationLoadSource for the replication Source status of this region server.
	ReplLoadSource []*ReplicationLoadSource `protobuf:"bytes,10,rep,name=replLoadSource" json:"replLoadSource,omitempty"`
	// *
	// The replicationLoadSink for the replication Sink status of this region server.
	ReplLoadSink     *ReplicationLoadSink `protobuf:"bytes,11,opt,name=replLoadSink" json:"replLoadSink,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *ServerLoad) Reset()                    { *m = ServerLoad{} }
func (m *ServerLoad) String() string            { return proto.CompactTextString(m) }
func (*ServerLoad) ProtoMessage()               {}
func (*ServerLoad) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{7} }

func (m *ServerLoad) GetNumberOfRequests() uint64 {
	if m != nil && m.NumberOfRequests != nil {
		return *m.NumberOfRequests
	}
	return 0
}

func (m *ServerLoad) GetTotalNumberOfRequests() uint64 {
	if m != nil && m.TotalNumberOfRequests != nil {
		return *m.TotalNumberOfRequests
	}
	return 0
}

func (m *ServerLoad) GetUsedHeap_MB() uint32 {
	if m != nil && m.UsedHeap_MB != nil {
		return *m.UsedHeap_MB
	}
	return 0
}

func (m *ServerLoad) GetMaxHeap_MB() uint32 {
	if m != nil && m.MaxHeap_MB != nil {
		return *m.MaxHeap_MB
	}
	return 0
}

func (m *ServerLoad) GetRegionLoads() []*RegionLoad {
	if m != nil {
		return m.RegionLoads
	}
	return nil
}

func (m *ServerLoad) GetCoprocessors() []*Coprocessor {
	if m != nil {
		return m.Coprocessors
	}
	return nil
}

func (m *ServerLoad) GetReportStartTime() uint64 {
	if m != nil && m.ReportStartTime != nil {
		return *m.ReportStartTime
	}
	return 0
}

func (m *ServerLoad) GetReportEndTime() uint64 {
	if m != nil && m.ReportEndTime != nil {
		return *m.ReportEndTime
	}
	return 0
}

func (m *ServerLoad) GetInfoServerPort() uint32 {
	if m != nil && m.InfoServerPort != nil {
		return *m.InfoServerPort
	}
	return 0
}

func (m *ServerLoad) GetReplLoadSource() []*ReplicationLoadSource {
	if m != nil {
		return m.ReplLoadSource
	}
	return nil
}

func (m *ServerLoad) GetReplLoadSink() *ReplicationLoadSink {
	if m != nil {
		return m.ReplLoadSink
	}
	return nil
}

type LiveServerInfo struct {
	Server           *ServerName `protobuf:"bytes,1,req,name=server" json:"server,omitempty"`
	ServerLoad       *ServerLoad `protobuf:"bytes,2,req,name=server_load,json=serverLoad" json:"server_load,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *LiveServerInfo) Reset()                    { *m = LiveServerInfo{} }
func (m *LiveServerInfo) String() string            { return proto.CompactTextString(m) }
func (*LiveServerInfo) ProtoMessage()               {}
func (*LiveServerInfo) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{8} }

func (m *LiveServerInfo) GetServer() *ServerName {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *LiveServerInfo) GetServerLoad() *ServerLoad {
	if m != nil {
		return m.ServerLoad
	}
	return nil
}

type ClusterStatus struct {
	HbaseVersion        *HBaseVersionFileContent `protobuf:"bytes,1,opt,name=hbase_version,json=hbaseVersion" json:"hbase_version,omitempty"`
	LiveServers         []*LiveServerInfo        `protobuf:"bytes,2,rep,name=live_servers,json=liveServers" json:"live_servers,omitempty"`
	DeadServers         []*ServerName            `protobuf:"bytes,3,rep,name=dead_servers,json=deadServers" json:"dead_servers,omitempty"`
	RegionsInTransition []*RegionInTransition    `protobuf:"bytes,4,rep,name=regions_in_transition,json=regionsInTransition" json:"regions_in_transition,omitempty"`
	ClusterId           *ClusterId               `protobuf:"bytes,5,opt,name=cluster_id,json=clusterId" json:"cluster_id,omitempty"`
	MasterCoprocessors  []*Coprocessor           `protobuf:"bytes,6,rep,name=master_coprocessors,json=masterCoprocessors" json:"master_coprocessors,omitempty"`
	Master              *ServerName              `protobuf:"bytes,7,opt,name=master" json:"master,omitempty"`
	BackupMasters       []*ServerName            `protobuf:"bytes,8,rep,name=backup_masters,json=backupMasters" json:"backup_masters,omitempty"`
	BalancerOn          *bool                    `protobuf:"varint,9,opt,name=balancer_on,json=balancerOn" json:"balancer_on,omitempty"`
	XXX_unrecognized    []byte                   `json:"-"`
}

func (m *ClusterStatus) Reset()                    { *m = ClusterStatus{} }
func (m *ClusterStatus) String() string            { return proto.CompactTextString(m) }
func (*ClusterStatus) ProtoMessage()               {}
func (*ClusterStatus) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{9} }

func (m *ClusterStatus) GetHbaseVersion() *HBaseVersionFileContent {
	if m != nil {
		return m.HbaseVersion
	}
	return nil
}

func (m *ClusterStatus) GetLiveServers() []*LiveServerInfo {
	if m != nil {
		return m.LiveServers
	}
	return nil
}

func (m *ClusterStatus) GetDeadServers() []*ServerName {
	if m != nil {
		return m.DeadServers
	}
	return nil
}

func (m *ClusterStatus) GetRegionsInTransition() []*RegionInTransition {
	if m != nil {
		return m.RegionsInTransition
	}
	return nil
}

func (m *ClusterStatus) GetClusterId() *ClusterId {
	if m != nil {
		return m.ClusterId
	}
	return nil
}

func (m *ClusterStatus) GetMasterCoprocessors() []*Coprocessor {
	if m != nil {
		return m.MasterCoprocessors
	}
	return nil
}

func (m *ClusterStatus) GetMaster() *ServerName {
	if m != nil {
		return m.Master
	}
	return nil
}

func (m *ClusterStatus) GetBackupMasters() []*ServerName {
	if m != nil {
		return m.BackupMasters
	}
	return nil
}

func (m *ClusterStatus) GetBalancerOn() bool {
	if m != nil && m.BalancerOn != nil {
		return *m.BalancerOn
	}
	return false
}

func init() {
	proto.RegisterType((*RegionState)(nil), "pb.RegionState")
	proto.RegisterType((*RegionInTransition)(nil), "pb.RegionInTransition")
	proto.RegisterType((*StoreSequenceId)(nil), "pb.StoreSequenceId")
	proto.RegisterType((*RegionStoreSequenceIds)(nil), "pb.RegionStoreSequenceIds")
	proto.RegisterType((*RegionLoad)(nil), "pb.RegionLoad")
	proto.RegisterType((*ReplicationLoadSink)(nil), "pb.ReplicationLoadSink")
	proto.RegisterType((*ReplicationLoadSource)(nil), "pb.ReplicationLoadSource")
	proto.RegisterType((*ServerLoad)(nil), "pb.ServerLoad")
	proto.RegisterType((*LiveServerInfo)(nil), "pb.LiveServerInfo")
	proto.RegisterType((*ClusterStatus)(nil), "pb.ClusterStatus")
	proto.RegisterEnum("pb.RegionState_State", RegionState_State_name, RegionState_State_value)
}

var fileDescriptor7 = []byte{
	// 1453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x57, 0x4d, 0x6f, 0xdb, 0x46,
	0x13, 0x7e, 0xf5, 0xe1, 0xaf, 0xa1, 0xbe, 0xbc, 0x8e, 0x1d, 0xc5, 0x2f, 0x90, 0x06, 0x2a, 0x90,
	0x1a, 0x49, 0xa0, 0xa6, 0x2e, 0x8a, 0x02, 0x69, 0xd1, 0x36, 0x92, 0xed, 0x44, 0x8d, 0xfc, 0x51,
	0xd2, 0x4d, 0x8f, 0x04, 0x45, 0xae, 0x6c, 0x36, 0x14, 0xc9, 0x70, 0xa9, 0x34, 0xe9, 0xb5, 0xff,
	0xa0, 0xd7, 0x5e, 0xf2, 0xdf, 0x7a, 0xe8, 0xa1, 0x3f, 0xa0, 0xd7, 0xce, 0xec, 0x2e, 0x29, 0x52,
	0x51, 0x82, 0x5e, 0x84, 0xe5, 0xf3, 0x3c, 0x33, 0xbb, 0x3b, 0x3b, 0x3b, 0x3b, 0x82, 0x9d, 0x61,
	0x30, 0x17, 0x29, 0x4f, 0xac, 0xd4, 0x49, 0xe7, 0xa2, 0x1f, 0x27, 0x51, 0x1a, 0xb1, 0x6a, 0x3c,
	0xd9, 0x37, 0x9e, 0x0e, 0x1c, 0xc1, 0x15, 0xb0, 0xdf, 0xd6, 0xaa, 0x91, 0xa7, 0x81, 0xcd, 0x13,
	0x4b, 0x8d, 0x7a, 0x7f, 0x55, 0xc1, 0x30, 0xf9, 0x95, 0x1f, 0x85, 0xe4, 0x82, 0xb3, 0x4f, 0xc1,
	0x48, 0xe4, 0xa7, 0xed, 0x87, 0xd3, 0xa8, 0x5b, 0xb9, 0x53, 0x3d, 0x30, 0x0e, 0x5b, 0xfd, 0x78,
	0xd2, 0x57, 0xaa, 0x11, 0xa2, 0x26, 0x24, 0xf9, 0x98, 0xdd, 0x87, 0x35, 0x41, 0x96, 0xdd, 0x2a,
	0x4a, 0x5b, 0x87, 0xbb, 0x0b, 0xa9, 0x74, 0xd8, 0x97, 0xbf, 0xa6, 0xd2, 0xb0, 0x1b, 0x52, 0x3c,
	0x8b, 0xbb, 0xb5, 0x3b, 0x95, 0x83, 0xba, 0xa9, 0x3e, 0x7a, 0x7f, 0x56, 0x60, 0x4d, 0xcd, 0x6e,
	0xc0, 0xc6, 0xf9, 0xc9, 0xc9, 0x78, 0x74, 0x76, 0xdc, 0xf9, 0x1f, 0xeb, 0x40, 0xe3, 0xe2, 0xf8,
	0xec, 0x68, 0x74, 0xf6, 0xc4, 0x3e, 0xc7, 0x41, 0xa7, 0x22, 0x69, 0x1c, 0x21, 0xd2, 0xa9, 0xb2,
	0x4d, 0xa8, 0x4b, 0xb8, 0xc6, 0xb6, 0xa1, 0x99, 0x09, 0x87, 0xe3, 0x73, 0xeb, 0xb8, 0x53, 0x27,
	0x25, 0x0d, 0x49, 0xb9, 0xc6, 0x00, 0xd6, 0x25, 0x7e, 0xd4, 0x59, 0x67, 0x4d, 0xd8, 0xb2, 0x2e,
	0xc6, 0xa3, 0xcb, 0x4b, 0xa2, 0x36, 0xd8, 0x16, 0xce, 0x4c, 0x9f, 0x9d, 0x4d, 0xd6, 0x06, 0xe3,
	0xe4, 0xf1, 0x68, 0x7c, 0x7c, 0xa4, 0x66, 0xdb, 0xa2, 0xf9, 0x35, 0xa0, 0xbc, 0x02, 0x79, 0x3d,
	0x3d, 0x36, 0x9f, 0x90, 0xa9, 0x41, 0x5e, 0xe9, 0x03, 0xbd, 0x36, 0x68, 0x05, 0xb9, 0x57, 0xfb,
	0xec, 0xf8, 0xa7, 0x4e, 0x93, 0xdc, 0x69, 0xad, 0x04, 0x5a, 0xbd, 0x97, 0xc0, 0xb2, 0x10, 0x5e,
	0x26, 0x4e, 0x28, 0xfc, 0x14, 0xc7, 0xec, 0x13, 0xa8, 0x8b, 0x98, 0xbb, 0x3a, 0xd0, 0x3b, 0x85,
	0xe8, 0x21, 0xea, 0x4f, 0x7d, 0x9e, 0x98, 0x52, 0xc0, 0x0e, 0xa1, 0xa1, 0x0f, 0x66, 0x11, 0x6e,
	0xe3, 0xb0, 0xbd, 0x14, 0x6e, 0x53, 0x9f, 0x9e, 0xfc, 0xe8, 0x59, 0xd0, 0xb6, 0xd2, 0x28, 0xe1,
	0x16, 0x7f, 0x39, 0xe7, 0xa1, 0xcb, 0x47, 0x1e, 0xfb, 0x08, 0x8c, 0xa9, 0x33, 0xf3, 0x83, 0x37,
	0x76, 0xe8, 0xcc, 0xb8, 0x9c, 0xb6, 0x61, 0x82, 0x82, 0xce, 0x10, 0x21, 0x81, 0xd0, 0x72, 0xdb,
	0xf7, 0xe4, 0x34, 0x75, 0x13, 0x44, 0xee, 0xa1, 0xf7, 0x7b, 0x05, 0xf6, 0xb2, 0x19, 0x4b, 0xbe,
	0x05, 0xfb, 0x12, 0xba, 0x81, 0x23, 0x52, 0x7b, 0x8a, 0xe9, 0x76, 0xcd, 0x3d, 0xbb, 0xe8, 0xa8,
	0x22, 0x1d, 0xed, 0x12, 0x7f, 0xa2, 0xe8, 0xc2, 0xaa, 0xbe, 0x85, 0x6d, 0x41, 0xce, 0xec, 0xf2,
	0xd4, 0xb5, 0x2c, 0x24, 0x4b, 0x33, 0x99, 0x6d, 0x51, 0x06, 0x7a, 0x6f, 0x37, 0x00, 0xd4, 0xa2,
	0xc6, 0x91, 0xe3, 0xb1, 0x6f, 0xa0, 0x93, 0x05, 0x2b, 0x0b, 0xe3, 0x87, 0x22, 0xdc, 0x4e, 0xca,
	0x00, 0xdb, 0x83, 0x75, 0x39, 0x83, 0xc0, 0x45, 0x54, 0x0e, 0x9a, 0xa6, 0xfe, 0x62, 0xb7, 0x01,
	0xe4, 0x68, 0xea, 0x07, 0xc8, 0xd5, 0x24, 0x57, 0x40, 0xd8, 0x57, 0xb0, 0xaf, 0xf6, 0x31, 0x0f,
	0xdd, 0x68, 0x16, 0xa3, 0x89, 0xa0, 0x30, 0xf8, 0xbf, 0x72, 0xfb, 0x74, 0xd0, 0xad, 0x4b, 0xfd,
	0x4d, 0xa9, 0xf8, 0xb1, 0x20, 0xb0, 0x90, 0x3f, 0x1d, 0xb0, 0x7b, 0x3a, 0x08, 0xe4, 0x2a, 0xb7,
	0x59, 0x93, 0x36, 0xed, 0x9c, 0xd0, 0xda, 0x03, 0xe8, 0xcc, 0xf8, 0x4c, 0xc7, 0x4c, 0x4b, 0xd7,
	0xa5, 0xb4, 0x95, 0xe1, 0x5a, 0xf9, 0x05, 0xdc, 0x5c, 0x78, 0xf5, 0x43, 0x8f, 0xbf, 0xce, 0x0d,
	0x36, 0xa4, 0xc1, 0x8d, 0x9c, 0x1e, 0x11, 0xab, 0xcd, 0xfa, 0xb0, 0x93, 0x70, 0xc7, 0xb3, 0x13,
	0x8a, 0xb1, 0x48, 0x85, 0xed, 0x46, 0xf3, 0x30, 0xed, 0x6e, 0xca, 0x7b, 0xbb, 0x4d, 0x94, 0xa9,
	0x99, 0x21, 0x11, 0xec, 0x21, 0xdc, 0xf8, 0x25, 0xf1, 0x53, 0xbe, 0x6c, 0xb0, 0x25, 0x0d, 0x98,
	0xe4, 0xde, 0xb1, 0x48, 0xa3, 0xd4, 0x09, 0x6c, 0x0a, 0x84, 0xe3, 0xa6, 0x7e, 0x78, 0x65, 0x3f,
	0x7b, 0x2e, 0xba, 0xa0, 0x2c, 0x24, 0x37, 0xcc, 0x29, 0x64, 0xf0, 0x0a, 0xec, 0xba, 0xf3, 0x24,
	0xe1, 0x61, 0x9a, 0xd9, 0x60, 0x70, 0xc9, 0xc4, 0x90, 0x26, 0x3b, 0x9a, 0x1c, 0x66, 0x1c, 0xd9,
	0xdc, 0x07, 0x96, 0x44, 0x51, 0x5a, 0xdc, 0xf9, 0xb3, 0x41, 0xb7, 0xa1, 0xa2, 0x4a, 0x4c, 0xbe,
	0xe9, 0x67, 0x03, 0xf6, 0x08, 0xf6, 0xd5, 0x92, 0xe8, 0x8a, 0xf9, 0xee, 0x92, 0x51, 0x53, 0x1a,
	0xed, 0x49, 0x85, 0x25, 0x05, 0x1f, 0xb2, 0x9d, 0x04, 0x51, 0x34, 0xcb, 0x6d, 0x5b, 0xef, 0xd8,
	0x0e, 0x88, 0xd7, 0xb6, 0x18, 0x0a, 0xda, 0x50, 0xc0, 0xd3, 0xf2, 0x0d, 0x68, 0xab, 0x50, 0x64,
	0x5c, 0xe1, 0xc2, 0x7c, 0x0c, 0x4d, 0xcf, 0x49, 0x1d, 0x3b, 0x88, 0x5c, 0x27, 0xf0, 0xd3, 0x37,
	0xdd, 0x0e, 0x4a, 0xab, 0x66, 0x83, 0xc0, 0xb1, 0xc6, 0x70, 0x49, 0xea, 0x3a, 0xce, 0x9c, 0x9f,
	0xa3, 0x24, 0x0f, 0x33, 0x5e, 0x8a, 0x54, 0x74, 0xb7, 0xc9, 0xf5, 0xa3, 0xca, 0x43, 0x75, 0x23,
	0x4f, 0x49, 0x31, 0xcc, 0x05, 0x97, 0x82, 0x5d, 0x64, 0x99, 0xbc, 0x72, 0x61, 0xec, 0xfd, 0x57,
	0x53, 0x65, 0xdb, 0xf0, 0x9d, 0x25, 0xf7, 0x7e, 0xab, 0xc0, 0x8e, 0xc9, 0xe3, 0xc0, 0x77, 0x9d,
	0x54, 0xdf, 0x53, 0xcb, 0x0f, 0x5f, 0x60, 0xa6, 0x31, 0xe7, 0x8a, 0x9f, 0x4f, 0xc7, 0xb8, 0x8e,
	0xc7, 0x31, 0xf2, 0xdc, 0x3b, 0x8f, 0x75, 0xb9, 0x58, 0xc1, 0xb0, 0xaf, 0xe1, 0x56, 0xea, 0xcf,
	0xb8, 0x45, 0x4f, 0x87, 0x58, 0x36, 0x53, 0xe5, 0xea, 0xfd, 0x82, 0xde, 0xdf, 0x15, 0xd8, 0x5d,
	0x5e, 0x45, 0x34, 0x4f, 0x5c, 0x4e, 0x77, 0x3e, 0xe6, 0xf8, 0x46, 0x1e, 0xc9, 0xb9, 0xb7, 0x4c,
	0xfd, 0x55, 0x5a, 0x9f, 0x75, 0xed, 0xc7, 0x71, 0x61, 0xa2, 0x15, 0x0c, 0xbb, 0x0b, 0x2d, 0x3a,
	0x75, 0x84, 0xa3, 0xab, 0x1f, 0xe6, 0x7c, 0xce, 0xb1, 0x4e, 0x54, 0xe9, 0x62, 0x96, 0x51, 0x3a,
	0x9d, 0x7c, 0x99, 0xcb, 0xde, 0xeb, 0xd2, 0xfb, 0x7b, 0x79, 0x9a, 0x23, 0x29, 0x6c, 0xc2, 0xb9,
	0xc2, 0x3a, 0x41, 0x16, 0x4b, 0x68, 0xef, 0x8f, 0x3a, 0x80, 0xc5, 0x93, 0x57, 0x3c, 0x91, 0x65,
	0xf1, 0x01, 0xb0, 0x70, 0x3e, 0x9b, 0xf0, 0xc4, 0x8e, 0xa6, 0xf9, 0x45, 0xc5, 0xed, 0x52, 0x96,
	0x75, 0x14, 0x73, 0x3e, 0xcd, 0x6e, 0x29, 0x55, 0x73, 0x95, 0xd1, 0x2b, 0x6c, 0xaa, 0xd2, 0x66,
	0x57, 0xf2, 0x67, 0xcb, 0x86, 0x77, 0xa0, 0x31, 0xa7, 0xba, 0x77, 0xcd, 0x9d, 0x98, 0xea, 0x8c,
	0xae, 0x93, 0x84, 0x3d, 0x45, 0x08, 0xab, 0xcb, 0x6d, 0x30, 0x66, 0xce, 0xeb, 0x5c, 0xa0, 0x0a,
	0xe3, 0x16, 0x42, 0x9a, 0xff, 0x2c, 0x7f, 0xec, 0x02, 0x5c, 0xb7, 0xc0, 0xdd, 0xd5, 0xca, 0x6d,
	0x08, 0x6d, 0x27, 0x7b, 0xeb, 0x68, 0x2c, 0xd8, 0xe7, 0xd0, 0x70, 0x23, 0xec, 0x69, 0x5c, 0x2c,
	0xa9, 0x51, 0x22, 0xb0, 0x1a, 0xd6, 0xb2, 0xf7, 0x71, 0xb8, 0xc0, 0xcd, 0x92, 0x88, 0x4a, 0x2e,
	0x46, 0x2c, 0x4a, 0x52, 0xba, 0xb5, 0xf8, 0x4b, 0x01, 0x97, 0x65, 0xb1, 0x4e, 0x6f, 0x02, 0x11,
	0x16, 0xe1, 0x97, 0x08, 0x63, 0xcc, 0x35, 0x64, 0xf3, 0xd0, 0x53, 0x4a, 0x55, 0x0d, 0x9b, 0x0a,
	0x3e, 0x0e, 0x3d, 0xa9, 0xc3, 0xd2, 0x4c, 0xad, 0x13, 0xde, 0x17, 0x8a, 0xbb, 0x4d, 0x94, 0xac,
	0x82, 0x98, 0x01, 0x84, 0xab, 0xe3, 0xb8, 0x40, 0x94, 0x3d, 0x56, 0xa7, 0xb8, 0xc8, 0x41, 0xac,
	0x7d, 0xb4, 0xe8, 0x5b, 0x6a, 0x9f, 0x2b, 0x92, 0xd4, 0x5c, 0x32, 0xc0, 0x07, 0xa7, 0x91, 0x23,
	0x78, 0x99, 0x64, 0x25, 0x34, 0x0e, 0x6f, 0xae, 0x72, 0x80, 0xb4, 0x59, 0x12, 0xf7, 0x7c, 0x68,
	0x8d, 0xfd, 0x57, 0x5c, 0xad, 0x48, 0x36, 0x73, 0x77, 0xf1, 0xdd, 0x93, 0x5f, 0xc5, 0xc6, 0x4f,
	0xf1, 0xd4, 0x1c, 0x98, 0x9a, 0xa5, 0x2e, 0x51, 0x6f, 0x8f, 0xce, 0x47, 0xf7, 0x22, 0x05, 0xb1,
	0x3c, 0x1e, 0x10, 0xf9, 0xb8, 0xf7, 0x4f, 0x0d, 0x9a, 0xa5, 0x56, 0x95, 0x7d, 0x07, 0xcd, 0xeb,
	0x09, 0xb6, 0xa8, 0x36, 0x4a, 0x04, 0xae, 0x51, 0xa6, 0xa1, 0x71, 0xf8, 0x7f, 0x72, 0x22, 0x7b,
	0xd7, 0xe7, 0x0a, 0x3f, 0xc1, 0x87, 0x69, 0x18, 0x85, 0x29, 0x16, 0x76, 0xb3, 0x21, 0x2d, 0x34,
	0x81, 0x2f, 0x5b, 0x23, 0xc0, 0xe5, 0xeb, 0x40, 0x0b, 0xdd, 0x2f, 0x30, 0x72, 0x50, 0xde, 0x96,
	0x69, 0x04, 0xf9, 0xb7, 0xa0, 0xdc, 0xf2, 0xe8, 0x65, 0xcb, 0xcc, 0x6a, 0x8b, 0xdc, 0x2a, 0xec,
	0xd4, 0x20, 0x4d, 0x66, 0xf2, 0x3d, 0xec, 0xaa, 0x54, 0x13, 0xf8, 0x24, 0xd8, 0x69, 0xde, 0xbd,
	0x61, 0xe2, 0x92, 0xed, 0x5e, 0xb1, 0x3d, 0x5e, 0xf4, 0x76, 0xe6, 0x8e, 0x36, 0x2a, 0x35, 0x7c,
	0x0f, 0x00, 0x5c, 0x15, 0x08, 0x2a, 0xa4, 0x6b, 0x72, 0xd3, 0x4d, 0x99, 0xa5, 0x59, 0x8f, 0x6e,
	0x6e, 0xb9, 0xd9, 0x10, 0xa3, 0xb4, 0x33, 0x73, 0xa4, 0xf8, 0xbf, 0x24, 0x37, 0x53, 0xda, 0x61,
	0x31, 0xc5, 0xf1, 0x48, 0x15, 0x2a, 0xf3, 0x7a, 0xc5, 0x91, 0x2a, 0x16, 0xa3, 0xd9, 0x9a, 0x38,
	0xee, 0x8b, 0x79, 0x6c, 0x2b, 0x40, 0x60, 0x76, 0xaf, 0x0a, 0x4c, 0x53, 0xa9, 0x4e, 0x95, 0x88,
	0xda, 0xc5, 0x89, 0x13, 0x38, 0x58, 0xe2, 0xb1, 0x40, 0x84, 0x32, 0xd1, 0x37, 0x4d, 0xc8, 0xa0,
	0xf3, 0x70, 0x70, 0x02, 0xf7, 0xa2, 0xe4, 0xaa, 0xef, 0xe0, 0xd3, 0x72, 0xcd, 0xfb, 0xd7, 0x8e,
	0x17, 0x45, 0x71, 0x5f, 0x9e, 0xa3, 0xfa, 0x07, 0x32, 0x99, 0x4f, 0xfb, 0x57, 0x3c, 0xe4, 0x09,
	0x36, 0xac, 0xde, 0xa0, 0xfc, 0x7f, 0xe6, 0x82, 0x04, 0xe2, 0x69, 0xe5, 0x6d, 0xa5, 0xf2, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x04, 0x73, 0x7c, 0x28, 0xea, 0x0c, 0x00, 0x00,
}
