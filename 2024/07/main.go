package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Equation struct {
	TestValue int
	Numbers   []int
}

func ParseEquation(s string) Equation {
	r := regexp.MustCompile(`^(\d+):\s+(.*)$`)
	m := r.FindStringSubmatch(s)

	tv, _ := strconv.Atoi(m[1])

	numbers := []int{}
	for _, f := range strings.Fields(m[2]) {
		v, err := strconv.Atoi(f)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, v)
	}

	e := Equation{tv, numbers}
	return e
}

func (eq Equation) Validate(operators []func(testvalue, subvalue int) (int, bool)) bool {
	if len(eq.Numbers) == 1 {
		return eq.Numbers[0] == eq.TestValue
	}

	for _, op := range operators {
		v, ok := op(eq.TestValue, eq.Numbers[len(eq.Numbers)-1])
		if !ok {
			continue
		}

		eqq := Equation{v, eq.Numbers[:len(eq.Numbers)-1]}
		if eqq.Validate(operators) {
			return true
		}
	}

	return false
}

type Solver struct {
	Equations []Equation
}

func NewSolver(input string) *Solver {
	eqs := []Equation{}
	for _, line := range strings.Split(input, "\n") {
		eqs = append(eqs, ParseEquation(line))
	}

	s := Solver{eqs}
	return &s
}

func (s *Solver) Part1() int {
	sum := 0
	ops := []func(testvalue, subvalue int) (int, bool){TrySum, TryMul}

	for _, eq := range s.Equations {
		if eq.Validate(ops) {
			sum += eq.TestValue
		}
	}

	return sum
}

func (s *Solver) Part2() int {
	sum := 0
	ops := []func(testvalue, subvalue int) (int, bool){TrySum, TryMul, TryConcat}

	for _, eq := range s.Equations {
		if eq.Validate(ops) {
			sum += eq.TestValue
		}
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

func EvaluatePossibleValues(nums []int, operators []func(a, b int) int, max int) []int {
	if len(nums) <= 1 {
		return nums
	}

	values := make(map[int]struct{})

	for _, op := range operators {
		vv := op(nums[0], nums[1])
		if vv > max {
			continue
		}

		for _, v := range EvaluatePossibleValues(append([]int{vv}, nums[2:]...), operators, max) {
			values[v] = struct{}{}
		}
	}

	valueslist := []int{}
	for v := range values {
		valueslist = append(valueslist, v)
	}
	return valueslist
}

func TrySum(testvalue, subvalue int) (int, bool) {
	if subvalue > testvalue {
		return testvalue, false
	}
	return testvalue - subvalue, true
}

func TryMul(testvalue, subvalue int) (int, bool) {
	if testvalue%subvalue != 0 {
		return testvalue, false
	}

	return testvalue / subvalue, true
}

func TryConcat(testvalue, subvalue int) (int, bool) {
	a := testvalue
	b := subvalue

	for {
		if a%10 != b%10 {
			return testvalue, false
		}

		a /= 10
		b /= 10

		if b == 0 {
			break
		}
	}

	return a, true
}
