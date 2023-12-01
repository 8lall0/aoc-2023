package day_01

import (
	"bufio"
	"io"
	"strconv"
	"strings"
	"unicode"
)

func Trebuchet(r io.Reader, first bool) (int, error) {
	var buffer string
	sum := 0

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		buffer = scanner.Text()
		var v int
		var err error
		if first {
			v, err = ScanString(buffer)
		} else {
			v, err = ScanStringAsLetters(buffer)
		}
		if err != nil {
			return 0, err
		}
		sum += v
	}

	return sum, nil
}

func ScanString(input string) (int, error) {
	first := -1
	last := -1
	runes := []rune(input)

	for _, c := range runes {
		if unicode.IsDigit(c) {
			v, err := strconv.Atoi(string(c))
			if err != nil {
				return 0, err
			}

			if first == -1 {
				first = v
			} else {
				last = v
			}
		}
	}

	if last == -1 {
		last = first
	}

	return (first * 10) + last, nil
}

func ScanStringAsLetters(input string) (int, error) {
	type Eval struct {
		index int
		value int
	}
	numCmp := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	strCmp := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	first := Eval{
		index: -1,
		value: 0,
	}
	last := Eval{
		index: -1,
		value: 0,
	}

	// compare numbers
	for i, n := range numCmp {
		ndx := strings.Index(input, n)

		if ndx == -1 {
			continue
		} else if first.index == -1 || ndx < first.index {
			first.index = ndx
			first.value = i + 1
			if last.index == -1 {
				last.index = ndx
				last.value = i + 1
			}
		} else if ndx > last.index {
			last.index = ndx
			last.value = i + 1
		}
	}

	// compare strings
	for i, n := range strCmp {
		ndx := strings.Index(input, n)
		if ndx == -1 {
			continue
		} else if first.index == -1 || ndx < first.index {
			first.index = ndx
			first.value = i + 1

			if last.index == -1 {
				last.index = ndx
				last.value = i + 1
			}
		} else if ndx > last.index {
			last.index = ndx
			last.value = i + 1
		}
	}

	return (first.value * 10) + last.value, nil
}
