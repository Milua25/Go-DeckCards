package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// method print Function on type deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// newDeck Function
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// deal Function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// convert the type deck to string
func (d deck) toString() string {
	return strings.Join(d, ",")
}

// saveToFile Method function
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// newDeckFromFile Function
func newDeckFromFile(filename string) deck {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("unable to read the file: %s", filename)
		os.Exit(1)
	}
	// type conversion of slice of byte to string
	deckString := string(data)
	deckSlice := deck(strings.Split(deckString, ","))
	return deckSlice
}

// shuffle Method Function
func (d deck) shuffle() {
	rand.New(rand.NewSource(time.Now().Unix()))
	deckLength := len(d) - 1
	for i, _ := range d {
		randNum := rand.Int() % deckLength
		d[i], d[randNum] = d[randNum], d[i]
	}
}
