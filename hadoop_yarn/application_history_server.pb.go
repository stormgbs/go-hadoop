// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/application_history_server.proto

/*
Package hadoop_yarn is a generated protocol buffer package.

It is generated from these files:
	server/application_history_server.proto

It has these top-level messages:
	ApplicationHistoryDataProto
	ApplicationStartDataProto
	ApplicationFinishDataProto
	ApplicationAttemptHistoryDataProto
	ApplicationAttemptStartDataProto
	ApplicationAttemptFinishDataProto
	ContainerHistoryDataProto
	ContainerStartDataProto
	ContainerFinishDataProto
*/
package hadoop_yarn

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ApplicationHistoryDataProto struct {
	ApplicationId          *ApplicationIdProto          `protobuf:"bytes,1,opt,name=application_id,json=applicationId" json:"application_id,omitempty"`
	ApplicationName        *string                      `protobuf:"bytes,2,opt,name=application_name,json=applicationName" json:"application_name,omitempty"`
	ApplicationType        *string                      `protobuf:"bytes,3,opt,name=application_type,json=applicationType" json:"application_type,omitempty"`
	User                   *string                      `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	Queue                  *string                      `protobuf:"bytes,5,opt,name=queue" json:"queue,omitempty"`
	SubmitTime             *int64                       `protobuf:"varint,6,opt,name=submit_time,json=submitTime" json:"submit_time,omitempty"`
	StartTime              *int64                       `protobuf:"varint,7,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	FinishTime             *int64                       `protobuf:"varint,8,opt,name=finish_time,json=finishTime" json:"finish_time,omitempty"`
	DiagnosticsInfo        *string                      `protobuf:"bytes,9,opt,name=diagnostics_info,json=diagnosticsInfo" json:"diagnostics_info,omitempty"`
	FinalApplicationStatus *FinalApplicationStatusProto `protobuf:"varint,10,opt,name=final_application_status,json=finalApplicationStatus,enum=hadoop.yarn.FinalApplicationStatusProto" json:"final_application_status,omitempty"`
	YarnApplicationState   *YarnApplicationStateProto   `protobuf:"varint,11,opt,name=yarn_application_state,json=yarnApplicationState,enum=hadoop.yarn.YarnApplicationStateProto" json:"yarn_application_state,omitempty"`
	XXX_unrecognized       []byte                       `json:"-"`
}

func (m *ApplicationHistoryDataProto) Reset()                    { *m = ApplicationHistoryDataProto{} }
func (m *ApplicationHistoryDataProto) String() string            { return proto.CompactTextString(m) }
func (*ApplicationHistoryDataProto) ProtoMessage()               {}
func (*ApplicationHistoryDataProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ApplicationHistoryDataProto) GetApplicationId() *ApplicationIdProto {
	if m != nil {
		return m.ApplicationId
	}
	return nil
}

func (m *ApplicationHistoryDataProto) GetApplicationName() string {
	if m != nil && m.ApplicationName != nil {
		return *m.ApplicationName
	}
	return ""
}

func (m *ApplicationHistoryDataProto) GetApplicationType() string {
	if m != nil && m.ApplicationType != nil {
		return *m.ApplicationType
	}
	return ""
}

func (m *ApplicationHistoryDataProto) GetUser() string {
	if m != nil && m.User != nil {
		return *m.User
	}
	return ""
}

func (m *ApplicationHistoryDataProto) GetQueue() string {
	if m != nil && m.Queue != nil {
		return *m.Queue
	}
	return ""
}

func (m *ApplicationHistoryDataProto) GetSubmitTime() int64 {
	if m != nil && m.SubmitTime != nil {
		return *m.SubmitTime
	}
	return 0
}

func (m *ApplicationHistoryDataProto) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *ApplicationHistoryDataProto) GetFinishTime() int64 {
	if m != nil && m.FinishTime != nil {
		return *m.FinishTime
	}
	return 0
}

func (m *ApplicationHistoryDataProto) GetDiagnosticsInfo() string {
	if m != nil && m.DiagnosticsInfo != nil {
		return *m.DiagnosticsInfo
	}
	return ""
}

func (m *ApplicationHistoryDataProto) GetFinalApplicationStatus() FinalApplicationStatusProto {
	if m != nil && m.FinalApplicationStatus != nil {
		return *m.FinalApplicationStatus
	}
	return FinalApplicationStatusProto_APP_UNDEFINED
}

func (m *ApplicationHistoryDataProto) GetYarnApplicationState() YarnApplicationStateProto {
	if m != nil && m.YarnApplicationState != nil {
		return *m.YarnApplicationState
	}
	return YarnApplicationStateProto_NEW
}

