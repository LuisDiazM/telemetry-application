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
	TOPIC_DEVICES  = "devices"
	PUSHER_CHANNEL = "devices"
	PUSHER_EVENT   = "status"
)

type PusherStatusData struct {
	DeviceId string `json:"device_id,omitempty"`
	Status   bool   `json:"status,omitempty"`
}

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
			var data PusherStatusData = PusherStatusData{DeviceId: requestData.DeviceId, Status: true}
			appl.Pusher.PublishEvent(PUSHER_CHANNEL, PUSHER_EVENT, data)
		}
		t2 := time.Since(t1)
		log.Info().Msg(fmt.Sprintf(`%s took %d ms`, string(m.Payload()), t2.Milliseconds()))
	}
	appl.MessagingBroker.Subscribe(TOPIC_DEVICES, manage)
}
