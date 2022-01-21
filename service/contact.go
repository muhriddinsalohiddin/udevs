package service

import (
	"context"

	"github.com/gofrs/uuid"
	pb "github.com/muhriddinsalohiddin/udevs/genproto"
	"google.golang.org/grpc/codes"
	"github.com/muhriddinsalohiddin/udevs/pkg/logger"
	"google.golang.org/grpc/status"
)

func (p *ProtoService) CreateContact(ctx context.Context, in *pb.Contact) (*pb.Contact, error) {
	id, err := uuid.NewV4()
	if err != nil {
		p.log.Error("Failed to generated uuid", logger.Error(err))
		return nil, status.Error(codes.Internal, "Failed to generate uuid")
	}
	in.Id = id.String()
	contact, err := p.storage.Contact().CreateContact(*in)
	if err != nil {
		p.log.Error("Failed to create contact", logger.Error(err))
		return nil, status.Error(codes.Internal, "Failed to create contact")
	}
	return &contact, nil
}
func (p *ProtoService) GetContact(ctx context.Context, in *pb.ByIdReq) (*pb.Contact, error) {
	contact, err := p.storage.Contact().GetContact(*in)
	if err != nil {
		p.log.Error("Failed to get contact", logger.Error(err))
		return nil, status.Error(codes.Internal, "Failed to get contact")
	}
	return &contact, nil
}
func (p *ProtoService) ListContact(ctx context.Context, in *pb.ListReq) (*pb.ListRespContact, error) {
	contacts, err := p.storage.Contact().ListContact(*in)
	if err != nil {
		p.log.Error("Failed to get list contact", logger.Error(err))
		return nil, status.Error(codes.Internal, "Failed to get list contact")
	}
	return &contacts, nil
}
func (p *ProtoService) UpdateContact(ctx context.Context, in *pb.Contact) (*pb.Contact, error) {
	contact, err := p.storage.Contact().UpdateContact(*in)
	if err != nil {
		p.log.Error("Failed to update contact", logger.Error(err))
		return nil, status.Error(codes.Internal, "Failed to update contact")
	}
	return &contact, nil
}
func (p *ProtoService) DeleteContact(ctx context.Context, in *pb.ByIdReq) (*pb.EmptyResp, error) {
	err := p.storage.Contact().DeleteContact(*in)
	if err != nil {
		p.log.Error("Failed to delete contact", logger.Error(err))
		return nil, status.Error(codes.Internal, "Failed to delete contact")
	}
	return &pb.EmptyResp{}, nil
}
