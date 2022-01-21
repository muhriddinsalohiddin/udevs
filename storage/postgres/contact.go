package postgres

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	pb "github.com/muhriddinsalohiddin/udevs/genproto"
)

type contactRepo struct {
	db *sqlx.DB
}

// NewTaskRepo
func NewContactRepo(db *sqlx.DB) *contactRepo {
	return &contactRepo{db: db}
}

func (c *contactRepo) CreateContact(in pb.Contact) (pb.Contact, error) {
	err := c.db.QueryRow(`
		INSERT INTO contacts (id,first_name,last_name,phone,email,position)
		VALUES ($1,$2,$3,$4,$5,$6)`,
		in.Id,
		in.FirstName,
		in.LastName,
		in.Phone,
		in.Email,
		in.Position,
	).Err()
	if err != nil {
		return pb.Contact{}, err
	}
	contact, err := c.GetContact(pb.ByIdReq{Id: in.Id})
	if err != nil {
		return pb.Contact{}, err
	}
	return contact, nil
}
func (c *contactRepo) GetContact(in pb.ByIdReq) (pb.Contact, error) {
	var contact pb.Contact
	err := c.db.QueryRow(`
		SELECT
			id,
			first_name,
			last_name,
			phone,
			email,
			position
		FROM contacts
		WHERE id = $1`, in.Id,
	).Scan(
		&contact.Id,
		&contact.FirstName,
		&contact.LastName,
		&contact.Phone,
		&contact.Email,
		&contact.Position,
	)
	if err != nil {
		return pb.Contact{}, err
	}
	return contact, nil
}
func (c *contactRepo) ListContact(in pb.ListReq) (pb.ListRespContact, error) {
	var list pb.ListRespContact
	offset := (in.Page - 1) * in.Limit
	rows, err := c.db.Query(`
		SELECT
			id,
			first_name,
			last_name,
			phone,
			email,
			position
		FROM contacts
		LIMIT $1
		OFFSET $2`,
		in.Limit,
		offset,
	)
	if err != nil {
		return pb.ListRespContact{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var contact pb.Contact
		rows.Scan(
			&contact.Id,
			&contact.FirstName,
			&contact.LastName,
			&contact.Phone,
			&contact.Email,
			&contact.Position,
		)
		list.Contacts = append(list.Contacts, &contact)
	}
	err = c.db.QueryRow(`
		SELECT COUNT(*)
		FROM contacts
	`).Scan(&list.Count)
	if err != nil {
		return pb.ListRespContact{}, err
	}
	return list, nil
}
func (c *contactRepo) UpdateContact(in pb.Contact) (pb.Contact, error) {
	result, err := c.db.Exec(`
		UPDATE contacts
		SET first_name = $1,
			last_name = $2,
			phone=$3,
			email=$4,
			position = $5
		WHERE id = $6`,
		in.FirstName,
		in.LastName,
		in.Phone,
		in.Email,
		in.Position,
		in.Id,
	)
	if err != nil {
		return pb.Contact{}, err
	}
	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Contact{}, sql.ErrNoRows
	}
	contact, err := c.GetContact(pb.ByIdReq{Id: in.Id})
	if err != nil {
		return pb.Contact{}, err
	}
	return contact, nil
}
func (c *contactRepo) DeleteContact(in pb.ByIdReq) error {
	result, err := c.db.Exec(`
		DELETE FROM contacts
		WHERE id = $1`,
		in.Id,
	)
	if err != nil {
		return err
	}
	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}
	return nil
}
