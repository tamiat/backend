package main

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/tamiat/backend/internals/config"
	"github.com/tamiat/backend/internals/handlers"
	"github.com/tamiat/backend/internals/routes"
)
var app config.AppConfig
func main(){
	/*app.USER = os.Getenv("USER")
	app.PASS = os.Getenv("PASS")
	app.HOST = os.Getenv("HOST")
	app.DBNAME = os.Getenv("DBNAME")
	app.DBPORT = os.Getenv("DBPORT")*/
	
	fmt.Println("starting app")
	routes.Routes()
}
