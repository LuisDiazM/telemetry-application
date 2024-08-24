package mongodb

type MongoSettings struct {
	Url         string `envconfig:"MONGO_URL" required:"true"`
	MaxPoolSize uint64 `envconfig:"MONGO_MAXPOOLSIZE" default:"20"`
}
