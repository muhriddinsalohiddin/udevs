package storage

import (
	"github.com/jmoiron/sqlx"

	"github.com/muhriddinsalohiddin/udevs/storage/postgres"
	"github.com/muhriddinsalohiddin/udevs/storage/repo"
)

// IStorage ...
type IStorage interface {
	Task() repo.TaskStorageI
	Contact() repo.ContactStorageI
}

type storagePg struct {
	db       *sqlx.DB
	taskRepo repo.TaskStorageI
	contactRepo repo.ContactStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		taskRepo: postgres.NewTaskRepo(db),
		contactRepo: postgres.NewContactRepo(db),
	}
}

func (s storagePg) Task() repo.TaskStorageI {
	return s.taskRepo
}
func (s storagePg) Contact() repo.ContactStorageI {
	return s.contactRepo
}
