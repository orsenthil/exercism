package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

// Define the Robot type here.

type Robot struct {
	name string
}

const (
	ordA           = 65
	alphabetLength = 26
	numberLimit    = 1000
	namePoolSize   = alphabetLength * alphabetLength * numberLimit
)

var namePool = generateNamePool(namePoolSize)

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		if len(namePool) == 0 {
			return "", errors.New("No name available")
		}
		r.Reset()
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = numberToText(namePool[0])
	namePool = namePool[1:]
}

func numberToText(num int) string {
	letterPart := num / numberLimit
	part1 := letterPart / alphabetLength
	part2 := letterPart % alphabetLength
	part3 := num % numberLimit
	return fmt.Sprintf("%v%v%03d", numberToLetter(part1), numberToLetter(part2), part3)
}

func generateNamePool(size int) []int {
	pool := make([]int, size)
	for i := 0; i < size; i++ {
		pool[i] = i
	}
	rand.Shuffle(size, func(i, j int) {
		pool[i], pool[j] = pool[j], pool[i]
	})
	return pool
}

func numberToLetter(number int) string {
	return string(rune(number + ordA))
}
