package main

import (
        log "github.com/sirupsen/logrus"
    "pogrom/config"
    "pogrom/internal/app"
)

func main() {
    config.Validate()

    level, err := log.ParseLevel(config.Args.LOG_LEVEL)
    if err != nil {
        log.Fatal(err)
    }

//    log.SetFormatter(&log.JSONFormatter{})
    log.SetLevel(level)

    app.Run()
}
