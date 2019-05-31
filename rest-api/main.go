package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/lucfek/go-exercises/rest-api/server"
)

func main() {
	var ip = flag.String("ip", "127.0.0.1:8000", "Ip address the server will run on")
	var db = flag.String("db", "postgres://testuser:testpass@localhost:5555/testdb?sslmode=disable", "Address of database the server will handle")
	flag.Parse()

	srv, err := server.New(*ip, *db)
	if err != nil {
		log.Println(err)
		return
	}
	defer srv.CloseDB()
	srv.Router.GET("/api/todos/", srv.Handler.GetAll)
	srv.Router.GET("/api/todos/:id", srv.Handler.Get)
	srv.Router.POST("/api/todos/", srv.Handler.Set)
	srv.Router.PATCH("/api/todos/:id", srv.Handler.Update)
	srv.Router.DELETE("/api/todos/:id", srv.Handler.Delete)

	fmt.Printf("Server is running on address: %s \n", *ip)
	srv.Run()
}
