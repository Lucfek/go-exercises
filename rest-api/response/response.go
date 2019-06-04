package response

import (
	"encoding/json"
	"net/http"
)

type Resp struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func Writer(w http.ResponseWriter, resp Resp) {
	b, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
