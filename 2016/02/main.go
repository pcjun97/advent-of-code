package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Instruction int

const (
	InstructionU Instruction = iota
	InstructionD
	InstructionL
	InstructionR
)

func (i Instruction) nextCoordinate(c aoc.Coordinate, g *aoc.Grid) aoc.Coordinate {
	cc := c

	switch i {
	case InstructionU:
		c.Y -= 1

	case InstructionD:
		c.Y += 1

	case InstructionL:
		c.X -= 1

	case InstructionR:
		c.X += 1
	}

	if g.Get(c) == nil {
		return cc
	}

	return c
}

type Solver struct {
	instructions [][]Instruction
}

func NewSolver(input string) *Solver {
	instructions := [][]Instruction{}
	for _, line := range strings.Split(input, "\n") {
		ins := []Instruction{}
		for _, c := range line {
			switch c {
			case 'U':
				ins = append(ins, InstructionU)
			case 'D':
				ins = append(ins, InstructionD)
			case 'L':
				ins = append(ins, InstructionL)
			case 'R':
				ins = append(ins, InstructionR)
			}
		}
		instructions = append(instructions, ins)
	}

	s := Solver{instructions}
	return &s
}

func (s *Solver) Part1() string {
	keypad := aoc.NewGrid()
	keypad.Add(aoc.NewNode(aoc.NewCoordinate(-1, -1), '1'))
	keypad.Add(aoc.NewNode(aoc.NewCoordinate(0, -1), '2'))
	keypad.Add(aoc.NewNode(aoc.NewCoordinate(1, -1), '3'))
	keypad.Add(aoc.NewNode(aoc.NewCoordinate(-1, 0), '4'))
	keypad.Add(aoc.NewNode(aoc.NewCoordinate(0, 0), '5'))
	keypad.Add(aoc.NewNode(aoc.NewCoordinate(1, 0), '6'))
	keypad.Add(aoc.NewNode(aoc.NewCoordinate(-1, 1), '7'))
	keypad.Add(aoc.NewNode(aoc.NewCoordinate(0, 1), '8'))
	keypad.Add(aoc.NewNode(aoc.NewCoordinate(1, 1), '9'))

	passcode := ""
	c := aoc.NewCoordinate(0, 0)

	for _, ins := range s.instructions {
		for _, i := range ins {
			c = i.nextCoordinate(c, keypad)
		}
		passcode = fmt.Sprintf("%s%c", passcode, keypad.Get(c).Value())
	}

	return passcode
}

func (s *Solver) Part2() string {
	keypad := aoc.NewGrid()
	keypad.Add(aoc.NewNode(aoc.NewCoordinate(0, -2), '1'))
	keypad.Add(aoc.NewNode(aoc.NewCoordinate(-1, -1), '2'))
	keypad.Add(aoc.NewNode(aoc.NewCoordinate(0, -1), '3'))
	keypad.Add(aoc.NewNode(aoc.NewCoordinate(1, -1), '4'))
	keypad.Add(aoc.NewNode(aoc.NewCoordinate(-2, 0), '5'))
	keypad.Add(aoc.NewNode(aoc.NewCoordinate(-1, 0), '6'))
	keypad.Add(aoc.NewNode(aoc.NewCoordinate(0, 0), '7'))
	keypad.Add(aoc.NewNode(aoc.NewCoordinate(1, 0), '8'))
	keypad.Add(aoc.NewNode(aoc.NewCoordinate(2, 0), '9'))
	keypad.Add(aoc.NewNode(aoc.NewCoordinate(-1, 1), 'A'))
	keypad.Add(aoc.NewNode(aoc.NewCoordinate(0, 1), 'B'))
	keypad.Add(aoc.NewNode(aoc.NewCoordinate(1, 1), 'C'))
	keypad.Add(aoc.NewNode(aoc.NewCoordinate(0, 2), 'D'))

	passcode := ""
	c := aoc.NewCoordinate(-2, 0)

	for _, ins := range s.instructions {
		for _, i := range ins {
			c = i.nextCoordinate(c, keypad)
		}
		passcode = fmt.Sprintf("%s%c", passcode, keypad.Get(c).Value())
	}

	return passcode
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
