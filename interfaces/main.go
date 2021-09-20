package main

import "fmt"

type bot interface{
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	ebot := englishBot{}
	sbot := spanishBot{}

	printGreeting(ebot)
	printGreeting(sbot)

}
//usage of inetrfaces
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	//VERY custom logic for generating an English logic
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	//VERY custom logic for generating a Spanish logic
	return "Hola!"
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb.spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

