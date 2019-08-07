package socketioemitter

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)

type Connector struct {
	redis *redis.Client
}

type ConnectorConfig struct {
	Host     string
	Password string
	DB       string
	Port     int
}

type IConnector interface {
	Adapter() *redis.Client
	Connect(ConnectorConfig) error
	PushEvent([]byte) error
}

func NewConnector(cc ConnectorConfig) (c *Connector, err error) {
	c = &Connector{}
	if err = c.Connect(cc); err != nil {
		return
	}

	return
}

func (c *Connector) Connect(cc ConnectorConfig) (err error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cc.Host, cc.Port),
		Password: cc.Password,
		DB:       0,
	})

	if _, err = client.Ping().Result(); err != nil {
		return
	}

	c.redis = client
	go c.handleSub()
	return
}

func (c *Connector) Adapter() *redis.Client {
	return c.redis
}

func (c *Connector) PushEvent(payload []byte) error {
	res := c.redis.Publish("channel", payload)
	return res.Err()
}

func (c *Connector) handleSub() {
	pubsub := c.redis.Subscribe("channel")
	defer pubsub.Close()

	for {
		msg, _ := pubsub.ReceiveTimeout(time.Second)
		switch m := msg.(type) {
		case *redis.Subscription:
			log.Printf("Subscription Message: %v to channel '%v'. %v total subscriptions.", m.Kind, m.Channel, m.Count)
			continue
		case *redis.Message:
			fmt.Println(m.Channel, m.Payload)
		}

	}
}
