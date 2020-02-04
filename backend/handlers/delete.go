package handlers

import (
	"fmt"
	"log"
	"net/http"

	"../db"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("Delete Handler path", r.URL.Path)

	uid := r.URL.Query().Get("id")
	delForm, err := db.DB.Prepare("DELETE FROM users WHERE user_id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(uid)
	log.Println("DELETE: User ID: " + uid)
	http.Redirect(w, r, "/", 301)
}
