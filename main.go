package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gunjan01/battleship/battleship"
)

const (
	// Player represents a battleship player.
	player int = 0
	// Opponent represents a battleship Opponent.
	opponent int = 1
)

func generateOutputFile(game battleship.Game) {
	outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatal("Error while opening file: ", err)
	}
	defer outputFile.Close()

	// Write the player boards to the output file
	scoreBoard := make(map[int]string)
	for index := range game.Players {
		player := game.Players[index]
		_, err := outputFile.WriteString("Player" + strconv.Itoa(index+1) + "\n")
		if err != nil {
			log.Fatal("Error while writing to a file: ", err)
		}

		scoreBoard[index] = "P" + strconv.Itoa(index+1) + ":" + strconv.Itoa(player.TotalPoints)

		for i := range player.Board.Board {
			var str string
			for j := range player.Board.Board[i] {
				if player.Board.Board[i][j] != "" {
					str = str + player.Board.Board[i][j] + " "
				} else {
					str = str + "_" + " "
				}
			}

			// Write the string to file
			_, err = outputFile.WriteString(str + "\n")
			if err != nil {
				log.Fatal("Error while writing to a file: ", err)
			}
		}

		_, err = outputFile.WriteString("\n")
		if err != nil {
			log.Fatal("Error while writing to a file: ", err)
		}
	}

	// Write point to the output file.
	for _, value := range scoreBoard {
		_, err = outputFile.WriteString(value + "\n")
		if err != nil {
			log.Fatal("Error while writing to a file: ", err)
		}
	}

	// Write the final result to the output file.
	var finalResult string = "It is a draw"
	if game.Players[player].TotalPoints > game.Players[opponent].TotalPoints {
		finalResult = "Player 1 wins"
	}

	if game.Players[opponent].TotalPoints > game.Players[player].TotalPoints {
		finalResult = "Player 2 wins"
	}

	_, err = outputFile.WriteString("\n" + finalResult + "\n")
	if err != nil {
		log.Fatal("Error while writing to a file: ", err)
	}
}

func generateRandomCoordinates(n int) []battleship.Coordinates {
	rand.Seed(time.Now().UnixNano())

	pairs := make([]battleship.Coordinates, n)
	for i := 0; i < n; i++ {
		pairs[i] = battleship.Coordinates{
			X: rand.Intn(5), // Adjust the range based on your needs
			Y: rand.Intn(5),
		}
	}
	return pairs
}

func main() {
	// Start the game
	var numberOfPlayers int = 2
	game := battleship.NewGame(numberOfPlayers)

	// This is to limit the gridsize to 10.
	gridSize := 4 + rand.Intn(6)

	totalShips := rand.Intn(gridSize)
	p1ShipPositions := generateRandomCoordinates(totalShips)
	p2ShipPositions := generateRandomCoordinates(totalShips)

	totalMissiles := rand.Intn(7)
	p1MissileMoves := generateRandomCoordinates(totalMissiles)
	p2MissileMoves := generateRandomCoordinates(totalMissiles)

	// Setup Player and opponent
	game.SetUpPlayer(player, gridSize, totalShips, totalMissiles, p1ShipPositions, p1MissileMoves)
	game.SetUpPlayer(opponent, gridSize, totalShips, totalMissiles, p2ShipPositions, p2MissileMoves)

	var wg sync.WaitGroup
	wg.Add(len(game.Players))

	// Two independent threads to fire missiles simulataneously.
	go func() {
		defer wg.Done()
		game.FireMissiles(player, opponent)
	}()

	go func() {
		defer wg.Done()
		game.FireMissiles(opponent, player)
	}()

	// Wait for the go routines to finish.
	wg.Wait()

	// generate the final output.
	generateOutputFile(game)
}
