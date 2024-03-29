package main

import (
	"GolangQuizlet/internal/app/handler"
	"GolangQuizlet/internal/app/repository"
	"GolangQuizlet/internal/app/service"
	"GolangQuizlet/internal/config"
	"GolangQuizlet/pkg/db"
	mwLogger "GolangQuizlet/pkg/logger"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
	//pool:=db.NewPool(cfg.)
	log := setupLogger(cfg.Env)
	dbConnect, err := db.DbConnection()
	if err != nil {
		log.Error("Error connect to db", err)
		os.Exit(1)
	}
	defer dbConnect.Close()

	quizRepo := repository.NewQuizRepository(dbConnect)
	quizService := service.NewQuizService(quizRepo)
	quizHandler := handler.NewQuizHandler(quizService)

	http.HandleFunc("/next", quizHandler.GetNextQuestion)
	http.HandleFunc("/check", quizHandler.CheckAnswer)

	fmt.Println("server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server: %v", err)
	}

	log.Info("starting GolangQuizlet", slog.String("env", cfg.Env))
	log.Debug("Debug messages are enabled")

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	//run server :
}
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
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