type ApplicationStartDataProto struct {
	ApplicationId    *ApplicationIdProto `protobuf:"bytes,1,opt,name=application_id,json=applicationId" json:"application_id,omitempty"`
	ApplicationName  *string             `protobuf:"bytes,2,opt,name=application_name,json=applicationName" json:"application_name,omitempty"`
	ApplicationType  *string             `protobuf:"bytes,3,opt,name=application_type,json=applicationType" json:"application_type,omitempty"`
	User             *string             `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	Queue            *string             `protobuf:"bytes,5,opt,name=queue" json:"queue,omitempty"`
	SubmitTime       *int64              `protobuf:"varint,6,opt,name=submit_time,json=submitTime" json:"submit_time,omitempty"`
	StartTime        *int64              `protobuf:"varint,7,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *ApplicationStartDataProto) Reset()                    { *m = ApplicationStartDataProto{} }
func (m *ApplicationStartDataProto) String() string            { return proto.CompactTextString(m) }
func (*ApplicationStartDataProto) ProtoMessage()               {}
func (*ApplicationStartDataProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ApplicationStartDataProto) GetApplicationId() *ApplicationIdProto {
	if m != nil {
		return m.ApplicationId
	}
	return nil
}

func (m *ApplicationStartDataProto) GetApplicationName() string {
	if m != nil && m.ApplicationName != nil {
		return *m.ApplicationName
	}
	return ""
}

func (m *ApplicationStartDataProto) GetApplicationType() string {
	if m != nil && m.ApplicationType != nil {
		return *m.ApplicationType
	}
	return ""
}

func (m *ApplicationStartDataProto) GetUser() string {
	if m != nil && m.User != nil {
		return *m.User
	}
	return ""
}

func (m *ApplicationStartDataProto) GetQueue() string {
	if m != nil && m.Queue != nil {
		return *m.Queue
	}
	return ""
}

func (m *ApplicationStartDataProto) GetSubmitTime() int64 {
	if m != nil && m.SubmitTime != nil {
		return *m.SubmitTime
	}
	return 0
}

func (m *ApplicationStartDataProto) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

type ApplicationFinishDataProto struct {
	ApplicationId          *ApplicationIdProto          `protobuf:"bytes,1,opt,name=application_id,json=applicationId" json:"application_id,omitempty"`
	FinishTime             *int64                       `protobuf:"varint,2,opt,name=finish_time,json=finishTime" json:"finish_time,omitempty"`
	DiagnosticsInfo        *string                      `protobuf:"bytes,3,opt,name=diagnostics_info,json=diagnosticsInfo" json:"diagnostics_info,omitempty"`
	FinalApplicationStatus *FinalApplicationStatusProto `protobuf:"varint,4,opt,name=final_application_status,json=finalApplicationStatus,enum=hadoop.yarn.FinalApplicationStatusProto" json:"final_application_status,omitempty"`
	YarnApplicationState   *YarnApplicationStateProto   `protobuf:"varint,5,opt,name=yarn_application_state,json=yarnApplicationState,enum=hadoop.yarn.YarnApplicationStateProto" json:"yarn_application_state,omitempty"`
	XXX_unrecognized       []byte                       `json:"-"`
}

func (m *ApplicationFinishDataProto) Reset()                    { *m = ApplicationFinishDataProto{} }
func (m *ApplicationFinishDataProto) String() string            { return proto.CompactTextString(m) }
func (*ApplicationFinishDataProto) ProtoMessage()               {}
func (*ApplicationFinishDataProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ApplicationFinishDataProto) GetApplicationId() *ApplicationIdProto {
	if m != nil {
		return m.ApplicationId
	}
	return nil
}

func (m *ApplicationFinishDataProto) GetFinishTime() int64 {
	if m != nil && m.FinishTime != nil {
		return *m.FinishTime
	}
	return 0
}

func (m *ApplicationFinishDataProto) GetDiagnosticsInfo() string {
	if m != nil && m.DiagnosticsInfo != nil {
		return *m.DiagnosticsInfo
	}
	return ""
}

func (m *ApplicationFinishDataProto) GetFinalApplicationStatus() FinalApplicationStatusProto {
	if m != nil && m.FinalApplicationStatus != nil {
		return *m.FinalApplicationStatus
	}
	return FinalApplicationStatusProto_APP_UNDEFINED
}

func (m *ApplicationFinishDataProto) GetYarnApplicationState() YarnApplicationStateProto {
	if m != nil && m.YarnApplicationState != nil {
		return *m.YarnApplicationState
	}
	return YarnApplicationStateProto_NEW
}

type ApplicationAttemptHistoryDataProto struct {
	ApplicationAttemptId        *ApplicationAttemptIdProto        `protobuf:"bytes,1,opt,name=application_attempt_id,json=applicationAttemptId" json:"application_attempt_id,omitempty"`
	Host                        *string                           `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	RpcPort                     *int32                            `protobuf:"varint,3,opt,name=rpc_port,json=rpcPort" json:"rpc_port,omitempty"`
	TrackingUrl                 *string                           `protobuf:"bytes,4,opt,name=tracking_url,json=trackingUrl" json:"tracking_url,omitempty"`
	DiagnosticsInfo             *string                           `protobuf:"bytes,5,opt,name=diagnostics_info,json=diagnosticsInfo" json:"diagnostics_info,omitempty"`
	FinalApplicationStatus      *FinalApplicationStatusProto      `protobuf:"varint,6,opt,name=final_application_status,json=finalApplicationStatus,enum=hadoop.yarn.FinalApplicationStatusProto" json:"final_application_status,omitempty"`
	MasterContainerId           *ContainerIdProto                 `protobuf:"bytes,7,opt,name=master_container_id,json=masterContainerId" json:"master_container_id,omitempty"`
	YarnApplicationAttemptState *YarnApplicationAttemptStateProto `protobuf:"varint,8,opt,name=yarn_application_attempt_state,json=yarnApplicationAttemptState,enum=hadoop.yarn.YarnApplicationAttemptStateProto" json:"yarn_application_attempt_state,omitempty"`
	XXX_unrecognized            []byte                            `json:"-"`
}

func (m *ApplicationAttemptHistoryDataProto) Reset()         { *m = ApplicationAttemptHistoryDataProto{} }
func (m *ApplicationAttemptHistoryDataProto) String() string { return proto.CompactTextString(m) }
func (*ApplicationAttemptHistoryDataProto) ProtoMessage()    {}
func (*ApplicationAttemptHistoryDataProto) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3}
}

func (m *ApplicationAttemptHistoryDataProto) GetApplicationAttemptId() *ApplicationAttemptIdProto {
	if m != nil {
		return m.ApplicationAttemptId
	}
	return nil
}

func (m *ApplicationAttemptHistoryDataProto) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

func (m *ApplicationAttemptHistoryDataProto) GetRpcPort() int32 {
	if m != nil && m.RpcPort != nil {
		return *m.RpcPort
	}
	return 0
}

func (m *ApplicationAttemptHistoryDataProto) GetTrackingUrl() string {
	if m != nil && m.TrackingUrl != nil {
		return *m.TrackingUrl
	}
	return ""
}

func (m *ApplicationAttemptHistoryDataProto) GetDiagnosticsInfo() string {
	if m != nil && m.DiagnosticsInfo != nil {
		return *m.DiagnosticsInfo
	}
	return ""
}

func (m *ApplicationAttemptHistoryDataProto) GetFinalApplicationStatus() FinalApplicationStatusProto {
	if m != nil && m.FinalApplicationStatus != nil {
		return *m.FinalApplicationStatus
	}
	return FinalApplicationStatusProto_APP_UNDEFINED
}

func (m *ApplicationAttemptHistoryDataProto) GetMasterContainerId() *ContainerIdProto {
	if m != nil {
		return m.MasterContainerId
	}
	return nil
}

func (m *ApplicationAttemptHistoryDataProto) GetYarnApplicationAttemptState() YarnApplicationAttemptStateProto {
	if m != nil && m.YarnApplicationAttemptState != nil {
		return *m.YarnApplicationAttemptState
	}
	return YarnApplicationAttemptStateProto_APP_ATTEMPT_NEW
}

type ApplicationAttemptStartDataProto struct {
	ApplicationAttemptId *ApplicationAttemptIdProto `protobuf:"bytes,1,opt,name=application_attempt_id,json=applicationAttemptId" json:"application_attempt_id,omitempty"`
	Host                 *string                    `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	RpcPort              *int32                     `protobuf:"varint,3,opt,name=rpc_port,json=rpcPort" json:"rpc_port,omitempty"`
	MasterContainerId    *ContainerIdProto          `protobuf:"bytes,4,opt,name=master_container_id,json=masterContainerId" json:"master_container_id,omitempty"`
	XXX_unrecognized     []byte                     `json:"-"`
}

func (m *ApplicationAttemptStartDataProto) Reset()         { *m = ApplicationAttemptStartDataProto{} }
func (m *ApplicationAttemptStartDataProto) String() string { return proto.CompactTextString(m) }
func (*ApplicationAttemptStartDataProto) ProtoMessage()    {}
func (*ApplicationAttemptStartDataProto) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4}
}

