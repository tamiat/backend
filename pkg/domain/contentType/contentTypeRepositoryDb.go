package contentType

import (
	"database/sql"
	"errors"
	"fmt"
)

type ContentTypeRepositoryDb struct {
	db *sql.DB
}

func (r ContentTypeRepositoryDb) Create(n string, cols string) (string, error) {
	var query = "INSERT INTO contentType (name) VALUES ('" + n + "')"
	_, err := r.db.Exec(query)
	if err != nil {
		return "", errors.New("Unexpected database error")
	}
	query = "CREATE TABLE " + n + " ( " + cols + " )"
	_, err = r.db.Exec(query)
	if err != nil {
		return "", errors.New("Unexpected database error")
	}
	query = `SELECT currval(pg_get_serial_sequence('contentType','id'));`
	row := r.db.QueryRow(query)
	var id string
	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		return "", errors.New("Unexpected database error")
	case nil:
		return id, nil
	default:
		fmt.Println(id)
		panic(err)
	}
}

func (r ContentTypeRepositoryDb) DeleteById(id string) error {
	var query = "SELECT name FROM contentType WHERE id=" + id
	fmt.Println(query)
	row := r.db.QueryRow(query)
	var name string
	err := row.Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("content type not found")
		} else {
			return errors.New("Unexpected database error")
		}
	}
	query = "DROP TABLE " + name
	_, err = r.db.Exec(query)
	if err != nil {
		return errors.New("Unexpected database error")
	}
	query = "DELETE FROM contentType" + " WHERE id=" + id
	_, err = r.db.Exec(query)
	if err != nil {
		return errors.New("Unexpected database error")
	}
	return nil
}

func NewContentTypeRepositoryDb(db *sql.DB) ContentTypeRepositoryDb {
	return ContentTypeRepositoryDb{db}
}
