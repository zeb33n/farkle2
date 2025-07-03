package core

import (
	"testing"
)

func runScore(dice []int, expectedScore int, expectedDice int, t *testing.T) {
	s, d, _ := Score(dice)
	if s != expectedScore {
		t.Errorf("Score: %d != %d for %v", s, expectedScore, dice)
	}
	if d != expectedDice {
		t.Errorf("Dice: %d != %d for %v", d, expectedDice, dice)
	}
}

func TestScore(t *testing.T) {
	runScore([]int{1, 1, 2, 2, 3, 3}, 1500, 0, t)
	runScore([]int{4, 3, 2, 1, 5, 6}, 2000, 0, t)
	runScore([]int{1, 1, 5, 4, 6, 2}, 250, 3, t)
	runScore([]int{6}, 0, 1, t)
	runScore([]int{5}, 50, 0, t)
}
