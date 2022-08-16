package main

import (
    "os"
    "os/signal"
    "syscall"

    log "github.com/sirupsen/logrus"
    "pogrom/config"
    "pogrom/internal/app"
)

func main() {
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)

    config.Validate()

    level, err := log.ParseLevel(config.Args.LOG_LEVEL)
    if err != nil {
        log.Fatal(err)
    }

//    log.SetFormatter(&log.JSONFormatter{})
    log.SetLevel(level)

    app.Run()

    <-c
    log.Println("Server gracefully shut down")
}
