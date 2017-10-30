package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Started new Game\n")
	cards := newDeck()
	//cards.print()
	cards.shuffle()
	h, d := deal(cards, 2)
	if h.value() > 21{
		main()
	}
	gameLoop(d, h)

}

func gameLoop(d deck, h deck){
	hv:= h.value()
	hn := h.names()
	//Value of cards is above 21
	if hv > 21 {
		fmt.Println("YOU LOST! REASON: BUST!\n")
		main()
	}
	fmt.Println("Your Cards: ", hn)
	fmt.Println("Value: ", hv)
	fmt.Println("1. Hit 2. Stay")
	suc := false
	reader := bufio.NewReader(os.Stdin)
	for suc == false{
		s, _ := reader.ReadString('\n')
		s = strings.Trim(s, "\n\r")
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Error: ",err)
			fmt.Println("Please enter an integer")
			continue
		}
		switch i {
		case 1:
			hit(d, h)
			suc = true
		case 2:
			stay(d,h)
			suc= true
		default:
			fmt.Println("Please choose a correct option!")
			continue
		}
	}
}

func hit(d deck, h deck){
	nh, nd := deal(d, 1)
	h = append(h, nh[0])
	d = nd
	gameLoop(d, h)
}

func stay(d deck, h deck){
	//Dealer should do shit
	dh, nd := deal(d, 2)
	//make dealer decisions
	for {
		n := dh.names()
		v := dh.value()
		fmt.Println("Dealer Cards: ", n)
		fmt.Println("Dealer Value: ", v)
		if dh.value() > 21 {
			fmt.Println("YOU WON!\n")
			break
		}
		if dh.value() > h.value(){
			fmt.Println("YOU LOST!\n")
			break
		}
		if dh.value() == h.value() && dh.value() < 16 {
			//hit
			nd, dh = dealerHit(nd, dh)
			continue
		}
		if dh.value() == h.value() && dh.value() >= 16{
			fmt.Println("DRAW!\n")
			break
		}
		if dh.value() < h.value(){
			//hit
			nd, dh = dealerHit(nd, dh)
			continue
		}
	}
	main()
}

func dealerHit(d deck, h deck) (deck, deck){
	dh, nd := deal(d, 1)
	h = append(h, dh[0])
	return nd, h
}