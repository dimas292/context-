package utility

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber(min, max int) (number int){

	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max - min) + min
}