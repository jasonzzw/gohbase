// Code generated by protoc-gen-go.
// source: MapReduce.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ScanMetrics struct {
	Metrics          []*NameInt64Pair `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *ScanMetrics) Reset()                    { *m = ScanMetrics{} }
func (m *ScanMetrics) String() string            { return proto.CompactTextString(m) }
func (*ScanMetrics) ProtoMessage()               {}
func (*ScanMetrics) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{0} }

func (m *ScanMetrics) GetMetrics() []*NameInt64Pair {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type TableSnapshotRegionSplit struct {
	Locations        []string     `protobuf:"bytes,2,rep,name=locations" json:"locations,omitempty"`
	Table            *TableSchema `protobuf:"bytes,3,opt,name=table" json:"table,omitempty"`
	Region           *RegionInfo  `protobuf:"bytes,4,opt,name=region" json:"region,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *TableSnapshotRegionSplit) Reset()                    { *m = TableSnapshotRegionSplit{} }
func (m *TableSnapshotRegionSplit) String() string            { return proto.CompactTextString(m) }
func (*TableSnapshotRegionSplit) ProtoMessage()               {}
func (*TableSnapshotRegionSplit) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{1} }

func (m *TableSnapshotRegionSplit) GetLocations() []string {
	if m != nil {
		return m.Locations
	}
	return nil
}

func (m *TableSnapshotRegionSplit) GetTable() *TableSchema {
	if m != nil {
		return m.Table
	}
	return nil
}

func (m *TableSnapshotRegionSplit) GetRegion() *RegionInfo {
	if m != nil {
		return m.Region
	}
	return nil
}

func init() {
	proto.RegisterType((*ScanMetrics)(nil), "pb.ScanMetrics")
	proto.RegisterType((*TableSnapshotRegionSplit)(nil), "pb.TableSnapshotRegionSplit")
}

var fileDescriptor15 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0x90, 0x4d, 0x4b, 0xc3, 0x30,
	0x18, 0xc7, 0xc9, 0xe6, 0x0b, 0x4b, 0xc1, 0x61, 0x4e, 0x41, 0x3c, 0x8c, 0x81, 0x22, 0x0a, 0x39,
	0x88, 0x78, 0xf0, 0xd8, 0xd3, 0x76, 0x98, 0x8c, 0xd4, 0x2f, 0xf0, 0x24, 0x7d, 0xd6, 0x16, 0xda,
	0x3c, 0x21, 0xc9, 0x3e, 0x83, 0x5f, 0xc3, 0x8f, 0x6a, 0xbb, 0x4e, 0x7b, 0x0b, 0xff, 0xb7, 0x1f,
	0x79, 0xf8, 0x72, 0x07, 0x5e, 0x63, 0x79, 0xb4, 0xa8, 0x7c, 0xa0, 0x44, 0x62, 0xe6, 0xcd, 0x5d,
	0xb6, 0xc9, 0x21, 0x9e, 0x85, 0xf5, 0x07, 0xcf, 0x0a, 0x0b, 0x6e, 0x87, 0x29, 0x34, 0x36, 0x8a,
	0x17, 0x7e, 0xdd, 0x8d, 0x4f, 0xc9, 0x56, 0xf3, 0xa7, 0xec, 0xf5, 0x56, 0x79, 0xa3, 0x3e, 0xa1,
	0xc3, 0xad, 0x4b, 0xef, 0x6f, 0x7b, 0x68, 0x82, 0xfe, 0x4b, 0xac, 0xbf, 0x19, 0x97, 0x5f, 0x60,
	0x5a, 0x2c, 0x1c, 0xf8, 0x58, 0x53, 0xd2, 0x58, 0x35, 0xe4, 0x0a, 0xdf, 0x36, 0x49, 0xdc, 0xf3,
	0x45, 0x4b, 0x16, 0x52, 0x2f, 0x44, 0x39, 0xeb, 0xb7, 0x16, 0x7a, 0x12, 0xc4, 0x03, 0xbf, 0x4c,
	0x43, 0x53, 0xce, 0x57, 0xac, 0xa7, 0x2c, 0x07, 0xca, 0x38, 0x65, 0x6b, 0xec, 0x40, 0x8f, 0xae,
	0x78, 0xe4, 0x57, 0xe1, 0xb4, 0x29, 0x2f, 0x4e, 0xb9, 0x9b, 0x21, 0x37, 0x52, 0xb6, 0xee, 0x40,
	0xfa, 0xec, 0xe6, 0x39, 0x7f, 0xa6, 0x50, 0x29, 0xf0, 0xd0, 0xd7, 0x55, 0x0d, 0x25, 0x91, 0x57,
	0xb5, 0xf9, 0xff, 0xa6, 0x39, 0x1e, 0x54, 0x85, 0x0e, 0x03, 0x24, 0x2c, 0xf3, 0xe9, 0x2a, 0xfb,
	0xc1, 0x8c, 0x1b, 0xf6, 0xc3, 0xd8, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0xa3, 0xaa, 0x1a,
	0x2c, 0x01, 0x00, 0x00,
}
