package main

import (
	"context"
	"log/slog"
	"os"

	_ "github.com/guil95/csv-parser/config"
	cliadapter "github.com/guil95/csv-parser/internal/parser/adapters/cli"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	logHandler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}).WithAttrs([]slog.Attr{slog.String("service", "parser")})

	slog.SetDefault(slog.New(logHandler))

	ctx := context.Background()
	cli := cliadapter.New()

	err := cli.Run(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to run cli adapter", "error", err)
		return
	}
}
