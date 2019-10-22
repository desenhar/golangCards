package main

func main() {
	//cards := newDeckFromFile("cards5")
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// fmt.Println(cards.toString())
	//cards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

