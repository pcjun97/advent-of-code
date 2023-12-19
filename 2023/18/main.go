package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Instruction struct {
	Direction
	count int
	code  string
}

func parsePlan(s string) []Instruction {
	instructions := []Instruction{}

	r := regexp.MustCompile(`(.) (.+) \(#(.+)\)`)
	for _, line := range strings.Split(s, "\n") {
		m := r.FindStringSubmatch(line)
		var d Direction
		switch m[1] {
		case "U":
			d = Up
		case "D":
			d = Down
		case "L":
			d = Left
		case "R":
			d = Right
		}

		count, _ := strconv.Atoi(m[2])

		ins := Instruction{d, count, m[3]}
		instructions = append(instructions, ins)
	}

	return instructions
}

func parseHexCode(s string) Instruction {
	r := regexp.MustCompile(`(.....)(.)`)
	m := r.FindStringSubmatch(s)
	var d Direction
	switch m[2] {
	case "0":
		d = Right
	case "1":
		d = Down
	case "2":
		d = Left
	case "3":
		d = Up
	}
	count, _ := strconv.ParseInt(m[1], 16, 0)
	ins := Instruction{d, int(count), ""}
	return ins
}

type Solver struct {
	plan []Instruction
}

func NewSolver(input string) *Solver {
	s := Solver{parsePlan(input)}
	return &s
}

func (s *Solver) Part1() int {
	coordinates := []aoc.Coordinate{}

	sum := 0
	current := aoc.NewCoordinate(0, 0)
	for _, ins := range s.plan {
		coordinates = append(coordinates, current)
		switch ins.Direction {
		case Up:
			current.Y -= ins.count
		case Down:
			current.Y += ins.count
		case Left:
			current.X -= ins.count
		case Right:
			current.X += ins.count
		}
		sum += ins.count
	}

	return polygonArea(coordinates) + (sum / 2) + 1
}

func (s *Solver) Part2() int {
	coordinates := []aoc.Coordinate{}

	sum := 0
	current := aoc.NewCoordinate(0, 0)
	for _, ins := range s.plan {
		coordinates = append(coordinates, current)
		hins := parseHexCode(ins.code)
		switch hins.Direction {
		case Up:
			current.Y -= hins.count
		case Down:
			current.Y += hins.count
		case Left:
			current.X -= hins.count
		case Right:
			current.X += hins.count
		}
		sum += hins.count
	}

	return polygonArea(coordinates) + (sum / 2) + 1
}

func polygonArea(coordinates []aoc.Coordinate) int {
	area := 0
	j := len(coordinates) - 1
	for i, c := range coordinates {
		width := c.X + coordinates[j].X
		height := c.Y - coordinates[j].Y
		area += width * height
		j = i
	}
	return (area / 2)
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
