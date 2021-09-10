package main

type santa struct {
	Position position
}

func (s *santa) move(direction string) position {
	switch direction {
	case "^":
		s.Position = s.Position.moveUp()
	case "v":
		s.Position = s.Position.moveDown()
	case ">":
		s.Position = s.Position.moveRight()
	case "<":
		s.Position = s.Position.moveLeft()
	default:
		panic("not a direction")
	}
	return s.Position
}