func (m *ApplicationAttemptStartDataProto) GetApplicationAttemptId() *ApplicationAttemptIdProto {
	if m != nil {
		return m.ApplicationAttemptId
	}
	return nil
}

func (m *ApplicationAttemptStartDataProto) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

func (m *ApplicationAttemptStartDataProto) GetRpcPort() int32 {
	if m != nil && m.RpcPort != nil {
		return *m.RpcPort
	}
	return 0
}

func (m *ApplicationAttemptStartDataProto) GetMasterContainerId() *ContainerIdProto {
	if m != nil {
		return m.MasterContainerId
	}
	return nil
}

type ApplicationAttemptFinishDataProto struct {
	ApplicationAttemptId        *ApplicationAttemptIdProto        `protobuf:"bytes,1,opt,name=application_attempt_id,json=applicationAttemptId" json:"application_attempt_id,omitempty"`
	TrackingUrl                 *string                           `protobuf:"bytes,2,opt,name=tracking_url,json=trackingUrl" json:"tracking_url,omitempty"`
	DiagnosticsInfo             *string                           `protobuf:"bytes,3,opt,name=diagnostics_info,json=diagnosticsInfo" json:"diagnostics_info,omitempty"`
	FinalApplicationStatus      *FinalApplicationStatusProto      `protobuf:"varint,4,opt,name=final_application_status,json=finalApplicationStatus,enum=hadoop.yarn.FinalApplicationStatusProto" json:"final_application_status,omitempty"`
	YarnApplicationAttemptState *YarnApplicationAttemptStateProto `protobuf:"varint,5,opt,name=yarn_application_attempt_state,json=yarnApplicationAttemptState,enum=hadoop.yarn.YarnApplicationAttemptStateProto" json:"yarn_application_attempt_state,omitempty"`
	XXX_unrecognized            []byte                            `json:"-"`
}

