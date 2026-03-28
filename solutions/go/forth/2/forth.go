// Package forth provides functionality for evaluating simple Forth expressions.
package forth
import (
	"errors"
	"strconv"
	"strings"
)
const (
	notEnoughStack = "not enough numbers in stack to perform operation"
)
var builtinOps = map[string]func([]int) ([]int, error){
	"+": func(stack []int) ([]int, error) {
		if len(stack) < 2 {
			return stack, errors.New(notEnoughStack)
		}
		stack[len(stack)-2] = stack[len(stack)-2] + stack[len(stack)-1]
		return stack[:len(stack)-1], nil
	},
	"-": func(stack []int) ([]int, error) {
		if len(stack) < 2 {
			return stack, errors.New(notEnoughStack)
		}
		stack[len(stack)-2] = stack[len(stack)-2] - stack[len(stack)-1]
		return stack[:len(stack)-1], nil
	},
	"*": func(stack []int) ([]int, error) {
		if len(stack) < 2 {
			return stack, errors.New(notEnoughStack)
		}
		stack[len(stack)-2] = stack[len(stack)-2] * stack[len(stack)-1]
		return stack[:len(stack)-1], nil
	},
	"/": func(stack []int) ([]int, error) {
		if len(stack) < 2 {
			return stack, errors.New(notEnoughStack)
		} else if stack[len(stack)-1] == 0 {
			return stack, errors.New("divided by zero")
		}
		stack[len(stack)-2] = stack[len(stack)-2] / stack[len(stack)-1]
		return stack[:len(stack)-1], nil
	},
	"dup": func(stack []int) ([]int, error) {
		if len(stack) < 1 {
			return stack, errors.New(notEnoughStack)
		}
		return append(stack, stack[len(stack)-1]), nil
	},
	"drop": func(stack []int) ([]int, error) {
		if len(stack) < 1 {
			return stack, errors.New(notEnoughStack)
		}
		return stack[0 : len(stack)-1], nil
	},
	"swap": func(stack []int) ([]int, error) {
		if len(stack) < 2 {
			return stack, errors.New(notEnoughStack)
		}
		stack[len(stack)-2], stack[len(stack)-1] = stack[len(stack)-1], stack[len(stack)-2]
		return stack, nil
	},
	"over": func(stack []int) ([]int, error) {
		if len(stack) < 2 {
			return stack, errors.New(notEnoughStack)
		}
		return append(stack, stack[len(stack)-2]), nil
	},
}
// Forth evaluates the given Forth inputs and returns the result.
func Forth(inputs []string) ([]int, error) {
	stack := make([]int, 0)
	userDefinedWords := map[string][]string{}
	for _, input := range inputs {
		if strings.HasPrefix(input, ":") && strings.HasSuffix(input, ";") {
			if err := parseUserDefinedWord(input, userDefinedWords); err != nil {
				return nil, err
			}
		} else {
			return runOps(stack, strings.Fields(input), userDefinedWords)
		}
	}
	return nil, nil // No op
}
func parseUserDefinedWord(input string, userDefinedWords map[string][]string) error {
	elements := strings.Fields(input)
	if _, err := strconv.Atoi(elements[1]); err == nil { // [0]: ":"
		return errors.New("cannot redefine numbers")
	}
	ops := make([]string, 0)
	// Expand the user defined words flat to avoid from getting modified by subsequent overrides
	for _, op := range elements[2 : len(elements)-1] { // [0]: ":", [1]: word Name, [len(element)-1]: ";"
		if wordOps, ok := userDefinedWords[strings.ToLower(op)]; ok {
			ops = append(ops, wordOps...)
		} else {
			ops = append(ops, op)
		}
	}
	userDefinedWords[strings.ToLower(elements[1])] = ops
	return nil
}
func runOps(stack []int, ops []string, userDefinedWords map[string][]string) ([]int, error) {
	for _, op := range ops {
		if n, err := strconv.Atoi(op); err == nil {
			stack = append(stack, n)
		} else if wordOps, ok := userDefinedWords[strings.ToLower(op)]; ok {
			if stack, err = runOps(stack, wordOps, userDefinedWords); err != nil {
				return stack, err
			}
		} else if builtinOp, ok := builtinOps[strings.ToLower(op)]; ok {
			if stack, err = builtinOp(stack); err != nil {
				return stack, err
			}
		} else {
			return stack, errors.New("Unrecognized operation")
		}
	}
	return stack, nil
}
