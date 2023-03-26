package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/lekht/news-formatter-service/config"
	"github.com/lekht/news-formatter-service/internal/api"
	"github.com/lekht/news-formatter-service/internal/format"
	"github.com/lekht/news-formatter-service/pkg/server"
)

func Run(cfg *config.Config) {
	f := format.New()
	api := api.New(f)
	router := api.Router()
	httpServer := server.New(router, server.Port(cfg.Server.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println(fmt.Errorf("app - Run - signal: " + s.String()))
	case err := <-httpServer.Notify():
		log.Println(fmt.Errorf("app - Run - server.Notify: %w", err))
	}

	err := httpServer.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("app - Run - server.Shutdown: %w", err))
	}
}
