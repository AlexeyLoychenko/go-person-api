package main

import (
	"log"

	"github.com/AlexeyLoychenko/person_api/config"
	"github.com/AlexeyLoychenko/person_api/internal/app"
)

func main() {
	//Init config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Can't read config!")
	}
	//Run app
	app.Run(cfg)

}
