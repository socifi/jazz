package cmd

import (
	"github.com/spf13/cobra"
)

var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "List configuration of jazz",
	Long:  `List configuration of jazz`,
	Run:   list,
}

func list(cmd *cobra.Command, args []string) {
	cfg.Print()
}

func init() {
	configCmd.AddCommand(configListCmd)
}
