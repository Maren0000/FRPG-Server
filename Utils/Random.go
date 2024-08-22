package Utils

import (
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func GenRandId() string {
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func IsHit(int) bool {
	if rand.Int63n(2) == 0 {
		return false
	} else {
		return true
	}
}

func RandomBGM(BGMs []int) int {
	return BGMs[rand.Intn(len(BGMs))]
}

func GenerateColorID() int64 {
	return rand.Int63n(3) + 1
}
