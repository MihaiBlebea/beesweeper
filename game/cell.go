package game

type cell struct {
	bee        bool
	x          int
	y          int
	discovered bool
	count      int
}

func newCell(bee bool, x, y int) *cell {
	return &cell{bee: bee, x: x, y: y}
}

func (c *cell) assignBee(hasBee bool) {
	c.bee = hasBee
}

func (c *cell) assignCount(count int) {
	c.count = count
}
