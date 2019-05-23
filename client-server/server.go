package main

import (
	"encoding/json"
	database "go-exercises/client-server/database"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	db := database.New()

	r.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		user := new(database.User)
		user.Name = r.FormValue("name")
		user.Surname = r.FormValue("surname")
		user.Email = r.FormValue("email")

		db.Set(*user)
		w.Write([]byte("added"))
	})

	r.HandleFunc("/users/{id}/d", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if id, err := strconv.ParseUint(vars["id"], 10, 64); err == nil {
			db.Delete(id)
		}
	})

	r.HandleFunc("/users/{id}/", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if id, err := strconv.ParseUint(vars["id"], 10, 64); err == nil {
			if user, err := db.Get(id); err != "" {
				w.Write([]byte(err))
			} else {
				if b, err := json.Marshal(user); err == nil {
					w.Write(b)
				}
			}
		}
	})

	log.Fatal(http.ListenAndServe(":8000", r))

}
