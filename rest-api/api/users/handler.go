package users

import (
	"log"

	"github.com/lucfek/go-exercises/rest-api/model"
)

type modelInter interface {
	Login(user model.User) error
	Register(user model.User) error
}

// Handler is a struct responsible for handling requests
type Handler struct {
	m   modelInter
	log *log.Logger
}

// New is a constructor of "Handler", it gets "Model" type Model as an argument and returns "Handler" type Handler
func New(m modelInter, e *log.Logger) Handler {
	return Handler{m: m, log: e}
}
