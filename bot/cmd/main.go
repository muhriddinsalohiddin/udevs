package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/muhriddinsalohiddin/udevs/bot/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Problem with loading env: %v", err)
	}

	port := os.Getenv("HTTP_PORT")
	r := api.Option()
	log.Println("Listening and serving HTTP on localhost:", port)
	err = r.Run(port)
	if err != nil {
		log.Fatalf("Problem with connecting gateway to port: %v", err)
	}
}
