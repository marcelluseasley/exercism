package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(list []string) FreqMap {
	resChannel := make(chan FreqMap, len(list))
	for _, s := range list {
		s := s
		go func() {
			resChannel <- Frequency(s)
		}()
	}

	res := make(FreqMap)
	for range list {
		for letter, freq := range <- resChannel {
			res[letter] += freq
		}
	}
	return res
}
