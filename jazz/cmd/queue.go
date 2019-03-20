// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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

	"github.com/spf13/cobra"
)

var (
	exclusive bool
)

// queueCmd represents the queue command
var queueCmd = &cobra.Command{
	Use:   "queue",
	Short: "Creates RabbitMQ queue",
	Long:  `Creates queue on RabbitMQ cluster`,
	Args:  cobra.ExactArgs(1),
	Run:   createQueue,
}

func createQueue(cmd *cobra.Command, args []string) {
	fmt.Println("Queue called")
	fmt.Println("durable:", durable)
	fmt.Println("auto-delete:", autoDelete)
	fmt.Println("exclusive:", exclusive)
}

func init() {
	createCmd.AddCommand(queueCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	queueCmd.PersistentFlags().BoolVarP(&exclusive, "exclusive", "e", false, "Queue is exclusive, default false")
	//	queueCmd.PersistentFlags().BoolVarP(&exclusive, "exclusive", "e", false, "Queue is exclusive, default false")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queueCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
