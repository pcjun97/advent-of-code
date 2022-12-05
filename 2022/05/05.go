package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func main() {
	input := aoc.ReadInput()

	crane9000 := NewCrane9000(input)
	crane9000.Run()
	fmt.Println(string(crane9000.TopCrates()))

	crane9001 := NewCrane9001(input)
	crane9001.Run()
	fmt.Println(string(crane9001.TopCrates()))
}

type Crate rune

type Stack []Crate

func (s *Stack) Peek() Crate {
	stack := *s
	return stack[len(stack)-1]
}

func (s *Stack) Push(c []Crate) {
	*s = append(*s, c...)
}

func (s *Stack) Pop(count int) []Crate {
	stack := *s
	crates := stack[len(stack)-count:]
	*s = stack[:len(stack)-count]
	return crates
}

type Procedure struct {
	count int
	from  int
	to    int
}

func NewProcedure(s string) Procedure {
	fields := strings.Fields(s)

	count, err := strconv.Atoi(fields[1])
	if err != nil {
		panic("error parsing string to int")
	}

	from, err := strconv.Atoi(fields[3])
	if err != nil {
		panic("error parsing string to int")
	}

	to, err := strconv.Atoi(fields[5])
	if err != nil {
		panic("error parsing string to int")
	}

	p := Procedure{
		count: count,
		from:  from - 1,
		to:    to - 1,
	}

	return p
}

type Crane struct {
	stacks     []*Stack
	procedures []Procedure
}

func NewCrane(input string) *Crane {
	lines := strings.Split(input, "\n")

	var n int
	for i, line := range lines {
		if len(line) == 0 {
			n = i
			break
		}
	}

	p1 := lines[:n]
	p2 := lines[n+1:]

	numberOfStacks := len(strings.Fields(p1[len(p1)-1]))

	stacks := make([]*Stack, numberOfStacks)
	for i := 0; i < numberOfStacks; i++ {
		stacks[i] = &Stack{}
	}

	for i := len(p1) - 2; i >= 0; i-- {
		line := p1[i]
		n = 0
		for j := 1; j < len(line); j += 4 {
			if line[j] != ' ' {
				crate := Crate(line[j])
				*stacks[n] = append(*stacks[n], crate)
			}
			n++
		}
	}

	procedures := []Procedure{}
	for _, line := range p2 {
		p := NewProcedure(line)
		procedures = append(procedures, p)
	}

	c := Crane{
		stacks:     stacks,
		procedures: procedures,
	}

	return &c
}

func (c *Crane) TopCrates() []Crate {
	crates := make([]Crate, len(c.stacks))
	for i, s := range c.stacks {
		crates[i] = s.Peek()
	}
	return crates
}

type Crane9000 struct {
	Crane
}

func NewCrane9000(input string) *Crane9000 {
	c := Crane9000{
		Crane: *NewCrane(input),
	}
	return &c
}

func (c *Crane9000) Run() {
	for _, p := range c.procedures {
		for i := 0; i < p.count; i++ {
			crate := c.stacks[p.from].Pop(1)
			c.stacks[p.to].Push(crate)
		}
	}
}

type Crane9001 struct {
	Crane
}

func NewCrane9001(input string) *Crane9001 {
	c := Crane9001{
		Crane: *NewCrane(input),
	}
	return &c
}

func (c *Crane9001) Run() {
	for _, p := range c.procedures {
		crates := c.stacks[p.from].Pop(p.count)
		c.stacks[p.to].Push(crates)
	}
}
