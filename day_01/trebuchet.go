package day_01

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"unicode"
)

func Trebuchet(r io.Reader) (int, error) {
	var buffer string
	sum := 0

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		buffer = scanner.Text()
		v, err := ScanString(buffer)
		if err != nil {
			return 0, err
		}
		fmt.Println(v)
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
