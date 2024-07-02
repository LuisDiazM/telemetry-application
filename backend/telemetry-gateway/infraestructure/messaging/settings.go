package messaging

type HiveMQSettings struct {
	Broker   string `envconfig:"HIVE_MQ_URL" required:"true"`
	Port     string `envconfig:"HIVE_MQ_PORT" required:"true"`
	Username string `envconfig:"HIVE_MQ_USER" required:"true"`
	Password string `envconfig:"HIVE_MQ_PASS" required:"true"`
	ClientId string `envconfig:"HIVE_CLIENT_ID" default:"x5pcFRxXMVjeBFv"`
}
