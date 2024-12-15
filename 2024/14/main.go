package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Robot struct {
	Position, Velocity aoc.Coordinate
}

func ParseRobot(s string) Robot {
	r := regexp.MustCompile(`^p=(\d+),(\d+)\s*v=(-?\d+),(-?\d+)$`)
	m := r.FindStringSubmatch(s)
	px, _ := strconv.Atoi(m[1])
	py, _ := strconv.Atoi(m[2])
	vx, _ := strconv.Atoi(m[3])
	vy, _ := strconv.Atoi(m[4])

	return Robot{aoc.NewCoordinate(px, py), aoc.NewCoordinate(vx, vy)}
}

func (r Robot) PositionAtSecond(t int) aoc.Coordinate {
	px := r.Position.X + (r.Velocity.X * t)
	py := r.Position.Y + (r.Velocity.Y * t)

	return aoc.NewCoordinate(px, py)
}

type Solver struct {
	Robots        []Robot
	Height, Width int
}

func NewSolver(input string, h, w int) *Solver {
	robots := []Robot{}

	for _, line := range strings.Split(input, "\n") {
		robots = append(robots, ParseRobot(line))
	}

	s := Solver{robots, h, w}
	return &s
}

func (s *Solver) RobotPositionAtSecond(r Robot, t int) aoc.Coordinate {
	p := r.PositionAtSecond(t)

	p.X %= s.Width
	if p.X < 0 {
		p.X += s.Width
	}

	p.Y %= s.Height
	if p.Y < 0 {
		p.Y += s.Height
	}

	return p
}

func (s *Solver) ToStringAtSecond(t int) string {
	output := []rune{}
	for y := 0; y < s.Height; y++ {
		for x := 0; x < s.Width; x++ {
			output = append(output, ' ')
		}
		output = append(output, '\n')
	}
	output = append(output, '\n')

	for _, robot := range s.Robots {
		c := s.RobotPositionAtSecond(robot, t)
		output[(c.Y*(s.Width+1))+c.X] = '#'
	}

	return string(output)
}

func (s *Solver) Part1() int {
	rcs := []aoc.Coordinate{}

	for _, robot := range s.Robots {
		rcs = append(rcs, s.RobotPositionAtSecond(robot, 100))
	}

	quadrants := []int{0, 0, 0, 0}

	mX := s.Width / 2
	mY := s.Height / 2

	for _, rc := range rcs {
		switch {
		case rc.X < mX && rc.Y < mY:
			quadrants[0]++

		case rc.X < mX && rc.Y > mY:
			quadrants[1]++

		case rc.X > mX && rc.Y < mY:
			quadrants[2]++

		case rc.X > mX && rc.Y > mY:
			quadrants[3]++
		}
	}

	product := 1
	for _, q := range quadrants {
		product *= q
	}

	return product
}

func (s *Solver) Part2() {
	i := 6750
	for {
		fmt.Println(s.ToStringAtSecond(i))
		fmt.Println(i)
		time.Sleep(time.Second)
		i++
	}
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input, 103, 101)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	s.Part2()
}
