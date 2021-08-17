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

//declare an instance from AppConfig struct to make db accessible in this package
var app *config.AppConfig

func SetConfig(a *config.AppConfig) {
	app = a
}

//this func is used for extracting the numbers which passes as id values
//it takes the value which client passes it as an id in the request (the value is a string) and it returns an array contain the values after parsing
//there is to cases:
//1- if the user wants to get a specific content so this function return an array with size 1
//if the string in case 1 it will return an array which contain the string itself
//2- if the user wants to get a range of contents
//if in case2, it will return an array with size 2 (first element is the start of the rang and second is the end)
func parseNums(str string) []string {
	var params []string
	if str[0] >= '0' && str[0] <= '9' { //case 1
		params = append(params, str)
	} else { //case 2
		ind := strings.IndexByte(str, ':')
		params = append(params, str[1:ind])
		params = append(params, str[ind+1:len(str)-1])
	}
	return params
}

// GetContent endpoint to get contents
func GetContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//regular expression to check if the string has numbers only	example: 1234
	pattern1, _ := regexp.Match(`^[0-9]+$`, []byte(params["id"]))
	//regular expression to check if the string in the pattern of this examples ([1:2], [35:40])
	pattern2, _ := regexp.Match(`^([)([0-9]+)[:]([0-9]+)(])$`, []byte(params["id"]))
	//if the string can't match with any RG, the response will be 400 (badrequest)
	if !pattern1 && !pattern2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idValues := parseNums(params["id"])
	var items []content.Content
	if len(idValues) == 1 {
		items = append(items, getOneContent(idValues[0])...)
	} else {
		items = append(items, getRangeOfContents(idValues)...)
	}
	if len(items) < 1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	for i := 0; i < len(items); i++ {
		json.NewEncoder(w).Encode(items[i])
	}
}

//utility function to get a specific content frim db
func getOneContent(id string) []content.Content {
	query := `SELECT id, title, details FROM contents WHERE id=$1`
	row := app.DB.QueryRow(query, id)
	var item content.Content
	switch err := row.Scan(&item.ID, &item.Title, &item.Details); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(item.ID, item.Title, item.Details)
	default:
		fmt.Println(item.ID)
		panic(err)
	}
	var res []content.Content
	if item.ID != "" {
		res = append(res, item)
	}
	return res
}

//utility function to get a specific range of contents
func getRangeOfContents(ids []string) []content.Content {
	var res []content.Content
	from, _ := strconv.Atoi(ids[0])
	to, _ := strconv.Atoi(ids[1])
	query := `SELECT id, title, details FROM contents WHERE id>=$1 AND id<=$2`
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

func PostContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newContent content.Content
	err := json.NewDecoder(r.Body).Decode(&newContent)
	if err != nil {
		panic(err)
	}
	query := `INSERT INTO contents (title, details) VALUES ($1, $2)`
	res, err := app.DB.Exec(query, newContent.Title, newContent.Details)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(res)
}
