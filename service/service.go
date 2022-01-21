package service

import (
	"github.com/muhriddinsalohiddin/udevs/pkg/logger"
	"github.com/muhriddinsalohiddin/udevs/storage"
)

type ProtoService struct {
	storage storage.IStorage
	log logger.Logger
}

func NewProtoService(storage storage.IStorage,log logger.Logger) *ProtoService {
	return &ProtoService{
		storage: storage,
		log: log,
	}
}
