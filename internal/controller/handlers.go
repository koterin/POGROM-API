package controller

import (
	"net/http"
	"strings"

	db "pogrom/internal/database"
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
	//	var data entity.ClientRequest

	log.Info("GET ", req.URL.Path)

	/*
		if err := ReadJson(req, &data); err != nil {
			sendResponse(w, http.StatusBadRequest, "error")

			return
		}
	*/

	itemId := strings.TrimPrefix(req.URL.Path, "/api/items/")

	// TODO: log.Debug()
	sendInfo(w, entity.Item{
		Id:        itemId,
		CreatedAt: "12 march 12:34:45"})
}

func GetUser(w http.ResponseWriter, req *http.Request) {
	log.Info("GET ", req.URL.Path)

	username := strings.TrimPrefix(req.URL.Path, "/api/users/")

	userInfo, err := db.FindUser(username)
	// TODO: Custom error here
	if err != nil {
		log.Error(err)
		sendResponse(w, http.StatusBadRequest, "error")

		return
	}

	// log.Debug()
	sendInfo(w, userInfo)
}
