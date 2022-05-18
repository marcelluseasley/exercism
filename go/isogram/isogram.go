package isogram

import (
	"strings"

)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	if word == "" {
		return true
	}
	
	wordMap := make(map[rune]int)

	for _, r := range word {
		if r == '-' || r == ' ' {
			continue
		}
		if _, ok := wordMap[r]; !ok {
			wordMap[r] = 1
		} else {
			return false
		}
	}

	return true


}
