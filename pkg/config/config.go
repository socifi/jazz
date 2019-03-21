// Config is package that takes care about configuration of clusters, their credentials and so on
package config

// NewConfig creates new configuration structure
func NewConfig() *Config {
	contexts := map[string]Context{}
	clusters := map[string]Cluster{}
	users := map[string]User{}
	return &Config{clusters, contexts, "", users}
}

// SwitchContext tries to switch to another existing context
func (c *Config) UseContext(ctx string) error {
	if _, ok := c.Contexts[ctx]; !ok {
		return NewTraverseError("Context does not exist", ctx)
	}
	c.CurrentContext = ctx
	return nil
}

// AddContext tries to add new context
func (c *Config) AddContext(name, cluster, user string) error {
	if _, ok := c.Contexts[name]; ok {
		return NewTraverseError("Context already exists", name)
	}
	if _, ok := c.Users[user]; !ok {
		return NewTraverseError("User does not exists", user)
	}
	if _, ok := c.Clusters[cluster]; !ok {
		return NewTraverseError("Cluster does not exists", cluster)
	}
	ctx := Context{cluster, user}
	c.Contexts[name] = ctx
	return nil
}

// ChangeContext either adds or changes context to desired state
func (c *Config) ChangeContext(name, cluster, user string) error {
	if _, ok := c.Users[user]; !ok {
		return NewTraverseError("User does not exists", user)
	}
	if _, ok := c.Clusters[cluster]; !ok {
		return NewTraverseError("Cluster does not exists", cluster)
	}
	ctx := Context{cluster, user}
	c.Contexts[name] = ctx
	return nil
}

// RemoveContext removes a context
func (c *Config) RemoveContext(name string) {
	delete(c.Contexts, name)
}

// AddUser tries to add new user
func (c *Config) AddUser(name, user, password string) error {
	if _, ok := c.Users[name]; ok {
		return NewTraverseError("User already exists", name)
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
