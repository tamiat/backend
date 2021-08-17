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

// NewConfig func to create a new instance of AppConfig struct
func NewConfig(db *sql.DB) *AppConfig {
	return &AppConfig{
		DB: db,
	}
}

// ConnectDB func to connect with db
func ConnectDB() *sql.DB {
	portNum, _ := strconv.Atoi(os.Getenv("PORTDB"))
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"), portNum, os.Getenv("USR"), os.Getenv("PASS"), os.Getenv("DBNAME"))
	var db *sql.DB
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