func (m *ApplicationAttemptFinishDataProto) Reset()         { *m = ApplicationAttemptFinishDataProto{} }
func (m *ApplicationAttemptFinishDataProto) String() string { return proto.CompactTextString(m) }
func (*ApplicationAttemptFinishDataProto) ProtoMessage()    {}
func (*ApplicationAttemptFinishDataProto) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5}
}

func (m *ApplicationAttemptFinishDataProto) GetApplicationAttemptId() *ApplicationAttemptIdProto {
	if m != nil {
		return m.ApplicationAttemptId
	}
	return nil
}

func (m *ApplicationAttemptFinishDataProto) GetTrackingUrl() string {
	if m != nil && m.TrackingUrl != nil {
		return *m.TrackingUrl
	}
	return ""
}

func (m *ApplicationAttemptFinishDataProto) GetDiagnosticsInfo() string {
	if m != nil && m.DiagnosticsInfo != nil {
		return *m.DiagnosticsInfo
	}
	return ""
}

func (m *ApplicationAttemptFinishDataProto) GetFinalApplicationStatus() FinalApplicationStatusProto {
	if m != nil && m.FinalApplicationStatus != nil {
		return *m.FinalApplicationStatus
	}
	return FinalApplicationStatusProto_APP_UNDEFINED
}

func (m *ApplicationAttemptFinishDataProto) GetYarnApplicationAttemptState() YarnApplicationAttemptStateProto {
	if m != nil && m.YarnApplicationAttemptState != nil {
		return *m.YarnApplicationAttemptState
	}
	return YarnApplicationAttemptStateProto_APP_ATTEMPT_NEW
}

