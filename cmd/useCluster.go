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
	"fmt"

	"github.com/magnusfurugard/flinkctl/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// useClusterCmd represents the useCluster command
var useClusterCmd = &cobra.Command{
	Use:     "use-cluster <cluster url>",
	Short:   "Change the currently in-use config for flinkctl",
	Example: `flinkctl config use-cluster https://localhost:123`,
	Args:    cobra.ExactValidArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		url := args[0]
		if !config.ConfigExists(url) {
			return fmt.Errorf("no such cluster exists in your config")
		}
		viper.SetDefault("current-cluster", url)
		viper.WriteConfig()
		fmt.Printf("current cluster updated: %v\n", config.GetCurrentName())
		return nil
	},
}

func init() {
	configCmd.AddCommand(useClusterCmd)
}
