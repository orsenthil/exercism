package luhn

import "strconv"

func Valid(id string) bool {
	var nummap = make(map[int]int)

	var idx = 0

	for _, c := range id {
		if string(c) == " " {
			continue
		}

		num, err := strconv.Atoi(string(c))

		if err != nil {
			return false
		}
		nummap[idx] = num
		idx += 1
	}

	if len(nummap) <= 1 {
		return false
	}

	for i := len(nummap) - 2; i >= 0; i -= 2 {
		if nummap[i]*2 > 9 {
			nummap[i] = (nummap[i] * 2) - 9
		} else {
			nummap[i] = nummap[i] * 2
		}
	}

	var result = 0

	for i := 0; i < len(nummap); i++ {
		result += nummap[i]
	}

	if result%10 == 0 {
		return true
	}

	return false
}
