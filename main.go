package main

import (
	"log"

	"github.com/MihaiBlebea/beesweeper/game"
)

// BeeSweeper _
type BeeSweeper struct {
	Game   *game.Game
	Screen *Screen
}

func main() {
	screen := &Screen{}

	gm := game.NewGame()

	bs := BeeSweeper{gm, screen}
	_ = bs
	err := screen.render()
	if err != nil {
		log.Panic(err)
	}

	// demo()
}
