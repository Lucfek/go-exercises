package main

import (
	"bufio"
	"flag"
	"fmt"
	database "go-exercises/client-server/database"
	"log"
	"net/http"
	"os"
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

func saving(db *database.Database, delay int) {
	for {
		err := db.Save()
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Duration(delay) * time.Minute)
	}
}

func main() {
	var ip = flag.String("ip", "127.0.0.1:8000", "Ip address the server will run on")
	var saveDelay = flag.Int("savedelay", 1, "Delay between saves in minuts")
	flag.Parse()

	var db *database.Database
	if askQue("Do you want to create new database? (This will overwritwe the old one)") {
		db = database.New()
	} else {
		var err error
		db, err = database.Load()
		if err != nil {
			log.Println(err)
			return
		}
	}

	go saving(db, *saveDelay)

	r := mux.NewRouter()
	r.HandleFunc("/users/", db.SetHandler).Methods("POST")
	r.HandleFunc("/users/{id}/", db.GetHandler).Methods("GET")
	r.HandleFunc("/users/{id}/", db.DelHandler).Methods("DELETE")

	srv := &http.Server{
		Handler:      r,
		Addr:         *ip,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Server is running on address: %s \n", *ip)
	log.Fatal(srv.ListenAndServe())

}
