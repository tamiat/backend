package config

import (
	"database/sql"
)

type AppConfig struct {
	Db *sql.DB
	USER string
	PASS string
	DBNAME string
	HOST string
	DBPORT string
}



