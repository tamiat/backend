package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tamiat/backend/internals/config"
	"github.com/tamiat/backend/internals/models"
	"log"
	"net/http"
	"strconv"
)

var app *config.AppConfig

func SetAppConfig(a *config.AppConfig) {
	app = a
}

func GetContentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["content_id"])
	fmt.Println(id)
	content,err := getContent(app.Db,id)
	if err!=nil {
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(content)
}
func getContent(db *sql.DB,id int) (models.Content, error) {
	var content models.Content
	row := db.QueryRow("SELECT id,title,details FROM contents WHERE id = $1",id)
	err := row.Scan(&content.Id, &content.Title, &content.Details)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("QueryRow returns", content.Id, content.Title, content.Details)

	row.Scan(&content.Id,&content.Title,&content.Details)
	return content, nil
}

func ConnectToDB()(*sql.DB,error){
	//TODO change to env variables
	//connecting to db
	var db *sql.DB
	dataSourceName:= fmt.Sprintf( "host={%s} port={%s} dbname={%s} user={%s} password ={%s}",app.HOST,app.DBPORT,app.DBNAME,app.USER,app.PASS)
	db, err := sql.Open("pgx", dataSourceName)
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to conect to db"))
		return db,err
	}
	//defer db.Close()
	log.Println("connected to db ")

	//test connection
	err = db.Ping()
	if err!=nil{
		log.Fatal("cannot ping db")
		return db,err
	}
	log.Println("pinged db")
	return db,err
}
