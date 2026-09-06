package main

import "fmt"

type player struct {
	card        string
	isProtected bool
}

func main() {
	deck := CreateShuffledDeck()

	for i := 0; i < len(deck); i++ {
		fmt.Println(deck[i])
	}

}
