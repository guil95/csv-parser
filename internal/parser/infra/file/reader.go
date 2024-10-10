package file

import (
	"context"
	"encoding/csv"
	"io"
	"log/slog"

	"github.com/guil95/csv-parser/internal/parser"
)

type csvReader struct {
	reader *csv.Reader
}

func NewCSVReader(r *csv.Reader) parser.ReaderFile {
	return &csvReader{reader: r}
}

func (c csvReader) GetHeader(ctx context.Context) ([]string, error) {
	header, err := c.reader.Read()
	if err != nil {
		slog.ErrorContext(ctx, "error reading header", "error", err)
		return nil, err
	}

	return header, nil
}

func (c csvReader) GetNextRecord(ctx context.Context) ([]string, error) {
	record, err := c.reader.Read()

	if err == io.EOF {
		slog.DebugContext(ctx, "EOF file", "error", err)
		return nil, err
	}

	if err != nil {
		slog.ErrorContext(ctx, "error reading record", "error", err)
		return nil, err
	}

	return record, nil
}

func (c csvReader) GenerateFile(ctx context.Context) error {
	return nil
}
