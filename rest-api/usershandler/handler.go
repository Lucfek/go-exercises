package usershandler

import (
	"log"

	"github.com/lucfek/go-exercises/rest-api/usersmodel"
)

type modeler interface {
	Add(user usersmodel.User) (usersmodel.User, error)
	Get(email string) (usersmodel.User, error)
}

// Handler is a struct responsible for handling requests
type Handler struct {
	m   modeler
	log *log.Logger
}

// New is a constructor of "Handler", it gets "Model" type Model as an argument and returns "Handler" type Handler
func New(m modeler, e *log.Logger) Handler {
	return Handler{m: m, log: e}
}
