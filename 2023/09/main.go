package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func parseList(s string) []int {
	list := []int{}
	r := regexp.MustCompile(`\S+`)
	for _, m := range r.FindAllString(s, -1) {
		n, _ := strconv.Atoi(m)
		list = append(list, n)
	}
	return list
}

type Solver struct {
	lists [][]int
}

func NewSolver(input string) *Solver {
	lists := [][]int{}
	for _, line := range strings.Split(input, "\n") {
		lists = append(lists, parseList(line))
	}
	s := Solver{lists}
	return &s
}

func (s *Solver) Part1() int {
	sum := 0
	for _, list := range s.lists {
		sum += NextValue(list)
	}
	return sum
}

func (s *Solver) Part2() int {
	sum := 0
	for _, list := range s.lists {
		sum += PreviousValue(list)
	}
	return sum
}

func NextValue(list []int) int {
	if ListEqual(list) {
		return list[0]
	}

	next := []int{}
	for i := 0; i < len(list)-1; i++ {
		next = append(next, list[i+1]-list[i])
	}
	return list[len(list)-1] + NextValue(next)
}

func PreviousValue(list []int) int {
	if ListEqual(list) {
		return list[0]
	}

	next := []int{}
	for i := 0; i < len(list)-1; i++ {
		next = append(next, list[i+1]-list[i])
	}
	return list[0] - PreviousValue(next)
}

func ListEqual(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i] != list[i+1] {
			return false
		}
	}
	return true
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
