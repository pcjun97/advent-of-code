package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type CompareType int

const (
	NA CompareType = iota
	LessThan
	GreaterThan
)

type Workflow struct {
	label string
	rules []Rule
}

func parseWorkflow(s string) Workflow {
	r := regexp.MustCompile(`(.+){(.+)}`)
	m := r.FindStringSubmatch(s)
	label := m[1]
	rules := []Rule{}
	for _, r := range strings.Split(m[2], ",") {
		rules = append(rules, parseRule(r))
	}
	return Workflow{label, rules}
}

func (w Workflow) Evaluate(p Part) string {
	for _, r := range w.rules {
		if r.Match(p) {
			return r.next
		}
	}

	return ""
}

type Rule struct {
	compareType CompareType
	value       int
	attr        string
	next        string
}

func parseRule(s string) Rule {
	if !strings.Contains(s, ":") {
		r := Rule{
			compareType: NA,
			next:        s,
		}
		return r
	}

	r := regexp.MustCompile(`(.+)(<|>)(.+):(.+)`)
	m := r.FindStringSubmatch(s)
	v, _ := strconv.Atoi(m[3])
	var compareType CompareType
	switch m[2] {
	case "<":
		compareType = LessThan
	case ">":
		compareType = GreaterThan
	}
	rule := Rule{
		compareType: compareType,
		value:       v,
		attr:        m[1],
		next:        m[4],
	}
	return rule
}

func (r Rule) Match(p Part) bool {
	if r.compareType == NA {
		return true
	}

	var v int
	switch r.attr {
	case "x":
		v = p.x
	case "m":
		v = p.m
	case "a":
		v = p.a
	case "s":
		v = p.s
	}

	if r.compareType == GreaterThan {
		return v > r.value
	} else {
		return v < r.value
	}
}

type Part struct {
	x, m, a, s int
}

func parsePart(s string) Part {
	attrs := make(map[string]int)
	attrs["x"] = 0
	attrs["m"] = 0
	attrs["a"] = 0
	attrs["s"] = 0

	r := regexp.MustCompile(`(.)=(.+)`)
	for _, attr := range strings.Split(s[1:len(s)-1], ",") {
		m := r.FindStringSubmatch(attr)
		v, _ := strconv.Atoi(m[2])
		attrs[m[1]] = v
	}

	return Part{attrs["x"], attrs["m"], attrs["a"], attrs["s"]}
}

func (p Part) Total() int {
	return p.x + p.m + p.a + p.s
}

type Solver struct {
	workflows map[string]Workflow
	parts     []Part
}

func NewSolver(input string) *Solver {
	workflows := make(map[string]Workflow)
	parts := []Part{}

	groups := strings.Split(input, "\n\n")

	for _, line := range strings.Split(groups[0], "\n") {
		w := parseWorkflow(line)
		workflows[w.label] = w
	}

	for _, line := range strings.Split(groups[1], "\n") {
		parts = append(parts, parsePart(line))
	}

	s := Solver{workflows, parts}
	return &s
}

func (s *Solver) Part1() int {
	sum := 0

	for _, p := range s.parts {
		if s.Evaluate(p) {
			sum += p.Total()
		}
	}

	return sum
}

func (s *Solver) Part2() int {
	defaultRange := aoc.NewRange(1, 4000)
	initial := Item{
		xrange: defaultRange,
		mrange: defaultRange,
		arange: defaultRange,
		srange: defaultRange,
		next:   "in",
	}

	cache := make(map[Item]struct{})
	list := []Item{initial}
	accepted := []Item{}
	for len(list) > 0 {
		item := list[0]
		list = list[1:]

		if _, ok := cache[item]; ok {
			continue
		}
		cache[item] = struct{}{}

		if item.next == "R" {
			continue
		}

		if item.next == "A" {
			accepted = append(accepted, item)
			continue
		}

		w := s.workflows[item.next]
		for _, rule := range w.rules {
			i := item
			i.next = rule.next

			switch rule.compareType {
			case LessThan:
				switch rule.attr {
				case "x":
					i.xrange.Max = rule.value - 1
					item.xrange.Min = rule.value
				case "m":
					i.mrange.Max = rule.value - 1
					item.mrange.Min = rule.value
				case "a":
					i.arange.Max = rule.value - 1
					item.arange.Min = rule.value
				case "s":
					i.srange.Max = rule.value - 1
					item.srange.Min = rule.value
				}
			case GreaterThan:
				switch rule.attr {
				case "x":
					i.xrange.Min = rule.value + 1
					item.xrange.Max = rule.value
				case "m":
					i.mrange.Min = rule.value + 1
					item.mrange.Max = rule.value
				case "a":
					i.arange.Min = rule.value + 1
					item.arange.Max = rule.value
				case "s":
					i.srange.Min = rule.value + 1
					item.srange.Max = rule.value
				}
			}

			if i.Valid() {
				list = append(list, i)
			}
			if !item.Valid() {
				break
			}
		}
	}

	sum := 0
	for _, item := range accepted {
		xmas := 1
		xmas *= item.xrange.Max - item.xrange.Min + 1
		xmas *= item.mrange.Max - item.mrange.Min + 1
		xmas *= item.arange.Max - item.arange.Min + 1
		xmas *= item.srange.Max - item.srange.Min + 1
		sum += xmas
	}
	return sum
}

func (s *Solver) Evaluate(p Part) bool {
	next := "in"
	for next != "A" && next != "R" {
		w := s.workflows[next]
		next = w.Evaluate(p)
	}

	if next == "A" {
		return true
	}
	return false
}

type Item struct {
	xrange aoc.Range
	mrange aoc.Range
	arange aoc.Range
	srange aoc.Range
	next   string
}

func (i Item) Valid() bool {
	if i.xrange.Max < i.xrange.Min {
		return false
	}
	if i.mrange.Max < i.mrange.Min {
		return false
	}
	if i.arange.Max < i.arange.Min {
		return false
	}
	if i.srange.Max < i.srange.Min {
		return false
	}

	return true
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
