package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func parseGrid(s string) *aoc.Grid {
	grid := aoc.NewGrid()
	for y, line := range strings.Split(s, "\n") {
		for x, r := range []byte(line) {
			c := aoc.NewCoordinate(x, y)
			v := r - '0'
			grid.Add(aoc.NewNode(c, int(v)))
		}
	}
	return grid
}

type Solver struct {
	grid  *aoc.Grid
	final aoc.Coordinate
}

func NewSolver(input string) *Solver {
	grid := parseGrid(input)
	maxX, maxY := grid.MaxX(), grid.MaxY()
	final := aoc.NewCoordinate(maxX, maxY)
	s := Solver{grid, final}
	return &s
}

func (s *Solver) Part1() int {
	return s.minHeatLoss(1, 3)
}

func (s *Solver) Part2() int {
	return s.minHeatLoss(4, 10)
}

type CacheKey struct {
	direction  Direction
	coordinate aoc.Coordinate
	count      int
}

func (s *Solver) minHeatLoss(minStraight, maxStraight int) int {
	cache := make(map[CacheKey]int)
	right := aoc.NewCoordinate(1, 0)
	down := aoc.NewCoordinate(0, 1)

	states := NewBinaryHeap()
	states.Push(State{Right, right, 1, s.grid.Get(right).Value()})
	states.Push(State{Down, down, 1, s.grid.Get(down).Value()})

	for states.Size() > 0 {
		state := states.Pop()
		if state.coordinate.X == s.final.X && state.coordinate.Y == s.final.Y {
			if state.count < minStraight {
				continue
			}
			return state.heatloss
		}

		key := CacheKey{state.direction, state.coordinate, state.count}
		if heatloss, ok := cache[key]; ok && heatloss <= state.heatloss {
			continue
		}
		cache[key] = state.heatloss
		next := s.nextStates(state, minStraight, maxStraight)
		for _, n := range next {
			states.Push(n)
		}
	}
	return -1
}

type State struct {
	direction  Direction
	coordinate aoc.Coordinate
	count      int
	heatloss   int
}

func (s *Solver) nextStates(state State, minStraight, maxStraight int) []State {
	states := []State{}

	if state.count < maxStraight {
		next := state
		next.count += 1
		switch state.direction {
		case Up:
			next.coordinate.Y -= 1
		case Down:
			next.coordinate.Y += 1
		case Left:
			next.coordinate.X -= 1
		case Right:
			next.coordinate.X += 1
		}

		if node := s.grid.Get(next.coordinate); node != nil {
			next.heatloss += node.Value()
			states = append(states, next)
		}
	}

	if state.count >= minStraight {
		left := state
		left.count = 1
		switch state.direction {
		case Up:
			left.direction = Left
			left.coordinate.X -= 1
		case Down:
			left.direction = Right
			left.coordinate.X += 1
		case Left:
			left.direction = Down
			left.coordinate.Y += 1
		case Right:
			left.direction = Up
			left.coordinate.Y -= 1
		}
		if node := s.grid.Get(left.coordinate); node != nil {
			left.heatloss += node.Value()
			states = append(states, left)
		}

		right := state
		right.count = 1
		switch state.direction {
		case Up:
			right.direction = Right
			right.coordinate.X += 1
		case Down:
			right.direction = Left
			right.coordinate.X -= 1
		case Left:
			right.direction = Up
			right.coordinate.Y -= 1
		case Right:
			right.direction = Down
			right.coordinate.Y += 1
		}
		if node := s.grid.Get(right.coordinate); node != nil {
			right.heatloss += node.Value()
			states = append(states, right)
		}
	}

	return states
}

type BinaryHeap struct {
	list []State
}

func NewBinaryHeap() *BinaryHeap {
	list := []State{}
	heap := BinaryHeap{list}
	return &heap
}

func (heap *BinaryHeap) Size() int {
	return len(heap.list)
}

func (heap *BinaryHeap) Push(state State) {
	heap.list = append(heap.list, state)
	heap.heapifyUp(len(heap.list) - 1)
}

func (heap *BinaryHeap) Pop() State {
	if len(heap.list) <= 0 {
		return State{}
	}
	state := heap.list[0]
	heap.list[0] = heap.list[len(heap.list)-1]
	heap.list = heap.list[:len(heap.list)-1]
	heap.heapifyDown(0)
	return state
}

func (heap *BinaryHeap) heapifyDown(i int) {
	if i < 0 || i >= len(heap.list) {
		return
	}

	left := (2 * i) + 1
	if len(heap.list) <= left {
		return
	}

	if heap.list[left].heatloss < heap.list[i].heatloss {
		heap.list[i], heap.list[left] = heap.list[left], heap.list[i]
		heap.heapifyDown(left)
	}

	right := (2 * i) + 2
	if len(heap.list) <= right {
		return
	}
	if heap.list[right].heatloss < heap.list[i].heatloss {
		heap.list[i], heap.list[right] = heap.list[right], heap.list[i]
		heap.heapifyDown(right)
	}
}

func (heap *BinaryHeap) heapifyUp(i int) {
	if i < 0 || i >= len(heap.list) {
		return
	}

	parent := (i - 1) / 2
	if parent < 0 {
		return
	}

	if heap.list[i].heatloss < heap.list[parent].heatloss {
		heap.list[i], heap.list[parent] = heap.list[parent], heap.list[i]
		heap.heapifyUp(parent)
	}
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
