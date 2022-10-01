package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// Which is a slice of strings

type deck []string

func newDeck() deck {
	cardSuits := []string{"Club", "Diamond", "Heart", "Spade"}
	cardValues := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	cards := deck{}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+"-"+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

func (cards deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range cards {
		ranI := r.Intn(len(cards) - 1)
		cards[i], cards[ranI] = cards[ranI], cards[i]
	}
}

func (cards deck) print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func (cards deck) toString() string {
	return strings.Join([]string(cards), ",")
}

func (cards deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(cards.toString()), 0666)
}
