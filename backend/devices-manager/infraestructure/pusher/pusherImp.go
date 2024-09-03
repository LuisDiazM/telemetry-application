package pusher

import (
	"log"

	"github.com/pusher/pusher-http-go"
)

type PusherImp struct {
	pusherClient *pusher.Client
}

func NewPusherImp(pusherSettings *PusherSettings) *PusherImp {
	pusherClient := pusher.Client{
		AppID:   pusherSettings.AppID,
		Key:     pusherSettings.Key,
		Secret:  pusherSettings.Secret,
		Cluster: pusherSettings.Cluster,
		Secure:  pusherSettings.Secure,
	}
	return &PusherImp{pusherClient: &pusherClient}
}

func (push *PusherImp) PublishEvent(channel string, event string, data interface{}) {

	err := push.pusherClient.Trigger(channel, event, data)
	if err != nil {
		log.Println(err.Error())
	}
}
