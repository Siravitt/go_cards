package main

func main() {
	// cards := newDeckFromFile("my_Cards")
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
