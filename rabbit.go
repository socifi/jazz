package rabbit

import(
	"github.com/streadway/amqp"
)

type Connection struct {
	c *amqp.Connection
}

type Queue struct {
	c *amqp.Channel
	q amqp.Queue
}

func Connect(dsn string) (*Connection, error) {
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, err
	}
	return &Connection{conn}, nil
}

func (c *Connection) ConnectToQueue(name string, recreate bool) (*Queue, error) {
	ch, err := c.c.Channel()
	if err != nil {
		return nil, err
	}
	q, err := ch.QueueDeclare(
		name,  // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}
	return &Queue{ch, q}, nil
}

func (c *Connection) Close() (error) {
	return c.c.Close()
}

func (q *Queue) Close() (error) {
	return q.c.Close()
}

func (q *Queue) SendMessage(msg string) (error) {
	return q.c.Publish("",
		q.q.Name, // routing key
		false,    // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(msg),
		})
}
