package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	configCmd.AddCommand(configRemoveCmd)
	configRemoveCmd.AddCommand(configRemoveClusterCmd)
	configRemoveCmd.AddCommand(configRemoveContextCmd)
	configRemoveCmd.AddCommand(configRemoveUserCmd)
}

var configRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove configuration of jazz",
	Long:  `Remove configuration of jazz`,
	Run:   remove,
}

var configRemoveClusterCmd = &cobra.Command{
	Use:   "cluster name url port",
	Short: "Remove existing cluster in configuration",
	Long:  `Remove existing cluster in configuration`,
	Args:  cobra.ExactArgs(1),
	Run:   removeCluster,
}

var configRemoveUserCmd = &cobra.Command{
	Use:   "user name username password",
	Short: "Remove existing user in configuration",
	Long:  `Remove existing user in configuration`,
	Args:  cobra.ExactArgs(1),
	Run:   removeUser,
}

var configRemoveContextCmd = &cobra.Command{
	Use:   "context name user cluster",
	Short: "Remove existing context in configuration",
	Long:  `Remove existing context in configuration`,
	Args:  cobra.ExactArgs(1),
	Run:   removeContext,
}

func remove(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func removeCluster(cmd *cobra.Command, args []string) {
	cfg.RemoveCluster(args[0])
}

func removeContext(cmd *cobra.Command, args []string) {
	cfg.RemoveContext(args[0])
}

func removeUser(cmd *cobra.Command, args []string) {
	cfg.RemoveUser(args[0])
}
