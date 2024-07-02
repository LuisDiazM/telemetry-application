package app

import (
	"context"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	topicDevices = "devices/data"
)

func (appl *Application) SaveDevicesData(ctx context.Context) {
	var manage mqtt.MessageHandler = func(c mqtt.Client, m mqtt.Message) {
		fmt.Println(m.Topic(), m.Payload())
	}
	appl.MessagingBroker.Subscribe(topicDevices, manage)
}
