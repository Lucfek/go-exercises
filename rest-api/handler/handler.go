package handler

import (
	"encoding/json"
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
	Update(id uint64) (model.Todo, error)
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

func (h Handler) handle(fn interface{}, param interface{}, w http.ResponseWriter, msg string) {
	var todo interface{}
	var err error

	switch f := fn.(type) {
	case func(id uint64) (model.Todo, error):
		sID := param.(string)
		id, err := strconv.ParseUint(sID, 10, 64)
		if err != nil {
			res := Response{
				Status: "Error",
				Data:   "Invalid Id",
			}
			h.respWriter(w, res)
			return
		}
		todo, err = f(id)
	case func() ([]model.Todo, error):
		todo, err = f()
	case func(todo model.Todo) (model.Todo, error):
		toDo, ok := param.(model.Todo)
		if !ok {
			res := Response{
				Status: "Error",
				Data:   "Internal server error",
			}
			h.respWriter(w, res)
			return
		}
		todo, err = f(toDo)
	default:
		res := Response{
			Status: "Error",
			Data:   "Internal server error",
		}
		h.respWriter(w, res)
		return
	}
	if err != nil {
		res := Response{
			Status: "Error",
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
	todo := model.Todo{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		res := Response{
			Status: "Error",
			Data:   err.Error(),
		}
		h.respWriter(w, res)
		return
	}
	h.handle(h.M.Set, todo, w, "SET")
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	h.handle(h.M.Get, p.ByName("id"), w, "GET")
}

func (h Handler) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h.handle(h.M.GetAll, "", w, "GETALL")
}

func (h Handler) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	h.handle(h.M.Update, p.ByName("id"), w, "UPDATE")
}

func (h Handler) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	h.handle(h.M.Delete, p.ByName("id"), w, "DELETE")
}

func New(m Model) (*Handler, error) {
	reg, err := regexp.Compile("^[0-9]+$")
	if err != nil {
		return &Handler{}, err
	}
	return &Handler{M: m, idRegex: reg}, nil
}
