package main

import (
	"log"

	"github.com/MihaiBlebea/beesweeper/game"
)

func main() {
	gm := game.NewGame(20, 20)

	screen := NewScreen(20, 20, 2, gm)

	err := screen.render()
	if err != nil {
		log.Panic(err)
	}

	// demo()
}
