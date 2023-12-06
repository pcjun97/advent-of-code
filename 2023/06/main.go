package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Race struct {
	time   int
	record int
}

func (r Race) PotentialWinCount() int {
	sum := 0
	for i := 0; i <= r.time; i++ {
		d := (r.time - i) * i
		if d > r.record {
			sum += 1
		}
	}
	return sum
}

func parseRaces(s string) []Race {
	r := regexp.MustCompile(`\d+`)
	lines := strings.Split(s, "\n")

	times := []int{}
	for _, m := range r.FindAllString(lines[0], -1) {
		t, _ := strconv.Atoi(m)
		times = append(times, t)
	}

	distances := []int{}
	for _, m := range r.FindAllString(lines[1], -1) {
		d, _ := strconv.Atoi(m)
		distances = append(distances, d)
	}

	races := []Race{}
	for i, d := range distances {
		r := Race{times[i], d}
		races = append(races, r)
	}

	return races
}

func parseRaceIgnoreWhitespace(s string) Race {
	r1 := regexp.MustCompile(`Time:\s(.*)\nDistance:\s(.*)`)
	m1 := r1.FindStringSubmatch(s)

	r2 := regexp.MustCompile(`\s*`)
	time, _ := strconv.Atoi(r2.ReplaceAllString(m1[1], ""))
	distance, _ := strconv.Atoi(r2.ReplaceAllString(m1[2], ""))
	race := Race{time, distance}
	return race
}

type Solver struct {
	races []Race
	race  Race
}

func NewSolver(input string) *Solver {
	races := parseRaces(input)
	race := parseRaceIgnoreWhitespace(input)
	s := Solver{races, race}
	return &s
}

func (s *Solver) Part1() int {
	margin := 1
	for _, r := range s.races {
		margin *= r.PotentialWinCount()
	}
	return margin
}

func (s *Solver) Part2() int {
	return s.race.PotentialWinCount()
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
