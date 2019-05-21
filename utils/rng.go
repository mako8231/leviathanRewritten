package utils

import (
	"math/rand"
	"time"
)

//Return a random number between max and min values
func RandomRange(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
