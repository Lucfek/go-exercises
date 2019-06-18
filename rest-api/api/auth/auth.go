package auth

import (
	"context"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/lucfek/go-exercises/rest-api/model"
	"github.com/lucfek/go-exercises/rest-api/response"
)

func JwtVerify(next func(w http.ResponseWriter, r *http.Request, p httprouter.Params)) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		header := r.Header.Get("Authorization") //Grab the token from the header
		AuthArr := strings.Split(header, " ")
		var token string
		if len(AuthArr) == 2 {
			token = AuthArr[1]
		} else {
			res := response.Resp{
				Status: "error",
				Data:   "Invalid token",
			}
			response.Writer(w, res)
			return
		}

		if token == "" {
			res := response.Resp{
				Status: "error",
				Data:   "Missing token",
			}
			response.Writer(w, res)
			return
		}
		tk := &model.Token{}

		_, err := jwt.ParseWithClaims(token, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			res := response.Resp{
				Status: "error",
				Data:   "There was an problem, please try again",
			}
			response.Writer(w, res)
			return
		}
		key := "user"
		ctx := context.WithValue(r.Context(), &key, tk)
		next(w, r.WithContext(ctx), p)
	}
}
