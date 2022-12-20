package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func main() {
	input := aoc.ReadInput()

	list := NewList(input, 1)
	list.Encrypt(1)
	sum := 0
	for _, n := range list.GroveCoordinates() {
		sum += n
	}
	fmt.Println(sum)

	list = NewList(input, 811589153)
	list.Encrypt(10)
	sum = 0
	for _, n := range list.GroveCoordinates() {
		sum += n
	}
	fmt.Println(sum)
}

type List struct {
	nodes []*Node
	zero  *Node
}

func NewList(s string, key int) *List {
	lines := strings.Split(s, "\n")
	nodes := make([]*Node, len(lines))

	var prev, zero *Node
	prev = nil

	for i, line := range strings.Split(s, "\n") {
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		node := Node{
			Prev:  prev,
			Value: value * key,
		}
		nodes[i] = &node

		if value == 0 {
			zero = &node
		}
		if prev != nil {
			prev.Next = &node
		}
		prev = &node
	}

	if len(nodes) > 0 {
		nodes[0].Prev = nodes[len(nodes)-1]
		nodes[len(nodes)-1].Next = nodes[0]
	}

	l := List{nodes, zero}
	return &l
}

func (l *List) Encrypt(n int) {
	for i := 0; i < n; i++ {
		for _, node := range l.nodes {
			l.Move(node)
		}
	}
}

func (l *List) String() string {
	node := l.zero
	s := make([]string, len(l.nodes))
	for i := 0; i < len(l.nodes); i++ {
		s[i] = strconv.Itoa(node.Value)
		node = node.Next
	}
	return "[" + strings.Join(s, " ") + "]"
}

func (l *List) Move(node *Node) {
	steps := node.Value % (len(l.nodes) - 1)
	if steps == 0 {
		return
	}

	if steps < 0 {
		steps = len(l.nodes) + steps - 1
	}

	prev := node.Prev
	next := node.Next
	prev.Next = node.Next
	next.Prev = node.Prev

	prev = node
	for i := 0; i < steps; i++ {
		prev = prev.Next
	}
	next = prev.Next

	prev.Next = node
	node.Prev = prev
	next.Prev = node
	node.Next = next
}

func (l *List) GroveCoordinates() [3]int {
	c := [3]int{}

	n1000 := 1000 % len(l.nodes)
	n2000 := 2000 % len(l.nodes)
	n3000 := 3000 % len(l.nodes)

	node := l.zero
	for i := 0; i < len(l.nodes); i++ {
		if i == n1000 {
			c[0] = node.Value
		}
		if i == n2000 {
			c[1] = node.Value
		}
		if i == n3000 {
			c[2] = node.Value
		}
		node = node.Next
	}

	return c
}

type Node struct {
	Prev, Next *Node
	Value      int
}
