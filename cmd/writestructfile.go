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
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// writestructfileCmd represents the writestructfile command
var writestructfileCmd = &cobra.Command{
	Use:   "writestructfile",
	Short: "Writes user input in file with specific format. User enter data with space.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		writedata(args[0], args[1:])
	},
}

func init() {
	rootCmd.AddCommand(writestructfileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// writestructfileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// writestructfileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Dataa struct {
	No       int
	Name     string
	Age      int
	PhNumber int
}

func writedata(data1 string, args []string) {

	//name := strings.Join(args, " ")

	//slice := make([]Dat, 0, 10)

	file, e := os.OpenFile(filepath.Join(data1), os.O_APPEND, 0644)
	if e != nil {
		fmt.Println("Error is = ", e)
	}
	ar := strings.Join(args, " ")
	s := strings.Split(ar, " ")
	var d Dat

	d.No, _ = strconv.Atoi(s[0])
	d.Name = s[1]
	d.Age, _ = strconv.Atoi(s[2])
	d.PhNumber, _ = strconv.Atoi(s[3])

	fmt.Fprintf(file, "%v|%v|%v|%v\n", d.No, d.Name, d.Age, d.PhNumber)

	file.Close()
}
