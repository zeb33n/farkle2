// Package core
package core

import (
	"maps"
	"reflect"
	"slices"
)

func Score(dice []int) (int, int, []int) {
	slices.Sort(dice)
	if reflect.DeepEqual(dice, []int{1, 2, 3, 4, 5, 6}) {
		return 2000, 0, []int{0, 1, 2, 3, 4, 5}
	}
	valueCounts := make(map[int]int)
	positions := make(map[int][]int)
	for i, e := range dice {
		valueCounts[e]++
		positions[e] = append(positions[e], i)
	}
	counts := slices.Collect(maps.Values(valueCounts))
	if (len(counts) == 2 && allEqual(counts, 3)) || (len(counts) == 3 && allEqual(counts, 2)) {
		return 1500, 0, []int{0, 1, 2, 3, 4, 5}
	}

	score := 0
	numDice := len(dice)
	scoringPositions := []int{}
	for value, count := range valueCounts {
		if count >= 4 {
			score += 1000 * (count - 3)
			numDice -= count
			scoringPositions = append(scoringPositions, positions[value]...)
		} else if count == 3 {
			if value == 1 {
				score += 300
			} else {
				score += value * 100
			}
			numDice -= count
			scoringPositions = append(scoringPositions, positions[value]...)
		} else {
			if value == 1 {
				score += count * 100
				numDice -= count
				scoringPositions = append(scoringPositions, positions[value]...)
			}
			if value == 5 {
				score += count * 50
				numDice -= count
				scoringPositions = append(scoringPositions, positions[value]...)
			}
		}
	}
	return score, numDice, scoringPositions
}

func allEqual(s []int, value int) bool {
	for _, e := range s {
		if e != value {
			return false
		}
	}
	return true
}
