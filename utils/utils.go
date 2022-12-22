package utils

import (
	"math/rand"
	"time"
)

func RandomInt() int {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	return random.Int()
}

func MakePoint() bool {
	randomNumber := RandomInt()

	return randomNumber%2 == 0
}
