package main

import (
	"bytes"
	"fmt"
	"github.com/socifi/jazz"
)

var dsn = "amqp://guest:guest@localhost:5672/"

var data = []byte(`
exchanges:
  exchange0:
    durable: true
    type: topic
  exchange1:
    durable: true
    type: topic
    bindings:
      - exchange: "exchange0"
        key: "key1"
      - exchange: "exchange0"
        key: "key2"
  exchange2:
    durable: true
    type: topic
    bindings:
      - exchange: "exchange0"
        key: "key3"
      - exchange: "exchange1"
        key: "key2"
  exchange3:
    durable: true
    type: topic
    bindings:
      - exchange: "exchange0"
        key: "key4"
queues:
  queue0:
    durable: true
    bindings:
      - exchange: "exchange0"
        key: "key4"
  queue1:
    durable: true
    bindings:
      - exchange: "exchange1"
        key: "key2"
  queue2:
    durable: true
    bindings:
      - exchange: "exchange1"
        key: "#"
  queue3:
    durable: true
    bindings:
      - exchange: "exchange2"
        key: "#"
  queue4:
    durable: true
    bindings:
      - exchange: "exchange3"
        key: "#"
  queue5:
    durable: true
    bindings:
      - exchange: "exchange0"
        key: "#"
`)

func main() {
	c, err := jazz.Connect(dsn)
	if err != nil {
		panic(fmt.Sprintf("Could not connect to RabbitMQ: %v", err.Error()))
	}

	reader := bytes.NewReader(data)
	scheme, err := jazz.DecodeYaml(reader)
	if err != nil {
		panic(fmt.Sprintf("Could not read YAML: %v", err.Error()))
	}
	// Create scheme
	err = c.CreateScheme(scheme)
	if err != nil {
		panic(fmt.Sprintf("Could not create scheme: %v", err.Error()))
	}

	// Handler function
	f := func(msg []byte) {
		fmt.Println(string(msg))
	}

	go c.ProcessQueue("queue0", f)
	go c.ProcessQueue("queue1", f)
	go c.ProcessQueue("queue2", f)
	go c.ProcessQueue("queue3", f)
	go c.ProcessQueue("queue4", f)
	go c.ProcessQueue("queue5", f)
	c.SendMessage("exchange0", "key1", "Hello World!")
	c.SendMessage("exchange0", "key2", "Hello!")
	c.SendMessage("exchange0", "key3", "World!")
	c.SendMessage("exchange0", "key4", "Hi!")
	c.SendMessage("exchange0", "key5", "Again!")

	// Be nice and clean up a little bit. Not advisable in production.
	err = c.DeleteScheme(scheme)
	if err != nil {
		panic(fmt.Sprintf("Could not delete scheme: %v", err.Error()))
	}
	c.Close()
}
