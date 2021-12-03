package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1-part1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	ifFirst := true
	var count int = 0
	var second int
	var first int
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		if ifFirst {
			second, _ = strconv.Atoi(scanner.Text())
			ifFirst = false
			continue
		}else {
			first = second
			second, _ = strconv.Atoi(scanner.Text())
			fmt.Printf("first %v", first)
			fmt.Println()
			fmt.Printf("second: %v", second)
			fmt.Println()
		}
		if second > first {
			count++
		}
		fmt.Printf("count: %v", count)
		fmt.Println()
	}

	print(count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
