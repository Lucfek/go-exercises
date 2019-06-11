package usershandler

import (
	"encoding/json"
	"net/http"
	"regexp"
	"unicode"

	"github.com/julienschmidt/httprouter"
	"github.com/lib/pq"
	"github.com/lucfek/go-exercises/rest-api/response"
	"github.com/lucfek/go-exercises/rest-api/usersmodel"
	"golang.org/x/crypto/bcrypt"
)

var validEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func isValidPass(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 7 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Register registers a user
func (h Handler) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	if !isValidPass(data.Password) {
		res := response.Resp{
			Status: "error",
			Data:   "Weak password",
		}
		response.Writer(w, res)
		return
	}

	if !validEmail.MatchString(data.Email) {
		res := response.Resp{
			Status: "error",
			Data:   "Email not valid",
		}
		response.Writer(w, res)
		return
	}

	hash, err := hashPassword(data.Password)
	if err != nil {
		h.log.Println(err)
		res := response.Resp{
			Status: "error",
			Data:   "There was an problem, please try again",
		}
		response.Writer(w, res)
		return
	}

	user, err := h.m.Add(usersmodel.User{Email: data.Email, Password: hash})
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				h.log.Println(err)
				res := response.Resp{
					Status: "error",
					Data:   "User with this email already exists",
				}
				response.Writer(w, res)
				return
			}
		}
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
		Data:   user,
	}
	response.Writer(w, res)
}
