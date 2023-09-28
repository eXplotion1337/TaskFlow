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

	storage, err := app.InitStorage(config)
	if err != nil {
		log.Fatal("fail create storage", err)
	}

	err = app.Run(config, storage)
	if err != nil {
		log.Fatal("fail run server", err)
	}

}
