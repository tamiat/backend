package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Content struct {
	gorm.Model
	Name    string
	Title   string
	Details string
}

var (
	content = &Content{Name: "Amgad", Title: "dummy", Details: "database test with go lang !"}
)
var db *gorm.DB
var err error

func main() {

	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)
	db, err = gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("succeded")
	}
	defer db.Close()

	//make migrations
	db.AutoMigrate(&Content{})
	db.Create(&content)

	r := mux.NewRouter()

	r.HandleFunc("/content/{id}", getData).Methods("GET")
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var content Content
	db.First(&content, params["id"])
	json.NewEncoder(w).Encode(content)
}
