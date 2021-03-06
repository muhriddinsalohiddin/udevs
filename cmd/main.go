package main

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/muhriddinsalohiddin/udevs/service"
	"github.com/muhriddinsalohiddin/udevs/storage"

	pb "github.com/muhriddinsalohiddin/udevs/genproto"
	"github.com/muhriddinsalohiddin/udevs/pkg/db"
	"github.com/muhriddinsalohiddin/udevs/pkg/logger"

	"github.com/muhriddinsalohiddin/udevs/config"
	// "github.com/muhriddinsalohiddin/udevs/syntax"
)

func main() {
	// syntax.Fibonacci(25)
	// syntax.FizzBuzz(3)
	// syntax.Palindrome("radar")
	// syntax.OddEvenSum(256,"even")
	// syntax.Duplicate([]int{1,2,3,5,66,7,32,3})
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "catalog-service")
	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres err", logger.Error(err))
	}
	pgStorage := storage.NewStoragePg(connDB)

	protoService := service.NewProtoService(pgStorage,log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Connection GRPC error", logger.Error(err))
	}
	s := grpc.NewServer()
	pb.RegisterServiceServer(s, protoService)
	reflection.Register(s)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening", logger.Error(err))
	}
}
