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
