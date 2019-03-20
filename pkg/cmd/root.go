package cmd

import (
	"fmt"
	"os"

	"github.com/howeyc/gopass"
	"github.com/socifi/jazz/pkg/config"
	"github.com/spf13/cobra"

	"github.com/streadway/amqp"
)

var (
	cfg      *config.Config
	host     string
	password []byte
	pass     bool
	user     string
	dsn      string
)

var rootCmd = &cobra.Command{
	Use:   "jazz",
	Short: "Simple RabbitMQ management tool",
	Long: `Jazz is a simple RabbitMQ management tool.

It allows you to create, change and bind exchanges and queues.
You can use either simple atomic commands or pass entire structure via yaml.`,
	PersistentPreRun: rootPersistentPreRunHook,
}

func rootPersistentPreRunHook(cmd *cobra.Command, args []string) {
	GetConfig()
	GetDsn()
}

func GetConfig() {
	var err error
	cfg, err = config.ParseConfig()
	if err != nil {
		fmt.Println("Error loading configuration:", err.Error())
		os.Exit(1)
	}
}

func GetDsn() {
	var err error
	var uri amqp.URI

	if host != "" {
		uri, err = amqp.ParseURI(host)
	} else {
		dsn = cfg.GetCurrentContextDsn()
		uri, err = amqp.ParseURI(dsn)
	}

	if err != nil {
		fmt.Println("Error parsing host")
		os.Exit(1)
	}

	if user != "" {
		uri.Username = user
	}
	if pass {
		fmt.Printf("Enter password >")
		password, err = gopass.GetPasswd()
		if err != nil {
			fmt.Println("Error obtaining password")
			os.Exit(1)
		}
		uri.Password = string(password)
	}
	dsn = uri.String()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jazz.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&pass, "password", "p", false, "Ask for password")
	rootCmd.PersistentFlags().StringVarP(&host, "host", "H", "", "RabbitMQ address")
	rootCmd.PersistentFlags().StringVarP(&user, "user", "u", "", "RabbitMQ user name")
}
