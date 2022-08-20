package controller

import (
	"net/http"
	"strings"

	"pogrom/internal/entity"

	log "github.com/sirupsen/logrus"
)

func PostHealthCheck(w http.ResponseWriter, req *http.Request) {
	var data entity.ClientRequest

	log.Info("POST /api/health-check")

	if err := ReadJson(req, &data); err != nil {
		log.Error("Health-check Error")
		sendResponse(w, http.StatusBadRequest, "error")

		return
	}

	sendResponse(w, http.StatusOK, "ok")
}

func GetItem(w http.ResponseWriter, req *http.Request) {
	var data entity.ClientRequest

	log.Info("GET /api/item")

	if err := ReadJson(req, &data); err != nil {
		sendResponse(w, http.StatusBadRequest, "error")

		return
	}

	// log.Debug()
	sendInfo(w, entity.Item{
		Id:        "1",
		Desc:      "description",
		Title:     "title",
		Status:    "active",
		Category:  "category",
		CreatedAt: "12 march 12:34:45"})
}

func GetUser(w http.ResponseWriter, req *http.Request) {
	log.Info("GET ", req.URL.Path)

	userId := strings.TrimPrefix(req.URL.Path, "/api/users/")

	if userId == "2" {
		sendResponse(w, http.StatusBadRequest, "error")

		return
	}

	// log.Debug()
	sendInfo(w, entity.UserInfo{
		Id:        "1",
		Username:  "username",
		Role:      "role",
		Status:    "active",
		Bidded:    "123.45",
		BidsWon:   "14",
		CreatedAt: "12:23:23"})
}
