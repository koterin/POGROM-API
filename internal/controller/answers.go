package controller

import (
    "encoding/json"
    "net/http"

    "pogrom/internal/entity"
)

func sendAnswer(w http.ResponseWriter, status int, msg string) {
    addBasicHeaders(w)
    w.WriteHeader(status)

    answer := entity.ServerResponse{
        Status:   status,
        Response: msg,
    }

    json.NewEncoder(w).Encode(answer)
}

func sendItem(w http.ResponseWriter, i entity.Items) {
    addBasicHeaders(w)
    w.WriteHeader(http.StatusOK)

    answer := entity.ServerResponse{
        Status:        http.StatusOK,
        ItemId:        i.Id,
        ItemDesc:      i.Desc,
        ItemTitle:     i.Title,
        ItemStatus:    i.Status,
        ItemCategory:  i.Category,
        ItemCreatedAt: i.CreatedAt,
    }

    json.NewEncoder(w).Encode(answer)
}
