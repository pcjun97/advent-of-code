package main

import (
	"fmt"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func main() {
	input := aoc.ReadInput()
	treesInput := strings.Split(input, "\n")
	trees := make([][]*Tree, len(treesInput))

	for i, line := range treesInput {
		trees[i] = make([]*Tree, len(line))
		for j, h := range line {
			height := int(h - '0')
			t := NewTree(height)

			if i > 0 {
				t.SetNeightbour(UP, trees[i-1][j])
			}

			if j > 0 {
				t.SetNeightbour(LEFT, trees[i][j-1])
			}

			trees[i][j] = t
		}
	}

	max := 0
	sum := 0
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			if trees[i][j].Visible() {
				sum++
			}
			if score := trees[i][j].ScenicScore(); score > max {
				max = score
			}
		}
	}

	fmt.Println(sum)
	fmt.Println(max)
}

type Direction int

const (
	UP Direction = iota
	LEFT
	DOWN
	RIGHT
)

type Tree struct {
	height     int
	neighbours [4]*Tree
}

func NewTree(height int) *Tree {
	t := Tree{
		height: height,
	}

	return &t
}

func (t *Tree) Visible() bool {
	for i, n := range t.neighbours {
		visible := true
		for n != nil {
			if n.height >= t.height {
				visible = false
				break
			}
			n = n.neighbours[i]
		}
		if visible {
			return true
		}
	}
	return false
}

func (t *Tree) ScenicScore() int {
	score := 1
	for direction, n := range t.neighbours {
		i := 0
		for n != nil && n.height < t.height {
			i++
			if n.height >= t.height {
				break
			}
			n = n.neighbours[direction]
		}
		score *= i
	}
	return score
}

func (t *Tree) SetNeightbour(direction Direction, neighbour *Tree) {
	t.neighbours[direction] = neighbour
	neighbour.neighbours[(direction+2)%4] = t
}
