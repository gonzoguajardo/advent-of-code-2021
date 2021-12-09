package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day9-part1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := make([][]int, 0)
	for scanner.Scan() {
		row := make([]int, 0)
		for _, char := range scanner.Text() {
			value, _ := strconv.Atoi(string(char))
			row = append(row, value)
		}
		matrix = append(matrix, row)
	}
	//printMatrix(matrix)

	sum := 0
	for r, row := range matrix {
		for c, number := range row {

			//check top
			if r != 0 {
				if number >= matrix[r-1][c] {
					continue
				}
			}

			//check bottom
			if r != len(matrix)-1 {
				if number >= matrix[r+1][c] {
					continue
				}
			}

			//check left
			if c != 0 {
				if number >= matrix[r][c-1] {
					continue
				}
			}

			//check right
			if c != len(matrix[0])-1 {
				if number >= matrix[r][c+1] {
					continue
				}
			}

			//fmt.Println(r, c)
			sum += number + 1

		}
	}
	fmt.Println(sum)
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}
