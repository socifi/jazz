package jazz

type Exchange struct {
	Durable    bool      `yaml:"durable"`
	Autodelete bool      `yaml:"autodelete"`
	Internal   bool      `yaml:"internal"`
	Nowait     bool      `yaml:"nowait"`
	Type       string    `yaml:"type"`
	Bindings   []Binding `yaml:"bindings"`
}

type Binding struct {
	Exchange string `yaml:"exchange"`
	Key string      `yaml:"key"`
	Nowait bool     `yaml:"nowait"`
}

type QueueSpec struct {
	Durable    bool      `yaml:"durable"`
	Autodelete bool      `yaml:"autodelete"`
	Nowait     bool      `yaml:"nowait"`
	Exclusive  bool      `yaml:"exclusive"`
	Bindings   []Binding `yaml:"bindings"`
}

type Settings struct {
	Exchanges map[string]Exchange  `yaml:"exchanges"`
	Queues    map[string]QueueSpec `yaml:"queues"`
}
