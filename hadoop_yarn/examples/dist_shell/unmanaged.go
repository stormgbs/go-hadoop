package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/stormgbs/go-hadoop/hadoop_common/security"
	"github.com/stormgbs/go-hadoop/hadoop_yarn"
	"github.com/stormgbs/go-hadoop/hadoop_yarn/conf"
	"github.com/stormgbs/go-hadoop/hadoop_yarn/yarn_client"
)

func main() {
	var err error

	// Create YarnConfiguration
	conf, _ := conf.NewYarnConfiguration()

	// Create YarnClient
	yarnClient, _ := yarn_client.CreateYarnClient(conf)

	// Create new application to get ApplicationSubmissionContext
	_, asc, _ := yarnClient.CreateNewApplication()

	// Some useful information
	queue := "default"
	appName := "simple-go-yarn-app"
	appType := "GO_HADOOP_UNMANAGED"
	unmanaged := true
	clc := hadoop_yarn.ContainerLaunchContextProto{}

	// Setup ApplicationSubmissionContext for the application
	asc.AmContainerSpec = &clc
	asc.ApplicationName = &appName
	asc.Queue = &queue
	asc.ApplicationType = &appType
	asc.UnmanagedAm = &unmanaged

	// Submit!
	err = yarnClient.SubmitApplication(asc)
	if err != nil {
		log.Fatal("yarnClient.SubmitApplication ", err)
	}
	log.Println("Successfully submitted unmanaged application: ", asc.ApplicationId)
	time.Sleep(1 * time.Second)

	appReport, err := yarnClient.GetApplicationReport(asc.ApplicationId)
	if err != nil {
		log.Fatal("yarnClient.GetApplicationReport ", err)
	}
	appState := appReport.GetYarnApplicationState()
	for appState != hadoop_yarn.YarnApplicationStateProto_ACCEPTED {
		log.Println("Application in state ", appState)
		time.Sleep(1 * time.Second)
		appReport, err = yarnClient.GetApplicationReport(asc.ApplicationId)
		appState = appReport.GetYarnApplicationState()
		if appState == hadoop_yarn.YarnApplicationStateProto_FAILED || appState == hadoop_yarn.YarnApplicationStateProto_KILLED {
			log.Fatal("Application in state ", appState)
		}
	}

	amRmToken := appReport.GetAmRmToken()

	if amRmToken != nil {
		savedAmRmToken := *amRmToken
		//in the unmanaged AM scenario, the returned token does NOT have the "service" field set
		service, _ := conf.GetRMSchedulerAddress()
		security.GetCurrentUser().AddUserTokenWithAlias(service, &savedAmRmToken)
		log.Printf("saved: service=%s, token=%v\n", service, amRmToken)
	}

	log.Println("Application in state ", appState)

	// Create AMRMClient
	var attemptId int32
	attemptId = 1
	applicationAttemptId := hadoop_yarn.ApplicationAttemptIdProto{ApplicationId: asc.ApplicationId, AttemptId: &attemptId}

	rmClient, _ := yarn_client.CreateAMRMClient(conf)
	log.Println("Created RM client: ", rmClient)

	// Wait for ApplicationAttempt to be in Launched state
	appAttemptReport, err := yarnClient.GetApplicationAttemptReport(&applicationAttemptId)
	appAttemptState := appAttemptReport.GetYarnApplicationAttemptState()
	for appAttemptState != hadoop_yarn.YarnApplicationAttemptStateProto_APP_ATTEMPT_LAUNCHED {
		log.Println("ApplicationAttempt in state ", appAttemptState)
		time.Sleep(1 * time.Second)
		appAttemptReport, err = yarnClient.GetApplicationAttemptReport(&applicationAttemptId)
		appAttemptState = appAttemptReport.GetYarnApplicationAttemptState()
	}
	log.Println("ApplicationAttempt in state ", appAttemptState)

	// Register with ResourceManager
	log.Println("About to register application master.")
	err = rmClient.RegisterApplicationMaster("", -1, "")
	if err != nil {
		log.Fatal("rmClient.RegisterApplicationMaster ", err)
	}
	log.Println("Successfully registered application master.")

	// Add resource requests
	const numContainers = int32(1)
	memory := int32(128)
	vcores := int32(1)
	resource := hadoop_yarn.ResourceProto{Memory: &memory, VirtualCores: &vcores}
	rmClient.AddRequest(1, "*", &resource, numContainers)

	// Now call ResourceManager.allocate
	allocateResponse, err := rmClient.Allocate()
	if err == nil {
		log.Println("allocateResponse: ", *allocateResponse)
	}
	log.Println("#containers allocated: ", len(allocateResponse.AllocatedContainers))

	numAllocatedContainers := int32(0)
	allocatedContainers := make([]*hadoop_yarn.ContainerProto, numContainers, numContainers)
	for numAllocatedContainers < numContainers {
		// Sleep for a while
		log.Println("Sleeping...")
		time.Sleep(3 * time.Second)
		log.Println("Sleeping... done!")

		// Try to get containers now...
		allocateResponse, err = rmClient.Allocate()
		if err != nil {
			log.Printf("rmClient.Allocate error: %v, trying ...", err)
			continue
		}
		bys, _ := json.Marshal(*allocateResponse)
		log.Println("allocateResponse: ", string(bys))

		for _, container := range allocateResponse.AllocatedContainers {
			allocatedContainers[numAllocatedContainers] = container
			numAllocatedContainers++
		}

		log.Println("#containers allocated: ", len(allocateResponse.AllocatedContainers))
		log.Println("Total #containers allocated so far: ", numAllocatedContainers)
	}
	log.Println("Final #containers allocated: ", numAllocatedContainers)

	// Now launch containers
	containerLaunchContext := hadoop_yarn.ContainerLaunchContextProto{Command: []string{"/bin/date"}}
	log.Println("containerLaunchContext: ", containerLaunchContext)
	for _, container := range allocatedContainers {
		log.Println("Launching container: ", *container, " ", *container.NodeId.Host, ":", *container.NodeId.Port)
		nmClient, err := yarn_client.CreateAMNMClient(*container.NodeId.Host, int(*container.NodeId.Port))
		if err != nil {
			log.Fatal("hadoop_yarn.DialContainerManagementProtocolService: ", err)
		}
		log.Printf("Successfully created nmClient: %v", nmClient)

		err = nmClient.StartContainer(container, &containerLaunchContext)
		if err != nil {
			log.Fatal("nmClient.StartContainer: ", err)
		}
	}

	// Wait for Containers to complete
	numCompletedContainers := int32(0)
	for numCompletedContainers < numContainers {
		// Sleep for a while
		log.Println("Sleeping...")
		time.Sleep(3 * time.Second)
		log.Println("Sleeping... done!")

		allocateResponse, err = rmClient.Allocate()
		if err == nil {
			log.Println("allocateResponse: ", *allocateResponse)
		}

		for _, containerStatus := range allocateResponse.CompletedContainerStatuses {
			log.Println("Completed container: ", *containerStatus, " exit-code: ", *containerStatus.ExitStatus)
			numCompletedContainers++
		}
	}
	log.Println("Containers complete: ", numCompletedContainers)

	// Unregister with ResourceManager
	log.Println("About to unregister application master.")
	finalStatus := hadoop_yarn.FinalApplicationStatusProto_APP_SUCCEEDED
	err = rmClient.FinishApplicationMaster(&finalStatus, "done", "")
	if err != nil {
		log.Fatal("rmClient.RegisterApplicationMaster ", err)
	}
	log.Println("Successfully unregistered application master.")
}
