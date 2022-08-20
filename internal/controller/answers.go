package controller

import (
	"encoding/json"
	"net/http"

	"pogrom/internal/entity"
)

func sendResponse(w http.ResponseWriter, status int, msg string) {
	addBasicHeaders(w)
	w.WriteHeader(status)

	answer := entity.ServerResponse{
		Status:   status,
		Response: msg,
	}

	json.NewEncoder(w).Encode(answer)
}

func sendInfo(w http.ResponseWriter, i ...entity.Response) {
	addBasicHeaders(w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(i)
}
