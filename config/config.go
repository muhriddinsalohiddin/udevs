package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost      string
	PostgresPort      int
	PostgresDatabase  string
	PostgresUser      string
	PostgresPassword  string
	RPCPort           string
}

func Load() Config {
	var c Config


	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "127.0.0.1"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "task"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "muhriddin"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "1"))


	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":50051"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
