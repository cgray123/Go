package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Value int
	Name  string
	Suit  string
}
type Player struct {
	Name     string
	Cards    []Card
	HasAce   map[bool]int
	IsBusted bool
}
//generates a deck of cards in order and return them in a slice of Card
func GenerateDeck() []Card {
	var deck []Card
	var suit string
	for x := 0; x < 4; x++ {
		if x == 0 {
			suit = "Spade"
		} else if x == 1 {
			suit = "Clubs"
		} else if x == 2 {
			suit = "Diamonds"
		} else if x == 3 {
			suit = "Hearts"
		}
		for i := 2; i < 15; i++ {
			if i <= 10 {
				deck = append(deck, Card{Name: fmt.Sprint(i), Value: i, Suit: suit})
			} else if i == 11 {
				deck = append(deck, Card{Name: "jack", Value: 10, Suit: suit})
			} else if i == 12 {
				deck = append(deck, Card{Name: "queen", Value: 10, Suit: suit})
			} else if i == 13 {
				deck = append(deck, Card{Name: "king", Value: 10, Suit: suit})
			} else {
				deck = append(deck, Card{Name: "ace", Value: 11, Suit: suit})
			}

		}
	}
	return deck
}
//Shuffles slices of Card
func ShuffleDeck(deck []Card) []Card {
	rand.Seed(time.Now().UnixNano())
	for i := len(deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
	return deck
}
//outputs the names and hands of everyone
func PrintTable(p []Player) {
	for i := 0; i < len(p); i++ {
		if p[i].Name == "Dealer" {
			fmt.Println(p[i].Name, "Hand:", p[i].Cards[0], "???")
		} else {
			fmt.Println(p[i].Name, "Hand:", p[i].Cards)
		}

	}
}
//takes in input from user must be 1-4
func inPlay() (int, error) {
	fmt.Println("Input 1-4 for amount of players")
	var pNum int
	fmt.Scanln(&pNum)
	if pNum > 4 || pNum < 1 {
		return 0, errors.New("input an int between 1 and 4")
	}
	return pNum, nil
}
//main calling method, sets up names and structs
func BlackJack() {
retry:
	pNum, err := inPlay()
	var players []Player
	if err != nil {
		fmt.Println(err)
		goto retry
	} else {
		for i := 1; i <= pNum; i++ {
			var name string

			fmt.Println("input player", i, "name")
			fmt.Scanln(&name)
			players = append(players, Player{Name: name, HasAce: make(map[bool]int)})
		}
		players = append(players, Player{Name: "Dealer", HasAce: make(map[bool]int)})
		Play(players)
	}

}
//deals out the first cards and runs the loop that handles all the hits
func Play(players []Player) {
	deck := ShuffleDeck(GenerateDeck())
	for x := 0; x < len(players); x++ {
		players[x].Cards = nil
		players[x].Cards = append(players[x].Cards, deck[x])
		players[x].Cards = append(players[x].Cards, deck[x+len(players)])
	}

	deck = deck[len(players)*2:]
	PrintTable(players)
	for i := 0; i < len(players); i++ {
	draw:
		err := Draw(players[i], deck)
		if err != nil {
			fmt.Println(err)
			goto draw
		}
	}
	End(players)
}
//output who wins, losses, and pushed/ then ask if they want to play again
func End(p []Player) {
	for i := 0; i < len(p)-1; i++ {
		if !p[i].IsBusted && Tot(p[i]) > Tot(p[len(p)-1]) {
			fmt.Println(p[i].Name, "has won")
		} else if !p[i].IsBusted && Tot(p[i]) == Tot(p[len(p)-1]) {
			fmt.Println(p[i].Name, "has pushed")
		} else {
			fmt.Println(p[i].Name, "has lost")
		}

	}
	fmt.Println("input y to play again")
	var input string
	fmt.Scanln(&input)
	if input == "y" {
		Play(p)
	}
}
//handles logic for starting a hit and staying 
func Draw(p Player, deck []Card) error {
	if p.Name == "Dealer" {
		if Tot(p) < 18 {
			Hit(p, deck)
		}
	} else {
		fmt.Println(p.Name, "Your total is", Tot(p), "would you like to 'hit' or 'stay'")
		var input string
		fmt.Scanln(&input)
		if input == "hit" {
			Hit(p, deck)

		} else if input == "stay" {
			fmt.Println("your Total is", Tot(p))
		} else {
			return errors.New("please input hit or stay")

		}
	}

	return nil
}
//return to card value
func Tot(p Player) int {
	var tot int
	for i := 0; i < len(p.Cards); i++ {
		tot += p.Cards[i].Value
		if p.Cards[i].Name == "ace" {
			p.HasAce[true] = i

		}

	}
	return tot
}
//give player a new card and handles busting
func Hit(p Player, deck []Card) {
	p.Cards = append(p.Cards, deck[0])
	fmt.Println(p.Name, "drew a", deck[0])
	deck = deck[0:]
	if Tot(p) > 21 {
		if v, ok := p.HasAce[true]; ok {
			p.Cards[v].Value = 1
			fmt.Println(p.Cards)
			Draw(p, deck)
		} else {
			fmt.Println(p.Name, "have busted")
			p.IsBusted = true
		}
	} else {
		Draw(p, deck)
	}
}
func main() {
	BlackJack()
}
