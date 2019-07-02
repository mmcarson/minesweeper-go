package board

import "strconv"

type Square struct {
	NeighboringMines int
	IsMine bool
	IsSeen bool
}

func NewSquare () *Square {
	s := new(Square)
	s.IsSeen = false
	s.IsMine = false
	s.NeighboringMines = 0
	return s
}

func (s *Square) Print () string {
	if !s.IsSeen {
		return "*"
	} else if s.IsMine {
		return "M"
	} else {
		return strconv.Itoa(s.NeighboringMines)
	}
}