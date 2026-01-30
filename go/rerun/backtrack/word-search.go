package backtrack

import (
	"unicode"
)

type Coordinates struct {
	r int
	c int
}

func makeCoord(r, c int) Coordinates {
	return Coordinates{
		r,
		c,
	}
}

func exist(board [][]byte, word string) bool {
	rMax, cMax := len(board), len(board[0])

	isValidCoord := func(coord Coordinates, currIdx int) bool {
		r, c := coord.r, coord.c
		if r < 0 || r >= rMax || c < 0 || c >= cMax {
			return false
		}
		if board[r][c] == '#' {
			return false
		}

		wLetter := byte(unicode.ToLower(rune(word[currIdx])))
		boardLetter := byte(unicode.ToLower(rune(board[r][c])))
		if wLetter != boardLetter {
			return false
		}
		return true
	}
	var backtrack func(coord Coordinates, currIdx int) bool
	backtrack = func(coord Coordinates, currIdx int) bool {
		if currIdx == len(word) {
			return true
		}
		if !isValidCoord(coord, currIdx) {
			return false
		}
		r, c := coord.r, coord.c
		temp := board[r][c]
		board[r][c] = '#'
		up := makeCoord(r+1, c)
		down := makeCoord(r-1, c)
		left := makeCoord(r, c+1)
		right := makeCoord(r, c-1)
		found := backtrack(up, currIdx+1) || backtrack(down, currIdx+1) || backtrack(left, currIdx+1) || backtrack(right, currIdx+1)
		board[r][c] = temp
		return found
	}

	for r := range board {
		for c := range board[r] {
			if backtrack(makeCoord(r, c), 0) {
				return true
			}
		}
	}
	return false
}
