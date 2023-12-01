package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type InstructionType int

const (
	TurnOn InstructionType = iota
	TurnOff
	Toggle
)

type Instruction struct {
	instructionType InstructionType
	from            Coordinate
	to              Coordinate
}

func parseInstruction(s string) (Instruction, error) {
	var zero Instruction
	i := Instruction{}

	r := regexp.MustCompile(`(.*) (.*),(.*) through (.*),(.*)`)
	m := r.FindStringSubmatch(s)

	switch m[1] {
	case "toggle":
		i.instructionType = Toggle
	case "turn on":
		i.instructionType = TurnOn
	case "turn off":
		i.instructionType = TurnOff
	default:
		return zero, fmt.Errorf("unknown instruction type '%s'", m[1])
	}

	fromX, err := strconv.Atoi(m[2])
	if err != nil {
		return zero, err
	}

	fromY, err := strconv.Atoi(m[3])
	if err != nil {
		return zero, err
	}

	toX, err := strconv.Atoi(m[4])
	if err != nil {
		return zero, err
	}

	toY, err := strconv.Atoi(m[5])
	if err != nil {
		return zero, err
	}

	i.from = Coordinate{fromX, fromY}
	i.to = Coordinate{toX, toY}

	return i, nil
}

type Coordinate struct {
	X int
	Y int
}

type Solver struct {
	instructions []Instruction
}

func NewSolver(input string) (*Solver, error) {
	instructions := []Instruction{}
	for _, line := range strings.Split(input, "\n") {
		i, err := parseInstruction(line)
		if err != nil {
			return nil, err
		}

		instructions = append(instructions, i)
	}

	s := Solver{instructions}
	return &s, nil
}

func (s *Solver) Part1() int {
	lights := make(map[Coordinate]bool)

	for _, ins := range s.instructions {
		var minX, maxX, minY, maxY int
		if ins.from.X > ins.to.X {
			maxX = ins.from.X
			minX = ins.to.X
		} else {
			minX = ins.from.X
			maxX = ins.to.X
		}

		if ins.from.Y > ins.to.Y {
			maxY = ins.from.Y
			minY = ins.to.Y
		} else {
			minY = ins.from.Y
			maxY = ins.to.Y
		}

		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				c := Coordinate{x, y}

				switch ins.instructionType {
				case Toggle:
					if _, ok := lights[c]; !ok {
						lights[c] = false
					}
					lights[c] = !lights[c]
				case TurnOn:
					lights[c] = true
				case TurnOff:
					lights[c] = false
				}
			}
		}
	}

	sum := 0
	for _, v := range lights {
		if v {
			sum += 1
		}
	}

	return sum
}

func (s *Solver) Part2() int {
	lights := make(map[Coordinate]int)

	for _, ins := range s.instructions {
		var minX, maxX, minY, maxY int
		if ins.from.X > ins.to.X {
			maxX = ins.from.X
			minX = ins.to.X
		} else {
			minX = ins.from.X
			maxX = ins.to.X
		}

		if ins.from.Y > ins.to.Y {
			maxY = ins.from.Y
			minY = ins.to.Y
		} else {
			minY = ins.from.Y
			maxY = ins.to.Y
		}

		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				c := Coordinate{x, y}

				switch ins.instructionType {
				case Toggle:
					if _, ok := lights[c]; !ok {
						lights[c] = 0
					}
					lights[c] += 2
				case TurnOn:
					lights[c] += 1
				case TurnOff:
					if lights[c] > 0 {
						lights[c] -= 1
					}
				}
			}
		}
	}

	sum := 0
	for _, v := range lights {
		sum += v
	}

	return sum
}

func main() {
	input := aoc.ReadInput()
	s, err := NewSolver(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
