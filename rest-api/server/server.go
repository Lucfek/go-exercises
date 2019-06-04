package server

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/lucfek/go-exercises/rest-api/handler"
	"github.com/lucfek/go-exercises/rest-api/model"
)

// Server struct
type Server struct {
	model   *model.Model
	Router  *httprouter.Router
	Handler *handler.Handler
	conf    *http.Server
}

// New takes server address and database address as parameters and returns Server struct
func New(srvAddr, dbAddr string) (*Server, error) {
	model, err := model.New(dbAddr)
	if err != nil {
		return &Server{}, err
	}
	handler := handler.New(model)
	router := httprouter.New()
	conf := &http.Server{
		Handler:      router,
		Addr:         srvAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return &Server{model: model, Handler: handler, Router: router, conf: conf}, nil
}

// CloseDB ends connection with database
func (s Server) CloseDB() {
	s.model.Close()
}

// Run starts the server
func (s Server) Run() {
	log.Fatal(s.conf.ListenAndServe())
}
