package jazz

import(
	"fmt"
	"gopkg.in/yaml.v2"
	"github.com/streadway/amqp"
)

type Connection struct {
	c *amqp.Connection
}

func Connect(dsn string) (*Connection, error) {
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, err
	}
	return &Connection{conn}, nil
}

func (c *Connection) CreateScheme(data []byte) (error) {
	ch, err := c.c.Channel()
	if err != nil {
		fmt.Println("jsem tu!")
		return err
	}

	s := Settings{}
	err = yaml.Unmarshal(data, &s)
	if err != nil {
		fmt.Println("jsem tady!")
		return err
	}

	// Create exchanges according to settings
	for name, e := range s.Exchanges {
		err = ch.ExchangeDeclarePassive(name, e.Type, e.Durable, e.Autodelete, e.Internal, e.Nowait, nil)
		if err != nil {
			ch, err = c.c.Channel()
			if err != nil {
				return err
			}

			err = ch.ExchangeDeclare(name, e.Type, e.Durable, e.Autodelete, e.Internal, e.Nowait, nil)
			if err != nil {
				return err
			}
		}
	}

	// Create queues according to settings
	for name, q := range s.Queues {
		_, err := ch.QueueDeclarePassive(name, q.Durable, q.Autodelete, q.Exclusive, q.Nowait, nil)
		if err != nil {
			ch, err = c.c.Channel()
			if err != nil {
				return err
			}

			_, err := ch.QueueDeclare(name, q.Durable, q.Autodelete, q.Exclusive, q.Nowait, nil)
			if err != nil {
				return err
			}
		}
	}

	// Create bindings now that everything is setup
	for name, e := range s.Exchanges {
		for _, b := range e.Bindings {
			fmt.Println(b.Exchange, b.Key)
			err = ch.ExchangeBind(name, b.Key, b.Exchange, b.Nowait, nil)
			if err != nil {
				return err
			}
		}
	}

	for name, q := range s.Queues {
		for _, b := range q.Bindings {
			err = ch.QueueBind(name, b.Key, b.Exchange, b.Nowait, nil)
		}
	}

	ch.Close()
	return nil
}

func (c *Connection) DeleteScheme(data []byte) (error) {
	ch, err := c.c.Channel()
	if err != nil {
		return err
	}

	s := Settings{}
	err = yaml.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	for name, _ := range s.Exchanges {
		err = ch.ExchangeDelete(name, false, false)
		if err != nil {
			return err
		}
	}

	for name, _ := range s.Queues {
		_, err = ch.QueueDelete(name, false, false, false)
		if err != nil {
			return err
		}
	}
	ch.Close()
	return nil
}

func (c *Connection) Close() (error) {
	return c.c.Close()
}

func (c *Connection) SendMessage(ex, key, msg string) (error) {
	ch, err := c.c.Channel()
	if err != nil {
		return err
	}

	err = ch.Publish(ex, key, false, false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(msg),
		})
	if err != nil {
		return err
	}
	return ch.Close()
}

func (c *Connection) SendBlob(ex, key string, msg []byte) (error) {
	ch, err := c.c.Channel()
	if err != nil {
		return err
	}

	err = ch.Publish(ex, key, false, false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/octet-stream",
			Body:         msg,
		})
	if err != nil {
		return err
	}
	return ch.Close()
}

func (c *Connection) ProcessQueue(name string, f func([]byte)) error {
	ch, err := c.c.Channel()
	if err != nil {
		return err
	}
	msgs, err := ch.Consume(name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}
	for d := range msgs {
		f(d.Body)
	}
	return nil
}
