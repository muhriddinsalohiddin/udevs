package service

import "github.com/muhriddinsalohiddin/udevs/storage"

type ProtoService struct {
	storage storage.IStorage
}

func NewProtoService(storage storage.IStorage) *ProtoService {
	return &ProtoService{
		storage: storage,
	}
}
