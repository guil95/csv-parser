package usecases

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"strings"

	"github.com/guil95/csv-parser/config"
	"github.com/guil95/csv-parser/internal/parser"
	"golang.org/x/sync/errgroup"
)

type (
	ParserUC interface {
		Parse(ctx context.Context) error
	}

	parserUC struct {
		reader      parser.ReaderFile
		writer      parser.WriterFile
		parserModel parser.ParserModel
	}
)

const (
	NameField = iota
	EmailField
	SalaryField
	IDField
)

var (
	nameFields   = map[string]bool{"name": true, "first": true, "n": true, "first name": true, "f. name": true}
	emailFields  = map[string]bool{"email": true, "e-mail": true}
	salaryFields = map[string]bool{"salary": true, "wage": true, "rate": true}
	idFields     = map[string]bool{"number": true, "id": true, "employee number": true, "emp id": true}
	fieldIndexes = map[int]int{NameField: -1, EmailField: -1, SalaryField: -1, IDField: -1}
)

func NewParserUC(reader parser.ReaderFile, writer parser.WriterFile) ParserUC {
	return &parserUC{
		reader:      reader,
		writer:      writer,
		parserModel: parser.NewParser(),
	}
}

func (uc parserUC) Parse(ctx context.Context) error {
	header, err := uc.reader.GetHeader(ctx)
	if err != nil {
		slog.Error("create parser error", "error", err)
		return err
	}

	err = uc.buildFieldIndexes(header)
	if err != nil {
		slog.Error("build field index error", "error", err)
		return err
	}

	for {
		line, err := uc.reader.GetNextRecord(ctx)
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		uc.parserModel.AddLine(parser.Line{
			ID:     strings.ReplaceAll(line[fieldIndexes[IDField]], " ", ""),
			Name:   strings.ReplaceAll(line[fieldIndexes[NameField]], " ", ""),
			Salary: strings.ReplaceAll(line[fieldIndexes[SalaryField]], " ", ""),
			Email:  strings.ReplaceAll(line[fieldIndexes[EmailField]], " ", ""),
		})

		if uc.parserModel.TotalLines() >= config.AppConfig.MaxFileLength {
			if err := uc.saveFile(ctx); err != nil {
				slog.ErrorContext(ctx, "save files error", "error", err)
				return err
			}
		}
	}

	err = uc.saveFile(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "save files error", "error", err)
		return err
	}

	return nil
}

func (uc parserUC) saveFile(ctx context.Context) error {
	eg, errCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		validLines := uc.parserModel.RetrieveValidLines()
		if len(validLines) > 0 {
			err := uc.writer.GenerateFile(errCtx, validLines, parser.ValidFilesPath)
			if err != nil {
				return err
			}
		}

		return nil
	})

	eg.Go(func() error {
		invalidLines := uc.parserModel.RetrieveInvalidLines()
		if len(invalidLines) > 0 {
			err := uc.writer.GenerateFile(errCtx, invalidLines, parser.InvalidFilesPath)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	uc.parserModel.CleanLines()

	return nil
}

func (uc parserUC) buildFieldIndexes(header []string) error {
	for i, field := range header {
		field = removeZWNBSP(strings.ToLower(strings.TrimSpace(field)))
		if nameFields[field] {
			fieldIndexes[NameField] = i
		}

		if emailFields[field] {
			fieldIndexes[EmailField] = i
		}

		if salaryFields[field] {
			fieldIndexes[SalaryField] = i
		}

		if idFields[field] {
			fieldIndexes[IDField] = i
		}
	}

	for _, v := range fieldIndexes {
		if v < 0 {
			return errors.New("invalid field found")
		}
	}

	return nil
}

func removeZWNBSP(field string) string {
	return strings.ReplaceAll(field, "\uFEFF", "")
}
