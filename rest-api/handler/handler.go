package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"github.com/julienschmidt/httprouter"
	"github.com/lucfek/go-exercises/rest-api/model"
)

type Model interface {
	Set(todo model.Todo) (model.Todo, error)
	GetAll() (model.Todo, error)
	Get(id string) (model.Todo, error)
	Update(id string) (model.Todo, error)
	Delete(id string) (model.Todo, error)
}

type Respone struct {
	Status string
	Data   interface{}
}

type Handler struct {
	M       Model
	idRegex *regexp.Regexp
}

func (h Handler) idCheck(id string) bool {
	return h.idRegex.MatchString(id)
}

func (h Handler) respWriter(w http.ResponseWriter, resp Respone) {
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

func (h Handler) Set(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	todo, err := h.M.Set(model.Todo{})
	if err != nil {
		res := Respone{
			Status: "Error",
			Data:   err,
		}
		h.respWriter(w, res)
		return
	}
	res := Respone{
		Status: "GETALL",
		Data:   todo,
	}
	h.respWriter(w, res)
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !h.idCheck(id) {
		err := errors.New("Invalid Id")
		res := Respone{
			Status: "Error",
			Data:   err,
		}
		h.respWriter(w, res)
		return
	}
	todo, err := h.M.Get(id)
	if err != nil {
		res := Respone{
			Status: "Error",
			Data:   err,
		}
		h.respWriter(w, res)
		return
	}
	res := Respone{
		Status: "GET",
		Data:   todo,
	}
	h.respWriter(w, res)
}

func (h Handler) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todo, err := h.M.GetAll()
	if err != nil {
		res := Respone{
			Status: "Error",
			Data:   err,
		}
		h.respWriter(w, res)
		return
	}
	res := Respone{
		Status: "GETALL",
		Data:   todo,
	}
	h.respWriter(w, res)
}

func (h Handler) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !h.idCheck(id) {
		err := errors.New("Invalid Id")
		res := Respone{
			Status: "Error",
			Data:   err,
		}
		h.respWriter(w, res)
		return
	}
	todo, err := h.M.Update(id)
	if err != nil {
		res := Respone{
			Status: "Error",
			Data:   err,
		}
		h.respWriter(w, res)
		return
	}
	res := Respone{
		Status: "UPDATE",
		Data:   todo,
	}
	h.respWriter(w, res)

}

func (h Handler) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !h.idCheck(id) {
		err := errors.New("Invalid Id")
		res := Respone{
			Status: "Error",
			Data:   err,
		}
		h.respWriter(w, res)
		return
	}
	todo, err := h.M.Delete(id)
	if err != nil {
		res := Respone{
			Status: "Error",
			Data:   err,
		}
		h.respWriter(w, res)
		return
	}
	res := Respone{
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
