package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/lucfek/go-exercises/rest-api/model"
)

type Model interface {
	Set(todo model.Todo) (model.Todo, error)
	Get(id uint64) (model.Todo, error)
	GetAll() ([]model.Todo, error)
	Update(id uint64, todo model.Todo) (model.Todo, error)
	Delete(id uint64) (model.Todo, error)
}

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type Handler struct {
	M       Model
	idRegex *regexp.Regexp
}
type updatePack struct {
	id   string
	todo model.Todo
}

func (h Handler) handle(fn interface{}, param interface{}, w http.ResponseWriter, msg string) {
	var todo interface{}
	var err error
	var id uint64

	switch f := fn.(type) {
	case func(id uint64) (model.Todo, error):
		sID := param.(string)
		id, err = strconv.ParseUint(sID, 10, 64)
		if err == nil {
			todo, err = f(id)
		}
	case func() ([]model.Todo, error):
		todo, err = f()
	case func(todo model.Todo) (model.Todo, error):
		toDo, ok := param.(model.Todo)
		if ok {
			todo, err = f(toDo)
		} else {
			err = errors.New("Internal server error")
		}
	case func(id uint64, todo model.Todo) (model.Todo, error):
		values, ok := param.(updatePack)
		id, err = strconv.ParseUint(values.id, 10, 64)
		if ok {
			if err == nil {
				todo, err = f(id, values.todo)
			}
		} else {
			err = errors.New("Internal server error")
		}
	default:
		err = errors.New("Internal server error")
	}
	if err != nil {
		res := Response{
			Status: "ERROR",
			Data:   err.Error(),
		}
		h.respWriter(w, res)
		return
	}
	res := Response{
		Status: msg,
		Data:   todo,
	}
	h.respWriter(w, res)
}

func (h Handler) respWriter(w http.ResponseWriter, resp Response) {
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

func (h Handler) Set(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error
	todo := model.Todo{}
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err == nil {
		todo, err = h.M.Set(todo)
	}
	if err != nil {
		res := Response{
			Status: "ERROR",
			Data:   err.Error(),
		}
		h.respWriter(w, res)
		return
	}
	h.handle(h.M.Set, todo, w, "SET")

}

func (h Handler) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error
	var id uint64
	var todo model.Todo
	id, err = strconv.ParseUint(p.ByName("id"), 10, 64)
	if err == nil {
		todo, err = h.M.Get(id)
	}
	if err != nil {
		res := Response{
			Status: "ERROR",
			Data:   err.Error(),
		}
		h.respWriter(w, res)
		return
	}
	res := Response{
		Status: "GET",
		Data:   todo,
	}
	h.respWriter(w, res)
}

func (h Handler) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todo, err := h.M.GetAll()
	if err != nil {
		res := Response{
			Status: "ERROR",
			Data:   err.Error(),
		}
		h.respWriter(w, res)
		return
	}
	res := Response{
		Status: "GETALL",
		Data:   todo,
	}
	h.respWriter(w, res)
}

func (h Handler) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error
	var id uint64
	todo := model.Todo{}
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err == nil {
		id, err = strconv.ParseUint(p.ByName("id"), 10, 64)
		if err == nil {
			todo, err = h.M.Update(id, todo)
		}
	}
	if err != nil {
		res := Response{
			Status: "ERROR",
			Data:   err.Error(),
		}
		h.respWriter(w, res)
		return
	}
	res := Response{
		Status: "UPDATE",
		Data:   todo,
	}
	h.respWriter(w, res)
}

func (h Handler) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error
	var id uint64
	var todo model.Todo
	id, err = strconv.ParseUint(p.ByName("id"), 10, 64)
	if err == nil {
		todo, err = h.M.Get(id)
	}
	if err != nil {
		res := Response{
			Status: "ERROR",
			Data:   err.Error(),
		}
		h.respWriter(w, res)
		return
	}
	res := Response{
		Status: "DELETE",
		Data:   todo,
	}
	h.respWriter(w, res)
}

func New(m Model) (*Handler, error) {
	reg, err := regexp.Compile("^[0-9]+$")
	if err != nil {
		return &Handler{}, err
	}
	return &Handler{M: m, idRegex: reg}, nil
}
