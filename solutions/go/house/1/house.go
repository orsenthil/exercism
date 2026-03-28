package house

import "fmt"

var subjects = []string{
	"house",
	"malt",
	"rat",
	"cat",
	"dog",
	"cow with the crumpled horn",
	"maiden all forlorn",
	"man all tattered and torn",
	"priest all shaven and shorn",
	"rooster that crowed in the morn",
	"farmer sowing his corn",
	"horse and the hound and the horn",
}

var actions = []string{
	"that Jack built.",
	"that lay in",
	"that ate",
	"that killed",
	"that worried",
	"that tossed",
	"that milked",
	"that kissed",
	"that married",
	"that woke",
	"that kept",
	"that belonged to",
}

func buildLine(idx int) string {
	if idx == 0 {
		return fmt.Sprintf("that Jack built.")
	}

	if idx == 1 {
		return fmt.Sprintf("%s the %s %s", actions[idx], subjects[idx-1], buildLine(idx-1))
	}

	return fmt.Sprintf("%s the %s\n%s", actions[idx], subjects[idx-1], buildLine(idx-1))
}

func Verse(v int) string {
	if v == 1 {
		return fmt.Sprintf("This is the house that Jack built.")
	}
	return fmt.Sprintf("This is the %s\n%s", subjects[v-1], buildLine(v-1))
}

func Song() string {
	song := Verse(1)
	for i := 2; i <= len(subjects); i++ {
		song += "\n\n" + Verse(i)
	}
	return song
}
