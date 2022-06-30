package main

import (
	"REST/internal/app"
	"REST/internal/config"
	"REST/pkg/logging"
)

func main() {

	cfg := config.GetConfig()
	logger := logging.GetLogger()

	app, err := app.NewApp(cfg, logger)
	if err != nil {
		logger.Fatalln(err.Error())
	}

	app.Run()
}
