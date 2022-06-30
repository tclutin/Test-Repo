package app

import (
	"REST/internal/config"
	"REST/internal/user"
	"REST/pkg/logging"
	"net/http"
	"time"
)

type App struct {
	router *http.ServeMux
	server *http.Server
	logger *logging.Logger
	cfg    *config.Config
}

func NewApp(config *config.Config, logger *logging.Logger) (App, error) {
	logger.Info("router initializing")
	router := http.NewServeMux()

	handler := user.NewHandler()
	handler.Register(router)

	return App{
		router: router,
		logger: logger,
		cfg:    config,
	}, nil
}

func (a *App) startHTTP() {
	a.logger.Info("start HTTP")

	a.server = &http.Server{
		Addr:              ":" + a.cfg.Listen.Port,
		Handler:           a.router,
		MaxHeaderBytes:    1 << 20,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	a.logger.Fatalln(a.server.ListenAndServe())
}

func (a *App) Run() {
	a.startHTTP()
}
