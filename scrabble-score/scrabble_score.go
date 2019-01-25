package scrabble

import (
	"strings"
)

func Score(word string) int {
	var total int
	for _, letter := range strings.ToLower(word) {
		switch letter {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			total += 1
		case 'd', 'g':
			total += 2
		case 'b', 'c', 'm', 'p':
			total += 3
		case 'f', 'h', 'v', 'w', 'y':
			total += 4
		case 'k':
			total += 5
		case 'j', 'x':
			total += 8
		case 'q', 'z':
			total += 10
		}
	}

	return total
}
