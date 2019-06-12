package todos

import (
	"log"

	"github.com/lucfek/go-exercises/rest-api/model"
)

type modelInter interface {
	Set(todo model.Todo) (model.Todo, error)
	Get(id uint64) (model.Todo, error)
	GetAll() ([]model.Todo, error)
	Update(id uint64, todo model.Todo) (model.Todo, error)
	Delete(id uint64) (model.Todo, error)
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
