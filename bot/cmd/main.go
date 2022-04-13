package main

import (
	"fmt"
	"log"

	"github.com/muhriddinsalohiddin/udevs/bot/api"
	"github.com/muhriddinsalohiddin/udevs/bot/config"
	"github.com/muhriddinsalohiddin/udevs/bot/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conf := config.Load()
	conn, err := grpc.Dial(fmt.Sprintf("%s%s", conf.RpcHost, conf.RpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Problem with connecting client to server: %v", err)
	}
	client.Conn = conn
	r := api.Option()
	log.Printf("Listening and serving HTTP on localhost:%v", conf.Port)

	if err = r.Run(conf.Port); err != nil {
		log.Fatalf("Problem with connecting gateway to port: %v", err)
	}
}
