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
// ConcurrentFrequency takes an array of strings and returns a frequency map of the total.
func ConcurrentFrequency(letters []string) FreqMap {
	total := make(FreqMap)
	channel := make(chan FreqMap)

	for _, l := range letters {
		go func(letter string) {
			channel <- Frequency(letter)
		}(l)
	}

	for i := 0; i < len(letters); i++ {
		total.count(<-channel)
	}

	return  total
}
// count adds the rune frequencies in the other map to this FreqMap
func (mapper FreqMap) count(channel FreqMap) {
	for r, f := range channel {
		mapper[r] += f
	}
}


