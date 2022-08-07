package app

import (
        "fmt"

	log "github.com/sirupsen/logrus"
	"pogrom/config"
)

func Run() {
    log.Warn("We're here")

    log.Info("host is", config.Args.HOST_URL)

    fmt.Println("hi, 1+1=", 1+1)
}
