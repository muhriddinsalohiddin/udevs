package service

import (
	"context"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/muhriddinsalohiddin/udevs/genproto"
	"github.com/muhriddinsalohiddin/udevs/pkg/logger"
)

func (p *ProtoService) CreateTask(ctx context.Context, in *pb.Task) (*pb.Task, error) {
	id, err := uuid.NewV4()
	if err != nil {
		p.log.Error("Failed to generated uuid", logger.Error(err))
		return nil, status.Error(codes.Internal, "Failed to generate uuid")
	}
	in.Id = id.String()
	task, err := p.storage.Task().CreateTask(*in)
	if err != nil {
		p.log.Error("Failed to create task", logger.Error(err))
		return nil, status.Error(codes.Internal, "Failed to create task")
	}
	return &task, nil
}
func (p *ProtoService) GetTask(ctx context.Context, in *pb.ByIdReq) (*pb.Task, error) {
	task, err := p.storage.Task().GetTask(*in)
	if err != nil {
		p.log.Error("Failed to get task", logger.Error(err))
		return nil, status.Error(codes.Internal, "Failed to get task")
	}
	return &task, nil
}
func (p *ProtoService) ListTask(ctx context.Context, in *pb.ListReq) (*pb.ListRespTask, error) {
	list, err := p.storage.Task().ListTask(*in)
	if err != nil {
		p.log.Error("Failed to get list task", logger.Error(err))
		return nil, status.Error(codes.Internal, "Failed to get list task")
	}
	return &list, nil
}
func (p *ProtoService) UpdateTask(ctx context.Context, in *pb.Task) (*pb.Task, error) {
	task, err := p.storage.Task().UpdateTask(*in)
	if err != nil {
		p.log.Error("Failed to update task", logger.Error(err))
		return nil, status.Error(codes.Internal, "Failed to update task")
	}
	return &task, nil
}
func (p *ProtoService) DeleteTask(ctx context.Context, in *pb.ByIdReq) (*pb.EmptyResp, error) {
	err := p.storage.Task().DeleteTask(*in)
	if err != nil {
		p.log.Error("Failed to delete task", logger.Error(err))
		return nil, status.Error(codes.Internal, "Failed to delete task")
	}
	return &pb.EmptyResp{}, nil
}
