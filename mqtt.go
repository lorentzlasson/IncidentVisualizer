package main

import (
	"encoding/json"
	"fmt"
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"os"
)

type Configuration struct {
	Org      string
	Username string
	Password string
}

var credentials = Configuration{}

// define a function for the default message handler
var f MQTT.MessageHandler = func(client *MQTT.MqttClient, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
	h.broadcast <- msg.Payload()
}

func readCredentials() {
	file, _ := os.Open("iotfcreds.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&credentials)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func initMqtt() {
	readCredentials()
	broker := fmt.Sprintf("tcp://%s.messaging.internetofthings.ibmcloud.com:1883", credentials.Org)
	opts := MQTT.NewClientOptions().AddBroker(broker)
	opts.SetUsername(credentials.Username)
	opts.SetPassword(credentials.Password)
	clientId := fmt.Sprintf("a:%s:%s", credentials.Org, "visualizer")
	opts.SetClientId(clientId)

	fmt.Print("Connecting to mqtt...")
	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	_, err := c.Start()
	if err != nil {
		panic(err)
	}

	fmt.Println(" success!")
	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	topic, err := MQTT.NewTopicFilter("iot-2/type/+/id/+/evt/incident/fmt/json", 0)
	if receipt, err := c.StartSubscription(f, topic); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		<-receipt
	}
}
