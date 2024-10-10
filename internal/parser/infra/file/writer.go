package file

import (
	"context"
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/guil95/csv-parser/internal/parser"
)

type csvWriter struct {
}

func NewCSVWriter() parser.WriterFile {
	return &csvWriter{}
}

func (c csvWriter) GenerateFile(ctx context.Context, lines []parser.Line, filePath string) error {
	filename := fmt.Sprintf("employees_%s_%v.csv", time.Now().Format("20060102_150405"), time.Now().Nanosecond())
	fullPath := filepath.Join(filePath, filename)

	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		slog.ErrorContext(ctx, "error creating directory", slog.String("path", fullPath), slog.String("err", err.Error()))
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	slog.Debug(fmt.Sprintf("generate file %s", fullPath))

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Employee name", "Employee salary", "Employee email", "Employee ID"}
	if err := writer.Write(header); err != nil {
		slog.ErrorContext(ctx, "failed to write csv header", "error", err)
		return err
	}

	for _, line := range lines {
		row := []string{line.Name, line.Salary, line.Email, line.ID}
		if err := writer.Write(row); err != nil {
			slog.ErrorContext(ctx, "failed to write csv row", "error", err)
			return err
		}

		slog.Debug(fmt.Sprintf("saving row %v", row))
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return err
		slog.ErrorContext(ctx, "failed to flush", "error", err)
	}

	return nil
}
