package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day10-part1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	score := 0

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
				ifPoped, currentScore := fifoQ.pop(s)
				if !ifPoped {
					//fmt.Println("hit", s, currentScore)
					score += currentScore
					continue scanner
				}
			default:
				fifoQ.append(s)
			}
		}

	}
	fmt.Println(score)

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
