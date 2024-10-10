package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	filerepository "github.com/guil95/csv-parser/internal/parser/infra/file"
	"github.com/guil95/csv-parser/internal/parser/usecases"
	"log/slog"
	"os"

	_ "github.com/guil95/csv-parser/config"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	logHandler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}).WithAttrs([]slog.Attr{slog.String("service", "parser")})

	slog.SetDefault(slog.New(logHandler))

	filePath := flag.String("file", "", "Path to the input file")

	flag.Parse()

	if *filePath == "" {
		fmt.Println("Please provide a file using -file flag")
		return
	}

	file, err := os.Open(*filePath)
	if err != nil {
		slog.Error("Error opening file", "err", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	repo := filerepository.NewCSVReader(reader)
	uc := usecases.NewParserUC(repo, filerepository.NewCSVWriter())

	err = uc.Parse(context.Background())
	if err != nil {
		return
	}
}
