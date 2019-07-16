package tooee

import ()

func Initiative(players int) map[int]int {
	var results map[int]int
	for i := 1; i <= players; i++ {
		roll := RollDice(DiceRoll{10, 1})
		results[i] = roll
	}
	return results
}

func Thaco(attacker Character, defender Character) (hit bool, epic bool) {
	thacoRoll := RollDice(DiceRoll{20, 1})
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
