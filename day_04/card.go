package day_04

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type ScratchCard struct {
	Winning   []int
	Yours     []int
	Instances int
}

func (s *ScratchCard) Evaluate() (int, int) {
	cnt := 0
	mul := 1

	for _, x := range s.Winning {
		for _, y := range s.Yours {
			if x == y {
				cnt++
				break
			}
		}
	}

	if cnt == 0 {
		return 0, 0
	}

	for i := 1; i < cnt; i++ {
		mul *= 2
	}

	return mul, cnt
}

func ScratchCardEvaluate(r io.Reader) (int, int, error) {
	c, err := loadCards(r)
	if err != nil {
		return 0, 0, nil
	}
	sum := 0

	for i, card := range c {
		x, y := card.Evaluate()
		sum += x
		for j := 1; j <= y && (i+j) < len(c); j++ {
			c[i+j].Instances += c[i].Instances
		}
	}
	sum2 := 0
	for _, card := range c {
		sum2 += card.Instances
	}

	return sum, sum2, nil
}

func loadCards(r io.Reader) ([]ScratchCard, error) {
	cards := make([]ScratchCard, 0)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		buffer := scanner.Text()
		ndx := strings.Index(buffer, ":")

		tmp := strings.Split(buffer[ndx+1:], "|")
		winNums := strings.Split(tmp[0], " ")
		yourNums := strings.Split(tmp[1], " ")
		winning := make([]int, 0)
		yours := make([]int, 0)

		for i := 0; i < len(winNums); i++ {
			if winNums[i] == "" {
				continue
			}
			x, err := strconv.Atoi(winNums[i])
			if err != nil {
				return nil, err
			}
			winning = append(winning, x)
		}

		for i := 0; i < len(yourNums); i++ {
			if yourNums[i] == "" {
				continue
			}
			y, err := strconv.Atoi(yourNums[i])
			if err != nil {
				return nil, err
			}
			yours = append(yours, y)
		}

		cards = append(cards, ScratchCard{Winning: winning, Yours: yours, Instances: 1})
	}

	return cards, nil
}
