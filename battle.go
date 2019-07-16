package tooee

import ()

func initiative(players int) map[int]int {
	var results map[int]int
	for i := 1; i <= players; i++ {
		roll := rollDice(diceRoll{10, 1})
		results[i] = roll
	}
	return results
}