type ContainerHistoryDataProto struct {
	ContainerId         *ContainerIdProto    `protobuf:"bytes,1,opt,name=container_id,json=containerId" json:"container_id,omitempty"`
	AllocatedResource   *ResourceProto       `protobuf:"bytes,2,opt,name=allocated_resource,json=allocatedResource" json:"allocated_resource,omitempty"`
	AssignedNodeId      *NodeIdProto         `protobuf:"bytes,3,opt,name=assigned_node_id,json=assignedNodeId" json:"assigned_node_id,omitempty"`
	Priority            *PriorityProto       `protobuf:"bytes,4,opt,name=priority" json:"priority,omitempty"`
	StartTime           *int64               `protobuf:"varint,5,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	FinishTime          *int64               `protobuf:"varint,6,opt,name=finish_time,json=finishTime" json:"finish_time,omitempty"`
	DiagnosticsInfo     *string              `protobuf:"bytes,7,opt,name=diagnostics_info,json=diagnosticsInfo" json:"diagnostics_info,omitempty"`
	ContainerExitStatus *int32               `protobuf:"varint,8,opt,name=container_exit_status,json=containerExitStatus" json:"container_exit_status,omitempty"`
	ContainerState      *ContainerStateProto `protobuf:"varint,9,opt,name=container_state,json=containerState,enum=hadoop.yarn.ContainerStateProto" json:"container_state,omitempty"`
	XXX_unrecognized    []byte               `json:"-"`
}

func (m *ContainerHistoryDataProto) Reset()                    { *m = ContainerHistoryDataProto{} }
func (m *ContainerHistoryDataProto) String() string            { return proto.CompactTextString(m) }
func (*ContainerHistoryDataProto) ProtoMessage()               {}
func (*ContainerHistoryDataProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ContainerHistoryDataProto) GetContainerId() *ContainerIdProto {
	if m != nil {
		return m.ContainerId
	}
	return nil
}

func (m *ContainerHistoryDataProto) GetAllocatedResource() *ResourceProto {
	if m != nil {
		return m.AllocatedResource
	}
	return nil
}

func (m *ContainerHistoryDataProto) GetAssignedNodeId() *NodeIdProto {
	if m != nil {
		return m.AssignedNodeId
	}
	return nil
}

func (m *ContainerHistoryDataProto) GetPriority() *PriorityProto {
	if m != nil {
		return m.Priority
	}
	return nil
}

func (m *ContainerHistoryDataProto) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *ContainerHistoryDataProto) GetFinishTime() int64 {
	if m != nil && m.FinishTime != nil {
		return *m.FinishTime
	}
	return 0
}

func (m *ContainerHistoryDataProto) GetDiagnosticsInfo() string {
	if m != nil && m.DiagnosticsInfo != nil {
		return *m.DiagnosticsInfo
	}
	return ""
}

func (m *ContainerHistoryDataProto) GetContainerExitStatus() int32 {
	if m != nil && m.ContainerExitStatus != nil {
		return *m.ContainerExitStatus
	}
	return 0
}

func (m *ContainerHistoryDataProto) GetContainerState() ContainerStateProto {
	if m != nil && m.ContainerState != nil {
		return *m.ContainerState
	}
	return ContainerStateProto_C_NEW
}

type ContainerStartDataProto struct {
	ContainerId       *ContainerIdProto `protobuf:"bytes,1,opt,name=container_id,json=containerId" json:"container_id,omitempty"`
	AllocatedResource *ResourceProto    `protobuf:"bytes,2,opt,name=allocated_resource,json=allocatedResource" json:"allocated_resource,omitempty"`
	AssignedNodeId    *NodeIdProto      `protobuf:"bytes,3,opt,name=assigned_node_id,json=assignedNodeId" json:"assigned_node_id,omitempty"`
	Priority          *PriorityProto    `protobuf:"bytes,4,opt,name=priority" json:"priority,omitempty"`
	StartTime         *int64            `protobuf:"varint,5,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	XXX_unrecognized  []byte            `json:"-"`
}

func (m *ContainerStartDataProto) Reset()                    { *m = ContainerStartDataProto{} }
func (m *ContainerStartDataProto) String() string            { return proto.CompactTextString(m) }
func (*ContainerStartDataProto) ProtoMessage()               {}
func (*ContainerStartDataProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ContainerStartDataProto) GetContainerId() *ContainerIdProto {
	if m != nil {
		return m.ContainerId
	}
	return nil
}

func (m *ContainerStartDataProto) GetAllocatedResource() *ResourceProto {
	if m != nil {
		return m.AllocatedResource
	}
	return nil
}

func (m *ContainerStartDataProto) GetAssignedNodeId() *NodeIdProto {
	if m != nil {
		return m.AssignedNodeId
	}
	return nil
}

func (m *ContainerStartDataProto) GetPriority() *PriorityProto {
	if m != nil {
		return m.Priority
	}
	return nil
}

func (m *ContainerStartDataProto) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

type ContainerFinishDataProto struct {
	ContainerId         *ContainerIdProto    `protobuf:"bytes,1,opt,name=container_id,json=containerId" json:"container_id,omitempty"`
	FinishTime          *int64               `protobuf:"varint,2,opt,name=finish_time,json=finishTime" json:"finish_time,omitempty"`
	DiagnosticsInfo     *string              `protobuf:"bytes,3,opt,name=diagnostics_info,json=diagnosticsInfo" json:"diagnostics_info,omitempty"`
	ContainerExitStatus *int32               `protobuf:"varint,4,opt,name=container_exit_status,json=containerExitStatus" json:"container_exit_status,omitempty"`
	ContainerState      *ContainerStateProto `protobuf:"varint,5,opt,name=container_state,json=containerState,enum=hadoop.yarn.ContainerStateProto" json:"container_state,omitempty"`
	XXX_unrecognized    []byte               `json:"-"`
}

