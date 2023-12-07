package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Solver struct {
	n int
}

func NewSolver(input string) *Solver {
	n, _ := strconv.Atoi(input)
	s := Solver{n}
	return &s
}

func (s *Solver) Part1() int {
	houses := make([]int, s.n/10)
	for i := 1; i <= s.n/10; i++ {
		for j := i; j <= len(houses); j += i {
			houses[j-1] += i * 10
		}
	}
	for i, h := range houses {
		if h >= s.n {
			return i + 1
		}
	}
	return 0
}

func (s *Solver) Part2() int {
	houses := make([]int, s.n/10)
	for i := 1; i <= s.n/10; i++ {
		for k := 1; k <= 50; k++ {
			j := (i * k) - 1
			if j >= len(houses) {
				break
			}
			houses[j] += i * 11
		}
	}
	for i, h := range houses {
		if h >= s.n {
			return i + 1
		}
	}
	return 0
}

func sumOfFactors(n int) int {
	sum := 0

	max := int(math.Sqrt(float64(n)))
	for i := 1; i <= max; i++ {
		if n%i != 0 {
			continue
		}
		sum += i

		if (n / i) != i {
			sum += n / i
		}
	}

	return sum
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
