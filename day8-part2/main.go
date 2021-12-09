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
	file, err := os.Open("day8-part2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// handle input

	answerMap := make(map[int]string)
	letterMap := make(map[string]string)
	var sum int

	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), "|")
		var uniqueSignals []string
		for _, uniqueSignal := range strings.Fields(lines[0]) {
			switch len(uniqueSignal) {
			case 2:
				answerMap[1] = uniqueSignal
			case 3:
				answerMap[7] = uniqueSignal
			case 4:
				answerMap[4] = uniqueSignal
			case 7:
				answerMap[8] = uniqueSignal
			default:
				uniqueSignals = append(uniqueSignals, uniqueSignal)
			}
		}

		workingOutput := make([]string, 0)

		for _, output := range strings.Fields(lines[1]) {
			workingOutput = append(workingOutput, output)

		}

		//fmt.Printf("uniqueSignals: %v\n", uniqueSignals)
		//fmt.Printf("working output: %v\n", workingOutput)
		//fmt.Printf("answer: %v\n", answerMap)
		//fmt.Printf("outputs: %v\n", outputs)

		letterMap["a"] = removeSubstring(answerMap[7], answerMap[1])

		//find 9
		fourPlusSeven := addUniques(answerMap[4], answerMap[7])
		uniqueSignalsCopy := make([]string, 0)
		for _, us := range uniqueSignals {
			if len(fourPlusSeven)+1 == len(us) && containsSubstring(us, fourPlusSeven) {
				answerMap[9] = us
				letterMap["g"] = removeSubstring(us, fourPlusSeven)
			} else {
				uniqueSignalsCopy = append(uniqueSignalsCopy, us)
			}
		}
		uniqueSignals = uniqueSignalsCopy
		//fmt.Printf("uniqueSignals: %v\n", uniqueSignals)

		letterMap["e"] = removeSubstring(answerMap[8], answerMap[9])

		//find zero
		almostZero := answerMap[1] + letterMap["a"] + letterMap["e"] + letterMap["g"]
		uniqueSignalsCopy = make([]string, 0)
		for _, us := range uniqueSignals {
			if len(almostZero)+1 == len(us) && containsSubstring(us, almostZero) {
				answerMap[0] = us
				letterMap["b"] = removeSubstring(us, almostZero)
			} else {
				uniqueSignalsCopy = append(uniqueSignalsCopy, us)
			}
		}
		uniqueSignals = uniqueSignalsCopy
		//fmt.Printf("uniqueSignals: %v\n", uniqueSignals)

		letterMap["d"] = removeSubstring(answerMap[8], answerMap[0])

		//find three
		three := answerMap[1] + letterMap["a"] + letterMap["d"] + letterMap["g"]
		uniqueSignalsCopy = make([]string, 0)
		for _, us := range uniqueSignals {
			if len(three) == len(us) && containsSubstring(three, us) {
				answerMap[3] = us
			} else {
				uniqueSignalsCopy = append(uniqueSignalsCopy, us)
			}
		}
		uniqueSignals = uniqueSignalsCopy

		//find six
		uniqueSignalsCopy = make([]string, 0)
		for _, us := range uniqueSignals {
			if len(us) == 6 {
				answerMap[6] = us
			} else {
				uniqueSignalsCopy = append(uniqueSignalsCopy, us)
			}
		}
		uniqueSignals = uniqueSignalsCopy

		//find c
		letterMap["c"] = removeSubstring(answerMap[8], answerMap[6])

		//2 and 5
		uniqueSignalsCopy = make([]string, 0)
		for _, us := range uniqueSignals {
			if containsSubstring(us, letterMap["c"]) {
				answerMap[2] = us
			} else {
				answerMap[5] = us
			}
		}
		uniqueSignals = uniqueSignalsCopy

		//fmt.Printf("answer: %v\n", answerMap)
		//fmt.Printf("letterMap: %v\n", letterMap)
		//fmt.Printf("uniqueSignals: %v\n", uniqueSignals)

		outputValueString := ""
		for _, output := range workingOutput {
			for k, v := range answerMap {
				if len(v) == len(output) && containsSubstring(output, v) {
					outputValueString += strconv.Itoa(k)
				}
			}
		}

		outputValue, _ := strconv.Atoi(outputValueString)
		sum += outputValue
		fmt.Println(outputValueString)

		//break
	}

	fmt.Println(sum)

}
func containsSubstring(str string, substr string) bool {
	for _, char := range substr {
		if !strings.ContainsRune(str, char) {
			return false
		}
	}
	return true
}
func removeSubstring(str string, substr string) string {
	for _, char := range substr {
		str = strings.ReplaceAll(str, string(char), "")
	}
	return str
}
func addUniques(string1 string, string2 string) string {
	var output string
	for _, char := range string1 {
		if !strings.ContainsRune(output, char) {
			output += string(char)
		}
	}
	for _, char := range string2 {
		if !strings.ContainsRune(output, char) {
			output += string(char)
		}
	}
	return output
}
