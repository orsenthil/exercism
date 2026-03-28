package lineup

import "fmt"

func Format(name string, number int) string {
	var suffix string
	suffix = getSuffix(number)
	return fmt.Sprintf("%s, you are the %d%s customer we serve today. Thank you!", name, number, suffix)
}


func getSuffix(number int) string {
	if number % 100 == 11 || number % 100 == 12 || number % 100 == 13 {
		return "th"
	}
	switch number % 10 {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}