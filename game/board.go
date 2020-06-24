package game

import (
	"math/rand"
	"time"
)

// Board interface
type Board interface {
}

type board struct {
	cells  cells
	height int
	width  int
}

type cells map[int]map[int]cell

func newBoard(height, width int) *board {
	b := &board{
		height: height,
		width:  width,
	}
	b.cells = b.generateCells(height, width)
	b.cells = b.generateBees(b.cells, 10)

	return b
}

func (b *board) generateCells(height, width int) cells {
	cells := make(map[int]map[int]cell)
	for x := 0; x < width; x++ {
		cells[x] = make(map[int]cell)

		for y := 0; y < height; y++ {
			cell := newCell(false, x, y)
			cells[x][y] = *cell
		}
	}

	return cells
}

func (b *board) generateBees(cells cells, count int) cells {

	for count > 0 {
		x := random(0, b.width)
		y := random(0, b.height)

		for i := 0; i < x; i++ {
			for j := 0; j < y; j++ {
				cell := newCell(false, x, y)
				cells[x][y] = *cell
			}
		}

		if _, found := b.cells[x][y]; found == true {
			if b.cells[x][y].bee == false {
				cell := b.cells[x][y]
				cell.bee = true
				b.cells[x][y] = cell
				count--
			}
		}
	}

	return b.cells
}

func random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
