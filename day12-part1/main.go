package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("day12-part1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pathMap := make(map[string][]string)
	for scanner.Scan() {
		text := scanner.Text()
		pathsInput := strings.Split(text, "-")
		addCaveSystem(pathsInput[0], pathsInput[1], pathMap)

	}
	fmt.Println("path map: ", pathMap)
	//walk through paths
	foundPaths := make([]string, 0)
	foundPaths = walkThroughPaths([]string{"start"}, pathMap, make(map[string]bool), foundPaths)
	//fmt.Println("found paths: ", foundPaths)
	fmt.Println("number of paths: ", len(foundPaths))
}

func walkThroughPaths(currentPath []string, pathMap map[string][]string, visitedSmallCaves map[string]bool, foundPaths []string) []string {
	lastPath := currentPath[len(currentPath)-1]
	if lastPath == "end" {
		visitedSmallCaves[lastPath] = true
		return []string{strings.Join(currentPath[:], ",")}
	}
	if unicode.IsLower(rune(lastPath[0])) {
		visitedSmallCaves[lastPath] = true
	}

	if val, ok := pathMap[lastPath]; ok {
		for _, nextPath := range val {
			if !visitedSmallCaves[nextPath] {
				visitedSmallCavesCopy := make(map[string]bool)
				for key, value := range visitedSmallCaves {
					visitedSmallCavesCopy[key] = value
				}

				paths := walkThroughPaths(append(currentPath, nextPath), pathMap, visitedSmallCavesCopy, foundPaths)
				for _, path := range paths {
					exists := false
					for _, foundPath := range foundPaths {
						if foundPath == path {
							exists = true
							break
						}
					}
					if !exists {
						foundPaths = append(foundPaths, path)
					}
				}
			}
			//fmt.Println("iterating through: ", nextPath)
		}
	}

	return foundPaths
}

func addCaveSystem(cave1 string, cave2 string, pathMap map[string][]string) {
	startCave := cave1
	endCave := cave2
	if val, ok := pathMap[startCave]; ok {
		val = append(val, endCave)
		pathMap[startCave] = val
	} else {
		pathMap[startCave] = []string{endCave}
	}
	startCave = cave2
	endCave = cave1
	if val, ok := pathMap[startCave]; ok {
		val = append(val, endCave)
		pathMap[startCave] = val
	} else {
		pathMap[startCave] = []string{endCave}
	}

}
