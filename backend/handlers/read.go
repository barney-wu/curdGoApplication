package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"../db"
	"../model"
)

func ReadHandler(w http.ResponseWriter, r *http.Request) {
	usr := model.User{}
	//res := []model.User{}
	var tmpl = template.Must(template.ParseGlob("frontend/*"))
	//pstr := strings.Split(r.URL.Path, "/")
	uid := r.URL.Query().Get("id")
	var rows *sql.Rows
	if len(uid) != 0 {
		rows, _ = db.DB.Query("SELECT * FROM users WHERE user_id=?", uid)
		for rows.Next() {
			var uid int
			var username string
			var password string
			var email string
			err := rows.Scan(&uid, &username, &password, &email)
			if err != nil {
				log.Fatal("ReadHandler", err)
			}
			usr.Id = uid
			usr.Username = username
			usr.Password = password
			usr.Email = email
			//res = append(res, usr)
		}
		tmpl.ExecuteTemplate(w, "Read", usr)
		fmt.Println("path", r.URL.Path)

	} else {
		fmt.Fprintf(w, "No user found")
	}

}
