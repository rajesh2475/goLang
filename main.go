package main

import "fmt"

func main() {
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, value := range values {
		if value%2 == 0 {
			fmt.Println(fmt.Sprintf("The value : %d is even", value))
		} else {
			fmt.Println(fmt.Sprintf("The value : %d is odd", value))
		}
	}
}

// func main() {
// 	cards := newDeck()
// 	// cards.print()
// 	// cardsHand, remainingCards := cards.deal(5)
// 	// cardsHand.print()
// 	// remainingCards.print()

// 	bytesData := toBytesArray(cards.toString())
// 	// fmt.Println(string([]byte(strings.Join(cards, ","))))
// 	writeToFile(bytesData)
// 	stringData := readFromFile()
// 	finalResult := stringToDeck(stringData)
// 	finalResult.print()
// 	// fmt.Println("1111111111111111111111111111111111111111111111")
// 	finalResult.shuffel().print()

// }
