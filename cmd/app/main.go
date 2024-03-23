package main

import (
	"GolangQuizlet/internal/config"
	"GolangQuizlet/internal/storage"
	"fmt"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info("starting GolangQuizlet", slog.String("env", cfg.Env))
	log.Debug("Debug messages are enabled")

	_, err := storage.DbConnection()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Db Successfully connected")

	//init storage: postgres
	//init router : chi, chi render
	//run server :
}
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
