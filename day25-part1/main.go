package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	area := make([][]string, 0)
	for scanner.Scan() {
		currentString := scanner.Text()
		currentLine := make([]string, 0)
		for _, char := range(currentString) {
			currentCharString := string(char)
			currentLine = append(currentLine, currentCharString)
		}
		area = append(area, currentLine)
	}
	n := 0
	area, n = move(area)
	log.Println(n + 1)
}

func move(area [][]string) ([][]string, int) {
	dest := make([][]string, len(area))
	for i := range area {
		dest[i] = make([]string, len(area[i]))
		copy(dest[i], area[i])
	}
	n := 0
	for {
		var movedEast bool
		var movedSouth bool
		dest, movedEast = moveEast(dest)
		dest, movedSouth = moveSouth(dest)
		if !movedEast && !movedSouth {
			break
		}
		n++
	}
	return dest, n

}

func moveSouth(area[][]string) ([][]string, bool){
	dest := make([][]string, len(area))
	for i := range area {
		dest[i] = make([]string, len(area[i]))
		copy(dest[i], area[i])
	}
	moved := false
	for col := 0; col < len(area[0]); col++ {
		row := 0
		for row < len(area) {
			char := area[row][col]
			if char == "v" {
				if row + 1 < len(area) && area[row + 1][col] == "." {
					dest[row][col] = "."
					dest[row + 1][col] = "v"
					row += 2
					moved = true
					continue
				} else if row + 1 == len(area) && area[0][col] == "." {
					dest[row][col] = "."
					dest[0][col] = "v"
					row += 2
					moved = true
					continue
				}
			}
			row++
		}
	}
	return dest, moved
}

func moveEast(area [][]string) ( [][]string, bool ) {
	dest := make([][]string, len(area))
	for i := range area {
		dest[i] = make([]string, len(area[i]))
		copy(dest[i], area[i])
	}
	moved := false
	for row, line := range(area) {
		col := 0
		for col < len(line){
			char := line[col]
			if char == ">" {
				if col + 1 < len(line) && area[row][col + 1] == "." {
					dest[row][col] = "."
					dest[row][col + 1] = ">"
					col += 2
					moved = true
				} else if col + 1 == len(line) && area[row][0] == "." {
					dest[row][col] = "."
					dest[row][0] = ">"
					col += 2
					moved = true
				}else{
					col++
				}
			}else {
				col++
			}
		}
	}
	return dest, moved
}

func print(area [][]string) {
	for _, line := range(area) {
		log.Println(line)
	}
}
