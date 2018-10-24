package main

import(
	"fmt"
	"github.com/socifi/jazz"
)

var dsn = "amqp://guest:guest@localhost:5672/"


var data = []byte(`
exchanges:
  change:
    durable: true
    type: topic
  change1:
    durable: true
    type: topic
    bindings:
      - exchange: "change"
        key: "key1"
      - exchange: "change"
        key: "key2"
  change2:
    durable: true
    type: topic
    bindings:
      - exchange: "change"
        key: "key3"
      - exchange: "change1"
        key: "key2"
  change3:
    durable: true
    type: topic
    bindings:
      - exchange: "change"
        key: "key4"
queues:
  queue1:
    durable: true
    bindings:
      - exchange: "change"
        key: "key4"
  queue2:
    durable: true
    bindings:
      - exchange: "change1"
        key: "key2"
  queue3:
    durable: true
    bindings:
      - exchange: "change1"
        key: "#"
  queue4:
    durable: true
    bindings:
      - exchange: "change2"
        key: "#"
  queue5:
    durable: true
    bindings:
      - exchange: "change3"
        key: "#"
  queue6:
    durable: true
    bindings:
      - exchange: "change"
        key: "#"
`)

func main() {
	c, err := jazz.Connect(dsn)
	if err != nil {
		panic(fmt.Sprintf("Could not connect to RabbitMQ: %v",err.Error()))
	}

	// Create scheme
	err = c.CreateScheme(data)
	if err != nil {
		panic(fmt.Sprintf("Could not create scheme: %v",err.Error()))
	}

	// Handler function
	f := func(msg []byte) {
		fmt.Println(string(msg))
	}

	go c.ProcessQueue("queue1", f)
	go c.ProcessQueue("queue2", f)
	go c.ProcessQueue("queue3", f)
	go c.ProcessQueue("queue4", f)
	go c.ProcessQueue("queue5", f)
	go c.ProcessQueue("queue6", f)
	c.SendMessage("change", "key1", "Hello World!")
	c.SendMessage("change", "key2", "Hello!")
	c.SendMessage("change", "key3", "World!")
	c.SendMessage("change", "key4", "Hi!")
	c.SendMessage("change", "key5", "Again!")

	// Be nice and clean up a little bit
	err = c.DeleteScheme(data)
	if err != nil {
		panic(fmt.Sprintf("Could not delete scheme: %v",err.Error()))
	}
	c.Close()
}
