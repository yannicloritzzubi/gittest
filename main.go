package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishbot struct{}
type spanishbot struct{}
type schweizerbot struct{}

func main() {

	eb := englishbot{}
	sb := spanishbot{}
	chb := schweizerbot{}

	printGreeting(eb)
	printGreeting(chb)
	fmt.Println(sb.getGreeting())

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishbot) getGreeting() string {
	// Very custom logic for generating an English greeting
	return "Hi There!"
}

func (sb spanishbot) getGreeting() string {
	return "Hola!"
}
