package content

import (
	"database/sql"
	"errors"
	"log"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// ContentRepositoryDb implements the ContentRepository interface
type ContentRepositoryDb struct {
	db *sql.DB
}

func (r ContentRepositoryDb) ReadAll() ([]Content, error) {
	findAllSQL := "SELECT id, title, details FROM contents"
	rows, err := r.db.Query(findAllSQL)
	if err != nil {
		log.Println("Unexpected database error" + err.Error())
		return nil, err
	}
	contents := make([]Content, 0)
	for rows.Next() {
		var c Content
		err := rows.Scan(&c.Id, &c.Title, &c.Details)
		if err != nil {
			log.Println("Unexpected database error" + err.Error())
			return nil, err
		}
		contents = append(contents, c)
	}
	return contents, nil
}
func (d ContentRepositoryDb) ReadById(id string) (*Content, error) {
	contentSQL := "SELECT id, title, details FROM contents WHERE id = $1"
	row := d.db.QueryRow(contentSQL, id)
	var c Content
	err := row.Scan(&c.Id, &c.Title, &c.Details)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("content not found")
		} else {
			log.Println("error while querying in repoDB" + err.Error())
			return nil, errors.New("Unexpected database error")
		}
	}
	return &c, nil
}

func (d ContentRepositoryDb) ReadRange(ids []string) ([]Content, error) {
	var res []Content
	query := `SELECT id, title, details FROM contents WHERE id>=$1 AND id<=$2`
	rows, err := d.db.Query(query, ids[0], ids[1])
	if err != nil {
		panic(err)
		return res, errors.New("Unexpected database error")
	}
	for rows.Next() {
		var item Content
		switch err := rows.Scan(&item.Id, &item.Title, &item.Details); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			res = append(res, item)
		default:
			panic(err)
		}
	}
	return res, nil
}

func (d ContentRepositoryDb) Create(newContent Content) (string, error) {
	query := `INSERT INTO contents (title, details) VALUES ($1, $2)`
	_, err := d.db.Exec(query, newContent.Title, newContent.Details)
	if err != nil {
		return "", errors.New("Unexpected database error")
	}
	query = `SELECT currval(pg_get_serial_sequence('contents','id'));`
	row := d.db.QueryRow(query)
	var id string
	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		return "", nil
	case nil:
		return id, nil
	default:
		fmt.Println(id)
		panic(err)
	}
}

func (d ContentRepositoryDb) DeleteById(id string) error {
	query := `DELETE FROM contents WHERE id=$1`
	_, err := d.db.Exec(query,id)
	if err != nil {
		return errors.New("Unexpected database error")
	}
	return nil
}

func (d ContentRepositoryDb) UpdateById(id string, UpdContent Content) error {
	query := `UPDATE contents SET title=$1, details=$2 Where id=$3`
	_, err := d.db.Exec(query, UpdContent.Title, UpdContent.Details, id)
	if err != nil {
		return errors.New("Unexpected database error")
	}
	return nil
}

func NewContentRepositoryDb(db *sql.DB) ContentRepositoryDb {
	return ContentRepositoryDb{db}
}
