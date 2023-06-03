package main

import (
	"fmt"
	send "mqtt-go-lang-example/sender"
	sub "mqtt-go-lang-example/subscribers"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var topic = "topic/test"

func main() {
	broker := "broker.emqx.io"
	port := 1883

	opts := mqtt.NewClientOptions()

	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername("emqx")
	opts.SetPassword("public")

	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {

		fmt.Printf("TOPIC: %s\n", msg.Topic())
		fmt.Printf("MSG: %s\n", string(msg.Payload()))
	})
	opts.OnConnect = func(c mqtt.Client) {
		fmt.Println("Connected")
	}

	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		fmt.Println("Error", err)
	}

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub := sub.NewSub(topic, client)
	pub := send.NewSender(topic, client)

	sub.Handler()
	pub.Handler()

	client.Disconnect(10000)

}

// func pub(c mqtt.Client) {
// 	num := 1000
// 	for i := 0; i < num; i++ {
// 		res, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + fmt.Sprintf("%d", i))
// 		if err != nil {
// 			panic(err)
// 		}

// 		defer res.Body.Close()

// 		bt, err := io.ReadAll(res.Body)

// 		if err != nil {
// 			panic(err)
// 		}

// 		token := c.Publish(topic, 0, false, string(bt))

// 		token.Wait()
// 	}
// }

// func sub(c mqtt.Client) {

// 	token := c.Subscribe(topic, 1, nil)
// 	token.Wait()
// 	fmt.Printf("Subscribing to %s\n", topic)
// }
