package jazz

// Exchange is structure with specification of properties of RabbitMQ exchange
type Exchange struct {
	Durable    bool      `yaml:"durable"`
	Autodelete bool      `yaml:"autodelete"`
	Internal   bool      `yaml:"internal"`
	Nowait     bool      `yaml:"nowait"`
	Type       string    `yaml:"type"`
	Bindings   []Binding `yaml:"bindings"`
}

// Binding specifies to which exchange should be an exchange or a queue binded
type Binding struct {
	Exchange string `yaml:"exchange"`
	Key      string `yaml:"key"`
	Nowait   bool   `yaml:"nowait"`
}

// QueueSpec is a specification of properties of RabbitMQ queue
type QueueSpec struct {
	Durable    bool      `yaml:"durable"`
	Autodelete bool      `yaml:"autodelete"`
	Nowait     bool      `yaml:"nowait"`
	Exclusive  bool      `yaml:"exclusive"`
	Bindings   []Binding `yaml:"bindings"`
}

// Settings is a specification of all queues and exchanges together with all bindings.
type Settings struct {
	Exchanges map[string]Exchange  `yaml:"exchanges"`
	Queues    map[string]QueueSpec `yaml:"queues"`
}
