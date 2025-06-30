package score

import (
	"maps"
	"reflect"
	"slices"
	"sort"
)

func Score(dice []int) (int, int, []int) {
	sort.Slice(dice, func(i, j int) bool { return dice[i] < dice[j] })
	if reflect.DeepEqual(dice, []int{1, 2, 3, 4, 5, 6}) {
		return 2000, 0, []int{0, 1, 2, 3, 4, 5}
	}
	value_counts := make(map[int]int)
	positions := make(map[int][]int)
	for i, e := range dice {
		value_counts[e]++
		positions[e] = append(positions[e], i)
	}
	counts := slices.Collect(maps.Values(value_counts))
	if (len(counts) == 2 && all_equal(counts, 3)) || (len(counts) == 3 && all_equal(counts, 2)) {
		return 1500, 0, []int{0, 1, 2, 3, 4, 5}
	}

	score := 0
	num_dice := len(dice)
	scoring_positions := []int{}
	for value, count := range value_counts {
		if count >= 4 {
			score += 1000 * (count - 3)
			num_dice -= count
			scoring_positions = positions[value]
		} else if count == 3 {
			if value == 1 {
				score += 300
			} else {
				score += value * 100
			}
			num_dice -= count
			scoring_positions = positions[value]
		} else {
			if value == 1 {
				score += count * 100
				num_dice -= count
				scoring_positions = positions[value]
			}
			if value == 5 {
				score += count * 50
				num_dice -= count
				scoring_positions = positions[value]
			}
		}
	}
	return score, num_dice, scoring_positions
}

func all_equal(s []int, value int) bool {
	for _, e := range s {
		if e != value {
			return false
		}
	}
	return true
}
