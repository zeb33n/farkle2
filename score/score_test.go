package score

import (
	"testing"
)

func TestScore(t *testing.T) {
	s, _ := Score([]int{1, 1, 2, 2, 3, 3})
	if s != 1500 {
		t.Errorf("Score: %d != 1500", s)
	}
}
