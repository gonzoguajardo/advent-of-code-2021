package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day13-part2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	board := newBoard()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}

		point := pointFromString(text)
		board.applyPoint(point)
		//board.applyLine(point, point)
	}

	//board.print()

	for scanner.Scan() {
		text := scanner.Text()
		fold := strings.Split(strings.Split(text, "fold along ")[1], "=")
		fmt.Println(fold)

		if fold[0] == "y" {
			y, _ := strconv.Atoi(fold[1])
			board.foldY(y)

		} else {
			x, _ := strconv.Atoi(fold[1])
			board.foldX(x)
		}
	}

	board.printDotsHashes()

	sum := 0
	for y := 0; y < len(board.boardSlice); y++ {
		for x := 0; x < len(board.boardSlice[y]); x++ {
			if board.boardSlice[y][x] {
				sum++
			}
		}
	}

	fmt.Println(sum)

}

type board struct {
	boardSlice [][]bool
}

func newBoard() *board {
	return &board{boardSlice: make([][]bool, 0)}
}

func (board *board) foldX(x int) {
	for i := x + 1; i < len(board.boardSlice[0]); i++ {
		foldedX := x - (i - x)
		for y := 0; y < len(board.boardSlice); y++ {
			if board.boardSlice[y][i] {
				board.boardSlice[y][foldedX] = true
			}
		}
	}

	for y := 0; y < len(board.boardSlice); y++ {
		board.boardSlice[y] = board.boardSlice[y][:x]
	}

}

func (board *board) foldY(y int) {
	for i := y + 1; i < len(board.boardSlice); i++ {
		foldedY := y - (i - y)
		for x := 0; x < len(board.boardSlice[i]); x++ {
			if board.boardSlice[i][x] {
				board.boardSlice[foldedY][x] = true
			}
		}
	}

	board.boardSlice = board.boardSlice[:y]
}

func (board *board) applyPoint(point *point) {
	board.increaseCapacityPoint(point)
	board.boardSlice[point.y][point.x] = true
}

func (board *board) increaseCapacityPoint(point *point) {

	currentYLength := len(board.boardSlice)
	maxY := max(point.y+1, currentYLength)
	var currentXLength int
	if currentYLength > 0 {
		currentXLength = len(board.boardSlice[0])
	}
	maxX := max(point.x+1, currentXLength)

	if maxY > currentYLength {
		for i := currentYLength; i < maxY; i++ {
			board.boardSlice = append(board.boardSlice, make([]bool, maxX))
		}
	}
	if maxX > currentXLength {
		additionalLength := maxX - currentXLength

		for i := 0; i < currentYLength; i++ {
			row := board.boardSlice[i]
			row = append(row, make([]bool, additionalLength)...)
			board.boardSlice[i] = row
		}
	}
}

func (board board) printDotsHashes() {
	for _, row := range board.boardSlice {
		for _, b := range row {
			if b {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (board board) print() {
	for _, row := range board.boardSlice {
		fmt.Println(row)
	}
	fmt.Println()
}

type point struct {
	x int
	y int
}

func pointFromString(input string) *point {
	var pointSlice []int
	for _, pointString := range strings.Split(input, ",") {
		pointInt, _ := strconv.Atoi(pointString)
		pointSlice = append(pointSlice, pointInt)
	}
	return &point{x: pointSlice[0], y: pointSlice[1]}
}

func min(value1 int, value2 int) int {
	if value1 < value2 {
		return value1
	} else {
		return value2
	}
}
func max(value1 int, value2 int) int {
	if value1 > value2 {
		return value1
	} else {
		return value2
	}
}
