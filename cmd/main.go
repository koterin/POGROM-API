package main

import (
	"pogrom/config"
	"pogrom/internal/app"
)

func main() {
	config.Validate()

	app.Run()
}
