package app

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/LuisDiazM/backend/telemetry-gateway/domain/devices/entities"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/phuslu/log"
)

const (
	topicDevices = "devices/data"
)

func (appl *Application) SaveDevicesData(ctx context.Context) {
	var manage mqtt.MessageHandler = func(c mqtt.Client, m mqtt.Message) {
		t1 := time.Now()
		var requestData entities.DeviceDataRequest
		err := json.Unmarshal(m.Payload(), &requestData)
		if err != nil {
			log.Error().Msg(err.Error())
		}

		appl.DeviceUsecase.SaveDeviceRegistries(ctx, requestData)
		t2 := time.Since(t1)
		log.Info().Msg(fmt.Sprintf(`%s took %d ms`, string(m.Payload()), t2.Milliseconds()))
	}
	appl.MessagingBroker.Subscribe(topicDevices, manage)
}
