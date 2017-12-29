package yarn_client

import (
	"github.com/stormgbs/go-hadoop/hadoop_yarn"
	yarn_conf "github.com/stormgbs/go-hadoop/hadoop_yarn/conf"
)

type NMClient struct {
	client hadoop_yarn.NodeManagerResourceUtilizationTrackerService
	conf   yarn_conf.YarnConfiguration
}

func CreateNMClient(conf yarn_conf.YarnConfiguration) (*NMClient, error) {
	c, err := hadoop_yarn.DialNodeManagerResourceUtilizationTrackerService(conf)
	return &NMClient{client: c, conf: conf}, err
}

func (client *NMClient) GetResourceUtilization() (*hadoop_yarn.ResourceUtilizationResponseProto, error) {
	req := &hadoop_yarn.ResourceUtilizationRequestProto{}
	resp := &hadoop_yarn.ResourceUtilizationResponseProto{}
	return resp, client.client.GetResourceUtilization(req, resp)
}
