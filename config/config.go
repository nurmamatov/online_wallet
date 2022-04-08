package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment string

	PostgresHost     string
	PostgresPassword string
	PostgresUser     string
	PostgresDatabase string
	PostgresPort     int

	HttpPort string
}

func Load() Config {
	cfg := Config{}

	cfg.Environment = cast.ToString(EnvOrDefoult("ENVIROMENT", "develop"))

	cfg.PostgresUser = cast.ToString(EnvOrDefoult("POSTGRES_USER", "khusniddin"))
	cfg.PostgresDatabase = cast.ToString(EnvOrDefoult("POSTGRES_DATABASE", "online_wallet"))
	cfg.PostgresPort = cast.ToInt(EnvOrDefoult("POSTGRES_PORT", 5432))
	cfg.PostgresPassword = cast.ToString(EnvOrDefoult("POSTGRES_PASSWORD", "1234"))
	cfg.PostgresHost = cast.ToString(EnvOrDefoult("POSTGRES_HOST", "localhost"))

	cfg.HttpPort = cast.ToString(EnvOrDefoult("HTTP_PORT", ":8080"))

	return cfg
}

func EnvOrDefoult(key string, value interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return value
}
