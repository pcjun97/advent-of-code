package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

const (
	OpNop = ""
	OpAnd = "AND"
	OpOr  = "OR"
	OpXor = "XOR"
)

type Gate struct {
	Name      string
	Operation string
	Value     bool
	Active    bool
	inputs    map[string]struct{}
	outputs   map[string]struct{}
}

func NewGate(name string) *Gate {
	gate := Gate{
		Name:      name,
		Operation: OpNop,
		inputs:    make(map[string]struct{}),
		outputs:   make(map[string]struct{}),
	}
	return &gate
}

func (g *Gate) Inputs() []string {
	list := []string{}
	for i := range g.inputs {
		list = append(list, i)
	}
	return list
}

func (g *Gate) Outputs() []string {
	list := []string{}
	for i := range g.outputs {
		list = append(list, i)
	}
	return list
}

func (g *Gate) AddInputs(names ...string) {
	for _, n := range names {
		g.inputs[n] = struct{}{}
	}
}

func (g *Gate) AddOutputs(names ...string) {
	for _, n := range names {
		g.outputs[n] = struct{}{}
	}
}

type Solver struct {
	Gates map[string]*Gate
}

func NewSolver(input string) *Solver {
	blocks := strings.Split(input, "\n\n")

	gates := make(map[string]*Gate)

	r := regexp.MustCompile(`(.*):\s+(\d)`)
	for _, line := range strings.Split(blocks[0], "\n") {
		m := r.FindStringSubmatch(line)
		name := m[1]
		v, _ := strconv.Atoi(m[2])
		gates[name] = NewGate(name)
		gates[name].Value = v == 1
		gates[name].Active = true
	}

	r = regexp.MustCompile(`(.*) (.*) (.*) -> (.*)`)
	for _, line := range strings.Split(blocks[1], "\n") {
		m := r.FindStringSubmatch(line)
		name := m[4]
		op := m[2]
		i1 := m[1]
		i2 := m[3]

		if _, ok := gates[name]; !ok {
			gates[name] = NewGate(name)
		}
		gates[name].Operation = op
		gates[name].AddInputs(i1, i2)

		if _, ok := gates[i1]; !ok {
			gates[i1] = NewGate(i1)
		}
		gates[i1].AddOutputs(name)

		if _, ok := gates[i2]; !ok {
			gates[i2] = NewGate(i2)
		}
		gates[i2].AddOutputs(name)
	}

	s := Solver{gates}
	return &s
}

func (s *Solver) InputsActive(name string) bool {
	g := s.Gates[name]

	for _, input := range g.Inputs() {
		if !s.Gates[input].Active {
			return false
		}
	}

	return true
}

func (s *Solver) LargestZ() string {
	maxz := 0
	for name := range s.Gates {
		if name[0] != 'z' {
			continue
		}

		v, _ := strconv.Atoi(name[1:])
		if v > maxz {
			maxz = v
		}
	}

	return fmt.Sprintf("z%0d", maxz)
}

func (s *Solver) Part1() int {
	tovisit := make(map[string]struct{})

	for _, g := range s.Gates {
		if !g.Active {
			continue
		}

		for _, o := range g.Outputs() {
			if !s.InputsActive(o) {
				continue
			}
			tovisit[o] = struct{}{}
		}
	}

	for len(tovisit) > 0 {
		var cur string
		for name := range tovisit {
			cur = name
			break
		}
		delete(tovisit, cur)

		g := s.Gates[cur]
		inputs := g.Inputs()
		v := s.Gates[inputs[0]].Value

		switch g.Operation {
		case OpAnd:
			for _, i := range inputs[1:] {
				v = v && s.Gates[i].Value
			}

		case OpOr:
			for _, i := range inputs[1:] {
				v = v || s.Gates[i].Value
			}

		case OpXor:
			for _, i := range inputs[1:] {
				v = v != s.Gates[i].Value
			}

		default:
			log.Fatalf("unknown operation: %s", g.Operation)
		}

		g.Value = v
		g.Active = true

		for _, o := range g.Outputs() {
			if !s.InputsActive(o) {
				continue
			}
			tovisit[o] = struct{}{}
		}
	}

	z := 0

	for name, g := range s.Gates {
		if name[0] != 'z' || !g.Value {
			continue
		}
		i, _ := strconv.Atoi(name[1:])
		z += int(math.Pow(2, float64(i)))
	}

	return z
}

func (s *Solver) Part2() string {
	wrong := []string{}

	for name := range s.Gates {
		if !s.IsValidGate(name) {
			wrong = append(wrong, name)
		}
	}

	sort.Strings(wrong)

	return strings.Join(wrong, ",")
}

func (s *Solver) IsValidGate(name string) bool {
	if name[0] == 'x' || name[0] == 'y' {
		return true
	}

	switch s.Gates[name].Operation {
	case OpAnd:
		return s.IsValidAnd(name)

	case OpXor:
		return s.IsValidXor(name)

	case OpOr:
		return s.IsValidCarry(name)

	default:
		return false
	}
}

func (s *Solver) IsValidXor(name string) bool {
	inputs := s.Gates[name].Inputs()

	if (inputs[0][0] == 'x' || inputs[0][0] == 'y') && inputs[0][1:] != "00" {
		outputs := s.Gates[name].Outputs()
		if len(outputs) != 2 {
			return false
		}

		o1 := s.Gates[outputs[0]]
		o2 := s.Gates[outputs[1]]

		if o1.Operation == OpXor && o2.Operation == OpAnd {
			return true
		}

		if o1.Operation == OpAnd && o2.Operation == OpXor {
			return true
		}

		return false
	}

	return name[0] == 'z'
}

func (s *Solver) IsValidAnd(name string) bool {
	inputs := s.Gates[name].Inputs()

	if (inputs[0][0] == 'x' || inputs[0][0] == 'y') && inputs[0][1:] == "00" {
		return s.IsValidCarry(name)
	}

	outputs := s.Gates[name].Outputs()
	if len(outputs) != 1 {
		return false
	}

	return s.Gates[outputs[0]].Operation == OpOr
}

func (s *Solver) IsValidCarry(name string) bool {
	if name == s.LargestZ() {
		return true
	}

	outputs := s.Gates[name].Outputs()
	if len(outputs) != 2 {
		return false
	}

	o1 := s.Gates[outputs[0]]
	o2 := s.Gates[outputs[1]]

	if o1.Operation == OpXor && o2.Operation == OpAnd {
		return true
	}

	if o1.Operation == OpAnd && o2.Operation == OpXor {
		return true
	}

	return false
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
