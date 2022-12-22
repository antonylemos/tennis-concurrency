package main

import (
	"fmt"
	"sync"
	"tennis-concurrency/utils"
)

var waitGroup sync.WaitGroup

func player(id int, match chan bool, pointsToWin int, playerPoints int) {
	defer waitGroup.Done()

	for {
		finish := <-match

		if !finish {
			return
		}

		fmt.Println("Jogador ", id, " começou a jogada!")

		playerScored := utils.MakePoint()

		if playerScored {
			playerPoints++

			fmt.Println("Jogador ", id, " marcou ", playerPoints, " pontos!")

			if playerPoints == pointsToWin {
				fmt.Println("Jogador vencedor: ", id)

				close(match)

				return
			}
		} else {
			fmt.Println("Jogador ", id, " errou!")
		}

		match <- true
	}
}

func main() {
	match := make(chan bool)

	PointsToWin := 10

	waitGroup.Add(2)

	go player(1, match, PointsToWin, 0)
	go player(2, match, PointsToWin, 0)

	match <- true

	waitGroup.Wait()
}
