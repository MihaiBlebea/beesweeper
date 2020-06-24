package game

import (
	"math/rand"
	"time"
)

// Game model
type Game struct {
	id      string
	playing bool
	turn    int
	board   board
}

// NewGame creates a new game struct
func NewGame() *Game {
	return &Game{
		id:      generateID(),
		playing: true,
		turn:    1,
		board:   *newBoard(10, 10),
	}
}

// GetBoard returns the game board
func (g *Game) GetBoard() Board {
	return g.board
}

func generateID() string {
	rand.Seed(time.Now().UnixNano())

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
