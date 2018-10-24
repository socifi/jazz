package rabbit

import(
	"testing"
)

func TestConnection(t *testing.T) {
	conn, err := Connect("amqp://guest:guest@localhost:5672/")
	if err != nil {
		t.Errorf("Could not connect to rabbitmq")
		return
	}
	q, err := conn.ConnectToQueue("pokusna", false)
	if err != nil {
		t.Errorf("Could not connect to rabbitmq")
		return
	}
	q.SendMessage("Hello World!")
	q.Close()
	conn.Close()
}
