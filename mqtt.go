package main

import (
	"fmt"
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"os"
)

// define a function for the default message handler
var f MQTT.MessageHandler = func(client *MQTT.MqttClient, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
	h.broadcast <- msg.Payload()
}

func initMqtt() {

	opts := MQTT.NewClientOptions().AddBroker("tcp://qchp0k.messaging.internetofthings.ibmcloud.com:1883")
	opts.SetUsername("a-qchp0k-s4ywm4ruua")
	opts.SetPassword("0MaeP8nAIRfYiT!-ub")
	opts.SetClientId("a:qchp0k:visualizer")

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	_, err := c.Start()
	if err != nil {
		panic(err)
	}

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
