package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseString(t *testing.T) {
	tests := []struct {
		input    string
		expected []customChar
		isError  bool
	}{
		{
			input: `qwe\\\3`,
			expected: []customChar{
				{'q', false},
				{'w', false},
				{'e', false},
				{'\\', false},
				{'3', false},
			},
		},
		{
			input:    `qw\ne`,
			expected: []customChar{},
			isError:  true,
		},
		{
			input:    `\`,
			expected: []customChar{},
			isError:  true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := parseString(tc.input)
			if tc.isError {
				require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackStareks(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		isError  bool
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "aaa10b", expected: "", isError: true},
		{input: "a1b2", expected: "abb"},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
		{input: "a", expected: "a"},
		{input: "1", expected: "", isError: true},
		// // uncomment if task with asterisk completed
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		{input: `qwe\\5a`, expected: `qwe\\\\\a`},
		{input: `\3a`, expected: `3a`},
		{input: `\\3a`, expected: `\\\a`},
		{input: `\`, expected: "", isError: true},
		{input: "\\", expected: "", isError: true},
		{input: `qw\ne`, expected: "", isError: true},
		{input: `qw\te`, expected: "", isError: true},
		{input: `qw\t\ne`, expected: "", isError: true},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := UnpackStareks(tc.input)
			if tc.isError {
				require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tc.expected, result)
		})
	}
}
