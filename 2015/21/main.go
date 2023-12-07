package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Item struct {
	cost, damage, armor int
}

var weapons = []Item{
	{8, 4, 0},
	{10, 5, 0},
	{25, 6, 0},
	{40, 7, 0},
	{74, 8, 0},
}

var armors = []Item{
	{13, 0, 1},
	{31, 0, 2},
	{53, 0, 3},
	{75, 0, 4},
	{102, 0, 5},
}

var rings = []Item{
	{25, 1, 0},
	{50, 2, 0},
	{100, 3, 0},
	{20, 0, 1},
	{40, 0, 2},
	{80, 0, 3},
}

type Character struct {
	hp, damage, armor int
}

func (c Character) CanBeat(cc Character) bool {
	dmg1 := c.damage - cc.armor
	if dmg1 <= 0 {
		dmg1 = 1
	}

	dmg2 := cc.damage - c.armor
	if dmg2 <= 0 {
		dmg2 = 1
	}

	r1 := c.hp / dmg2
	if c.hp%dmg2 > 0 {
		r1 += 1
	}

	r2 := cc.hp / dmg1
	if cc.hp%dmg1 > 0 {
		r2 += 1
	}

	return r1 >= r2
}

func parseCharacter(s string) Character {
	r := regexp.MustCompile(`Hit Points: (.*)\nDamage: (.*)\nArmor: (.*)$`)
	m := r.FindStringSubmatch(s)
	hp, _ := strconv.Atoi(m[1])
	damage, _ := strconv.Atoi(m[2])
	armor, _ := strconv.Atoi(m[3])
	c := Character{hp, damage, armor}
	return c
}

type Solver struct {
	boss Character
}

func NewSolver(input string) *Solver {
	s := Solver{parseCharacter(input)}
	return &s
}

func (s *Solver) Part1() int {
	sums := []int{}
	for _, items := range s.ItemsCombinations() {
		damage := 0
		armor := 0
		cost := 0
		for _, item := range items {
			damage += item.damage
			armor += item.armor
			cost += item.cost
		}
		c := Character{100, damage, armor}
		if c.CanBeat(s.boss) {
			sums = append(sums, cost)
		}
	}

	min := math.MaxInt
	for _, cost := range sums {
		if cost < min {
			min = cost
		}
	}
	return min
}

func (s *Solver) Part2() int {
	sums := []int{}
	for _, items := range s.ItemsCombinations() {
		damage := 0
		armor := 0
		cost := 0
		for _, item := range items {
			damage += item.damage
			armor += item.armor
			cost += item.cost
		}
		c := Character{100, damage, armor}
		if !c.CanBeat(s.boss) {
			sums = append(sums, cost)
		}
	}

	max := math.MinInt
	for _, cost := range sums {
		if cost > max {
			max = cost
		}
	}
	return max
}

func (s *Solver) ItemsCombinations() [][]Item {
	com := [][]Item{}

	for w := 0; w < len(weapons); w++ {
		items1 := []Item{weapons[w]}
		for a := -1; a < len(armors); a++ {
			items2 := items1
			if a >= 0 {
				items2 = append([]Item{armors[a]}, items2...)
			}

			for r1 := -1; r1 < len(rings); r1++ {
				items3 := items2
				if r1 >= 0 {
					items3 = append([]Item{rings[r1]}, items3...)
				}
				com = append(com, items3)

				for r2 := r1 + 1; r2 < len(rings); r2++ {
					items4 := append([]Item{rings[r2]}, items3...)
					com = append(com, items4)
				}
			}
		}
	}

	return com
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
