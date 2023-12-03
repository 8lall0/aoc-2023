package day_03

import (
	"strings"
	"testing"
)

var test = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestCountDigits(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{1024, 4},
		{1, 1},
	}

	for _, test := range tests {
		result := countDigits(test.input)
		if result != test.expected {
			t.Errorf("For input %d, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestReadLine(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{"467..114..", []int{467, 467, 467, DOT, DOT, 114, 114, 114, DOT, DOT}},
		{"...*......", []int{DOT, DOT, DOT, SYMBOL, DOT, DOT, DOT, DOT, DOT, DOT}},
		{"..35..633.", []int{DOT, DOT, 35, 35, DOT, DOT, 633, 633, 633, DOT}},
		{"......#...", []int{DOT, DOT, DOT, DOT, DOT, DOT, SYMBOL, DOT, DOT, DOT}},
	}

	for _, test := range tests {
		result := readLine(test.input, len(test.input))
		if len(result) != len(test.expected) {
			t.Errorf("For input %s, expected %v, but got %v", test.input, test.expected, result)
		} else {
			for i := 0; i < len(result); i++ {
				if result[i] != test.expected[i] {
					t.Errorf("For input %s, expected %v, but got %v", test.input, test.expected, result)
					break
				}
			}
		}

	}
}

func TestEngineMatrix(t *testing.T) {
	tests := []struct {
		input     string
		expected  int
		expected2 int
	}{
		{test, 4361, 467835},
	}

	for _, test := range tests {
		read := strings.NewReader(test.input)
		result, result2, _ := EngineMatrix(read)
		if result != test.expected {
			t.Errorf("For input %s, expected %v, but got %v", test.input, test.expected, result)
		}
		if result2 != test.expected2 {
			t.Errorf("For input %s, expected %v, but got %v", test.input, test.expected2, result2)
		}
	}
}

/*
func TestGameParser(t *testing.T) {
	tests := []struct {
		input    string
		expected game
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", game{game: 1, red: 4, blue: 6, green: 2}},
	}

	for _, test := range tests {
		result, err := gameParser(test.input)
		if result != test.expected || err != nil {
			t.Errorf("For input %s, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}*/
