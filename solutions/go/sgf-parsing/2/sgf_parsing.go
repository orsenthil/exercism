package sgfparsing

import "errors"

// Node represents an SGF node with properties and child nodes.
type Node struct {
	Properties map[string][]string
	Children   []*Node
}

type parser struct {
	input string
	pos int
}

func (p *parser) current() byte {
	if p.pos >= len(p.input) {
		return 0
	}
	return p.input[p.pos]
}

func (p *parser) parseValue() string {
	p.pos++ // skip [
	var value []byte
	for p.current() != ']' {
		switch p.current() {
		case '\\':
			p.pos++ // skip \
			next := p.current()
			switch next {
			case '\n':
				// do nothing
			case '\t', ' ':
				value = append(value, ' ')
			default:
				value = append(value, next)
			}
		case '\t':
			value = append(value, ' ')
		default:
			value = append(value, p.current())
		}
		p.pos++ // skip the current character
	}
	p.pos++ // skip ]
	return string(value)
}

func (p * parser) parseProperty() (string, []string, error) {
	var key string
	var values []string
	for p.current() >= 'A' && p.current() <= 'Z' {
		key += string(p.current())
		p.pos++
	}
	if p.current() >= 'a' && p.current() <= 'z' {
		return "", nil, errors.New("property must be in uppercase")
	}
	if p.current() != '[' {
		return "", nil, errors.New("property without delimiter")
	}
	for p.current() == '[' {
		values = append(values, p.parseValue())
	}
	return key, values, nil
}

func (p * parser) parseNode() (*Node, error) {
	p.pos++ // skip ;

	node := &Node{
		Properties: make(map[string][]string),
		Children:   []*Node{},
	}

	for p.current() >= 'A' && p.current() <= 'Z' {
		key, values, err := p.parseProperty()
		if err != nil {
			return nil, err
		}
		node.Properties[key] = values
	}

	if p.current() >= 'a' && p.current() <= 'z' {
		return nil, errors.New("property must be in uppercase")
	}

	return node, nil
}

func (p * parser) parseTree() (*Node, error) {
	p.pos++ // skip (
	if p.current() != ';' {
		return nil, errors.New("tree with no nodes")
	}
	root, err := p.parseNode()
	if err != nil {
		return nil, err
	}

	currentNode := root
	for p.current() == '('  || p.current() == ';'{
		if p.current() == ';' {
			node, err := p.parseNode()
			if err != nil {
				return nil, err
			}
			currentNode.Children = append(currentNode.Children, node)
			currentNode = node
		} else if p.current() == '(' {
			childTree, err := p.parseTree()
			if err != nil {
				return nil, err
			}
			currentNode.Children = append(currentNode.Children, childTree)
		}
	}
	p.pos++ // skip )
	return root, nil
}

// Parse decodes an SGF string and returns the root node of the tree.
func Parse(encoded string) (*Node, error) {
	if len(encoded) == 0 {
		return nil, errors.New("tree missing")
	}
	p := &parser{input: encoded, pos: 0}
	return p.parseTree()
}
