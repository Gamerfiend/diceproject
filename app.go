package main

import "diceproject/game"

func main() {
	playGame()
}

func playGame() {
	game := &game.Game{}
	game.NewGame()
	game.Play()
}
