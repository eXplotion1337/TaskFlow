package main

import (
	"TaskFlow/internal/app"
	"log"
)

func main() {

	config, err := app.InitConfig()
	if err != nil {
		log.Fatal("fail load config", err)
	}

	err = app.Run(config)
	if err != nil {
		log.Fatal("fail run server", err)
	}

}
