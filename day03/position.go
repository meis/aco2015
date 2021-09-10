package main

type position struct {
	x int
	y int
}

func (p position) moveUp() position {
	return position{p.x + 1, p.y}
}

func (p position) moveDown() position {
	return position{p.x - 1, p.y}
}

func (p position) moveRight() position {
	return position{p.x, p.y + 1}
}

func (p position) moveLeft() position {
	return position{p.x, p.y - 1}
}
