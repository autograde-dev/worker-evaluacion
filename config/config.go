package config

type Config struct {
	MaxWorkers int
	MaxQueue   int
	Port       string
}

func Load() Config {
	return Config{
		MaxWorkers: 4,
		MaxQueue:   20,
		Port:       ":8081",
	}
}