func (m *ContainerFinishDataProto) Reset()                    { *m = ContainerFinishDataProto{} }
func (m *ContainerFinishDataProto) String() string            { return proto.CompactTextString(m) }
func (*ContainerFinishDataProto) ProtoMessage()               {}
func (*ContainerFinishDataProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ContainerFinishDataProto) GetContainerId() *ContainerIdProto {
	if m != nil {
		return m.ContainerId
	}
	return nil
}

func (m *ContainerFinishDataProto) GetFinishTime() int64 {
	if m != nil && m.FinishTime != nil {
		return *m.FinishTime
	}
	return 0
}

func (m *ContainerFinishDataProto) GetDiagnosticsInfo() string {
	if m != nil && m.DiagnosticsInfo != nil {
		return *m.DiagnosticsInfo
	}
	return ""
}

func (m *ContainerFinishDataProto) GetContainerExitStatus() int32 {
	if m != nil && m.ContainerExitStatus != nil {
		return *m.ContainerExitStatus
	}
	return 0
}

func (m *ContainerFinishDataProto) GetContainerState() ContainerStateProto {
	if m != nil && m.ContainerState != nil {
		return *m.ContainerState
	}
	return ContainerStateProto_C_NEW
}

func init() {
	proto.RegisterType((*ApplicationHistoryDataProto)(nil), "hadoop.yarn.ApplicationHistoryDataProto")
	proto.RegisterType((*ApplicationStartDataProto)(nil), "hadoop.yarn.ApplicationStartDataProto")
	proto.RegisterType((*ApplicationFinishDataProto)(nil), "hadoop.yarn.ApplicationFinishDataProto")
	proto.RegisterType((*ApplicationAttemptHistoryDataProto)(nil), "hadoop.yarn.ApplicationAttemptHistoryDataProto")
	proto.RegisterType((*ApplicationAttemptStartDataProto)(nil), "hadoop.yarn.ApplicationAttemptStartDataProto")
	proto.RegisterType((*ApplicationAttemptFinishDataProto)(nil), "hadoop.yarn.ApplicationAttemptFinishDataProto")
	proto.RegisterType((*ContainerHistoryDataProto)(nil), "hadoop.yarn.ContainerHistoryDataProto")
	proto.RegisterType((*ContainerStartDataProto)(nil), "hadoop.yarn.ContainerStartDataProto")
	proto.RegisterType((*ContainerFinishDataProto)(nil), "hadoop.yarn.ContainerFinishDataProto")
}

