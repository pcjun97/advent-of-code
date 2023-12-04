package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Card struct {
	id      int
	numbers []int
	winning map[int]struct{}
}

func (c Card) WinningCount() int {
	count := 0
	for _, n := range c.numbers {
		if _, ok := c.winning[n]; ok {
			count += 1
		}
	}
	return count
}

func parseCard(s string) Card {
	r := regexp.MustCompile(`Card (.*): (.*) \| (.*)`)
	m := r.FindStringSubmatch(s)
	id, _ := strconv.Atoi(m[1])

	numbers := []int{}
	winning := make(map[int]struct{})

	for _, n := range strings.Split(m[2], " ") {
		if len(n) == 0 {
			continue
		}
		x, _ := strconv.Atoi(n)
		winning[x] = struct{}{}
	}

	for _, n := range strings.Split(m[3], " ") {
		if len(n) == 0 {
			continue
		}
		x, _ := strconv.Atoi(n)
		numbers = append(numbers, x)
	}

	c := Card{id, numbers, winning}
	return c
}

type Solver struct {
	cards []Card
}

func NewSolver(input string) *Solver {
	cards := []Card{}
	for _, line := range strings.Split(input, "\n") {
		c := parseCard(line)
		cards = append(cards, c)
	}
	s := Solver{cards}
	return &s
}

func (s *Solver) Part1() int {
	sum := 0
	for _, c := range s.cards {
		count := c.WinningCount()
		if count > 0 {
			sum += int(math.Pow(2, float64(count-1)))
		}
	}

	return sum
}

func (s *Solver) Part2() int {
	cardCounts := make([]int, len(s.cards))
	for i := range cardCounts {
		cardCounts[i] = 1
	}

	for i, c := range s.cards {
		wc := c.WinningCount()
		for j := 1; j <= wc && i+j < len(s.cards); j++ {
			cardCounts[i+j] += cardCounts[i]
		}
	}

	sum := 0
	for _, c := range cardCounts {
		sum += c
	}

	return sum
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
