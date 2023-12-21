package aoc

import (
	"math"
)

type Coordinate struct {
	X, Y int
}

func NewCoordinate(x, y int) Coordinate {
	c := Coordinate{x, y}
	return c
}

func (c Coordinate) ManhattanDistance(cc Coordinate) int {
	dx := c.X - cc.X
	if dx < 0 {
		dx *= -1
	}

	dy := c.Y - cc.Y
	if dy < 0 {
		dy *= -1
	}

	return dx + dy
}

func (c Coordinate) Neighbors4Way() []Coordinate {
	result := []Coordinate{}
	result = append(result, Coordinate{c.X, c.Y - 1})
	result = append(result, Coordinate{c.X, c.Y + 1})
	result = append(result, Coordinate{c.X - 1, c.Y})
	result = append(result, Coordinate{c.X + 1, c.Y})
	return result
}

type Node struct {
	Coordinate
	v int
}

func NewNode(c Coordinate, v int) *Node {
	node := Node{c, v}
	return &node
}

func (n *Node) Set(v int) {
	n.v = v
}

func (n *Node) Value() int {
	return n.v
}

type Grid struct {
	nodes map[Coordinate]*Node
}

func NewGrid() *Grid {
	g := Grid{make(map[Coordinate]*Node)}
	return &g
}

func (g *Grid) Add(node *Node) {
	g.nodes[node.Coordinate] = node
}

func (g *Grid) Get(c Coordinate) *Node {
	n, ok := g.nodes[c]
	if !ok {
		return nil
	}
	return n
}

func (g *Grid) Nodes() []*Node {
	nodes := []*Node{}
	for _, n := range g.nodes {
		nodes = append(nodes, n)
	}
	return nodes
}

func (g *Grid) MaxX() int {
	max := math.MinInt
	for _, n := range g.nodes {
		if n.X > max {
			max = n.X
		}
	}
	return max
}

func (g *Grid) MaxY() int {
	max := math.MinInt
	for _, n := range g.nodes {
		if n.Y > max {
			max = n.Y
		}
	}
	return max
}

func (g *Grid) MinX() int {
	min := math.MaxInt
	for _, n := range g.nodes {
		if n.X < min {
			min = n.X
		}
	}
	return min
}

func (g *Grid) MinY() int {
	min := math.MaxInt
	for _, n := range g.nodes {
		if n.Y < min {
			min = n.Y
		}
	}
	return min
}

func (g *Grid) Neighbors8Way(node *Node) []*Node {
	neighbors := []*Node{}
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}

			c := Coordinate{node.Coordinate.X + x, node.Coordinate.Y + y}
			if neighbor, ok := g.nodes[c]; ok {
				neighbors = append(neighbors, neighbor)
			}
		}
	}
	return neighbors
}

func (g *Grid) Neighbors4Way(node *Node) []*Node {
	neighbors := []*Node{}

	c := Coordinate{node.X - 1, node.Y}
	if n, ok := g.nodes[c]; ok {
		neighbors = append(neighbors, n)
	}

	c = Coordinate{node.X + 1, node.Y}
	if n, ok := g.nodes[c]; ok {
		neighbors = append(neighbors, n)
	}

	c = Coordinate{node.X, node.Y - 1}
	if n, ok := g.nodes[c]; ok {
		neighbors = append(neighbors, n)
	}

	c = Coordinate{node.X, node.Y + 1}
	if n, ok := g.nodes[c]; ok {
		neighbors = append(neighbors, n)
	}

	return neighbors
}