func init() { proto.RegisterFile("server/application_history_server.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 827 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x97, 0x41, 0x4e, 0xc3, 0x46,
	0x14, 0x86, 0xe5, 0xc4, 0x26, 0xc9, 0x33, 0x0d, 0x60, 0x28, 0x35, 0x50, 0x20, 0x78, 0x41, 0xc3,
	0xa2, 0xa9, 0x94, 0x45, 0xd7, 0x85, 0x52, 0xd4, 0x2c, 0x8a, 0x50, 0xa0, 0x8b, 0x4a, 0x48, 0xd6,
	0x60, 0x4f, 0x92, 0x51, 0x13, 0x8f, 0x3b, 0x33, 0xa9, 0xc8, 0x0d, 0xba, 0xea, 0x19, 0x2a, 0x55,
	0x2a, 0x07, 0xe8, 0x09, 0x7a, 0xa5, 0x1e, 0xa0, 0xaa, 0x32, 0x63, 0x27, 0x63, 0x27, 0x29, 0x81,
	0x06, 0xb5, 0x0b, 0x36, 0x28, 0xbc, 0xff, 0xbd, 0x37, 0xe3, 0xff, 0x7d, 0x1e, 0xdb, 0xf0, 0x09,
	0xc7, 0xec, 0x47, 0xcc, 0x3e, 0x43, 0x71, 0xdc, 0x27, 0x01, 0x12, 0x84, 0x46, 0x7e, 0x8f, 0x70,
	0x41, 0xd9, 0xc8, 0x57, 0x52, 0x23, 0x66, 0x54, 0x50, 0xc7, 0xee, 0xa1, 0x90, 0xd2, 0xb8, 0x31,
	0x42, 0x2c, 0xda, 0xdf, 0x1a, 0xff, 0xf5, 0xa5, 0xc0, 0x95, 0xee, 0xfd, 0x6e, 0xc2, 0xc1, 0xf9,
	0xb4, 0xc9, 0xd7, 0xaa, 0xc7, 0x25, 0x12, 0xe8, 0x46, 0xd6, 0x5f, 0x41, 0x55, 0x5f, 0x83, 0x84,
	0xae, 0x51, 0x33, 0xea, 0x76, 0xf3, 0xb8, 0xa1, 0x35, 0x6e, 0x68, 0x1d, 0x5a, 0xa1, 0x2c, 0x6c,
	0x7f, 0x80, 0xf4, 0x98, 0x73, 0x06, 0x9b, 0x7a, 0x9f, 0x08, 0x0d, 0xb0, 0x5b, 0xa8, 0x19, 0xf5,
	0x4a, 0x7b, 0x43, 0x8b, 0x5f, 0xa3, 0x01, 0xce, 0xa7, 0x8a, 0x51, 0x8c, 0xdd, 0xe2, 0x4c, 0xea,
	0xdd, 0x28, 0xc6, 0x8e, 0x03, 0xe6, 0x90, 0x63, 0xe6, 0x9a, 0x52, 0x96, 0xbf, 0x9d, 0x1d, 0xb0,
	0x7e, 0x18, 0xe2, 0x21, 0x76, 0x2d, 0x19, 0x54, 0xff, 0x38, 0xc7, 0x60, 0xf3, 0xe1, 0xc3, 0x80,
	0x08, 0x5f, 0x90, 0x01, 0x76, 0xd7, 0x6a, 0x46, 0xbd, 0xd8, 0x06, 0x15, 0xba, 0x23, 0x03, 0xec,
	0x1c, 0x02, 0x70, 0x81, 0x58, 0xa2, 0x97, 0xa4, 0x5e, 0x91, 0x11, 0x29, 0x1f, 0x83, 0xdd, 0x21,
	0x11, 0xe1, 0x3d, 0xa5, 0x97, 0x55, 0xbd, 0x0a, 0xc9, 0x84, 0x33, 0xd8, 0x0c, 0x09, 0xea, 0x46,
	0x94, 0x0b, 0x12, 0x70, 0x9f, 0x44, 0x1d, 0xea, 0x56, 0xd4, 0xae, 0xb5, 0x78, 0x2b, 0xea, 0x50,
	0xe7, 0x01, 0xdc, 0x0e, 0x89, 0x50, 0xdf, 0xd7, 0x2f, 0x93, 0x0b, 0x24, 0x86, 0xdc, 0x85, 0x9a,
	0x51, 0xaf, 0x36, 0xeb, 0x19, 0x77, 0xaf, 0xc6, 0xc9, 0x9a, 0xc5, 0xb7, 0x32, 0x55, 0xd9, 0xbc,
	0xdb, 0x99, 0x2b, 0x3a, 0xf7, 0xb0, 0x2b, 0x87, 0x9d, 0x5f, 0x02, 0xbb, 0xb6, 0x5c, 0xe1, 0x34,
	0xb3, 0xc2, 0x77, 0x88, 0x45, 0xb9, 0x1e, 0x58, 0xf5, 0xdf, 0x19, 0xcd, 0x91, 0xbc, 0x5f, 0x0b,
	0xb0, 0x97, 0x0d, 0x32, 0xf1, 0xce, 0x4c, 0x8e, 0x19, 0xef, 0xcf, 0x02, 0xec, 0x6b, 0x57, 0x79,
	0x25, 0x61, 0x59, 0xbd, 0x4d, 0x39, 0x34, 0x0b, 0x4b, 0xa1, 0x59, 0x7c, 0x39, 0x9a, 0xe6, 0x9b,
	0xa3, 0x69, 0xad, 0x00, 0xcd, 0xdf, 0x4c, 0xf0, 0xb4, 0xe0, 0xb9, 0x10, 0x78, 0x10, 0x8b, 0x99,
	0x73, 0xed, 0x1e, 0x76, 0xf5, 0xf5, 0x91, 0x4a, 0x9b, 0x0e, 0xe1, 0x74, 0xd1, 0x10, 0x92, 0x86,
	0xe9, 0x2c, 0x76, 0xd0, 0x1c, 0x69, 0xcc, 0x58, 0x8f, 0x72, 0x91, 0xd0, 0x2a, 0x7f, 0x3b, 0x7b,
	0x50, 0x66, 0x71, 0xe0, 0xc7, 0x94, 0x09, 0xe9, 0xbe, 0xd5, 0x2e, 0xb1, 0x38, 0xb8, 0xa1, 0x4c,
	0x38, 0x27, 0xb0, 0x2e, 0x18, 0x0a, 0xbe, 0x27, 0x51, 0xd7, 0x1f, 0xb2, 0x7e, 0x82, 0xa6, 0x9d,
	0xc6, 0xbe, 0x65, 0xfd, 0xb9, 0x33, 0xb4, 0x5e, 0x3e, 0xc3, 0xb5, 0x15, 0xcd, 0xf0, 0x1b, 0xd8,
	0x1e, 0x20, 0x2e, 0x30, 0xf3, 0x03, 0x1a, 0x09, 0x44, 0x22, 0xcc, 0xc6, 0xde, 0x95, 0xa4, 0x77,
	0x87, 0x99, 0xf6, 0x5f, 0xa6, 0x09, 0xa9, 0x65, 0x5b, 0xaa, 0x52, 0x8b, 0x3b, 0x0c, 0x8e, 0x66,
	0x90, 0x48, 0x47, 0xa2, 0xd0, 0x28, 0xcb, 0x8d, 0x7f, 0xfa, 0x4f, 0x68, 0x24, 0xf6, 0x6b, 0x84,
	0x1c, 0x8c, 0x16, 0x67, 0x78, 0x7f, 0x19, 0x50, 0x9b, 0xab, 0xe9, 0x47, 0xd9, 0xff, 0x0a, 0x93,
	0x05, 0xa6, 0x9b, 0xaf, 0x33, 0xdd, 0x7b, 0x2a, 0xc2, 0xc9, 0xec, 0x8e, 0xf3, 0xa7, 0xd4, 0xdb,
	0x3a, 0x90, 0x27, 0xbf, 0xb0, 0x1c, 0xf9, 0xff, 0xe1, 0xe9, 0xf5, 0x3c, 0xaa, 0xd6, 0xca, 0x51,
	0xfd, 0xd9, 0x84, 0xbd, 0xc9, 0xe4, 0x66, 0x8e, 0xb2, 0x2f, 0x60, 0x3d, 0xc3, 0x83, 0xb1, 0x0c,
	0x0f, 0x76, 0xa0, 0xdd, 0x7e, 0x2d, 0x70, 0x50, 0xbf, 0x4f, 0x03, 0x24, 0x70, 0xe8, 0x33, 0xcc,
	0xe9, 0x90, 0x05, 0xea, 0x41, 0x62, 0x37, 0xf7, 0x33, 0x7d, 0xda, 0x89, 0x98, 0x40, 0x35, 0xa9,
	0x4a, 0xe3, 0xce, 0x05, 0x6c, 0x22, 0xce, 0x49, 0x37, 0xc2, 0xa1, 0x1f, 0xd1, 0x10, 0x8f, 0x37,
	0x54, 0x94, 0x8d, 0xdc, 0x4c, 0xa3, 0x6b, 0x1a, 0xe2, 0x74, 0x2f, 0xd5, 0xb4, 0x42, 0x05, 0x9d,
	0xcf, 0xa1, 0x1c, 0x33, 0x42, 0x19, 0x11, 0xa3, 0x04, 0xee, 0xec, 0x26, 0x6e, 0x12, 0x51, 0x55,
	0x4f, 0x72, 0x73, 0x8f, 0x63, 0xeb, 0x99, 0x57, 0xb8, 0xb5, 0xa5, 0x9e, 0x93, 0xa5, 0xf9, 0xa4,
	0x35, 0xe1, 0xc3, 0xa9, 0xe7, 0xf8, 0x91, 0x88, 0x14, 0xb3, 0xb2, 0xbc, 0x65, 0xb7, 0x27, 0xe2,
	0x57, 0x8f, 0x44, 0x24, 0xe4, 0xb4, 0x60, 0x63, 0x5a, 0xa3, 0x50, 0xa9, 0x48, 0x54, 0x6a, 0xf3,
	0x47, 0xa5, 0xd1, 0x51, 0x0d, 0x32, 0x41, 0xef, 0x8f, 0x02, 0x7c, 0xa4, 0xe7, 0xe9, 0x47, 0xd6,
	0x3b, 0x0e, 0x4b, 0xe1, 0xe0, 0x3d, 0x15, 0xc0, 0x9d, 0xf8, 0x90, 0x3f, 0xf5, 0xfe, 0xbd, 0x89,
	0xab, 0x7c, 0x2b, 0x5b, 0x48, 0x9b, 0xf9, 0x22, 0xda, 0xac, 0xd7, 0xd1, 0x76, 0x71, 0x09, 0x1f,
	0x53, 0xd6, 0x6d, 0xa0, 0x18, 0x05, 0x3d, 0x9c, 0xa9, 0x96, 0xdf, 0x90, 0x17, 0x47, 0xb3, 0x1f,
	0x90, 0xb7, 0xf2, 0x1b, 0x54, 0xb6, 0xe3, 0x3f, 0x19, 0xc6, 0x2f, 0x86, 0xf1, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x58, 0x80, 0x01, 0xd0, 0xb3, 0x0e, 0x00, 0x00,
}
