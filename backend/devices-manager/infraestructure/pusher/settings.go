package pusher

type PusherSettings struct {
	AppID   string `envconfig:"PUSHER_APP_ID" required:"true"`
	Key     string `envconfig:"PUSHER_KEY" required:"true"`
	Secret  string `envconfig:"PUSHER_SECRET" required:"true"`
	Cluster string `envconfig:"PUSHER_CLUSTER" required:"true"`
	Secure  bool   `envconfig:"SECURE" default:"true"`
}
