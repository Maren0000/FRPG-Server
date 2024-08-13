package Utils

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func GenRandId() string {
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func IsHit(int) bool {
	rand.Seed(time.Now().UTC().UnixNano())
	if rand.Int63n(2) == 0 {
		return false
	} else {
		return true
	}
}
