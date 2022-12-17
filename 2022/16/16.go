package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func main() {
	input := aoc.ReadInput()

	valves := NewValves(input)
	fmt.Println(valves.MaxPressure("AA", 1, 30))
	fmt.Println(valves.MaxPressure("AA", 2, 26))
}

type Valves map[string]*Valve

func NewValves(s string) Valves {
	valves := make(Valves)

	lines := [][]string{}
	for _, line := range strings.Split(s, "\n") {
		fields := strings.Split(line, " ")
		lines = append(lines, fields)
		valves[fields[1]] = NewValve(fields[1])
	}

	for _, line := range lines {
		v := valves[line[1]]

		flowrate, err := strconv.Atoi(line[4][strings.IndexRune(line[4], '=')+1 : len(line[4])-1])
		if err != nil {
			panic(err)
		}

		v.SetFlowRate(flowrate)
		for _, n := range line[9:] {
			v.AddNeighbour(valves[strings.Trim(n, ",")], 1)
		}
	}

	valves.Djikstra()
	return valves
}

func (valves Valves) Djikstra() {
	for _, v := range valves {
		visit := make(map[string]bool)

		for name := range valves {
			visit[name] = true
		}
		delete(visit, v.label)

		for len(visit) > 0 {
			var nearest *Valve
			min := math.MaxInt
			for n, cost := range v.neighbours {
				if visit[n.label] && cost < min {
					nearest = n
					min = cost
				}
			}

			for n, cost := range nearest.neighbours {
				if n == v {
					continue
				}

				if _, ok := v.neighbours[n]; !ok || min+cost < v.neighbours[n] {
					v.neighbours[n] = min + cost
				}
			}

			delete(visit, nearest.label)
		}
	}
}

func (valves Valves) MaxPressure(start string, parallel, limit int) int {
	if _, ok := valves[start]; !ok {
		return 0
	}

	visit := make(map[string]bool)

	for label, v := range valves {
		if v.flowrate > 0 {
			visit[label] = true
		}
	}

	if _, ok := visit[start]; ok {
		delete(visit, start)
	}

	solver := NewSolver(valves, start, parallel, limit)
	return solver.MaxValue(nil)
}

type Solver struct {
	cache    map[int]int
	valves   map[int]*Valve
	limit    int
	parallel int
}

func NewSolver(valves Valves, start string, parallel, limit int) *Solver {
	s := Solver{
		cache:    make(map[int]int),
		valves:   make(map[int]*Valve),
		parallel: parallel,
		limit:    limit,
	}

	s.valves[0] = valves[start]
	i := 1
	for _, v := range valves {
		if v.flowrate > 0 {
			s.valves[i] = v
			i = i << 1
		}
	}

	s.DFS(0, 0, 0, 0)
	return &s
}

func (s *Solver) DFS(visited, t, value, last int) {
	if val, ok := s.cache[visited]; !ok || value > val {
		s.cache[visited] = value
	}

	for i := 0; i < len(s.valves)-1; i++ {
		mask := 1 << i
		if visited&mask != 0 {
			continue
		}

		v := s.valves[mask]
		prev := s.valves[last]
		cost := prev.neighbours[v] + 1
		remaining := s.limit - t - cost
		if remaining > 0 {
			s.DFS(visited|mask, t+cost, value+remaining*v.flowrate, mask)
		}
	}
}

func (s *Solver) MaxValue(others []int) int {
	if others == nil {
		others = []int{}
	}

	if len(others) == s.parallel {
		value := 0
		for _, v := range others {
			value += s.cache[v]
		}
		return value
	}

	max := 0
	sets := s.MutuallyExclusiveSets(others)
	for i := range sets {
		val := s.MaxValue(append(others, sets[i]))
		if val > max {
			max = val
		}
	}
	return max
}

func (s *Solver) MutuallyExclusiveSets(others []int) []int {
	sets := []int{}
	for visited := range s.cache {
		overlap := false
		for _, v := range others {
			if (visited & v) != 0 {
				overlap = true
			}
		}
		if !overlap {
			sets = append(sets, visited)
		}
	}
	return sets
}

type Valve struct {
	label      string
	flowrate   int
	neighbours map[*Valve]int
}

func NewValve(label string) *Valve {
	v := Valve{
		label:      label,
		neighbours: make(map[*Valve]int),
	}

	return &v
}

func (v *Valve) AddNeighbour(neighbour *Valve, distance int) {
	v.neighbours[neighbour] = distance
}

func (v *Valve) SetFlowRate(flowrate int) {
	v.flowrate = flowrate
}
