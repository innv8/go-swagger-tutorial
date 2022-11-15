package middleware

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, code int, response interface{}) {
	resBytes, _ := json.Marshal(response)
	w.Header().Set("content-type", "application/json")
	w.Header().Set("connection", "close")
	_, _ = w.Write(resBytes)
}
