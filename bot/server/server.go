package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/reflection"

	"github.com/muhriddinsalohiddin/udevs/bot/bot"
	"github.com/muhriddinsalohiddin/udevs/bot/config"
	pb "github.com/muhriddinsalohiddin/udevs/bot/genproto"
	"github.com/muhriddinsalohiddin/udevs/bot/handlers"
	"google.golang.org/grpc"
)

var List []handlers.Message

type Server struct {
	pb.UnimplementedBotServer
}

func (*Server) Sender(ctx context.Context, input *pb.Content) (*empty.Empty, error) {

	var message handlers.Message

	message.Priority = input.Priority
	message.Text = input.Text

	List = append(List, message)

	return &empty.Empty{}, nil
}

func main() {
	go CuncurrentSending()
	conf := config.Load()
	lis, err := net.Listen("tcp", conf.RpcPort)
	if err != nil {
		log.Fatalf("Problem with connecting to tcp: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBotServer(s, &Server{})
	reflection.Register(s)

	log.Println("Server is running in port :" + conf.RpcPort)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Problem with connecting to GRPC Server: %v", err)
	}
}

func CuncurrentSending() {
	var isSend bool
	for {
		isSend = false
		for i, message := range List {
			if message.Priority == "high" {
				err := bot.BotService(List[i].Text)
				if err != nil {
					log.Fatalf("Problem with sending message to bot: %v", err)
				} else {
					isSend = true
					Remove(i)
					time.Sleep(time.Second * 5)

					break
				}
			}
		}

		if isSend {
			continue
		}

		for i, message := range List {
			if message.Priority == "medium" {
				err := bot.BotService(List[i].Text)
				if err != nil {
					log.Fatalf("Problem with sending message to bot: %v", err)
				} else {
					isSend = true
					Remove(i)
					time.Sleep(time.Second * 5)

					break
				}
			}
		}
		if isSend {
			continue
		}
		for i, message := range List {
			if message.Priority == "low" {
				err := bot.BotService(List[i].Text)
				if err != nil {
					log.Fatalf("Problem with sending message to bot: %v", err)
				} else {
					isSend = true
					Remove(i)
					time.Sleep(time.Second * 5)
					break
				}
			}
		}

	}
}

func Remove(i int) {
	List = append(List[:i], List[i+1:]...)
}
