package hadoop_yarn

import proto "github.com/golang/protobuf/proto"
import json "encoding/json"
import math "math"
import "github.com/stormgbs/go-hadoop"

import hadoop_ipc_client "github.com/stormgbs/go-hadoop/hadoop_common/ipc/client"
import yarn_conf "github.com/stormgbs/go-hadoop/hadoop_yarn/conf"
import "github.com/nu7hatch/gouuid"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf
var RESOURCE_MANAGER_ADMINISTRATION_PROTOCOL = "org.apache.hadoop.yarn.server.api.ResourceManagerAdministrationProtocolPB"

func init() {
}

type ResourceManagerAdministrationProtocolService interface {
	RefreshQueues(in *RefreshQueuesRequestProto, out *RefreshQueuesResponseProto) error
	RefreshNodes(in *RefreshNodesRequestProto, out *RefreshNodesResponseProto) error
	RefreshSuperUserGroupsConfiguration(in *RefreshSuperUserGroupsConfigurationRequestProto, out *RefreshSuperUserGroupsConfigurationResponseProto) error
	RefreshUserToGroupsMappings(in *RefreshUserToGroupsMappingsRequestProto, out *RefreshUserToGroupsMappingsResponseProto) error
	RefreshAdminAcls(in *RefreshAdminAclsRequestProto, out *RefreshAdminAclsResponseProto) error
	RefreshServiceAcls(in *RefreshServiceAclsRequestProto, out *RefreshServiceAclsResponseProto) error
	GetGroupsForUser(in *GetGroupsForUserRequestProto, out *GetGroupsForUserResponseProto) error
	UpdateNodeResource(in *UpdateNodeResourceRequestProto, out *UpdateNodeResourceResponseProto) error
	AddToClusterNodeLabels(in *AddToClusterNodeLabelsRequestProto, out *AddToClusterNodeLabelsResponseProto) error
	RemoveFromClusterNodeLabels(in *RemoveFromClusterNodeLabelsRequestProto, out *RemoveFromClusterNodeLabelsResponseProto) error
	ReplaceLabelsOnNodes(in *ReplaceLabelsOnNodeRequestProto, out *ReplaceLabelsOnNodeResponseProto) error
	CheckForDecommissioningNodes(in *CheckForDecommissioningNodesRequestProto, out *CheckForDecommissioningNodesResponseProto) error
}

type ResourceManagerAdministrationProtocolServiceClient struct {
	*hadoop_ipc_client.Client
}

func (c *ResourceManagerAdministrationProtocolServiceClient) RefreshQueues(in *RefreshQueuesRequestProto, out *RefreshQueuesResponseProto) error {
	return c.Call(gohadoop.GetCalleeRPCRequestHeaderProto(&RESOURCE_MANAGER_ADMINISTRATION_PROTOCOL), in, out)
}

func (c *ResourceManagerAdministrationProtocolServiceClient) RefreshNodes(in *RefreshNodesRequestProto, out *RefreshNodesResponseProto) error {
	return c.Call(gohadoop.GetCalleeRPCRequestHeaderProto(&RESOURCE_MANAGER_ADMINISTRATION_PROTOCOL), in, out)
}

func (c *ResourceManagerAdministrationProtocolServiceClient) RefreshSuperUserGroupsConfiguration(in *RefreshSuperUserGroupsConfigurationRequestProto, out *RefreshSuperUserGroupsConfigurationResponseProto) error {
	return c.Call(gohadoop.GetCalleeRPCRequestHeaderProto(&RESOURCE_MANAGER_ADMINISTRATION_PROTOCOL), in, out)
}

func (c *ResourceManagerAdministrationProtocolServiceClient) RefreshUserToGroupsMappings(in *RefreshUserToGroupsMappingsRequestProto, out *RefreshUserToGroupsMappingsResponseProto) error {
	return c.Call(gohadoop.GetCalleeRPCRequestHeaderProto(&RESOURCE_MANAGER_ADMINISTRATION_PROTOCOL), in, out)
}

