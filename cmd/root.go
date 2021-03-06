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
	"os"
	"strings"

	c "github.com/magnusfurugard/flinkctl/cluster"
	"github.com/magnusfurugard/flinkctl/config"
	"github.com/magnusfurugard/flinkctl/tools"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cl           c.Cluster
	outputFormat string
)

func Print(t interface{}) {
	tools.Printer(outputFormat, t)
}

func InitCluster() {
	conf, err := config.GetCurrent()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	hostString := strings.TrimSpace(conf.URL)
	cl = c.New(hostString)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:       "flinkctl",
	Short:     "Manage Flink applications.",
	ValidArgs: []string{"cluster", "completion", "config", "describe", "edit", "get", "help", "submit-job"},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "outputFormat", "o", "table", "output format, supports: json,table")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(home)
	viper.SetConfigName(".flinkctl")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.ReadInConfig()
}
