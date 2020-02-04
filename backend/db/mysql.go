package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

//InitMySQL sets max lifetime to 5 mins
func InitMySQL(connectionString string) {
	var err error
	DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Println(err)
	}
	// Set the maximum lifetime of a connection to 5 minutes. Setting it to 0
	// means that there is no maximum lifetime and the connection is reused
	// forever (which is the default behavior).
	DB.SetConnMaxLifetime(time.Minute * 5)

	if err = DB.Ping(); err != nil {
		log.Println(err)
	}
}
