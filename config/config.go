package config

type Config struct {
	MaxWorkers int
	MaxQueue   int
	Port       string
}

type ConfigMinio struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

type ConfigRabbitMq struct {
	User         string
	Pass         string
	Host         string
	Port         string
	NameQueueEva string
}

func Load() Config {
	return Config{
		MaxWorkers: 4,
		MaxQueue:   20,
		Port:       ":8081",
	}
}

func LoadConfMinio() ConfigMinio {
	return ConfigMinio{
		Endpoint:        "host.docker.internal:9000",
		AccessKeyID:     "wDeP32IOLrUwrlmFLwOc",
		SecretAccessKey: "VzwSovPOwHeAe68asfwzBMjgCpbNzLvCXiUpNpCi",
		UseSSL:          false,
	}
}

func LoadConfRabbitMq() ConfigRabbitMq {
	return ConfigRabbitMq{
		User:         "admin",
		Pass:         "admin",
		Host:         "host.docker.internal",
		Port:         "5673",
		NameQueueEva: "notification",
	}
}
