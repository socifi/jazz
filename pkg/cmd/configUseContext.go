package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var configUseContextCmd = &cobra.Command{
	Use:   "use-context context",
	Short: "Switch current context",
	Long:  `Switch current context`,
	Args:  cobra.ExactArgs(1),
	Run:   useContext,
}

func useContext(cmd *cobra.Command, args []string) {
	err := cfg.UseContext(args[0])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
	configCmd.AddCommand(configUseContextCmd)
}
