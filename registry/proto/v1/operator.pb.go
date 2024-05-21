// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: operator.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var (
	file_operator_proto_rawDescOnce sync.Once
	file_operator_proto_rawDescData = file_operator_proto_rawDesc
)

var File_operator_proto protoreflect.FileDescriptor

var file_operator_proto_depIdxs = []int32{
	2, // 0: registry.node_operator.v1.NodeOperatorRecord.rewardable_nodes:type_name -> registry.node_operator.v1.NodeOperatorRecord.RewardableNodesEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

var file_operator_proto_goTypes = []interface{}{
	(*NodeOperatorRecord)(nil),         // 0: registry.node_operator.v1.NodeOperatorRecord
	(*RemoveNodeOperatorsPayload)(nil), // 1: registry.node_operator.v1.RemoveNodeOperatorsPayload
	nil,                                // 2: registry.node_operator.v1.NodeOperatorRecord.RewardableNodesEntry
}

var file_operator_proto_msgTypes = make([]protoimpl.MessageInfo, 3)

var file_operator_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x9f, 0x03, 0x0a, 0x12,
	0x4e, 0x6f, 0x64, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x3b, 0x0a, 0x1a, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x5f, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x17, 0x6e, 0x6f, 0x64, 0x65, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x49, 0x64, 0x12,
	0x25, 0x0a, 0x0e, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x6c, 0x6c,
	0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x1a, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x17, 0x6e, 0x6f, 0x64, 0x65,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61,
	0x6c, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x64, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x63, 0x49, 0x64, 0x12, 0x6d, 0x0a, 0x10, 0x72, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x42, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x61, 0x62,
	0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x36, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x69, 0x70, 0x76, 0x36, 0x88, 0x01, 0x01,
	0x1a, 0x42, 0x0a, 0x14, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x69, 0x70, 0x76, 0x36, 0x22, 0x55, 0x0a,
	0x1a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x37, 0x0a, 0x18, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x5f, 0x74, 0x6f,
	0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x15, 0x6e,
	0x6f, 0x64, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x54, 0x6f, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

func file_operator_proto_init() {
	if File_operator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_operator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeOperatorRecord); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_operator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveNodeOperatorsPayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_operator_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_operator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_operator_proto_goTypes,
		DependencyIndexes: file_operator_proto_depIdxs,
		MessageInfos:      file_operator_proto_msgTypes,
	}.Build()
	File_operator_proto = out.File
	file_operator_proto_rawDesc = nil
	file_operator_proto_goTypes = nil
	file_operator_proto_depIdxs = nil
}

func file_operator_proto_rawDescGZIP() []byte {
	file_operator_proto_rawDescOnce.Do(func() {
		file_operator_proto_rawDescData = protoimpl.X.CompressGZIP(file_operator_proto_rawDescData)
	})
	return file_operator_proto_rawDescData
}

func init() { file_operator_proto_init() }

// A record for a node operator. Each node operator is associated with a
// unique principal id, a.k.a. NOID.
//
// Note that while a node operator might host nodes for more than
// one funding partner, its principal ID must be unique.
type NodeOperatorRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The principal id of the node operator. This principal is the entity that
	// is able to add and remove nodes.
	//
	// This must be unique across NodeOperatorRecords.
	NodeOperatorPrincipalId []byte `protobuf:"bytes,1,opt,name=node_operator_principal_id,json=nodeOperatorPrincipalId,proto3" json:"node_operator_principal_id,omitempty"`
	// The remaining number of nodes that could be added by this node operator.
	// This number should never go below 0.
	NodeAllowance uint64 `protobuf:"varint,2,opt,name=node_allowance,json=nodeAllowance,proto3" json:"node_allowance,omitempty"`
	// The principal id of this node operator's provider.
	NodeProviderPrincipalId []byte `protobuf:"bytes,3,opt,name=node_provider_principal_id,json=nodeProviderPrincipalId,proto3" json:"node_provider_principal_id,omitempty"`
	// The ID of the data center where this Node Operator hosts nodes.
	DcId string `protobuf:"bytes,4,opt,name=dc_id,json=dcId,proto3" json:"dc_id,omitempty"`
	// A map from node type to the number of nodes for which the associated Node
	// Provider should be rewarded.
	RewardableNodes map[string]uint32 `protobuf:"bytes,5,rep,name=rewardable_nodes,json=rewardableNodes,proto3" json:"rewardable_nodes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Ipv6            *string           `protobuf:"bytes,6,opt,name=ipv6,proto3,oneof" json:"ipv6,omitempty"`
}

// Deprecated: Use NodeOperatorRecord.ProtoReflect.Descriptor instead.
func (*NodeOperatorRecord) Descriptor() ([]byte, []int) {
	return file_operator_proto_rawDescGZIP(), []int{0}
}

func (x *NodeOperatorRecord) GetDcId() string {
	if x != nil {
		return x.DcId
	}
	return ""
}

func (x *NodeOperatorRecord) GetIpv6() string {
	if x != nil && x.Ipv6 != nil {
		return *x.Ipv6
	}
	return ""
}

func (x *NodeOperatorRecord) GetNodeAllowance() uint64 {
	if x != nil {
		return x.NodeAllowance
	}
	return 0
}

func (x *NodeOperatorRecord) GetNodeOperatorPrincipalId() []byte {
	if x != nil {
		return x.NodeOperatorPrincipalId
	}
	return nil
}

func (x *NodeOperatorRecord) GetNodeProviderPrincipalId() []byte {
	if x != nil {
		return x.NodeProviderPrincipalId
	}
	return nil
}

func (x *NodeOperatorRecord) GetRewardableNodes() map[string]uint32 {
	if x != nil {
		return x.RewardableNodes
	}
	return nil
}

func (*NodeOperatorRecord) ProtoMessage() {}

func (x *NodeOperatorRecord) ProtoReflect() protoreflect.Message {
	mi := &file_operator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *NodeOperatorRecord) Reset() {
	*x = NodeOperatorRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeOperatorRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

// The payload of a request to remove Node Operator records from the Registry
type RemoveNodeOperatorsPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeOperatorsToRemove [][]byte `protobuf:"bytes,1,rep,name=node_operators_to_remove,json=nodeOperatorsToRemove,proto3" json:"node_operators_to_remove,omitempty"`
}

// Deprecated: Use RemoveNodeOperatorsPayload.ProtoReflect.Descriptor instead.
func (*RemoveNodeOperatorsPayload) Descriptor() ([]byte, []int) {
	return file_operator_proto_rawDescGZIP(), []int{1}
}

func (x *RemoveNodeOperatorsPayload) GetNodeOperatorsToRemove() [][]byte {
	if x != nil {
		return x.NodeOperatorsToRemove
	}
	return nil
}
func (*RemoveNodeOperatorsPayload) ProtoMessage() {}
func (x *RemoveNodeOperatorsPayload) ProtoReflect() protoreflect.Message {
	mi := &file_operator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *RemoveNodeOperatorsPayload) Reset() {
	*x = RemoveNodeOperatorsPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}
func (x *RemoveNodeOperatorsPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}
