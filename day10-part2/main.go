package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("day10-part2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	firstScore := 0
	secondScores := make([]int, 0)
scanner:
	for scanner.Scan() {
		fifoQ := fifoQ{make([]string, 0)}
		text := scanner.Text()
		//fmt.Println(text)
		for _, char := range text {
			s := string(char)
			//fmt.Println(s)
			switch s {
			case ")", "]", "}", ">":
				ifPopped, currentScore := fifoQ.pop(s)
				if !ifPopped {
					//fmt.Println("hit", s, currentScore)
					firstScore += currentScore
					continue scanner
				}
			default:
				fifoQ.append(s)
			}
		}

		secondScore := 0
		//fmt.Println("queue: ", fifoQ.queue)
		for len(fifoQ.queue) > 0 {
			secondScore = secondScore*5 + fifoQ.incompletePop()
			//fmt.Println("queue: ", fifoQ.queue)
			//fmt.Println("second score:", secondScore)

		}
		secondScores = append(secondScores, secondScore)

		//fmt.Println("second score:", secondScore)

	}
	fmt.Println("first score", firstScore)
	sort.Ints(secondScores)
	fmt.Println("second scores: ", secondScores)
	fmt.Println("middle second score: ", secondScores[len(secondScores)/2])
}

type fifoQ struct {
	queue []string
}

func (fifoQ *fifoQ) append(char string) {
	fifoQ.queue = append(fifoQ.queue, char)
}
func (fifoQ *fifoQ) pop(end string) (bool, int) {
	//fmt.Println(fifoQ.queue)
	switch end {
	case ")":
		if fifoQ.queue[len(fifoQ.queue)-1] == "(" {
			fifoQ.removeLast()
			return true, 0
		} else {
			//fmt.Println("invalid bracket", fifoQ.queue)
			return false, 3
		}
	case "]":
		if fifoQ.queue[len(fifoQ.queue)-1] == "[" {
			fifoQ.removeLast()
			return true, 0
		} else {
			return false, 57
		}
	case "}":
		if fifoQ.queue[len(fifoQ.queue)-1] == "{" {
			fifoQ.removeLast()
			return true, 0
		} else {
			return false, 1197
		}
	case ">":
		if fifoQ.queue[len(fifoQ.queue)-1] == "<" {
			fifoQ.removeLast()
			return true, 0
		} else {
			//fmt.Println("invalid bracket", fifoQ.queue)
			return false, 25137
		}
	default:
		fmt.Println("no match", end)
	}
	return false, 0
}

func (fifoQ *fifoQ) removeLast() {
	//fmt.Println("before remove", fifoQ.queue)
	fifoQ.queue = fifoQ.queue[:len(fifoQ.queue)-1]
	//fmt.Println("after remove", fifoQ.queue)
}

func (fifoQ *fifoQ) incompletePop() int {
	end := fifoQ.queue[len(fifoQ.queue)-1]
	switch end {
	case "(":
		fifoQ.removeLast()
		return 1
	case "[":
		fifoQ.removeLast()
		return 2
	case "{":
		fifoQ.removeLast()
		return 3
	case "<":
		fifoQ.removeLast()
		return 4
	default:
		fmt.Println("no match", end)
	}
	return 0
}
