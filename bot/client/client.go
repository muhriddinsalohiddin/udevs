package client

import (
	"context"
	"fmt"
	"log"

	"github.com/muhriddinsalohiddin/udevs/bot/config"
	pb "github.com/muhriddinsalohiddin/udevs/bot/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Stub(message, priority string) (status int) {
	conf := config.Load()
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", conf.RpcHost, conf.RpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Problem with connecting client to server: %s", err)
		return 400
	}

	defer conn.Close()

	stub := pb.NewBotClient(conn)

	_, err = stub.Sender(
		context.Background(),
		&pb.Content{Text: message, Priority: priority},
	)

	if err != nil {
		log.Printf("Problem with sending request to server: %v", err)
		return 400
	}
	return 200
}
