package main

import (
	"maps"
	"math/rand"
	"reflect"
	"slices"
	"sort"
)

type Player struct {
	name  string
	score int
}

type GameState struct {
	num_dice       int
	current_score  int
	current_player int
	players        []Player
}

func take_turn(gamestate GameState) GameState {
	dice := make([]int, gamestate.num_dice)
	for i := 0; i < gamestate.num_dice; i++ {
		dice[i] = rand.Intn(6) + 1
		println(dice[i])
	}
	round_score, num_dice := score(dice)
	println(round_score)
	println(num_dice)
	return gamestate
}

func score(dice []int) (int, int) {
	sort.Slice(dice, func(i, j int) bool { return dice[i] < dice[j] })
	if reflect.DeepEqual(dice, []int{1, 2, 3, 4, 5, 6}) {
		return 2000, 0
	}
	value_counts := make(map[int]int)
	for i := range dice {
		value_counts[dice[i]]++
	}
	counts := slices.Collect(maps.Values(value_counts))
	if (len(counts) == 2 && all_equal(counts, 3)) || (len(counts) == 2 && all_equal(counts, 2)) {
		return 1500, 0
	}

	score := 0
	num_dice := len(dice)
	for value, count := range value_counts {
		if count >= 4 {
			score += 1000 * (count - 3)
			num_dice -= count
		} else if count == 3 {
			if value == 1 {
				score += 300
			} else {
				score += value * 100
			}
			num_dice -= count
		} else {
			if value == 1 {
				score += count * 100
				num_dice -= count
			}
			if value == 5 {
				score += count * 50
				num_dice -= count
			}
		}
	}
	return score, num_dice
}

func all_equal(s []int, value int) bool {
	for i := range s {
		if i != value {
			return false
		}
	}
	return true
}

func main() {
	gamestate := GameState{6, 0, 0, []Player{{name: "bob", score: 0}}}
	take_turn(gamestate)
}
