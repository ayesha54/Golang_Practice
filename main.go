package main

func main() {
	// cards := newDeck()

	//cards.print()

	// hand, remainingCards := deal(cards, 3)

	// hand.print()
	// fmt.Println("\n")
	// remainingCards.print()

	//fmt.Println(cards.toString())
	// cards.saveToFile("mycards")
	cards := newDeckFromFile("mycards")
	cards.shuffle()
	cards.print()

}
