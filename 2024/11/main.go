package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Stone struct {
	Value int
}

func (s Stone) Next() []Stone {
	if s.Value == 0 {
		return []Stone{{1}}
	}

	length := int(math.Log10(float64(s.Value)) + 1)
	if length%2 == 0 {
		half := int(math.Pow10(length / 2))
		return []Stone{{s.Value / half}, {s.Value % half}}
	}

	return []Stone{{s.Value * 2024}}
}

type Solver struct {
	Stones []int
}

func NewSolver(input string) *Solver {
	stones := []int{}

	for _, field := range strings.Fields(input) {
		v, _ := strconv.Atoi(field)
		stones = append(stones, v)
	}

	s := Solver{stones}
	return &s
}

func (s *Solver) StoneCountAfterBlink(blinks int) int {
	sc := make(map[int]int)
	for _, stone := range s.Stones {
		if _, ok := sc[stone]; !ok {
			sc[stone] = 0
		}
		sc[stone]++
	}

	for i := 0; i < blinks; i++ {
		next := make(map[int]int)

		for stone, count := range sc {
			length := int(math.Log10(float64(stone)) + 1)

			switch {
			case stone == 0:
				if _, ok := next[1]; !ok {
					next[1] = 0
				}
				next[1] += count

			case length%2 == 0:
				half := int(math.Pow10(length / 2))

				v := stone / half
				if _, ok := next[v]; !ok {
					next[v] = 0
				}
				next[v] += count

				v = stone % half
				if _, ok := next[v]; !ok {
					next[v] = 0
				}
				next[v] += count

			default:
				v := stone * 2024
				if _, ok := next[v]; !ok {
					next[v] = 0
				}
				next[v] += count
			}
		}

		sc = next
	}

	sum := 0
	for _, v := range sc {
		sum += v
	}

	return sum
}

func (s *Solver) Part1() int {
	return s.StoneCountAfterBlink(25)
}

func (s *Solver) Part2() int {
	return s.StoneCountAfterBlink(75)
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
