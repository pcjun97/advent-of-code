package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func main() {
	input := aoc.ReadInput()

	pairs := NewPairs(input)
	sum := 0
	for i, p := range pairs {
		if p.Compare() == 1 {
			sum += i + 1
		}
	}
	fmt.Println(sum)

	signal := NewSignal(input)
	fmt.Println(signal.DecoderKey())
}

type Signal struct {
	packets  []*Data
	dividers [2]*Data
}

func NewSignal(s string) Signal {
	d1 := NewData("[[2]]")
	d2 := NewData("[[6]]")

	packets := []*Data{&d1, &d2}

	for _, packet := range strings.Split(s, "\n") {
		if len(packet) == 0 {
			continue
		}
		d := NewData(packet)
		packets = append(packets, &d)
	}

	sort.Slice(packets, func(i, j int) bool {
		return packets[i].Compare(*packets[j]) == 1
	})

	signal := Signal{
		packets:  packets,
		dividers: [2]*Data{&d1, &d2},
	}

	return signal
}

func (s Signal) DecoderKey() int {
	key := 1
	for i, p := range s.packets {
		if p == s.dividers[0] || p == s.dividers[1] {
			key *= i + 1
		}
	}
	return key
}

type Pairs []Pair

func NewPairs(s string) Pairs {
	pairs := Pairs{}
	for _, pair := range strings.Split(s, "\n\n") {
		pairs = append(pairs, NewPair(pair))
	}
	return pairs
}

type Pair [2]Data

func NewPair(s string) Pair {
	var p Pair
	lines := strings.Split(s, "\n")
	p[0] = NewData(lines[0])
	p[1] = NewData(lines[1])
	return p
}

func (p Pair) Compare() int {
	return p[0].Compare(p[1])
}

type Data struct {
	Integer int
	List    []Data
}

func NewData(s string) Data {
	d, _ := ParseData(s, 0)
	return d
}

func ParseData(s string, i int) (Data, int) {
	d := Data{
		Integer: -1,
		List:    nil,
	}

	if s[i] != '[' {
		end := i
		for s[end] != ']' && s[end] != ',' {
			end++
		}

		val, err := strconv.Atoi(s[i:end])
		if err != nil {
			panic("failed to parse data: " + s)
		}
		d.Integer = val
		return d, end
	}

	l := []Data{}
	i++
	for s[i] != ']' {
		if s[i] != ',' && s[i] != ' ' {
			var e Data
			e, i = ParseData(s, i)
			l = append(l, e)
		} else {
			i++
		}
	}

	d.List = l
	return d, i + 1
}

func (d Data) Compare(o Data) int {
	if d.Integer >= 0 && o.Integer >= 0 {
		if d.Integer > o.Integer {
			return -1
		}
		if d.Integer < o.Integer {
			return 1
		}
		return 0
	}

	if d.Integer < 0 && o.Integer < 0 {
		min := len(d.List)
		if len(o.List) < min {
			min = len(o.List)
		}

		for i := 0; i < min; i++ {
			v := d.List[i].Compare(o.List[i])
			if v != 0 {
				return v
			}
		}

		if len(d.List) < len(o.List) {
			return 1
		}
		if len(d.List) > len(o.List) {
			return -1
		}
		return 0
	}

	if d.Integer >= 0 {
		e := Data{
			Integer: -1,
			List:    []Data{d},
		}
		return e.Compare(o)
	}

	if o.Integer >= 0 {
		e := Data{
			Integer: -1,
			List:    []Data{o},
		}
		return d.Compare(e)
	}

	return 0
}
