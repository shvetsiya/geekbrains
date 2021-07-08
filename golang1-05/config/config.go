package config

type Config struct {
	Port        uint16 `yaml:"port,omitempty"`
	DbURL       string `yaml:"db_url,omitempty"`
	JaegerURL   string `yaml:"jaeger_url,omitempty"`
	SentryURL   string `yaml:"sentry_url,omitempty"`
	KafkaBroker string `yaml:"kafka_broker,omitempty"`
	SomeAppID   string `yaml:"some_app_id,omitempty"`
	SomeAppKey  string `yaml:"some_app_key,omitempty"`
}

func (conf *Config) IsValid() bool {
	v := new(Validator)
	v.IsPortValid(conf.Port, 0, 65535)
	v.IsDbURLValid(conf.DbURL, "postgres://db-user:db-password@petstore-db")
	v.IsURLValid(conf.JaegerURL, "jaeger")
	v.IsURLValid(conf.SentryURL, "sentry")
	v.IsKafkaBrokerValid(conf.KafkaBroker, "kafka")
	return v.IsValid()
}
