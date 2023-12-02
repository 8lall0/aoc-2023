package main

import (
	"aoc_2023/day_01"
	"fmt"
	"io"
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

	v, err := day_01.Trebuchet(f, true)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("First problem: ", v)
	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		log.Println(err)
	}
	v, err = day_01.Trebuchet(f, false)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Second problem: ", v)
}
