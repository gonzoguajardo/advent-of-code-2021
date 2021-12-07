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
	file, err := os.Open("day7-part2/input.txt")
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
	//fmt.Println(crabs)

	dp := make([][2]int, len(crabs))
	//left side
	var leftSum int
	var leftPositions [][2]int //[pos, count]
	for i := 1; i < len(crabs); i++ {
		var currentLeftSum int
		if crabs[i-1] > 0 {
			leftPositions = append(leftPositions, [2]int{i - 1, crabs[i-1]})
		}
		for _, position := range leftPositions {
			currentLeftSum += position[1] * (i - position[0])
		}
		leftSum += currentLeftSum
		dp[i][0] = leftSum
		//fmt.Printf("i: %v\n", i)
		//fmt.Printf("currentLeftSum: %v\n", currentLeftSum)
		//fmt.Printf("leftSum: %v\n", leftSum)
	}
	//right side
	var rightSum int
	var rightPositions [][2]int //[pos, count]
	for i := len(crabs) - 2; i >= 0; i-- {
		var currentRightSum int
		if crabs[i+1] > 0 {
			rightPositions = append(rightPositions, [2]int{i + 1, crabs[i+1]})
		}
		for _, position := range rightPositions {
			currentRightSum += position[1] * (position[0] - i)
		}
		rightSum += currentRightSum
		dp[i][1] = rightSum
		//fmt.Printf("i: %v\n", i)
		//fmt.Printf("currentRightSum: %v\n", currentRightSum)
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
