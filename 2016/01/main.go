package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Direction int

const (
	DirectionR Direction = iota
	DirectionL
)

type Cardinal int

const (
	North Cardinal = iota
	East
	South
	West
)

type Instruction struct {
	Direction
	Blocks int
}

func (i Instruction) nextCardinal(car Cardinal) Cardinal {
	switch i.Direction {
	case DirectionL:
		car = (car + 3) % 4
	case DirectionR:
		car = (car + 1) % 4
	default:
		log.Fatalf("unrecognize direction '%d'", i.Direction)
	}

	return car
}

func (i Instruction) nextCoordinate(c aoc.Coordinate, car Cardinal) aoc.Coordinate {
	switch car {
	case North:
		c.Y += i.Blocks
	case South:
		c.Y -= i.Blocks
	case East:
		c.X += i.Blocks
	case West:
		c.X -= i.Blocks
	default:
		log.Fatalf("unrecognize direction '%d'", c)
	}

	return c
}

func parseInstruction(s string) Instruction {
	r := regexp.MustCompile(`(.)(.*)`)
	m := r.FindStringSubmatch(s)

	var d Direction

	switch m[1] {
	case "R":
		d = DirectionR
	case "L":
		d = DirectionL
	default:
		log.Fatal("unrecognize direction")
	}

	b, err := strconv.Atoi(m[2])
	if err != nil {
		log.Fatal(err)
	}

	return Instruction{d, b}
}

type Solver struct {
	instructions []Instruction
}

func NewSolver(input string) *Solver {
	instructions := []Instruction{}
	for _, i := range strings.Split(input, ", ") {
		instructions = append(instructions, parseInstruction(i))
	}

	s := Solver{instructions}
	return &s
}

func (s *Solver) Part1() int {
	start := aoc.NewCoordinate(0, 0)
	end := aoc.NewCoordinate(0, 0)

	c := North
	for _, ins := range s.instructions {
		c = ins.nextCardinal(c)
		end = ins.nextCoordinate(end, c)
	}

	return start.ManhattanDistance(end)
}

func (s *Solver) Part2() int {
	visited := make(map[aoc.Coordinate]struct{})

	start := aoc.NewCoordinate(0, 0)
	cur := aoc.NewCoordinate(0, 0)

	c := North
	visited[cur] = struct{}{}

	for _, ins := range s.instructions {
		c = ins.nextCardinal(c)

		ii := Instruction{ins.Direction, 1}
		for i := 0; i < ins.Blocks; i++ {
			cur = ii.nextCoordinate(cur, c)
			if _, ok := visited[cur]; ok {
				return start.ManhattanDistance(cur)
			}
			visited[cur] = struct{}{}
		}
	}

	return -1
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
