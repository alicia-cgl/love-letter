package main

import "fmt"

type player struct {
	name        string
	cards       []card
	isProtected bool
}

func CreatePlayerQueue(numOfPlayers int) []player {
	queue := make([]player, numOfPlayers)

	for i := 0; i < numOfPlayers; i++ {
		queue[i] = player{name: fmt.Sprintf("Player %d", i+1), isProtected: false}
	}

	return queue
}
