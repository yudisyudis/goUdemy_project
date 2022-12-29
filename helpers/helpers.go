package helpers

import (
	"math/rand"
	"time"
)

func AcakNomor(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}