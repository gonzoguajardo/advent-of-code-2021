package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day3-part2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// I guess I need this in memory to loop more than once? is there a O(n) solution?
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	//fmt.Println(input)

	// and then we need to loop through twice since the filtering is different? meh, there is probably a better way...
	o2Value := reduceToValue(input, true)
	co2Value := reduceToValue(input, false)
	fmt.Println(o2Value * co2Value)

}

func reduceToValue(input []string, favorOne bool) int64 {
	filteredSlice := make([]string, len(input))
	copy(filteredSlice, input)
	var currentIndex int
	for len(filteredSlice) > 1 {
		var zeroCount int
		var oneCount int
		var zeroArray []string
		var oneArray []string

		for _, currentString := range filteredSlice {
			if currentString[currentIndex] == 48 { //zero
				zeroCount++
				zeroArray = append(zeroArray, currentString)
			} else {
				oneCount++
				oneArray = append(oneArray, currentString)
			}
		}

		if favorOne {
			if oneCount >= zeroCount {
				filteredSlice = make([]string, len(oneArray))
				copy(filteredSlice, oneArray)
			} else {
				filteredSlice = make([]string, len(zeroArray))
				copy(filteredSlice, zeroArray)
			}
		} else {
			if oneCount < zeroCount {
				filteredSlice = make([]string, len(oneArray))
				copy(filteredSlice, oneArray)
			} else {
				filteredSlice = make([]string, len(zeroArray))
				copy(filteredSlice, zeroArray)
			}
		}

		currentIndex++
	}
	value, _ := strconv.ParseInt(filteredSlice[0], 2, 64)
	return value
}
