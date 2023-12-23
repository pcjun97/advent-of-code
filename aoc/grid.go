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
	nodes                  map[Coordinate]*Node
	minX, maxX, minY, maxY int
}

func NewGrid() *Grid {
	g := Grid{make(map[Coordinate]*Node), math.MaxInt, math.MinInt, math.MaxInt, math.MinInt}
	return &g
}

func (g *Grid) Add(node *Node) {
	g.nodes[node.Coordinate] = node

	x, y := node.Coordinate.X, node.Coordinate.Y
	if x > g.maxX {
		g.maxX = x
	}
	if x < g.minX {
		g.minX = x
	}
	if y > g.maxY {
		g.maxY = y
	}
	if y < g.minY {
		g.minY = y
	}
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
	return g.maxX
}

func (g *Grid) MaxY() int {
	return g.maxY
}

func (g *Grid) MinX() int {
	return g.minX
}

func (g *Grid) MinY() int {
	return g.minY
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
