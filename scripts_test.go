package main

import (
	"testing"
)

func TestGetFileExtension(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"test.py", "py"},
		{"test.txt", "txt"},
		{"test", ""},
		{"oh.my.god", "god"},
	}

	for _, testCase := range cases {
		result := getFileExtension(testCase.input)
		if result != testCase.expected {
			t.Errorf("getFileExtension(%q) = %q; expected %q", testCase.input, result, testCase.expected)
		}
	}
}
