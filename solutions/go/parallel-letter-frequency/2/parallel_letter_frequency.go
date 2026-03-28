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
func ConcurrentFrequency(l []string) FreqMap {
	comm := make(chan FreqMap)
	defer close(comm)

	for _, s := range l {
		go func(s string) {
			comm <- Frequency(s)
		}(s)
	}

	output := <-comm

	for i := 1; i < len(l); i++ {
		for k, v := range <-comm {
			output[k] += v
		}
	}

	return output
}
