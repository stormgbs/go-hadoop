package yarn_client

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/stormgbs/go-hadoop/hadoop_yarn"
	yarn_conf "github.com/stormgbs/go-hadoop/hadoop_yarn/conf"
)

type RMClient struct {
	client hadoop_yarn.ResourceManagerAdministrationProtocolService
	conf   yarn_conf.YarnConfiguration
}

func CreateRMClient(conf yarn_conf.YarnConfiguration) (*RMClient, error) {
	c, err := hadoop_yarn.DialResourceManagerAdministrationProtocolService(conf)
	return &RMClient{client: c, conf: conf}, err
}

func (c *RMClient) UpdateNodeResource(nodeHost string, nodePort int32, memoryMB int32, virtualCpuCores int32) error {
	request := hadoop_yarn.UpdateNodeResourceRequestProto{
		NodeResourceMap: []*hadoop_yarn.NodeResourceMapProto{
			&hadoop_yarn.NodeResourceMapProto{
				NodeId: &hadoop_yarn.NodeIdProto{Host: &nodeHost, Port: &nodePort},
				ResourceOption: &hadoop_yarn.ResourceOptionProto{
					Resource: &hadoop_yarn.ResourceProto{
						Memory:       &memoryMB,
						VirtualCores: &virtualCpuCores,
					},
				},
			},
		},
	}
	response := hadoop_yarn.UpdateNodeResourceResponseProto{}
	return c.client.UpdateNodeResource(&request, &response)
}

func (c *RMClient) UpdateResourceOfGivenNode(memoryMB int32, virtualCpuCores int32) error {
	nodeId, _ := c.conf.Get(yarn_conf.NM_ADDRESS, yarn_conf.DEFAULT_NM_ADDRESS)
	hostAndPort := strings.Split(nodeId, ":")
	if len(hostAndPort) != 2 {
		return errors.New("illegal nodeId format[host:port].nodeId=" + nodeId)
	}
	port, err := strconv.Atoi(hostAndPort[1])
	if err != nil {
		return errors.New(fmt.Sprintf("%s%d", "number format error.port=", port))
	}
	return c.UpdateNodeResource(hostAndPort[0], int32(port), memoryMB, virtualCpuCores)
}
