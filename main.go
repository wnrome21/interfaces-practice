package main

import "fmt"

type englishBot struct {}

type spanishBot struct {}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	//Very custon logic for agenerating an english greeting
	return "Hello There!"
}

func (spanishBot) getGreeting() string {
	// Custom logic for generating an english greeting
	return "Qué Tál?"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting( sb spanishBot) {
	fmt.Println(sb.getGreeting())
}