package usershandler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lucfek/go-exercises/rest-api/response"
	"golang.org/x/crypto/bcrypt"
)

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

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

	user, err := h.m.Get(data.Email)
	if err == sql.ErrNoRows {
		h.log.Println(err)
		res := response.Resp{
			Status: "error",
			Data:   "There is no user with such email",
		}
		response.Writer(w, res)
		return
	} else if err != nil {
		h.log.Println(err)
		res := response.Resp{
			Status: "error",
			Data:   "There was an problem, please try again",
		}
		response.Writer(w, res)
		return
	}

	if !checkPasswordHash(data.Password, user.Password) {
		res := response.Resp{
			Status: "error",
			Data:   "Wrong login data",
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
