package app

import (
    "net/http"
    // "database/sql"
    "os"
    "os/signal"
    "syscall"

    log "github.com/sirupsen/logrus"
    "pogrom/config"
    "pogrom/internal/controller"
)

func Run() {
    log.Info("host is", config.Args.HOST_URL)

 /*   db_pass,err := os.ReadFile(service.Args.DB_PASSWORD_FILE)
    if err != nil {
        log.Fatal(err)
    }

    dbhandler.DB, err = sql.Open("postgres", service.ConfigDB(string(db_pass)))
    if err != nil {
        log.Fatal("00: DB credentials are invalid: ", err)
    }

    err = dbhandler.DB.Ping()
    if err != nil {
        log.Fatal("01: DB connection is not established, ping errored: ", err)
    }

    log.Println("Successfully connected to DB")
*/
    shutdown := make(chan os.Signal, 1)
    signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

    http.HandleFunc("/api/health-check", controller.PostHealthCheck)

    log.Fatal(http.ListenAndServe(config.Args.PORT, nil))

    <-shutdown
    log.Println("Server gracefully shut down")
}
