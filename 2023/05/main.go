package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Range struct {
	s int
	r int
}

func (r Range) InRange(n int) bool {
	return n >= r.s && n < (r.s+r.r)
}

func (r Range) IsOverlap(rr Range) bool {
	if rr.s >= r.s && rr.s < (r.s+r.r) {
		return true
	}
	if r.s >= rr.s && r.s < (rr.s+rr.r) {
		return true
	}
	return false
}

func (r Range) Intersection(rr Range) []Range {
	if !r.IsOverlap(rr) {
		return nil
	}

	result := []Range{}

	minStart := r.s
	maxStart := rr.s
	if maxStart < minStart {
		minStart, maxStart = maxStart, minStart
	}

	minEnd := r.s + r.r
	maxEnd := rr.s + rr.r
	if maxEnd < minEnd {
		minEnd, maxEnd = maxEnd, minEnd
	}

	intersect := Range{maxStart, minEnd - maxStart}
	result = append(result, intersect)

	if r.s < intersect.s {
		result = append(result, Range{r.s, intersect.s - r.s})
	}
	if r.s+r.r > intersect.s+intersect.r {
		s := intersect.s + intersect.r
		result = append(result, Range{s, r.s + r.r - s})
	}

	return result
}

type RangeMapEntry struct {
	r Range
	t int
}

type RangeMap struct {
	m []RangeMapEntry
}

func NewRangeMap() *RangeMap {
	rm := RangeMap{
		m: []RangeMapEntry{},
	}
	return &rm
}

func (m *RangeMap) Add(k, v, r int) {
	t := v - k
	m.m = append(m.m, RangeMapEntry{Range{k, r}, t})
}

func (m *RangeMap) Get(k int) (int, bool) {
	for _, mm := range m.m {
		if mm.r.InRange(k) {
			return k + mm.t, true
		}
	}
	return 0, false
}

func (m *RangeMap) MapRanges(r []Range) []Range {
	ranges := []Range{}
	next := []Range{}
	remaining := r

	for _, mm := range m.m {
		for _, rr := range remaining {
			it := rr.Intersection(mm.r)
			if it == nil {
				next = append(next, rr)
				continue
			}
			ranges = append(ranges, Range{it[0].s + mm.t, it[0].r})
			next = append(next, it[1:]...)
		}
		remaining = next
		next = []Range{}
	}
	ranges = append(ranges, remaining...)

	return ranges
}

func parseMap(s string) *RangeMap {
	lines := strings.Split(s, "\n")
	r := regexp.MustCompile(`(.*) (.*) (.*)`)

	rm := NewRangeMap()
	for _, line := range lines[1:] {
		m := r.FindStringSubmatch(line)

		v, _ := strconv.Atoi(m[1])
		k, _ := strconv.Atoi(m[2])
		r, _ := strconv.Atoi(m[3])
		rm.Add(k, v, r)
	}
	return rm
}

func parseSeeds(s string) []int {
	seeds := []int{}
	for _, ss := range strings.Split(s, " ")[1:] {
		seed, _ := strconv.Atoi(ss)
		seeds = append(seeds, seed)
	}
	return seeds
}

type Solver struct {
	seeds                 []int
	seedToSoil            *RangeMap
	soilToFertilizer      *RangeMap
	fertilizerToWater     *RangeMap
	waterToLight          *RangeMap
	lightToTemperature    *RangeMap
	temperatureToHumidity *RangeMap
	humidityToLocation    *RangeMap
}

func NewSolver(input string) *Solver {
	groups := strings.Split(input, "\n\n")
	s := Solver{
		seeds:                 parseSeeds(groups[0]),
		seedToSoil:            parseMap(groups[1]),
		soilToFertilizer:      parseMap(groups[2]),
		fertilizerToWater:     parseMap(groups[3]),
		waterToLight:          parseMap(groups[4]),
		lightToTemperature:    parseMap(groups[5]),
		temperatureToHumidity: parseMap(groups[6]),
		humidityToLocation:    parseMap(groups[7]),
	}
	return &s
}

func (s *Solver) Part1() int {
	locations := []int{}
	for _, seed := range s.seeds {
		locations = append(locations, s.GetLocation(seed))
	}

	min := locations[0]
	for _, location := range locations[1:] {
		if location < min {
			min = location
		}
	}

	return min
}

func (s *Solver) Part2() int {
	ranges := []Range{}
	for i := 0; i < len(s.seeds); i += 2 {
		ranges = append(ranges, Range{s.seeds[i], s.seeds[i+1]})
	}
	ranges = s.GetLocationRange(ranges)

	min := ranges[0]
	for _, r := range ranges {
		if r.s < min.s {
			min = r
		}
	}

	return min.s
}

func (s *Solver) GetLocationRange(ranges []Range) []Range {
	ranges = s.seedToSoil.MapRanges(ranges)
	ranges = s.soilToFertilizer.MapRanges(ranges)
	ranges = s.fertilizerToWater.MapRanges(ranges)
	ranges = s.waterToLight.MapRanges(ranges)
	ranges = s.lightToTemperature.MapRanges(ranges)
	ranges = s.temperatureToHumidity.MapRanges(ranges)
	ranges = s.humidityToLocation.MapRanges(ranges)
	return ranges
}

func (s *Solver) GetLocation(seed int) int {
	soil, ok := s.seedToSoil.Get(seed)
	if !ok {
		soil = seed
	}

	fertilizer, ok := s.soilToFertilizer.Get(soil)
	if !ok {
		fertilizer = soil
	}

	water, ok := s.fertilizerToWater.Get(fertilizer)
	if !ok {
		water = fertilizer
	}

	light, ok := s.waterToLight.Get(water)
	if !ok {
		light = water
	}

	temperature, ok := s.lightToTemperature.Get(light)
	if !ok {
		temperature = light
	}

	humidity, ok := s.temperatureToHumidity.Get(temperature)
	if !ok {
		humidity = temperature
	}

	location, ok := s.humidityToLocation.Get(humidity)
	if !ok {
		location = humidity
	}

	return location
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)
	fmt.Println(s.Part1())
	fmt.Println(s.Part2())
}
