package cache

type RedisSettings struct {
	RedisHost string `envconfig:"REDIS_HOST" default:":6379"`
	RedisPass string `envconfig:"REDIS_PASSWORD" default:""`
}
