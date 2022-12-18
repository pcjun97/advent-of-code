package main

import (
	"fmt"

	"github.com/pcjun97/advent-of-code/aoc"
)

var ROCK_PATTERNS [][]byte = [][]byte{
	{
		0b0011110,
	},
	{
		0b0001000,
		0b0011100,
		0b0001000,
	},
	{
		0b0011100,
		0b0000100,
		0b0000100,
	},
	{
		0b0010000,
		0b0010000,
		0b0010000,
		0b0010000,
	},
	{
		0b0011000,
		0b0011000,
	},
}

func main() {
	input := aoc.ReadInput()

	c := NewChamber(input)

	h := c.DropRocks(2022)
	fmt.Println(h)

	h = c.DropRocks(1000000000000)
	fmt.Println(h)
}

type Chamber struct {
	chamber   []byte
	jet       []Direction
	jetIndex  int
	rockIndex int
	cache     map[[2]int][][2]int
}

func NewChamber(s string) *Chamber {
	jet := []Direction{}
	for _, i := range s {
		if i == '<' {
			jet = append(jet, DIR_LEFT)
		} else if i == '>' {
			jet = append(jet, DIR_RIGHT)
		}
	}

	c := Chamber{
		chamber:   []byte{},
		jet:       jet,
		rockIndex: 0,
		jetIndex:  0,
		cache:     make(map[[2]int][][2]int),
	}
	return &c
}

func (c *Chamber) DropRocks(count int) int {
	c.chamber = []byte{}
	c.rockIndex = 0
	c.jetIndex = 0
	c.cache = make(map[[2]int][][2]int)

	var i int
	var cachedHeight, cachedDropped int
	for i = 0; i < count; i++ {
		ri := c.rockIndex
		ji := c.jetIndex
		c.DropOneRock()
		if cached, ok := c.Cache(ri, ji, len(c.chamber), i+1); ok {
			cachedHeight = cached[0]
			cachedDropped = cached[1]
			i++
			break
		}
	}

	if i == count {
		return len(c.chamber)
	}

	last := len(c.chamber)
	final := last

	h := len(c.chamber) - cachedHeight
	n := i - cachedDropped

	count -= i
	final += (count / n) * h
	count = count % n

	for count > 0 {
		c.DropOneRock()
		count--
	}

	final += len(c.chamber) - last
	return final
}

func (c *Chamber) DropOneRock() {
	for i := 0; i < 7; i++ {
		c.chamber = append(c.chamber, 0)
	}

	rock := Rock(ROCK_PATTERNS[c.rockIndex])
	c.rockIndex = (c.rockIndex + 1) % len(ROCK_PATTERNS)

	var i int
	for i = len(c.chamber) - 4; i >= 0; i-- {
		tmp := make(Rock, len(rock))
		copy(tmp, rock)

		tmp.Move(c.jet[c.jetIndex])
		c.jetIndex = (c.jetIndex + 1) % len(c.jet)

		if !tmp.Collide(c.chamber[i : i+len(tmp)]) {
			rock = tmp
		}

		if i == 0 || rock.Collide(c.chamber[i-1:i-1+len(rock)]) {
			break
		}
	}

	for j := range rock {
		c.chamber[i+j] = c.chamber[i+j] | rock[j]
	}

	if i+len(rock) > len(c.chamber)-7 {
		c.chamber = c.chamber[:i+len(rock)]
	} else {
		c.chamber = c.chamber[:len(c.chamber)-7]
	}
}

func (c *Chamber) Cache(rockIndex, jetIndex, height, dropped int) ([2]int, bool) {
	key := [2]int{rockIndex, jetIndex}
	cached, ok := c.cache[key]
	if !ok {
		c.cache[key] = [][2]int{{height, dropped}}
		return [2]int{}, false
	}
	for i := len(cached) - 1; i > 0; i-- {
		m1 := string(c.chamber[cached[i][0]:height])
		for j := i - 1; j >= 0; j-- {
			m2 := string(c.chamber[cached[j][0]:cached[i][0]])
			if m1 == m2 {
				return cached[i], true
			}
		}
	}
	c.cache[key] = append(c.cache[key], [2]int{height, dropped})
	return [2]int{}, false
}

type Rock []byte

func (r Rock) Move(dir Direction) {
	var edge byte
	switch dir {
	case DIR_RIGHT:
		edge = 1
	case DIR_LEFT:
		edge = 0b1000000
	}

	for i := range r {
		if (r[i] & edge) != 0 {
			return
		}
	}

	for i := range r {
		switch dir {
		case DIR_RIGHT:
			r[i] = r[i] >> 1
		case DIR_LEFT:
			r[i] = r[i] << 1
		}
	}
}

func (r Rock) Collide(chamber []byte) bool {
	for i := range r {
		if (r[i] & chamber[i]) != 0 {
			return true
		}
	}
	return false
}

type Direction bool

const (
	DIR_RIGHT Direction = true
	DIR_LEFT  Direction = false
)
