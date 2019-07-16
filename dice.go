package tooee

import (
	"math/rand"
	"time"
)

type diceRoll struct {
	Sides int
	Rolls int
}

func rollDice(dice diceRoll) int {
	rand.Seed(time.Now().UnixNano())
	min := dice.Rolls
	max := dice.Sides * dice.Rolls
	var roll int
	for i := 0; i <= max; i++ {
		newRoll := rand.Intn(max-min) + min
		roll = roll + newRoll
	}
	return roll
}
