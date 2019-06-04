package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"

	_ "github.com/lib/pq" //Database driver
	"github.com/lucfek/go-exercises/rest-api/dbhandler"
	"github.com/lucfek/go-exercises/rest-api/dbmodel"
)

var ipAddr, dbAddr string

func init() {
	flag.StringVar(&ipAddr, "ip", "127.0.0.1:8000", "Ip address the server will run on")
	flag.StringVar(&dbAddr, "db", "postgres://testuser:testpass@localhost:5555/testdb?sslmode=disable", "Address of database the server will handle")
}

func main() {

	db, err := sql.Open("postgres", dbAddr)
	if err != nil {
		return
	}
	defer db.Close()

	model := dbmodel.New(db)
	router := httprouter.New()
	errLog := log.New(os.Stderr,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	flag.Parse()
	handler := dbhandler.New(model, errLog)
	httpSrv := &http.Server{
		Handler:      router,
		Addr:         ipAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	router.GET("/api/todos/", handler.GetAll)
	router.GET("/api/todos/:id/", handler.Get)
	router.POST("/api/todos/", handler.Set)
	router.PATCH("/api/todos/:id/", handler.Update)
	router.DELETE("/api/todos/:id/", handler.Delete)

	fmt.Printf("Server is running on address: %s \n", ipAddr)
	log.Fatal(httpSrv.ListenAndServe())
}
