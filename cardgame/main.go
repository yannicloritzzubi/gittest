package main

func main() {
	/*cards := newDeck()
		hand, remainingCards := deal(cards, 4)

		hand.print()
		remainingCards.print()
	ss
		cards.saveToFile("testFile.txt")
		fmt.Println(cards.toString())*/

	cards := newDeckFromFile("testFile.txt")
	cards.shuffle()
	cards.print()

}

func returnString() string {
	return "test"
}
xadsf