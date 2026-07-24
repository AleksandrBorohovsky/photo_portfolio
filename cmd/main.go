package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/AleksandrBorohovsky/photo_portfolio/internal/config"
	"github.com/AleksandrBorohovsky/photo_portfolio/internal/handlers"
	"github.com/AleksandrBorohovsky/photo_portfolio/pkg/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.NewLogger(cfg.Env)

	log.Info("config loaded", slog.Any("config", cfg))

	router := handlers.NewRouter()

	go func() {
		if err := http.ListenAndServe(fmt.Sprintf("localhost:%d", cfg.Port), router); err != nil {
			panic("Panic in server listener")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	sig := <-stop
	log.Info("stopping application", slog.String("signal", sig.String()))
}
