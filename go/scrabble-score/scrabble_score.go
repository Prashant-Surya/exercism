package scrabble

import (
	"strings"
)

func Score(s string) int {
	var values = []int{
		1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10,
	}
	s = strings.ToUpper(s)
	ascii := []rune(s)
	var a int = 65
	var score int = 0
	for i := 0; i < len(s); i++ {
		score += values[int(ascii[i])-a]
	}
	return score
}
