package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func init() {
	configCmd.AddCommand(configChangeCmd)
	configChangeCmd.AddCommand(configChangeClusterCmd)
	configChangeCmd.AddCommand(configChangeContextCmd)
	configChangeCmd.AddCommand(configChangeUserCmd)
}

var configChangeCmd = &cobra.Command{
	Use:   "change",
	Short: "Change configuration of jazz",
	Long:  `Change configuration of jazz`,
	Run:   change,
}

var configChangeClusterCmd = &cobra.Command{
	Use:   "cluster name url port",
	Short: "Change existing cluster in configuration",
	Long:  `Change existing cluster in configuration`,
	Args:  cobra.ExactArgs(3),
	Run:   changeCluster,
}

var configChangeUserCmd = &cobra.Command{
	Use:   "user name username password",
	Short: "Change existing user in configuration",
	Long:  `Change existing user in configuration`,
	Args:  cobra.ExactArgs(3),
	Run:   changeUser,
}

var configChangeContextCmd = &cobra.Command{
	Use:   "context name user cluster",
	Short: "Change existing context in configuration",
	Long:  `Change existing context in configuration`,
	Args:  cobra.ExactArgs(3),
	Run:   changeContext,
}

func change(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func changeCluster(cmd *cobra.Command, args []string) {
	p, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	cfg.ChangeCluster(args[0], args[1], p)
}

func changeContext(cmd *cobra.Command, args []string) {
	err := cfg.ChangeContext(args[0], args[1], args[2])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func changeUser(cmd *cobra.Command, args []string) {
	cfg.ChangeUser(args[0], args[1], args[2])
}
