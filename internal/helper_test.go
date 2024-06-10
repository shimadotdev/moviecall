package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatGenres(t *testing.T) {

	tests := []struct {
		name     string
		genres   []struct{ Name string }
		expected string
	}{
		{
			name:     "single genre",
			genres:   []struct{ Name string }{{Name: "Action"}},
			expected: "Action",
		},
		{
			name:     "multiple genres",
			genres:   []struct{ Name string }{{Name: "Action"}, {Name: "Comedy"}, {Name: "Drama"}},
			expected: "Action, Comedy, Drama",
		},
		{
			name:     "no genres",
			genres:   []struct{ Name string }{},
			expected: "",
		},
		{
			name:     "genre with spaces",
			genres:   []struct{ Name string }{{Name: "Science Fiction"}},
			expected: "Science Fiction",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatGenres(tt.genres)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestEllipsizeString(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		maxLen   int
		expected string
	}{
		{
			name:     "string shorter than max length",
			str:      "Hello",
			maxLen:   10,
			expected: "Hello",
		},
		{
			name:     "string exactly max length",
			str:      "HelloWorld",
			maxLen:   10,
			expected: "HelloWorld",
		},
		{
			name:     "string longer than max length",
			str:      "HelloWorld",
			maxLen:   5,
			expected: "Hello...",
		},
		{
			name:     "empty string",
			str:      "",
			maxLen:   5,
			expected: "",
		},
		{
			name:     "max length zero",
			str:      "Hello",
			maxLen:   0,
			expected: "...",
		},
		{
			name:     "max length less than string length with spaces",
			str:      "Hello World",
			maxLen:   5,
			expected: "Hello...",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EllipsizeString(tt.str, tt.maxLen)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestConvertString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "single word",
			input:    "Hello",
			expected: "Hello",
		},
		{
			name:     "multiple words",
			input:    "Hello World",
			expected: "Hello+World",
		},
		{
			name:     "leading and trailing spaces",
			input:    "  Hello World  ",
			expected: "Hello+World",
		},
		{
			name:     "multiple spaces between words",
			input:    "Hello   World",
			expected: "Hello+World",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertString(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
