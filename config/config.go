package config

import (
	"fmt"
	"net/url"
	"os"

	"github.com/alexflint/go-arg"
	log "github.com/sirupsen/logrus"
)

var Args struct {
	LOG_LEVEL        string `arg:"required,env"`
	PORT             string `arg:"required,env"`
	DB_PASSWORD_FILE string `arg:"required,env"`
	DB_PORT          int    `arg:"required,env"`
	DB_HOST          string `arg:"required,env"`
	DB_USER          string `arg:"required,env"`
	DB_NAME          string `arg:"required,env"`
	HOST_URL         string `arg:"required,env"`
}

func Validate() {
	if err := arg.Parse(&Args); err != nil {
		log.Fatal(err)
	}

	file, err := os.ReadFile(Args.DB_PASSWORD_FILE)
	if err != nil {
		log.Info("Path to or .db-secret itself is invalid")
		log.Fatal(err)
	}

	if len(file) == 0 {
		log.Fatal(".db-secret is empty")
	}

	_, err = url.Parse(Args.HOST_URL)
	if err != nil {
		log.Info("HOST_URL env is invalid. Required format is: https://test.domain.com")
		log.Fatal(err)
	}

	_, err = url.ParseRequestURI(Args.HOST_URL)
	if err != nil {
		log.Info("HOST_URL env is invalid. Required format is: https://test.domain.com")
		log.Fatal(err)
	}
}

func ConfigDB(pass string) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Args.DB_HOST, Args.DB_PORT, Args.DB_USER, pass, Args.DB_NAME)
}
