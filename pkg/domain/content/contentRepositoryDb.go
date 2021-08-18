package content

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"os"
)

type ContentRepositoryDb struct {
	db *sql.DB
}

func (r ContentRepositoryDb) FindAll() ([]Content, error) {
	findAllSQL := "SELECT id, title, details FROM contents"
	rows, err := r.db.Query(findAllSQL)
	if err != nil {
		log.Println("error while querying " + err.Error())
		return nil, err
	}
	contents := make([]Content, 0)
	for rows.Next() {
		var c Content
		err := rows.Scan(&c.Id, &c.Title, &c.Details)
		if err != nil {
			log.Println("error while querying " + err.Error())
			return nil, err
		}
		contents = append(contents, c)
	}
	return contents, nil
}
func (d ContentRepositoryDb) ById(id string) (*Content, error) {
	contentSQL := "SELECT id, title, details FROM contents WHERE id = $1"
	row := d.db.QueryRow(contentSQL, id)
	var c Content
	err := row.Scan(&c.Id, &c.Title, &c.Details)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("content not found")
		} else {
			log.Println("error while querying in repoDB" + err.Error())
			return nil, errors.New("unexpected database error")
		}
	}
	return &c, nil
}

func NewContentRepositoryDb() ContentRepositoryDb {

	dataSourceName := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
		os.Getenv("HOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBNAME"),
		os.Getenv("USER"),
		os.Getenv("PASS"))
	db, err := sql.Open("pgx", dataSourceName)
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to conect to db"))
		panic(err)
	}
	log.Println("connected to db ")

	//test connection
	err = db.Ping()
	if err != nil {
		log.Fatal("cannot ping db")
		panic(err)
	}
	log.Println("pinged db")
	return ContentRepositoryDb{db}
}
