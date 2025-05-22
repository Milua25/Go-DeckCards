package main

func main() {

	cards := newDeckFromFile("myCards")
	cards.shuffle()
	cards.print()
}
