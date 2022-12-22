package main

import (
	"fmt"
	"sync"
	"tennis-concurrency/utils"
)

var waitGroup sync.WaitGroup

func player() {
	randomNumber := utils.RandomInt()

	if randomNumber%2 != 0 {
		fmt.Println(randomNumber)
	} else {
		fmt.Println("Missed")
	}
}

func main() {
	player()
}
