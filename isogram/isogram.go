package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	letters := make(map[rune]bool)
	for _, letter := range strings.ToLower(word) {
		if letter == ' ' || letter == '-' {
			continue
		}
		if letters[letter] {
			return false
		} else {
			letters[letter] = true
		}
	}
	return true
}
