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
	file, err := os.Open("day6-part1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// handle input
	pond := make([]int64, 9)
	for scanner.Scan() {
		for _, indexString := range strings.Split(scanner.Text(), ",") {
			index, _ := strconv.Atoi(indexString)
			pond[index]++
		}
	}
	for days := 1; days < 81; days++ {
		pondCopy := make([]int64, len(pond))
		copy(pondCopy, pond)
		pondCopy[6] = pond[0] + pond[7]
		pondCopy[8] = pond[0]
		for everyOther := 1; everyOther < 9; everyOther++ {
			if everyOther == 7 {
				continue
			}
			pondCopy[everyOther-1] = pond[everyOther]
		}
		pond = pondCopy
	}
	fmt.Println(pond)
	var sum int64
	for _, day := range pond {
		sum += day
	}
	fmt.Println(sum)
}
