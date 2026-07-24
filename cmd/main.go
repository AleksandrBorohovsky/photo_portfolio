package main

import (
	"log/slog"

	"github.com/AleksandrBorohovsky/photo_portfolio/internal/config"
	"github.com/AleksandrBorohovsky/photo_portfolio/pkg/logger"
)

func main() {
	cfg := config.MustLoad()
	log := logger.NewLogger(cfg.Env)
	log.Info("config loaded", slog.Any("config", cfg))
}
