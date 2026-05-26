package main

func main() {

	// =================================== Section 1
	// cards := newDeck()

	// // hand, remainingCards := deal(cards, 5)

	// // hand.print()
	// // remainingCards.print()

	// cards.saveToFile("my_cards.txt")

	// =================================== Section 2

	// cards := newDeckFromFile("my_cards.txt")
	// cards.print()

	// =================================== Section 3

	cards := newDeck()
	// cards.shuffle()
	cards.custom_shuffle()
	cards.print()

}
