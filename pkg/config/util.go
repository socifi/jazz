package config

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/streadway/amqp"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Utility functions for configuration, mainly loading, saving and

// Initializes default configuration
func initConfig() {
	defaultConfig := NewConfig()
	defaultConfig.AddCluster("local", "localhost", 5672)
	defaultConfig.AddUser("guest", "guest", "guest")
	defaultConfig.AddContext("guest@local", "local", "guest")
	defaultConfig.UseContext("guest@local")

	defaultConfig.SaveCofig()
}

// ParseConfig get configuration location and parses it's content
func ParseConfig() (*Config, error) {
	c, err := GetConfigHome()
	if err != nil {
		return nil, err
	}
Open:
	f, err := os.Open(filepath.Join(c, "config"))
	if os.IsNotExist(err) {
		// If config does not exist, create it
		initConfig()
		goto Open
	} else if err != nil {
		return nil, err
	}
	defer f.Close()
	return Parse(f)
}

// GetConfigHome retrieves JAZZ_HOME
func GetConfigHome() (string, error) {
	jazz := os.Getenv("JAZZ_HOME")
	if jazz == "" {
		home, err := homedir.Dir()
		if err != nil {
			return "", err
		}
		jazz = filepath.Join(home, ".jazz")
	}
	return jazz, nil
}

// Parse parses yaml configuration from given io.Reader to Config structure
func Parse(r io.Reader) (*Config, error) {
	c := &Config{}
	d := yaml.NewDecoder(r)
	return c, d.Decode(c)
}

// SaveCofig saves configuration to JAZZ_HOME
func (c Config) SaveCofig() error {
	cfg, err := GetConfigHome()
	if err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(cfg, "config"))
	if err != nil {
		return err
	}
	defer f.Close()
	return c.Save(f)
}

// Save saves configuration to given io.Writer
func (c Config) Save(w io.Writer) error {
	e := yaml.NewEncoder(w)
	defer e.Close()
	return e.Encode(c)
}

func (c Config) GetCurrentContextDsn() string {
	return c.GetContextDsn(c.CurrentContext)
}

// GetCurrentContextDsn returns dsn for current contexts' cluster and user
func (c Config) GetContextDsn(name string) string {
	ctx := c.Contexts[name]
	usr := c.Users[ctx.User]
	clu := c.Clusters[ctx.Cluster]

	var uri amqp.URI
	if strings.HasPrefix(clu.Server, "amqp://") {
		uri, _ = amqp.ParseURI(clu.Server)
	} else {
		uri, _ = amqp.ParseURI(strings.Join([]string{"amqp://", clu.Server}, ""))
	}

	uri.Port = clu.Port
	uri.Username = usr.User
	uri.Password = usr.Password
	return uri.String()
}

func (c Config) Print() {
	fmt.Println("Current context:", c.CurrentContext)
	fmt.Println("Available contexts:")
	for k, _ := range c.Contexts {
		fmt.Printf("  %v: %v\n", k, c.GetContextDsn(k))
	}
	fmt.Println("Available clusters:")
	for k, v := range c.Clusters {
		fmt.Printf("  %v: %v:%v\n", k, v.Server, v.Port)
	}
	fmt.Println("Available users:")
	for k, v := range c.Users {
		fmt.Printf("  %v: %v:%v\n", k, v.User, v.Password)
	}
}
