package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("day8-part1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// handle input
	var count int

	var wordMap map[int][]string
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), "|")
		var text string

		wordMap = make(map[int][]string)
		text = lines[0]
		for _, word := range strings.Fields(text) {
			if word == "|" {
				continue
			}
			wordMap[len(word)] = append(wordMap[len(word)], word)
		}

		text = lines[1]
		for _, word := range strings.Fields(text) {
			length := len(word)
			if length == 2 || length == 3 || length == 4 || length == 7 {
				if len(wordMap[length]) == 1 {
					clean := true
					for _, char := range wordMap[length][0] {
						if !strings.ContainsRune(word, char) {
							fmt.Printf("breaking: %v %v\n", wordMap[length][0], word)
							clean = false
							break
						}
					}
					if clean {
						count++
					}
				}
			}
		}
	}
	fmt.Println(count)

}
