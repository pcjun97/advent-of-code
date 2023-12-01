package main

import (
	"crypto/md5"
	"fmt"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Miner struct {
	key    string
	zeroes string
}

func NewMiner(key string, difficulty uint) Miner {
	zeroes := ""
	var i uint
	for i = 0; i < difficulty; i++ {
		zeroes += "0"
	}
	m := Miner{key, zeroes}
	return m
}

func (m Miner) Mine() int {
	i := 0
	for {
		data := []byte(fmt.Sprintf("%s%d", m.key, i))
		hash := fmt.Sprintf("%x", md5.Sum(data))
		if hash[:len(m.zeroes)] == m.zeroes {
			return i
		}
		i++
	}
}

type Solver struct {
	key string
}

func NewSolver(input string) *Solver {
	s := Solver{input}
	return &s
}

func (s *Solver) Part1() int {
	m := NewMiner(s.key, 5)
	return m.Mine()
}

func (s *Solver) Part2() int {
	m := NewMiner(s.key, 6)
	return m.Mine()
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
