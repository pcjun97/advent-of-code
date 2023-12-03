package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type CubeSet struct {
	Red   int
	Green int
	Blue  int
}

func parseCubeSet(s string) CubeSet {
	cs := CubeSet{0, 0, 0}

	for _, cube := range strings.Split(s, ", ") {
		r := regexp.MustCompile("^(.*) (.*)$")
		m := r.FindStringSubmatch(cube)

		count, _ := strconv.Atoi(m[1])
		color := m[2]

		switch color {
		case "red":
			cs.Red = count
		case "green":
			cs.Green = count
		case "blue":
			cs.Blue = count
		}
	}

	return cs
}

type Game struct {
	ID       int
	CubeSets []CubeSet
}

func parseGame(s string) *Game {
	cubesets := []CubeSet{}

	r := regexp.MustCompile(`Game (.*): (.*)$`)
	m := r.FindStringSubmatch(s)
	id, _ := strconv.Atoi(m[1])
	cubesetsStr := m[2]

	for _, csStr := range strings.Split(cubesetsStr, "; ") {
		cs := parseCubeSet(csStr)
		cubesets = append(cubesets, cs)
	}

	g := Game{
		ID:       id,
		CubeSets: cubesets,
	}

	return &g
}

type Solver struct {
	games []*Game
}

func NewSolver(input string) *Solver {
	games := []*Game{}
	for _, line := range strings.Split(input, "\n") {
		game := parseGame(line)
		games = append(games, game)
	}

	s := Solver{games}
	return &s
}

func (s *Solver) Part1() int {
	isGameValid := func(g *Game) bool {
		for _, cs := range g.CubeSets {
			if cs.Red > 12 || cs.Green > 13 || cs.Blue > 14 {
				return false
			}
		}
		return true
	}

	sum := 0
	for _, game := range s.games {
		if isGameValid(game) {
			sum += game.ID
		}
	}

	return sum
}

func (s *Solver) Part2() int {
	minCubeSet := func(g *Game) CubeSet {
		min := CubeSet{0, 0, 0}
		for _, cs := range g.CubeSets {
			if cs.Red > min.Red {
				min.Red = cs.Red
			}
			if cs.Green > min.Green {
				min.Green = cs.Green
			}
			if cs.Blue > min.Blue {
				min.Blue = cs.Blue
			}
		}
		return min
	}

	sum := 0
	for _, game := range s.games {
		min := minCubeSet(game)
		power := min.Red * min.Green * min.Blue
		sum += power
	}

	return sum
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
