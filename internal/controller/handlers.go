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
		sendResponse(w, http.StatusBadRequest, "not ok")

		return
	}

	sendResponse(w, http.StatusOK, "ok")
}

func GetItem(w http.ResponseWriter, req *http.Request) {
	var data entity.ClientRequest

	log.Info("GET /api/item")

	if err := ReadJson(req, &data); err != nil {
		sendResponse(w, http.StatusBadRequest, "not ok")

		return
	}

	sendInfo(w, entity.Item{"id1",
		"some desc",
		"testItem",
		"active",
		"ceramics",
		"12:34:34"})
}

func GetUser(w http.ResponseWriter, req *http.Request) {
	log.Info("GET ", req.URL.Path)

	userId := strings.TrimPrefix(req.URL.Path, "/api/users/")

	if userId == "2" {
		sendResponse(w, http.StatusBadRequest, "not ok")

		return
	}

	sendInfo(w, entity.UserInfo{"1", "username", "bidder",
		"active", "234", "3", "12:23:23"})
}
