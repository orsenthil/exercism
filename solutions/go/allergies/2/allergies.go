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

// Allergies returns a list of allergens that the person is allergic to
// based on the score
func Allergies(allergies uint) []string {
	var allergens []string
	var score uint
	for score = 1; score <= 128; score *= 2 {
		// score&allergies == score is a bitwise AND operation
		// if the result is equal to score, then the allergen is present
		// INSIGHT(senthil): this is a bitwise AND operation to check if a bit is set
		// Examples
		// 1 & 1 = 1
		// 2 & 1 = 0
		// 4 & 1 = 0
		// 8 & 1 = 0
		// 16 & 1 = 0
		// 32 & 1 = 0
		// 64 & 1 = 0
		if score&allergies == score {
			allergens = append(allergens, scoreToAllergen[score])
		}
	}
	return allergens
}

// AllergicTo returns true if the person is allergic to the allergen
func AllergicTo(allergies uint, allergen string) bool {
	// INSIGHT(senthil): This is a golang way to check if a key exists in a map
	if value, found := allergenToScore[allergen]; found {
		return allergies&value == value
	}
	return false
}
