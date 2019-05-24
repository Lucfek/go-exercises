package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	database "go-exercises/client-server/database"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func askQue(question string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question + " (Y/N):")

	switch answear, _, _ := reader.ReadRune(); answear {
	case 'Y', 'y':
		return true
	case 'N', 'n':
		return false
	}
	askQue(question)
	return false
}

var db *database.Database

func setHandler(w http.ResponseWriter, r *http.Request) {
	user := new(database.User)
	user.Name = r.FormValue("name")
	user.Surname = r.FormValue("surname")
	user.Email = r.FormValue("email")

	err := db.Set(*user)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("User succesfuly added"))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	user, err := db.Get(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	b, err := json.Marshal(user)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(b)
}
func delHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	err = db.Delete(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}
func saving(delay int) {
	for {
		db.Save()
		time.Sleep(time.Duration(delay) * time.Minute)
	}
}

func main() {
	var ip = flag.String("ip", "127.0.0.1:8000", "Ip address the server will run on")
	var saveDelay = flag.Int("savedelay", 1, "Delay between saves in minuts")
	flag.Parse()

	if askQue("Do you want to create new database? (This will overwritwe the old one)") {
		db = database.New()
	} else {
		db = database.Load()
	}

	go saving(*saveDelay)

	r := mux.NewRouter()
	r.HandleFunc("/users/", setHandler).Methods("POST")
	r.HandleFunc("/users/{id}/", getHandler).Methods("GET")
	r.HandleFunc("/users/{id}/", delHandler).Methods("DELETE")

	srv := &http.Server{
		Handler:      r,
		Addr:         *ip,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server is running")
	log.Fatal(srv.ListenAndServe())

}
