package api

import (
	"encoding/json"
	"net/http"
)

func CommandFromBody[TReq any](r *http.Request) (TReq, error) {
	var req TReq
	return req, json.NewDecoder(r.Body).Decode(&req)
}

func JSONResponse(w http.ResponseWriter, status int, data Json) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

type Json map[string]interface{}
