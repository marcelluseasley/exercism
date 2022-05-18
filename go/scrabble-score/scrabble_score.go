package scrabble

import "strings"

func Score(word string) int {
	var score int
	word = strings.ToLower(word)
	for _, letter := range word {

		if strings.ContainsRune("aeioulnrst", letter) {
			score++
		} else if strings.ContainsRune("dg", letter) {
			score += 2
		} else if strings.ContainsRune("bcmp", letter) {
			score += 3
		} else if strings.ContainsRune("fhvwy", letter) {
			score += 4
		} else if strings.ContainsRune("k", letter) {
			score += 5
		} else if strings.ContainsRune("jx", letter) {
			score += 8
		} else if strings.ContainsRune("qz", letter) {
			score += 10
		}

	}
  return score
}
