package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/pcjun97/advent-of-code/aoc"
)

func main() {
	input := aoc.ReadInput()

	sum := 0
	for _, line := range strings.Split(input, "\n") {
		s := NewSnafuFromString(line)
		sum += s.Int()
	}
	fmt.Println(sum)

	s := NewSnafuFromInt(sum)
	fmt.Println(s)
}

type Snafu []int

func NewSnafuFromString(s string) Snafu {
	snafu := Snafu{}
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '2':
			snafu = append(snafu, 2)
		case '1':
			snafu = append(snafu, 1)
		case '0':
			snafu = append(snafu, 0)
		case '-':
			snafu = append(snafu, -1)
		case '=':
			snafu = append(snafu, -2)
		}
	}
	return snafu
}

func NewSnafuFromInt(n int) Snafu {
	snafu := Snafu{}
	for n != 0 {
		r := n % 5
		if r <= 2 {
			snafu = append(snafu, r)
		} else {
			d := 5 - r
			snafu = append(snafu, -d)
			n += d
		}
		n /= 5
	}
	return snafu
}

func (s Snafu) Int() int {
	sum := 0
	for i, v := range s {
		sum += int(math.Pow(5, float64(i))) * v
	}
	return sum
}

func (snafu Snafu) String() string {
	s := ""
	for i := len(snafu) - 1; i >= 0; i-- {
		switch snafu[i] {
		case -1:
			s += "-"
		case -2:
			s += "="
		default:
			s += string('0' + snafu[i])
		}
	}
	return s
}
