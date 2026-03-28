package kindergarten

import (
	"fmt"
	"sort"
	"strings"
)

// Define the Garden type here.

type Garden struct {
	diagram  string
	children []string
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	// define a new Garden
	g := new(Garden)

	// check for an empty diagram
	if len(diagram) == 0 {
		return nil, fmt.Errorf("empty diagram")
	}

	// if the first line of the diagram is not empty, return an error
	if diagram[0] != '\n' {
		return nil, fmt.Errorf("first line of diagram must be empty")
	}

	// check for mismatched rows
	parts := strings.Split(diagram, "\n")
	if len(parts) != 3 {
		return nil, fmt.Errorf("mismatched rows")
	}

	// check for mismatched columns
	if len(parts[1]) != len(parts[2]) {
		return nil, fmt.Errorf("mismatched columns")
	}

	// odd number of cups
	if len(parts[1])%2 != 0 || len(parts[2])%2 != 0 {
		return nil, fmt.Errorf("odd number of cups")
	}

	// duplicate children names
	for i := 0; i < len(children); i++ {
		for j := i + 1; j < len(children); j++ {
			if children[i] == children[j] {
				return nil, fmt.Errorf("duplicate children names")
			}
		}
	}

	// invalid cup codes
	for i := 0; i < len(parts[1]); i++ {
		if parts[1][i] != 'R' && parts[1][i] != 'G' && parts[1][i] != 'C' && parts[1][i] != 'V' {
			return nil, fmt.Errorf("invalid cup codes")
		}
	}

	for i := 0; i < len(parts[2]); i++ {
		if parts[2][i] != 'R' && parts[2][i] != 'G' && parts[2][i] != 'C' && parts[2][i] != 'V' {
			return nil, fmt.Errorf("invalid cup codes")
		}
	}

	// sorted children
	sortedChildren := make([]string, len(children))
	copy(sortedChildren, children)
	sort.Strings(sortedChildren)

	// assign the diagram and children to the Garden

	g.diagram = diagram
	g.children = sortedChildren

	// return the new Garden
	return g, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	// define the plants slice
	var plants []string

	// define the ok boolean
	var ok bool

	// define the cups string
	var cups string

	// define the cup codes string
	var cupCodes string

	// define the cup codes map
	cupCodesMap := map[byte]string{
		'R': "radishes",
		'G': "grass",
		'C': "clover",
		'V': "violets",
	}

	// get the two row
	parts := strings.Split(g.diagram, "\n")

	for i := 0; i < len(parts); i++ {
		parts[i] = strings.TrimSpace(parts[i])
	}

	for i := 0; i < len(g.children); i++ {
		if g.children[i] == child {
			cups = parts[1][2*i : 2*i+2]
			// next row
			cups += parts[2][2*i : 2*i+2]

			ok = true

			break
		}
	}

	if !ok {
		return nil, ok
	}

	// get the cup codes for the child
	for i := 0; i < len(cups); i++ {
		cupCodes += string(cups[i])
	}

	// get the plants for the child
	for i := 0; i < len(cupCodes); i++ {
		plants = append(plants, cupCodesMap[cupCodes[i]])
	}

	// return the plants and ok
	return plants, ok

}
