package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func main() {
	input := aoc.ReadInput()
	solver := NewSolver(input)

	fmt.Println(solver.TotalQualityLevel(24))

	val := 1
	for _, n := range solver.FirstNMaxGeodes(3, 32) {
		val *= n
	}
	fmt.Println(val)
}

type Resource int

const (
	ORE Resource = iota
	CLAY
	OBSEDIAN
	GEODE
)

type Recipe struct {
	Item        Resource
	Ingredients map[Resource]int
}

type Blueprint struct {
	ID      int
	Recipes map[Resource]Recipe
	max     map[Resource]int
}

func NewBlueprint(s string) Blueprint {
	fields := strings.Split(s, " ")
	id, err := strconv.Atoi(strings.TrimRight(fields[1], ":"))
	if err != nil {
		panic(err)
	}

	recipes := make(map[Resource]Recipe)
	var n1, n2 int

	n1, err = strconv.Atoi(fields[6])
	if err != nil {
		panic(err)
	}
	recipes[ORE] = Recipe{
		Item:        ORE,
		Ingredients: map[Resource]int{ORE: n1},
	}

	n1, err = strconv.Atoi(fields[12])
	if err != nil {
		panic(err)
	}
	recipes[CLAY] = Recipe{
		Item:        CLAY,
		Ingredients: map[Resource]int{ORE: n1},
	}

	n1, err = strconv.Atoi(fields[18])
	if err != nil {
		panic(err)
	}
	n2, err = strconv.Atoi(fields[21])
	if err != nil {
		panic(err)
	}
	recipes[OBSEDIAN] = Recipe{
		Item:        OBSEDIAN,
		Ingredients: map[Resource]int{ORE: n1, CLAY: n2},
	}

	n1, err = strconv.Atoi(fields[27])
	if err != nil {
		panic(err)
	}
	n2, err = strconv.Atoi(fields[30])
	if err != nil {
		panic(err)
	}
	recipes[GEODE] = Recipe{
		Item:        GEODE,
		Ingredients: map[Resource]int{ORE: n1, OBSEDIAN: n2},
	}

	max := make(map[Resource]int)
	for _, recipe := range recipes {
		for res, n := range recipe.Ingredients {
			if _, ok := max[res]; !ok || n > max[res] {
				max[res] = n
			}
		}
	}

	b := Blueprint{
		ID:      id,
		Recipes: recipes,
		max:     max,
	}

	return b
}

func (b Blueprint) Max(r Resource) int {
	if n, ok := b.max[r]; ok {
		return n
	}
	return 0
}

type State struct {
	Robots    map[Resource]int
	Resources map[Resource]int
}

func (s *State) Build(r Recipe) {
	for resource, n := range r.Ingredients {
		s.Resources[resource] -= n
	}

	if _, ok := s.Robots[r.Item]; ok {
		s.Robots[r.Item]++
	} else {
		s.Robots[r.Item] = 1
	}
}

func (s *State) ResourceCount(r Resource) int {
	if n, ok := s.Resources[r]; ok {
		return n
	}
	return 0
}

func (s *State) RobotCount(r Resource) int {
	if n, ok := s.Robots[r]; ok {
		return n
	}
	return 0
}

func (s *State) Copy() State {
	c := State{
		Robots:    make(map[Resource]int),
		Resources: make(map[Resource]int),
	}
	for k, v := range s.Robots {
		c.Robots[k] = v
	}
	for k, v := range s.Resources {
		c.Resources[k] = v
	}
	return c
}

func (s *State) HasRobots(r Recipe) bool {
	for robot := range r.Ingredients {
		if s.RobotCount(robot) == 0 {
			return false
		}
	}
	return true
}

func (s *State) WaitTime(r Recipe) int {
	max := 0
	for resource, count := range r.Ingredients {
		if s.RobotCount(resource) == 0 {
			return -1
		}
		diff := count - s.ResourceCount(resource)
		wait := diff / s.RobotCount(resource)
		if diff%s.RobotCount(resource) > 0 {
			wait += 1
		}
		if wait > max {
			max = wait
		}
	}
	return max
}

func (s *State) Progress(step int) {
	for r, n := range s.Robots {
		if _, ok := s.Resources[r]; ok {
			s.Resources[r] += n * step
		} else {
			s.Resources[r] = n * step
		}
	}
}

type Solver struct {
	blueprints []Blueprint
	cache      map[CacheKey][][4]int
	cacheMax   int
}

