package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pcjun97/advent-of-code/aoc"
)

type Range3D struct {
	x, y, z aoc.Range
}

func parseRange3D(s string) Range3D {
	r := regexp.MustCompile(`(.+),(.+),(.+)~(.+),(.+),(.+)`)
	m := r.FindStringSubmatch(s)
	x1, _ := strconv.Atoi(m[1])
	y1, _ := strconv.Atoi(m[2])
	z1, _ := strconv.Atoi(m[3])
	x2, _ := strconv.Atoi(m[4])
	y2, _ := strconv.Atoi(m[5])
	z2, _ := strconv.Atoi(m[6])

	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	if z1 > z2 {
		z1, z2 = z2, z1
	}

	b := Range3D{
		x: aoc.NewRange(x1, x2),
		y: aoc.NewRange(y1, y2),
		z: aoc.NewRange(z1, z2),
	}

	return b
}

func (r Range3D) Above() Range3D {
	z := r.z.Max
	return Range3D{r.x, r.y, aoc.NewRange(z+1, z+1)}
}

func (r Range3D) Below() Range3D {
	z := r.z.Min
	return Range3D{r.x, r.y, aoc.NewRange(z-1, z-1)}
}

func (r Range3D) Overlap(rr Range3D) bool {
	return r.x.Overlap(rr.x) && r.y.Overlap(rr.y) && r.z.Overlap(rr.z)
}

type Solver struct {
	bricks []Range3D
}

func NewSolver(input string) *Solver {
	bricks := []Range3D{}
	for _, line := range strings.Split(input, "\n") {
		bricks = append(bricks, parseRange3D(line))
	}
	s := Solver{bricks}
	return &s
}

func (s *Solver) Part1() int {
	bricks := dropBricks(s.bricks)

	supported := make(map[Range3D][]Range3D)
	supporting := make(map[Range3D][]Range3D)
	for _, brick := range bricks {
		supported[brick] = []Range3D{}
		supportedSet := make(map[Range3D]struct{})
		below := brick.Below()
		for _, b := range bricks {
			if b.Overlap(below) {
				supportedSet[b] = struct{}{}
			}
		}
		for b := range supportedSet {
			supported[brick] = append(supported[brick], b)
		}

		supporting[brick] = []Range3D{}
		supportingSet := make(map[Range3D]struct{})
		above := brick.Above()
		for _, b := range bricks {
			if b.Overlap(above) {
				supportingSet[b] = struct{}{}
			}
		}
		for b := range supportingSet {
			supporting[brick] = append(supporting[brick], b)
		}
	}

	count := 0
	for _, brick := range bricks {
		ok := true
		for _, b := range supporting[brick] {
			if len(supported[b]) <= 1 {
				ok = false
				break
			}
		}
		if ok {
			count += 1
		}
	}

	return count
}

func (s *Solver) Part2() int {
	bricks := dropBricks(s.bricks)

	supported := make(map[Range3D][]Range3D)
	for _, brick := range bricks {
		supported[brick] = []Range3D{}
		supportedSet := make(map[Range3D]struct{})
		below := brick.Below()
		for _, b := range bricks {
			if b.Overlap(below) {
				supportedSet[b] = struct{}{}
			}
		}
		for b := range supportedSet {
			supported[brick] = append(supported[brick], b)
		}
	}

	lenWithout := func(list []Range3D, without map[Range3D]struct{}) int {
		count := 0
		for _, b := range list {
			if _, ok := without[b]; !ok {
				count += 1
			}
		}
		return count
	}

	sum := 0
	for _, brick := range bricks {
		removed := make(map[Range3D]struct{})
		removed[brick] = struct{}{}

		change := true
		for change {
			change = false
			for _, b := range bricks {
				if b.z.Min == 1 {
					continue
				}

				if _, ok := removed[b]; ok {
					continue
				}

				if lenWithout(supported[b], removed) == 0 {
					removed[b] = struct{}{}
					change = true
				}
			}
		}

		sum += len(removed) - 1
	}

	return sum
}

func dropBricks(bricks []Range3D) []Range3D {
	result := make([]Range3D, len(bricks))
	copy(result, bricks)

	moved := true
	for moved {
		moved = false
		for i := range result {
			candrop := true
			for candrop {
				if result[i].z.Min == 1 {
					candrop = false
					break
				}

				below := result[i].Below()
				for _, b := range result {
					if b.Overlap(below) {
						candrop = false
						break
					}
				}

				if candrop {
					result[i].z.Max -= 1
					result[i].z.Min -= 1
					moved = true
				}
			}
		}
	}

	return result
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(s.Part2(), time.Since(start).String())
}
