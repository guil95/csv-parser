package file

import (
	"context"
	"github.com/guil95/csv-parser/internal/parser"
)

type ReaderFile interface {
	GetHeader(ctx context.Context) ([]string, error)
	GetNextRecord(ctx context.Context) ([]string, error)
}

type WriterFile interface {
	GenerateFile(ctx context.Context, lines []parser.Line, filePath string) error
}
