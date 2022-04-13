package client

import (
	"context"
	"log"

	pb "github.com/muhriddinsalohiddin/udevs/bot/genproto"
	"google.golang.org/grpc"
)

var Conn *grpc.ClientConn

func Stub(message, priority string) (status int) {
	stub := pb.NewBotClient(Conn)
	_, err := stub.Sender(
		context.Background(),
		&pb.Content{Text: message, Priority: priority},
	)
	if err != nil {
		log.Printf("Problem with sending request to server: %v", err)
		return 400
	}
	return 200
}
