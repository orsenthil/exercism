// Package dotdsl provides a parser for the DOT language, which is used to describe graphs.
// It allows you to parse DOT files and extract the graph structure, including nodes, edges, and their properties.
package dotdsl

import (
	"errors"
	"maps"
	"slices"
	"strconv"
	"strings"
)

// Properties holds the properties of a node or edge.
// The values can be int, bool or string.
type Properties map[string]any

// Graph stores the parts of a dot graph.
// All entities are stored as a Properties map (`nil` Properties when none set)
// attrs is the Properties for the entire Graph, vs a specific node or edge.
type Graph struct {
	nodes map[string]Properties
	edges map[string]Properties
	attrs Properties
}

// Parse creates a Graph from a text blob.
func Parse(data string) (*Graph, error) {
	toks := tokenize(stripComments(data))
	if len(toks) < 2 || toks[0] != "graph" || toks[1] != "{" {
		return nil, errors.New("invalid graph")
	}

	stmts := splitStatements(innerTokens(toks))

	g := &Graph{}

	for _, stmt := range stmts {
		ids, attrs, err := parseStatement(stmt)
		if err != nil {
			return nil, err
		}
		switch {
		case len(ids) == 0:
			if g.attrs == nil {
				g.attrs = Properties{}
			}
			mergeProps(g.attrs, attrs)
		case len(ids) == 1:
			if g.nodes == nil {
				g.nodes = map[string]Properties{}
			}
			if g.nodes[ids[0]] == nil {
				g.nodes[ids[0]] = attrs
			} else {
				mergeProps(g.nodes[ids[0]], attrs)
			}
		default:
			if g.nodes == nil {
				g.nodes = map[string]Properties{}
			}
			if g.edges == nil {
				g.edges = map[string]Properties{}
			}
			for i := range ids {
				if g.nodes[ids[i]] == nil {
					g.nodes[ids[i]] = nil
				}
			}
			for i := 0; i+1 < len(ids); i++ {
				a, b := ids[i], ids[i+1]
				if a > b {
					a, b = b, a
				}
				key := "{" + a + " " + b + "}"
				if g.edges[key] == nil {
					g.edges[key] = attrs
				} else {
					mergeProps(g.edges[key], attrs)
				}
			}
		}

	}
	return g, nil
}

func mergeProps(dst, src Properties) {
	maps.Copy(dst, src)
}

func stripComments(data string) string {
	lines := strings.Split(data, "\n")
	for i, line := range lines {
		// find the earlist comment marker and truncate
		c1 := strings.Index(line, "//")
		if c1 != -1 {
			line = line[:c1]
			lines[i] = line
		}

		c2 := strings.Index(line, "#")
		if c2 != -1 {
			line = line[:c2]
			lines[i] = line
		}
	}

	return strings.Join(lines, " ")
}

func isWordChar(b byte) bool {
	return ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9')
}

func innerTokens(tokens []string) []string {
	var first int
	var last int
	for idx, token := range tokens {
		if token == "{" {
			first = idx
			continue
		}
		if token == "}" {
			last = idx
			break
		}
	}

	return tokens[first+1 : last]
}

func splitStatements(tokens []string) [][]string {
	var statements [][]string
	var current []string
	for _, token := range tokens {
		if token == ";" {
			statements = append(statements, current)
			current = nil
		} else {
			current = append(current, token)
		}
	}
	return statements
}

func tokenize(data string) []string {
	var tokens []string
	i := 0
	for i < len(data) {
		if data[i] == ' ' {
			i += 1
			continue
		}

		if data[i] == '-' {
			if i < len(data)-1 {
				i += 1
				if data[i] == '-' {
					tokens = append(tokens, "--")
					i += 1
					continue
				}
			}
		}

		if data[i] == '"' {
			start := i
			i++ // skip opening quote
			for i < len(data) && data[i] != '"' {
				i++
			}
			i++ // skip closing quote
			tokens = append(tokens, data[start:i])
			continue
		}

		if isWordChar(data[i]) {
			start := i
			for i < len(data) && isWordChar(data[i]) {
				i++
			}
			tokens = append(tokens, data[start:i])
			continue
		}
		tokens = append(tokens, string(data[i]))

		i++
	}

	return tokens
}

func parseAttrs(tokens []string) (Properties, error) {
	if len(tokens) == 0 || len(tokens)%3 != 0 {
		return nil, errors.New("invalid attribute")
	}
	props := Properties{}
	for i := 0; i < len(tokens); i += 3 {
		if tokens[i+1] != "=" {
			return nil, errors.New("invalid attribute")
		}
		key := tokens[i]
		valueTok := tokens[i+2]
		props[key] = parseValue(valueTok)
	}

	return props, nil
}

func parseValue(valueTok string) any {
	if strings.HasPrefix(valueTok, `"`) {
		return valueTok[1 : len(valueTok)-1] // strip quotes
	}
	if valueTok == "true" {
		return true
	}
	if valueTok == "false" {
		return false
	}
	n, err := strconv.Atoi(valueTok)
	if err != nil {
		return valueTok
	}
	return n
}

func parseStatement(stmt []string) (ids []string, attrs Properties, err error) {
	first := slices.Index(stmt, "[")

	region := stmt
	if first != -1 {
		region = stmt[:first]
		attrs, err = parseAttrs(stmt[first+1 : len(stmt)-1])
		if err != nil {
			return nil, nil, err
		}
	}
	if len(region) > 0 {
		if len(region)%2 == 0 {
			return nil, nil, errors.New("invalid edge")
		}
		for i := 1; i < len(region); i += 2 {
			if region[i] != "--" {
				return nil, nil, errors.New("invalid edge")
			}
		}
	}

	for i := 0; i < len(region); i += 2 {
		id := region[i]
		for j := 0; j < len(id); j++ {
			if !isWordChar(id[j]) {
				return nil, nil, errors.New("node name must be alphanumeric")
			}
		}
		ids = append(ids, region[i])
	}

	return ids, attrs, nil
}
