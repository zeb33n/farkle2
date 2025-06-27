package score

import (
	"testing"
)

func runScore(dice []int, expected_score int, expected_dice int, t *testing.T) {
	s, d := Score(dice)
	if s != expected_score {
		t.Errorf("Score: %d != %d", s, expected_score)
	}
	if d != expected_dice {
		t.Errorf("Dice: %d != %d", d, expected_dice)
	}
}

func TestScore(t *testing.T) {
	runScore([]int{1, 1, 2, 2, 3, 3}, 1500, 0, t)
	runScore([]int{4, 3, 2, 1, 5, 6}, 2000, 0, t)
	runScore([]int{1, 1, 5, 4, 6, 2}, 250, 3, t)
}
