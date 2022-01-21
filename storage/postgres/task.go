package postgres

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	pb "github.com/muhriddinsalohiddin/udevs/genproto"
)

type taskRepo struct {
	db *sqlx.DB
}

// NewTaskRepo
func NewTaskRepo(db *sqlx.DB) *taskRepo {
	return &taskRepo{db: db}
}

func (t *taskRepo) CreateTask(in pb.Task) (pb.Task, error) {
	err := t.db.QueryRow(`
		INSERT INTO tasks (
			id,
			name,
			status,
			priority,
			created_by,
			created_at,
			due_date)
		VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		in.Id,
		in.Name,
		in.Status,
		in.Priority,
		in.CreatedBy,
		time.Now().UTC(),
		in.DueDate,
	).Err()
	if err != nil {
		return pb.Task{}, err
	}
	task, err := t.GetTask(pb.ByIdReq{Id: in.Id})
	if err != nil {
		return pb.Task{}, err
	}
	return task, nil
}
func (t *taskRepo) GetTask(in pb.ByIdReq) (pb.Task, error) {
	var task pb.Task
	err := t.db.QueryRow(`
		SELECT id,
			name,
			status,
			priority,
			created_by,
			created_at,
			due_date 
		FROM tasks
		WHERE id = $1`,
		in.Id,
	).Scan(
		&task.Id,
		&task.Name,
		&task.Status,
		&task.Priority,
		&task.CreatedBy,
		&task.CreatedAt,
		&task.DueDate,
	)
	if err != nil {
		return pb.Task{}, err
	}
	return task, nil
}
func (t *taskRepo) ListTask(in pb.ListReq) (pb.ListRespTask, error) {
	offset := (in.Page - 1) * in.Limit
	var list pb.ListRespTask
	rows, err := t.db.Query(`
		SELECT id,
			name,
			status,
			priority,
			created_by,
			created_at,
			due_date 
		FROM tasks
		LIMIT $1
		OFFSET $2`,
		in.Limit,
		offset,
	)
	if err != nil {
		return pb.ListRespTask{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var task pb.Task
		rows.Scan(
			&task.Id,
			&task.Name,
			&task.Status,
			&task.Priority,
			&task.CreatedBy,
			&task.CreatedAt,
			&task.DueDate,
		)
		list.Tasks = append(list.Tasks, &task)
	}
	err = t.db.QueryRow(`
		SELECT COUNT(*)
		FROM tasks
	`).Scan(&list.Count)
	if err != nil {
		return pb.ListRespTask{}, err
	}
	return list, nil
}
func (t *taskRepo) UpdateTask(in pb.Task) (pb.Task, error) {
	result, err := t.db.Exec(`
		UPDATE tasks
		SET name = $1,
			status = $2,
			priority = $3,
			created_by = $4,
			created_at = $5,
			due_date = $6
		WHERE id = $7`,
		in.Name,
		in.Status,
		in.Priority,
		in.CreatedBy,
		in.CreatedAt,
		in.DueDate,
		in.Id,
	)
	if err != nil {
		return pb.Task{}, err
	}
	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Task{}, sql.ErrNoRows
	}
	task, err := t.GetTask(pb.ByIdReq{Id: in.Id})
	if err != nil {
		return pb.Task{}, err
	}
	return task, nil
}
func (t *taskRepo) DeleteTask(in pb.ByIdReq) error {
	result, err := t.db.Exec(`
		DELETE FROM tasks
		WHERE id = $1`, in.Id,
	)
	if err != nil {
		return err
	}
	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}
	return nil
}
