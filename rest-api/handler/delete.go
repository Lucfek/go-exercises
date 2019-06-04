package handler

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/lucfek/go-exercises/rest-api/response"
)

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
