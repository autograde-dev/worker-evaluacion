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

func Load() Config {
	return Config{
		MaxWorkers: 4,
		MaxQueue:   20,
		Port:       ":8081",
	}
}

func LoadConfMinio() ConfigMinio {
	return ConfigMinio{
		Endpoint:        "localhost:9000",
		AccessKeyID:     "wDeP32IOLrUwrlmFLwOc",
		SecretAccessKey: "VzwSovPOwHeAe68asfwzBMjgCpbNzLvCXiUpNpCi",
		UseSSL:          false,
	}
}
