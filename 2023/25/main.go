package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Graph struct {
	vertex []string
	edges  [][2]string
}

func parseGraph(s string) Graph {
	vertex := []string{}
	edges := [][2]string{}

	vertexSet := make(map[string]struct{})
	r := regexp.MustCompile(`(.+): (.+)`)
	for _, line := range strings.Split(s, "\n") {
		m := r.FindStringSubmatch(line)
		v := m[1]
		vertexSet[v] = struct{}{}

		for _, neighbour := range strings.Split(m[2], " ") {
			vertexSet[neighbour] = struct{}{}

			edges = append(edges, [2]string{v, neighbour})
		}
	}

	for v := range vertexSet {
		vertex = append(vertex, v)
	}

	g := Graph{vertex, edges}
	return g
}

func (g Graph) BFSPaths() [][]string {
	paths := [][]string{}

	neighbours := make(map[string][]string)
	for _, edge := range g.edges {
		if _, ok := neighbours[edge[0]]; !ok {
			neighbours[edge[0]] = []string{}
		}
		neighbours[edge[0]] = append(neighbours[edge[0]], edge[1])

		if _, ok := neighbours[edge[1]]; !ok {
			neighbours[edge[1]] = []string{}
		}
		neighbours[edge[1]] = append(neighbours[edge[1]], edge[0])
	}

	var bfs func(string)
	bfs = func(from string) {
		visited := make(map[string]struct{})
		tovisit := [][]string{{from}}
		for len(tovisit) > 0 {
			path := tovisit[0]
			tovisit = tovisit[1:]
			if _, ok := visited[path[0]]; ok {
				continue
			}
			visited[path[0]] = struct{}{}

			if len(path) > 1 {
				paths = append(paths, path)
			}

			for _, n := range neighbours[path[0]] {
				p := append([]string{n}, path...)
				tovisit = append(tovisit, p)
			}
		}
	}

	for _, v := range g.vertex {
		bfs(v)
	}

	return paths
}

func (g Graph) Groups() [][]string {
	groups := [][]string{}

	ref := make(map[string][]string)
	for _, edge := range g.edges {
		if _, ok := ref[edge[0]]; !ok {
			ref[edge[0]] = []string{}
		}
		ref[edge[0]] = append(ref[edge[0]], edge[1])

		if _, ok := ref[edge[1]]; !ok {
			ref[edge[1]] = []string{}
		}
		ref[edge[1]] = append(ref[edge[1]], edge[0])
	}

	visited := make(map[string]struct{})
	for _, v := range g.vertex {
		if _, ok := visited[v]; ok {
			continue
		}

		group := make(map[string]struct{})
		tovisit := []string{v}
		for len(tovisit) > 0 {
			cur := tovisit[0]
			tovisit = tovisit[1:]
			if _, ok := group[cur]; ok {
				continue
			}

			tovisit = append(tovisit, ref[cur]...)
			group[cur] = struct{}{}
			visited[cur] = struct{}{}
		}

		list := []string{}
		for n := range group {
			list = append(list, n)
		}
		groups = append(groups, list)
	}

	return groups
}

type Solver struct {
	graph Graph
}

func NewSolver(input string) *Solver {
	s := Solver{parseGraph(input)}
	return &s
}

func (s *Solver) Part1() int {
	g := s.graph
	for i := 0; i < 3; i++ {
		vertex := make([]string, len(g.vertex))
		copy(vertex, g.vertex)

		edges := make([][2]string, len(g.edges))
		copy(edges, g.edges)

		g = Graph{vertex, edges}
		count := make(map[[2]string]int)
		for _, edge := range g.edges {
			count[edge] = 0
		}

		for _, path := range g.BFSPaths() {
			for j := 0; j < len(path)-1; j++ {
				a := path[j]
				b := path[j+1]
				if _, ok := count[[2]string{a, b}]; !ok {
					a, b = b, a
				}
				count[[2]string{a, b}] += 1
			}
		}

		max := g.edges[0]
		for edge, c := range count {
			if c > count[max] {
				max = edge
			}
		}

		edges = [][2]string{}
		for _, e := range g.edges {
			if e == max {
				continue
			}
			edges = append(edges, e)
		}
		g.edges = edges
	}

	groups := g.Groups()
	return len(groups[0]) * len(groups[1])
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())
}
