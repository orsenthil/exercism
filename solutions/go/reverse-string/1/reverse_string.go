package reverse

func Reverse(input string) string {
	var result string
	for _, v := range input {
		result = string(v) + result
	}
	return result
}
