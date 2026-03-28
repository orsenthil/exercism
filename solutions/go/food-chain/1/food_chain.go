package foodchain

import "fmt"

var animals = []string{"fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}

var reactions = map[string]string{
	"spider": "It wriggled and jiggled and tickled inside her.",
	"bird":   "How absurd to swallow a bird!",
	"cat":    "Imagine that, to swallow a cat!",
	"dog":    "What a hog, to swallow a dog!",
	"goat":   "Just opened her throat and swallowed a goat!",
	"cow":    "I don't know how she swallowed a cow!",
	"horse":  "She's dead, of course!",
}

// Example helper concept (you'll need to implement the logic)
func buildChain(index int) string {
	animal := animals[index]
	switch animal {
	case "fly":
		return "I don't know why she swallowed the fly. Perhaps she'll die."
	case "spider":
		spider := "She swallowed the spider to catch the fly."
		previous := buildChain(index - 1)
		return spider + "\n" + previous
	case "bird":
		bird := "She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her."
		previous := buildChain(index - 1)
		return bird + "\n" + previous
	case "cat":
		cat := "She swallowed the cat to catch the bird."
		previous := buildChain(index - 1)
		return cat + "\n" + previous
	case "dog":
		dog := "She swallowed the dog to catch the cat."
		previous := buildChain(index - 1)
		return dog + "\n" + previous
	case "goat":
		goat := "She swallowed the goat to catch the dog."
		previous := buildChain(index - 1)
		return goat + "\n" + previous
	case "cow":
		cow := "She swallowed the cow to catch the goat."
		previous := buildChain(index - 1)
		return cow + "\n" + previous
	}
	return ""
}

func Verse(v int) string {
	animal := animals[v-1]
	verse := fmt.Sprintf("I know an old lady who swallowed a %s.", animal)
	if reaction, exists := reactions[animal]; exists {
		verse += "\n" + reaction
	}
	if animal != "horse" {
		verse += "\n" + buildChain(v-1)
	}
	return verse
}

func Verses(start, end int) string {
	result := ""
	for i := start; i <= end; i++ {
		if i == start {
			result += Verse(i)
		} else {
			result += "\n\n" + Verse(i)
		}
	}

	return result
}

func Song() string {
	return Verses(1, 8)
}
