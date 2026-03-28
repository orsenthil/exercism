package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

import "fmt"

type Greeter interface {
    LanguageName() string
    Greet(visitor string) string
}

type Italian struct {
}

type Portuguese struct {
}

func (italian Italian) LanguageName() string {
    return "Italian"
}

func (italian Italian) Greet(visitor string) string {
	language := italian.LanguageName()
    return fmt.Sprintf("I can speak %s: Ciao %s!", language, visitor)
}

func (portuguese Portuguese) LanguageName() string {
    return "Portuguese"
}

func (portuguese Portuguese) Greet(visitor string) string {
	language := portuguese.LanguageName()
    return fmt.Sprintf("I can speak %s: Olá %s!", language, visitor)

}
func SayHello(visitor string, greeter Greeter) string {
    return greeter.Greet(visitor)
}