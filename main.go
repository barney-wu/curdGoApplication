package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"./backend/config"
	"./backend/db"
	"./backend/routes"
	"github.com/gorilla/handlers"
)

func main() {
	var file string
	//load config
	flag.StringVar(&file, "loadConfig", "./config/dev.json", "select env")
	flag.Parse()
	var configuration = config.LoadConfig(file)

	//initial mysql driver
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", configuration.DbUsername, configuration.DbPassword, configuration.DbHost, configuration.DbName)
	db.InitMySQL(connectionString)

	//register to routers
	router := routes.NewRouter()

	headersOk := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS", "PUT", "DELETE"})

	log.Fatal(http.ListenAndServe(":"+configuration.HTTPPort, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