func (c *ResourceManagerAdministrationProtocolServiceClient) RefreshAdminAcls(in *RefreshAdminAclsRequestProto, out *RefreshAdminAclsResponseProto) error {
	return c.Call(gohadoop.GetCalleeRPCRequestHeaderProto(&RESOURCE_MANAGER_ADMINISTRATION_PROTOCOL), in, out)
}

func (c *ResourceManagerAdministrationProtocolServiceClient) RefreshServiceAcls(in *RefreshServiceAclsRequestProto, out *RefreshServiceAclsResponseProto) error {
	return c.Call(gohadoop.GetCalleeRPCRequestHeaderProto(&RESOURCE_MANAGER_ADMINISTRATION_PROTOCOL), in, out)
}

func (c *ResourceManagerAdministrationProtocolServiceClient) GetGroupsForUser(in *GetGroupsForUserRequestProto, out *GetGroupsForUserResponseProto) error {
	return c.Call(gohadoop.GetCalleeRPCRequestHeaderProto(&RESOURCE_MANAGER_ADMINISTRATION_PROTOCOL), in, out)
}

func (c *ResourceManagerAdministrationProtocolServiceClient) UpdateNodeResource(in *UpdateNodeResourceRequestProto, out *UpdateNodeResourceResponseProto) error {
	return c.Call(gohadoop.GetCalleeRPCRequestHeaderProto(&RESOURCE_MANAGER_ADMINISTRATION_PROTOCOL), in, out)
}

func (c *ResourceManagerAdministrationProtocolServiceClient) AddToClusterNodeLabels(in *AddToClusterNodeLabelsRequestProto, out *AddToClusterNodeLabelsResponseProto) error {
	return c.Call(gohadoop.GetCalleeRPCRequestHeaderProto(&RESOURCE_MANAGER_ADMINISTRATION_PROTOCOL), in, out)
}

func (c *ResourceManagerAdministrationProtocolServiceClient) RemoveFromClusterNodeLabels(in *RemoveFromClusterNodeLabelsRequestProto, out *RemoveFromClusterNodeLabelsResponseProto) error {
	return c.Call(gohadoop.GetCalleeRPCRequestHeaderProto(&RESOURCE_MANAGER_ADMINISTRATION_PROTOCOL), in, out)
}

func (c *ResourceManagerAdministrationProtocolServiceClient) ReplaceLabelsOnNodes(in *ReplaceLabelsOnNodeRequestProto, out *ReplaceLabelsOnNodeResponseProto) error {
	return c.Call(gohadoop.GetCalleeRPCRequestHeaderProto(&RESOURCE_MANAGER_ADMINISTRATION_PROTOCOL), in, out)
}

func (c *ResourceManagerAdministrationProtocolServiceClient) CheckForDecommissioningNodes(in *CheckForDecommissioningNodesRequestProto, out *CheckForDecommissioningNodesResponseProto) error {
	return c.Call(gohadoop.GetCalleeRPCRequestHeaderProto(&RESOURCE_MANAGER_ADMINISTRATION_PROTOCOL), in, out)
}

// DialResourceManagerAdministrationProtocolService connects to an ResourceManagerAdministrationProtocolService at the specified network address.
func DialResourceManagerAdministrationProtocolService(conf yarn_conf.YarnConfiguration) (ResourceManagerAdministrationProtocolService, error) {
	clientId, _ := uuid.NewV4()
	ugi, _ := gohadoop.CreateSimpleUGIProto()
	adminServerAddress, _ := conf.GetRMAdminAddress()
	c := &hadoop_ipc_client.Client{ClientId: clientId, Ugi: ugi, ServerAddress: adminServerAddress}
	return &ResourceManagerAdministrationProtocolServiceClient{c}, nil
}
