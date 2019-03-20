// Config is package that takes care about configuration of clusters, their credentials and so on
package config

import (
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"path/filepath"
)

// ParseConfig get configuration location and parses it's content
func ParseConfig() (*Config, error) {
	c, err := GetConfigHome()
	if err != nil {
		return nil, err
	}
	f, err := os.Open(filepath.Join(c, "config"))
	if err != nil {
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

// SwitchContext tries to switch to another existing context
func (c *Config) SwitchContext(ctx string) error {
	if _, ok := c.Contexts[ctx]; !ok {
		return NewTraverseError("Key does not exist", ctx)
	}
	c.CurrentContext = ctx
	return nil
}

// AddContext tries to add new context
func (c *Config) AddContext(name, cluster, user string) error {
	if _, ok := c.Contexts[name]; ok {
		return NewTraverseError("Key exists", name)
	}
	ctx := Context{cluster, user}
	c.Contexts[name] = ctx
	return nil
}

// ChangeContext either adds or changes context to desired state
func (c *Config) ChangeContext(name, cluster, user string) {
	ctx := Context{cluster, user}
	c.Contexts[name] = ctx
}

// RemoveContext removes a context
func (c *Config) RemoveContext(name string) {
	delete(c.Contexts, name)
}

// AddUser tries to add new user
func (c *Config) AddUser(name, user, password string) error {
	if _, ok := c.Users[name]; ok {
		return NewTraverseError("Key exists", name)
	}
	u := User{user, password}
	c.Users[name] = u
	return nil
}

// ChangeUser either adds or changes user to desired state
func (c *Config) ChangeUser(name, user, password string) {
	u := User{user, password}
	c.Users[name] = u
}

// RemoveUser removes a user
func (c *Config) RemoveUser(name string) {
	delete(c.Users, name)
}

// AddCluster tries to add new cluster
func (c *Config) AddCluster(name, url string, port int) error {
	if _, ok := c.Clusters[name]; ok {
		return NewTraverseError("Key exists", name)
	}
	s := Cluster{url, port}
	c.Clusters[name] = s
	return nil
}

// ChangeCluster either adds or changes cluster to desired state
func (c *Config) ChangeCluster(name, url string, port int) {
	s := Cluster{url, port}
	c.Clusters[name] = s
}

// RemoveCluster removes a cluster
func (c *Config) RemoveCluster(name string) {
	delete(c.Clusters, name)
}
