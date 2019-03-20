package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// ConfigCmd represents the config command
var configCmd = &cobra.Command{
	Use:               "config",
	Short:             "Configuration of jazz",
	Long:              `You can configure mainly clusters and login information`,
	Run:               run,
	PersistentPostRun: configPersistentPostRunHook,
}

func configPersistentPostRunHook(cmd *cobra.Command, args []string) {
	saveConfig()
}

func saveConfig() {
	err := cfg.SaveCofig()
	if err != nil {
		fmt.Println("Error saving configuration:", err.Error())
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func init() {
	rootCmd.AddCommand(configCmd)
}
