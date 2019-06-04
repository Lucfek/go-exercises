package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lucfek/go-exercises/rest-api/model"
	"github.com/lucfek/go-exercises/rest-api/response"
)

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