func NewSolver(s string) *Solver {
	lines := strings.Split(s, "\n")
	b := make([]Blueprint, len(lines))
	for i, line := range lines {
		b[i] = NewBlueprint(line)
	}
	solver := Solver{
		blueprints: b,
	}
	return &solver
}

func (s *Solver) TotalQualityLevel(limit int) int {
	sum := 0
	for _, b := range s.blueprints {
		sum += s.QualityLevel(b, limit)
	}
	return sum
}

func (s *Solver) FirstNMaxGeodes(n, limit int) []int {
	l := []int{}
	for i := 0; i < n && i < len(s.blueprints); i++ {
		state := State{
			Robots:    make(map[Resource]int),
			Resources: make(map[Resource]int),
		}
		state.Robots[ORE] = 1
		s.cacheMax = 0
		s.cache = make(map[CacheKey][][4]int)
		l = append(l, s.DFS(s.blueprints[i], state, limit, -1, ""))
	}
	return l
}

func (s *Solver) QualityLevel(b Blueprint, limit int) int {
	state := State{
		Robots:    make(map[Resource]int),
		Resources: make(map[Resource]int),
	}
	state.Robots[ORE] = 1
	s.cacheMax = 0
	s.cache = make(map[CacheKey][][4]int)
	return b.ID * s.DFS(b, state, limit, -1, "")
}

func (s *Solver) DFS(b Blueprint, state State, rem int, build Resource, indent string) int {
	state.Progress(1)
	if robot, ok := b.Recipes[build]; ok {
		state.Build(robot)
	}
	rem -= 1

	if hit := s.Cache(state, rem); hit {
		return -1
	}

	if rem == 1 {
		state.Progress(1)
		if state.ResourceCount(GEODE) > s.cacheMax {
			s.cacheMax = state.ResourceCount(GEODE)
		}
		return state.ResourceCount(GEODE)
	}

	// total number of geode if we build geode for every following minutes
	n := state.ResourceCount(GEODE)
	for i := 0; i < rem; i++ {
		n += state.RobotCount(GEODE) + i
	}

	// abandon if cannot be better than cached max
	if n <= s.cacheMax {
		return -1
	}

	// check if enough materials to build geode for every remaining minutes
	sufficient := true
	for r, n := range b.Recipes[GEODE].Ingredients {
		if (state.ResourceCount(r)+(state.RobotCount(r)*(rem-2)))/(rem-1) < n {
			sufficient = false
		}
	}
	if sufficient {
		if n > s.cacheMax {
			s.cacheMax = n
		}
		return n
	}

	max := 0
	if state.HasRobots(b.Recipes[GEODE]) {
		t := state.WaitTime(b.Recipes[GEODE])
		if t >= 0 && rem-t > 1 {
			c := state.Copy()
			c.Progress(t)
			val := s.DFS(b, c, rem-t, GEODE, indent+"| ")
			if val > max {
				max = val
			}
		}
	}

	robots := []Resource{OBSEDIAN, CLAY, ORE}
	for _, robot := range robots {
		if !state.HasRobots(b.Recipes[robot]) {
			continue
		}

		// enough to build any robot every minute
		if (state.ResourceCount(robot)+state.RobotCount(robot)*(rem-2))/(rem-1) >= b.Max(robot) {
			continue
		}

		t := state.WaitTime(b.Recipes[robot])
		if t < 0 || rem-t <= 1 {
			continue
		}

		c := state.Copy()
		c.Progress(t)
		val := s.DFS(b, c, rem-t, robot, indent+"| ")
		if val > max {
			max = val
		}
	}

	return max
}

func (s *Solver) Cache(state State, rem int) bool {
	key := NewCacheKey(state, rem)

	order := [4]Resource{ORE, CLAY, OBSEDIAN, GEODE}
	resources := [4]int{}

	for i, r := range order {
		resources[i] = state.ResourceCount(r)
	}

	if _, ok := s.cache[key]; !ok {
		s.cache[key] = [][4]int{resources}
		return false
	}

	for _, r := range s.cache[key] {
		smaller := true
		for i, n := range r {
			if resources[i] > n {
				smaller = false
			}
		}
		if smaller {
			return true
		}
	}
	s.cache[key] = append(s.cache[key], resources)
	return false
}

type CacheKey struct {
	t      int
	robots [4]int
}

func NewCacheKey(state State, rem int) CacheKey {
	order := [4]Resource{ORE, CLAY, OBSEDIAN, GEODE}

	robots := [4]int{}
	for i, r := range order {
		robots[i] = state.RobotCount(r)
	}

	key := CacheKey{rem, robots}
	return key
}
