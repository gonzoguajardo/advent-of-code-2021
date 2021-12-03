package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day2-part1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var horizontal int
	var depth int
	var aim int

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		currentLine := scanner.Text()
		split := strings.Split(currentLine, " ")
		value, _ := strconv.Atoi(split[1])

		if split[0] == "forward" {
			horizontal += value
			depth = depth + (value * aim)
		} else if split[0] == "down" {
			aim += value
			//depth += value
		} else if split[0] == "up" {
			aim -= value
			//depth -= value
		}

		//break
	}

	position := horizontal * depth
	fmt.Println(position)


}
