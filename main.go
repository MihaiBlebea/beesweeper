package main

import (
	"log"

	"github.com/MihaiBlebea/beesweeper/game"
)

func main() {
	gm := game.NewGame()

	screen := &Screen{20, 20, 40, 40, 2, gm}

	err := screen.render()
	if err != nil {
		log.Panic(err)
	}

	// demo()
}
