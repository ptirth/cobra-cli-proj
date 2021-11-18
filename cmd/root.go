/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobra-cli-proj",
	Short: "A brief description of your application",
	Long: `This application has few capabilities listed below:

cobra-cli-proj is a CLI tool that empowers users to, make directories, files.
It also allows users to add data to files, print all data from the file and 
also search through the file to print specific data.

Examples:
	- cobra-cli-proj.exe mkd Test  - Makes directory named Test. You can also assign full path while creating directory
	- cobra-cli-proj.exe mkf test.txt  -  Makes test.txt file at current location. Can also provde full location.
	- cobra-cli-proj.exe appendF "Hi, This is test file" test.txt  -  Appends user input text to the file.
	- cobra-cli-proj.exe readF test.txt  -  Reads the content of file and prints it on the console.
	- cobra-cli-proj.exe writestructfile .\Data.txt 9 Patel 52 9984126238 - Write user given input to the file and saves it in specific format using | to seperate data.
	- cobra-cli-proj.exe datastruct .\Data.txt - Prints data from given file on console using structs to properly formatted.
	- cobra-cli-proj.exe searchdata Tirth .\Data.txt - Searches the given file for "Tirth" and prints the row with all data.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra-cli-proj.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra-cli-proj" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra-cli-proj")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
