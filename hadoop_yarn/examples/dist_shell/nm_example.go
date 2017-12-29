package main

import (
	"log"

	"github.com/stormgbs/go-hadoop/hadoop_yarn"
	"github.com/stormgbs/go-hadoop/hadoop_yarn/conf"
	"github.com/stormgbs/go-hadoop/hadoop_yarn/yarn_client"
)

func main() {
	var err error
	var resp *hadoop_yarn.ResourceUtilizationResponseProto
	// Create YarnConfiguration
	conf, _ := conf.NewYarnConfiguration()

	// Create RMClient
	nmClient, _ := yarn_client.CreateNMClient(conf)

	resp, err = nmClient.GetResourceUtilization()

	if err != nil {
		log.Fatal("rmClient.UpdateNodeResource: ", err)
	} else {
		nodeStatus := resp.NodeStatus
		containersStatus := nodeStatus.ContainersStatuses
		size := len(containersStatus)
		log.Println("container status size:", size)
		if size == 0 {
			log.Println(nodeStatus)
		}
		for _, item := range containersStatus {
			log.Println(item.ContainerId, item.Capability, item.Utilization, item.Priority)
			log.Println(item.Priority.Priority)
		}
	}
}
