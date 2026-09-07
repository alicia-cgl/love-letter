package main

import (
	"fmt"
	"slices"
)

func playerRound(player *player) {
	var chosenCard int

	fmt.Printf("%s, escolha qual carta jogar:\n", player.name)
	for i := 0; i < len(player.cards); i++ {
		fmt.Printf("\t%d. %s [%d]\n", i+1, player.cards[i].name, player.cards[i].power)
	}

	_, err := fmt.Scanln(&chosenCard)
	if err != nil {
		fmt.Println("Erro lendo entrada do usuário: ", err)
	}

	card := player.cards[chosenCard-1]
	player.cards = slices.Delete(player.cards, chosenCard-1, chosenCard)
	card.effect()

}

func main() {
	deck := CreateShuffledDeck()

	numOfPlayers := 2
	playerQueue := CreatePlayerQueue(numOfPlayers)

	round := 0
	for len(playerQueue) > 0 || len(deck) > 0 {
		p := &playerQueue[round%numOfPlayers]
		if round < len(playerQueue) {
			p.cards = append(p.cards, deck[0])
			deck = deck[1:]
		}

		p.cards = append(p.cards, deck[0])
		deck = deck[1:]
		playerRound(p)

		round++
	}

}
