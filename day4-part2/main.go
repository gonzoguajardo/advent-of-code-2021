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
	file, err := os.Open("day4-part2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lineCount int
	var bingoNumbers []string
	var boards []bingoBoard
	var currentBoard bingoBoard

	// handle input
	for scanner.Scan() {
		if lineCount == 0 {
			bingoNumbers = strings.Split(scanner.Text(), ",")
		} else {
			currentLine := scanner.Text()
			if currentLine == "" {
				if lineCount != 1 {
					boards = append(boards, currentBoard)
				}
				currentBoard = *newBingoBoard()
			} else {
				currentBoard.loadLine(scanner.Text())
			}
		}
		lineCount++
	}
	boards = append(boards, currentBoard)

	// play bingo
	var numberOfWins int
main:
	for _, numberString := range bingoNumbers {
		number, _ := strconv.Atoi(numberString)
		for index, board := range boards {
			if !board.won {
				if board.applyAndCheck(number) {
					fmt.Printf("%v board won\n", index)
					numberOfWins++
					boards[index] = board
					if numberOfWins == len(boards) {
						var unmarkedSum int
						for k, v := range board.referenceMap {
							if v[0] != -1 {
								unmarkedSum += k
							}
						}
						fmt.Println(number * unmarkedSum)
						break main
					}

				}
			}

		}
	}
}

type bingoBoard struct {
	referenceMap map[int][]int
	board        [][]bool
	won          bool
}

func (bb *bingoBoard) loadLine(line string) {
	split := strings.Fields(line)
	for index, numberString := range split {
		number, _ := strconv.Atoi(numberString)
		bb.referenceMap[number] = []int{len(bb.board), index}
	}
	bb.board = append(bb.board, make([]bool, len(split)))
}

func (bb *bingoBoard) applyAndCheck(number int) bool {
	if val, ok := bb.referenceMap[number]; ok {
		bb.board[val[0]][val[1]] = true
		bb.referenceMap[number] = []int{-1, -1} //marking map we have a hit
		didWin := bb.checkRow(val) || bb.checkCol(val)
		if didWin {
			bb.won = true
		}
		return didWin
	}
	return false
}

func (bb *bingoBoard) checkRow(hit []int) bool {
	for i := 0; i < 5; i++ {
		if !bb.board[hit[0]][i] {
			return false
		}
	}
	return true
}

func (bb *bingoBoard) checkCol(hit []int) bool {
	for i := 0; i < 5; i++ {
		if !bb.board[i][hit[1]] {
			return false
		}
	}
	return true
}

func (bb bingoBoard) String() string {
	return fmt.Sprintln(bb.won)
}

func newBingoBoard() *bingoBoard {
	return &bingoBoard{referenceMap: map[int][]int{}}
}
