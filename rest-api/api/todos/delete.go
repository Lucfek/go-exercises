package todos

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/lucfek/go-exercises/rest-api/response"
)

// Delete is responsible for handling "DELETE" Requests
func (h Handler) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
	todo, err := h.m.Delete(id)
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
