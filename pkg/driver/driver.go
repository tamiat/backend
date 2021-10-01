package driver

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func GetDbConnetion() (*gorm.DB, *sql.DB) {
	dataSourceName := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
		os.Getenv("HOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBNAME"),
		os.Getenv("USER"),
		os.Getenv("PASS"))
	sqlDB, err := sql.Open("pgx", dataSourceName)
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to conect to db"))
		panic(err)
	}
	log.Println("connected to db ")

	//test connection
	/*err = db.Ping()
	if err != nil {
		log.Fatal("cannot ping db")
		panic(err)
	}*/
	log.Println("pinged db")
	return db, sqlDB
}