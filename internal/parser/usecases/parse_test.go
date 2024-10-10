package usecases

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/guil95/csv-parser/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	GenerateFileMethod  = "GenerateFile"
	GetNextRecordMethod = "GetNextRecord"
	GetHeaderMethod     = "GetHeader"
)

var ctx = context.Background()

func TestParserUC_Parse(t *testing.T) {
	tests := []struct {
		name         string
		header       []string
		records      [][]string
		expectError  bool
		readerError  error
		writerError  error
		expectedLogs string
	}{
		{
			name:        "successful parse with valid and invalid lines",
			header:      []string{"first", "email", "wage", "id"},
			records:     [][]string{{"John", "john@example.com", "1000", "1"}, {"Doe", "doe@example", "2000", "2"}},
			expectError: false,
		},
		{
			name:        "successful parse with valid and invalid lines with different header",
			header:      []string{"first", "e-mail", "salary", "id"},
			records:     [][]string{{"John", "john@example.com", "1000", "1"}, {"Doe", "doe@example", "2000", "2"}},
			expectError: false,
		},
		{
			name:        "error on header retrieval",
			header:      nil,
			records:     nil,
			readerError: errors.New("failed to get header"),
			expectError: true,
		},
		{
			name:        "error on getting next record",
			header:      []string{"first", "email", "wage", "id"},
			records:     nil,
			readerError: errors.New("failed to get record"),
			expectError: true,
		},
		{
			name:        "error on saving valid file",
			header:      []string{"first", "email", "wage", "id"},
			records:     [][]string{{"John", "john@example.com", "1000", "1"}},
			writerError: errors.New("failed to save file"),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockReader := new(mocks.ReaderFile)
			mockWriter := new(mocks.WriterFile)

			if tt.readerError != nil {
				mockReader.On(GetHeaderMethod, ctx).Return(nil, tt.readerError).Once()
			} else {
				mockReader.On(GetHeaderMethod, ctx).Return(tt.header, nil).Once()
				mockReader.On(GetNextRecordMethod, ctx).Return(tt.records[0], nil).Once()
				if len(tt.records) > 1 {
					mockReader.On(GetNextRecordMethod, ctx).Return(tt.records[1], io.EOF).Once()
				} else {
					mockReader.On(GetNextRecordMethod, ctx).Return(nil, io.EOF).Once()
				}
			}

			if tt.writerError != nil {
				mockWriter.On(GenerateFileMethod, mock.Anything, mock.Anything, mock.Anything).Return(tt.writerError).Once()
			} else {
				mockWriter.On(GenerateFileMethod, mock.Anything, mock.Anything, mock.Anything).Return(nil).Twice()
			}

			parserUC := NewParserUC(mockReader, mockWriter)

			err := parserUC.Parse(ctx)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
