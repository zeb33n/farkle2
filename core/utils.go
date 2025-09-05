package core

import (
	"os"
)

type TourType int

const (
	ROUNDROBIN TourType = iota
	ELIMINATION
)

func WaitForKeyPress(verbose bool) string {
	b := make([]byte, 1)
	os.Stdin.Read(b)
	s := string(b)
	if verbose {
		if s != "\n" {
			print(s)
		}
	}
	return s
}

func nextSquare(i int) int {
	ui := uint32(i)
	// bit twiddling fun
	ui--
	ui |= ui >> 1
	ui |= ui >> 2
	ui |= ui >> 4
	ui |= ui >> 8
	ui |= ui >> 16
	ui++
	return int(ui)
}
