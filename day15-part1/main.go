package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"math"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	matrix := make([][]int, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		current := make([]int, 0)
		for _, char := range text {
			s := string(char)
			i, _ := strconv.Atoi(s)
			current = append(current, i)
		}
		matrix = append(matrix, current)
	}

	dp := make(map[[2]int]int)
	fmt.Println(r(0, 0, matrix, 0 - matrix[0][0], dp))
}

func r(x int, y int, matrix [][]int, risk int, dp map[[2]int]int) int {
	risk += matrix[y][x]
	if y == len(matrix)-1 && x == len(matrix[0])-1 {
		return risk
	}

	current := [2]int{x, y}
	if val, ok := dp[current]; ok && risk > val {
		// fmt.Println("short")
		return math.MaxInt16
	} else {
		dp[current] = risk
	}
	// down
	down := math.MaxInt16
	right := math.MaxInt16
	if y != len(matrix)-1 {
		down = r(x, y+1, matrix, risk, dp)
	}
	// right
	if x != len(matrix[0])-1 {
		right = r(x+1, y, matrix, risk, dp)
	}
	if down < right {
		return down
	}
	return right
}
