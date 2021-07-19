package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	decks := deck{}
	cardSuite := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValue := []string{"One", "Two", "Three"}
	for _, suite := range cardSuite {
		for _, value := range cardValue {
			decks = append(decks, value+" "+suite)
		}
	}
	return decks
}

func (d deck) print() {
	for index, value := range d {
		fmt.Println(fmt.Sprintf("Index :%d, value :%s", index, value))
	}
}

func (d deck) deal(splitNum int) (deck, deck) {
	return d[:splitNum], d[splitNum:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func toBytesArray(value string) []byte {
	return []byte(value)
}

func writeToFile(value []byte) {
	// File Name
	// arrayBytes
	// Permission
	// returns if any error
	err := ioutil.WriteFile("hello", value, 0644)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func readFromFile() string {
	// returns value and error
	value, err := ioutil.ReadFile("hello")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	return string(value)
}

func stringToDeck(value string) deck {
	return deck(strings.Split(string(value), ","))
}

func (d deck) shuffel() deck {
	length := len(d)
	source :=
		rand.NewSource(time.Now().UnixNano())
	randType := rand.New(source)

	// result =: rand.Shuffle(length, func(i, j int) {
	// 	temp := d[j]
	// 	d[j] = d[i]
	// 	d[i] = temp
	// })
	for index, value := range d {
		randVal := randType.Intn(length - 1)
		temp := d[randVal]
		d[randVal] = value
		d[index] = temp
	}

	return d
}
