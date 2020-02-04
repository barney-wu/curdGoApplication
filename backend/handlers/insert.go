package handlers

import (
	"fmt"
	"log"
	"net/http"

	"../db"
)

func InsertHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("Insert Handler path", r.URL.Path)
	if r.Method == "POST" {
		username := r.FormValue("Username")
		password := r.FormValue("Password")
		email := r.FormValue("Email")
		insForm, err := db.DB.Prepare("INSERT INTO users(username, password, email) VALUES(?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(username, password, email)
		log.Println("Create User: Name: " + username)
	}
	http.Redirect(w, r, "/", 301)
}
