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
	UncoverCell(x, y int) (gameOver bool)
	ToggleFlag(x, y int)
	Won() bool
}

type board struct {
	cells      cells
	cellCountH int
	cellCountW int
	beeCount   int
}

type cells map[int]map[int]cell

func newBoard(cellCountH, cellCountW, beeCount int) *board {
	b := &board{
		cellCountH: cellCountH,
		cellCountW: cellCountW,
		beeCount:   beeCount,
	}
	b.cells = b.generateCells(cellCountH, cellCountW)
	b.cells = b.generateBees(b.cells, b.beeCount)

	b.countBees()

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

func (b *board) countBees() {
	w := b.GetCellCountW()
	h := b.GetCellCountH()

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			total := 0

			if cell := b.GetCell(x, y); cell.HasBee() {
				total++
			}

			// TODO: This needs to be refactored
			if x > 0 && y > 0 && b.cells[x-1][y-1].bee == true {
				total++
			}
			if x > 0 && b.cells[x-1][y].bee == true {
				total++
			}
			if x > 0 && y < h-1 && b.cells[x-1][y+1].bee == true {
				total++
			}

			if y > 0 && b.cells[x][y-1].bee == true {
				total++
			}
			if y < h-1 && b.cells[x][y+1].bee == true {
				total++
			}

			if x < w-1 && y > 0 && b.cells[x+1][y-1].bee == true {
				total++
			}
			if x < w-1 && b.cells[x+1][y].bee == true {
				total++
			}
			if x < w-1 && y < h-1 && b.cells[x+1][y+1].bee == true {
				total++
			}
			cell := b.cells[x][y]
			cell.count = total
			b.cells[x][y] = cell
		}
	}
}

func (b *board) UncoverCell(x, y int) (gameOver bool) {
	cell := b.cells[x][y]
	if cell.IsDiscovered() == true {
		return false
	}

	if cell.HasBee() == true {
		return true
	}

	cell.count = 0
	cell.showCount = false
	cell.discovered = true
	b.cells[x][y] = cell

	w := b.GetCellCountW()
	h := b.GetCellCountH()

	if x > 0 && y > 0 {
		if b.cells[x-1][y-1].count == 0 {
			b.UncoverCell(x-1, y-1)
		} else {
			b.showCountInCell(x-1, y-1)
		}
	}

	if x > 0 {
		if b.cells[x-1][y].count == 0 {
			b.UncoverCell(x-1, y)
		} else {
			b.showCountInCell(x-1, y)
		}
	}

	if x > 0 && y < h-1 {
		if b.cells[x-1][y+1].count == 0 {
			b.UncoverCell(x-1, y+1)
		} else {
			b.showCountInCell(x-1, y+1)
		}
	}

	if y > 0 {
		if b.cells[x][y-1].count == 0 {
			b.UncoverCell(x, y-1)
		} else {
			b.showCountInCell(x, y-1)
		}
	}

	if y < h-1 {
		if b.cells[x][y+1].count == 0 {
			b.UncoverCell(x, y+1)
		} else {
			b.showCountInCell(x, y+1)
		}
	}

	if x < w-1 && y > 0 {
		if b.cells[x+1][y-1].count == 0 {
			b.UncoverCell(x+1, y-1)
		} else {
			b.showCountInCell(x+1, y-1)
		}
	}

	if x < w-1 {
		if b.cells[x+1][y].count == 0 {
			b.UncoverCell(x+1, y)
		} else {
			b.showCountInCell(x+1, y)
		}
	}

	if x < w-1 && y < h-1 {
		if b.cells[x+1][y+1].count == 0 {
			b.UncoverCell(x+1, y+1)
		} else {
			b.showCountInCell(x+1, y+1)
		}
	}

	return false
}

func (b *board) ToggleFlag(x, y int) {
	cell := b.cells[x][y]
	cell.flagged = !cell.flagged
	b.cells[x][y] = cell
}

func (b *board) Won() bool {
	w := b.GetCellCountW()
	h := b.GetCellCountH()

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if cell := b.GetCell(x, y); cell.HasBee() && cell.IsFlagged() == false {
				return false
			}
		}
	}

	return true
}

func (b *board) showCountInCell(x, y int) {
	cell := b.cells[x][y]
	cell.showCount = true
	b.cells[x][y] = cell
}

func (b *board) unselect(x, y int) Cell {
	cell := b.cells[x][y]
	cell.selected = false

	b.cells[x][y] = cell

	return &cell
}

func random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
