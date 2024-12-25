package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Network struct {
	set map[string]struct{}
}

func NewNetwork() *Network {
	set := make(map[string]struct{})
	return &Network{set}
}

func (n *Network) Add(computer string) {
	n.set[computer] = struct{}{}
}

func (n *Network) Size() int {
	return len(n.set)
}

func (n *Network) HasComputerWithPrefix(prefix string) bool {
	for c := range n.set {
		if strings.HasPrefix(c, prefix) {
			return true
		}
	}
	return false
}

type Solver struct {
	Connections [][2]string
}

func NewSolver(input string) *Solver {
	connections := [][2]string{}
	r := regexp.MustCompile(`^(.*)-(.*)$`)

	for _, line := range strings.Split(input, "\n") {
		m := r.FindStringSubmatch(line)
		conn := [2]string{m[1], m[2]}
		connections = append(connections, conn)
	}

	s := Solver{connections}
	return &s
}

func (s *Solver) Neighbors() map[string][]string {
	neighbours := make(map[string]map[string]struct{})

	for _, conn := range s.Connections {
		if _, ok := neighbours[conn[0]]; !ok {
			neighbours[conn[0]] = make(map[string]struct{})
		}
		if _, ok := neighbours[conn[1]]; !ok {
			neighbours[conn[1]] = make(map[string]struct{})
		}
		neighbours[conn[0]][conn[1]] = struct{}{}
		neighbours[conn[1]][conn[0]] = struct{}{}
	}

	list := make(map[string][]string)
	for c, n := range neighbours {
		list[c] = []string{}
		for nn := range n {
			list[c] = append(list[c], nn)
		}
	}

	return list
}

func StringSliceContains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}

	return false
}

func (s *Solver) Part1() int {
	visited := make(map[string]struct{})

	sets := [][]string{}

	neighbours := s.Neighbors()

	var dfs func(candidates, set []string)
	dfs = func(candidates, set []string) {
		for i, c := range candidates {
			if _, ok := visited[c]; ok {
				continue
			}

			ok := true
			for _, s := range set {
				if !StringSliceContains(neighbours[c], s) {
					ok = false
					break
				}
			}
			if !ok {
				continue
			}

			newset := append([]string{c}, set...)

			if len(newset) == 3 {
				sets = append(sets, newset)
				continue
			}

			dfs(candidates[i+1:], newset)
		}
	}

	for c, ns := range neighbours {
		dfs(ns, []string{c})
		visited[c] = struct{}{}
	}

	setsStartsWithT := [][]string{}
	for _, set := range sets {
		if set[0][0] == 't' || set[1][0] == 't' || set[2][0] == 't' {
			setsStartsWithT = append(setsStartsWithT, set)
		}
	}

	return len(setsStartsWithT)
}

func (s *Solver) Part2() string {
	visited := make(map[string]struct{})

	maxset := []string{}

	neighbours := s.Neighbors()

	var dfs func(candidates, set []string)
	dfs = func(candidates, set []string) {
		for i, c := range candidates {
			if _, ok := visited[c]; ok {
				continue
			}

			ok := true
			for _, s := range set {
				if !StringSliceContains(neighbours[c], s) {
					ok = false
					break
				}
			}
			if !ok {
				continue
			}

			newset := append([]string{c}, set...)

			if len(newset) > len(maxset) {
				maxset = newset
			}

			dfs(candidates[i+1:], newset)
		}
	}

	for c, ns := range neighbours {
		dfs(ns, []string{c})
		visited[c] = struct{}{}
	}

	sort.Strings(maxset)

	return strings.Join(maxset, ",")
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
