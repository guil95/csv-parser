package parser

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/guil95/csv-parser/validator"
)

type Line struct {
	ID     string `json:"Employee ID" validate:"required,not_blank"`
	Name   string `json:"Employee name" validate:"required,not_blank"`
	Salary string `json:"Employee salary" validate:"required,not_blank"`
	Email  string `json:"Employee email" validate:"required,not_blank"`
}

func (l *Line) IsValid() bool {
	err := validator.Validate(l)
	return err == nil
}

type Parser struct {
	invalidLines []Line
	validLines   []Line
	emails       map[string]bool
}

const (
	InvalidFilesPath = "data/invalid_files"
	ValidFilesPath   = "data/valid_files"
)

func NewParser() *Parser {
	return &Parser{invalidLines: []Line{}, validLines: []Line{}, emails: make(map[string]bool)}
}

func (p *Parser) AddLine(line Line) {
	if line.Email == "" || !p.emails[line.Email] {
		if line.IsValid() {
			color.Green(fmt.Sprintf("valid line added to file %v", line))
			p.validLines = append(p.validLines, line)
		} else {
			color.Red(fmt.Sprintf("invalid line added to file %v", line))
			p.invalidLines = append(p.invalidLines, line)
		}

		p.emails[line.Email] = true
	}
}

func (p *Parser) TotalLines() int {
	var lines = append(p.invalidLines, p.validLines...)

	return len(lines)
}

func (p *Parser) RetrieveInvalidLines() []Line {
	return p.invalidLines
}

func (p *Parser) RetrieveValidLines() []Line {
	return p.validLines
}

func (p *Parser) CleanLines() {
	p.invalidLines = []Line{}
	p.validLines = []Line{}
}
