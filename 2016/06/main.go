package main

import (
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Solver struct {
	messages []string
}

func NewSolver(input string) *Solver {
	s := Solver{strings.Split(input, "\n")}
	return &s
}

func (s *Solver) Part1() string {
	m := ""

	if len(s.messages) <= 0 {
		log.Fatal("no message parsed")
	}

	for i := 0; i < len(s.messages[0]); i++ {
		cm := make(map[byte]int)

		for _, message := range s.messages {
			c := message[i]
			if _, ok := cm[c]; !ok {
				cm[c] = 0
			}
			cm[c]++
		}

		var maxc byte
		max := 0
		for c, x := range cm {
			if x > max {
				max = x
				maxc = c
			}
		}

		m = fmt.Sprintf("%s%c", m, maxc)
	}

	return m
}

func (s *Solver) Part2() string {
	m := ""

	if len(s.messages) <= 0 {
		log.Fatal("no message parsed")
	}

	for i := 0; i < len(s.messages[0]); i++ {
		cm := make(map[byte]int)

		for _, message := range s.messages {
			c := message[i]
			if _, ok := cm[c]; !ok {
				cm[c] = 0
			}
			cm[c]++
		}

		var minc byte
		min := math.MaxInt
		for c, x := range cm {
			if x < min {
				min = x
				minc = c
			}
		}

		m = fmt.Sprintf("%s%c", m, minc)
	}

	return m
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
