package utils

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ErrorWithPrefix(w http.ResponseWriter, err error, status int, prefix string) {
	JsonResponse(w, map[string]string{"error": prefix + ": " + err.Error()}, status)
}

func InternalServerError(w http.ResponseWriter, err error) {
	JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
}

func BadRequest(w http.ResponseWriter, err error) {
	JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
}

func NotFound(w http.ResponseWriter, err error) {
	JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusNotFound)
}

func Ok(w http.ResponseWriter, data interface{}) {
	JsonResponse(w, data, http.StatusOK)
}

func IsMultiPartForm(r *http.Request) bool {
	return r.Header.Get("Content-Type") == "multipart/form-data"
}

func IsPost(r *http.Request) bool {
	return r.Method == "POST"
}

func SetAccessControl(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
