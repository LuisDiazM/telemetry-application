package app

import (
	"context"
	"encoding/json"
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
		type deviceRequest struct {
			Status   bool   `json:"status,omitempty"`
			DeviceId string `json:"device_id,omitempty"`
		}
		var requestData deviceRequest
		err := json.Unmarshal(m.Payload(), &requestData)
		if err != nil {
			log.Error().Msg(err.Error())
		}
		if err == nil {
			appl.DevicesUsecase.UpdateDeviceStatus(ctx, requestData.DeviceId)
		}
		t2 := time.Since(t1)
		log.Info().Msg(fmt.Sprintf(`%s took %d ms`, string(m.Payload()), t2.Milliseconds()))
	}
	appl.MessagingBroker.Subscribe(topicDevices, manage)
}
