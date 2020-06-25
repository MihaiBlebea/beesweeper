package game

// Cell interface
type Cell interface {
	HasBee() bool
	IsSelected() bool
	GetCount() int
	IsDiscovered() bool
	ShouldShowCount() bool
	IsFlagged() bool
}

type cell struct {
	bee        bool
	x          int
	y          int
	discovered bool
	count      int
	selected   bool
	showCount  bool
	flagged    bool
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

// IsDiscovered returns true if the cell is discovered by the user
func (c *cell) IsDiscovered() bool {
	return c.discovered
}

// IsFlagged returns true if the cell is flagged
func (c *cell) IsFlagged() bool {
	return c.flagged
}

// GetCount returns the count for the cell
func (c *cell) GetCount() int {
	return c.count
}

func (c *cell) ShouldShowCount() bool {
	return c.showCount
}

func (c *cell) assignBee(hasBee bool) {
	c.bee = hasBee
}

func (c *cell) assignCount(count int) {
	c.count = count
}
