package controller

import (
	"net/http"
	"os"
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
	log.Info("GET ", req.URL.Path)

	itemId := strings.TrimPrefix(req.URL.Path, "/api/items/")

	itemInfo, err := db.GetItemInfo(itemId)
	if err != nil {
		log.Error("Error in GetItem: ", err)
		sendResponse(w, http.StatusBadRequest, "No such item")

		return
	}

	log.Debug(itemInfo)
	sendInfo(w, itemInfo)
}

func GetImage(w http.ResponseWriter, req *http.Request) {
	log.Info("GET ", req.URL.Path)

	imageId := strings.TrimPrefix(req.URL.Path, "/api/images/")

	img, err := os.ReadFile("/images/" + imageId + ".png")
	if err != nil {
		log.Error("Error in GetItem: ", err)
		sendResponse(w, http.StatusBadRequest, "Error getting image")

		return
	}

	sendImage(w, img)
}

func GetUser(w http.ResponseWriter, req *http.Request) {
	log.Info("GET ", req.URL.Path)

	username := strings.TrimPrefix(req.URL.Path, "/api/users/")

	userInfo, err := db.FindUser(username)
	if err != nil {
		log.Error("Error in GetUser: ", err)
		sendResponse(w, http.StatusBadRequest, "No such user")

		return
	}

	log.Debug(userInfo)
	sendInfo(w, userInfo)
}
