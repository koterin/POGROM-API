package app

import (
	"database/sql"
	"net/http"
	"os"

	"pogrom/config"
	"pogrom/internal/controller"
	db "pogrom/internal/database"

	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"
)

func Run() {
	level, err := log.ParseLevel(config.Args.LOG_LEVEL)
	if err != nil {
		log.Fatal(err)
	}

	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(level)

	log.Info("host is", config.Args.HOST_URL)

	db_pass, err := os.ReadFile(config.Args.DB_PASSWORD_FILE)
	if err != nil {
		log.Fatal(err)
	}

	db.DB, err = sql.Open("postgres", config.ConfigDB(string(db_pass)))
	if err != nil {
		log.Fatal("00: DB credentials are invalid: ", err)
	}

	err = db.DB.Ping()
	if err != nil {
		log.Fatal("01: DB connection is not established, ping errored: ", err)
	}

	log.Info("Successfully connected to DB")

	http.HandleFunc("/api/health-check", controller.PostHealthCheck)
	http.HandleFunc("/api/items/", controller.GetItem)
	http.HandleFunc("/api/users/", controller.GetUser)

	log.Fatal(http.ListenAndServe(config.Args.PORT, nil))
}
