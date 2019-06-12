package users

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lucfek/go-exercises/rest-api/model"
	"github.com/lucfek/go-exercises/rest-api/response"
)

// Login logs in a user
func (h Handler) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := userData{}
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

	if data.Email == "" || data.Password == "" {
		res := response.Resp{
			Status: "error",
			Data:   "Empty values",
		}
		response.Writer(w, res)
		return
	}

	_, err = h.m.Login(model.User{Email: data.Email, Password: data.Password})
	if err, ok := err.(model.UserError); ok {
		h.log.Println(err)
		res := response.Resp{
			Status: "error",
			Data:   err.Msg,
		}
		response.Writer(w, res)
		return
	}
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
		Data:   true,
	}
	response.Writer(w, res)
}
