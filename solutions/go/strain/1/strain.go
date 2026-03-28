package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var result Ints
	for _, v := range i {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var result Ints
	for _, v := range i {
		if !filter(v) {
			result = append(result, v)
		}
	}
	return result
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var result Lists
	for _, v := range l {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var result Strings
	for _, v := range s {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}
