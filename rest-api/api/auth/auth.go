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

		var header = r.Header.Get("x-access-token") //Grab the token from the header

		header = strings.TrimSpace(header)

		if header == "" {
			res := response.Resp{
				Status: "error",
				Data:   "Missing token",
			}
			response.Writer(w, res)
			return
		}
		tk := &model.Token{}

		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
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
