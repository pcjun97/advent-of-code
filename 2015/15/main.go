package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func parseIngredient(s string) Ingredient {
	r := regexp.MustCompile(`(.*): capacity (.*), durability (.*), flavor (.*), texture (.*), calories (.*)`)
	m := r.FindStringSubmatch(s)
	name := m[1]
	capacity, _ := strconv.Atoi(m[2])
	durability, _ := strconv.Atoi(m[3])
	flavor, _ := strconv.Atoi(m[4])
	texture, _ := strconv.Atoi(m[5])
	calories, _ := strconv.Atoi(m[6])
	ing := Ingredient{name, capacity, durability, flavor, texture, calories}
	return ing
}

type Solver struct {
	ingredients []Ingredient
}

func NewSolver(input string) *Solver {
	ingredients := []Ingredient{}
	for _, line := range strings.Split(input, "\n") {
		ingredients = append(ingredients, parseIngredient(line))
	}
	s := Solver{ingredients}
	return &s
}

func (s *Solver) Part1() int {
	com := generateCombinations(100, len(s.ingredients), nil)

	max := math.MinInt
	for _, c := range com {
		score := s.CookieScore(c)
		if score > max {
			max = score
		}
	}

	return max
}

func (s *Solver) Part2() int {
	com := generateCombinations(100, len(s.ingredients), nil)

	max := math.MinInt
	for _, c := range com {
		score := s.CookieScore(c)
		calories := s.CookieCalories(c)
		if calories == 500 && score > max {
			max = score
		}
	}

	return max
}

func (s *Solver) CookieScore(count []int) int {
	capacity := 0
	for i, ing := range s.ingredients {
		capacity += count[i] * ing.capacity
	}
	if capacity < 0 {
		capacity = 0
	}

	durability := 0
	for i, ing := range s.ingredients {
		durability += count[i] * ing.durability
	}
	if durability < 0 {
		durability = 0
	}

	flavor := 0
	for i, ing := range s.ingredients {
		flavor += count[i] * ing.flavor
	}
	if flavor < 0 {
		flavor = 0
	}

	texture := 0
	for i, ing := range s.ingredients {
		texture += count[i] * ing.texture
	}
	if texture < 0 {
		texture = 0
	}

	return capacity * durability * flavor * texture
}

func (s *Solver) CookieCalories(count []int) int {
	calories := 0
	for i, ing := range s.ingredients {
		calories += count[i] * ing.calories
	}
	if calories < 0 {
		calories = 0
	}
	return calories
}

func generateCombinations(left int, n int, current []int) [][]int {
	if left == 0 {
		for len(current) < n {
			current = append(current, 0)
		}
		return [][]int{current}
	}

	if len(current) == n-1 {
		return [][]int{append(current, left)}
	}

	com := [][]int{}
	for i := 0; i <= left; i++ {
		next := make([]int, len(current))
		copy(next, current)
		next = append(next, i)
		com = append(com, generateCombinations(left-i, n, next)...)
	}

	return com
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
