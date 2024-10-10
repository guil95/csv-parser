package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLine_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		line     Line
		expected bool
	}{
		{
			name:     "Valid Line",
			line:     Line{ID: "1", Name: "John Doe", Salary: "50000", Email: "john@example.com"},
			expected: true,
		},
		{
			name:     "Missing Name",
			line:     Line{ID: "2", Salary: "60000", Email: "jane@example.com"},
			expected: false,
		},
		{
			name:     "Missing Email",
			line:     Line{ID: "3", Name: "Jane Doe", Salary: "70000"},
			expected: false,
		},
		{
			name:     "Empty Salary",
			line:     Line{ID: "4", Name: "Alice", Email: "alice@example.com"},
			expected: false,
		},
		{
			name:     "Blank Fields",
			line:     Line{ID: "", Name: "", Salary: "", Email: ""},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.line.IsValid())
		})
	}
}

func TestParser_AddLine(t *testing.T) {
	p := NewParser()

	tests := []struct {
		name      string
		line      Line
		valid     bool
		total     int
		invalid   int
		validated int
	}{
		{
			name:      "Add valid line",
			line:      Line{ID: "1", Name: "John Doe", Salary: "50000", Email: "john@example.com"},
			valid:     true,
			total:     1,
			invalid:   0,
			validated: 1,
		},
		{
			name:      "Add invalid line (missing email)",
			line:      Line{ID: "2", Name: "Jane Doe", Salary: "60000"},
			valid:     false,
			total:     2,
			invalid:   1,
			validated: 1,
		},
		{
			name:      "Add duplicate email",
			line:      Line{ID: "3", Name: "Alice", Salary: "70000", Email: "john@example.com"},
			valid:     false,
			total:     2,
			invalid:   1,
			validated: 1,
		},
		{
			name:      "Add another valid line",
			line:      Line{ID: "4", Name: "Bob", Salary: "80000", Email: "bob@example.com"},
			valid:     true,
			total:     3,
			invalid:   1,
			validated: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p.AddLine(tt.line)

			assert.Equal(t, tt.total, p.TotalLines())
			assert.Equal(t, tt.invalid, len(p.RetrieveInvalidLines()))
			assert.Equal(t, tt.validated, len(p.RetrieveValidLines()))
		})
	}
}

func TestParser_CleanLines(t *testing.T) {
	p := NewParser()

	p.AddLine(Line{ID: "1", Name: "John Doe", Salary: "50000", Email: "john@example.com"})
	p.AddLine(Line{ID: "2", Name: "Jane Doe", Salary: "60000"})

	assert.Equal(t, 2, p.TotalLines())

	p.CleanLines()

	assert.Equal(t, 0, p.TotalLines())
	assert.Empty(t, p.RetrieveInvalidLines())
	assert.Empty(t, p.RetrieveValidLines())
}
