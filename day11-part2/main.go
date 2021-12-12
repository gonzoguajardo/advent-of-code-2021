package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day11-part1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ocean := newOcean()
	for scanner.Scan() {
		text := scanner.Text()
		row := make([]int, 0)
		for _, char := range text {
			n, _ := strconv.Atoi(string(char))
			row = append(row, n)
		}
		ocean.matrix = append(ocean.matrix, row)
		//fmt.Println(text)
	}

	fmt.Println("Before any steps: ")
	ocean.print()
	count := 0
	maxStep := 100
	for step := 1; step <= maxStep; step++ {
		count += ocean.step()
		fmt.Println("After step ", step)
		//ocean.print()
	}
	println("count: ", count)

}

type ocean struct {
	matrix [][]int
}

func newOcean() *ocean {
	return &ocean{matrix: make([][]int, 0)}
}
func (ocean *ocean) print() {
	for _, row := range ocean.matrix {
		fmt.Println(row)
	}
	fmt.Println()
}
func (ocean *ocean) step() int {
	needsToFlash := make([][2]int, 0)
	for r, row := range ocean.matrix {
		for c, n := range row {
			if n == 9 {
				//fmt.Println(r, c)
				needsToFlash = append(needsToFlash, [2]int{r, c})
				//ocean.matrix[r][c] = 0
			} else {
				ocean.matrix[r][c]++
			}
		}
	}
	//fmt.Println("needs to flash: ", needsToFlash)
	hasFlashed := make([][]bool, len(ocean.matrix))
	for index, _ := range hasFlashed {
		hasFlashed[index] = make([]bool, len(ocean.matrix[index]))
	}
	for _, preFlash := range needsToFlash {
		hasFlashed[preFlash[0]][preFlash[1]] = true
	}
	for _, toFlash := range needsToFlash {
		//fmt.Println("original flash loop", toFlash)
		flash(toFlash[0], toFlash[1], ocean.matrix, hasFlashed, true)
		//break
	}
	count := 0
	for _, row := range hasFlashed {
		for _, hasFlash := range row {
			if hasFlash {
				count++
			}
		}
	}
	return count
	//ocean.print()

}

func flash(r int, c int, matrix [][]int, hasFlashed [][]bool, origFlasher bool) {
	//fmt.Println("stepping through: ", r, c)
	if r < 0 || r == len(matrix) || c < 0 || c == len(matrix) {
		return
	}
	if !origFlasher && hasFlashed[r][c] {
		return
	} else {
		//flash
		//fmt.Println("before flash")
		if matrix[r][c] == 9 {
			hasFlashed[r][c] = true
			//fmt.Println("flashing: ", r, c)
			//top left
			flash(r-1, c-1, matrix, hasFlashed, false)
			//top
			flash(r-1, c, matrix, hasFlashed, false)
			//top right
			flash(r-1, c+1, matrix, hasFlashed, false)
			//left
			flash(r, c-1, matrix, hasFlashed, false)
			//right
			flash(r, c+1, matrix, hasFlashed, false)
			//bottom left
			flash(r+1, c-1, matrix, hasFlashed, false)
			//bottom
			flash(r+1, c, matrix, hasFlashed, false)
			//bottom right
			flash(r+1, c+1, matrix, hasFlashed, false)
			matrix[r][c] = 0
		} else {
			matrix[r][c]++
		}
		//fmt.Println("after flash")
	}
}
