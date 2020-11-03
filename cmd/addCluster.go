/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"net/url"

	"github.com/magnusfurugard/flinkctl/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var headers []string

// addClusterCmd represents the addCluster command
var addClusterCmd = &cobra.Command{
	Use:   "add-cluster <url:port>",
	Short: "Add a new cluster to your flinkctl config",
	Example: `flinkctl config add-cluster https://localhost:123
flinkctl config add-cluster https://localhost:567 --headers="Authorization: Basic Zm9v,Content-Type: application/json"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("add-cluster requires exactly 1 positional argument, not %v", len(args))
		}

		u, err := url.Parse(args[0])
		if err != nil {
			return err
		}

		currentConfig := config.Get()
		newConfig := config.ClusterConfig{URL: u.String(), Headers: headers}
		if len(currentConfig.Clusters) == 0 {
			viper.SetDefault("clusters", newConfig)
			viper.SetDefault("current-cluster", u.String())
			viper.SafeWriteConfig()
			viper.ReadInConfig()
			fmt.Printf("Config file created: %v\n", viper.ConfigFileUsed())
		} else {
			viper.SetDefault("clusters", append(currentConfig.Clusters, newConfig))
			viper.SetDefault("current-cluster", u.String())
			viper.WriteConfig()
			fmt.Printf("current-cluster updated: %v\n", u.String())
		}
		return nil
	},
}

func init() {
	configCmd.AddCommand(addClusterCmd)
	addClusterCmd.Flags().StringSliceVar(&headers, "headers", []string{}, "additional headers to pass when calling this cluster")
}
