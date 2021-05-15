package services

import (
	"encoding/json"
	"net/http"
)

func handleResponse(w http.ResponseWriter, b []byte, statusCode int) {
	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(statusCode)
	w.Write(b)
}

func marshalFormat(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}
