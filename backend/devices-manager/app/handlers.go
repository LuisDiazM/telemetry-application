package app

import (
	"context"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/phuslu/log"
)

const (
	topicDevices = "devices/status"
)

func (appl *Application) NotifyStatusWS(ctx context.Context) {
	var manage mqtt.MessageHandler = func(c mqtt.Client, m mqtt.Message) {
		t1 := time.Now()

		t2 := time.Since(t1)
		log.Info().Msg(fmt.Sprintf(`%s took %d ms`, string(m.Payload()), t2.Milliseconds()))
	}
	appl.MessagingBroker.Subscribe(topicDevices, manage)
}
