package day_03

import (
	"bufio"
	"io"
	"strconv"
)

const (
	DOT    = -1
	SYMBOL = -2
	ENGINE = -3
)

func EngineMatrix(r io.Reader) (int, int, error) {
	sum := 0
	sum2 := 0

	mtx := loadMatrix(r)
	col, row := len(mtx), len(mtx[0])

	for i := 0; i < col; i++ {
		lastDigit := 0
		for j := 0; j < row; j++ {
			if mtx[i][j] < 0 {
				lastDigit = 0
				continue
			}

			if mtx[i][j] == lastDigit {
				continue
			}

			if checkAdjacency(mtx, row, col, i, j) {
				sum += mtx[i][j]
				lastDigit = mtx[i][j]
			}

		}
	}

	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			if mtx[i][j] == ENGINE {
				sum2 += getAdjacencyEngine(mtx, row, col, i, j)
			}
		}
	}

	return sum, sum2, nil
}

func loadMatrix(r io.Reader) [][]int {
	mtx := make([][]int, 0)

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	buffer := scanner.Text()

	n := len(buffer)
	ndx := 1

	mtx = append(mtx, readLine(buffer, n))
	for scanner.Scan() {
		buffer := scanner.Text()
		mtx = append(mtx, readLine(buffer, n))
		ndx++
	}

	return mtx
}

func readLine(input string, n int) []int {
	row := make([]int, n)
	runes := []rune(input)

	for i := 0; i < len(runes); i++ {
		switch runes[i] {
		case '.':
			row[i] = DOT
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			var j, tmpNum int
			for j, tmpNum = 0, 0; i+j < len(runes) && (runes[i+j] >= '0' && runes[i+j] <= '9'); j++ {
				x, err := strconv.Atoi(string(runes[i+j]))
				if err != nil {
					panic(err)
				}
				tmpNum = tmpNum*10 + x
			}

			for j = 0; i+j < len(runes) && (runes[i+j] >= '0' && runes[i+j] <= '9'); j++ {
				row[i+j] = tmpNum
			}

			i += j - 1
		case '*':
			row[i] = ENGINE
		default:
			row[i] = SYMBOL
		}
	}

	return row
}

func checkAdjacency(mtx [][]int, row, col, i, j int) bool {
	var startRow, startCol, endRow, endCol int

	if i == 0 {
		startCol = i
		endCol = i + 1
	} else if i == col-1 {
		startCol = i - 1
		endCol = i
	} else {
		startCol = i - 1
		endCol = i + 1
	}

	if j == 0 {
		startRow = j
		endRow = j + 1
	} else if j == row-1 {
		startRow = j - 1
		endRow = j
	} else {
		startRow = j - 1
		endRow = j + 1
	}

	for k := startCol; k <= endCol; k++ {
		for l := startRow; l <= endRow; l++ {
			if mtx[k][l] == SYMBOL || mtx[k][l] == ENGINE {
				return true
			}
		}
	}

	return false
}

func getAdjacencyEngine(mtx [][]int, row, col, i, j int) int {
	var startRow, startCol, endRow, endCol int

	if i == 0 {
		startCol = i
		endCol = i + 1
	} else if i == col-1 {
		startCol = i - 1
		endCol = i
	} else {
		startCol = i - 1
		endCol = i + 1
	}

	if j == 0 {
		startRow = j
		endRow = j + 1
	} else if j == row-1 {
		startRow = j - 1
		endRow = j
	} else {
		startRow = j - 1
		endRow = j + 1
	}

	cntEngine := 0
	sum := 1

	for k := startCol; k <= endCol; k++ {
		lastNum := 0
		for l := startRow; l <= endRow; l++ {
			if mtx[k][l] > 0 && lastNum != mtx[k][l] {
				lastNum = mtx[k][l]
				cntEngine++
				if cntEngine > 2 {
					return 0
				}
				sum *= mtx[k][l]
			}
		}
	}

	if cntEngine != 2 {
		return 0
	}

	return sum
}

func countDigits(n int) int {
	ret := 0

	for n > 0 {
		n = n / 10
		ret++
	}

	return ret
}
