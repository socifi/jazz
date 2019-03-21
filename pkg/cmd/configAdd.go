package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func init() {
	configCmd.AddCommand(configAddCmd)
	configAddCmd.AddCommand(configAddClusterCmd)
	configAddCmd.AddCommand(configAddUserCmd)
	configAddCmd.AddCommand(configAddContextCmd)
}

var configAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new configuration to jazz",
	Long:  `Add new configuration to jazz`,
	Run:   add,
}

var configAddClusterCmd = &cobra.Command{
	Use:   "cluster name url port",
	Short: "Add new cluster to configuration",
	Long:  `Add new cluster to configuration`,
	Args:  cobra.ExactArgs(3),
	Run:   addCluster,
}

var configAddUserCmd = &cobra.Command{
	Use:   "user name username password",
	Short: "Add new user to configuration",
	Long:  `Add new user to configuration`,
	Args:  cobra.ExactArgs(3),
	Run:   addUser,
}

var configAddContextCmd = &cobra.Command{
	Use:   "context name user cluster",
	Short: "Add new user to configuration",
	Long:  `Add new user to configuration`,
	Args:  cobra.ExactArgs(3),
	Run:   addContext,
}

func add(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func addCluster(cmd *cobra.Command, args []string) {
	p, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	err = cfg.AddCluster(args[0], args[1], p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func addUser(cmd *cobra.Command, args []string) {
	err := cfg.AddUser(args[0], args[1], args[2])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func addContext(cmd *cobra.Command, args []string) {
	err := cfg.AddContext(args[0], args[1], args[2])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
