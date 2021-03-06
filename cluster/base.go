package cluster

import (
	"encoding/json"
	"io/ioutil"

	"github.com/magnusfurugard/flinkctl/tools"
	"github.com/parnurzeal/gorequest"
)

type ClusterConfig struct {
	Features struct {
		WebSubmit bool `json:"web-submit" header:"web-submit"`
	} `json:"features" header:"features"`
	FlinkRevision   string `json:"flink-revision" header:"flink-revision"`
	FlinkVersion    string `json:"flink-version" header:"flink-version"`
	RefreshInterval int64  `json:"refresh-interval" header:"refresh-interval"`
	TimezoneName    string `json:"timezone-name" header:"timezone-name"`
	TimezoneOffset  int64  `json:"timezone-offset" header:"timezone-offset"`
}

type ClusterOverview struct {
	FlinkCommit    string `json:"flink-commit" header:"flink-commit"`
	FlinkVersion   string `json:"flink-version" header:"flink-version"`
	JobsCancelled  int64  `json:"jobs-cancelled" header:"jobs-cancelled"`
	JobsFailed     int64  `json:"jobs-failed" header:"jobs-failed"`
	JobsFinished   int64  `json:"jobs-finished" header:"jobs-finished"`
	JobsRunning    int64  `json:"jobs-running" header:"jobs-running"`
	SlotsAvailable int64  `json:"slots-available" header:"slots-available"`
	SlotsTotal     int64  `json:"slots-total" header:"slots-total"`
	Taskmanagers   int64  `json:"taskmanagers" header:"taskmanagers"`
}

func (c *Cluster) GetConfig() (ClusterConfig, error) {
	re, _, _ := tools.ApplyHeadersToRequest(gorequest.New().Get(c.ConfigURL.String())).End()
	//TODO: Error management
	body, err := ioutil.ReadAll(re.Body)
	if err != nil {
		return ClusterConfig{}, err
	}

	s := ClusterConfig{}
	err = json.Unmarshal(body, &s)
	if err != nil {
		return ClusterConfig{}, err
	}
	return s, nil
}

func (c *Cluster) GetOverview() (ClusterOverview, error) {
	re, _, _ := tools.ApplyHeadersToRequest(gorequest.New().Get(c.OverviewURL.String())).End()
	//TODO: Error management
	body, err := ioutil.ReadAll(re.Body)
	if err != nil {
		return ClusterOverview{}, err
	}

	s := ClusterOverview{}
	err = json.Unmarshal(body, &s)
	if err != nil {
		return ClusterOverview{}, err
	}
	return s, nil

}
