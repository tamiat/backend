package contentType

import (
	"database/sql"
	"fmt"

	"github.com/tamiat/backend/pkg/errs"
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
			return "", errs.ErrContentTypeNotFound
		} else {
			return "", errs.ErrDb
		}
	}
	return name, nil
}

func (r ContentTypeRepositoryDb) isColExists(tableName string, colName string) error {
	query := "SELECT COUNT(*) FROM information_schema.columns WHERE table_name= '" + tableName + "' and column_name='" + colName + "'"
	row := r.db.QueryRow(query)
	var numOfCols int
	err := row.Scan(&numOfCols)
	fmt.Println(numOfCols, colName, query)
	if err != nil {
		return errs.ErrDb
	}
	if numOfCols == 0 {
		return errs.ErrColumnNotFound
	}
	return nil
}

func (r ContentTypeRepositoryDb) Create(n string, cols string) (string, error) {
	var query = "INSERT INTO contentType (name) VALUES ('" + n + "')"
	_, err := r.db.Exec(query)
	if err != nil {
		return "", errs.ErrDb
	}
	query = "CREATE TABLE " + n + " ( " + cols + " )"
	_, err = r.db.Exec(query)
	if err != nil {
		return "", errs.ErrDb
	}
	query = `SELECT currval(pg_get_serial_sequence('contentType','id'));`
	row := r.db.QueryRow(query)
	var id string
	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		return "", errs.ErrDb
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
		return errs.ErrDb
	}
	query = "DELETE FROM contentType" + " WHERE id=" + id
	_, err = r.db.Exec(query)
	if err != nil {
		return errs.ErrDb
	}
	return nil
}

func (r ContentTypeRepositoryDb) UpdateColName(id string, oldName string, newName string) error {
	name, err := r.isTableExists(id)
	if err != nil {
		return err
	}
	err = r.isColExists(name, oldName)
	if err != nil {
		return err
	}
	query := "ALTER TABLE " + name + " RENAME COLUMN " + oldName + " TO " + newName
	_, err = r.db.Exec(query)
	if err != nil {
		return errs.ErrDb
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
		return errs.ErrDb
	}
	return nil
}

func (r ContentTypeRepositoryDb) DeleteCol(id string, col string) error {
	name, err := r.isTableExists(id)
	if err != nil {
		return err
	}
	err = r.isColExists(name, col)
	if err != nil {
		return err
	}
	query := "ALTER TABLE " + name + " DROP COLUMN " + col
	fmt.Println(query)
	_, err = r.db.Exec(query)
	if err != nil {
		return errs.ErrDb
	}
	return nil
}

func NewContentTypeRepositoryDb(db *sql.DB) ContentTypeRepositoryDb {
	return ContentTypeRepositoryDb{db}
}