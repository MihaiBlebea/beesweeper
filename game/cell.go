package game

// Cell interface
type Cell interface {
	HasBee() bool
	IsSelected() bool
}

type cell struct {
	bee        bool
	x          int
	y          int
	discovered bool
	count      int
	selected   bool
}

func newCell(bee bool, x, y int) *cell {
	return &cell{bee: bee, x: x, y: y}
}

// HasBee returns true if a bee is present on this cell
func (c *cell) HasBee() bool {
	return c.bee
}

// IsSelected returns true if the cell is selected by the user
func (c *cell) IsSelected() bool {
	return c.selected
}

func (c *cell) assignBee(hasBee bool) {
	c.bee = hasBee
}

func (c *cell) assignCount(count int) {
	c.count = count
}
