package sub

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Subscribe struct {
	topic string
	c     mqtt.Client
}

func NewSub(topic string, c mqtt.Client) *Subscribe {
	return &Subscribe{
		topic,
		c,
	}
}

func (s *Subscribe) Handler() {

	token := s.c.Subscribe(s.topic, 1, nil)

	token.Wait()

	fmt.Println("Subscribed:", s.topic)

}
