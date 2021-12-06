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
	file, err := os.Open("day5-part1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// handle input
	board := newBoard()
	for scanner.Scan() {
		points := strings.Split(scanner.Text(), " -> ")
		startPoint := pointFromString(points[0])
		endPoint := pointFromString(points[1])
		board.applyLine(startPoint, endPoint)

	}
	//board.print()

	// determine number of points with 2 or more
	var numberGreaterThanTwo int
	for _, row := range board.boardSlice {
		for _, value := range row {
			if value >= 2 {
				numberGreaterThanTwo++
			}
		}
	}
	fmt.Println(numberGreaterThanTwo)

}

type board struct {
	boardSlice [][]int
}

func newBoard() *board {
	return &board{boardSlice: make([][]int, 0)}
}
func (board *board) applyLine(startPoint *point, endPoint *point) {
	//capacity increase if needed
	board.increaseCapacity(startPoint, endPoint)

	var minValue, maxValue int
	if startPoint.x == endPoint.x {
		minValue = min(startPoint.y, endPoint.y)
		maxValue = max(startPoint.y+1, endPoint.y+1)
		for i := minValue; i < maxValue; i++ {
			board.boardSlice[i][startPoint.x]++
		}

	} else if startPoint.y == endPoint.y {
		minValue = min(startPoint.x, endPoint.x)
		maxValue = max(startPoint.x+1, endPoint.x+1)
		for i := minValue; i < maxValue; i++ {
			board.boardSlice[startPoint.y][i]++
		}
	}
}
func (board *board) increaseCapacity(startPoint *point, endPoint *point) {

	currentYLength := len(board.boardSlice)
	maxY := max(max(startPoint.y+1, endPoint.y+1), currentYLength)
	var currentXLength int
	if currentYLength > 0 {
		currentXLength = len(board.boardSlice[0])
	}
	maxX := max(max(startPoint.x+1, endPoint.x+1), currentXLength)

	if maxY > currentYLength {
		for i := currentYLength; i < maxY; i++ {
			board.boardSlice = append(board.boardSlice, make([]int, maxX))
		}
	}
	if maxX > currentXLength {
		additionalLength := maxX - currentXLength

		for i := 0; i < currentYLength; i++ {
			row := board.boardSlice[i]
			row = append(row, make([]int, additionalLength)...)
			board.boardSlice[i] = row
		}
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
