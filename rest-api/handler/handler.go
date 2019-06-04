package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/lucfek/go-exercises/rest-api/model"
	"github.com/lucfek/go-exercises/rest-api/response"
)

type modelInter interface {
	Set(todo model.Todo) (model.Todo, error)
	Get(id uint64) (model.Todo, error)
	GetAll() ([]model.Todo, error)
	Update(id uint64, todo model.Todo) (model.Todo, error)
	Delete(id uint64) (model.Todo, error)
}

type postData struct {
	Name string
	Desc string
}

// Handler is a struct responsible for handling requests
type Handler struct {
	m modelInter
}

// Set is responsible for handling "SET" Requests
func (h Handler) Set(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res := response.Resp{}
	data := postData{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		res.Status = "ERROR"
		res.Data = err
		response.Writer(w, res)
		return
	}
	todo, err := h.m.Set(model.Todo{Name: data.Name, Description: data.Desc})
	if err != nil {
		res.Status = "ERROR"
		res.Data = err
		response.Writer(w, res)
		return
	}

	res.Status = "SUCCES"
	res.Data = todo
	response.Writer(w, res)

}

// Get is responsible for handling "GET" Requests
func (h Handler) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res := response.Resp{}
	id, err := strconv.ParseUint(p.ByName("id"), 10, 64)
	if err != nil {
		res.Status = "ERROR"
		res.Data = err
		response.Writer(w, res)
		return
	}
	todo, err := h.m.Get(id)
	if err != nil {
		res.Status = "ERROR"
		res.Data = err
		response.Writer(w, res)
		return
	}
	res.Status = "SUCCES"
	res.Data = todo
	response.Writer(w, res)

}

// GetAll is responsible for handling "GETALL" Requests
func (h Handler) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res := response.Resp{}
	todos, err := h.m.GetAll()
	if err != nil {
		res.Status = "ERROR"
		res.Data = err
		response.Writer(w, res)
		return
	}
	res.Status = "SUCCES"
	res.Data = todos
	response.Writer(w, res)
}

// Update is responsible for handling "UPDATE" Requests
func (h Handler) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res := response.Resp{}
	data := postData{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		res.Status = "ERROR"
		res.Data = err
		response.Writer(w, res)
		return
	}
	id, err := strconv.ParseUint(p.ByName("id"), 10, 64)
	if err != nil {
		res.Status = "ERROR"
		res.Data = err
		response.Writer(w, res)
		return
	}
	todo, err := h.m.Update(id, model.Todo{Name: data.Name, Description: data.Desc})
	if err != nil {
		res.Status = "ERROR"
		res.Data = err
		response.Writer(w, res)
		return
	}
	res.Status = "SUCCES"
	res.Data = todo
	response.Writer(w, res)
}

// Delete is responsible for handling "DELETE" Requests
func (h Handler) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res := response.Resp{}
	id, err := strconv.ParseUint(p.ByName("id"), 10, 64)
	if err != nil {
		res.Status = "ERROR"
		res.Data = err
		response.Writer(w, res)
		return
	}
	todo, err := h.m.Delete(id)
	if err != nil {
		res.Status = "ERROR"
		res.Data = err
		response.Writer(w, res)
		return
	}
	res.Status = "SUCCES"
	res.Data = todo
	response.Writer(w, res)
}

// New is a constructor of "Handler", it gets "Model" type Model as an argument and returns "Handler" type Handler
func New(m modelInter) Handler {
	return Handler{m: m}
}
