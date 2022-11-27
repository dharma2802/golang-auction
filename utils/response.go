package utils

import (
	"encoding/json"
	"net/http"

	"github.com/seller-app/auction/entities"
)

func InternalError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	var response entities.DefaultResponse
	w.WriteHeader(http.StatusInternalServerError)
	response.Status = http.StatusInternalServerError
	response.Message = message
	json.NewEncoder(w).Encode(response)
}

func NotFound(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	var response entities.DefaultResponse
	w.WriteHeader(http.StatusNotFound)
	response.Status = http.StatusNotFound
	response.Message = message
	json.NewEncoder(w).Encode(response)
}

func BadRequest(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	var response entities.DefaultResponse
	w.WriteHeader(http.StatusBadRequest)
	response.Status = http.StatusBadRequest
	response.Message = message
	json.NewEncoder(w).Encode(response)
}

func Ok(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
