package messaging

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/phuslu/log"
)

type MessagingBroker struct {
	Settings *HiveMQSettings
	Client   mqtt.Client
}

func NewMessagingBroker(settings *HiveMQSettings) *MessagingBroker {
	brokerTLS := fmt.Sprintf("tls://%s:%s", settings.Broker, settings.Port)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerTLS)
	opts.SetClientID(settings.ClientId)
	opts.SetUsername(settings.Username)
	opts.SetPassword(settings.Password)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return &MessagingBroker{
		Settings: settings,
		Client:   client,
	}
}

func (hive *MessagingBroker) Subscribe(topic string, callback mqtt.MessageHandler) {
	token := hive.Client.Subscribe(topic, 2, callback)
	token.Wait()
	// Check for errors during subscribe
	if token.Error() != nil {
		log.Error().Msg(fmt.Sprintf(`failed to subscribe topic %s %s`, topic, token.Error().Error()))
		panic(token.Error())
	}
	log.Info().Msg(fmt.Sprintf(`"Subscribed to topic: %s"`, topic))
}
