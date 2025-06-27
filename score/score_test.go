package score

import (
	"testing"
)

func runScore(dice []int, expected_score int, expected_dice int, t *testing.T) {
	s, d := Score(dice)
	if s != expected_score {
		t.Errorf("Score: %d != %d for %v", s, expected_score, dice)
	}
	if d != expected_dice {
		t.Errorf("Dice: %d != %d for %v", d, expected_dice, dice)
	}
}

func TestScore(t *testing.T) {
	// runScore([]int{1, 1, 2, 2, 3, 3}, 1500, 0, t)
	// runScore([]int{4, 3, 2, 1, 5, 6}, 2000, 0, t)
	// runScore([]int{1, 1, 5, 4, 6, 2}, 250, 3, t)
	// runScore([]int{6}, 0, 1, t)
	runScore([]int{5}, 50, 0, t)
}
