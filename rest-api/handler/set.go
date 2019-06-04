package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lucfek/go-exercises/rest-api/model"
	"github.com/lucfek/go-exercises/rest-api/response"
)

// Set is responsible for handling "SET" Requests
func (h Handler) Set(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := postData{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		res := response.Resp{
			Status: "error",
			Data:   "There was an problem, please try again",
		}
		response.Writer(w, res)
		return
	}
	todo, err := h.m.Set(model.Todo{Name: data.Name, Description: data.Desc})
	if err != nil {
		log.Println(err)
		res := response.Resp{
			Status: "error",
			Data:   "There was an problem, please try again",
		}
		response.Writer(w, res)
		return
	}

	res := response.Resp{
		Status: "succes",
		Data:   todo,
	}
	response.Writer(w, res)

}
