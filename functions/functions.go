package functions

import (
	"math/rand"
	"time"
)

func RandomValue(nb int) int {
	rand.Seed(time.Now().UnixNano())
	randomNb := rand.Intn(nb)
	return randomNb
}