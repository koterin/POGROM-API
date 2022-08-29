package controller

import (
	"encoding/json"
	"net/http"

	"pogrom/internal/entity"
)

func sendResponse(w http.ResponseWriter, status int, msg string) {
	addBasicHeaders(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(entity.ServerResponse{
		Status:   status,
		Response: msg,
	})
}

func sendInfo(w http.ResponseWriter, i ...entity.Response) {
	addBasicHeaders(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(i)
}

func sendImage(w http.ResponseWriter, img []byte) {
	addBasicHeaders(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(img)
}
