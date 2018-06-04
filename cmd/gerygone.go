// Copyright © 2018 Albert Moreno <albert.moreno@protonmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// gerygoneCmd represents the base command when called without any subcommands
var gerygoneCmd = &cobra.Command{
	Version: "0.0.1-dev",
	Use:     "gerygone",
	Short:   "Quick and Easy infrastructure validation",
	Long: `Gerygone is a YAML based tool for validating a infrastructure’s configuration. 
It eases the process of writing tests by allowing the user to generate tests from the current system state. 
Once the test suite is written they can be executed, waited-on, or served as a health endpoint.`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the gerygoneCmd.
func Execute() {
	if err := gerygoneCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	gerygoneCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gerygone.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	gerygoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".gerygone" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gerygone")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
