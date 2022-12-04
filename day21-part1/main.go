package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var player1Location, player2Location int
	var player1, player2 *player
	first := true
	for scanner.Scan() {
		currentString := scanner.Text()
		spaceSplit := strings.Split(currentString, " ")
		if first {
			player1Location , _ = strconv.Atoi(spaceSplit[4])
			player1 = &player{
				Location: player1Location,
			}
			first = false;
		} else {
			player2Location , _ = strconv.Atoi(spaceSplit[4])
			player2 = &player{
				Location: player2Location,
			}
		}
	}


	currentDice := 1
	var currentAdditionalScore int

	for {
		currentAdditionalScore = currentDice + (currentDice + 1) + (currentDice + 2)
		if player1.takeTurn(currentAdditionalScore) {
			log.Println(player2.Score * (currentDice + 2))
			break
		}
		currentDice += 3

		currentAdditionalScore = currentDice + (currentDice + 1) + (currentDice + 2)
		if player2.takeTurn(currentAdditionalScore) {
			log.Println(player1.Score * (currentDice + 2))
			break
		}
		currentDice += 3
	}

}

type player struct {
	Score int
	Location int
}

func (player *player) takeTurn(additionalScore int) (bool) {
	player.Location = (player.Location + additionalScore) % 10
	if player.Location == 0 {
		player.Location = 10
	}
	player.Score += player.Location
	return player.Score >= 1000
}
