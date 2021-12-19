package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
	number int
}
type spanishBot struct{}

func main() {
	eb := englishBot{number: 9799}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (en englishBot) getGreeting() string {
	msg := fmt.Sprintf("Hi There! %d", en.number)
	return msg
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
