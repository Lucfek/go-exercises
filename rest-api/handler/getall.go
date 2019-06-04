package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lucfek/go-exercises/rest-api/response"
)

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
