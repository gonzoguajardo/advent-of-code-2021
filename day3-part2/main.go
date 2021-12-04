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

	o2 := make([]string, len(input))
	copy(o2, input)
	var currentIndex int
	for len(o2) > 1 {
		var zeroCount int
		var oneCount int
		var zeroArray []string
		var oneArray []string

		for _, currentString := range o2 {
			if currentString[currentIndex] == 48 { //zero
				zeroCount++
				zeroArray = append(zeroArray, currentString)
			} else {
				oneCount++
				oneArray = append(oneArray, currentString)
			}
		}

		if oneCount >= zeroCount {
			o2 = make([]string, len(oneArray))
			copy(o2, oneArray)
		} else {
			o2 = make([]string, len(zeroArray))
			copy(o2, zeroArray)
		}

		currentIndex++
	}

	// yes, you should move this to a method but meh lazzy
	co2 := make([]string, len(input))
	copy(co2, input)
	currentIndex = 0
	for len(co2) > 1 {
		var zeroCount int
		var oneCount int
		var zeroArray []string
		var oneArray []string

		for _, currentString := range co2 {
			if currentString[currentIndex] == 48 { //zero
				zeroCount++
				zeroArray = append(zeroArray, currentString)
			} else {
				oneCount++
				oneArray = append(oneArray, currentString)
			}
		}

		if oneCount < zeroCount {
			co2 = make([]string, len(oneArray))
			copy(co2, oneArray)
		} else {
			co2 = make([]string, len(zeroArray))
			copy(co2, zeroArray)
		}

		currentIndex++
	}

	o2value, _ := strconv.ParseInt(o2[0], 2, 64)
	co2value, _ := strconv.ParseInt(co2[0], 2, 64)
	fmt.Println(o2value * co2value)

}
