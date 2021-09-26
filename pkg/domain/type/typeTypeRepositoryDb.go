package _type

import (
	"database/sql"

	"github.com/tamiat/backend/pkg/errs"
)

type TypeRepositoryDb struct {
	db *sql.DB
}

func (r TypeRepositoryDb) Create(newType Type) (int, error) {
	row := r.db.QueryRow(`INSERT INTO types (name) VALUES ($1) RETURNING id`, newType.Name)
	var id int
	err := row.Scan(&id)
	return id, err
}

func (r TypeRepositoryDb) Read() ([]Type, error) {
	Types := make([]Type, 0)
	rows, err := r.db.Query("SELECT id, name FROM types")
	if err != nil {
		return Types, err
	}
	for rows.Next() {
		var _type Type
		switch err := rows.Scan(&_type.ID, &_type.Name); err {
		case sql.ErrNoRows:
			return Types, sql.ErrNoRows
		case nil:
			Types = append(Types, _type)
		default:
			return Types, errs.ErrDb
		}
	}
	return Types, nil
}

func (r TypeRepositoryDb) Update(_type Type, id string) error {
	var name string
	row := r.db.QueryRow("SELECT name FROM types WHERE id= $1", id)
	err := row.Scan(&name)
	if err != nil {
		return err
	}
	_, err = r.db.Query("UPDATE types SET name= $1 WHERE id=$2", _type.Name, id)
	if err != nil {
		return err
	}
	return nil
}

func (r TypeRepositoryDb) Delete(id string) error {
	var name string
	row := r.db.QueryRow("SELECT name FROM types WHERE id=$1", id)
	err := row.Scan(&name)
	if err != nil {
		return err
	}
	_, err = r.db.Query("DELETE FROM types WHERE id= $1", id)
	if err != nil {
		return err
	}
	return nil
}



func NewTypeRepositoryDb(db *sql.DB) TypeRepositoryDb {
	return TypeRepositoryDb{db}
}
