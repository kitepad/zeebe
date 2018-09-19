// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Partition_PartitionBrokerRole int32

const (
	Partition_LEADER Partition_PartitionBrokerRole = 0
	Partition_FOLLOW Partition_PartitionBrokerRole = 1
)

var Partition_PartitionBrokerRole_name = map[int32]string{
	0: "LEADER",
	1: "FOLLOW",
}
var Partition_PartitionBrokerRole_value = map[string]int32{
	"LEADER": 0,
	"FOLLOW": 1,
}

func (x Partition_PartitionBrokerRole) String() string {
	return proto.EnumName(Partition_PartitionBrokerRole_name, int32(x))
}
func (Partition_PartitionBrokerRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_gateway_551de5e543f276fa, []int{1, 0}
}

type WorkflowRequestObject_ResourceType int32

const (
	WorkflowRequestObject_BPMN WorkflowRequestObject_ResourceType = 0
	WorkflowRequestObject_YAML WorkflowRequestObject_ResourceType = 1
)

var WorkflowRequestObject_ResourceType_name = map[int32]string{
	0: "BPMN",
	1: "YAML",
}
var WorkflowRequestObject_ResourceType_value = map[string]int32{
	"BPMN": 0,
	"YAML": 1,
}

func (x WorkflowRequestObject_ResourceType) String() string {
	return proto.EnumName(WorkflowRequestObject_ResourceType_name, int32(x))
}
func (WorkflowRequestObject_ResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_gateway_551de5e543f276fa, []int{4, 0}
}

type HealthRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthRequest) Reset()         { *m = HealthRequest{} }
func (m *HealthRequest) String() string { return proto.CompactTextString(m) }
func (*HealthRequest) ProtoMessage()    {}
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_gateway_551de5e543f276fa, []int{0}
}
func (m *HealthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthRequest.Unmarshal(m, b)
}
func (m *HealthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthRequest.Marshal(b, m, deterministic)
}
func (dst *HealthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthRequest.Merge(dst, src)
}
func (m *HealthRequest) XXX_Size() int {
	return xxx_messageInfo_HealthRequest.Size(m)
}
func (m *HealthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthRequest proto.InternalMessageInfo

type Partition struct {
	PartitionId          int32                         `protobuf:"varint,1,opt,name=partitionId,proto3" json:"partitionId,omitempty"`
	Role                 Partition_PartitionBrokerRole `protobuf:"varint,3,opt,name=role,proto3,enum=gateway_protocol.Partition_PartitionBrokerRole" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *Partition) Reset()         { *m = Partition{} }
func (m *Partition) String() string { return proto.CompactTextString(m) }
func (*Partition) ProtoMessage()    {}
func (*Partition) Descriptor() ([]byte, []int) {
	return fileDescriptor_gateway_551de5e543f276fa, []int{1}
}
func (m *Partition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Partition.Unmarshal(m, b)
}
func (m *Partition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Partition.Marshal(b, m, deterministic)
}
func (dst *Partition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Partition.Merge(dst, src)
}
func (m *Partition) XXX_Size() int {
	return xxx_messageInfo_Partition.Size(m)
}
func (m *Partition) XXX_DiscardUnknown() {
	xxx_messageInfo_Partition.DiscardUnknown(m)
}

var xxx_messageInfo_Partition proto.InternalMessageInfo

func (m *Partition) GetPartitionId() int32 {
	if m != nil {
		return m.PartitionId
	}
	return 0
}

func (m *Partition) GetRole() Partition_PartitionBrokerRole {
	if m != nil {
		return m.Role
	}
	return Partition_LEADER
}

type BrokerInfo struct {
	Host                 string       `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                 int32        `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Partitions           []*Partition `protobuf:"bytes,3,rep,name=partitions,proto3" json:"partitions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BrokerInfo) Reset()         { *m = BrokerInfo{} }
func (m *BrokerInfo) String() string { return proto.CompactTextString(m) }
func (*BrokerInfo) ProtoMessage()    {}
func (*BrokerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_gateway_551de5e543f276fa, []int{2}
}
func (m *BrokerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BrokerInfo.Unmarshal(m, b)
}
func (m *BrokerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BrokerInfo.Marshal(b, m, deterministic)
}
func (dst *BrokerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BrokerInfo.Merge(dst, src)
}
func (m *BrokerInfo) XXX_Size() int {
	return xxx_messageInfo_BrokerInfo.Size(m)
}
func (m *BrokerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BrokerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BrokerInfo proto.InternalMessageInfo

func (m *BrokerInfo) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *BrokerInfo) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *BrokerInfo) GetPartitions() []*Partition {
	if m != nil {
		return m.Partitions
	}
	return nil
}

type HealthResponse struct {
	Brokers              []*BrokerInfo `protobuf:"bytes,1,rep,name=brokers,proto3" json:"brokers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *HealthResponse) Reset()         { *m = HealthResponse{} }
func (m *HealthResponse) String() string { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()    {}
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_gateway_551de5e543f276fa, []int{3}
}
func (m *HealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthResponse.Unmarshal(m, b)
}
func (m *HealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthResponse.Marshal(b, m, deterministic)
}
func (dst *HealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResponse.Merge(dst, src)
}
func (m *HealthResponse) XXX_Size() int {
	return xxx_messageInfo_HealthResponse.Size(m)
}
func (m *HealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResponse proto.InternalMessageInfo

func (m *HealthResponse) GetBrokers() []*BrokerInfo {
	if m != nil {
		return m.Brokers
	}
	return nil
}

type WorkflowRequestObject struct {
	Name                 string                             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 WorkflowRequestObject_ResourceType `protobuf:"varint,2,opt,name=type,proto3,enum=gateway_protocol.WorkflowRequestObject_ResourceType" json:"type,omitempty"`
	Definition           []byte                             `protobuf:"bytes,3,opt,name=definition,proto3" json:"definition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *WorkflowRequestObject) Reset()         { *m = WorkflowRequestObject{} }
func (m *WorkflowRequestObject) String() string { return proto.CompactTextString(m) }
func (*WorkflowRequestObject) ProtoMessage()    {}
func (*WorkflowRequestObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_gateway_551de5e543f276fa, []int{4}
}
func (m *WorkflowRequestObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowRequestObject.Unmarshal(m, b)
}
func (m *WorkflowRequestObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowRequestObject.Marshal(b, m, deterministic)
}
func (dst *WorkflowRequestObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowRequestObject.Merge(dst, src)
}
func (m *WorkflowRequestObject) XXX_Size() int {
	return xxx_messageInfo_WorkflowRequestObject.Size(m)
}
func (m *WorkflowRequestObject) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowRequestObject.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowRequestObject proto.InternalMessageInfo

func (m *WorkflowRequestObject) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WorkflowRequestObject) GetType() WorkflowRequestObject_ResourceType {
	if m != nil {
		return m.Type
	}
	return WorkflowRequestObject_BPMN
}

func (m *WorkflowRequestObject) GetDefinition() []byte {
	if m != nil {
		return m.Definition
	}
	return nil
}

type DeployWorkflowRequest struct {
	Workflows            []*WorkflowRequestObject `protobuf:"bytes,1,rep,name=workflows,proto3" json:"workflows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *DeployWorkflowRequest) Reset()         { *m = DeployWorkflowRequest{} }
func (m *DeployWorkflowRequest) String() string { return proto.CompactTextString(m) }
func (*DeployWorkflowRequest) ProtoMessage()    {}
func (*DeployWorkflowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_gateway_551de5e543f276fa, []int{5}
}
func (m *DeployWorkflowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployWorkflowRequest.Unmarshal(m, b)
}
func (m *DeployWorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployWorkflowRequest.Marshal(b, m, deterministic)
}
func (dst *DeployWorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployWorkflowRequest.Merge(dst, src)
}
func (m *DeployWorkflowRequest) XXX_Size() int {
	return xxx_messageInfo_DeployWorkflowRequest.Size(m)
}
func (m *DeployWorkflowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployWorkflowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeployWorkflowRequest proto.InternalMessageInfo

func (m *DeployWorkflowRequest) GetWorkflows() []*WorkflowRequestObject {
	if m != nil {
		return m.Workflows
	}
	return nil
}

type WorkflowResponseObject struct {
	BpmnProcessId        string   `protobuf:"bytes,1,opt,name=bpmnProcessId,proto3" json:"bpmnProcessId,omitempty"`
	Version              int32    `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	WorkflowKey          int64    `protobuf:"varint,3,opt,name=workflowKey,proto3" json:"workflowKey,omitempty"`
	ResourceName         string   `protobuf:"bytes,4,opt,name=resourceName,proto3" json:"resourceName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowResponseObject) Reset()         { *m = WorkflowResponseObject{} }
func (m *WorkflowResponseObject) String() string { return proto.CompactTextString(m) }
func (*WorkflowResponseObject) ProtoMessage()    {}
func (*WorkflowResponseObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_gateway_551de5e543f276fa, []int{6}
}
func (m *WorkflowResponseObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowResponseObject.Unmarshal(m, b)
}
func (m *WorkflowResponseObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowResponseObject.Marshal(b, m, deterministic)
}
func (dst *WorkflowResponseObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowResponseObject.Merge(dst, src)
}
func (m *WorkflowResponseObject) XXX_Size() int {
	return xxx_messageInfo_WorkflowResponseObject.Size(m)
}
func (m *WorkflowResponseObject) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowResponseObject.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowResponseObject proto.InternalMessageInfo

func (m *WorkflowResponseObject) GetBpmnProcessId() string {
	if m != nil {
		return m.BpmnProcessId
	}
	return ""
}

func (m *WorkflowResponseObject) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *WorkflowResponseObject) GetWorkflowKey() int64 {
	if m != nil {
		return m.WorkflowKey
	}
	return 0
}

func (m *WorkflowResponseObject) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

type DeployWorkflowResponse struct {
	Workflows            []*WorkflowResponseObject `protobuf:"bytes,1,rep,name=workflows,proto3" json:"workflows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *DeployWorkflowResponse) Reset()         { *m = DeployWorkflowResponse{} }
func (m *DeployWorkflowResponse) String() string { return proto.CompactTextString(m) }
func (*DeployWorkflowResponse) ProtoMessage()    {}
func (*DeployWorkflowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_gateway_551de5e543f276fa, []int{7}
}
func (m *DeployWorkflowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployWorkflowResponse.Unmarshal(m, b)
}
func (m *DeployWorkflowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployWorkflowResponse.Marshal(b, m, deterministic)
}
func (dst *DeployWorkflowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployWorkflowResponse.Merge(dst, src)
}
func (m *DeployWorkflowResponse) XXX_Size() int {
	return xxx_messageInfo_DeployWorkflowResponse.Size(m)
}
func (m *DeployWorkflowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployWorkflowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeployWorkflowResponse proto.InternalMessageInfo

func (m *DeployWorkflowResponse) GetWorkflows() []*WorkflowResponseObject {
	if m != nil {
		return m.Workflows
	}
	return nil
}

func init() {
	proto.RegisterType((*HealthRequest)(nil), "gateway_protocol.HealthRequest")
	proto.RegisterType((*Partition)(nil), "gateway_protocol.Partition")
	proto.RegisterType((*BrokerInfo)(nil), "gateway_protocol.BrokerInfo")
	proto.RegisterType((*HealthResponse)(nil), "gateway_protocol.HealthResponse")
	proto.RegisterType((*WorkflowRequestObject)(nil), "gateway_protocol.WorkflowRequestObject")
	proto.RegisterType((*DeployWorkflowRequest)(nil), "gateway_protocol.DeployWorkflowRequest")
	proto.RegisterType((*WorkflowResponseObject)(nil), "gateway_protocol.WorkflowResponseObject")
	proto.RegisterType((*DeployWorkflowResponse)(nil), "gateway_protocol.DeployWorkflowResponse")
	proto.RegisterEnum("gateway_protocol.Partition_PartitionBrokerRole", Partition_PartitionBrokerRole_name, Partition_PartitionBrokerRole_value)
	proto.RegisterEnum("gateway_protocol.WorkflowRequestObject_ResourceType", WorkflowRequestObject_ResourceType_name, WorkflowRequestObject_ResourceType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayClient interface {
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	DeployWorkflow(ctx context.Context, in *DeployWorkflowRequest, opts ...grpc.CallOption) (*DeployWorkflowResponse, error)
}

type gatewayClient struct {
	cc *grpc.ClientConn
}

func NewGatewayClient(cc *grpc.ClientConn) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/gateway_protocol.Gateway/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeployWorkflow(ctx context.Context, in *DeployWorkflowRequest, opts ...grpc.CallOption) (*DeployWorkflowResponse, error) {
	out := new(DeployWorkflowResponse)
	err := c.cc.Invoke(ctx, "/gateway_protocol.Gateway/DeployWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
type GatewayServer interface {
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
	DeployWorkflow(context.Context, *DeployWorkflowRequest) (*DeployWorkflowResponse, error)
}

func RegisterGatewayServer(s *grpc.Server, srv GatewayServer) {
	s.RegisterService(&_Gateway_serviceDesc, srv)
}

func _Gateway_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway_protocol.Gateway/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeployWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeployWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway_protocol.Gateway/DeployWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeployWorkflow(ctx, req.(*DeployWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gateway_protocol.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _Gateway_Health_Handler,
		},
		{
			MethodName: "DeployWorkflow",
			Handler:    _Gateway_DeployWorkflow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway.proto",
}

func init() { proto.RegisterFile("gateway.proto", fileDescriptor_gateway_551de5e543f276fa) }

var fileDescriptor_gateway_551de5e543f276fa = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xcd, 0x36, 0x69, 0x82, 0xa7, 0x49, 0x88, 0x16, 0xb5, 0x32, 0xa5, 0x02, 0xb3, 0x42, 0xc2,
	0x17, 0x8c, 0x14, 0x10, 0x17, 0x4e, 0x0d, 0x4d, 0x49, 0x45, 0xd2, 0x44, 0x2b, 0xa4, 0x0a, 0x0e,
	0x40, 0xec, 0x4e, 0x68, 0xa8, 0xeb, 0x75, 0xd7, 0x5b, 0xa2, 0xf0, 0x31, 0xfc, 0x07, 0xa7, 0xfe,
	0x1a, 0xf2, 0xda, 0x6e, 0xec, 0x24, 0x94, 0xde, 0x66, 0x5e, 0x66, 0xe7, 0xbd, 0x79, 0x33, 0x31,
	0x34, 0xbe, 0x8f, 0x15, 0xce, 0xc6, 0x73, 0x27, 0x94, 0x42, 0x09, 0xda, 0x4a, 0xd3, 0xaf, 0x3a,
	0xf5, 0x84, 0xcf, 0xee, 0x43, 0xa3, 0x87, 0x63, 0x5f, 0x9d, 0x71, 0xbc, 0xbc, 0xc2, 0x48, 0xb1,
	0x6b, 0x02, 0xc6, 0x68, 0x2c, 0xd5, 0x54, 0x4d, 0x45, 0x40, 0x2d, 0xd8, 0x0a, 0xb3, 0xe4, 0xe8,
	0xd4, 0x24, 0x16, 0xb1, 0x37, 0x79, 0x1e, 0xa2, 0x7b, 0x60, 0x28, 0x11, 0x4e, 0xbd, 0xe3, 0xf1,
	0x05, 0x9a, 0x1b, 0x16, 0xb1, 0x0d, 0xbe, 0x00, 0xe8, 0x3b, 0xa8, 0x48, 0xe1, 0xa3, 0x59, 0xb6,
	0x88, 0xdd, 0x6c, 0xbf, 0x74, 0x96, 0xf9, 0x9d, 0x1b, 0xaa, 0x45, 0xd4, 0x91, 0xe2, 0x1c, 0x25,
	0x17, 0x3e, 0x72, 0xfd, 0x98, 0xbd, 0x80, 0x07, 0x6b, 0x7e, 0xa4, 0x00, 0xd5, 0x7e, 0x77, 0xff,
	0xa0, 0xcb, 0x5b, 0xa5, 0x38, 0x3e, 0x1c, 0xf6, 0xfb, 0xc3, 0x93, 0x16, 0x61, 0x97, 0x00, 0x49,
	0xd5, 0x51, 0x30, 0x11, 0x94, 0x42, 0xe5, 0x4c, 0x44, 0x4a, 0x4b, 0x37, 0xb8, 0x8e, 0x63, 0x2c,
	0x14, 0x52, 0x69, 0xb9, 0x9b, 0x5c, 0xc7, 0xf4, 0x2d, 0xc0, 0xcd, 0x58, 0x91, 0x59, 0xb6, 0xca,
	0xf6, 0x56, 0xfb, 0xd1, 0x2d, 0x7a, 0x79, 0xae, 0x9c, 0xf5, 0xa0, 0x99, 0xb9, 0x18, 0x85, 0x22,
	0x88, 0x90, 0xbe, 0x81, 0x9a, 0xab, 0x45, 0x44, 0x26, 0xd1, 0xbd, 0xf6, 0x56, 0x7b, 0x2d, 0x54,
	0xf2, 0xac, 0x98, 0xfd, 0x21, 0xb0, 0x7d, 0x22, 0xe4, 0xf9, 0xc4, 0x17, 0xb3, 0x74, 0x25, 0x43,
	0xf7, 0x07, 0x7a, 0x5a, 0x74, 0x10, 0x7b, 0x9c, 0x0e, 0x12, 0xc7, 0xb4, 0x07, 0x15, 0x35, 0x0f,
	0x13, 0xdf, 0x9b, 0xed, 0xd7, 0xab, 0x14, 0x6b, 0x5b, 0x39, 0x1c, 0x23, 0x71, 0x25, 0x3d, 0xfc,
	0x38, 0x0f, 0x91, 0xeb, 0x0e, 0xf4, 0x31, 0xc0, 0x29, 0x4e, 0xa6, 0x81, 0x1e, 0x48, 0xaf, 0xab,
	0xce, 0x73, 0x08, 0x63, 0x50, 0xcf, 0xbf, 0xa2, 0xf7, 0xa0, 0xd2, 0x19, 0x0d, 0x8e, 0x5b, 0xa5,
	0x38, 0xfa, 0xb4, 0x3f, 0xe8, 0xb7, 0x08, 0xfb, 0x02, 0xdb, 0x07, 0x18, 0xfa, 0x62, 0xbe, 0xc4,
	0x4a, 0xbb, 0x60, 0xcc, 0x52, 0x28, 0xb3, 0xe3, 0xf9, 0x1d, 0xb5, 0xf2, 0xc5, 0x4b, 0xf6, 0x9b,
	0xc0, 0xce, 0xa2, 0x28, 0x31, 0x3a, 0x35, 0xe7, 0x19, 0x34, 0xdc, 0xf0, 0x22, 0x18, 0x49, 0xe1,
	0x61, 0x14, 0xa5, 0x97, 0x6a, 0xf0, 0x22, 0x48, 0x4d, 0xa8, 0xfd, 0x44, 0x19, 0xc5, 0x13, 0x26,
	0xab, 0xcf, 0xd2, 0xf8, 0xce, 0x33, 0x9e, 0x0f, 0x38, 0xd7, 0xf3, 0x97, 0x79, 0x1e, 0xa2, 0x0c,
	0xea, 0x32, 0x35, 0x40, 0x9f, 0x7a, 0x45, 0x13, 0x14, 0x30, 0xf6, 0x0d, 0x76, 0x96, 0x0d, 0x48,
	0xcf, 0xe1, 0x70, 0xd5, 0x01, 0xfb, 0x36, 0x07, 0xf2, 0xc3, 0xe5, 0x2c, 0x68, 0x5f, 0x13, 0xa8,
	0xbd, 0x4f, 0x9e, 0xd1, 0x01, 0x54, 0x93, 0xa3, 0xa3, 0x4f, 0x56, 0x5b, 0x15, 0xfe, 0xd4, 0xbb,
	0xd6, 0xbf, 0x0b, 0x12, 0x26, 0x56, 0xa2, 0x08, 0xcd, 0xa2, 0x78, 0xba, 0x66, 0x47, 0x6b, 0xf7,
	0xbb, 0x6b, 0xff, 0xbf, 0x30, 0xa3, 0xe9, 0x3c, 0x85, 0x87, 0x53, 0xe1, 0xfc, 0x42, 0x74, 0xd1,
	0x29, 0x7c, 0x9c, 0x3c, 0xe1, 0x8f, 0x4a, 0x9f, 0x37, 0x42, 0xd7, 0xad, 0xea, 0xfc, 0xd5, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xbb, 0xec, 0x58, 0xbd, 0x04, 0x00, 0x00,
}