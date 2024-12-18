package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

const (
	ADV int = iota
	BXL
	BST
	JNZ
	BXC
	OUT
	BDV
	CDV
)

type Computer struct {
	InstructionPointer              int
	RegisterA, RegisterB, RegisterC int
	Output                          []int
}

func NewComputer(a, b, c int) *Computer {
	com := Computer{
		InstructionPointer: 0,
		RegisterA:          a,
		RegisterB:          b,
		RegisterC:          c,
		Output:             []int{},
	}

	return &com
}

func (c *Computer) RunProgram(program []int) []int {
	for c.InstructionPointer < len(program)-1 {
		c.Execute(program[c.InstructionPointer], program[c.InstructionPointer+1])
	}

	return c.Output
}

func (c *Computer) Execute(instruction, operand int) {
	combo := operand
	switch operand {
	case 4:
		combo = c.RegisterA
	case 5:
		combo = c.RegisterB
	case 6:
		combo = c.RegisterC
	}

	switch instruction {
	case ADV:
		c.RegisterA = c.RegisterA / int(math.Pow(2, float64(combo)))

	case BXL:
		c.RegisterB = c.RegisterB ^ operand

	case BST:
		c.RegisterB = combo % 8

	case JNZ:
		if c.RegisterA != 0 {
			c.InstructionPointer = operand
			return
		}

	case BXC:
		c.RegisterB = c.RegisterB ^ c.RegisterC

	case OUT:
		c.Output = append(c.Output, combo%8)

	case BDV:
		c.RegisterB = c.RegisterA / int(math.Pow(2, float64(combo)))

	case CDV:
		c.RegisterC = c.RegisterA / int(math.Pow(2, float64(combo)))
	}

	c.InstructionPointer += 2
}

type Solver struct {
	A, B, C      int
	Instructions []int
}

func NewSolver(input string) *Solver {
	lines := strings.Split(input, "\n")
	r := regexp.MustCompile(`\d+`)
	a, _ := strconv.Atoi(r.FindString(lines[0]))
	b, _ := strconv.Atoi(r.FindString(lines[1]))
	c, _ := strconv.Atoi(r.FindString(lines[2]))

	r = regexp.MustCompile(`(\d,?)+`)
	m := r.FindString(lines[4])

	ins := []int{}
	for _, i := range strings.Split(m, ",") {
		v, _ := strconv.Atoi(i)
		ins = append(ins, v)
	}

	s := Solver{a, b, c, ins}
	return &s
}

func (s *Solver) Part1() string {
	com := NewComputer(s.A, s.B, s.C)
	outputs := []string{}
	for _, o := range com.RunProgram(s.Instructions) {
		outputs = append(outputs, fmt.Sprint(o))
	}
	return strings.Join(outputs, ",")
}

func (s *Solver) Part2() int {
	v, ok := s.PossibleRegisterA(0)
	if !ok {
		log.Fatal("stupid")
	}

	return v
}

func (s *Solver) PossibleRegisterA(base int) (int, bool) {
	var zero int

	i := 0
	for {
		v := (base * 8) + i
		com := NewComputer(v, s.B, s.C)
		output := com.RunProgram(s.Instructions)

		if IntSliceHasSuffix(s.Instructions, output) {
			if len(s.Instructions) == len(output) {
				return v, true
			}

			if vv, ok := s.PossibleRegisterA(v); ok {
				return vv, true
			}
		}

		i++
		if i >= 8 {
			return zero, false
		}
	}
}

func IntSliceHasSuffix(a, b []int) bool {
	if len(b) > len(a) {
		return false
	}

	for i := 1; i <= len(b); i++ {
		if a[len(a)-i] != b[len(b)-i] {
			return false
		}
	}

	return true
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
