package allergies

var scoreToAllergen = map[uint]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

var allergenToScore = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func Allergies(allergies uint) []string {
	var allergens []string
	var score uint
	for score = 1; score <= 128; score *= 2 {
		if score&allergies == score {
			allergens = append(allergens, scoreToAllergen[score])
		}
	}
	return allergens
}

func AllergicTo(allergies uint, allergen string) bool {
	if value, found := allergenToScore[allergen]; found {
		return allergies&value == value
	}
	return false
}
