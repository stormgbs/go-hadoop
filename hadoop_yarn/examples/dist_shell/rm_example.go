package main

import (
	"log"

	"github.com/stormgbs/go-hadoop/hadoop_yarn/conf"
	"github.com/stormgbs/go-hadoop/hadoop_yarn/yarn_client"
)

func main() {
	var err error

	// Create YarnConfiguration
	conf, _ := conf.NewYarnConfiguration()

	// Create RMClient
	rmClient, _ := yarn_client.CreateRMClient(conf)

	// Some useful information
	var (
		//host string = "node004"
		//TODO: Get port of node from RM
		//port     int32 = 31233
		memoryMB int32 = 4096
		vcores   int32 = 4
	)

	// Update node resource
	//err = rmClient.UpdateNodeResource(host, port, memoryMB, vcores)
	err = rmClient.UpdateResourceOfGivenNode(memoryMB, vcores)
	if err != nil {
		log.Fatal("rmClient.UpdateNodeResource: ", err)
	}

}
