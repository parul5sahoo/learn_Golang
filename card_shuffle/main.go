package main


//import "fmt"

func main() {
	//var card string = "Ace of Spades"
	cards := newDeck()
	cards.shuffle()

	//hand, remainingDeck := deal(cards, 5)

	// // hand.print()
	// // remainingDeck.print()

	// greeting := "Hi there!"

	// fmt.Println([]byte(greeting))

	//var card string = "Ace of Spades"
	//cards := newDeck()
	//cards.saveToFile("Gocards")

	// cards := newDeckFromFile("Gocards")
	cards.print()
}

