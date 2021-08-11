package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/marwangalal746/backend/src/pkg/config"
	"github.com/marwangalal746/backend/src/pkg/content"
	"net/http"
	"strconv"
)

var app *config.AppConfig

func SetConfig(a *config.AppConfig) {
	app = a
}

func GetContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	query := `SELECT * FROM contents WHERE id=$1`
	intVar, _ := strconv.Atoi(params["id"])
	row := app.DB.QueryRow(query, intVar)
	var item content.Content
	switch err := row.Scan(&item.ID, &item.Title, &item.Details); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(item.ID, item.Title, item.Details)
	default:
		panic(err)

	}
	json.NewEncoder(w).Encode(item)
}
