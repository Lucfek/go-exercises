package handle

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lucfek/go-exercises/client-server/database"
)

type Storage interface {
	Get(id uint64) (database.User, error)
	Set(user database.User) database.User
	Delete(id uint64) (database.User, error)
}

type Handler struct {
	Storage Storage
}

type response struct {
	Succ bool
	Msg  string
	User database.User
}

func respWriter(w http.ResponseWriter, succ bool, msg string, user database.User) {
	resp := response{
		Succ: succ,
		Msg:  msg,
		User: user,
	}
	if data, err := json.Marshal(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}

}

// Set handles POST requests
func (h Handler) Set(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := database.User{
		Name:    r.FormValue("name"),
		Surname: r.FormValue("surname"),
		Email:   r.FormValue("email"),
	}
	respWriter(w, true, "Added succesfully", h.Storage.Set(user))
}

// Get handles GET requests
func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		respWriter(w, false, err.Error(), database.User{})
		return
	}
	user, err := h.Storage.Get(id)
	if err != nil {
		respWriter(w, false, err.Error(), database.User{})
		return
	}

	respWriter(w, true, "Obtaind succesfully", user)
}

// Del handles DELETE requests
func (h Handler) Del(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		respWriter(w, false, err.Error(), database.User{})
		return
	}
	user, err := h.Storage.Delete(id)
	if err != nil {
		respWriter(w, false, err.Error(), database.User{})
		return
	}
	respWriter(w, true, "Deleted succesfully", user)
}

func New(s Storage) Handler {
	return Handler{Storage: s}
}
