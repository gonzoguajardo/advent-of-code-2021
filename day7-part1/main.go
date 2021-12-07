package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day7-part1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// handle input
	crabs := make([]int, 0)
	for scanner.Scan() {
		for _, indexString := range strings.Split(scanner.Text(), ",") {
			index, _ := strconv.Atoi(indexString)
			if index >= len(crabs) {
				crabs = append(crabs, make([]int, index-len(crabs)+1)...)
			}
			crabs[index]++
		}
	}

	dp := make([][2]int, len(crabs))
	//left side
	var leftCount int
	var leftSum int
	for i := 1; i < len(crabs); i++ {
		leftCount += crabs[i-1]
		leftSum += leftCount
		dp[i][0] = leftSum
		//fmt.Printf("i: %v\n", i)
		//fmt.Printf("rightCount: %v\n", leftCount)
		//fmt.Printf("rightSum: %v\n", leftSum)
	}
	//right side
	var rightCount int
	var rightSum int
	for i := len(crabs) - 2; i >= 0; i-- {
		rightCount += crabs[i+1]
		rightSum += rightCount
		dp[i][1] = rightSum
		//fmt.Printf("i: %v\n", i)
		//fmt.Printf("rightCount: %v\n", rightCount)
		//fmt.Printf("rightSum: %v\n", rightSum)
	}
	minSum := math.MaxInt
	for _, sums := range dp {
		currentSum := sums[0] + sums[1]
		if currentSum < minSum {
			minSum = currentSum
		}
	}
	fmt.Println(minSum)
}
