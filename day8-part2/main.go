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

	sourceOfTruth := make([][][2]string, 7)
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), "|")
		var uniqueSignals []string
		for _, uniqueSignal := range strings.Fields(lines[0]) {
			switch len(uniqueSignal) {
			case 2:
				truth := make([][2]string, 0)
				truth = append(truth, [2]string{uniqueSignal, originalCode[1]})
				sourceOfTruth[1] = truth
			case 3:
				truth := make([][2]string, 0)
				truth = append(truth, [2]string{uniqueSignal, originalCode[7]})
				sourceOfTruth[2] = truth
			case 4:
				truth := make([][2]string, 0)
				truth = append(truth, [2]string{uniqueSignal, originalCode[4]})
				sourceOfTruth[3] = truth
			case 7:
				truth := make([][2]string, 0)
				truth = append(truth, [2]string{uniqueSignal, originalCode[8]})
				sourceOfTruth[6] = truth
			default:
				uniqueSignals = append(uniqueSignals, uniqueSignal)
			}
		}

		workingOutput := make([]string, 0)
		outputs := make([]int, 4)
		var numberOfOutputs int

		for index, output := range strings.Fields(lines[1]) {
			switch len(output) {
			case 2:
				numberOfOutputs++
				outputs[index] = 1
			case 3:
				numberOfOutputs++
				outputs[index] = 7
			case 4:
				numberOfOutputs++
				outputs[index] = 4
			case 7:
				numberOfOutputs++
				outputs[index] = 8
			default:
				workingOutput = append(workingOutput, output)
			}
		}

		//ez solve for a
		for _, char := range sourceOfTruth[2][0][0] {
			if !strings.ContainsRune(sourceOfTruth[1][0][0], char) {
				truth := make([][2]string, 0)
				truth = append(truth, [2]string{string(char), "a"})
				sourceOfTruth[0] = truth
			}
		}

		//ez solve for 2nd pair
		sourceOfTruth[1] = append(sourceOfTruth[1], [2]string{removeSubstring(sourceOfTruth[3][0][0], sourceOfTruth[2][0][0]),
			removeSubstring(sourceOfTruth[3][0][1], sourceOfTruth[2][0][1])})

		//loop through and find other solutions
	signal:
		for len(uniqueSignals) > 0 {
			for _, uniqueSignal := range uniqueSignals {
				output := ""
				currentLengthCheck := 1
				for len(uniqueSignal) > 0 {
					if currentLengthCheck > 6 {
						break
					}
					//fmt.Printf("unique signal: %v\n", uniqueSignal)
					//fmt.Printf("output: %v\n", output)
					//fmt.Printf("currentLengthCheck: %v\n", currentLengthCheck)
					for _, truth := range sourceOfTruth[currentLengthCheck-1] {
						if containsSubstring(uniqueSignal, truth[0]) {
							uniqueSignal = removeSubstring(uniqueSignal, truth[0])
							output += truth[1]
							fmt.Println(uniqueSignal)
							fmt.Println(output)
						}
					}
					currentLengthCheck++

					if len(uniqueSignal) == 1 {
					originalCode:
						for _, oc := range originalCode {
							var count int
							if containsSubstring(oc, output) && len(output)+1 == len(oc) {
								count++
							}
							if count == 1 {
								for _, check := range sourceOfTruth[0] {
									if check[0] == uniqueSignal {
										continue originalCode
									}
								}

								elems := [2]string{uniqueSignal, removeSubstring(oc, output)}
								fmt.Println(elems)
								sourceOfTruth[0] = append(sourceOfTruth[0], elems)
								fmt.Printf("source of truth: %v\n", sourceOfTruth)

								if len(sourceOfTruth[0]) == 7 {
									break signal
								}
							}
						}

					}
				}

			}
		}

		for indexOutput, output := range workingOutput {
			fmt.Println(output)
			converted := convert(output, sourceOfTruth[0])
			fmt.Println(converted)
			for index, check := range originalCode {
				if check == converted {
					outputs[indexOutput] = index
				}
			}
		}

		fmt.Printf("source of truth: %v\n", sourceOfTruth)
		fmt.Printf("uniqueSignals: %v\n", uniqueSignals)
		fmt.Printf("working output: %v\n", workingOutput)
		fmt.Printf("outputs: %v\n", outputs)

	}

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

func convert(str string, keys [][2]string) string {
	var output string
	for _, char := range str {
		charString := string(char)
		for _, key := range keys {
			if charString == key[0] {
				output += key[1]
			}
		}
	}
	return output
}
