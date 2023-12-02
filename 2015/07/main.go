package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type GateType int

const (
	SignalGate GateType = iota
	AndGate
	LShiftGate
	NotGate
	OrGate
	RShiftGate
)

type Gate struct {
	gateType GateType
	input1   string
	input2   string
	output   string
}

func parseGate(s string) Gate {
	g := Gate{}

	r := regexp.MustCompile(`(.*) -> (.*)`)
	m := r.FindStringSubmatch(s)
	inputs := m[1]
	g.output = m[2]

	r = regexp.MustCompile(`(.*) (.*) (.*)`)
	m = r.FindStringSubmatch(inputs)
	if m != nil {
		switch m[2] {
		case "AND":
			g.gateType = AndGate
		case "OR":
			g.gateType = OrGate
		case "LSHIFT":
			g.gateType = LShiftGate
		case "RSHIFT":
			g.gateType = RShiftGate
		}
		g.input1 = m[1]
		g.input2 = m[3]
		return g
	}

	r = regexp.MustCompile(`NOT (.*)`)
	m = r.FindStringSubmatch(inputs)
	if m != nil {
		g.gateType = NotGate
		g.input1 = m[1]
		return g
	}

	g.gateType = SignalGate
	g.input1 = inputs
	return g
}

type Solver struct {
	gates map[string]Gate
	cache map[string]uint16
}

func NewSolver(input string) *Solver {
	s := Solver{
		gates: make(map[string]Gate),
		cache: make(map[string]uint16),
	}

	for _, line := range strings.Split(input, "\n") {
		gate := parseGate(line)
		s.gates[gate.output] = gate
	}

	return &s
}

func (s *Solver) valueOf(gate string) uint16 {
	if v, ok := s.cache[gate]; ok {
		return v
	}

	value, err := strconv.ParseUint(gate, 10, 16)
	if err == nil {
		s.cache[gate] = uint16(value)
		return uint16(value)
	}

	var v uint16
	g := s.gates[gate]
	switch g.gateType {
	case AndGate:
		v = s.valueOf(g.input1) & s.valueOf(g.input2)
	case OrGate:
		v = s.valueOf(g.input1) | s.valueOf(g.input2)
	case NotGate:
		v = ^s.valueOf(g.input1)
	case LShiftGate:
		v = s.valueOf(g.input1) << s.valueOf(g.input2)
	case RShiftGate:
		v = s.valueOf(g.input1) >> s.valueOf(g.input2)
	default:
		v = s.valueOf(g.input1)
	}

	s.cache[gate] = v
	return v
}

func (s *Solver) Part1() int {
	s.cache = make(map[string]uint16)
	return int(s.valueOf("a"))
}

func (s *Solver) Part2() int {
	s.cache = make(map[string]uint16)
	a := s.valueOf("a")

	s.cache = make(map[string]uint16)
	s.cache["b"] = a
	return int(s.valueOf("a"))
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
