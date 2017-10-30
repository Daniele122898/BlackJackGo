package main

import (
	"fmt"
	//"io/ioutil"
	"math/rand"
	//"os"
	//"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type card struct {
	name string
	value int
}
type deck []card

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for i, value := range cardValues {
			if i > 9 {
				i = 9
			}
			aC := card {name: value + " of " +suit, value: i+1}
			cards = append(cards, aC)
		}
	}

	return cards
}

func (d deck) print() {
	fmt.Println(d.toString())
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	dS := ""
	for i, card := range d{
		//dS += string(i) + ". "+card.name +" : "+ string(card.value) +" \n"
		dS+= fmt.Sprintf("%v. %s : %v \n", i,card.name, card.value)
	}
	return dS
}

//func (d deck) saveToFile(filename string) error {
//	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
//}
//
//func newDeckFromFile(filename string) deck {
//	bs, err := ioutil.ReadFile(filename)
//	if err != nil {
//		// Option #1 - log the error and return a call to newDeck()
//		// Option #2 - Log the error and entirely quit the program
//		fmt.Println("Error:", err)
//		os.Exit(1)
//	}
//
//	s := strings.Split(string(bs), ",")
//	return deck(s)
//}

func (d deck) names() string{
	s := ""
	for _, card := range d{
		s += fmt.Sprintf("%s, ", card.name)
	}
	return s
}

func (d deck) value() int{
	v := 0
	ace := false
	for _, card := range d{
		v += card.value
		if card.value == 1{
			ace = true
		}
	}
	if ace && v <= 11{
		v += 10
	}
	return v
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}