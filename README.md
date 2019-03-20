# Jazz

Abstraction layer for quick and simple rabbitMQ connection, messaging and administration. Inspired by Jazz Jackrabbit and his eternal hatred towards slow turtles.

<p align="center">
    <img src="https://upload.wikimedia.org/wikipedia/en/b/b4/Jazz_Jackrabbit.jpg" alt="Jazz Jackrabbit">
</p>


## Usage

This library contains three major parts - exchange/queue scheme creation, publishing of messages and consuming of messages. The greatest benefit of this partitioning is that each part might be in separate application. Also due to dedicated administration part, publishing and consuming of messages is simplified to great extent.

### Step 1: Connect to rabbit

```golang
import(
	"github.com/socifi/jazz"
)

var dsn = "amqp://guest:guest@localhost:5672/"

func main() {
	// ...

	c, err := jazz.Connect(dsn)
	if err != nil {
		t.Errorf("Could not connect to RabbitMQ: %v", err.Error())
		return
	}

	//...
}
```

### Step 2: Create scheme

Scheme specification is done via structure `Settings` which can be easily specified in YAML. So generally you need to decode YAML and then create all queues and exchanges

It can be something really crazy like this!

```golang
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
	// ...

	reader := bytes.NewReader(data)
	scheme, err := DecodeYaml(reader)
	if err != nil {
		t.Errorf("Could not read YAML: %v", err.Error())
		return
	}

	err = c.CreateScheme(scheme)
	if err != nil {
		t.Errorf("Could not create scheme: %v", err.Error())
		return
	}

	//...

	// Be nice and delete scheme (Not advisable in ).
	err = c.DeleteScheme(scheme)
	if err != nil {
		t.Errorf("Could not delete scheme: %v", err.Error())
		return
	}
}
```

### Step 3: Publish and/or consume messages

You can process each queue in separate application or everything together like this:

```golang
func main() {
	// ...

	f := func(msg []byte) {
		fmt.Println(string(msg))
	}

	go c.ProcessQueue("queue1", f)
	go c.ProcessQueue("queue2", f)
	go c.ProcessQueue("queue3", f)
	go c.ProcessQueue("queue4", f)
	go c.ProcessQueue("queue5", f)
	go c.ProcessQueue("queue6", f)
	c.SendMessage("exchange0", "key1", "Hello World!")
	c.SendMessage("exchange0", "key2", "Hello!")
	c.SendMessage("exchange0", "key3", "World!")
	c.SendMessage("exchange0", "key4", "Hi!")
	c.SendMessage("exchange0", "key5", "Again!")

	//...
}
```

## Notes

<sub><sup>No copyright infringement intended. The name Jazz Jackrabbit and artwork of Jazz Jackrabbit is intelectual property of Epic MegaGames and was taken over from [wikipedia](https://en.wikipedia.org/wiki/File:Jazz_Jackrabbit.jpg)</sup></sub>
