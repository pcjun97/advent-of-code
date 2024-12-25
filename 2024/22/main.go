package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Solver struct {
	SecretNumbers []int
}

func NewSolver(input string) *Solver {
	numbers := []int{}

	for _, line := range strings.Split(input, "\n") {
		n, _ := strconv.Atoi(line)
		numbers = append(numbers, n)
	}

	s := Solver{numbers}
	return &s
}

func (s *Solver) Part1() int {
	sum := 0

	for _, n := range s.SecretNumbers {
		sum += NthSecretNumber(n, 2000)
	}

	return sum
}

func (s *Solver) Part2() int {
	sumbananas := make(map[[4]int]int)

	for _, num := range s.SecretNumbers {
		n := num
		bananas := make(map[[4]int]int)
		diffs := []int{}

		for i := 0; i < 2000; i++ {
			next := NextSecretNumber(n)

			diff := (next % 10) - (n % 10)
			diffs = append(diffs, diff)

			n = next

			if len(diffs) > 4 {
				diffs = diffs[1:]
			}

			if len(diffs) == 4 {
				key := [4]int{
					diffs[0],
					diffs[1],
					diffs[2],
					diffs[3],
				}
				if _, ok := bananas[key]; !ok {
					bananas[key] = n % 10
				}
			}
		}

		for diffs, v := range bananas {
			if _, ok := sumbananas[diffs]; !ok {
				sumbananas[diffs] = 0
			}
			sumbananas[diffs] += v
		}
	}

	maxbananas := 0
	for _, v := range sumbananas {
		if v > maxbananas {
			maxbananas = v
		}
	}
	return maxbananas
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}

func NthSecretNumber(initial, n int) int {
	v := initial
	for i := 0; i < n; i++ {
		v = NextSecretNumber(v)
	}
	return v
}

func NextSecretNumber(n int) int {
	n = Prune(Mix(n, n*64))
	n = Prune(Mix(n, n/32))
	n = Prune(Mix(n, n*2048))
	return n
}

func Mix(a, b int) int {
	return a ^ b
}

func Prune(n int) int {
	return n % 16777216
}
