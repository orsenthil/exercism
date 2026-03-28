package twelve

import (
	"fmt"
	"strings"
)

func Verse(i int) string {
	var day string
	var line string
	var lyrics []string

	switch i {
	case 1:
		day = "first"
	case 2:
		day = "second"
	case 3:
		day = "third"
	case 4:
		day = "fourth"
	case 5:
		day = "fifth"
	case 6:
		day = "sixth"
	case 7:
		day = "seventh"
	case 8:
		day = "eighth"
	case 9:
		day = "ninth"
	case 10:
		day = "tenth"
	case 11:
		day = "eleventh"
	case 12:
		day = "twelfth"
	}

	lyrics = append(lyrics, fmt.Sprintf("On the %s day of Christmas my true love gave to me:", day))

	for n := i; n >= 1; n-- {
		switch n {
		case 1:
			line = " a Partridge in a Pear Tree"
		case 2:
			line = " two Turtle Doves"
		case 3:
			line = " three French Hens"
		case 4:
			line = " four Calling Birds"
		case 5:
			line = " five Gold Rings"
		case 6:
			line = " six Geese-a-Laying"
		case 7:
			line = " seven Swans-a-Swimming"
		case 8:
			line = " eight Maids-a-Milking"
		case 9:
			line = " nine Ladies Dancing"
		case 10:
			line = " ten Lords-a-Leaping"
		case 11:
			line = " eleven Pipers Piping"
		case 12:
			line = " twelve Drummers Drumming"
		}
		lyrics = append(lyrics, line)
		if n == 1 {
			lyrics = append(lyrics, ".")
		} else {
			lyrics = append(lyrics, ",")
		}
		if n == 2 {
			lyrics = append(lyrics, " and")
		}
	}
	return strings.Join(lyrics, "")

}

func Song() string {
	var lyrics []string
	for i := 1; i <= 12; i++ {
		lyrics = append(lyrics, Verse(i))
	}
	return strings.Join(lyrics, "\n")
}
