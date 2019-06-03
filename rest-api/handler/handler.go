package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/lucfek/go-exercises/rest-api/model"
)

type modelInter interface {
	Set(todo model.Todo) (model.Todo, error)
	Get(id uint64) (model.Todo, error)
	GetAll() ([]model.Todo, error)
	Update(id uint64, todo model.Todo) (model.Todo, error)
	Delete(id uint64) (model.Todo, error)
}

type response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// Handler is a struct responsible for handling requests
type Handler struct {
	m modelInter
}

func (h Handler) respWriter(w http.ResponseWriter, resp response) {
	b, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

// Set is responsible for handling "SET" Requests
func (h Handler) Set(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error
	todo := model.Todo{}
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err == nil {
		todo, err = h.m.Set(todo)
	}
	if err != nil {
		res := response{
			Status: "ERROR",
			Data:   err.Error(),
		}
		h.respWriter(w, res)
		return
	}
	res := response{
		Status: "GET",
		Data:   todo,
	}
	h.respWriter(w, res)

}

// Get is responsible for handling "GET" Requests
func (h Handler) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error
	var id uint64
	var todo model.Todo
	id, err = strconv.ParseUint(p.ByName("id"), 10, 64)
	if err == nil {
		todo, err = h.m.Get(id)
	}
	if err != nil {
		res := response{
			Status: "ERROR",
			Data:   err.Error(),
		}
		h.respWriter(w, res)
		return
	}
	res := response{
		Status: "GET",
		Data:   todo,
	}
	h.respWriter(w, res)
}

// GetAll is responsible for handling "GETALL" Requests
func (h Handler) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todo, err := h.m.GetAll()
	if err != nil {
		res := response{
			Status: "ERROR",
			Data:   err.Error(),
		}
		h.respWriter(w, res)
		return
	}
	res := response{
		Status: "GETALL",
		Data:   todo,
	}
	h.respWriter(w, res)
}

// Update is responsible for handling "UPDATE" Requests
func (h Handler) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error
	var id uint64
	todo := model.Todo{}
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err == nil {
		id, err = strconv.ParseUint(p.ByName("id"), 10, 64)
		if err == nil {
			todo, err = h.m.Update(id, todo)
		}
	}
	if err != nil {
		res := response{
			Status: "ERROR",
			Data:   err.Error(),
		}
		h.respWriter(w, res)
		return
	}
	res := response{
		Status: "UPDATE",
		Data:   todo,
	}
	h.respWriter(w, res)
}

// Delete is responsible for handling "DELETE" Requests
func (h Handler) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error
	var id uint64
	var todo model.Todo
	id, err = strconv.ParseUint(p.ByName("id"), 10, 64)
	if err == nil {
		todo, err = h.m.Get(id)
	}
	if err != nil {
		res := response{
			Status: "ERROR",
			Data:   err.Error(),
		}
		h.respWriter(w, res)
		return
	}
	res := response{
		Status: "DELETE",
		Data:   todo,
	}
	h.respWriter(w, res)
}

// New is a constructor of "Handler", it gets "Model" type Model as an argument and returns "Handler" type Handler
func New(m modelInter) *Handler {
	return &Handler{m: m}
}
