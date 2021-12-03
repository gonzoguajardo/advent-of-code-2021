package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1-part2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	ifFirstTwo := 0
	var count int = 0
	var previousSum int
	var currentSum int
	var previous int
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		if ifFirstTwo < 3 {
			current, _ := strconv.Atoi(scanner.Text())
			previousSum += current
			if ifFirstTwo != 0 {
				currentSum += current
			}
			if ifFirstTwo == 2 {
				previous = current
			}
			ifFirstTwo++
			continue
		}else {
			current, _ := strconv.Atoi(scanner.Text())

			currentSum += current

			fmt.Printf("previousSum %v", previousSum)
			fmt.Println()
			fmt.Printf("currentSum: %v", currentSum)
			fmt.Println()

			if currentSum > previousSum {
				count++
			}

			previousSum = currentSum
			currentSum = currentSum - (currentSum - current - previous)
			previous = current


		}

		fmt.Printf("count: %v", count)
		fmt.Println()
	}

	print(count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
