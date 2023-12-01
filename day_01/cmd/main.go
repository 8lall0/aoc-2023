package main

import (
	day_01 "aoc_01"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Println(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)

	v, err := day_01.Trebuchet(f)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(v)
}
