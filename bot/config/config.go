package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BotTonken string
	ChatId    string
	Port      string
	RpcHost   string
	RpcPort   string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Problem with loading env: %v", err)
	}
	var c Config
	c.BotTonken = getOrReturnDefault("TELEGRAM_APITOKEN", "5194091320:AAGD6ChEhuOVN2UTIPIRRDzYxlOQbUiwwxI")
	c.ChatId = getOrReturnDefault("CHAT_ID", "334116637")
	c.Port = getOrReturnDefault("HTTP_PORT", ":8080")
	c.RpcPort = getOrReturnDefault("RPC_PORT", ":9500")
	c.RpcHost = getOrReturnDefault("RPC_HOST", "localhost")
	return c
}
func getOrReturnDefault(key, defaultValue string) string {

	if _, exists := os.LookupEnv(key); exists {
		return os.Getenv(key)
	}
	return defaultValue
}
