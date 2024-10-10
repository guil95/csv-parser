package parser

import (
	"context"
)

type ReaderFile interface {
	GetHeader(ctx context.Context) ([]string, error)
	GetNextRecord(ctx context.Context) ([]string, error)
}

type WriterFile interface {
	GenerateFile(ctx context.Context, lines []Line, filePath string) error
}
