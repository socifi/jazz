package cmd

import (
	"github.com/spf13/cobra"
)

var (
	durable    bool
	autoDelete bool
	wait       bool
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates new object in RabbitMQ",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	createCmd.PersistentFlags().BoolVarP(&durable, "durable", "d", true, "Durability, default is true")
	createCmd.PersistentFlags().BoolVarP(&autoDelete, "auto-delete", "a", false, "Autodelete, default is false")
	createCmd.PersistentFlags().BoolVarP(&wait, "nowait", "n", false, "Waiting, default is true (nowait is false)")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
