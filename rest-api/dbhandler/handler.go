package dbhandler

import "github.com/lucfek/go-exercises/rest-api/dbmodel"

type modelInter interface {
	Set(todo dbmodel.Todo) (dbmodel.Todo, error)
	Get(id uint64) (dbmodel.Todo, error)
	GetAll() ([]dbmodel.Todo, error)
	Update(id uint64, todo dbmodel.Todo) (dbmodel.Todo, error)
	Delete(id uint64) (dbmodel.Todo, error)
}

type errLogger interface {
	Println(v ...interface{})
}

// Handler is a struct responsible for handling requests
type Handler struct {
	m   modelInter
	log errLogger
}

// New is a constructor of "Handler", it gets "Model" type Model as an argument and returns "Handler" type Handler
func New(m modelInter, e errLogger) Handler {
	return Handler{m: m, log: e}
}
