package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Reindeer struct {
	name  string
	speed int
	fly   int
	rest  int
}

func parseReindeer(s string) Reindeer {
	r := regexp.MustCompile(`(.*) can fly (.*) km/s for (.*) seconds, but then must rest for (.*) seconds.`)
	m := r.FindStringSubmatch(s)
	name := m[1]
	speed, _ := strconv.Atoi(m[2])
	fly, _ := strconv.Atoi(m[3])
	rest, _ := strconv.Atoi(m[4])

	reindeer := Reindeer{name, speed, fly, rest}
	return reindeer
}

func (r Reindeer) Distance(t int) int {
	cycle := r.fly + r.rest
	count := t / cycle
	d := count * r.speed * r.fly

	remaining := t % cycle
	if remaining > r.fly {
		remaining = r.fly
	}

	d += remaining * r.speed
	return d
}

type Solver struct {
	reindeers []Reindeer
}

func NewSolver(input string) *Solver {
	reindeers := []Reindeer{}
	for _, line := range strings.Split(input, "\n") {
		reindeers = append(reindeers, parseReindeer(line))
	}
	s := Solver{reindeers}
	return &s
}

func (s *Solver) Part1(t int) int {
	max := math.MinInt
	for _, r := range s.reindeers {
		d := r.Distance(t)
		if d > max {
			max = d
		}
	}
	return max
}

func (s *Solver) Part2(t int) int {
	score := make(map[Reindeer]int)
	for _, r := range s.reindeers {
		score[r] = 0
	}

	for i := 1; i <= t; i++ {
		distances := make(map[Reindeer]int)
		max := math.MinInt
		for _, r := range s.reindeers {
			d := r.Distance(i)
			distances[r] = d
			if d > max {
				max = d
			}
		}
		for r, d := range distances {
			if d == max {
				score[r] += 1
			}
		}
	}

	max := math.MinInt
	for _, score := range score {
		if score > max {
			max = score
		}
	}

	return max
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1(2503))
	fmt.Println(s.Part2(2503))
}
