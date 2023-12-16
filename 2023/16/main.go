package main

import (
	"fmt"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Beam struct {
	direction  Direction
	coordinate aoc.Coordinate
}

const (
	EmptySpace = iota
	MirrorForward
	MirrorBackward
	SplitterHorizontal
	SplitterVertical
)

func beams(grid *aoc.Grid, initial Beam) []Beam {
	maxX, maxY := grid.MaxX(), grid.MaxY()

	beams := make(map[Beam]struct{})

	current := []Beam{initial}
	for len(current) > 0 {
		next := []Beam{}
		for _, beam := range current {
			if _, ok := beams[beam]; ok {
				continue
			}

			x, y := beam.coordinate.X, beam.coordinate.Y
			if x < 0 || x > maxX || y < 0 || y > maxY {
				continue
			}

			beams[beam] = struct{}{}
			tile := grid.Get(beam.coordinate).Value()

			switch beam.direction {
			case Up:
				switch tile {
				case MirrorForward:
					next = append(next, Beam{Right, aoc.NewCoordinate(x+1, y)})
				case MirrorBackward:
					next = append(next, Beam{Left, aoc.NewCoordinate(x-1, y)})
				case SplitterHorizontal:
					next = append(next, Beam{Left, aoc.NewCoordinate(x-1, y)}, Beam{Right, aoc.NewCoordinate(x+1, y)})
				default:
					next = append(next, Beam{Up, aoc.NewCoordinate(x, y-1)})
				}
			case Down:
				switch tile {
				case MirrorForward:
					next = append(next, Beam{Left, aoc.NewCoordinate(x-1, y)})
				case MirrorBackward:
					next = append(next, Beam{Right, aoc.NewCoordinate(x+1, y)})
				case SplitterHorizontal:
					next = append(next, Beam{Left, aoc.NewCoordinate(x-1, y)}, Beam{Right, aoc.NewCoordinate(x+1, y)})
				default:
					next = append(next, Beam{Down, aoc.NewCoordinate(x, y+1)})
				}
			case Left:
				switch tile {
				case MirrorForward:
					next = append(next, Beam{Down, aoc.NewCoordinate(x, y+1)})
				case MirrorBackward:
					next = append(next, Beam{Up, aoc.NewCoordinate(x, y-1)})
				case SplitterVertical:
					next = append(next, Beam{Down, aoc.NewCoordinate(x, y+1)}, Beam{Up, aoc.NewCoordinate(x, y-1)})
				default:
					next = append(next, Beam{Left, aoc.NewCoordinate(x-1, y)})
				}
			case Right:
				switch tile {
				case MirrorForward:
					next = append(next, Beam{Up, aoc.NewCoordinate(x, y-1)})
				case MirrorBackward:
					next = append(next, Beam{Down, aoc.NewCoordinate(x, y+1)})
				case SplitterVertical:
					next = append(next, Beam{Down, aoc.NewCoordinate(x, y+1)}, Beam{Up, aoc.NewCoordinate(x, y-1)})
				default:
					next = append(next, Beam{Right, aoc.NewCoordinate(x+1, y)})
				}
			}
		}
		current = next
	}

	result := []Beam{}
	for beam := range beams {
		result = append(result, beam)
	}
	return result
}

func energizedCoordinatesCount(bb []Beam) int {
	coordinates := make(map[aoc.Coordinate]struct{})
	for _, beam := range bb {
		coordinates[beam.coordinate] = struct{}{}
	}
	return len(coordinates)
}

func parseGrid(s string) *aoc.Grid {
	grid := aoc.NewGrid()
	for y, line := range strings.Split(s, "\n") {
		for x, r := range []byte(line) {
			c := aoc.NewCoordinate(x, y)
			tile := EmptySpace
			switch r {
			case '/':
				tile = MirrorForward
			case '\\':
				tile = MirrorBackward
			case '|':
				tile = SplitterVertical
			case '-':
				tile = SplitterHorizontal
			}
			grid.Add(aoc.NewNode(c, tile))
		}
	}
	return grid
}

type Solver struct {
	grid *aoc.Grid
}

func NewSolver(input string) *Solver {
	s := Solver{parseGrid(input)}
	return &s
}

func (s *Solver) Part1() int {
	bb := beams(s.grid, Beam{Right, aoc.NewCoordinate(0, 0)})
	return energizedCoordinatesCount(bb)
}

func (s *Solver) Part2() int {
	max := 0
	maxX, maxY := s.grid.MaxX(), s.grid.MaxY()
	for x := 0; x <= maxX; x++ {
		bb := beams(s.grid, Beam{Down, aoc.NewCoordinate(x, 0)})
		count := energizedCoordinatesCount(bb)
		if count > max {
			max = count
		}

		bb = beams(s.grid, Beam{Up, aoc.NewCoordinate(x, maxY)})
		count = energizedCoordinatesCount(bb)
		if count > max {
			max = count
		}
	}
	for y := 0; y <= maxY; y++ {
		bb := beams(s.grid, Beam{Right, aoc.NewCoordinate(0, y)})
		count := energizedCoordinatesCount(bb)
		if count > max {
			max = count
		}

		bb = beams(s.grid, Beam{Left, aoc.NewCoordinate(maxX, y)})
		count = energizedCoordinatesCount(bb)
		if count > max {
			max = count
		}
	}
	return max
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
