/*
Copyright © 2020 Magnus Furugård <magnus.furugard@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"io/ioutil"

	"github.com/magnusfurugard/flinkctl/tools"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

type JarsResponse struct {
	Address string `json:"address"`
	Files   []struct {
		Entry []struct {
			Description string `json:"description" header:"entry-description"`
			Name        string `json:"name" header:"entry-name"`
		} `json:"entry"`
		ID       string `json:"id" header:"id"`
		Name     string `json:"name" header:"name"`
		Uploaded int64  `json:"uploaded" header:"uploaded"`
	} `json:"files"`
}

var getJarsCmd = &cobra.Command{
	Use:    "jars",
	Short:  "List all uploaded jars in your cluster",
	PreRun: func(cmd *cobra.Command, args []string) { InitCluster() },
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, _, _ := tools.ApplyHeadersToRequest(gorequest.New().Get(cl.Jars.URL.String())).End()
		//TODO: Error management
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		s := JarsResponse{}
		err = json.Unmarshal(body, &s)
		if err != nil {
			return err
		}
		Print(s.Files)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getJarsCmd)
}
