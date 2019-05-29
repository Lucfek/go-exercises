package main

import (
	"flag"
	"fmt"
	"go-exercises/client-server/database"
	"go-exercises/client-server/handle"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	var ip = flag.String("ip", "127.0.0.1:8000", "Ip address the server will run on")
	var saveDelay = flag.Int("savedelay", 1, "Delay between saves in minuts")
	var newDb = flag.Bool("new", false, "Set as true if you want to create new database this will overwrite the old one")
	var file = flag.String("file", "database.json", "Name of database file")
	flag.Parse()

	db, err := database.New(*file)
	if err != nil {
		log.Println(err)
		return
	}
	if !*newDb {
		if err := db.Load(); err != nil {
			log.Println(err)
			return
		}
	}

	go db.Saving(*saveDelay)

	handler := handle.New(db)

	r := mux.NewRouter()
	r.HandleFunc("/users/", handler.Set).Methods("POST")
	r.HandleFunc("/users/{id}/", handler.Get).Methods("GET")
	r.HandleFunc("/users/{id}/", handler.Del).Methods("DELETE")

	srv := &http.Server{
		Handler:      r,
		Addr:         *ip,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Server is running on address: %s \n", *ip)
	log.Fatal(srv.ListenAndServe())

}
