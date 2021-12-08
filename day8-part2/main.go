package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("day8-part2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// handle input
	var count int

	originalCode := []string{
		"abcefg",
		"cf",
		"acdeg",
		"acdfg",
		"bcdf",
		"abdfg",
		"abdefg",
		"acf",
		"abcdefg",
		"abcdfg",
	}

	var wordMap map[int][]string
	sourceOfTruth := make([][][2]string, 7)
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
						if length == 7 {
							answerMap := make([][2]string, 0)
							answerMap = append(answerMap, [2]string{originalCode[8], word})
							sourceOfTruth[7-1] = answerMap
						} else if length == 4 {
							answerMap := make([][2]string, 0)
							answerMap = append(answerMap, [2]string{originalCode[4], word})
							sourceOfTruth[4-1] = answerMap
						}

						count++
					}
				}
			}
		}

		fmt.Println(sourceOfTruth)
		text = lines[0]
		for _, word := range strings.Fields(text) {
			fmt.Println(word)
		}

		break
	}

}
