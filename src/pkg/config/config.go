package config

import (
	"database/sql"
	"fmt"
	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"strconv"
	_ "strconv"
)

type AppConfig struct {
	DB *sql.DB
}

func NewConfig(db *sql.DB) *AppConfig {
	return &AppConfig{
		DB: db,
	}
}

func ConnectDB() *sql.DB {
	portNum, _ := strconv.Atoi(os.Getenv("PORTDB"))
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"),portNum , os.Getenv("USR"), os.Getenv("PASS"), os.Getenv("DBNAME"))
	var db *sql.DB
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}