package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Solver struct {
	listleft  []int
	listright []int
}

func NewSolver(input string) *Solver {
	listleft := []int{}
	listright := []int{}
	r := regexp.MustCompile(`^(\d+)\s+(\d+)$`)

	for _, line := range strings.Split(input, "\n") {
		m := r.FindStringSubmatch(line)

		l1, err := strconv.Atoi(m[1])
		if err != nil {
			log.Fatal(err)
		}

		l2, err := strconv.Atoi(m[2])
		if err != nil {
			log.Fatal(err)
		}

		listleft = append(listleft, l1)
		listright = append(listright, l2)
	}

	s := Solver{listleft, listright}
	return &s
}

func (s *Solver) Part1() int {
	ls1 := []int{}
	ls1 = append(ls1, s.listleft...)

	ls2 := []int{}
	ls2 = append(ls2, s.listright...)

	sort.Ints(ls1)
	sort.Ints(ls2)

	sum := 0
	for i := range ls1 {
		diff := ls1[i] - ls2[i]
		if diff < 0 {
			diff *= -1
		}
		sum += diff
	}

	return sum
}

func (s *Solver) Part2() int {
	m := make(map[int]int)
	for _, v := range s.listright {
		if _, ok := m[v]; !ok {
			m[v] = 0
		}
		m[v]++
	}

	sum := 0
	for _, v := range s.listleft {
		sum += v * m[v]
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
