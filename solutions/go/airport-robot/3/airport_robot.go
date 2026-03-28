package airportrobot

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

func (portuguese Portuguese) LanguageName() string {
	return "Portuguese"
}

func (italian Italian) Greet(visitor string) string {
	return fmt.Sprintf("I can speak %s: Ciao %s!", italian.LanguageName(), visitor)
}

func (portuguese Portuguese) Greet(visitor string) string {
	return fmt.Sprintf("I can speak %s: Olá %s!", portuguese.LanguageName(), visitor)
}

func SayHello(visitor string, greeter Greeter) string {
	return greeter.Greet(visitor)
}
