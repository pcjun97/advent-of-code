package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func main() {
	input := aoc.ReadInput()

	cave := NewCave(input, false)
	cave.SandFall()
	fmt.Println(cave.SandCount())

	caveWithFloor := NewCave(input, true)
	caveWithFloor.SandFall()
	fmt.Println(caveWithFloor.SandCount())
}

type Coordinate struct {
	X, Y int
}

type Content int

const (
	AIR Content = iota
	ROCK
	SAND
	SOURCE
	VOID
)

type Cave struct {
	grid             map[int]map[int]Content
	minX, maxX, maxY int
	source           Coordinate
	floor            bool
}

func NewCave(s string, floor bool) *Cave {
	minX := 500
	maxX := 500
	maxY := 0

	lines := strings.Split(s, "\n")
	cg := make([][]Coordinate, len(lines))

	for i, line := range lines {
		coordinatesStr := strings.Split(line, " -> ")
		cg[i] = make([]Coordinate, len(coordinatesStr))
		for j, c := range coordinatesStr {
			xy := strings.Split(c, ",")

			x, err := strconv.Atoi(xy[0])
			if err != nil {
				panic(err)
			}

			y, err := strconv.Atoi(xy[1])
			if err != nil {
				panic(err)
			}

			cg[i][j] = Coordinate{x, y}

			if x < minX {
				minX = x
			}

			if x > maxX {
				maxX = x
			}

			if y > maxY {
				maxY = y
			}
		}
	}

	grid := make(map[int]map[int]Content)

	cave := Cave{
		grid:   grid,
		minX:   minX,
		maxX:   maxX,
		maxY:   maxY,
		source: Coordinate{500, 0},
		floor:  floor,
	}

	cave.setTile(cave.source, SOURCE)

	for _, line := range cg {
		for i := 0; i < len(line)-1; i++ {
			cave.addRockPath(line[i], line[i+1])
		}
	}

	return &cave
}

func (cave *Cave) setTile(tile Coordinate, content Content) {
	if _, ok := cave.grid[tile.X]; !ok {
		cave.grid[tile.X] = make(map[int]Content)
	}
	cave.grid[tile.X][tile.Y] = content

	if tile.X < cave.minX {
		cave.minX = tile.X
	}

	if tile.X > cave.maxX {
		cave.maxX = tile.X
	}
}

func (cave *Cave) Tile(tile Coordinate) Content {
	if tile.X < 0 || tile.Y < 0 {
		return VOID
	}

	if !cave.floor && (tile.X > cave.maxX || tile.X < cave.minX || tile.Y > cave.maxY) {
		return VOID
	}

	if cave.floor && tile.Y == cave.maxY+2 {
		return ROCK
	}

	if _, ok := cave.grid[tile.X]; !ok {
		return AIR
	}

	if _, ok := cave.grid[tile.X][tile.Y]; !ok {
		return AIR
	}

	return cave.grid[tile.X][tile.Y]
}

func (cave *Cave) addRockPath(start Coordinate, end Coordinate) {
	if start.X != end.X && start.Y != end.Y {
		return
	}

	var i, j int

	if start.X == end.X {
		if start.Y < end.Y {
			i = start.Y
			j = end.Y
		} else {
			i = end.Y
			j = start.Y
		}

		for i <= j {
			cave.setTile(Coordinate{start.X, i}, ROCK)
			i += 1
		}

		return
	}

	if start.X < end.X {
		i = start.X
		j = end.X
	} else {
		i = end.X
		j = start.X
	}

	for i <= j {
		cave.setTile(Coordinate{i, start.Y}, ROCK)
		i += 1
	}
}

func (cave *Cave) nextSandPosition() Coordinate {
	c := cave.source

	for {
		c.Y += 1
		if cave.Tile(c) == VOID {
			return Coordinate{-1, -1}
		}
		if cave.Tile(c) == AIR {
			continue
		}

		c.X -= 1
		if cave.Tile(c) == VOID {
			return Coordinate{-1, -1}
		}
		if cave.Tile(c) == AIR {
			continue
		}

		c.X += 2
		if cave.Tile(c) == VOID {
			return Coordinate{-1, -1}
		}
		if cave.Tile(c) == AIR {
			continue
		}

		c.Y -= 1
		c.X -= 1
		return c
	}
}

func (cave *Cave) SandFall() {
	for {
		p := cave.nextSandPosition()

		if cave.Tile(p) == VOID {
			break
		}

		if p.X == cave.source.X && p.Y == cave.source.Y {
			break
		}

		cave.setTile(p, SAND)
	}
}

func (cave *Cave) SandCount() int {
	sum := 0
	for i := range cave.grid {
		for j := range cave.grid[i] {
			if cave.grid[i][j] == SAND {
				sum++
			}
		}
	}

	if cave.floor {
		return sum + 1
	}
	return sum
}
