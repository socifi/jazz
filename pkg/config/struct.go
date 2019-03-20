package config

// Cluster contains server specification (URI)
type Cluster struct {
	Server string `yaml:"url"`
	Port   int    `yaml:"port"`
}

// Context contains bindings between servers and users
type Context struct {
	Cluster string `yaml:"cluster"`
	User    string `yaml:"user"`
}

// User contains user data, user name and password
type User struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Config struct {
	Clusters       map[string]Cluster `yaml:"clusters"`
	Contexts       map[string]Context `yaml:"contexts"`
	CurrentContext string             `yaml:"current-context"`
	Users          map[string]User    `yaml:"users"`
}
