package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	_ "github.com/lib/pq" //Database driver
	"github.com/lucfek/go-exercises/rest-api/handler"
	"github.com/lucfek/go-exercises/rest-api/model"
)

var ip, dbAddr string

func init() {
	flag.StringVar(&ip, "ip", "127.0.0.1:8000", "Ip address the server will run on")
	flag.StringVar(&dbAddr, "db", "postgres://testuser:testpass@localhost:5555/testdb?sslmode=disable", "Address of database the server will handle")
	flag.Parse()
}

func main() {

	db, err := sql.Open("postgres", dbAddr)
	if err != nil {
		return
	}
	defer db.Close()

	model := model.New(db)
	router := httprouter.New()
	handler := handler.New(model)
	httpSrv := &http.Server{
		Handler:      router,
		Addr:         ip,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err != nil {
		log.Println(err)
		return
	}

	router.GET("/api/todos/", handler.GetAll)
	router.GET("/api/todos/:id", handler.Get)
	router.POST("/api/todos/", handler.Set)
	router.PATCH("/api/todos/:id", handler.Update)
	router.DELETE("/api/todos/:id", handler.Delete)

	fmt.Printf("Server is running on address: %s \n", ip)
	log.Fatal(httpSrv.ListenAndServe())
}
