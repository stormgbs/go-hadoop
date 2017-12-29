package yarn_client

import (
	"github.com/stormgbs/go-hadoop/hadoop_yarn"
	yarn_conf "github.com/stormgbs/go-hadoop/hadoop_yarn/conf"
	"log"
	"time"
)

type AMRMClientAsync struct {
	client     *AMRMClient
	intervalMs int
	handler    CallBackHandler
	stop       chan bool
}

func CreateAMRMClientAsync(conf yarn_conf.YarnConfiguration, intervalMs int, handler CallBackHandler) (*AMRMClientAsync, error) {
	syncClient, err := CreateAMRMClient(conf)
	return &AMRMClientAsync{client: syncClient, intervalMs: intervalMs, handler: handler, stop: make(chan bool, 1)}, err
}

func (c *AMRMClientAsync) RegisterApplicationMaster(host string, port int32, url string) error {
	err := c.client.RegisterApplicationMaster(host, port, url)

	if err != nil {
		log.Println("failed to register application master!")
		return err
	}

	log.Println("starting periodic allocate routine with interval(ms): ", c.intervalMs)
	go c.periodicAllocate()

	return nil
}

func (c *AMRMClientAsync) FinishApplicationMaster(finalStatus *hadoop_yarn.FinalApplicationStatusProto, message string, url string) error {
	c.stop <- true
	return c.client.FinishApplicationMaster(finalStatus, message, url)
}

func (c *AMRMClientAsync) ReleaseAssignedContainer(containerId *hadoop_yarn.ContainerIdProto) {
	c.client.ReleaseAssignedContainer(containerId)
}

func (c *AMRMClientAsync) AddRequest(priority int32, resourceName string, capability *hadoop_yarn.ResourceProto, numContainers int32) error {
	return c.client.AddRequest(priority, resourceName, capability, numContainers)
}

//Periodically call the  the Resource Manager with an allocate call
func (c *AMRMClientAsync) periodicAllocate() {
	for {
		select {
		case <-c.stop:
			log.Println("stop notification received. stopping allocate routine.")
			return
		default:
		}

		response, err := c.client.Allocate()

		if err != nil {
			log.Println("async allocate failed. terminating periodic allocate. error: ", err)
			return
		}

		if command := response.GetAMCommand(); command != 0 {
			switch command {
			case hadoop_yarn.AMCommandProto_AM_RESYNC, hadoop_yarn.AMCommandProto_AM_SHUTDOWN:
				log.Println("received command: %s. invoking shutdown and stopping allocate routine.", command)
				c.handler.OnShutDownRequest()
				return
			default:
				log.Printf("unknown command: %s. stopping allocate routine.", command)
				return
			}
		}

		if completedContainers := response.CompletedContainerStatuses; completedContainers != nil {
			if len(completedContainers) > 0 {
				c.handler.OnContainersCompleted(completedContainers)
			}
		}

		if allocatedContainers := response.AllocatedContainers; allocatedContainers != nil {
			if len(allocatedContainers) > 0 {
				c.handler.OnContainersAllocated(allocatedContainers)
			}
		}

		if updatedNodes := response.UpdatedNodes; updatedNodes != nil {
			if len(updatedNodes) > 0 {
				c.handler.OnNodesUpdated(updatedNodes)
			}
		}

		time.Sleep(time.Duration(c.intervalMs) * time.Millisecond)
	}
}

type CallBackHandler interface {
	OnContainersCompleted(completedContainers []*hadoop_yarn.ContainerStatusProto)
	OnContainersAllocated(allocatedContainers []*hadoop_yarn.ContainerProto)
	OnShutDownRequest()
	OnNodesUpdated(updatedNodes []*hadoop_yarn.NodeReportProto)
	GetProgress() float64
	OnError(err error)
}
