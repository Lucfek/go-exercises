package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/lucfek/go-exercises/rest-api/model"
	"github.com/lucfek/go-exercises/rest-api/response"
)

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
