package alphabetBoardPath

import (
	"strings"
)

func alphabetBoardPath(target string) string {
	paths := make([]string, 0)
	currentPosition := NewPosition(0, 0)
	for _, char := range target {
		targetPosition := NewPositionFromChar(char)
		paths = append(paths, currentPosition.MoveTo(targetPosition))
	}

	return strings.Join(paths, "")
}

func NewPositionFromChar(char rune) *Position {
	offset := int(char - 'a')
	x := offset % 5
	y := offset / 5
	return NewPosition(x, y)
}

func NewPosition(x int, y int) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}

type Position struct {
	X int
	Y int
}

func (p *Position) MoveTo(target *Position) string {
	path := ""
	for p.X != target.X || p.Y != target.Y {
		offsetX := target.X - p.X
		offsetY := target.Y - p.Y
		if offsetY > 0 && !p.isOnYDownBoundary() {
			path += "D"
			p.Y++
			continue
		} else if offsetY < 0 {
			path += "U"
			p.Y--
			continue
		}

		if offsetX > 0 {
			path += "R"
			p.X++
			continue
		} else if offsetX < 0 {
			path += "L"
			p.X--
			continue
		}
	}

	path += "!"
	return path
}

func (p *Position) isOnYDownBoundary() bool {
	if p.X == 0 && p.Y == 5 {
		return true
	} else if p.X != 0 && p.Y == 4 {
		return true
	}
	return false
}
