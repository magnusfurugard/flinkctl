package cluster

// TODO: Once we have generics, refactor all Get* functions :)

import (
	"encoding/json"
	"io/ioutil"

	"github.com/magnusfurugard/flinkctl/tools"
	"github.com/parnurzeal/gorequest"
)

type Jobs struct {
	Jobs []struct {
		ID     string `json:"id" header:"id"`
		Status string `json:"status" header:"status"`
	} `json:"jobs" header:"jobs"`
}

type JobsOverview struct {
	Jobs []struct {
		Duration         int64  `json:"duration" header:"duration"`
		EndTime          int64  `json:"end-time" header:"end-time"`
		Jid              string `json:"jid" header:"jid"`
		LastModification int64  `json:"last-modification" header:"last-modification"`
		Name             string `json:"name" header:"name"`
		StartTime        int64  `json:"start-time" header:"start-time"`
		State            string `json:"state" header:"state"`
		Tasks            struct {
			Canceled    int64 `json:"canceled" header:"canceled"`
			Canceling   int64 `json:"canceling" header:"canceling"`
			Created     int64 `json:"created" header:"created"`
			Deploying   int64 `json:"deploying" header:"deploying"`
			Failed      int64 `json:"failed" header:"failed"`
			Finished    int64 `json:"finished" header:"finished"`
			Reconciling int64 `json:"reconciling" header:"reconciling"`
			Running     int64 `json:"running" header:"running"`
			Scheduled   int64 `json:"scheduled" header:"scheduled"`
			Total       int64 `json:"total" header:"total"`
		} `json:"tasks" header:"tasks"`
	} `json:"jobs" header:"jobs"`
}

func (c *Cluster) GetJobs() (Jobs, error) {
	re, _, _ := tools.ApplyHeadersToRequest(gorequest.New().Get(c.Jobs.URL.String())).End()
	//TODO: Error management
	body, err := ioutil.ReadAll(re.Body)
	if err != nil {
		return Jobs{}, err
	}

	s := Jobs{}
	json.Unmarshal(body, &s)
	return s, nil
}

func (c *Cluster) GetJobsOverview() (JobsOverview, error) {
	// TODO: Fix .duration to be a duration
	// TODO: .last-modification and start-time shuold be datetimes (UTC)
	re, _, _ := tools.ApplyHeadersToRequest(gorequest.New().Get(c.Jobs.OverviewURL.String())).End()
	//TODO: Error management
	body, err := ioutil.ReadAll(re.Body)
	if err != nil {
		return JobsOverview{}, err
	}

	s := JobsOverview{}
	json.Unmarshal(body, &s)
	return s, nil
}
