package main

import (
	"fmt"
	"github.com/{{cookiecutter.github_account}}/{{cookiecutter.github_repo_name}}/cmd/bootstrap"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/{{cookiecutter.github_account}}/{{cookiecutter.github_repo_name}}/infra/database"
)

func main() {
	router := gin.Default()

	logHandler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}).WithAttrs([]slog.Attr{slog.String("service", "foods")})

	slog.SetDefault(slog.New(logHandler))

	db := database.NewPostgresDB()
	bootstrap.Domain(router, db)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	slog.Info(fmt.Sprintf("Listening on port %s", s.Addr))

	err := s.ListenAndServe()
	if err != nil {
		return
	}
}
