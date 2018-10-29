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

Scheme specification is done via YAML string.

It can be something really crazy like this!

```golang
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
	// ...

	err = c.CreateScheme(data)
	if err != nil {
		t.Errorf("Could not create scheme: %v", err.Error())
		return
	}

	//...
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
	c.SendMessage("change", "key1", "Hello World!")
	c.SendMessage("change", "key2", "Hello!")
	c.SendMessage("change", "key3", "World!")
	c.SendMessage("change", "key4", "Hi!")
	c.SendMessage("change", "key5", "Again!")

	//...
}
```

## Notes

<sub><sup>No copyright infringement intended. The name Jazz Jackrabbit and artwork of Jazz Jackrabbit is intelectual property of Epic MegaGames and was taken over from [wikipedia](https://en.wikipedia.org/wiki/File:Jazz_Jackrabbit.jpg)</sup></sub>
