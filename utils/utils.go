package utils

// Import packages
import (
	"math/rand"
	"time"
)

// Generate a random number
func RandomInt() int {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	return random.Int()
}

// Make a point
func MakePoint() bool {
	randomNumber := RandomInt()

	// If the random number is even, make a point
	return randomNumber%2 == 0
}
