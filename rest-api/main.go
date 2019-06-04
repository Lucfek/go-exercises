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

type server struct {
	model   model.Model
	router  *httprouter.Router
	handler handler.Handler
	httpSrv *http.Server
}

func init() {
	flag.StringVar(&ip, "ip", "127.0.0.1:8000", "Ip address the server will run on")
	flag.StringVar(&dbAddr, "db", "postgres://testuser:testpass@localhost:5555/testdb?sslmode=disable", "Address of database the server will handle")
	flag.Parse()
}

var ip, dbAddr string

func main() {

	db, err := sql.Open("postgres", dbAddr)
	if err != nil {
		return
	}
	defer db.Close()
	srv := server{}
	srv.model = model.New(db)
	srv.router = httprouter.New()
	srv.handler = handler.New(srv.model)
	srv.httpSrv = &http.Server{
		Handler:      srv.router,
		Addr:         ip,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err != nil {
		log.Println(err)
		return
	}

	srv.router.GET("/api/todos/", srv.handler.GetAll)
	srv.router.GET("/api/todos/:id", srv.handler.Get)
	srv.router.POST("/api/todos/", srv.handler.Set)
	srv.router.PATCH("/api/todos/:id", srv.handler.Update)
	srv.router.DELETE("/api/todos/:id", srv.handler.Delete)

	fmt.Printf("Server is running on address: %s \n", ip)
	log.Fatal(srv.httpSrv.ListenAndServe())
}
