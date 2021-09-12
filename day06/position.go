package main

type position struct {
	x int
	y int
}

func getPositionRange(startX int, startY int, endX int, endY int) []position {
	r := []position{}

	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			r = append(r, position{x, y})
		}
	}

	return r
}
