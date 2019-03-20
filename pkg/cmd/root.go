package cmd

import (
	"fmt"
	"os"

	"path/filepath"
	"github.com/howeyc/gopass"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	dsn      string
	cfgFile  string
	password []byte
	pass     bool
	user     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jazz",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		if dsn == "" {
			panic("You need to fill RabbitMQ dsn")
		}

		if pass {
			fmt.Printf("Enter password >")
			password, err = gopass.GetPasswd()
			if err != nil {
				panic(err.Error())
			}
			fmt.Println(string(password))
		}
		dsn, err = getFullDSN(user, string(password), dsn)
		if err != nil {
			panic(err.Error())
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jazz.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&pass, "password", "p", false, "Ask for password")
	rootCmd.PersistentFlags().StringVarP(&dsn, "dsn", "D", "", "RabbitMQ address")
	rootCmd.PersistentFlags().StringVarP(&user, "user", "u", "", "RabbitMQ user name")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		
		// Search for config in home directory with name ".jazz".
		viper.AddConfigPath(filepath.Join(home, ".jazz"))
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	viper.Debug()
}
