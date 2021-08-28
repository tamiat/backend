package contentType

import (
	"database/sql"
	"errors"
	"fmt"
)

type ContentTypeRepositoryDb struct {
	db *sql.DB
}

func (r ContentTypeRepositoryDb) isTableExists(id string) (string, error) {
	var query = "SELECT name FROM contentType WHERE id=" + id
	fmt.Println(query)
	row := r.db.QueryRow(query)
	var name string
	err := row.Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("content type not found")
		} else {
			return "", errors.New("Unexpected database error")
		}
	}
	return name, nil
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
	name, err := r.isTableExists(id)
	if err != nil {
		return err
	}
	query := "DROP TABLE " + name
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

func (r ContentTypeRepositoryDb) UpdateColName(id string, oldName string, newName string) error {
	name, err := r.isTableExists(id)
	if err != nil {
		return err
	}
	query := "SELECT COUNT(*) FROM information_schema.columns WHERE table_name= '" + name + "' and column_name='" + oldName + "'"
	row := r.db.QueryRow(query)
	var colName int
	err = row.Scan(&colName)
	if err != nil {
		return errors.New("Unexpected database error")
	}
	if colName == 0 {
		return errors.New("column not found")
	}
	query = "ALTER TABLE " + name + " RENAME COLUMN " + oldName + " TO " + newName
	_, err = r.db.Exec(query)
	if err != nil {
		return errors.New("Unexpected database error")
	}
	return nil
}

func (r ContentTypeRepositoryDb) AddCol(id string, col string) error {
	name, err := r.isTableExists(id)
	if err != nil {
		return err
	}
	query := "ALTER TABLE " + name + " ADD COLUMN " + col
	fmt.Println(query)
	_, err = r.db.Exec(query)
	if err != nil {
		return errors.New("Unexpected database error")
	}
	return nil
}

func NewContentTypeRepositoryDb(db *sql.DB) ContentTypeRepositoryDb {
	return ContentTypeRepositoryDb{db}
}
