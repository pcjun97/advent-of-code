package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type City struct {
	id        int
	name      string
	distances map[string]int
}

type Solver struct {
	cities map[string]*City
}

func parseCities(s string) map[string]*City {
	cities := make(map[string]*City)

	r := regexp.MustCompile(`(.*) to (.*) = (.*)`)
	for _, line := range strings.Split(s, "\n") {
		m := r.FindStringSubmatch(line)
		d, _ := strconv.Atoi(m[3])

		c1Name := m[1]
		c2Name := m[2]

		c1, ok := cities[c1Name]
		if !ok {
			c1 = &City{
				name:      c1Name,
				distances: make(map[string]int),
			}
			cities[c1Name] = c1
		}

		c2, ok := cities[c2Name]
		if !ok {
			c2 = &City{
				name:      c2Name,
				distances: make(map[string]int),
			}
			cities[c2Name] = c2
		}

		c1.distances[c2.name] = d
		c2.distances[c1.name] = d
	}

	id := 1
	for _, c := range cities {
		c.id = id
		id += 1
	}

	return cities
}

func NewSolver(input string) *Solver {
	cities := parseCities(input)
	s := Solver{cities}
	return &s
}

func (s *Solver) Part1() int {
	min := s.HeldKarp()
	return min
}

func (s *Solver) Part2() int {
	max := s.HeldKarpMax()
	return max
}

func (s *Solver) HeldKarp() int {
	cache := make(map[CacheKey]int)
	cities := []*City{}
	for _, c := range s.cities {
		cities = append(cities, c)
		for last, d := range c.distances {
			p := s.cities[last]
			key := CacheKey{
				hash: getCacheKeyHash([]*City{c, p}),
				last: p.id,
			}
			cache[key] = d
		}
	}

	for i := 3; i <= len(cities); i++ {
		for _, subset := range getSubsets(cities, i, nil) {
			for j, c := range subset {
				others := []*City{}
				for k, cc := range subset {
					if k == j {
						continue
					}
					others = append(others, cc)
				}
				hash := getCacheKeyHash(others)
				min := -1
				for _, last := range others {
					key := CacheKey{hash, last.id}
					d := cache[key] + c.distances[last.name]
					if min < 0 || d < min {
						min = d
					}
				}
				key := CacheKey{
					hash: getCacheKeyHash(subset),
					last: c.id,
				}
				cache[key] = min
			}
		}
	}

	min := -1
	hash := getCacheKeyHash(cities)
	for _, c := range cities {
		key := CacheKey{hash, c.id}
		d := cache[key]
		if _, ok := cache[key]; !ok {
			log.Fatal(key)
		}
		if min < 0 || d < min {
			min = d
		}
	}

	return min
}

func (s *Solver) HeldKarpMax() int {
	cache := make(map[CacheKey]int)
	cities := []*City{}
	for _, c := range s.cities {
		cities = append(cities, c)
		for last, d := range c.distances {
			p := s.cities[last]
			key := CacheKey{
				hash: getCacheKeyHash([]*City{c, p}),
				last: p.id,
			}
			cache[key] = d
		}
	}

	for i := 3; i <= len(cities); i++ {
		for _, subset := range getSubsets(cities, i, nil) {
			for j, c := range subset {
				others := []*City{}
				for k, cc := range subset {
					if k == j {
						continue
					}
					others = append(others, cc)
				}
				hash := getCacheKeyHash(others)
				max := 0
				for _, last := range others {
					key := CacheKey{hash, last.id}
					d := cache[key] + c.distances[last.name]
					if d > max {
						max = d
					}
				}
				key := CacheKey{
					hash: getCacheKeyHash(subset),
					last: c.id,
				}
				cache[key] = max
			}
		}
	}

	max := 0
	hash := getCacheKeyHash(cities)
	for _, c := range cities {
		key := CacheKey{hash, c.id}
		d := cache[key]
		if _, ok := cache[key]; !ok {
			log.Fatal(key)
		}
		if d > max {
			max = d
		}
	}

	return max
}

func getSubsets(cities []*City, n int, subset []*City) [][]*City {
	if len(subset) == n {
		return [][]*City{subset}
	}
	if len(cities)+len(subset) < n {
		return nil
	}
	subsets := [][]*City{}
	for i, c := range cities {
		ss := make([]*City, len(subset))
		copy(ss, subset)
		ss = append(ss, c)
		subsets = append(subsets, getSubsets(cities[i+1:], n, ss)...)
		subsets = append(subsets, getSubsets(cities[i+1:], n, subset)...)
	}
	return subsets
}

type CacheKey struct {
	hash int
	last int
}

func getCacheKeyHash(cities []*City) int {
	key := 0
	for _, c := range cities {
		key |= int(math.Pow(2, float64(c.id)))
	}
	return key
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
