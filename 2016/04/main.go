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

type LetterCount struct {
	Letter rune
	Count  int
}

type Room struct {
	Name     string
	SectorID int
	Checksum string
}

func parseRoom(s string) Room {
	r := regexp.MustCompile(`^(.*)-(\d+)\[(.*)\]$`)
	m := r.FindStringSubmatch(s)

	name := m[1]
	checksum := m[3]

	id, err := strconv.Atoi(m[2])
	if err != nil {
		log.Fatal(err)
	}

	return Room{name, id, checksum}
}

func (r Room) IsReal() bool {
	m := make(map[rune]int)

	for _, l := range r.Name {
		if l == '-' {
			continue
		}

		if _, ok := m[l]; !ok {
			m[l] = 0
		}

		m[l]++
	}

	lc := []LetterCount{}
	for l, c := range m {
		lc = append(lc, LetterCount{l, c})
	}

	sort.Slice(lc, func(i, j int) bool {
		if lc[i].Count == lc[j].Count {
			return lc[i].Letter < lc[j].Letter
		}

		return lc[i].Count > lc[j].Count
	})

	checksum := ""
	for _, l := range lc {
		checksum = fmt.Sprintf("%s%c", checksum, l.Letter)
		if len(checksum) >= 5 {
			break
		}
	}

	return r.Checksum == checksum
}

func (r Room) Decrypt() string {
	d := ""

	for _, rr := range r.Name {
		c := ' '

		if rr != '-' {
			c = ((rr - 'a' + rune(r.SectorID)) % 26) + 'a'
		}

		d = fmt.Sprintf("%s%c", d, c)
	}

	return d
}

type Solver struct {
	Rooms []Room
}

func NewSolver(input string) *Solver {
	rooms := []Room{}
	for _, line := range strings.Split(input, "\n") {
		rooms = append(rooms, parseRoom(line))
	}

	s := Solver{rooms}
	return &s
}

func (s *Solver) Part1() int {
	sum := 0
	for _, room := range s.Rooms {
		if room.IsReal() {
			sum += room.SectorID
		}
	}

	return sum
}

type Part2Output struct {
	Name     string
	SectorID int
}

func (s *Solver) Part2() []Part2Output {
	output := []Part2Output{}

	for _, r := range s.Rooms {
		if !r.IsReal() {
			continue
		}

		o := Part2Output{r.Decrypt(), r.SectorID}
		output = append(output, o)
	}

	return output
}

func main() {
	input := aoc.ReadInput()
	s := NewSolver(input)

	start := time.Now()
	fmt.Println(s.Part1(), time.Since(start).String())

	start = time.Now()
	fmt.Println(time.Since(start).String())

	for _, r := range s.Part2() {
		fmt.Printf("%s (%d)\n", r.Name, r.SectorID)
	}
}
