package controller

import (
    "net/http"

    "pogrom/internal/entity"

    log "github.com/sirupsen/logrus"
)

func PostHealthCheck(w http.ResponseWriter, req *http.Request) {
    var data entity.ClientRequest

    log.Info("POST /api/health-check")

    if err := ReadJson(req, &data); err != nil {
        sendAnswer(w, http.StatusBadRequest, "not ok")

        return
    }

    sendAnswer(w, http.StatusOK, "ok")
}

func GetItem(w http.ResponseWriter, req *http.Request) {
    var data entity.ClientRequest

    log.Info("GET /api/item")

    if err := ReadJson(req, &data); err != nil {
        sendAnswer(w, http.StatusBadRequest, "not ok")

        return
    }

    sendItem(w, entity.Items{"id1",
        "some desc",
        "testItem",
        "active",
        "ceramics",
        "12:34:34"})
}

func GetUser(w http.ResponseWriter, req *http.Request) {
    var data entity.ClientRequest

    log.Info("GET ", req.URL.Path)

    if err := ReadJson(req, &data); err != nil {
        sendAnswer(w, http.StatusBadRequest, "not ok")

        return
    }

    sendItem(w, entity.Items{"id1",
        "some desc",
        "testItem",
        "active",
        "ceramics",
        "12:34:34"})
}
