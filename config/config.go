package config

import (
	"os"
	"strconv"
)

// Config struct para configuración general
type Config struct {
	MaxWorkers int
	MaxQueue   int
	Port       string
}

// ConfigMinio struct para configuración de Minio
type ConfigMinio struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

// ConfigRabbitMq struct para configuración de RabbitMQ
type ConfigRabbitMq struct {
	User         string
	Pass         string
	Host         string
	Port         string
	NameQueueEva string
}

// Load carga la configuración general desde variables de entorno
func Load() Config {
	maxWorkers, _ := strconv.Atoi(getEnv("MAX_WORKERS", "4"))
	maxQueue, _ := strconv.Atoi(getEnv("MAX_QUEUE", "20"))
	port := getEnv("PORT", ":8081")

	return Config{
		MaxWorkers: maxWorkers,
		MaxQueue:   maxQueue,
		Port:       port,
	}
}

// LoadConfMinio carga la configuración de Minio desde variables de entorno
func LoadConfMinio() ConfigMinio {
	endpoint := getEnv("MINIO_ENDPOINT", "host.docker.internal:9000")
	accessKeyID := getEnv("MINIO_ACCESS_KEY", "lo_del_secrto")
	secretAccessKey := getEnv("MINIO_SECRET_KEY", "esto_debe_ser_un_secreto")
	useSSL, _ := strconv.ParseBool(getEnv("MINIO_USE_SSL", "false"))

	return ConfigMinio{
		Endpoint:        endpoint,
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
		UseSSL:          useSSL,
	}
}

// LoadConfRabbitMq carga la configuración de RabbitMQ desde variables de entorno
func LoadConfRabbitMq() ConfigRabbitMq {
	user := getEnv("RABBITMQ_USER", "admin")
	pass := getEnv("RABBITMQ_PASS", "admin")
	host := getEnv("RABBITMQ_HOST", "host.docker.internal")
	port := getEnv("RABBITMQ_PORT", "5673")
	nameQueueEva := getEnv("RABBITMQ_QUEUE_NAME_EVA", "notification")

	return ConfigRabbitMq{
		User:         user,
		Pass:         pass,
		Host:         host,
		Port:         port,
		NameQueueEva: nameQueueEva,
	}
}

// getEnv obtiene el valor de una variable de entorno o un valor por defecto si no está definida
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
