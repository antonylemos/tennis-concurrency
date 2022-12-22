package main

// Import packages
import (
	"fmt"
	"sync"
	"tennis-concurrency/utils"
)

// Create a wait group
var waitGroup sync.WaitGroup

// Player function
func player(id int, match chan bool, pointsToWin int, playerPoints int) {
	// Defer the waitGroup.Done() function
	defer waitGroup.Done()

	// Infinite loop
	for {
		// Wait for the channel to be available
		finish := <-match

		// If the match is finished, return
		if !finish {
			return
		}

		// Start the player's turn
		fmt.Println("Jogador ", id, " começou a jogada!")

		// Make a point
		playerScored := utils.MakePoint()

		// If the player scored, increment the player's points
		if playerScored {
			playerPoints++

			fmt.Println("Jogador ", id, " marcou ", playerPoints, " pontos!")

			// If the player has reached the number of points to win, print the winner and return
			if playerPoints == pointsToWin {
				fmt.Println("Jogador vencedor: ", id)

				close(match)

				return
			}
		} 
		// If the player didn't score, print the player's error
		else {
			fmt.Println("Jogador ", id, " errou!")
		}

		// Start the other player's turn
		match <- true
	}
}

// Main function
func main() {
	// Create a channel
	match := make(chan bool)

	// Number of points to win
	PointsToWin := 10

	// Add two goroutines to wait group
	waitGroup.Add(2)

	// Start two players
	go player(1, match, PointsToWin, 0)
	go player(2, match, PointsToWin, 0)

	// Start the match
	match <- true

	// Wait for the match to finish
	waitGroup.Wait()
}
