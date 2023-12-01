package day_01

import (
	"strings"
	"testing"
)

func TestScanString(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}

	for _, test := range tests {
		result, err := ScanString(test.input)
		if result != test.expected || err != nil {
			t.Errorf("For input %s, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}

func TestTrebuchet(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`, 142},
	}

	for _, test := range tests {
		read := strings.NewReader(test.input)
		result, err := Trebuchet(read)
		if result != test.expected || err != nil {
			t.Errorf("For input %s, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}
