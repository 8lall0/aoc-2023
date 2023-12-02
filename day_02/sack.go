package day_02

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type game struct {
	game  int
	red   int
	green int
	blue  int
}

func (g *game) isValid() bool {
	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	return g.red <= maxRed && g.blue <= maxBlue && g.green <= maxGreen

}
func SackPick(r io.Reader) (int, int, error) {
	var buffer string
	sum := 0
	sum2 := 0

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		buffer = scanner.Text()
		g, err := gameParser(buffer)
		if err != nil {
			return 0, 0, err
		}
		if g.isValid() {
			sum += g.game
		}

		sum2 += g.red * g.blue * g.green
	}

	return sum, sum2, nil
}

func gameParser(input string) (game, error) {
	gameStruct := game{}

	ndx := strings.Index(input, ":")
	if ndx == -1 {
		return game{}, fmt.Errorf("invalid input string")
	}

	_, err := fmt.Sscanf(input[:ndx+1], "Game %d:", &gameStruct.game)
	if err != nil {
		return game{}, err
	}

	tries := strings.Split(input[ndx+1:], ";")
	for _, s := range tries {
		extractions := strings.Split(s, ",")
		for _, e := range extractions {
			var num int
			var color string
			_, err := fmt.Sscanf(e, "%d %s", &num, &color)
			if err != nil {
				return game{}, err
			}
			switch color {
			case "blue":
				if gameStruct.blue < num {
					gameStruct.blue = num
				}
			case "red":
				if gameStruct.red < num {
					gameStruct.red = num
				}
			case "green":
				if gameStruct.green < num {
					gameStruct.green = num
				}
			default:
				return game{}, fmt.Errorf("invalid color")
			}
		}
	}

	return gameStruct, nil
}
