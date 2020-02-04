package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"../db"
	"../model"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	usr := model.User{}
	res := []model.User{}

	var tmpl = template.Must(template.ParseGlob("frontend/*"))
	r.ParseForm()
	fmt.Println("Index path", r.URL.Path)
	rows, err := db.DB.Query("SELECT * FROM users ORDER BY user_id DESC")
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var uid int
		var username string
		var password string
		var email string
		err := rows.Scan(&uid, &username, &password, &email)
		if err != nil {
			log.Fatal("IndexHandler", err)
		}
		usr.Id = uid
		usr.Username = username
		usr.Password = password
		usr.Email = email
		res = append(res, usr)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	fmt.Println("path", r.URL.Path)
}
