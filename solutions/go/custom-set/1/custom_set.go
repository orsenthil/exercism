package stringset

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.

type Set struct {
	elements *[]string
}

func New() Set {
	empty := make([]string, 0)
	return Set{&empty}
}

func NewFromSlice(l []string) Set {
	var newSet Set
	newSet = New()
	for _, elem := range l {
		newSet.Add(elem)
	}
	return newSet
}

func (s Set) String() string {
	var str string
	str = "{"
	for i, elem := range *s.elements {
		if i == len(*s.elements)-1 {
			str += "\"" + elem + "\""
		} else {
			str += "\"" + elem + "\", "
		}
	}
	str += "}"
	return str
}

func (s Set) IsEmpty() bool {
	if len(*s.elements) == 0 {
		return true
	}

	return false
}

func (s Set) Has(elem string) bool {
	for _, e := range *s.elements {
		if e == elem {
			return true
		}
	}

	return false
}

func (s Set) Add(elem string) {
	if !s.Has(elem) {
		*s.elements = append(*s.elements, elem)
	}
}

func Subset(s1, s2 Set) bool {
	var isSubset bool
	isSubset = true
	for _, e1 := range *s1.elements {
		if !s2.Has(e1) {
			isSubset = false
			return isSubset
		}
	}
	return isSubset
}

func Disjoint(s1, s2 Set) bool {
	for _, e1 := range *s1.elements {
		if s2.Has(e1) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(*s1.elements) != len(*s2.elements) {
		return false
	}

	for _, e1 := range *s1.elements {
		if !s2.Has(e1) {
			return false
		}
	}
	return true
}

func Intersection(s1, s2 Set) Set {
	var intersection Set
	intersection = New()
	for _, e1 := range *s1.elements {
		if s2.Has(e1) {
			intersection.Add(e1)
		}
	}
	return intersection
}

func Difference(s1, s2 Set) Set {
	var difference Set
	difference = New()
	for _, e1 := range *s1.elements {
		if !s2.Has(e1) {
			difference.Add(e1)
		}
	}
	return difference
}

func Union(s1, s2 Set) Set {
	var union Set
	union = New()
	for _, e1 := range *s1.elements {
		union.Add(e1)
	}
	for _, e2 := range *s2.elements {
		union.Add(e2)
	}
	return union
}
