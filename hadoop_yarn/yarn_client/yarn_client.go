package yarn_client

import (
  "github.com/stormgbs/go-hadoop/hadoop_yarn"
  yarn_conf "github.com/stormgbs/go-hadoop/hadoop_yarn/conf"
)

type YarnClient struct {
  client hadoop_yarn.ApplicationClientProtocolService
}

func CreateYarnClient(conf yarn_conf.YarnConfiguration) (*YarnClient, error) {
  c, err := hadoop_yarn.DialApplicationClientProtocolService(conf)
  return &YarnClient{client: c}, err
}

func (c *YarnClient) CreateNewApplication () (*hadoop_yarn.GetNewApplicationResponseProto, *hadoop_yarn.ApplicationSubmissionContextProto, error) {
  request := hadoop_yarn.GetNewApplicationRequestProto{}
  response := hadoop_yarn.GetNewApplicationResponseProto{}
  err := c.client.GetNewApplication(&request, &response)
  if err != nil {
    return nil, nil, err
  }

  return &response, &hadoop_yarn.ApplicationSubmissionContextProto{ApplicationId: response.ApplicationId}, nil
}

func (c *YarnClient) SubmitApplication (asc *hadoop_yarn.ApplicationSubmissionContextProto) (error) {
  request := hadoop_yarn.SubmitApplicationRequestProto{ApplicationSubmissionContext: asc}
  response := hadoop_yarn.SubmitApplicationResponseProto{}
  return c.client.SubmitApplication(&request, &response)
}

func (c *YarnClient) GetApplicationReport (applicationId *hadoop_yarn.ApplicationIdProto) (*hadoop_yarn.ApplicationReportProto, error) {
  request := hadoop_yarn.GetApplicationReportRequestProto{ApplicationId: applicationId}
  response := hadoop_yarn.GetApplicationReportResponseProto{}
  err := c.client.GetApplicationReport(&request, &response)
  return response.ApplicationReport, err
}

func (c *YarnClient) GetApplicationAttemptReport (applicationAttemptId *hadoop_yarn.ApplicationAttemptIdProto) (*hadoop_yarn.ApplicationAttemptReportProto, error) {
  request := hadoop_yarn.GetApplicationAttemptReportRequestProto{ApplicationAttemptId: applicationAttemptId}
  response := hadoop_yarn.GetApplicationAttemptReportResponseProto{}
  err := c.client.GetApplicationAttemptReport(&request, &response)
  return response.ApplicationAttemptReport, err
}

