package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type SpringCondition int

const (
	Unknown SpringCondition = iota
	Operational
	Damaged
)

func parseSpringConditions(s string) []SpringCondition {
	sc := []SpringCondition{}
	for _, r := range []byte(s) {
		switch r {
		case '?':
			sc = append(sc, Unknown)
		case '.':
			sc = append(sc, Operational)
		case '#':
			sc = append(sc, Damaged)
		}
	}
	return sc
}

type SpringRow struct {
	conditions []SpringCondition
	records    []int
}

func parseSringRow(s string) SpringRow {
	r := regexp.MustCompile(`(.*) (.*)`)
	m := r.FindStringSubmatch(s)
	conditions := parseSpringConditions(m[1])
	records := parseRecords(m[2])
	row := SpringRow{conditions, records}
	return row
}

func parseRecords(s string) []int {
	records := []int{}
	for _, ss := range strings.Split(s, ",") {
		r, _ := strconv.Atoi(ss)
		records = append(records, r)
	}
	return records
}

type Solver struct {
	rows []SpringRow
}

func NewSolver(input string) *Solver {
	rows := []SpringRow{}
	for _, line := range strings.Split(input, "\n") {
		rows = append(rows, parseSringRow(line))
	}
	s := Solver{rows}
	return &s
}

func (s *Solver) Part1() int {
	sum := 0
	for _, row := range s.rows {
		sum += countPossibleArragements(row.records, row.conditions)
	}
	return sum
}

func (s *Solver) Part2() int {
	sum := 0
	for _, row := range s.rows {
		records := []int{}
		for i := 0; i < 5; i++ {
			records = append(records, row.records...)
		}

		conditions := []SpringCondition{}
		conditions = append(conditions, row.conditions...)
		for i := 0; i < 4; i++ {
			conditions = append(conditions, Unknown)
			conditions = append(conditions, row.conditions...)
		}

		sum += countPossibleArragements(records, conditions)
	}
	return sum
}

func countPossibleArragements(records []int, conditions []SpringCondition) int {
	conditions = append([]SpringCondition{Operational}, conditions...)

	type key struct {
		sum    int
		length int
	}
	cache := make(map[key]int)

	sum := func(records []int) int {
		s := 0
		for _, r := range records {
			s += r
		}
		return s
	}

	var count func([]int, []SpringCondition) int
	count = func(records []int, conditions []SpringCondition) int {
		k := key{sum(records), len(conditions)}
		if v, ok := cache[k]; ok {
			return v
		}

		if len(records) == 0 {
			if allOperational(conditions) {
				cache[k] = 1
				return 1
			} else {
				cache[k] = 0
				return 0
			}
		}

		if len(conditions) <= 1 {
			cache[k] = 0
			return 0
		}

		if conditions[0] == Damaged {
			cache[k] = 0
			return 0
		}

		n := 0
		switch conditions[1] {
		case Operational:
			n = count(records, conditions[1:])
		case Damaged:
			if firstNDamaged(conditions[1:], records[0]) {
				n = count(records[1:], conditions[records[0]+1:])
			}
		case Unknown:
			n = count(records, conditions[1:])
			if firstNDamaged(conditions[1:], records[0]) {
				n += count(records[1:], conditions[records[0]+1:])
			}
		}

		cache[k] = n
		return n
	}

	return count(records, conditions)
}

func allOperational(conditions []SpringCondition) bool {
	for _, c := range conditions {
		if c == Damaged {
			return false
		}
	}
	return true
}

func firstNDamaged(conditions []SpringCondition, n int) bool {
	if len(conditions) < n {
		return false
	}

	for i := 0; i < n; i++ {
		if conditions[i] == Operational {
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
