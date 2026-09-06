package main

import "math/rand/v2"

type card struct {
	name   string
	value  int
	effect func()
}

func addCards(deck []card, name string, value int, effect func(), n int) []card {
	for i := 0; i < n; i++ {
		deck = append(deck, card{name, value, effect})
	}
	return deck
}

func CreateShuffledDeck() []card {
	var deck []card

	deck = addCards(deck, "Assassino", -1, func() {}, 2)
	deck = addCards(deck, "Bobo da Corte", 0, func() {}, 2)
	deck = addCards(deck, "Guarda", 1, func() {}, 8)
	deck = addCards(deck, "Espião", 2, func() {}, 3)
	deck = addCards(deck, "Barão", 3, func() {}, 3)
	deck = addCards(deck, "Dama", 4, func() {}, 3)
	deck = addCards(deck, "Príncipe", 5, func() {}, 2)
	deck = addCards(deck, "Rei", 6, func() {}, 1)
	deck = addCards(deck, "Condessa", 7, func() {}, 1)
	deck = addCards(deck, "Princesa", 8, func() {}, 1)

	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	return deck
}
