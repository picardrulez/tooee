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

func thaco(attacker character, defender character) (hit bool, epic bool) {
	thacoRoll := rollDice(diceRoll{20, 1})
	target := attacker.Thaco - defender.AC
	if thacoRoll >= target {
		if thacoRoll == 20 {
			return true, true
		}
		return true, false
	} else {
		if thacoRoll == 1 {
			return false, true
		} else {
			return false, true
		}
	}
}
