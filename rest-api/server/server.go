package server

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/lucfek/go-exercises/rest-api/handler"
	"github.com/lucfek/go-exercises/rest-api/model"
)

type Server struct {
	model   *model.Model
	Router  *httprouter.Router
	Handler *handler.Handler
	conf    *http.Server
}

func New(srvAddr, dbAddr string) (*Server, error) {
	model, err := model.New(dbAddr)
	if err != nil {
		return &Server{}, err
	}
	handler, err := handler.New(model)
	if err != nil {
		return &Server{}, err
	}
	router := httprouter.New()
	conf := &http.Server{
		Handler:      router,
		Addr:         srvAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return &Server{model: model, Handler: handler, Router: router, conf: conf}, nil
}
func (s Server) CloseDB() {
	s.model.Close()
}

func (s Server) Run() {
	log.Fatal(s.conf.ListenAndServe())
}
