package main

import (
	"aoc_2023/day_04"
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

	res, res2, _ := day_04.ScratchCardEvaluate(f)

	fmt.Println("First problem:", res)
	fmt.Println("Second problem:", res2)
}
