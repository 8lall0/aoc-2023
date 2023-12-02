package day_02

import (
	"strings"
	"testing"
)

var test = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

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
}

func TestSackPick(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{test, 8},
	}

	for _, test := range tests {
		read := strings.NewReader(test.input)
		result, err := SackPick(read)
		if result != test.expected || err != nil {
			t.Errorf("For input %s, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}
