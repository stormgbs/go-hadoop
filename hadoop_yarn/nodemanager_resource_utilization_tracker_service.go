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
var NODE_MANAGER_RESOURCE_UTILIZATION_TRACKER = "org.apache.hadoop.yarn.server.api.ResourceUtilizationTrackerPB"

func init() {
}

type NodeManagerResourceUtilizationTrackerService interface {
	GetResourceUtilization(req *ResourceUtilizationRequestProto, resp *ResourceUtilizationResponseProto) error
}

type NodeManagerResourceUtilizationTrackerServiceClient struct {
	*hadoop_ipc_client.Client
}

func (c NodeManagerResourceUtilizationTrackerServiceClient) GetResourceUtilization(req *ResourceUtilizationRequestProto,
	resp *ResourceUtilizationResponseProto) error {
	return c.Call(gohadoop.GetCalleeRPCRequestHeaderProto(&NODE_MANAGER_RESOURCE_UTILIZATION_TRACKER), req, resp)
}

func DialNodeManagerResourceUtilizationTrackerService(conf yarn_conf.YarnConfiguration) (NodeManagerResourceUtilizationTrackerService, error) {
	clientId, _ := uuid.NewV4()
	ugi, _ := gohadoop.CreateSimpleUGIProto()
	serviceAddress, _ := conf.GetNMResourceUtilizationTrackerAddress()
	c := &hadoop_ipc_client.Client{ClientId: clientId, Ugi: ugi, ServerAddress: serviceAddress}
	return &NodeManagerResourceUtilizationTrackerServiceClient{c}, nil
}
