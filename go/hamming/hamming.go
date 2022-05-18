package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("not equal length or length of zero present")
	}
	if len(a) == 0 {
		return 0, nil
	}
	nonMatches := 0

	for i, ch := range a {
		if ch != rune(b[i]) {
			nonMatches++
		}

	}
	return nonMatches, nil
}