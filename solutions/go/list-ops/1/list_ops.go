package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	var result int = initial
	for _, value := range s {
		result = fn(result, value)
	}
	return result
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	// Implementing Foldr
	var result int = initial
	for _, value := range s.Reverse() {
		result = fn(value, result)
	}
	return result
}

func (s IntList) Filter(fn func(int) bool) IntList {
	// Implementing Filter using Foldl
	var result IntList = make([]int, 0)
	for _, value := range s {
		if fn(value) {
			result = append(result, value)
		}
	}
	return result
}

func (s IntList) Map(fn func(int) int) IntList {
	// Implementing Map using Foldl
	var result IntList = make([]int, 0)
	for _, value := range s {
		result = append(result, fn(value))
	}
	return result
}

func (s IntList) Reverse() IntList {
	// Implementing Reverse
	var result IntList = make([]int, 0)
	for _, value := range s {
		result = append(IntList{value}, result...)
	}
	return result
}

func (s IntList) Append(lst IntList) IntList {
	// Implementing Append
	var result IntList = make([]int, 0)
	for _, value := range s {
		result = append(result, value)
	}
	for _, value := range lst {
		result = append(result, value)
	}
	return result
}

func (s IntList) Concat(lists []IntList) IntList {
	// Implementing Concat
	var result IntList = make([]int, 0)
	for _, value := range s {
		result = append(result, value)
	}
	for _, list := range lists {
		for _, value := range list {
			result = append(result, value)
		}
	}
	return result
}

func (s IntList) Length() int {
	var result int = 0
	for range s {
		result++
	}
	return result
}
