package series

func All(n int, s string) []string {
	var result []string
	for i := 0; i <= len(s)-n; i++ {
		result = append(result, s[i:i+n])
	}
	return result
}

func UnsafeFirst(n int, s string) string {
	var result string
	for i := 0; i < n; i++ {
		result += string(s[i])
	}
	return result
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	return UnsafeFirst(n, s), true
}
