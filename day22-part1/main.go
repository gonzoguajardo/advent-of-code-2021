package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	cubeCount := 0
	cubeMap := make(map[[3]int]bool)
	for scanner.Scan() {
		currentLine := scanner.Text()

		// handle status
		on := false
		spaceSplitSlice := strings.Split(currentLine, " ")
		if spaceSplitSlice[0] == "on" {
			on = true
		}

		// handle cubes
		commaSplitSlice := strings.Split(spaceSplitSlice[1], ",")
		xStart, xEnd := handleCommaSplit(commaSplitSlice[0])
		yStart, yEnd := handleCommaSplit(commaSplitSlice[1])
		zStart, zEnd := handleCommaSplit(commaSplitSlice[2])

		for xIndex := max(xStart, -50); xIndex <= min(xEnd, 50); xIndex ++ {
			for yIndex := max(yStart, -50); yIndex <= min(yEnd, 50); yIndex ++ {
				for zIndex := max(zStart, -50); zIndex <= min(zEnd, 50); zIndex ++ {
					if on {
						val, ok := cubeMap[[3]int{xIndex, yIndex, zIndex}]
						if ok {
							if !val {
								cubeCount++
							}
						} else {
							cubeCount++
						}
						cubeMap[[3]int{xIndex, yIndex, zIndex}] = true
					}
					if !on {
						val, ok := cubeMap[[3]int{xIndex, yIndex, zIndex}]
						if ok && val {
							cubeCount--
							cubeMap[[3]int{xIndex, yIndex, zIndex}] = false
						}
					}
				}
			}
		}
	}
	log.Println(cubeCount)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func handleCommaSplit(commaSplit string) (int, int) {
	dotsSplit := strings.Split(commaSplit, "..")
	startString := dotsSplit[0][2:]
	start , _ := strconv.Atoi(startString)
	endString := dotsSplit[1]
	end , _ := strconv.Atoi(endString)
	return start, end
}
