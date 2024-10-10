package cli

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/fatih/color"
	fileadapter "github.com/guil95/csv-parser/internal/parser/adapters/file"
	port "github.com/guil95/csv-parser/internal/parser/ports/cli"
	"github.com/guil95/csv-parser/internal/parser/usecases"
)

type cli struct {
	filePath string
}

func New() port.CLI {
	filePath := flag.String("file", "", "Path to the input file")

	flag.Parse()

	if *filePath == "" {
		fmt.Println("Please provide a file using -file flag")
		return nil
	}

	return &cli{filePath: *filePath}
}

func (c cli) Run(ctx context.Context) error {
	color.Blue("parse csv file started")
	file, err := os.Open(c.filePath)
	if err != nil {
		slog.Error("Error opening file", "err", err)
		return nil
	}
	defer func() {
		file.Close()
		color.Blue("parse csv file done")
	}()

	reader := csv.NewReader(file)
	fileReader := fileadapter.NewCSVReader(reader)
	uc := usecases.NewParserUC(fileReader, fileadapter.NewCSVWriter())

	return uc.Parse(ctx)
}
