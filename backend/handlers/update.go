package handlers

import (
	"fmt"
	"log"
	"net/http"

	"../db"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("Update Handler path", r.URL.Path)
	if r.Method == "POST" {
		uid := r.FormValue("user_id")
		username := r.FormValue("Username")
		password := r.FormValue("Password")
		email := r.FormValue("Email")
		insForm, err := db.DB.Prepare("UPDATE users SET username=?, password=?, email=? WHERE user_id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(username, password, email, uid)
		log.Println("UPDATE: User Name: " + username)
	}
	http.Redirect(w, r, "/", 301)
}
