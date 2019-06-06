package dbhandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/lucfek/go-exercises/rest-api/dbmodel"
	"github.com/lucfek/go-exercises/rest-api/response"
)

// Update is responsible for handling "UPDATE" Requests
func (h Handler) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	data := postData{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		h.log.Println(err)
		res := response.Resp{
			Status: "error",
			Data:   "There was an problem, please try again",
		}
		response.Writer(w, res)
		return
	}
	if data.Name == "" || data.Desc == "" {
		res := response.Resp{
			Status: "error",
			Data:   "Empty post values",
		}
		response.Writer(w, res)
		return
	}
	id, err := strconv.ParseUint(p.ByName("id"), 10, 64)
	if err != nil {
		h.log.Println(err)
		res := response.Resp{
			Status: "error",
			Data:   "There was an problem, please try again",
		}
		response.Writer(w, res)
		return
	}
	todo, err := h.m.Update(id, dbmodel.Todo{Name: data.Name, Description: data.Desc})
	if err != nil {
		h.log.Println(err)
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
