package game

import (
	"math/rand"
	"time"
)

// Board interface
type Board interface {
	GetCellCountH() int
	GetCellCountW() int
	GetCell(x, y int) Cell
	SetSelected(x, y int) Cell
	UnselectAll()
}

type board struct {
	cells      cells
	cellCountH int
	cellCountW int
}

type cells map[int]map[int]cell

func newBoard(cellCountH, cellCountW int) *board {
	b := &board{
		cellCountH: cellCountH,
		cellCountW: cellCountW,
	}
	b.cells = b.generateCells(cellCountH, cellCountW)
	b.cells = b.generateBees(b.cells, 10)

	return b
}

func (b *board) GetCellCountH() int {
	return b.cellCountH
}

func (b *board) GetCellCountW() int {
	return b.cellCountW
}

func (b *board) GetCell(x, y int) Cell {
	cell := b.cells[x][y]

	return &cell
}

func (b *board) SetSelected(x, y int) Cell {
	cell := b.cells[x][y]
	cell.selected = true

	b.cells[x][y] = cell

	return &cell
}

func (b *board) unselect(x, y int) Cell {
	cell := b.cells[x][y]
	cell.selected = false

	b.cells[x][y] = cell

	return &cell
}

func (b *board) UnselectAll() {
	for x := 0; x < b.cellCountW; x++ {
		for y := 0; y < b.cellCountH; y++ {
			b.unselect(x, y)
		}
	}
}

func (b *board) generateCells(cellCountH, cellCountW int) cells {
	cells := make(map[int]map[int]cell)
	for x := 0; x < cellCountW; x++ {
		cells[x] = make(map[int]cell)

		for y := 0; y < cellCountH; y++ {
			cell := newCell(false, x, y)
			cells[x][y] = *cell
		}
	}

	return cells
}

func (b *board) generateBees(cells cells, count int) cells {

	for count > 0 {
		x := random(0, b.cellCountW)
		y := random(0, b.cellCountH)

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
