package send

import (
	"fmt"
	"io"
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Sender struct {
	c     mqtt.Client
	topic string
}

func NewSender(topic string, c mqtt.Client) *Sender {
	return &Sender{
		c,
		topic,
	}
}

func (s *Sender) Handler() {
	postsSize := 2000

	for i := 0; i <= postsSize; i++ {
		res, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + fmt.Sprintf("%d", i))
		if err != nil {
			panic(err)
		}

		defer res.Body.Close()

		bt, err := io.ReadAll(res.Body)

		if err != nil {
			panic(err)
		}

		token := s.c.Publish(s.topic, 1, false, string(bt))

		token.Wait()

	}

}
