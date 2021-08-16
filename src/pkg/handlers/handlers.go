package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/marwangalal746/backend/src/pkg/config"
	"github.com/marwangalal746/backend/src/pkg/content"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var app *config.AppConfig

func SetConfig(a *config.AppConfig) {
	app = a
}

func parseNums(str string) []string {
	var params []string
	if str[0] >= '0' && str[0] <= '9' {
		params = append(params, str)
	} else {
		ind := strings.IndexByte(str, ':')
		params = append(params, str[1:ind])
		params = append(params, str[ind+1:len(str)-1])
	}
	return params
}

func GetContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	pattern1, _ := regexp.Match(`^[0-9]+$`, []byte(params["id"]))
	pattern2, _ := regexp.Match(`^([)([0-9]+)[:]([0-9]+)(])$`, []byte(params["id"]))
	if !pattern1 && !pattern2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//fmt.Println(params["id"])
	idValues := parseNums(params["id"])
	var items []content.Content
	if len(idValues) == 1 {
		items = append(items, getOneContent(idValues[0])...)
	} else {
		items = append(items, getRangeOfContents(idValues)...)
	}
	for i := 0; i < len(items); i++ {
		json.NewEncoder(w).Encode(items[i])
	}
}

func getOneContent(id string) []content.Content {
	query := `SELECT * FROM contents WHERE id=$1`
	row := app.DB.QueryRow(query, id)
	var item content.Content
	switch err := row.Scan(&item.ID, &item.Title, &item.Details); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(item.ID, item.Title, item.Details)
	default:
		panic(err)
	}
	res := []content.Content{
		item,
	}
	return res
}

func getRangeOfContents(ids []string) []content.Content {
	var res []content.Content
	from, _ := strconv.Atoi(ids[0])
	to, _ := strconv.Atoi(ids[1])
	query := `SELECT * FROM contents WHERE id>=$1 AND id<=$2`
	rows, err := app.DB.Query(query, from, to)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var item content.Content
		switch err := rows.Scan(&item.ID, &item.Title, &item.Details); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			res = append(res, item)
		default:
			panic(err)
		}
	}
	return res
}
