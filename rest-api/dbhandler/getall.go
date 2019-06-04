package dbhandler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lucfek/go-exercises/rest-api/response"
)

// GetAll is responsible for handling "GETALL" Requests
func (h Handler) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todos, err := h.m.GetAll()
	if err != nil {
		h.errLog.Println(err)
		res := response.Resp{
			Status: "error",
			Data:   "There was an problem, please try again",
		}
		response.Writer(w, res)
		return
	}
	res := response.Resp{
		Status: "succes",
		Data:   todos,
	}
	response.Writer(w, res)
}
