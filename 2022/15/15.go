package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func main() {
	input := aoc.ReadInput()
	sensors := NewSensors(input)

	fmt.Println(sensors.EmptyCount(2000000))
	signal, err := sensors.UnknownPosition(Position{0, 0}, Position{4000000, 4000000})
	if err != nil {
		panic(err)
	}

	fmt.Println(signal.X*4000000 + signal.Y)
}

type Position struct {
	X, Y int
}

func (p Position) ManhattanDistance(other Position) int {
	dX := p.X - other.X
	if dX < 0 {
		dX = -dX
	}

	dY := p.Y - other.Y
	if dY < 0 {
		dY = -dY
	}

	return dX + dY
}

type Range struct {
	Start, End int
}

func NewRange(start int, end int) Range {
	if end < start {
		start, end = end, start
	}

	return Range{start, end}
}

func (r Range) Overlap(other Range) bool {
	return !(r.Start > other.End || other.Start > r.End)
}

func (r Range) Merge(other Range) (Range, error) {
	if !r.Overlap(other) {
		return Range{0, 0}, errors.New("ranges do not overlap")
	}

	start := r.Start
	if other.Start < start {
		start = other.Start
	}

	end := r.End
	if other.End > end {
		end = other.End
	}

	return Range{start, end}, nil
}

type Ranges []Range

func (ranges *Ranges) Add(r Range) {
	merge := []Range{}

	for _, rg := range *ranges {
		if r.Overlap(rg) {
			r, _ = r.Merge(rg)
		} else {
			merge = append(merge, rg)
		}
	}
	*ranges = append(merge, r)
}

type Sensor struct {
	Position
	Beacon Position
}

func NewSensor(p Position, beacon Position) *Sensor {
	s := Sensor{
		Position: p,
		Beacon:   beacon,
	}

	return &s
}

func (s *Sensor) XRange(y int) Ranges {
	d := s.ManhattanDistance(s.Beacon)
	if y > s.Y+d || y < s.Y-d {
		return nil
	}

	dY := s.Y - y
	if dY < 0 {
		dY = -dY
	}

	dX := d - dY

	r := Range{s.X - dX, s.X + dX}
	return Ranges{r}
}

func (s *Sensor) XRangeEmpty(y int) Ranges {
	r := s.XRange(y)
	if len(r) == 0 {
		return r
	}

	if s.Beacon.Y == y {
		if r[0].Start == r[0].End {
			return Ranges{}
		}

		if s.Beacon.X == r[0].Start {
			r[0].Start += 1
		} else {
			r[0].End -= 1
		}
	}

	if s.Y != y {
		return r
	}

	if s.X == r[0].Start {
		r[0].Start += 1
	}

	if s.X == r[0].End {
		r[0].End -= 1
	}

	r1 := Range{r[0].Start, s.X - 1}
	r2 := Range{s.X + 1, r[0].End}
	return Ranges{r1, r2}
}

type Sensors []*Sensor

func NewSensors(s string) Sensors {
	lines := strings.Split(s, "\n")
	sensors := make([]*Sensor, len(lines))

	for i, line := range lines {
		fields := strings.Split(line, " ")

		sxStr := fields[2]
		sx, err := strconv.Atoi(sxStr[2 : len(sxStr)-1])
		if err != nil {
			panic(err)
		}

		syStr := fields[3]
		sy, err := strconv.Atoi(syStr[2 : len(syStr)-1])
		if err != nil {
			panic(err)
		}

		bxStr := fields[8]
		bx, err := strconv.Atoi(bxStr[2 : len(bxStr)-1])
		if err != nil {
			panic(err)
		}

		by, err := strconv.Atoi(fields[9][2:])
		if err != nil {
			panic(err)
		}

		sensor := Position{sx, sy}
		beacon := Position{bx, by}
		sensors[i] = NewSensor(sensor, beacon)
	}

	return sensors
}

func (sensors Sensors) EmptyXRanges(y int) Ranges {
	ranges := Ranges{}
	for _, s := range sensors {
		for _, r := range s.XRangeEmpty(y) {
			ranges.Add(r)
		}
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].End < ranges[j].Start
	})

	return ranges
}

func (sensors Sensors) XRanges(y int) Ranges {
	ranges := Ranges{}
	for _, s := range sensors {
		for _, r := range s.XRange(y) {
			ranges.Add(r)
		}
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].End < ranges[j].Start
	})

	return ranges
}

func (sensors Sensors) EmptyCount(row int) int {
	sum := 0
	for _, r := range sensors.EmptyXRanges(row) {
		sum += r.End - r.Start + 1
	}

	return sum
}

func (sensors Sensors) UnknownPosition(p1 Position, p2 Position) (Position, error) {
	minX := p1.X
	if p2.X < minX {
		minX = p2.X
	}

	minY := p1.Y
	if p2.Y < minY {
		minY = p2.Y
	}

	maxX := p1.X
	if p2.X > maxX {
		maxX = p2.X
	}

	maxY := p1.Y
	if p2.Y > maxY {
		maxY = p2.Y
	}

	p := Position{minX, minY}

	for p.Y <= maxY {
		for _, r := range sensors.XRanges(p.Y) {
			if r.End < minX || r.Start > maxX {
				continue
			} else if p.X > maxX {
				break
			} else if p.X >= r.Start && p.X <= r.End {
				p.X = r.End + 1
			} else {
				return p, nil
			}
		}
		p.X = minX
		p.Y += 1
	}

	return Position{}, errors.New("signal not found")
}
