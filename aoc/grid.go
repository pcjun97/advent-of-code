package aoc

type Coordinate struct {
	X, Y int
}

func NewCoordinate(x, y int) Coordinate {
	c := Coordinate{x, y}
	return c
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
