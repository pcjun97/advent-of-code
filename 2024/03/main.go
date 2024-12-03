package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Solver struct {
	Code string
}

func NewSolver(input string) *Solver {

	s := Solver{input}
	return &s
}

func (s *Solver) Part1() int {
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	sum := 0
	code := s.Code

	for {
		loc := r.FindStringIndex(code)
		if loc == nil {
			break
		}

		mulstr := code[loc[0]:loc[1]]
		code = code[loc[1]:]

		m := r.FindStringSubmatch(mulstr)
		v1, err := strconv.Atoi(m[1])
		if err != nil {
			log.Fatal(err)
		}

		v2, err := strconv.Atoi(m[2])
		if err != nil {
			log.Fatal(err)
		}

		sum += v1 * v2
	}

	return sum
}

func (s *Solver) Part2() int {
	mul := true

	r := regexp.MustCompile(`(mul\(\d+,\d+\))|(do\(\))|(don't\(\))`)
	rmul := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	sum := 0
	code := s.Code

	for {
		loc := r.FindStringIndex(code)
		if loc == nil {
			break
		}

		f := code[loc[0]:loc[1]]
		code = code[loc[1]:]

		switch {
		case f == "do()":
			mul = true

		case f == "don't()":
			mul = false

		case mul:
			m := rmul.FindStringSubmatch(f)
			v1, err := strconv.Atoi(m[1])
			if err != nil {
				log.Fatal(err)
			}

			v2, err := strconv.Atoi(m[2])
			if err != nil {
				log.Fatal(err)
			}
			sum += v1 * v2
		}
	}

	return sum
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
