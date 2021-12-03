package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day3-part1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var counts [][2]int
	first := true
	for scanner.Scan() {
		currentLine := scanner.Text()

		if first {
			counts = make([][2]int, len(currentLine))
			first = false
		}

		for index, char := range currentLine {
			if first {
				counts[index] = [2]int{0, 0}
				first = false
			}
			if char == 48 {
				counts[index][0]++
			} else {
				counts[index][1]++
			}
		}
		//break
	}

	//fmt.Println(counts)

	var gammaString string
	var epsilonString string
	for _, currentDigit := range counts {
		if currentDigit[0] > currentDigit[1] {
			gammaString += "0"
			epsilonString += "1"
		} else {
			gammaString += "1"
			epsilonString += "0"
		}
	}

	gamma, _ := strconv.ParseInt(gammaString, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonString, 2, 64)

	fmt.Println(gamma * epsilon)

}
