package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type ClawMachine struct {
	A, B, Prize aoc.Coordinate
}

func ParseClawMachine(s string) ClawMachine {
	lines := strings.Split(s, "\n")
	if len(lines) != 3 {
		log.Fatal("more than 3 lines in the block")
	}

	r := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)

	m := r.FindStringSubmatch(lines[0])
	x, _ := strconv.Atoi(m[1])
	y, _ := strconv.Atoi(m[2])
	a := aoc.NewCoordinate(x, y)

	m = r.FindStringSubmatch(lines[1])
	x, _ = strconv.Atoi(m[1])
	y, _ = strconv.Atoi(m[2])
	b := aoc.NewCoordinate(x, y)

	r = regexp.MustCompile(`X=(\d+), Y=(\d+)`)
	m = r.FindStringSubmatch(lines[2])
	x, _ = strconv.Atoi(m[1])
	y, _ = strconv.Atoi(m[2])
	p := aoc.NewCoordinate(x, y)

	return ClawMachine{a, b, p}
}

func (cm ClawMachine) CombinationAB() (int, int) {
	a := ((cm.Prize.Y * cm.B.X) - (cm.Prize.X * cm.B.Y)) / ((cm.A.Y * cm.B.X) - (cm.A.X * cm.B.Y))
	b := (cm.Prize.X - (a * cm.A.X)) / cm.B.X

	if (a*cm.A.X+b*cm.B.X != cm.Prize.X) || (a*cm.A.Y+b*cm.B.Y != cm.Prize.Y) {
		return -1, -1
	}

	return a, b
}

type Solver struct {
	ClawMachines []ClawMachine
}

func NewSolver(input string) *Solver {
	cms := []ClawMachine{}

	for _, block := range strings.Split(input, "\n\n") {
		cms = append(cms, ParseClawMachine(block))
	}

	s := Solver{cms}
	return &s
}

func (s *Solver) Part1() int {
	tokens := 0

	for _, cm := range s.ClawMachines {
		a, b := cm.CombinationAB()
		if a < 0 || b < 0 {
			continue
		}

		tokens += (a * 3) + b
	}

	return tokens
}

func (s *Solver) Part2() int {
	tokens := 0

	for _, cm := range s.ClawMachines {
		tmp := cm

		tmp.Prize.X += 10000000000000
		tmp.Prize.Y += 10000000000000
		a, b := tmp.CombinationAB()

		if a < 0 || b < 0 {
			continue
		}

		tokens += (a * 3) + b
	}

	return tokens
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
