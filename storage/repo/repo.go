package repo

import (
	pb "github.com/muhriddinsalohiddin/udevs/genproto"
)

// ContactStorageI
type ContactStorageI interface {
	// CRUD for Contact
	CreateContact(pb.Contact) (pb.Contact, error)
	GetContact(pb.ByIdReq) (pb.Contact, error)
	ListContact(pb.ListReq) (pb.ListRespContact, error)
	UpdateContact(pb.Contact) (pb.Contact, error)
	DeleteContact(pb.ByIdReq) error
	
}
// TaskStorageI
type TaskStorageI interface {
	CreateTask(pb.Task) (pb.Task, error)
	GetTask(pb.ByIdReq) (pb.Task, error)
	ListTask(pb.ListReq) (pb.ListRespTask, error)
	UpdateTask(pb.Task) (pb.Task, error)
	DeleteTask(pb.ByIdReq) error
}